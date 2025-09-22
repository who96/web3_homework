// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

/**
 * @title AuctionProxy
 * @notice Transparent proxy for the SimpleAuction contract
 * @dev Uses OpenZeppelin's TransparentUpgradeableProxy for secure upgrades
 */
contract AuctionProxy is TransparentUpgradeableProxy {
    /**
     * @dev Initialize the proxy with implementation and admin
     * @param implementation The initial implementation contract address
     * @param admin The admin address that can upgrade the proxy
     * @param data The initialization data to call on the implementation
     */
    constructor(
        address implementation,
        address admin,
        bytes memory data
    ) TransparentUpgradeableProxy(implementation, admin, data) {}
}