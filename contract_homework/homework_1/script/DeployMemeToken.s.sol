// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/MemeToken.sol";

/**
 * @title DeployMemeToken
 * @notice 部署MemeToken合约的脚本
 * @dev 使用方法:
 *   本地测试: forge script script/DeployMemeToken.s.sol --rpc-url http://localhost:8545 --broadcast
 *   测试网: forge script script/DeployMemeToken.s.sol --rpc-url $SEPOLIA_RPC_URL --broadcast --verify
 *   主网: forge script script/DeployMemeToken.s.sol --rpc-url $MAINNET_RPC_URL --broadcast --verify
 */
contract DeployMemeToken is Script {
    // 部署参数
    uint256 constant INITIAL_TAX_RATE = 300; // 3%

    function run() external returns (MemeToken) {
        // 从环境变量读取部署者私钥和WALLET1地址
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address wallet1 = vm.envAddress("WALLET1");

        require(wallet1 != address(0), "WALLET1 not set in .env");

        address deployer = vm.addr(deployerPrivateKey);

        console.log("===========================================");
        console.log("Deploying MemeToken...");
        console.log("===========================================");
        console.log("Deployer:", deployer);
        console.log("Tax Recipient (WALLET1):", wallet1);
        console.log("Initial Tax Rate:", INITIAL_TAX_RATE, "basis points (3%)");
        console.log("===========================================");

        vm.startBroadcast(deployerPrivateKey);

        // 部署合约
        // 参数: initialOwner, taxRecipient, taxRate
        MemeToken token = new MemeToken(deployer, wallet1, INITIAL_TAX_RATE);

        console.log("Token deployed at:", address(token));
        console.log("Token name:", token.name());
        console.log("Token symbol:", token.symbol());
        console.log("Total supply:", token.totalSupply() / 10 ** 18, "FK");
        console.log("Max transaction amount:", token.maxTransactionAmount() / 10 ** 18, "FK");
        console.log("===========================================");

        // 代币分配说明
        console.log("Token Distribution:");
        console.log("- Deployer (owner):", token.balanceOf(deployer) / 10 ** 18, "FK (100%)");
        console.log("");
        console.log("Next steps:");
        console.log("1. Transfer 210,000 FK to a temp address for liquidity");
        console.log("2. Add liquidity on Uniswap V2: 210,000 FK + 0.1 ETH");
        console.log("3. Add the Uniswap Pair address to tax/limit whitelist");
        console.log("4. Transfer remaining 20,790,000 FK to WALLET1");
        console.log("===========================================");

        vm.stopBroadcast();

        // 验证部署
        require(token.owner() == deployer, "Owner not set correctly");
        require(token.taxRecipient() == wallet1, "Tax recipient not set correctly");
        require(token.taxRate() == INITIAL_TAX_RATE, "Tax rate not set correctly");
        require(token.isExemptFromTax(deployer), "Deployer not exempt from tax");
        require(token.isExemptFromLimit(deployer), "Deployer not exempt from limit");

        console.log("Deployment verification passed!");

        return token;
    }
}
