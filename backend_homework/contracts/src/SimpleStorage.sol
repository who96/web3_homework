// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

contract SimpleStorage {
    uint256 private storedValue;

    event ValueChanged(uint256 oldValue, uint256 newValue, address indexed changer);

    constructor(uint256 _initialValue) {
        storedValue = _initialValue;
    }

    function set(uint256 _value) public {
        uint256 oldValue = storedValue;
        storedValue = _value;
        emit ValueChanged(oldValue, _value, msg.sender);
    }

    function get() public view returns (uint256) {
        return storedValue;
    }

    function increment() public {
        uint256 oldValue = storedValue;
        storedValue++;
        emit ValueChanged(oldValue, storedValue, msg.sender);
    }
}