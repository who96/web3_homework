// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/MemeToken.sol";

// Uniswap V2 接口
interface IUniswapV2Router02 {
    function addLiquidityETH(
        address token,
        uint256 amountTokenDesired,
        uint256 amountTokenMin,
        uint256 amountETHMin,
        address to,
        uint256 deadline
    ) external payable returns (uint256 amountToken, uint256 amountETH, uint256 liquidity);

    function factory() external pure returns (address);
}

interface IUniswapV2Factory {
    function getPair(address tokenA, address tokenB) external view returns (address pair);
}

/**
 * @title AddLiquidity
 * @notice 为MemeToken添加Uniswap V2流动性的脚本
 * @dev 使用方法:
 *   forge script script/AddLiquidity.s.sol --rpc-url $SEPOLIA_RPC_URL --broadcast
 *
 * 需要在.env中设置:
 *   - PRIVATE_KEY: 部署者私钥
 *   - MEME_TOKEN_ADDRESS: 已部署的MemeToken合约地址
 */
contract AddLiquidity is Script {
    // Uniswap V2 Router地址
    // Mainnet: 0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D
    // Sepolia: 需要部署或使用已有的Router
    address constant UNISWAP_V2_ROUTER = 0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D;

    // WETH地址
    // Mainnet: 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2
    // Sepolia: 0x7b79995e5f793A07Bc00c21412e50Ecae098E7f9
    address constant WETH = 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2;

    // 流动性参数
    uint256 constant TOKEN_AMOUNT = 210_000 * 10 ** 18; // 210,000 FK (1%)
    uint256 constant ETH_AMOUNT = 0.1 ether; // 0.1 ETH

    function run() external {
        // 读取环境变量
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address tokenAddress = vm.envAddress("MEME_TOKEN_ADDRESS");

        require(tokenAddress != address(0), "MEME_TOKEN_ADDRESS not set");

        address deployer = vm.addr(deployerPrivateKey);
        MemeToken token = MemeToken(tokenAddress);
        IUniswapV2Router02 router = IUniswapV2Router02(UNISWAP_V2_ROUTER);

        console.log("===========================================");
        console.log("Adding Liquidity to Uniswap V2");
        console.log("===========================================");
        console.log("Token:", tokenAddress);
        console.log("Router:", UNISWAP_V2_ROUTER);
        console.log("Deployer:", deployer);
        console.log("Token amount:", TOKEN_AMOUNT / 10 ** 18, "FK");
        console.log("ETH amount:", ETH_AMOUNT / 1 ether, "ETH");
        console.log("===========================================");

        // 检查余额
        uint256 tokenBalance = token.balanceOf(deployer);
        uint256 ethBalance = deployer.balance;

        console.log("Deployer token balance:", tokenBalance / 10 ** 18, "FK");
        console.log("Deployer ETH balance:", ethBalance / 1 ether, "ETH");

        require(tokenBalance >= TOKEN_AMOUNT, "Insufficient token balance");
        require(ethBalance >= ETH_AMOUNT, "Insufficient ETH balance");

        vm.startBroadcast(deployerPrivateKey);

        // 1. 授权Router使用代币
        console.log("\nStep 1: Approving router...");
        token.approve(UNISWAP_V2_ROUTER, TOKEN_AMOUNT);
        console.log("Approved!");

        // 2. 添加流动性
        console.log("\nStep 2: Adding liquidity...");
        (uint256 amountToken, uint256 amountETH, uint256 liquidity) = router.addLiquidityETH{
            value: ETH_AMOUNT
        }(
            tokenAddress,
            TOKEN_AMOUNT,
            TOKEN_AMOUNT * 95 / 100, // 允许5%滑点
            ETH_AMOUNT * 95 / 100,
            deployer,
            block.timestamp + 300
        );

        console.log("Liquidity added!");
        console.log("- Token amount:", amountToken / 10 ** 18, "FK");
        console.log("- ETH amount:", amountETH / 1 ether, "ETH");
        console.log("- LP tokens received:", liquidity);

        // 3. 获取Pair地址
        IUniswapV2Factory factory = IUniswapV2Factory(router.factory());
        address pair = factory.getPair(tokenAddress, WETH);
        console.log("\nUniswap Pair address:", pair);

        console.log("\n===========================================");
        console.log("IMPORTANT: Add this pair to whitelist!");
        console.log("===========================================");
        console.log("Run the following commands as owner:");
        console.log("");
        console.log("cast send", tokenAddress);
        console.log("  'setTaxExempt(address,bool)'");
        console.log("  ", pair, "true");
        console.log("  --rpc-url $RPC_URL --private-key $PRIVATE_KEY");
        console.log("");
        console.log("cast send", tokenAddress);
        console.log("  'setLimitExempt(address,bool)'");
        console.log("  ", pair, "true");
        console.log("  --rpc-url $RPC_URL --private-key $PRIVATE_KEY");
        console.log("===========================================");

        vm.stopBroadcast();
    }
}
