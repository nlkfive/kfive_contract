// SPDX-License-Identifier: MIT

pragma solidity 0.8.4;

import "../common/AccessControl.sol";
import "../BEP20/IBEP20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/IERC721Enumerable.sol";
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

    enum BoxType{ DIAMON, GOLD, SILVER, BROZEN }
    struct DrawInfo {
        BoxType boxType;
        address receiver;
    }

    IBEP20 private _acceptToken;
    uint256 _endedAt;
    mapping (BoxType => uint256) _price;
    mapping (BoxType => uint256) _rewardCount;
    mapping (BoxType => uint256[]) _rewardList;
    mapping (uint256 => DrawInfo) _vrfRandomMap;
    IERC721Enumerable private _magicBoxReward;

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

    error InvalidContract();
    error NoReward();
    error DrawFailed();
    error TooLate(uint256 time);

    event DrawSuccessful(BoxType boxType, address sender, uint256 requestId);
    event RewardAdded(BoxType boxType, uint256 rewardIndex, uint256 nftRewardId);
    event RewardTransfered(address receiver, uint256 requestId, uint256 randomness, uint256 nftId);

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
        _price[BoxType.DIAMON] = 20;
        _price[BoxType.GOLD] = 10;
        _price[BoxType.SILVER] = 5;
        _price[BoxType.BROZEN] = 0;
        _endedAt = endedAt;

        _magicBoxReward = IERC721Enumerable(magicBoxReward);
        _vrfCoordinator = VRFCoordinatorV2Interface(vrfCoordinator);
        s_keyHash = keyHash;
        s_subscriptionId = subscriptionId;
    }

    /** 
     * @dev Draw diamon box
     */
    function draw(
        BoxType boxType
    ) 
        external
        whenNotPaused
        onlyBefore
        returns (uint256 requestId)
    {
        address sender = _msgSender();

        if (_rewardCount[boxType] == 0) {
            revert NoReward();
        }

        if (!_acceptToken.transferFrom(sender, address(this), _price[boxType])) revert DrawFailed();

        requestId = _vrfCoordinator.requestRandomWords(
            s_keyHash,
            s_subscriptionId,
            requestConfirmations,
            callbackGasLimit,
            numWords
        );
        // storing requestId
        DrawInfo storage drawInfo = _vrfRandomMap[requestId];
        {
            drawInfo.boxType = boxType;
            drawInfo.receiver = sender;
        }

        emit DrawSuccessful(boxType, sender, requestId);
    }

    /** 
     * @dev Handle on result random
     */
    function fulfillRandomWords(uint256 requestId, uint256[] memory randomWords) internal override 
    {
        BoxType boxType = _vrfRandomMap[requestId].boxType;
        if (_rewardCount[boxType] == 0) {
            revert NoReward();
        }
        address receiver = _vrfRandomMap[requestId].receiver;
        uint256 randomness = randomWords[0];
        uint256 totalRemainReward = _rewardCount[boxType];

        uint256 nftId = _rewardList[boxType][randomness % totalRemainReward];

        delete _rewardList[boxType][randomness % totalRemainReward];
        _magicBoxReward.safeTransferFrom(
            address(this),
            receiver,
            nftId
        );

        emit RewardTransfered(receiver, requestId, randomness, nftId);
    }

    /**
     * @dev Add reward.
     */
    function addReward(BoxType boxType, uint256 nftRewardId) 
        external 
        whenNotPaused
        onlyRole(ADMIN_ROLE)
    {
        uint256 rewardIndex = _rewardCount[boxType];

        _magicBoxReward.safeTransferFrom(
            _msgSender(),
            address(this),
            nftRewardId
        );
        _rewardCount[boxType] = rewardIndex.add(1);
        _rewardList[boxType].push(nftRewardId);

        emit RewardAdded(boxType, rewardIndex, nftRewardId);
    }

    modifier onlyBefore() {
        if (block.timestamp >= _endedAt) revert TooLate(_endedAt);
        _;
    }

    function remainReward(BoxType boxType) external view returns (uint256) {
        return _rewardCount[boxType];
    }

    function getEndedAt() external view returns (uint256) {
        return _endedAt;
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
}