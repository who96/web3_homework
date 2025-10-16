// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title MockDEX
 * @notice 简化的DEX模拟，用于在Sepolia测试网演示MemeToken的流动性功能
 * @dev 仅用于测试，不用于生产环境
 *
 * 功能：
 * - 模拟流动性池（token-ETH pair）
 * - 简单的恒定乘积AMM (x * y = k)
 * - 支持添加/移除流动性
 * - 支持swap交易
 */
contract MockDEX {
    IERC20 public immutable token;

    uint256 public tokenReserve;
    uint256 public ethReserve;

    mapping(address => uint256) public liquidityShares;
    uint256 public totalShares;

    event LiquidityAdded(address indexed provider, uint256 tokenAmount, uint256 ethAmount, uint256 shares);
    event LiquidityRemoved(address indexed provider, uint256 tokenAmount, uint256 ethAmount, uint256 shares);
    event Swap(address indexed user, bool tokenToEth, uint256 amountIn, uint256 amountOut);

    constructor(address _token) {
        token = IERC20(_token);
    }

    /**
     * @notice 添加流动性
     * @param tokenAmount 要添加的代币数量
     */
    function addLiquidity(uint256 tokenAmount) external payable returns (uint256 shares) {
        require(tokenAmount > 0 && msg.value > 0, "Invalid amounts");

        // 转入代币
        require(token.transferFrom(msg.sender, address(this), tokenAmount), "Transfer failed");

        // 计算LP份额
        if (totalShares == 0) {
            // 首次添加流动性
            shares = sqrt(tokenAmount * msg.value);
        } else {
            // 按比例计算份额
            uint256 sharesByToken = (tokenAmount * totalShares) / tokenReserve;
            uint256 sharesByEth = (msg.value * totalShares) / ethReserve;
            shares = sharesByToken < sharesByEth ? sharesByToken : sharesByEth;
        }

        require(shares > 0, "Insufficient liquidity minted");

        // 更新状态
        tokenReserve += tokenAmount;
        ethReserve += msg.value;
        liquidityShares[msg.sender] += shares;
        totalShares += shares;

        emit LiquidityAdded(msg.sender, tokenAmount, msg.value, shares);
    }

    /**
     * @notice 移除流动性
     * @param shares 要移除的LP份额
     */
    function removeLiquidity(uint256 shares) external returns (uint256 tokenAmount, uint256 ethAmount) {
        require(shares > 0 && shares <= liquidityShares[msg.sender], "Invalid shares");

        // 计算可取回的数量
        tokenAmount = (shares * tokenReserve) / totalShares;
        ethAmount = (shares * ethReserve) / totalShares;

        require(tokenAmount > 0 && ethAmount > 0, "Insufficient liquidity burned");

        // 更新状态
        liquidityShares[msg.sender] -= shares;
        totalShares -= shares;
        tokenReserve -= tokenAmount;
        ethReserve -= ethAmount;

        // 转出代币和ETH
        require(token.transfer(msg.sender, tokenAmount), "Token transfer failed");
        (bool success,) = msg.sender.call{value: ethAmount}("");
        require(success, "ETH transfer failed");

        emit LiquidityRemoved(msg.sender, tokenAmount, ethAmount, shares);
    }

    /**
     * @notice ETH换Token
     */
    function swapEthForToken(uint256 minTokenOut) external payable returns (uint256 tokenOut) {
        require(msg.value > 0, "Invalid ETH amount");

        // 恒定乘积公式: (x + dx) * (y - dy) = x * y
        // dy = y * dx / (x + dx)
        // 收取0.3%手续费
        uint256 ethIn = (msg.value * 997) / 1000; // 0.3% fee
        tokenOut = (tokenReserve * ethIn) / (ethReserve + ethIn);

        require(tokenOut >= minTokenOut, "Slippage too high");
        require(tokenOut <= tokenReserve, "Insufficient liquidity");

        // 更新储备
        ethReserve += msg.value;
        tokenReserve -= tokenOut;

        // 转出代币
        require(token.transfer(msg.sender, tokenOut), "Transfer failed");

        emit Swap(msg.sender, false, msg.value, tokenOut);
    }

    /**
     * @notice Token换ETH
     * @param tokenIn 输入的代币数量
     * @param minEthOut 最小输出ETH
     */
    function swapTokenForEth(uint256 tokenIn, uint256 minEthOut) external returns (uint256 ethOut) {
        require(tokenIn > 0, "Invalid token amount");

        // 转入代币
        require(token.transferFrom(msg.sender, address(this), tokenIn), "Transfer failed");

        // 恒定乘积公式
        uint256 tokenInAfterFee = (tokenIn * 997) / 1000; // 0.3% fee
        ethOut = (ethReserve * tokenInAfterFee) / (tokenReserve + tokenInAfterFee);

        require(ethOut >= minEthOut, "Slippage too high");
        require(ethOut <= ethReserve, "Insufficient liquidity");

        // 更新储备
        tokenReserve += tokenIn;
        ethReserve -= ethOut;

        // 转出ETH
        (bool success,) = msg.sender.call{value: ethOut}("");
        require(success, "ETH transfer failed");

        emit Swap(msg.sender, true, tokenIn, ethOut);
    }

    /**
     * @notice 查询ETH换Token的输出
     */
    function getTokenOut(uint256 ethIn) external view returns (uint256) {
        if (ethIn == 0 || ethReserve == 0) return 0;
        uint256 ethInAfterFee = (ethIn * 997) / 1000;
        return (tokenReserve * ethInAfterFee) / (ethReserve + ethInAfterFee);
    }

    /**
     * @notice 查询Token换ETH的输出
     */
    function getEthOut(uint256 tokenIn) external view returns (uint256) {
        if (tokenIn == 0 || tokenReserve == 0) return 0;
        uint256 tokenInAfterFee = (tokenIn * 997) / 1000;
        return (ethReserve * tokenInAfterFee) / (tokenReserve + tokenInAfterFee);
    }

    /**
     * @notice 平方根函数（Babylonian method）
     */
    function sqrt(uint256 y) internal pure returns (uint256 z) {
        if (y > 3) {
            z = y;
            uint256 x = y / 2 + 1;
            while (x < z) {
                z = x;
                x = (y / x + x) / 2;
            }
        } else if (y != 0) {
            z = 1;
        }
    }

    receive() external payable {}
}
