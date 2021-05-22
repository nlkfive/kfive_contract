// SPDX-License-Identifier: MIT

pragma solidity 0.8.4;

import "./Admin.sol";
import "./BEP20/IBEP20.sol";

contract TokenRegistry is Admin {
    struct TokenMetadata {
        address token;
        string symbol;
        string name;
        uint8 decimals;
        bool valid;
    }

    mapping(address => TokenMetadata) public tokenMapping;

    // events to track onchain-offchain relationships
    event __addNewToken(bytes32 offchain);

    /**
     * @dev function to add a new token
     * @param token the address of IBEP20 token
     * @param symbol the symbol of IBEP20 token
     * @param name the name of IBEP20 token
     * @param decimals the decimals of IBEP20 token
     */
    function addNewToken(
        address token,
        string memory symbol,
        string memory name,
        uint8 decimals,
        bytes32 offchain
    ) public onlyAdmin {
        TokenMetadata memory tokenData =
            TokenMetadata({
                token: token,
                symbol: symbol,
                name: name,
                decimals: decimals,
                valid: true
            });

        tokenMapping[token] = tokenData;
        emit __addNewToken(offchain);
    }

    /**
     * @dev function to transfer IBEP20 token
     * @param token the address of IBEP20 token
     * @param from the address of sender
     * @param to the address of recipient
     * @param amount the amount that sender sends to recipient
     */
    function transferToken(
        address token,
        address from,
        address to,
        uint256 amount
    ) public {
        IBEP20(token).transferFrom(from, to, amount);
    }

    /**
     * @dev function to check IBEP20 token exist or not
     * @param token the address of IBEP20 token
     */
    function tokenIsExisted(address token) public view returns (bool) {
        TokenMetadata storage tokenMetadata = tokenMapping[token];
        return tokenMetadata.valid;
    }

    /**
     * @dev function to get balance of a wallet address
     * @param token the address of IBEP20 token
     * @param addr the wallet address
     */
    function getBalanceOf(address token, address addr)
        public
        view
        returns (uint256)
    {
        return IBEP20(token).balanceOf(addr);
    }

    /**
     * @dev function to get IBEP20 token information
     * @param token the address of IBEP20 token
     */
    function getTokenByAddr(address token)
        public
        view
        returns (
            string memory,
            string memory,
            uint8
        )
    {
        TokenMetadata storage tokenData = tokenMapping[token];
        return (tokenData.symbol, tokenData.name, tokenData.decimals);
    }
}
