# Foundry 命令简化

# 默认网络
NETWORK ?= sepolia

# 编译合约
build:
	forge build

# 运行测试
test:
	forge test -vvv

# 运行特定测试
test-match:
	forge test --match-test $(TEST) -vvv

# Gas 报告
gas:
	forge test --gas-report

# 覆盖率报告
coverage:
	forge coverage

# 格式化代码
format:
	forge fmt

# 本地部署（Anvil）
deploy-local:
	forge script script/Deploy.s.sol --rpc-url http://localhost:8545 --broadcast

# SimpleNFT 测试网部署
deploy-testnet:
	forge script script/Deploy.s.sol --rpc-url $(NETWORK) --broadcast --verify

# BeggingContract 部署
deploy-beggar:
	forge script script/DeployBeggar.s.sol --rpc-url $(NETWORK) --broadcast --verify

# BeggingContract 快速部署（不验证）
deploy-beggar-fast:
	forge create --rpc-url $(NETWORK) --private-key $(PRIVATE_KEY) contracts/getMoney.sol:BeggingContract

# 主网部署（危险！）
deploy-mainnet:
	@echo "⚠️  警告：即将部署到主网！"
	@echo "按 Ctrl+C 取消，或按 Enter 继续..."
	@read dummy
	forge script script/Deploy.s.sol --rpc-url mainnet --broadcast --verify

# 验证合约
verify:
	forge verify-contract $(CONTRACT) --chain $(NETWORK)

# 清理
clean:
	forge clean
	rm -rf out cache

# 安装依赖
install:
	forge install

# 启动本地节点
anvil:
	anvil --fork-url $(NETWORK)

# BeggingContract 测试（仅本地）
test-beggar:
	forge test --match-contract BeggingContract -vvv

# BeggingContract Sepolia交互测试（需要CONTRACT_ADDRESS环境变量）
test-beggar-sepolia:
	./test_sepolia_beggar.sh $(CONTRACT_ADDRESS)

# BeggingContract 提现测试（仅owner）
withdraw-beggar:
	@echo "⚠️  提现操作！合约所有资金将转移到owner地址"
	@echo "合约地址: $(CONTRACT_ADDRESS)"
	@echo "按 Enter 继续，Ctrl+C 取消..."
	@read dummy
	cast send $(CONTRACT_ADDRESS) "withdraw()" --rpc-url $(NETWORK) --private-key $(PRIVATE_KEY)

# 查看BeggingContract信息
beggar-info:
	@echo "合约地址: $(CONTRACT_ADDRESS)"
	@echo "网络: $(NETWORK)"
	@echo ""
	@echo "合约owner:"
	@OWNER_RAW=$$(cast call $(CONTRACT_ADDRESS) "owner()" --rpc-url $(NETWORK)); \
	OWNER="0x$${OWNER_RAW: -40}"; \
	echo "  Raw: $$OWNER_RAW"; \
	echo "  Address: $$OWNER"
	@echo ""
	@echo "排行榜:"
	@cast call $(CONTRACT_ADDRESS) "getTopDonors()" --rpc-url $(NETWORK)
	@echo ""
	@echo "合约余额:"
	@cast balance $(CONTRACT_ADDRESS) --rpc-url $(NETWORK)

.PHONY: build test test-match gas coverage format deploy-local deploy-testnet deploy-beggar deploy-beggar-fast deploy-mainnet verify clean install anvil test-beggar test-beggar-sepolia withdraw-beggar beggar-info