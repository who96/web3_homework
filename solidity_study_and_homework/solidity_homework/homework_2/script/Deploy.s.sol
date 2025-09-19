// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../contracts/SimpleNFT.sol";

contract DeployScript is Script {
    function run() public {
        // 从环境变量读取私钥（自动处理 0x 前缀）
        string memory privateKeyStr = vm.envString("PRIVATE_KEY");
        uint256 deployerPrivateKey;

        // 检查是否有 0x 前缀
        if (bytes(privateKeyStr).length > 2 &&
            bytes(privateKeyStr)[0] == 0x30 &&
            bytes(privateKeyStr)[1] == 0x78) {
            // 有 0x 前缀，直接解析
            deployerPrivateKey = vm.parseUint(privateKeyStr);
        } else {
            // 没有 0x 前缀，添加后解析
            deployerPrivateKey = vm.parseUint(string.concat("0x", privateKeyStr));
        }

        // 开始广播交易
        vm.startBroadcast(deployerPrivateKey);

        // 部署合约
        SimpleNFT nft = new SimpleNFT("EveoneCat", "EVC");

        // 铸造第一个 NFT
        string memory tokenURI = "https://coffee-capable-gayal-349.mypinata.cloud/ipfs/bafkreigoq5pytutaclrstytly5hucmc42tkzpxqzd7iyjkh4p5cy4rt27a";
        uint256 tokenId = nft.mintNFT(msg.sender, tokenURI);

        // 停止广播
        vm.stopBroadcast();

        // 输出部署信息
        console.log("SimpleNFT deployed to:", address(nft));
        console.log("First NFT minted with tokenId:", tokenId);
        console.log("Owner:", nft.ownerOf(tokenId));
    }
}