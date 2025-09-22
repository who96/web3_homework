# NFT Auction Market E2E Testing Makefile
# Author: Linus-style systematic testing approach
# Usage: make help

# Load environment variables
include .env
export

# Network configuration
NETWORK = sepolia
RPC_URL = $(SEPOLIA_RPC_URL)

# Contract addresses (updated for simplified proxy architecture)
NFT_CONTRACT = 0x6d61687CDE7F12A9c31aD833b305EF0c65EA830b
AUCTION_CONTRACT = 0x687265FBABe670a18D0274478984D6c9a03CCCb6
PRICE_FEED_CONTRACT = 0x15c8CED44bbdc1fec603BB9147EA51Dcc0809d7a

# Proxy architecture addresses
PROXY_ADMIN = 0x17D768939362FdfE8e3EC586A15e258E270d7BE7
IMPLEMENTATION_V1 = 0x9b1e5223Cb5BDb82bA05F75130330c55f3445169
IMPLEMENTATION_V2 = 0x4e8229e669b3c45adFa6Ce1421d002e013EdB122

# Test parameters
AUCTION_DURATION = 120
RESERVE_PRICE = 0.0001ether
BID_AMOUNT_1 = 0.0001ether
BID_AMOUNT_2 = 0.0002ether
BID_AMOUNT_3 = 0.0003ether
BID_AMOUNT_4 = 0.0004ether

# Wallet addresses (calculated from private keys)
WALLET1 = 0x5D4a329B262AC7a1d9Ae0f4C54171dF61E2c0b69
WALLET2 = 0xF4076C4a38971D71812B298A6aA9213C5425fa51
WALLET3 = 0x3b8802408e3da17Bc66aA6a8dAb537FD49403355
WALLET4 = 0x1c43bbdD79C85F72dEd3be8dc8a3BB3395e2BAA0
WALLET5 = 0xFD74f145EB49CB9B84834CFAd306A2a1C67d1A22
WALLET6 = 0x912E32baE81a95C5e39952a9E2E9085e579e7940

.PHONY: help clean build deploy test-e2e verify-deployment

help:
	@echo "NFT Auction Market Testing Commands"
	@echo "=================================="
	@echo "make build           - Compile contracts"
	@echo "make deploy          - Deploy all contracts to Sepolia"
	@echo "make test-e2e        - Run complete E2E test"
	@echo "make verify-deployment - Verify contract deployment"
	@echo "make clean           - Clean build artifacts"
	@echo ""
	@echo "E2E Test Phases:"
	@echo "make phase1          - Create NFT and auction"
	@echo "make phase2          - Execute bidding phase"
	@echo "make phase3          - Test price feed"
	@echo "make phase4          - Test platform wallet"
	@echo "make phase5          - End auction and settle"
	@echo "make phase6          - Test contract upgrade"
	@echo ""
	@echo "Utilities:"
	@echo "make show-wallets    - Display all wallet addresses"
	@echo "make check-balances  - Check all wallet balances"
	@echo ""
	@echo "Verification Commands:"
	@echo "make verify-contract-version - Check current contract version"
	@echo "make verify-implementation   - Check current implementation address"
	@echo "make verify-emergency-pause  - Check emergency pause state"
	@echo "make verify-upgrade-complete - Complete upgrade verification"

clean:
	@echo "🧹 Cleaning build artifacts..."
	forge clean

build:
	@echo "🔨 Building contracts..."
	forge build

deploy: build
	@echo "🚀 Deploying contracts to Sepolia..."
	forge script script/DeployContracts.s.sol:DeployContracts \
		--rpc-url $(RPC_URL) \
		--broadcast \
		--verify \
		--etherscan-api-key $(ETHERSCAN_API_KEY) \
		--legacy
	@echo "✅ Deployment complete. Update contract addresses in Makefile!"

verify-deployment:
	@echo "🔍 Verifying deployment..."
	@echo "NFT Contract: $(NFT_CONTRACT)"
	@cast call $(NFT_CONTRACT) "totalSupply()" --rpc-url $(RPC_URL)
	@echo "Auction Contract: $(AUCTION_CONTRACT)"
	@cast call $(AUCTION_CONTRACT) "auctionCounter()" --rpc-url $(RPC_URL)
	@echo "Price Feed Contract: $(PRICE_FEED_CONTRACT)"
	@cast call $(PRICE_FEED_CONTRACT) "DECIMALS()" --rpc-url $(RPC_URL)

