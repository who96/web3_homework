// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../contracts/getMoney.sol";

contract DeployBeggar is Script {
    function run() external {
        // 获取私钥 (支持有无0x前缀)
        string memory privateKeyStr = vm.envString("PRIVATE_KEY");
        uint256 deployerPrivateKey;

        // 检查是否有0x前缀
        if (bytes(privateKeyStr).length > 2 &&
            bytes(privateKeyStr)[0] == 0x30 &&
            bytes(privateKeyStr)[1] == 0x78) {
            deployerPrivateKey = vm.parseUint(privateKeyStr);
        } else {
            deployerPrivateKey = vm.parseUint(string.concat("0x", privateKeyStr));
        }

        address deployer = vm.addr(deployerPrivateKey);
        console.log("Deploying BeggingContract with account:", deployer);
        console.log("Account balance:", deployer.balance);

        vm.startBroadcast(deployerPrivateKey);

        BeggingContract beggar = new BeggingContract();
        console.log("BeggingContract deployed to:", address(beggar));
        console.log("Owner:", beggar.owner());

        // 验证部署
        require(beggar.owner() == deployer, "Owner verification failed");
        console.log("Deployment verification: SUCCESS");

        vm.stopBroadcast();

        // 显示部署后的信息
        console.log("\n=== Deployment Summary ===");
        console.log("Contract Address:", address(beggar));
        console.log("Owner Address:", deployer);
        console.log("Network: Sepolia Testnet");
        console.log("Working Hours: 9:00-18:00 UTC+8");
        console.log("\n=== Next Steps ===");
        console.log("1. Verify contract on Etherscan");
        console.log("2. Test donate() function during working hours");
        console.log("3. Check ranking with getTopDonors()");
    }
}