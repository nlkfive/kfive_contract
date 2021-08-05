// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./Configurable.sol";

contract TokenAdmin is Configurable {
    address private superAdmin;
    mapping(address => AdminInfo) public admin;

    struct AdminInfo {
        bool status;
        uint256 maxIssuingTokenPerTime;
        uint256 maxTotalIssuingToken;
        uint256 remainingIssuingToken;
    }

    constructor() {
        superAdmin = msg.sender;
        addAdmin(msg.sender, cap, cap);
    }

    modifier onlyAdmin() {
        require(admin[msg.sender].status, "not admin");
        _;
    }

    modifier onlySuperAdmin() {
        require(msg.sender == superAdmin, "not super admin");
        _;
    }

    function addAdmin(
        address a,
        uint256 maxIssuingTokenPerTime,
        uint256 maxTotalIssuingToken
    ) public onlySuperAdmin {
        require(a != address(0), "address is nil");
        require(
            maxIssuingTokenPerTime <= maxTotalIssuingToken,
            "maxIssuingTokenPerTime must be less than or equal to maxTotalIssuingToken"
        );
        require(
            maxTotalIssuingToken <= cap,
            "maxTotalIssuingToken must be less than or equal to cap value"
        );

        AdminInfo memory i;
        i.status = true;
        i.maxIssuingTokenPerTime = maxIssuingTokenPerTime;
        i.maxTotalIssuingToken = maxTotalIssuingToken;
        i.remainingIssuingToken = maxTotalIssuingToken;

        admin[a] = i;
    }

    function removeAdmin(address a) public onlySuperAdmin {
        require(a != address(0), "address is nil");
        admin[a].status = false;
    }
}