show-wallets:
	@echo "💳 Wallet Addresses:"
	@echo "Wallet1 (seller):    $(WALLET1)"
	@echo "Wallet2 (bidder1):   $(WALLET2)"
	@echo "Wallet3 (bidder2):   $(WALLET3)"
	@echo "Wallet4 (bidder3):   $(WALLET4)"
	@echo "Wallet5 (bidder4):   $(WALLET5)"
	@echo "Wallet6 (platform):  $(WALLET6)"

check-balances:
	@echo "💰 Current Balances:"
	@echo -n "Wallet1: "; cast balance $(WALLET1) --rpc-url $(RPC_URL) --ether
	@echo -n "Wallet2: "; cast balance $(WALLET2) --rpc-url $(RPC_URL) --ether
	@echo -n "Wallet3: "; cast balance $(WALLET3) --rpc-url $(RPC_URL) --ether
	@echo -n "Wallet4: "; cast balance $(WALLET4) --rpc-url $(RPC_URL) --ether
	@echo -n "Wallet5: "; cast balance $(WALLET5) --rpc-url $(RPC_URL) --ether
	@echo -n "Wallet6: "; cast balance $(WALLET6) --rpc-url $(RPC_URL) --ether

# Phase 1: Create NFT and Auction
phase1:
	@echo "🎯 Phase 1: Create NFT and Auction"
	@echo "1.1 Minting NFT..."
	cast send $(NFT_CONTRACT) "mint(address)" $(WALLET1) \
		--private-key $(PRIVATE_KEY) --rpc-url $(RPC_URL) --legacy
	@echo "1.2 Approving auction contract..."
	cast send $(NFT_CONTRACT) "setApprovalForAll(address,bool)" $(AUCTION_CONTRACT) true \
		--private-key $(PRIVATE_KEY) --rpc-url $(RPC_URL) --legacy
	@echo "1.3 Creating auction..."
	cast send $(AUCTION_CONTRACT) "createAuction(address,uint256,uint256,uint256)" \
		$(NFT_CONTRACT) 0 $(AUCTION_DURATION) $(RESERVE_PRICE) \
		--private-key $(PRIVATE_KEY) --rpc-url $(RPC_URL) --legacy
	@echo "✅ Phase 1 complete"

# Phase 2: Bidding Phase
phase2:
	@echo "🎯 Phase 2: Bidding Phase"
	@echo "2.1 Wallet2 bids $(BID_AMOUNT_1)..."
	cast send $(AUCTION_CONTRACT) "placeBid(uint256)" 0 \
		--value $(BID_AMOUNT_1) --private-key $(PRIVATE_KEY_2) --rpc-url $(RPC_URL) --legacy
	@echo "2.2 Wallet3 bids $(BID_AMOUNT_2)..."
	cast send $(AUCTION_CONTRACT) "placeBid(uint256)" 0 \
		--value $(BID_AMOUNT_2) --private-key $(PRIVATE_KEY_3) --rpc-url $(RPC_URL) --legacy
	@echo "2.3 Wallet4 bids $(BID_AMOUNT_3)..."
	cast send $(AUCTION_CONTRACT) "placeBid(uint256)" 0 \
		--value $(BID_AMOUNT_3) --private-key $(PRIVATE_KEY_4) --rpc-url $(RPC_URL) --legacy
	@echo "2.4 Wallet5 bids $(BID_AMOUNT_4)..."
	cast send $(AUCTION_CONTRACT) "placeBid(uint256)" 0 \
		--value $(BID_AMOUNT_4) --private-key $(PRIVATE_KEY_5) --rpc-url $(RPC_URL) --legacy
	@echo "✅ Phase 2 complete"

# Phase 3: Test Price Feed
phase3:
	@echo "🎯 Phase 3: Test Price Feed"
	@echo "3.1 Testing getEthUsdPrice for reserve price..."
	@cast call $(PRICE_FEED_CONTRACT) "getEthUsdPrice(uint256)" $(RESERVE_PRICE) --rpc-url $(RPC_URL)
	@echo "3.2 Testing getEthUsdPrice for 1 ETH..."
	@cast call $(PRICE_FEED_CONTRACT) "getEthUsdPrice(uint256)" 1ether --rpc-url $(RPC_URL)
	@echo "3.3 Testing auction price in USD..."
	@cast call $(AUCTION_CONTRACT) "getAuctionPriceInUSD(uint256)" 0 --rpc-url $(RPC_URL)
	@echo "✅ Phase 3 complete"

