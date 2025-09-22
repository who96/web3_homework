// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

/**
 * @title AuctionProxyAdmin
 * @notice Admin contract for managing AuctionProxy upgrades
 * @dev Uses OpenZeppelin's ProxyAdmin for secure upgrade management
 */
contract AuctionProxyAdmin is ProxyAdmin {
    /**
     * @dev Initialize the proxy admin with an initial owner
     * @param initialOwner The address that will own this ProxyAdmin
     */
    constructor(address initialOwner) ProxyAdmin(initialOwner) {}

    /**
     * @dev Upgrade the implementation of a proxy to a new version
     * @param proxy The proxy to upgrade
     * @param implementation The new implementation contract
     */
    function upgradeProxy(
        address proxy,
        address implementation
    ) external onlyOwner {
        upgradeAndCall(
            ITransparentUpgradeableProxy(proxy),
            implementation,
            ""
        );
    }

    /**
     * @dev Upgrade the implementation and call a function on the new implementation
     * @param proxy The proxy to upgrade
     * @param implementation The new implementation contract
     * @param data The data to call on the new implementation
     */
    function upgradeProxyAndCall(
        address proxy,
        address implementation,
        bytes calldata data
    ) external onlyOwner {
        upgradeAndCall(
            ITransparentUpgradeableProxy(proxy),
            implementation,
            data
        );
    }
}