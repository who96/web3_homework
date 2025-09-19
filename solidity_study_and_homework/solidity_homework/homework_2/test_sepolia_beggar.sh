#!/bin/bash

# Sepolia BeggingContract 测试脚本
# 使用方法: ./test_sepolia_beggar.sh <contract_address>

set -e

if [ -z "$1" ]; then
    echo "用法: $0 <contract_address>"
    echo "示例: $0 0x1234..."
    exit 1
fi

CONTRACT_ADDRESS="$1"
RPC_URL="sepolia"

# 检查环境变量
if [ -z "$PRIVATE_KEY" ]; then
    echo "错误: 请设置PRIVATE_KEY环境变量"
    exit 1
fi

echo "=== 🧪 BeggingContract Sepolia 测试 ==="
echo "合约地址: $CONTRACT_ADDRESS"
echo "网络: Sepolia"
echo ""

# 获取部署者地址
YOUR_ADDRESS=$(cast wallet address --private-key "$PRIVATE_KEY")
echo "测试地址: $YOUR_ADDRESS"

echo ""
echo "=== 📋 基础信息检查 ==="

echo "1. 检查合约owner:"
OWNER_RAW=$(cast call "$CONTRACT_ADDRESS" "owner()" --rpc-url "$RPC_URL")
# ABI编码的address是32字节，提取后20字节（40个十六进制字符）
OWNER="0x${OWNER_RAW: -40}"
echo "   Owner: $OWNER"

echo "2. 检查账户余额:"
BALANCE=$(cast balance "$YOUR_ADDRESS" --rpc-url "$RPC_URL")
echo "   余额: $BALANCE"

echo "3. 检查合约余额:"
CONTRACT_BALANCE=$(cast balance "$CONTRACT_ADDRESS" --rpc-url "$RPC_URL")
echo "   合约余额: $CONTRACT_BALANCE"

echo ""
echo "=== 🕐 时间检查 ==="

# 获取当前UTC+8时间
CURRENT_UTC_SECONDS=$(date -u +%s)
CURRENT_UTC8_HOUR=$(( (CURRENT_UTC_SECONDS / 3600 + 8) % 24 ))

echo "当前UTC+8时间: ${CURRENT_UTC8_HOUR}点"

if [ "$CURRENT_UTC8_HOUR" -ge 9 ] && [ "$CURRENT_UTC8_HOUR" -lt 18 ]; then
    echo "✅ 当前在工作时间内 (9:00-18:00 UTC+8)，可以捐赠"
    CAN_DONATE=true
else
    echo "❌ 当前在工作时间外，无法捐赠"
    CAN_DONATE=false
fi

echo ""
echo "=== 💰 捐赠测试 ==="

if [ "$CAN_DONATE" = true ]; then
    echo "尝试捐赠 0.01 ETH..."

    # 捐赠前查询
    DONATION_BEFORE=$(cast call "$CONTRACT_ADDRESS" "getDonation(address)" "$YOUR_ADDRESS" --rpc-url "$RPC_URL")
    echo "捐赠前金额: $DONATION_BEFORE"

    # 执行捐赠
    TX_HASH=$(cast send "$CONTRACT_ADDRESS" "donate()" --value 0.01ether --rpc-url "$RPC_URL" --private-key "$PRIVATE_KEY" 2>&1 | grep -o "0x[a-fA-F0-9]*" | head -1 || true)

    if [ -n "$TX_HASH" ]; then
        echo "✅ 捐赠成功! 交易hash: $TX_HASH"

        # 等待交易确认
        echo "等待交易确认..."
        sleep 3

        # 捐赠后查询
        DONATION_AFTER=$(cast call "$CONTRACT_ADDRESS" "getDonation(address)" "$YOUR_ADDRESS" --rpc-url "$RPC_URL")
        echo "捐赠后金额: $DONATION_AFTER"

    else
        echo "❌ 捐赠失败"
    fi
else
    echo "⏰ 跳过捐赠测试（时间限制）"
fi

echo ""
echo "=== 🏆 排行榜查询 ==="

echo "查询当前捐赠排行榜..."
LEADERBOARD=$(cast call "$CONTRACT_ADDRESS" "getTopDonors()" --rpc-url "$RPC_URL")
echo "排行榜数据: $LEADERBOARD"

echo ""
echo "=== 🔍 个人捐赠查询 ==="

MY_DONATION=$(cast call "$CONTRACT_ADDRESS" "getDonation(address)" "$YOUR_ADDRESS" --rpc-url "$RPC_URL")
echo "我的总捐赠: $MY_DONATION"

echo ""
echo "=== 👑 权限测试 ==="

# 将地址转换为小写进行比较（避免大小写问题）
OWNER_LOWER=$(echo "$OWNER" | tr '[:upper:]' '[:lower:]')
YOUR_ADDRESS_LOWER=$(echo "$YOUR_ADDRESS" | tr '[:upper:]' '[:lower:]')

echo "比较地址:"
echo "   Owner: $OWNER_LOWER"
echo "   Your:  $YOUR_ADDRESS_LOWER"

if [ "$OWNER_LOWER" = "$YOUR_ADDRESS_LOWER" ]; then
    echo "✅ 你是合约owner，可以测试withdraw功能"
    CONTRACT_BALANCE=$(cast balance "$CONTRACT_ADDRESS" --rpc-url "$RPC_URL")
    echo "合约当前余额: $CONTRACT_BALANCE"