# Phase 4: Test Platform Wallet
phase4:
	@echo "🎯 Phase 4: Test Platform Wallet Function"
	@echo "4.1 Setting wallet6 as new platform wallet..."
	cast send $(AUCTION_CONTRACT) "setPlatformWallet(address)" $(WALLET6) \
		--private-key $(PRIVATE_KEY) --rpc-url $(RPC_URL) --legacy
	@echo "✅ Phase 4 complete - platform wallet changed"

# Phase 5: End Auction and Settlement
phase5:
	@echo "🎯 Phase 5: End Auction and Settlement"
	@echo "5.1 Ending auction..."
	cast send $(AUCTION_CONTRACT) "endAuction(uint256)" 0 \
		--private-key $(PRIVATE_KEY) --rpc-url $(RPC_URL) --legacy
	@echo "5.2 Winner claims NFT..."
	cast send $(AUCTION_CONTRACT) "claimNFT(uint256)" 0 \
		--private-key $(PRIVATE_KEY_5) --rpc-url $(RPC_URL) --legacy
	@echo "5.3 Seller claims funds..."
	cast send $(AUCTION_CONTRACT) "claimFunds(uint256)" 0 \
		--private-key $(PRIVATE_KEY) --rpc-url $(RPC_URL) --legacy
	@echo "5.4 Failed bidders claim refunds..."
	-cast send $(AUCTION_CONTRACT) "withdrawRefund()" \
		--private-key $(PRIVATE_KEY_2) --rpc-url $(RPC_URL) --legacy
	-cast send $(AUCTION_CONTRACT) "withdrawRefund()" \
		--private-key $(PRIVATE_KEY_3) --rpc-url $(RPC_URL) --legacy
	-cast send $(AUCTION_CONTRACT) "withdrawRefund()" \
		--private-key $(PRIVATE_KEY_4) --rpc-url $(RPC_URL) --legacy
	@echo "✅ Phase 5 complete"

# Phase 6: Contract Upgrade Testing
phase6:
	@echo "🎯 Phase 6: Contract Upgrade Testing"
	@echo "6.1 Verifying current contract version..."
	@cast call $(AUCTION_CONTRACT) "getVersion()" --rpc-url $(RPC_URL)
	@echo "6.2 Recording pre-upgrade state..."
	@echo -n "Auction counter: "; cast call $(AUCTION_CONTRACT) "auctionCounter()" --rpc-url $(RPC_URL)
	@echo "6.3 Deploying V2 implementation if needed..."
	@echo "V2 Implementation already at: $(IMPLEMENTATION_V2)"
	@echo "6.4 Executing upgrade to V2..."
	@cast send $(PROXY_ADMIN) "upgradeAndCall(address,address,bytes)" \
		$(AUCTION_CONTRACT) $(IMPLEMENTATION_V2) "0x" \
		--private-key $(PRIVATE_KEY) --rpc-url $(RPC_URL) --legacy
	@echo "6.5 Verifying upgrade success..."
	@cast call $(AUCTION_CONTRACT) "getVersion()" --rpc-url $(RPC_URL)
	@echo -n "Auction counter after upgrade: "; cast call $(AUCTION_CONTRACT) "auctionCounter()" --rpc-url $(RPC_URL)
	@echo "6.6 Testing V2 emergency pause feature..."
	@echo -n "Initial pause state: "; cast call $(AUCTION_CONTRACT) "isEmergencyPaused()" --rpc-url $(RPC_URL)
	@echo "6.7 Toggling emergency pause..."
	@cast send $(AUCTION_CONTRACT) "toggleEmergencyPause()" \
		--private-key $(PRIVATE_KEY) --rpc-url $(RPC_URL) --legacy
	@echo -n "Pause state after toggle: "; cast call $(AUCTION_CONTRACT) "isEmergencyPaused()" --rpc-url $(RPC_URL)
	@echo "6.8 Testing auction creation while paused (should fail)..."
	-cast send $(AUCTION_CONTRACT) "createAuction(address,uint256,uint256,uint256)" \
		$(NFT_CONTRACT) 1 $(AUCTION_DURATION) $(RESERVE_PRICE) \
		--private-key $(PRIVATE_KEY) --rpc-url $(RPC_URL) --legacy
	@echo "6.9 Toggling pause off..."
	@cast send $(AUCTION_CONTRACT) "toggleEmergencyPause()" \
		--private-key $(PRIVATE_KEY) --rpc-url $(RPC_URL) --legacy
	@echo -n "Final pause state: "; cast call $(AUCTION_CONTRACT) "isEmergencyPaused()" --rpc-url $(RPC_URL)
	@echo "✅ Phase 6 complete - Contract upgraded successfully!"

