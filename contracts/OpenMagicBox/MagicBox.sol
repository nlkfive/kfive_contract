// SPDX-License-Identifier: MIT

pragma solidity 0.8.4;

import "../common/AccessControl.sol";
import "../BEP20/IBEP20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "../NFT/IKfiveNFT.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "@chainlink/contracts/src/v0.8/VRFConsumerBaseV2.sol";
import "@chainlink/contracts/src/v0.8/interfaces/VRFCoordinatorV2Interface.sol";
import "@openzeppelin/contracts/utils/Address.sol";

contract MagicBox is 
    VRFConsumerBaseV2,
    KfiveAccessControl, 
    IERC721Receiver
{

    using SafeMath for uint256;
    using Address for address;

    uint256 private constant EMPTY = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF;

    enum BoxType{ DIAMON, GOLD, SILVER, BROZEN }
    struct BoxInfo {
        BoxType boxType;
        address receiver;
        uint256 randomness;
    }

    IBEP20 private _acceptToken;
    uint256 _endedAt;
    mapping (BoxType => uint256) _price;
    mapping (BoxType => uint256[]) _rewardList;
    mapping (uint256 => BoxInfo) _vrfRandomMap;
    mapping (uint256 => uint256) _receivedMap;
    mapping (address => uint256) _listRequested;
    IKfiveNFT private _magicBoxReward;

    VRFCoordinatorV2Interface private _vrfCoordinator;

    // The contract will contain multiple objects. 
    // Each oracle job has a unique key hash that identifies the tasks that it should perform. 
    // The contract will store the Key Hash that identifies Chainlink VRF and the fee amount to use in the request.
    bytes32 private s_keyHash;
    uint64 private s_subscriptionId;

    // The default is 3, but you can set this higher.
    uint16 private constant requestConfirmations = 3;
    // Depends on the number of requested values that you want sent to the
    // fulfillRandomWords() function. Storing each word costs about 20,000 gas,
    // so 100,000 is a safe default for this example contract. Test and adjust
    // this limit based on the network that you select, the size of the request,
    // and the processing of the callback request in the fulfillRandomWords()
    // function.
    uint32 private constant callbackGasLimit = 100000;

    // For this example, retrieve 2 random values in one request.
    // Cannot exceed VRFCoordinatorV2.MAX_NUM_WORDS.
    uint32 private constant numWords =  1;

    error InvalidContract(); // 0x6eefed20
    error NoReward(); // 0x6e992686
    error OpenFailed(); // 0xffd586b9
    error TooLate(uint256 time);
    error NotOwner(); // 0x30cd7471
    error BoxRequested(); // 0x8b46d2c6
    error RandomnessNotGenerated(); // 0xd93e896f

    event RequestBoxSuccessful(BoxType boxType, address sender, uint256 requestId);
    event RewardAdded(BoxType boxType, uint256 rewardIndex, uint256 nftRewardId);
    event RewardTransfered(address receiver, uint256 rewardIndex, uint256 tokenId);
    event BoxReceived(uint256 requestId, uint256 randomness);

    constructor(
        address acceptToken, // LF5
        address vrfCoordinator,
        address magicBoxReward,
        uint256 endedAt,
        uint64 subscriptionId,
        bytes32 keyHash
    ) VRFConsumerBaseV2(vrfCoordinator) {
        if(!acceptToken.isContract()) revert InvalidContract();
        if(!vrfCoordinator.isContract()) revert InvalidContract();
        if(!magicBoxReward.isContract()) revert InvalidContract();

        if (block.timestamp >= endedAt) revert TooLate(endedAt);

        _acceptToken = IBEP20(acceptToken);
        _price[BoxType.DIAMON] = 40;
        _price[BoxType.GOLD] = 15;
        _price[BoxType.SILVER] = 5;
        _price[BoxType.BROZEN] = 1;
        _endedAt = endedAt;

        _magicBoxReward = IKfiveNFT(magicBoxReward);
        _vrfCoordinator = VRFCoordinatorV2Interface(vrfCoordinator);
        s_keyHash = keyHash;
        s_subscriptionId = subscriptionId;
    }

    /** 
     * @dev Get random box
     */
    function requestBox(
        BoxType boxType
    ) 
        external
        whenNotPaused
        onlyBefore
        returns (uint256 requestId)
    {
        address sender = _msgSender();
        if (_listRequested[sender] != 0){
            revert BoxRequested();
        }

        requestId = _vrfCoordinator.requestRandomWords(
            s_keyHash,
            s_subscriptionId,
            requestConfirmations,
            callbackGasLimit,
            numWords
        );
        _listRequested[sender] = requestId;
        // storing requestId
        BoxInfo storage boxInfo = _vrfRandomMap[requestId];
        {
            boxInfo.boxType = boxType;
            boxInfo.receiver = sender;
            boxInfo.randomness = EMPTY;
        }

        emit RequestBoxSuccessful(boxType, sender, requestId);
    }

    /** 
     * @dev Handle on result random
     */
    function fulfillRandomWords(uint256 requestId, uint256[] memory randomWords) internal override 
    {
        uint256 randomness = randomWords[0];
        _vrfRandomMap[requestId].randomness = randomness;
        emit BoxReceived(requestId, randomness);
    }
    
    /** 
     * @dev Handle on result random
     */
    function openBox(uint256 requestId) 
        external
        whenNotPaused
        onlyBefore
    {
        address sender = _msgSender();
        BoxInfo memory boxInfo = _vrfRandomMap[requestId];
        delete _vrfRandomMap[requestId];
        delete _listRequested[sender];

        if (boxInfo.receiver != sender){
            revert NotOwner();
        }
        if (boxInfo.randomness == EMPTY){
            revert RandomnessNotGenerated();
        }
        BoxType boxType = boxInfo.boxType;
        uint256 totalRemainReward = _rewardList[boxType].length;
        if (totalRemainReward == 0) {
            revert NoReward();
        }

        if (!_acceptToken.transferFrom(sender, address(this), _price[boxType])) revert OpenFailed();

        uint256 rewardIndex = boxInfo.randomness.mod(totalRemainReward);
        _getReward(boxType, sender, rewardIndex);
    }

    /**
     * @dev Add reward.
     */
    function addReward(BoxType boxType, uint256 nftRewardId, string memory tokenURI) 
        external 
        whenNotPaused
        onlyRole(ADMIN_ROLE)
    {
        _magicBoxReward.mint(
            address(this),
            nftRewardId,
            tokenURI
        );
        _rewardList[boxType].push(nftRewardId);

        emit RewardAdded(boxType, _rewardList[boxType].length, nftRewardId);
    }

    modifier onlyBefore() {
        if (block.timestamp >= _endedAt) revert TooLate(_endedAt);
        _;
    }

    function remainReward(BoxType boxType) external view returns (uint256) {
        return _rewardList[boxType].length;
    }

    function listReward(BoxType boxType) external view returns (uint256[] memory) {
        return _rewardList[boxType];
    }

    function getRequestId(address sender) external view returns (uint256) {
        return _listRequested[sender];
    }

    function getEndedAt() external view returns (uint256) {
        return _endedAt;
    }

    function _getReward(
        BoxType boxType, 
        address receiver, 
        uint256 rewardIndex
    ) internal {
        uint256 nftId = _rewardList[boxType][rewardIndex];
        _rewardList[boxType] = remove(rewardIndex, boxType);

        _magicBoxReward.safeTransferFrom(
            address(this),
            receiver,
            nftId
        );

        emit RewardTransfered(receiver, rewardIndex, nftId);
    }

    function onERC721Received(
        address,
        address,
        uint256,
        bytes calldata
    ) external pure override returns (bytes4) {
        return
            bytes4(
                keccak256("onERC721Received(address,address,uint256,bytes)")
            );
    }

    function remove(uint256 index, BoxType boxType) internal returns(uint256[] memory) {
        uint256[] storage rewardList = _rewardList[boxType];
        if (index >= rewardList.length) return rewardList;

        for (uint256 i = index; i< rewardList.length-1; i++){
            rewardList[i] = rewardList[i+1];
        }
        rewardList.pop();
        return rewardList;
    }
}