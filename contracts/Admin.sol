// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

contract Admin {
    address private superAdmin;
    mapping(address => bool) public admin;

    constructor() {
        superAdmin = msg.sender;
        addAdmin(msg.sender);
    }

    modifier onlyAdmin() {
        require(admin[msg.sender], "not admin");
        _;
    }

    modifier onlySuperAdmin() {
        require(msg.sender == superAdmin, "not super admin");
        _;
    }

    function addAdmin(address a) public onlySuperAdmin {
        require(a != address(0), "address is nil");
        admin[a] = true;
    }

    function removeAdmin(address a) public onlySuperAdmin {
        require(a != address(0), "address is nil");
        admin[a] = false;
    }
}