# Wait for auction to end (manual step)
wait-auction:
	@echo "⏰ Waiting for auction to end ($(AUCTION_DURATION) seconds)..."
	@echo "Run 'make phase5' after waiting or check auction status:"
	@echo "cast call $(AUCTION_CONTRACT) \"isAuctionActive(uint256)\" 0 --rpc-url $(RPC_URL)"

# Complete E2E test
test-e2e: phase1
	@echo "⏳ Waiting 30 seconds before bidding phase..."
	sleep 30
	$(MAKE) phase2
	$(MAKE) phase3
	$(MAKE) phase4
	@echo "⏰ Please wait $(AUCTION_DURATION) seconds for auction to end, then run 'make phase5' followed by 'make phase6'"
	@echo "Or check auction status: cast call $(AUCTION_CONTRACT) \"isAuctionActive(uint256)\" 0 --rpc-url $(RPC_URL)"

# Manual E2E test (user controlled timing)
test-e2e-manual:
	@echo "🎯 Manual E2E Test - Execute phases individually:"
	@echo "1. make phase1"
	@echo "2. make phase2"
	@echo "3. make phase3"
	@echo "4. make phase4"
	@echo "5. Wait $(AUCTION_DURATION) seconds"
	@echo "6. make phase5"
	@echo "7. make phase6"

# Verification commands
verify-nft-owner:
	@echo "🔍 NFT Owner:"
	@cast call $(NFT_CONTRACT) "ownerOf(uint256)" 0 --rpc-url $(RPC_URL)

verify-auction-state:
	@echo "🔍 Auction State:"
	@cast call $(AUCTION_CONTRACT) "getAuction(uint256)" 0 --rpc-url $(RPC_URL)

verify-platform-fee:
	@echo "🔍 Platform Fee (3% = 300/10000):"
	@cast call $(AUCTION_CONTRACT) "PLATFORM_FEE()" --rpc-url $(RPC_URL)

verify-contract-version:
	@echo "🔍 Contract Version:"
	@cast call $(AUCTION_CONTRACT) "getVersion()" --rpc-url $(RPC_URL)

verify-proxy-admin:
	@echo "🔍 Proxy Admin:"
	@cast call $(PROXY_ADMIN) "owner()" --rpc-url $(RPC_URL)

verify-implementation:
	@echo "🔍 Current Implementation:"
	@cast call $(PROXY_ADMIN) "getProxyImplementation(address)" $(AUCTION_CONTRACT) --rpc-url $(RPC_URL)

verify-emergency-pause:
	@echo "🔍 Emergency Pause State:"
	@cast call $(AUCTION_CONTRACT) "isEmergencyPaused()" --rpc-url $(RPC_URL)

verify-upgrade-complete:
	@echo "🔍 Complete Upgrade Verification:"
	@echo "Contract Version:"
	@cast call $(AUCTION_CONTRACT) "getVersion()" --rpc-url $(RPC_URL)
	@echo "Current Implementation:"
	@cast call $(PROXY_ADMIN) "getProxyImplementation(address)" $(AUCTION_CONTRACT) --rpc-url $(RPC_URL)
	@echo "Emergency Pause Available:"
	@cast call $(AUCTION_CONTRACT) "isEmergencyPaused()" --rpc-url $(RPC_URL)

# Update contract addresses after deployment
update-addresses:
	@echo "📝 Update the following variables in Makefile after deployment:"
	@echo "NFT_CONTRACT = [new_nft_address]"
	@echo "AUCTION_CONTRACT = [new_auction_address]"
	@echo "PRICE_FEED_CONTRACT = [new_price_feed_address]"