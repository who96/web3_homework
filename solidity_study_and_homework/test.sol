// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;


contract AccessControl {

    event RoleGranted(bytes32 indexed role, address indexed account);
    event RoleRevoked(bytes32 indexed role, address indexed account);

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant USER_ROLE = keccak256("USER_ROLE");


    constructor() {
        _grantRole(ADMIN_ROLE, msg.sender);
    }

    modifier onlyRole(bytes32 role) {
        require(roles[role][msg.sender], "Not authorized");
        _;
    }

   mapping(bytes32=>mapping(address=>bool)) public roles;

   function _grantRole(bytes32 role, address account) internal {
      roles[role][account] = true;
      emit RoleGranted(role, account);
   }

   function _revokeRole(bytes32 role, address account) internal {
      roles[role][account] = false;
      emit RoleRevoked(role, account);
   }
   
   function grantRole(bytes32 role, address account) public onlyRole(ADMIN_ROLE) {
      _grantRole(role, account);
   }

   function revokeRole(bytes32 role, address account) public onlyRole(ADMIN_ROLE) {
      _revokeRole(role, account);
   }
}