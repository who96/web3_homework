// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";
import {SimpleStorage} from "../src/SimpleStorage.sol";

contract SimpleStorageTest is Test {
    SimpleStorage public storageContract;

    function setUp() public {
        storageContract = new SimpleStorage(42);
    }

    function testInitialValue() public view {
        assertEq(storageContract.get(), 42);
    }

    function testSet() public {
        storageContract.set(100);
        assertEq(storageContract.get(), 100);
    }

    function testIncrement() public {
        storageContract.increment();
        assertEq(storageContract.get(), 43);
    }

    function testEventEmission() public {
        vm.expectEmit(true, true, false, true);
        emit SimpleStorage.ValueChanged(42, 100, address(this));
        storageContract.set(100);
    }
}