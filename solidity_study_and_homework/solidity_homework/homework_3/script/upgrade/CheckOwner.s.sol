// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract CheckOwner is Script {
    address constant AUCTION_PROXY = 0x5CAfBDf31623F3bfdc9Eefcdd6999719D67f6AbB;

    function run() external view {
        // Get the actual ProxyAdmin created internally by TransparentUpgradeableProxy
        bytes32 adminSlot = 0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;
        address actualProxyAdmin = address(uint160(uint256(vm.load(AUCTION_PROXY, adminSlot))));

        console.log("ProxyAdmin address:", actualProxyAdmin);

        ProxyAdmin proxyAdmin = ProxyAdmin(actualProxyAdmin);
        address owner = proxyAdmin.owner();

        console.log("ProxyAdmin owner:", owner);
        console.log("Current script account:", msg.sender);
        console.log("Are they the same?", owner == msg.sender);
    }
}