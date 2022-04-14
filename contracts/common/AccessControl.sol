// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/access/AccessControlEnumerable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";

contract KfiveAccessControl is AccessControlEnumerable, Ownable, Pausable {
    constructor() {
        _setupRole(DEFAULT_ADMIN_ROLE, _msgSender());
        _setupRole(ADMIN_ROLE, _msgSender());
        _setRoleAdmin(PAUSER_ROLE, ADMIN_ROLE);
    }

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    function transferOwnership(address newOwner) public override onlyOwner {
        _grantRole(ADMIN_ROLE, newOwner);
        _grantRole(DEFAULT_ADMIN_ROLE, newOwner);

        _revokeRole(ADMIN_ROLE, owner());
        _revokeRole(DEFAULT_ADMIN_ROLE, owner());

        super.transferOwnership(newOwner);
    }

    function _grantRole(bytes32 role, address account) internal override {
        super._grantRole(role, account);
        if (role == ADMIN_ROLE || role == DEFAULT_ADMIN_ROLE) {
            super._grantRole(PAUSER_ROLE, account);
            if (role == DEFAULT_ADMIN_ROLE) {
                super._grantRole(ADMIN_ROLE, account);
            }
        }
    }

    function _revokeRole(bytes32 role, address account) internal override {
        super._revokeRole(role, account);
        if (role == ADMIN_ROLE || role == DEFAULT_ADMIN_ROLE) {
            super._revokeRole(PAUSER_ROLE, account);
            if (role == DEFAULT_ADMIN_ROLE) {
                super._grantRole(ADMIN_ROLE, account);
            }
        }
    }

    /**
     * @dev Pauses all actions.
     *
     * Requirements:
     *
     * - the caller must have the `PAUSER_ROLE`.
     */
    function pause() public onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /**
     * @dev Unpauses all actions.
     *
     * Requirements:
     *
     * - the caller must have the `PAUSER_ROLE`.
     */
    function unpause() public onlyRole(PAUSER_ROLE) {
        _unpause();
    }
}
