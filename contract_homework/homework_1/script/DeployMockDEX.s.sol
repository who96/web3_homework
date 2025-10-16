// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/MockDEX.sol";
import "../src/MemeToken.sol";

/**
 * @title DeployMockDEX
 * @notice 在Sepolia部署MockDEX并添加流动性
 * @dev 使用方法:
 *   forge script script/DeployMockDEX.s.sol --rpc-url $SEPOLIA_RPC_URL --broadcast -vvv
 */
contract DeployMockDEX is Script {
    // Sepolia上已部署的MemeToken地址
    address constant MEME_TOKEN = 0x61a33158B1541AD0fc87DF41075ac6A40CC52498;

    // 流动性参数
    uint256 constant TOKEN_AMOUNT = 210_000 * 10 ** 18; // 210,000 FK
    uint256 constant ETH_AMOUNT = 0.1 ether; // 0.1 ETH

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);

        MemeToken token = MemeToken(MEME_TOKEN);

        console.log("==========================================");
        console.log("Deploying MockDEX on Sepolia");
        console.log("==========================================");
        console.log("Deployer:", deployer);
        console.log("Token:", MEME_TOKEN);
        console.log("Token balance:", token.balanceOf(deployer) / 10 ** 18, "FK");
        console.log("ETH balance:", deployer.balance / 1 ether, "ETH");
        console.log("==========================================");

        require(token.balanceOf(deployer) >= TOKEN_AMOUNT, "Insufficient token balance");
        require(deployer.balance >= ETH_AMOUNT, "Insufficient ETH balance");

        vm.startBroadcast(deployerPrivateKey);

        // 1. 部署MockDEX
        console.log("\nStep 1: Deploying MockDEX...");
        MockDEX dex = new MockDEX(MEME_TOKEN);
        console.log("MockDEX deployed at:", address(dex));

        // 2. 授权DEX使用代币
        console.log("\nStep 2: Approving DEX...");
        token.approve(address(dex), TOKEN_AMOUNT);
        console.log("Approved!");

        // 3. 添加流动性
        console.log("\nStep 3: Adding liquidity...");
        uint256 shares = dex.addLiquidity{value: ETH_AMOUNT}(TOKEN_AMOUNT);
        console.log("Liquidity added!");
        console.log("- Token amount:", TOKEN_AMOUNT / 10 ** 18, "FK");
        console.log("- ETH amount:", ETH_AMOUNT / 1 ether, "ETH");
        console.log("- LP shares:", shares);

        // 4. 查询储备
        console.log("\nReserves:");
        console.log("- Token reserve:", dex.tokenReserve() / 10 ** 18, "FK");
        console.log("- ETH reserve:", dex.ethReserve() / 1 ether, "ETH");

        console.log("\n==========================================");
        console.log("IMPORTANT: Add DEX to whitelist!");
        console.log("==========================================");
        console.log("Run the following command as owner:");
        console.log("");
        console.log("cast send", MEME_TOKEN);
        console.log("  'setTaxExempt(address,bool)'");
        console.log("  ", address(dex), "true");
        console.log("  --rpc-url $SEPOLIA_RPC_URL --private-key $PRIVATE_KEY");
        console.log("");
        console.log("cast send", MEME_TOKEN);
        console.log("  'setLimitExempt(address,bool)'");
        console.log("  ", address(dex), "true");
        console.log("  --rpc-url $SEPOLIA_RPC_URL --private-key $PRIVATE_KEY");
        console.log("==========================================");

        vm.stopBroadcast();
    }
}
