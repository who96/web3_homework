#!/bin/bash

echo "=== Go-Ethereum 交易生命周期演示 ==="
echo ""
echo "环境检查..."

# 检查 Go
if ! command -v go &> /dev/null; then
    echo "❌ Go 未安装"
    exit 1
fi
echo "✅ Go $(go version | awk '{print $3}')"

# 检查 Foundry
if ! command -v forge &> /dev/null; then
    echo "❌ Foundry 未安装"
    exit 1
fi
echo "✅ Foundry $(forge --version | head -1)"

# 检查 .env
if [ ! -f .env ]; then
    echo "❌ .env 文件不存在"
    exit 1
fi
echo "✅ .env 配置文件"

echo ""
echo "运行合约测试..."
make forge-test

echo ""
echo "启动主程序..."
go run main.go