else
    echo "ℹ️  你不是合约owner，无法测试withdraw功能"
fi

echo ""
echo "=== 💸 提现功能测试 ==="

if [ "$OWNER_LOWER" = "$YOUR_ADDRESS_LOWER" ]; then
    CONTRACT_BALANCE=$(cast balance "$CONTRACT_ADDRESS" --rpc-url "$RPC_URL")
    CONTRACT_BALANCE_WEI=$(cast to-wei "$CONTRACT_BALANCE" ether 2>/dev/null || echo "0")

    echo "合约余额: $CONTRACT_BALANCE"

    if [ "$CONTRACT_BALANCE_WEI" -gt 0 ]; then
        echo "💰 合约有余额，测试提现功能..."

        # 记录owner提现前余额
        OWNER_BALANCE_BEFORE=$(cast balance "$YOUR_ADDRESS" --rpc-url "$RPC_URL")
        echo "Owner提现前余额: $OWNER_BALANCE_BEFORE"

        # 执行提现
        echo "执行withdraw()..."
        WITHDRAW_TX=$(cast send "$CONTRACT_ADDRESS" "withdraw()" --rpc-url "$RPC_URL" --private-key "$PRIVATE_KEY" 2>&1 | grep -o "0x[a-fA-F0-9]*" | head -1 || true)

        if [ -n "$WITHDRAW_TX" ]; then
            echo "✅ 提现成功! 交易hash: $WITHDRAW_TX"

            # 等待交易确认
            echo "等待交易确认..."
            sleep 3

            # 检查提现后余额
            OWNER_BALANCE_AFTER=$(cast balance "$YOUR_ADDRESS" --rpc-url "$RPC_URL")
            CONTRACT_BALANCE_AFTER=$(cast balance "$CONTRACT_ADDRESS" --rpc-url "$RPC_URL")

            echo "Owner提现后余额: $OWNER_BALANCE_AFTER"
            echo "合约提现后余额: $CONTRACT_BALANCE_AFTER"

            # 验证提现结果
            if [ "$CONTRACT_BALANCE_AFTER" = "0" ]; then
                echo "✅ 合约余额已清空"
            else
                echo "⚠️  合约仍有余额: $CONTRACT_BALANCE_AFTER"
            fi
        else
            echo "❌ 提现失败"
        fi
    else
        echo "💰 合约余额为0，无法测试提现功能"
        echo "ℹ️  可以先捐赠一些ETH再测试提现"
    fi
else
    echo "🔒 非owner账户，测试拒绝访问..."

    # 测试非owner调用withdraw会被拒绝
    echo "尝试以非owner身份调用withdraw()（应该失败）..."
    WITHDRAW_RESULT=$(cast send "$CONTRACT_ADDRESS" "withdraw()" --rpc-url "$RPC_URL" --private-key "$PRIVATE_KEY" 2>&1 || true)

    if echo "$WITHDRAW_RESULT" | grep -q "Not owner"; then
        echo "✅ 正确拒绝非owner的提现请求"
    elif echo "$WITHDRAW_RESULT" | grep -q "revert"; then
        echo "✅ 交易被回滚（权限检查正常）"
    else
        echo "❌ 意外结果: $WITHDRAW_RESULT"
    fi
fi

echo ""
echo "=== 📊 测试总结 ==="
echo "合约地址: $CONTRACT_ADDRESS"
echo "测试地址: $YOUR_ADDRESS"
echo "是否为owner: $([ "$OWNER_LOWER" = "$YOUR_ADDRESS_LOWER" ] && echo "是" || echo "否")"
echo "当前时间: ${CURRENT_UTC8_HOUR}点 (UTC+8)"
echo "可否捐赠: $([ "$CAN_DONATE" = true ] && echo "是" || echo "否")"
echo "个人捐赠: $MY_DONATION"

echo ""
echo "🎉 测试完成!"

# 提供有用的命令
echo ""
echo "=== 💡 有用的命令 ==="
echo "# 基础查询"
echo "cast call $CONTRACT_ADDRESS \"owner()\" --rpc-url $RPC_URL"
echo "cast call $CONTRACT_ADDRESS \"getTopDonors()\" --rpc-url $RPC_URL"
echo "cast balance $CONTRACT_ADDRESS --rpc-url $RPC_URL"
echo ""
echo "# 捐赠（工作时间内）"
echo "cast send $CONTRACT_ADDRESS \"donate()\" --value 0.01ether --rpc-url $RPC_URL --private-key \$PRIVATE_KEY"
echo ""
echo "# 提现（仅owner）"
echo "cast send $CONTRACT_ADDRESS \"withdraw()\" --rpc-url $RPC_URL --private-key \$PRIVATE_KEY"
echo ""
echo "# 监控和调试"
if [ -n "$TX_HASH" ]; then
echo "查看交易详情: cast receipt $TX_HASH --rpc-url $RPC_URL"
fi
if [ -n "$WITHDRAW_TX" ]; then
echo "查看提现交易: cast receipt $WITHDRAW_TX --rpc-url $RPC_URL"
fi
echo "监控事件: cast logs --address $CONTRACT_ADDRESS --rpc-url $RPC_URL"
echo "在Etherscan查看: https://sepolia.etherscan.io/address/$CONTRACT_ADDRESS"