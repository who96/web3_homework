.PHONY: help build run test forge-test forge-build abigen contract-deploy contract-interact clean

help:
	@echo "Available commands:"
	@echo "  make build            - 构建 Go 项目"
	@echo "  make run              - 运行主程序"
	@echo "  make test             - 运行 Go 测试"
	@echo "  make forge-test       - 运行 Foundry 合约测试"
	@echo "  make forge-build      - 编译 Solidity 合约"
	@echo "  make abigen           - 生成合约 Go 绑定代码"
	@echo "  make contract-deploy  - 部署 SimpleStorage 合约"
	@echo "  make contract-interact ADDR=<address> - 与合约交互"
	@echo "  make clean            - 清理构建文件"

build:
	go build -o bin/go-eth-demo main.go

run:
	go run main.go

test:
	go test ./...

forge-test:
	cd contracts && forge test -vvv

forge-build:
	cd contracts && forge build

abigen:
	@echo "=== 生成合约 Go 绑定 ==="
	@echo "1. 编译合约..."
	@cd contracts && forge build
	@echo "2. 提取 ABI 和 Bytecode..."
	@mkdir -p contracts/bindings
	@jq '.abi' contracts/out/SimpleStorage.sol/SimpleStorage.json > contracts/bindings/SimpleStorage.abi
	@jq -r '.bytecode.object' contracts/out/SimpleStorage.sol/SimpleStorage.json > contracts/bindings/SimpleStorage.bin
	@echo "3. 生成 Go 绑定代码..."
	@$(shell go env GOPATH)/bin/abigen \
		--abi contracts/bindings/SimpleStorage.abi \
		--bin contracts/bindings/SimpleStorage.bin \
		--pkg bindings \
		--type SimpleStorage \
		--out contracts/bindings/SimpleStorage.go
	@echo "✅ Go 绑定代码生成完成: contracts/bindings/SimpleStorage.go"

contract-deploy:
	@echo "=== 部署 SimpleStorage 合约到 Sepolia ==="
	@cd examples/contract && go run deploy.go

contract-interact:
	@if [ -z "$(ADDR)" ]; then \
		echo "错误: 请提供合约地址"; \
		echo "用法: make contract-interact ADDR=0x..."; \
		exit 1; \
	fi
	@echo "=== 与 SimpleStorage 合约交互 ==="
	@cd examples/contract && go run interact.go $(ADDR)

clean:
	rm -rf bin/ contracts/out/ contracts/cache/ contracts/bindings/*.abi contracts/bindings/*.bin