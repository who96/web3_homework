// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fukuastake

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// FukuastakeMetaData contains all meta data concerning the Fukuastake contract.
var FukuastakeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ETH_PID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addPool\",\"inputs\":[{\"name\":\"_stTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poolWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minDepositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_unstakeLockedBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_withUpdate\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositEth\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"endBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fukuaPerBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fukuaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMultiplier\",\"inputs\":[{\"name\":\"_from\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_to\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"multiplier\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_fukuaToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_startBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_endBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_fukuaPerBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"massUpdatePools\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseClaim\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseWithdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingFukua\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingFukuaByBlockNumber\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pool\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"stTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"poolWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastRewardBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accFukuaPerShare\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minDepositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unstakeLockedBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEndBlock\",\"inputs\":[{\"name\":\"_endBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFukuaPerBlock\",\"inputs\":[{\"name\":\"_fukuaPerBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFukuaToken\",\"inputs\":[{\"name\":\"_fukuaToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPoolWeight\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_poolWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_withUpdate\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStartBlock\",\"inputs\":[{\"name\":\"_startBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakingBalance\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"startBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalPoolWeight\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseClaim\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseWithdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstake\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePool\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePool\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minDepositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_unstakeLockedBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"user\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"stAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finishedFukua\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pendingFukua\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAmount\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"requestAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pendingWithdrawAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AddPool\",\"inputs\":[{\"name\":\"stTokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"poolWeight\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"lastRewardBlock\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"minDepositAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"unstakeLockedBlocks\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Claim\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"poolId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"fukuaReward\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"poolId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauseClaim\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauseWithdraw\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestUnstake\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"poolId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetEndBlock\",\"inputs\":[{\"name\":\"endBlock\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetFukuaPerBlock\",\"inputs\":[{\"name\":\"newFukuaPerBlock\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetFukuaToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIERC20\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetPoolWeight\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"poolWeight\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"totalPoolWeight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetStartBlock\",\"inputs\":[{\"name\":\"startBlock\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnpauseClaim\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnpauseWithdraw\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatePool\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"lastRewardBlock\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"totalFukuaReward\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatePoolInfo\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"minDepositAmount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"unstakeLockedBlocks\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"poolId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a0604052306080523480156012575f5ffd5b5060805161346f6100395f395f8181612970015281816129990152612b2d015261346f5ff3fe6080604052600436106102bf575f3560e01c806375b238fc1161016f578063b6d9d919116100d8578063de065caa11610092578063f35e4a6e1161006d578063f35e4a6e1461082f578063fad07ece1461084e578063fe3131121461086d578063ff423357146108cb575f5ffd5b8063de065caa146107dd578063e26c25ce146107f1578063e2bbb15814610810575f5ffd5b8063b6d9d9191461072e578063b908afa81461074d578063bfc3ebba146106ab578063c713aa9414610780578063d547741f1461079f578063d86c0444146107be575f5ffd5b80639e2c8a5b116101295780639e2c8a5b1461064f5780639fac8a6a1461066e578063a217fddf146106ab578063a2e8fca2146106be578063ab5e124a146106d3578063ad3cb1cc146106f1575f5ffd5b806375b238fc146105aa5780637ef64c46146105ca5780638456cb59146105e95780638dbb1e3a146105fd5780638ff095f91461061c57806391d1485414610630575f5ffd5b80633ee521e21161022b57806351eb05a6116101e55780635c0ab30b116101c05780635c0ab30b146105405780635c975abb1461055f5780636155e3de14610582578063630b5ba114610596575f5ffd5b806351eb05a6146104f957806352d1902d146105185780635bb6d0071461052c575f5ffd5b80633ee521e2146104785780633f4ba83a14610497578063439370b1146104ab57806348cd4cb1146104b35780634ec81af1146104c75780634f1ef286146104e6575f5ffd5b80632e1a7d4d1161027c5780632e1a7d4d146103815780632f2ff15d146103a25780632f3ffb9f146103c157806336568abe146103da57806337849b3c146103f9578063379607f514610459575f5ffd5b806301ffc9a7146102c357806302559004146102f7578063081e3eda1461031a578063083c63231461032e5780631154823414610343578063248a9ca314610362575b5f5ffd5b3480156102ce575f5ffd5b506102e26102dd366004612ede565b6108ff565b60405190151581526020015b60405180910390f35b348015610302575f5ffd5b5061030c60045481565b6040519081526020016102ee565b348015610325575f5ffd5b5060055461030c565b348015610339575f5ffd5b5061030c60015481565b34801561034e575f5ffd5b5061030c61035d366004612f19565b610935565b34801561036d575f5ffd5b5061030c61037c366004612f47565b61096c565b34801561038c575f5ffd5b506103a061039b366004612f47565b61098c565b005b3480156103ad575f5ffd5b506103a06103bc366004612f19565b610bab565b3480156103cc575f5ffd5b506003546102e29060ff1681565b3480156103e5575f5ffd5b506103a06103f4366004612f19565b610bcd565b348015610404575f5ffd5b5061043e610413366004612f19565b600660209081525f928352604080842090915290825290208054600182015460029092015490919083565b604080519384526020840192909252908201526060016102ee565b348015610464575f5ffd5b506103a0610473366004612f47565b610c05565b348015610483575f5ffd5b506103a0610492366004612f5e565b610d44565b3480156104a2575f5ffd5b506103a0610db4565b6103a0610dd6565b3480156104be575f5ffd5b5061030c5f5481565b3480156104d2575f5ffd5b506103a06104e1366004612f79565b610e99565b6103a06104f4366004612fc5565b611059565b348015610504575f5ffd5b506103a0610513366004612f47565b611074565b348015610523575f5ffd5b5061030c6111f8565b348015610537575f5ffd5b506103a0611213565b34801561054b575f5ffd5b5061030c61055a366004612f19565b6112bc565b34801561056a575f5ffd5b505f51602061341a5f395f51905f525460ff166102e2565b34801561058d575f5ffd5b506103a06112da565b3480156105a1575f5ffd5b506103a061137c565b3480156105b5575f5ffd5b5061030c5f5160206133ba5f395f51905f5281565b3480156105d5575f5ffd5b5061030c6105e436600461308b565b61139a565b3480156105f4575f5ffd5b506103a06114b7565b348015610608575f5ffd5b5061030c6106173660046130c0565b6114d6565b348015610627575f5ffd5b506103a06115f2565b34801561063b575f5ffd5b506102e261064a366004612f19565b61169b565b34801561065a575f5ffd5b506103a06106693660046130c0565b6116d1565b348015610679575f5ffd5b50600354610693906201000090046001600160a01b031681565b6040516001600160a01b0390911681526020016102ee565b3480156106b6575f5ffd5b5061030c5f81565b3480156106c9575f5ffd5b5061030c60025481565b3480156106de575f5ffd5b506003546102e290610100900460ff1681565b3480156106fc575f5ffd5b50610721604051806040016040528060058152602001640352e302e360dc1b81525081565b6040516102ee91906130e0565b348015610739575f5ffd5b506103a0610748366004613122565b6118ab565b348015610758575f5ffd5b5061030c7fcab03bc4dbcc648cd59d6bbe9f848d1e9092f914016aa290ee92e18700d1e6f981565b34801561078b575f5ffd5b506103a061079a366004612f47565b611b98565b3480156107aa575f5ffd5b506103a06107b9366004612f19565b611c03565b3480156107c9575f5ffd5b506103a06107d8366004613174565b611c1f565b3480156107e8575f5ffd5b506103a0611cc4565b3480156107fc575f5ffd5b506103a061080b366004612f47565b611d68565b34801561081b575f5ffd5b506103a061082a3660046130c0565b611df5565b34801561083a575f5ffd5b506103a0610849366004612f47565b611ef2565b348015610859575f5ffd5b506103a061086836600461319d565b611f5b565b348015610878575f5ffd5b5061088c610887366004612f47565b61207b565b604080516001600160a01b0390981688526020880196909652948601939093526060850191909152608084015260a083015260c082015260e0016102ee565b3480156108d6575f5ffd5b506108ea6108e5366004612f19565b6120d2565b604080519283526020830191909152016102ee565b5f6001600160e01b03198216637965db0b60e01b148061092f57506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f82610940816121ad565b5f8481526006602090815260408083206001600160a01b038716845290915290205491505b5092915050565b5f9081525f5160206133fa5f395f51905f52602052604090206001015490565b6109946121ec565b8061099e816121ad565b6109a661221e565b5f600583815481106109ba576109ba6131d3565b5f918252602080832086845260068252604080852033865290925290832060079092020192509080805b6003840154811015610a615743846003018281548110610a0657610a066131d3565b905f5260205f2090600202016001015411610a6157836003018181548110610a3057610a306131d3565b905f5260205f2090600202015f015483610a4a91906131fb565b925081610a568161320e565b9250506001016109e4565b505f5b6003840154610a74908390613226565b811015610ade5760038401610a8983836131fb565b81548110610a9957610a996131d3565b905f5260205f209060020201846003018281548110610aba57610aba6131d3565b5f918252602090912082546002909202019081556001918201549082015501610a64565b505f5b81811015610b235783600301805480610afc57610afc613239565b5f8281526020812060025f1990930192830201818155600190810191909155915501610ae1565b508115610b5e5783546001600160a01b0316610b4857610b433383612266565b610b5e565b8354610b5e906001600160a01b0316338461237d565b4386336001600160a01b03167f02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca9485604051610b9b91815260200190565b60405180910390a4505050505050565b610bb48261096c565b610bbd816123dc565b610bc783836123e6565b50505050565b6001600160a01b0381163314610bf65760405163334bd91960e11b815260040160405180910390fd5b610c008282612487565b505050565b610c0d6121ec565b80610c17816121ad565b610c1f612500565b5f60058381548110610c3357610c336131d3565b5f9182526020808320868452600682526040808520338652909252922060079091029091019150610c6384611074565b5f81600201548260010154670de0b6b3a76400008560030154855f0154610c8a919061324d565b610c949190613264565b610c9e9190613226565b610ca891906131fb565b90505f8115610cdb575f6002840155610cc1338361254a565b905081811015610cdb57610cd58183613226565b60028401555b60038401548354670de0b6b3a764000091610cf59161324d565b610cff9190613264565b6001840155604051818152869033907f34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7906020015b60405180910390a3505050505050565b5f5160206133ba5f395f51905f52610d5b816123dc565b6003805462010000600160b01b031916620100006001600160a01b0385811682029290921792839055604051920416907f861966917ccb4b480494600b4105dff4d2cb14675074324df6acb1281b7a8ac0905f90a25050565b5f5160206133ba5f395f51905f52610dcb816123dc565b610dd36125f4565b50565b610dde6121ec565b5f60055f81548110610df257610df26131d3565b5f918252602090912060079091020180549091506001600160a01b031615610e355760405162461bcd60e51b8152600401610e2c90613283565b60405180910390fd5b60058101543490811015610e8b5760405162461bcd60e51b815260206004820152601b60248201527f6465706f73697420616d6f756e7420697320746f6f20736d616c6c00000000006044820152606401610e2c565b610e955f82612653565b5050565b5f610ea2612935565b805490915060ff600160401b820416159067ffffffffffffffff165f81158015610ec95750825b90505f8267ffffffffffffffff166001148015610ee55750303b155b905081158015610ef3575080155b15610f115760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610f3b57845460ff60401b1916600160401b1785555b868811158015610f4a57505f86115b610f8b5760405162461bcd60e51b8152602060048201526012602482015271696e76616c696420706172616d657465727360701b6044820152606401610e2c565b610f9361295d565b610f9b61295d565b610fa361295d565b610fad5f336123e6565b50610fd87fcab03bc4dbcc648cd59d6bbe9f848d1e9092f914016aa290ee92e18700d1e6f9336123e6565b50610ff05f5160206133ba5f395f51905f52336123e6565b50610ffa89610d44565b5f88905560018790556002869055831561104e57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b611061612965565b61106a82612a09565b610e958282612a33565b8061107e816121ad565b5f60058381548110611092576110926131d3565b905f5260205f2090600702019050806002015443116110b057505050565b5f5f6110ce83600101546110c88560020154436114d6565b90612aef565b91509150816110ef5760405162461bcd60e51b8152600401610e2c906132ba565b60045480151592509004816111165760405162461bcd60e51b8152600401610e2c906132ba565b600483015480156111be575f8061113584670de0b6b3a7640000612aef565b91509150816111565760405162461bcd60e51b8152600401610e2c906132ba565b82151591508290048161117b5760405162461bcd60e51b8152600401610e2c906132ba565b5f5f611194838960030154612b1290919063ffffffff16565b91509150816111b55760405162461bcd60e51b8152600401610e2c906132ba565b60038801555050505b436002850181905560405183815287907ff5d2d72d9b25d6853afd7d0554a113b705234b6a68bb36b7f14366299463241190602001610d34565b5f611201612b22565b505f5160206133da5f395f51905f5290565b5f5160206133ba5f395f51905f5261122a816123dc565b60035460ff166112875760405162461bcd60e51b815260206004820152602260248201527f776974686472617720686173206265656e20616c726561647920756e70617573604482015261195960f21b6064820152608401610e2c565b6003805460ff191690556040517f1c84bcaead48b692cc46b9b12e9a068951a59c99a2e2bf10b00b60b403cf12e2905f90a150565b5f826112c7816121ad565b6112d284844361139a565b949350505050565b5f5160206133ba5f395f51905f526112f1816123dc565b60035460ff16156113445760405162461bcd60e51b815260206004820181905260248201527f776974686472617720686173206265656e20616c7265616479207061757365646044820152606401610e2c565b6003805460ff191660011790556040517f8099f593a6aaecd68b6494933cd71f703376ac3975be83692e1b7d800abf6837905f90a150565b6005545f5b81811015610e955761139281611074565b600101611381565b5f836113a5816121ad565b5f600586815481106113b9576113b96131d3565b5f91825260208083208984526006825260408085206001600160a01b038b168652909252922060036007909202909201908101546004820154600283015492945090918711801561140957508015155b1561146b575f61141d8560020154896114d6565b90505f600454866001015483611433919061324d565b61143d9190613264565b90508261145282670de0b6b3a764000061324d565b61145c9190613264565b61146690856131fb565b935050505b600283015460018401548454670de0b6b3a76400009061148c90869061324d565b6114969190613264565b6114a09190613226565b6114aa91906131fb565b9998505050505050505050565b5f5160206133ba5f395f51905f526114ce816123dc565b610dd3612b6b565b5f818311156115175760405162461bcd60e51b815260206004820152600d60248201526c696e76616c696420626c6f636b60981b6044820152606401610e2c565b5f54831015611525575f5492505b6001548211156115355760015491505b818311156115985760405162461bcd60e51b815260206004820152602a60248201527f656e6420626c6f636b206d7573742062652067726561746572207468616e20736044820152697461727420626c6f636b60b01b6064820152608401610e2c565b5f6115ab60025485856110c89190613226565b92509050806109655760405162461bcd60e51b81526020600482015260136024820152726d756c7469706c696572206f766572666c6f7760681b6044820152606401610e2c565b5f5160206133ba5f395f51905f52611609816123dc565b600354610100900460ff16156116615760405162461bcd60e51b815260206004820152601d60248201527f636c61696d20686173206265656e20616c7265616479207061757365640000006044820152606401610e2c565b6003805461ff0019166101001790556040517f6d73d6b34c378ab3bf6630206d60b7882801b91d03ee20d016ff0d5054db81e1905f90a150565b5f9182525f5160206133fa5f395f51905f52602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6116d96121ec565b816116e3816121ad565b6116eb61221e565b5f600584815481106116ff576116ff6131d3565b5f918252602080832087845260068252604080852033865290925292208054600790920290920192508411156117775760405162461bcd60e51b815260206004820181905260248201527f4e6f7420656e6f756768207374616b696e6720746f6b656e2062616c616e63656044820152606401610e2c565b61178085611074565b5f8160010154670de0b6b3a76400008460030154845f01546117a2919061324d565b6117ac9190613264565b6117b69190613226565b905080156117d4578082600201546117ce91906131fb565b60028301555b84156118395781546117e7908690613226565b8255604080518082019091528581526006840154600384019190602082019061181090436131fb565b90528154600181810184555f938452602093849020835160029093020191825592909101519101555b8483600401546118499190613226565b600484015560038301548254670de0b6b3a7640000916118689161324d565b6118729190613264565b6001830155604051858152869033907fc80277265097707f6f12a4ac4c09d46c9926e2eea2536f63616cb04d9fcad7d690602001610d34565b5f5160206133ba5f395f51905f526118c2816123dc565b600554156118f5576001600160a01b0386166118f05760405162461bcd60e51b8152600401610e2c90613283565b61191c565b6001600160a01b0386161561191c5760405162461bcd60e51b8152600401610e2c90613283565b5f831161196b5760405162461bcd60e51b815260206004820152601e60248201527f696e76616c6964207769746864726177206c6f636b656420626c6f636b7300006044820152606401610e2c565b60015443106119ac5760405162461bcd60e51b815260206004820152600d60248201526c105b1c9958591e48195b991959609a1b6044820152606401610e2c565b81156119ba576119ba61137c565b5f5f5443116119ca575f546119cc565b435b9050856004546119dc91906131fb565b6004556040805160e0810182526001600160a01b0389811680835260208084018b81528486018781525f606087018181526080880182815260a089018f815260c08a018f815260058054600181018255955299517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0600790950294850180546001600160a01b03191691909a161790985593517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db183015591517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db282015590517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db382015590517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db482015592517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db584015592517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db6909201919091558251888152918201879052839289927f0fa296fce13e7a0e622b3a892e66220c248337289483a3cfa4130cde0caa1346910160405180910390a450505050505050565b5f5160206133ba5f395f51905f52611baf816123dc565b815f541115611bd05760405162461bcd60e51b8152600401610e2c906132dc565b600182905560405182907f1132c5baccb51da3d049fabc819697dc845fa224ad59d9b555507d6446b40850905f90a25050565b611c0c8261096c565b611c15816123dc565b610bc78383612487565b5f5160206133ba5f395f51905f52611c36816123dc565b83611c40816121ad565b8360058681548110611c5457611c546131d3565b905f5260205f209060070201600501819055508260058681548110611c7b57611c7b6131d3565b905f5260205f209060070201600601819055508284867f30dffdedaa3e3b4849298233f7cd71d229956e875ab09270498c96b7cf9181fd60405160405180910390a45050505050565b5f5160206133ba5f395f51905f52611cdb816123dc565b600354610100900460ff16611d325760405162461bcd60e51b815260206004820152601f60248201527f636c61696d20686173206265656e20616c726561647920756e706175736564006044820152606401610e2c565b6003805461ff00191690556040517fe72cb12952f056e3e7496019725f20a13108ca420f67f1ee9c9cdab73fb8ce85905f90a150565b5f5160206133ba5f395f51905f52611d7f816123dc565b5f8211611dc25760405162461bcd60e51b815260206004820152601160248201527034b73b30b634b2103830b930b6b2ba32b960791b6044820152606401610e2c565b600282905560405182907fda51cd4dfe7787695fefa2693693bfac4e20eb62914567a95f6672a110406440905f90a25050565b611dfd6121ec565b81611e07816121ad565b825f03611e565760405162461bcd60e51b815260206004820152601f60248201527f6465706f736974206e6f7420737570706f727420455448207374616b696e67006044820152606401610e2c565b5f60058481548110611e6a57611e6a6131d3565b905f5260205f209060070201905080600501548311611ecb5760405162461bcd60e51b815260206004820152601b60248201527f6465706f73697420616d6f756e7420697320746f6f20736d616c6c00000000006044820152606401610e2c565b8215611ee8578054611ee8906001600160a01b0316333086612bb3565b610bc78484612653565b5f5160206133ba5f395f51905f52611f09816123dc565b600154821115611f2b5760405162461bcd60e51b8152600401610e2c906132dc565b5f82815560405183917f63b90b79f11a0f132bcb2c4a4ddd44abda45c1308a83b2919318df7f5f8b7be491a25050565b5f5160206133ba5f395f51905f52611f72816123dc565b83611f7c816121ad565b5f8411611fc15760405162461bcd60e51b81526020600482015260136024820152721a5b9d985b1a59081c1bdbdb081dd95a59da1d606a1b6044820152606401610e2c565b8215611fcf57611fcf61137c565b8360058681548110611fe357611fe36131d3565b905f5260205f209060070201600101546004546120009190613226565b61200a91906131fb565b6004819055508360058681548110612024576120246131d3565b905f5260205f2090600702016001018190555083857f4b8fa3d6a87cb21d1bf4978bf60628ae358a28ac7f39de1751a481c6dd95761760045460405161206c91815260200190565b60405180910390a35050505050565b6005818154811061208a575f80fd5b5f91825260209091206007909102018054600182015460028301546003840154600485015460058601546006909601546001600160a01b039095169650929491939092919087565b5f5f836120de816121ad565b5f8581526006602090815260408083206001600160a01b03881684529091528120905b60038201548110156121a35743826003018281548110612123576121236131d3565b905f5260205f209060020201600101541161216a5781600301818154811061214d5761214d6131d3565b905f5260205f2090600202015f01548461216791906131fb565b93505b81600301818154811061217f5761217f6131d3565b905f5260205f2090600202015f01548561219991906131fb565b9450600101612101565b5050509250929050565b6005548110610dd35760405162461bcd60e51b815260206004820152600b60248201526a1a5b9d985b1a59081c1a5960aa1b6044820152606401610e2c565b5f51602061341a5f395f51905f525460ff161561221c5760405163d93c066560e01b815260040160405180910390fd5b565b60035460ff161561221c5760405162461bcd60e51b81526020600482015260126024820152711dda5d1a191c985dc81a5cc81c185d5cd95960721b6044820152606401610e2c565b5f5f836001600160a01b0316836040515f6040518083038185875af1925050503d805f81146122b0576040519150601f19603f3d011682016040523d82523d5f602084013e6122b5565b606091505b5091509150816123075760405162461bcd60e51b815260206004820152601860248201527f455448207472616e736665722063616c6c206661696c656400000000000000006044820152606401610e2c565b805115610bc757808060200190518101906123229190613326565b610bc75760405162461bcd60e51b815260206004820152602660248201527f455448207472616e73666572206f7065726174696f6e20646964206e6f7420736044820152651d58d8d9595960d21b6064820152608401610e2c565b6040516001600160a01b03838116602483015260448201839052610c0091859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050612bec565b610dd38133612c58565b5f5f5160206133fa5f395f51905f526123ff848461169b565b61247e575f848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556124343390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061092f565b5f91505061092f565b5f5f5160206133fa5f395f51905f526124a0848461169b565b1561247e575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061092f565b600354610100900460ff161561221c5760405162461bcd60e51b815260206004820152600f60248201526e18db185a5b481a5cc81c185d5cd959608a1b6044820152606401610e2c565b6003546040516370a0823160e01b81523060048201525f918291620100009091046001600160a01b0316906370a0823190602401602060405180830381865afa158015612599573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906125bd9190613341565b90505f8184116125cd57836125cf565b815b905080156112d2576003546112d2906201000090046001600160a01b0316868361237d565b6125fc612c91565b5f51602061341a5f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b5f60058381548110612667576126676131d3565b5f918252602080832086845260068252604080852033865290925292206007909102909101915061269784611074565b8054156127d457600382015481545f9182916126b291612aef565b91509150816126d35760405162461bcd60e51b8152600401610e2c90613358565b60019150670de0b6b3a764000090045f5f6126fb856001015484612cc090919063ffffffff16565b91509150816127585760405162461bcd60e51b815260206004820152602360248201527f6163635374616b65207375622066696e697368656446756b7561206f766572666044820152626c6f7760e81b6064820152608401610e2c565b80156127cf575f5f612777838860020154612b1290919063ffffffff16565b91509150816127c85760405162461bcd60e51b815260206004820152601a60248201527f757365722070656e64696e6746756b7561206f766572666c6f770000000000006044820152606401610e2c565b6002870155505b505050505b82156128385780545f9081906127ea9086612b12565b91509150816128345760405162461bcd60e51b815260206004820152601660248201527575736572207374416d6f756e74206f766572666c6f7760501b6044820152606401610e2c565b8255505b5f5f612851858560040154612b1290919063ffffffff16565b91509150816128a25760405162461bcd60e51b815260206004820152601b60248201527f706f6f6c207374546f6b656e416d6f756e74206f766572666c6f7700000000006044820152606401610e2c565b60048401819055600384015483545f9182916128bd91612aef565b91509150816128de5760405162461bcd60e51b8152600401610e2c90613358565b60019150670de0b6b3a7640000900460018501819055604051878152889033907f90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a159060200160405180910390a35050505050505050565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0061092f565b61221c612cd0565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806129eb57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166129df5f5160206133da5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b1561221c5760405163703e46dd60e11b815260040160405180910390fd5b7fcab03bc4dbcc648cd59d6bbe9f848d1e9092f914016aa290ee92e18700d1e6f9610e95816123dc565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612a8d575060408051601f3d908101601f19168201909252612a8a91810190613341565b60015b612ab557604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610e2c565b5f5160206133da5f395f51905f528114612ae557604051632a87526960e21b815260048101829052602401610e2c565b610c008383612cf5565b8181028281048214831517905f90612b0683151590565b81029150509250929050565b80820182811015905f9082612b06565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461221c5760405163703e46dd60e11b815260040160405180910390fd5b612b736121ec565b5f51602061341a5f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833612635565b6040516001600160a01b038481166024830152838116604483015260648201839052610bc79186918216906323b872dd906084016123aa565b5f5f60205f8451602086015f885af180612c0b576040513d5f823e3d81fd5b50505f513d91508115612c22578060011415612c2f565b6001600160a01b0384163b155b15610bc757604051635274afe760e01b81526001600160a01b0385166004820152602401610e2c565b612c62828261169b565b610e955760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610e2c565b5f51602061341a5f395f51905f525460ff1661221c57604051638dfc202b60e01b815260040160405180910390fd5b80820382811115905f9082612b06565b612cd8612d4a565b61221c57604051631afcd79f60e31b815260040160405180910390fd5b612cfe82612d63565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612d4257610c008282612dc6565b610e95612e38565b5f612d53612935565b54600160401b900460ff16919050565b806001600160a01b03163b5f03612d9857604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610e2c565b5f5160206133da5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051612de291906133a3565b5f60405180830381855af49150503d805f8114612e1a576040519150601f19603f3d011682016040523d82523d5f602084013e612e1f565b606091505b5091509150612e2f858383612e57565b95945050505050565b341561221c5760405163b398979f60e01b815260040160405180910390fd5b606082612e6c57612e6782612eb6565b612eaf565b8151158015612e8357506001600160a01b0384163b155b15612eac57604051639996b31560e01b81526001600160a01b0385166004820152602401610e2c565b50805b9392505050565b805115612ec557805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b5f60208284031215612eee575f5ffd5b81356001600160e01b031981168114612eaf575f5ffd5b6001600160a01b0381168114610dd3575f5ffd5b5f5f60408385031215612f2a575f5ffd5b823591506020830135612f3c81612f05565b809150509250929050565b5f60208284031215612f57575f5ffd5b5035919050565b5f60208284031215612f6e575f5ffd5b8135612eaf81612f05565b5f5f5f5f60808587031215612f8c575f5ffd5b8435612f9781612f05565b966020860135965060408601359560600135945092505050565b634e487b7160e01b5f52604160045260245ffd5b5f5f60408385031215612fd6575f5ffd5b8235612fe181612f05565b9150602083013567ffffffffffffffff811115612ffc575f5ffd5b8301601f8101851361300c575f5ffd5b803567ffffffffffffffff81111561302657613026612fb1565b604051601f8201601f19908116603f0116810167ffffffffffffffff8111828210171561305557613055612fb1565b60405281815282820160200187101561306c575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f5f5f6060848603121561309d575f5ffd5b8335925060208401356130af81612f05565b929592945050506040919091013590565b5f5f604083850312156130d1575f5ffd5b50508035926020909101359150565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b8015158114610dd3575f5ffd5b5f5f5f5f5f60a08688031215613136575f5ffd5b853561314181612f05565b9450602086013593506040860135925060608601359150608086013561316681613115565b809150509295509295909350565b5f5f5f60608486031215613186575f5ffd5b505081359360208301359350604090920135919050565b5f5f5f606084860312156131af575f5ffd5b833592506020840135915060408401356131c881613115565b809150509250925092565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b8082018082111561092f5761092f6131e7565b5f6001820161321f5761321f6131e7565b5060010190565b8181038181111561092f5761092f6131e7565b634e487b7160e01b5f52603160045260245ffd5b808202811582820484141761092f5761092f6131e7565b5f8261327e57634e487b7160e01b5f52601260045260245ffd5b500490565b6020808252601d908201527f696e76616c6964207374616b696e6720746f6b656e2061646472657373000000604082015260600190565b6020808252600890820152676f766572666c6f7760c01b604082015260600190565b6020808252602a908201527f737461727420626c6f636b206d75737420626520736d616c6c6572207468616e60408201526920656e6420626c6f636b60b01b606082015260800190565b5f60208284031215613336575f5ffd5b8151612eaf81613115565b5f60208284031215613351575f5ffd5b5051919050565b6020808252602b908201527f75736572207374416d6f756e74206d756c2061636346756b756150657253686160408201526a7265206f766572666c6f7760a81b606082015260800190565b5f82518060208501845e5f92019182525091905056fe589d473ba17c0f47d494622893831497bad25919b9afb8e33e9521b8963fccde360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300a2646970667358221220c799ef699ed69a71f34e32ce9fc5023226d5e987b39bf37856a1bd8b66cd049564736f6c634300081e0033",
}

// FukuastakeABI is the input ABI used to generate the binding from.
// Deprecated: Use FukuastakeMetaData.ABI instead.
var FukuastakeABI = FukuastakeMetaData.ABI

// FukuastakeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FukuastakeMetaData.Bin instead.
var FukuastakeBin = FukuastakeMetaData.Bin

// DeployFukuastake deploys a new Ethereum contract, binding an instance of Fukuastake to it.
func DeployFukuastake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Fukuastake, error) {
	parsed, err := FukuastakeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FukuastakeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Fukuastake{FukuastakeCaller: FukuastakeCaller{contract: contract}, FukuastakeTransactor: FukuastakeTransactor{contract: contract}, FukuastakeFilterer: FukuastakeFilterer{contract: contract}}, nil
}

// Fukuastake is an auto generated Go binding around an Ethereum contract.
type Fukuastake struct {
	FukuastakeCaller     // Read-only binding to the contract
	FukuastakeTransactor // Write-only binding to the contract
	FukuastakeFilterer   // Log filterer for contract events
}

// FukuastakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type FukuastakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FukuastakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FukuastakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FukuastakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FukuastakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FukuastakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FukuastakeSession struct {
	Contract     *Fukuastake       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FukuastakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FukuastakeCallerSession struct {
	Contract *FukuastakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FukuastakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FukuastakeTransactorSession struct {
	Contract     *FukuastakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FukuastakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type FukuastakeRaw struct {
	Contract *Fukuastake // Generic contract binding to access the raw methods on
}

// FukuastakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FukuastakeCallerRaw struct {
	Contract *FukuastakeCaller // Generic read-only contract binding to access the raw methods on
}

// FukuastakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FukuastakeTransactorRaw struct {
	Contract *FukuastakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFukuastake creates a new instance of Fukuastake, bound to a specific deployed contract.
func NewFukuastake(address common.Address, backend bind.ContractBackend) (*Fukuastake, error) {
	contract, err := bindFukuastake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fukuastake{FukuastakeCaller: FukuastakeCaller{contract: contract}, FukuastakeTransactor: FukuastakeTransactor{contract: contract}, FukuastakeFilterer: FukuastakeFilterer{contract: contract}}, nil
}

// NewFukuastakeCaller creates a new read-only instance of Fukuastake, bound to a specific deployed contract.
func NewFukuastakeCaller(address common.Address, caller bind.ContractCaller) (*FukuastakeCaller, error) {
	contract, err := bindFukuastake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FukuastakeCaller{contract: contract}, nil
}

// NewFukuastakeTransactor creates a new write-only instance of Fukuastake, bound to a specific deployed contract.
func NewFukuastakeTransactor(address common.Address, transactor bind.ContractTransactor) (*FukuastakeTransactor, error) {
	contract, err := bindFukuastake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FukuastakeTransactor{contract: contract}, nil
}

// NewFukuastakeFilterer creates a new log filterer instance of Fukuastake, bound to a specific deployed contract.
func NewFukuastakeFilterer(address common.Address, filterer bind.ContractFilterer) (*FukuastakeFilterer, error) {
	contract, err := bindFukuastake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FukuastakeFilterer{contract: contract}, nil
}

// bindFukuastake binds a generic wrapper to an already deployed contract.
func bindFukuastake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FukuastakeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fukuastake *FukuastakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fukuastake.Contract.FukuastakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fukuastake *FukuastakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fukuastake.Contract.FukuastakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fukuastake *FukuastakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fukuastake.Contract.FukuastakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fukuastake *FukuastakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fukuastake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fukuastake *FukuastakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fukuastake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fukuastake *FukuastakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fukuastake.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Fukuastake *FukuastakeCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Fukuastake *FukuastakeSession) ADMINROLE() ([32]byte, error) {
	return _Fukuastake.Contract.ADMINROLE(&_Fukuastake.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Fukuastake *FukuastakeCallerSession) ADMINROLE() ([32]byte, error) {
	return _Fukuastake.Contract.ADMINROLE(&_Fukuastake.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Fukuastake *FukuastakeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Fukuastake *FukuastakeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Fukuastake.Contract.DEFAULTADMINROLE(&_Fukuastake.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Fukuastake *FukuastakeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Fukuastake.Contract.DEFAULTADMINROLE(&_Fukuastake.CallOpts)
}

// ETHPID is a free data retrieval call binding the contract method 0xbfc3ebba.
//
// Solidity: function ETH_PID() view returns(uint256)
func (_Fukuastake *FukuastakeCaller) ETHPID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "ETH_PID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ETHPID is a free data retrieval call binding the contract method 0xbfc3ebba.
//
// Solidity: function ETH_PID() view returns(uint256)
func (_Fukuastake *FukuastakeSession) ETHPID() (*big.Int, error) {
	return _Fukuastake.Contract.ETHPID(&_Fukuastake.CallOpts)
}

// ETHPID is a free data retrieval call binding the contract method 0xbfc3ebba.
//
// Solidity: function ETH_PID() view returns(uint256)
func (_Fukuastake *FukuastakeCallerSession) ETHPID() (*big.Int, error) {
	return _Fukuastake.Contract.ETHPID(&_Fukuastake.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Fukuastake *FukuastakeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Fukuastake *FukuastakeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Fukuastake.Contract.UPGRADEINTERFACEVERSION(&_Fukuastake.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Fukuastake *FukuastakeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Fukuastake.Contract.UPGRADEINTERFACEVERSION(&_Fukuastake.CallOpts)
}

// UPGRADEROLE is a free data retrieval call binding the contract method 0xb908afa8.
//
// Solidity: function UPGRADE_ROLE() view returns(bytes32)
func (_Fukuastake *FukuastakeCaller) UPGRADEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "UPGRADE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADEROLE is a free data retrieval call binding the contract method 0xb908afa8.
//
// Solidity: function UPGRADE_ROLE() view returns(bytes32)
func (_Fukuastake *FukuastakeSession) UPGRADEROLE() ([32]byte, error) {
	return _Fukuastake.Contract.UPGRADEROLE(&_Fukuastake.CallOpts)
}

// UPGRADEROLE is a free data retrieval call binding the contract method 0xb908afa8.
//
// Solidity: function UPGRADE_ROLE() view returns(bytes32)
func (_Fukuastake *FukuastakeCallerSession) UPGRADEROLE() ([32]byte, error) {
	return _Fukuastake.Contract.UPGRADEROLE(&_Fukuastake.CallOpts)
}

// ClaimPaused is a free data retrieval call binding the contract method 0xab5e124a.
//
// Solidity: function claimPaused() view returns(bool)
func (_Fukuastake *FukuastakeCaller) ClaimPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "claimPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ClaimPaused is a free data retrieval call binding the contract method 0xab5e124a.
//
// Solidity: function claimPaused() view returns(bool)
func (_Fukuastake *FukuastakeSession) ClaimPaused() (bool, error) {
	return _Fukuastake.Contract.ClaimPaused(&_Fukuastake.CallOpts)
}

// ClaimPaused is a free data retrieval call binding the contract method 0xab5e124a.
//
// Solidity: function claimPaused() view returns(bool)
func (_Fukuastake *FukuastakeCallerSession) ClaimPaused() (bool, error) {
	return _Fukuastake.Contract.ClaimPaused(&_Fukuastake.CallOpts)
}

// EndBlock is a free data retrieval call binding the contract method 0x083c6323.
//
// Solidity: function endBlock() view returns(uint256)
func (_Fukuastake *FukuastakeCaller) EndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "endBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndBlock is a free data retrieval call binding the contract method 0x083c6323.
//
// Solidity: function endBlock() view returns(uint256)
func (_Fukuastake *FukuastakeSession) EndBlock() (*big.Int, error) {
	return _Fukuastake.Contract.EndBlock(&_Fukuastake.CallOpts)
}

// EndBlock is a free data retrieval call binding the contract method 0x083c6323.
//
// Solidity: function endBlock() view returns(uint256)
func (_Fukuastake *FukuastakeCallerSession) EndBlock() (*big.Int, error) {
	return _Fukuastake.Contract.EndBlock(&_Fukuastake.CallOpts)
}

// FukuaPerBlock is a free data retrieval call binding the contract method 0xa2e8fca2.
//
// Solidity: function fukuaPerBlock() view returns(uint256)
func (_Fukuastake *FukuastakeCaller) FukuaPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "fukuaPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FukuaPerBlock is a free data retrieval call binding the contract method 0xa2e8fca2.
//
// Solidity: function fukuaPerBlock() view returns(uint256)
func (_Fukuastake *FukuastakeSession) FukuaPerBlock() (*big.Int, error) {
	return _Fukuastake.Contract.FukuaPerBlock(&_Fukuastake.CallOpts)
}

// FukuaPerBlock is a free data retrieval call binding the contract method 0xa2e8fca2.
//
// Solidity: function fukuaPerBlock() view returns(uint256)
func (_Fukuastake *FukuastakeCallerSession) FukuaPerBlock() (*big.Int, error) {
	return _Fukuastake.Contract.FukuaPerBlock(&_Fukuastake.CallOpts)
}

// FukuaToken is a free data retrieval call binding the contract method 0x9fac8a6a.
//
// Solidity: function fukuaToken() view returns(address)
func (_Fukuastake *FukuastakeCaller) FukuaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "fukuaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FukuaToken is a free data retrieval call binding the contract method 0x9fac8a6a.
//
// Solidity: function fukuaToken() view returns(address)
func (_Fukuastake *FukuastakeSession) FukuaToken() (common.Address, error) {
	return _Fukuastake.Contract.FukuaToken(&_Fukuastake.CallOpts)
}

// FukuaToken is a free data retrieval call binding the contract method 0x9fac8a6a.
//
// Solidity: function fukuaToken() view returns(address)
func (_Fukuastake *FukuastakeCallerSession) FukuaToken() (common.Address, error) {
	return _Fukuastake.Contract.FukuaToken(&_Fukuastake.CallOpts)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256 multiplier)
func (_Fukuastake *FukuastakeCaller) GetMultiplier(opts *bind.CallOpts, _from *big.Int, _to *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "getMultiplier", _from, _to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256 multiplier)
func (_Fukuastake *FukuastakeSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _Fukuastake.Contract.GetMultiplier(&_Fukuastake.CallOpts, _from, _to)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256 multiplier)
func (_Fukuastake *FukuastakeCallerSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _Fukuastake.Contract.GetMultiplier(&_Fukuastake.CallOpts, _from, _to)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Fukuastake *FukuastakeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Fukuastake *FukuastakeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Fukuastake.Contract.GetRoleAdmin(&_Fukuastake.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Fukuastake *FukuastakeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Fukuastake.Contract.GetRoleAdmin(&_Fukuastake.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Fukuastake *FukuastakeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Fukuastake *FukuastakeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Fukuastake.Contract.HasRole(&_Fukuastake.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Fukuastake *FukuastakeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Fukuastake.Contract.HasRole(&_Fukuastake.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Fukuastake *FukuastakeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Fukuastake *FukuastakeSession) Paused() (bool, error) {
	return _Fukuastake.Contract.Paused(&_Fukuastake.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Fukuastake *FukuastakeCallerSession) Paused() (bool, error) {
	return _Fukuastake.Contract.Paused(&_Fukuastake.CallOpts)
}

// PendingFukua is a free data retrieval call binding the contract method 0x5c0ab30b.
//
// Solidity: function pendingFukua(uint256 _pid, address _user) view returns(uint256)
func (_Fukuastake *FukuastakeCaller) PendingFukua(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "pendingFukua", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingFukua is a free data retrieval call binding the contract method 0x5c0ab30b.
//
// Solidity: function pendingFukua(uint256 _pid, address _user) view returns(uint256)
func (_Fukuastake *FukuastakeSession) PendingFukua(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Fukuastake.Contract.PendingFukua(&_Fukuastake.CallOpts, _pid, _user)
}

// PendingFukua is a free data retrieval call binding the contract method 0x5c0ab30b.
//
// Solidity: function pendingFukua(uint256 _pid, address _user) view returns(uint256)
func (_Fukuastake *FukuastakeCallerSession) PendingFukua(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Fukuastake.Contract.PendingFukua(&_Fukuastake.CallOpts, _pid, _user)
}

// PendingFukuaByBlockNumber is a free data retrieval call binding the contract method 0x7ef64c46.
//
// Solidity: function pendingFukuaByBlockNumber(uint256 _pid, address _user, uint256 _blockNumber) view returns(uint256)
func (_Fukuastake *FukuastakeCaller) PendingFukuaByBlockNumber(opts *bind.CallOpts, _pid *big.Int, _user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "pendingFukuaByBlockNumber", _pid, _user, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingFukuaByBlockNumber is a free data retrieval call binding the contract method 0x7ef64c46.
//
// Solidity: function pendingFukuaByBlockNumber(uint256 _pid, address _user, uint256 _blockNumber) view returns(uint256)
func (_Fukuastake *FukuastakeSession) PendingFukuaByBlockNumber(_pid *big.Int, _user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _Fukuastake.Contract.PendingFukuaByBlockNumber(&_Fukuastake.CallOpts, _pid, _user, _blockNumber)
}

// PendingFukuaByBlockNumber is a free data retrieval call binding the contract method 0x7ef64c46.
//
// Solidity: function pendingFukuaByBlockNumber(uint256 _pid, address _user, uint256 _blockNumber) view returns(uint256)
func (_Fukuastake *FukuastakeCallerSession) PendingFukuaByBlockNumber(_pid *big.Int, _user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _Fukuastake.Contract.PendingFukuaByBlockNumber(&_Fukuastake.CallOpts, _pid, _user, _blockNumber)
}

// Pool is a free data retrieval call binding the contract method 0xfe313112.
//
// Solidity: function pool(uint256 ) view returns(address stTokenAddress, uint256 poolWeight, uint256 lastRewardBlock, uint256 accFukuaPerShare, uint256 stTokenAmount, uint256 minDepositAmount, uint256 unstakeLockedBlocks)
func (_Fukuastake *FukuastakeCaller) Pool(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StTokenAddress      common.Address
	PoolWeight          *big.Int
	LastRewardBlock     *big.Int
	AccFukuaPerShare    *big.Int
	StTokenAmount       *big.Int
	MinDepositAmount    *big.Int
	UnstakeLockedBlocks *big.Int
}, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "pool", arg0)

	outstruct := new(struct {
		StTokenAddress      common.Address
		PoolWeight          *big.Int
		LastRewardBlock     *big.Int
		AccFukuaPerShare    *big.Int
		StTokenAmount       *big.Int
		MinDepositAmount    *big.Int
		UnstakeLockedBlocks *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StTokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PoolWeight = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccFukuaPerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StTokenAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.MinDepositAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.UnstakeLockedBlocks = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Pool is a free data retrieval call binding the contract method 0xfe313112.
//
// Solidity: function pool(uint256 ) view returns(address stTokenAddress, uint256 poolWeight, uint256 lastRewardBlock, uint256 accFukuaPerShare, uint256 stTokenAmount, uint256 minDepositAmount, uint256 unstakeLockedBlocks)
func (_Fukuastake *FukuastakeSession) Pool(arg0 *big.Int) (struct {
	StTokenAddress      common.Address
	PoolWeight          *big.Int
	LastRewardBlock     *big.Int
	AccFukuaPerShare    *big.Int
	StTokenAmount       *big.Int
	MinDepositAmount    *big.Int
	UnstakeLockedBlocks *big.Int
}, error) {
	return _Fukuastake.Contract.Pool(&_Fukuastake.CallOpts, arg0)
}

// Pool is a free data retrieval call binding the contract method 0xfe313112.
//
// Solidity: function pool(uint256 ) view returns(address stTokenAddress, uint256 poolWeight, uint256 lastRewardBlock, uint256 accFukuaPerShare, uint256 stTokenAmount, uint256 minDepositAmount, uint256 unstakeLockedBlocks)
func (_Fukuastake *FukuastakeCallerSession) Pool(arg0 *big.Int) (struct {
	StTokenAddress      common.Address
	PoolWeight          *big.Int
	LastRewardBlock     *big.Int
	AccFukuaPerShare    *big.Int
	StTokenAmount       *big.Int
	MinDepositAmount    *big.Int
	UnstakeLockedBlocks *big.Int
}, error) {
	return _Fukuastake.Contract.Pool(&_Fukuastake.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Fukuastake *FukuastakeCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Fukuastake *FukuastakeSession) PoolLength() (*big.Int, error) {
	return _Fukuastake.Contract.PoolLength(&_Fukuastake.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Fukuastake *FukuastakeCallerSession) PoolLength() (*big.Int, error) {
	return _Fukuastake.Contract.PoolLength(&_Fukuastake.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Fukuastake *FukuastakeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Fukuastake *FukuastakeSession) ProxiableUUID() ([32]byte, error) {
	return _Fukuastake.Contract.ProxiableUUID(&_Fukuastake.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Fukuastake *FukuastakeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Fukuastake.Contract.ProxiableUUID(&_Fukuastake.CallOpts)
}

// StakingBalance is a free data retrieval call binding the contract method 0x11548234.
//
// Solidity: function stakingBalance(uint256 _pid, address _user) view returns(uint256)
func (_Fukuastake *FukuastakeCaller) StakingBalance(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "stakingBalance", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakingBalance is a free data retrieval call binding the contract method 0x11548234.
//
// Solidity: function stakingBalance(uint256 _pid, address _user) view returns(uint256)
func (_Fukuastake *FukuastakeSession) StakingBalance(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Fukuastake.Contract.StakingBalance(&_Fukuastake.CallOpts, _pid, _user)
}

// StakingBalance is a free data retrieval call binding the contract method 0x11548234.
//
// Solidity: function stakingBalance(uint256 _pid, address _user) view returns(uint256)
func (_Fukuastake *FukuastakeCallerSession) StakingBalance(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Fukuastake.Contract.StakingBalance(&_Fukuastake.CallOpts, _pid, _user)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Fukuastake *FukuastakeCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Fukuastake *FukuastakeSession) StartBlock() (*big.Int, error) {
	return _Fukuastake.Contract.StartBlock(&_Fukuastake.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Fukuastake *FukuastakeCallerSession) StartBlock() (*big.Int, error) {
	return _Fukuastake.Contract.StartBlock(&_Fukuastake.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Fukuastake *FukuastakeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Fukuastake *FukuastakeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Fukuastake.Contract.SupportsInterface(&_Fukuastake.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Fukuastake *FukuastakeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Fukuastake.Contract.SupportsInterface(&_Fukuastake.CallOpts, interfaceId)
}

// TotalPoolWeight is a free data retrieval call binding the contract method 0x02559004.
//
// Solidity: function totalPoolWeight() view returns(uint256)
func (_Fukuastake *FukuastakeCaller) TotalPoolWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "totalPoolWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPoolWeight is a free data retrieval call binding the contract method 0x02559004.
//
// Solidity: function totalPoolWeight() view returns(uint256)
func (_Fukuastake *FukuastakeSession) TotalPoolWeight() (*big.Int, error) {
	return _Fukuastake.Contract.TotalPoolWeight(&_Fukuastake.CallOpts)
}

// TotalPoolWeight is a free data retrieval call binding the contract method 0x02559004.
//
// Solidity: function totalPoolWeight() view returns(uint256)
func (_Fukuastake *FukuastakeCallerSession) TotalPoolWeight() (*big.Int, error) {
	return _Fukuastake.Contract.TotalPoolWeight(&_Fukuastake.CallOpts)
}

// User is a free data retrieval call binding the contract method 0x37849b3c.
//
// Solidity: function user(uint256 , address ) view returns(uint256 stAmount, uint256 finishedFukua, uint256 pendingFukua)
func (_Fukuastake *FukuastakeCaller) User(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	StAmount      *big.Int
	FinishedFukua *big.Int
	PendingFukua  *big.Int
}, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "user", arg0, arg1)

	outstruct := new(struct {
		StAmount      *big.Int
		FinishedFukua *big.Int
		PendingFukua  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FinishedFukua = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PendingFukua = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// User is a free data retrieval call binding the contract method 0x37849b3c.
//
// Solidity: function user(uint256 , address ) view returns(uint256 stAmount, uint256 finishedFukua, uint256 pendingFukua)
func (_Fukuastake *FukuastakeSession) User(arg0 *big.Int, arg1 common.Address) (struct {
	StAmount      *big.Int
	FinishedFukua *big.Int
	PendingFukua  *big.Int
}, error) {
	return _Fukuastake.Contract.User(&_Fukuastake.CallOpts, arg0, arg1)
}

// User is a free data retrieval call binding the contract method 0x37849b3c.
//
// Solidity: function user(uint256 , address ) view returns(uint256 stAmount, uint256 finishedFukua, uint256 pendingFukua)
func (_Fukuastake *FukuastakeCallerSession) User(arg0 *big.Int, arg1 common.Address) (struct {
	StAmount      *big.Int
	FinishedFukua *big.Int
	PendingFukua  *big.Int
}, error) {
	return _Fukuastake.Contract.User(&_Fukuastake.CallOpts, arg0, arg1)
}

// WithdrawAmount is a free data retrieval call binding the contract method 0xff423357.
//
// Solidity: function withdrawAmount(uint256 _pid, address _user) view returns(uint256 requestAmount, uint256 pendingWithdrawAmount)
func (_Fukuastake *FukuastakeCaller) WithdrawAmount(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (struct {
	RequestAmount         *big.Int
	PendingWithdrawAmount *big.Int
}, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "withdrawAmount", _pid, _user)

	outstruct := new(struct {
		RequestAmount         *big.Int
		PendingWithdrawAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PendingWithdrawAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// WithdrawAmount is a free data retrieval call binding the contract method 0xff423357.
//
// Solidity: function withdrawAmount(uint256 _pid, address _user) view returns(uint256 requestAmount, uint256 pendingWithdrawAmount)
func (_Fukuastake *FukuastakeSession) WithdrawAmount(_pid *big.Int, _user common.Address) (struct {
	RequestAmount         *big.Int
	PendingWithdrawAmount *big.Int
}, error) {
	return _Fukuastake.Contract.WithdrawAmount(&_Fukuastake.CallOpts, _pid, _user)
}

// WithdrawAmount is a free data retrieval call binding the contract method 0xff423357.
//
// Solidity: function withdrawAmount(uint256 _pid, address _user) view returns(uint256 requestAmount, uint256 pendingWithdrawAmount)
func (_Fukuastake *FukuastakeCallerSession) WithdrawAmount(_pid *big.Int, _user common.Address) (struct {
	RequestAmount         *big.Int
	PendingWithdrawAmount *big.Int
}, error) {
	return _Fukuastake.Contract.WithdrawAmount(&_Fukuastake.CallOpts, _pid, _user)
}

// WithdrawPaused is a free data retrieval call binding the contract method 0x2f3ffb9f.
//
// Solidity: function withdrawPaused() view returns(bool)
func (_Fukuastake *FukuastakeCaller) WithdrawPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Fukuastake.contract.Call(opts, &out, "withdrawPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawPaused is a free data retrieval call binding the contract method 0x2f3ffb9f.
//
// Solidity: function withdrawPaused() view returns(bool)
func (_Fukuastake *FukuastakeSession) WithdrawPaused() (bool, error) {
	return _Fukuastake.Contract.WithdrawPaused(&_Fukuastake.CallOpts)
}

// WithdrawPaused is a free data retrieval call binding the contract method 0x2f3ffb9f.
//
// Solidity: function withdrawPaused() view returns(bool)
func (_Fukuastake *FukuastakeCallerSession) WithdrawPaused() (bool, error) {
	return _Fukuastake.Contract.WithdrawPaused(&_Fukuastake.CallOpts)
}

// AddPool is a paid mutator transaction binding the contract method 0xb6d9d919.
//
// Solidity: function addPool(address _stTokenAddress, uint256 _poolWeight, uint256 _minDepositAmount, uint256 _unstakeLockedBlocks, bool _withUpdate) returns()
func (_Fukuastake *FukuastakeTransactor) AddPool(opts *bind.TransactOpts, _stTokenAddress common.Address, _poolWeight *big.Int, _minDepositAmount *big.Int, _unstakeLockedBlocks *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "addPool", _stTokenAddress, _poolWeight, _minDepositAmount, _unstakeLockedBlocks, _withUpdate)
}

// AddPool is a paid mutator transaction binding the contract method 0xb6d9d919.
//
// Solidity: function addPool(address _stTokenAddress, uint256 _poolWeight, uint256 _minDepositAmount, uint256 _unstakeLockedBlocks, bool _withUpdate) returns()
func (_Fukuastake *FukuastakeSession) AddPool(_stTokenAddress common.Address, _poolWeight *big.Int, _minDepositAmount *big.Int, _unstakeLockedBlocks *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Fukuastake.Contract.AddPool(&_Fukuastake.TransactOpts, _stTokenAddress, _poolWeight, _minDepositAmount, _unstakeLockedBlocks, _withUpdate)
}

// AddPool is a paid mutator transaction binding the contract method 0xb6d9d919.
//
// Solidity: function addPool(address _stTokenAddress, uint256 _poolWeight, uint256 _minDepositAmount, uint256 _unstakeLockedBlocks, bool _withUpdate) returns()
func (_Fukuastake *FukuastakeTransactorSession) AddPool(_stTokenAddress common.Address, _poolWeight *big.Int, _minDepositAmount *big.Int, _unstakeLockedBlocks *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Fukuastake.Contract.AddPool(&_Fukuastake.TransactOpts, _stTokenAddress, _poolWeight, _minDepositAmount, _unstakeLockedBlocks, _withUpdate)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _pid) returns()
func (_Fukuastake *FukuastakeTransactor) Claim(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "claim", _pid)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _pid) returns()
func (_Fukuastake *FukuastakeSession) Claim(_pid *big.Int) (*types.Transaction, error) {
	return _Fukuastake.Contract.Claim(&_Fukuastake.TransactOpts, _pid)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _pid) returns()
func (_Fukuastake *FukuastakeTransactorSession) Claim(_pid *big.Int) (*types.Transaction, error) {
	return _Fukuastake.Contract.Claim(&_Fukuastake.TransactOpts, _pid)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Fukuastake *FukuastakeTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Fukuastake *FukuastakeSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Fukuastake.Contract.Deposit(&_Fukuastake.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Fukuastake *FukuastakeTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Fukuastake.Contract.Deposit(&_Fukuastake.TransactOpts, _pid, _amount)
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns()
func (_Fukuastake *FukuastakeTransactor) DepositEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "depositEth")
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns()
func (_Fukuastake *FukuastakeSession) DepositEth() (*types.Transaction, error) {
	return _Fukuastake.Contract.DepositEth(&_Fukuastake.TransactOpts)
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns()
func (_Fukuastake *FukuastakeTransactorSession) DepositEth() (*types.Transaction, error) {
	return _Fukuastake.Contract.DepositEth(&_Fukuastake.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Fukuastake *FukuastakeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Fukuastake *FukuastakeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fukuastake.Contract.GrantRole(&_Fukuastake.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Fukuastake *FukuastakeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fukuastake.Contract.GrantRole(&_Fukuastake.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address _fukuaToken, uint256 _startBlock, uint256 _endBlock, uint256 _fukuaPerBlock) returns()
func (_Fukuastake *FukuastakeTransactor) Initialize(opts *bind.TransactOpts, _fukuaToken common.Address, _startBlock *big.Int, _endBlock *big.Int, _fukuaPerBlock *big.Int) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "initialize", _fukuaToken, _startBlock, _endBlock, _fukuaPerBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address _fukuaToken, uint256 _startBlock, uint256 _endBlock, uint256 _fukuaPerBlock) returns()
func (_Fukuastake *FukuastakeSession) Initialize(_fukuaToken common.Address, _startBlock *big.Int, _endBlock *big.Int, _fukuaPerBlock *big.Int) (*types.Transaction, error) {
	return _Fukuastake.Contract.Initialize(&_Fukuastake.TransactOpts, _fukuaToken, _startBlock, _endBlock, _fukuaPerBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address _fukuaToken, uint256 _startBlock, uint256 _endBlock, uint256 _fukuaPerBlock) returns()
func (_Fukuastake *FukuastakeTransactorSession) Initialize(_fukuaToken common.Address, _startBlock *big.Int, _endBlock *big.Int, _fukuaPerBlock *big.Int) (*types.Transaction, error) {
	return _Fukuastake.Contract.Initialize(&_Fukuastake.TransactOpts, _fukuaToken, _startBlock, _endBlock, _fukuaPerBlock)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Fukuastake *FukuastakeTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Fukuastake *FukuastakeSession) MassUpdatePools() (*types.Transaction, error) {
	return _Fukuastake.Contract.MassUpdatePools(&_Fukuastake.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Fukuastake *FukuastakeTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _Fukuastake.Contract.MassUpdatePools(&_Fukuastake.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Fukuastake *FukuastakeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Fukuastake *FukuastakeSession) Pause() (*types.Transaction, error) {
	return _Fukuastake.Contract.Pause(&_Fukuastake.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Fukuastake *FukuastakeTransactorSession) Pause() (*types.Transaction, error) {
	return _Fukuastake.Contract.Pause(&_Fukuastake.TransactOpts)
}

// PauseClaim is a paid mutator transaction binding the contract method 0x8ff095f9.
//
// Solidity: function pauseClaim() returns()
func (_Fukuastake *FukuastakeTransactor) PauseClaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "pauseClaim")
}

// PauseClaim is a paid mutator transaction binding the contract method 0x8ff095f9.
//
// Solidity: function pauseClaim() returns()
func (_Fukuastake *FukuastakeSession) PauseClaim() (*types.Transaction, error) {
	return _Fukuastake.Contract.PauseClaim(&_Fukuastake.TransactOpts)
}

// PauseClaim is a paid mutator transaction binding the contract method 0x8ff095f9.
//
// Solidity: function pauseClaim() returns()
func (_Fukuastake *FukuastakeTransactorSession) PauseClaim() (*types.Transaction, error) {
	return _Fukuastake.Contract.PauseClaim(&_Fukuastake.TransactOpts)
}

// PauseWithdraw is a paid mutator transaction binding the contract method 0x6155e3de.
//
// Solidity: function pauseWithdraw() returns()
func (_Fukuastake *FukuastakeTransactor) PauseWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "pauseWithdraw")
}

// PauseWithdraw is a paid mutator transaction binding the contract method 0x6155e3de.
//
// Solidity: function pauseWithdraw() returns()
func (_Fukuastake *FukuastakeSession) PauseWithdraw() (*types.Transaction, error) {
	return _Fukuastake.Contract.PauseWithdraw(&_Fukuastake.TransactOpts)
}

// PauseWithdraw is a paid mutator transaction binding the contract method 0x6155e3de.
//
// Solidity: function pauseWithdraw() returns()
func (_Fukuastake *FukuastakeTransactorSession) PauseWithdraw() (*types.Transaction, error) {
	return _Fukuastake.Contract.PauseWithdraw(&_Fukuastake.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Fukuastake *FukuastakeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Fukuastake *FukuastakeSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Fukuastake.Contract.RenounceRole(&_Fukuastake.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Fukuastake *FukuastakeTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Fukuastake.Contract.RenounceRole(&_Fukuastake.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Fukuastake *FukuastakeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Fukuastake *FukuastakeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fukuastake.Contract.RevokeRole(&_Fukuastake.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Fukuastake *FukuastakeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fukuastake.Contract.RevokeRole(&_Fukuastake.TransactOpts, role, account)
}

// SetEndBlock is a paid mutator transaction binding the contract method 0xc713aa94.
//
// Solidity: function setEndBlock(uint256 _endBlock) returns()
func (_Fukuastake *FukuastakeTransactor) SetEndBlock(opts *bind.TransactOpts, _endBlock *big.Int) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "setEndBlock", _endBlock)
}

// SetEndBlock is a paid mutator transaction binding the contract method 0xc713aa94.
//
// Solidity: function setEndBlock(uint256 _endBlock) returns()
func (_Fukuastake *FukuastakeSession) SetEndBlock(_endBlock *big.Int) (*types.Transaction, error) {
	return _Fukuastake.Contract.SetEndBlock(&_Fukuastake.TransactOpts, _endBlock)
}

// SetEndBlock is a paid mutator transaction binding the contract method 0xc713aa94.
//
// Solidity: function setEndBlock(uint256 _endBlock) returns()
func (_Fukuastake *FukuastakeTransactorSession) SetEndBlock(_endBlock *big.Int) (*types.Transaction, error) {
	return _Fukuastake.Contract.SetEndBlock(&_Fukuastake.TransactOpts, _endBlock)
}

// SetFukuaPerBlock is a paid mutator transaction binding the contract method 0xe26c25ce.
//
// Solidity: function setFukuaPerBlock(uint256 _fukuaPerBlock) returns()
func (_Fukuastake *FukuastakeTransactor) SetFukuaPerBlock(opts *bind.TransactOpts, _fukuaPerBlock *big.Int) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "setFukuaPerBlock", _fukuaPerBlock)
}

// SetFukuaPerBlock is a paid mutator transaction binding the contract method 0xe26c25ce.
//
// Solidity: function setFukuaPerBlock(uint256 _fukuaPerBlock) returns()
func (_Fukuastake *FukuastakeSession) SetFukuaPerBlock(_fukuaPerBlock *big.Int) (*types.Transaction, error) {
	return _Fukuastake.Contract.SetFukuaPerBlock(&_Fukuastake.TransactOpts, _fukuaPerBlock)
}

// SetFukuaPerBlock is a paid mutator transaction binding the contract method 0xe26c25ce.
//
// Solidity: function setFukuaPerBlock(uint256 _fukuaPerBlock) returns()
func (_Fukuastake *FukuastakeTransactorSession) SetFukuaPerBlock(_fukuaPerBlock *big.Int) (*types.Transaction, error) {
	return _Fukuastake.Contract.SetFukuaPerBlock(&_Fukuastake.TransactOpts, _fukuaPerBlock)
}

// SetFukuaToken is a paid mutator transaction binding the contract method 0x3ee521e2.
//
// Solidity: function setFukuaToken(address _fukuaToken) returns()
func (_Fukuastake *FukuastakeTransactor) SetFukuaToken(opts *bind.TransactOpts, _fukuaToken common.Address) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "setFukuaToken", _fukuaToken)
}

// SetFukuaToken is a paid mutator transaction binding the contract method 0x3ee521e2.
//
// Solidity: function setFukuaToken(address _fukuaToken) returns()
func (_Fukuastake *FukuastakeSession) SetFukuaToken(_fukuaToken common.Address) (*types.Transaction, error) {
	return _Fukuastake.Contract.SetFukuaToken(&_Fukuastake.TransactOpts, _fukuaToken)
}

// SetFukuaToken is a paid mutator transaction binding the contract method 0x3ee521e2.
//
// Solidity: function setFukuaToken(address _fukuaToken) returns()
func (_Fukuastake *FukuastakeTransactorSession) SetFukuaToken(_fukuaToken common.Address) (*types.Transaction, error) {
	return _Fukuastake.Contract.SetFukuaToken(&_Fukuastake.TransactOpts, _fukuaToken)
}

// SetPoolWeight is a paid mutator transaction binding the contract method 0xfad07ece.
//
// Solidity: function setPoolWeight(uint256 _pid, uint256 _poolWeight, bool _withUpdate) returns()
func (_Fukuastake *FukuastakeTransactor) SetPoolWeight(opts *bind.TransactOpts, _pid *big.Int, _poolWeight *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "setPoolWeight", _pid, _poolWeight, _withUpdate)
}

// SetPoolWeight is a paid mutator transaction binding the contract method 0xfad07ece.
//
// Solidity: function setPoolWeight(uint256 _pid, uint256 _poolWeight, bool _withUpdate) returns()
func (_Fukuastake *FukuastakeSession) SetPoolWeight(_pid *big.Int, _poolWeight *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Fukuastake.Contract.SetPoolWeight(&_Fukuastake.TransactOpts, _pid, _poolWeight, _withUpdate)
}

// SetPoolWeight is a paid mutator transaction binding the contract method 0xfad07ece.
//
// Solidity: function setPoolWeight(uint256 _pid, uint256 _poolWeight, bool _withUpdate) returns()
func (_Fukuastake *FukuastakeTransactorSession) SetPoolWeight(_pid *big.Int, _poolWeight *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Fukuastake.Contract.SetPoolWeight(&_Fukuastake.TransactOpts, _pid, _poolWeight, _withUpdate)
}

// SetStartBlock is a paid mutator transaction binding the contract method 0xf35e4a6e.
//
// Solidity: function setStartBlock(uint256 _startBlock) returns()
func (_Fukuastake *FukuastakeTransactor) SetStartBlock(opts *bind.TransactOpts, _startBlock *big.Int) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "setStartBlock", _startBlock)
}

// SetStartBlock is a paid mutator transaction binding the contract method 0xf35e4a6e.
//
// Solidity: function setStartBlock(uint256 _startBlock) returns()
func (_Fukuastake *FukuastakeSession) SetStartBlock(_startBlock *big.Int) (*types.Transaction, error) {
	return _Fukuastake.Contract.SetStartBlock(&_Fukuastake.TransactOpts, _startBlock)
}

// SetStartBlock is a paid mutator transaction binding the contract method 0xf35e4a6e.
//
// Solidity: function setStartBlock(uint256 _startBlock) returns()
func (_Fukuastake *FukuastakeTransactorSession) SetStartBlock(_startBlock *big.Int) (*types.Transaction, error) {
	return _Fukuastake.Contract.SetStartBlock(&_Fukuastake.TransactOpts, _startBlock)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Fukuastake *FukuastakeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Fukuastake *FukuastakeSession) Unpause() (*types.Transaction, error) {
	return _Fukuastake.Contract.Unpause(&_Fukuastake.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Fukuastake *FukuastakeTransactorSession) Unpause() (*types.Transaction, error) {
	return _Fukuastake.Contract.Unpause(&_Fukuastake.TransactOpts)
}

// UnpauseClaim is a paid mutator transaction binding the contract method 0xde065caa.
//
// Solidity: function unpauseClaim() returns()
func (_Fukuastake *FukuastakeTransactor) UnpauseClaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "unpauseClaim")
}

// UnpauseClaim is a paid mutator transaction binding the contract method 0xde065caa.
//
// Solidity: function unpauseClaim() returns()
func (_Fukuastake *FukuastakeSession) UnpauseClaim() (*types.Transaction, error) {
	return _Fukuastake.Contract.UnpauseClaim(&_Fukuastake.TransactOpts)
}

// UnpauseClaim is a paid mutator transaction binding the contract method 0xde065caa.
//
// Solidity: function unpauseClaim() returns()
func (_Fukuastake *FukuastakeTransactorSession) UnpauseClaim() (*types.Transaction, error) {
	return _Fukuastake.Contract.UnpauseClaim(&_Fukuastake.TransactOpts)
}

// UnpauseWithdraw is a paid mutator transaction binding the contract method 0x5bb6d007.
//
// Solidity: function unpauseWithdraw() returns()
func (_Fukuastake *FukuastakeTransactor) UnpauseWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "unpauseWithdraw")
}

// UnpauseWithdraw is a paid mutator transaction binding the contract method 0x5bb6d007.
//
// Solidity: function unpauseWithdraw() returns()
func (_Fukuastake *FukuastakeSession) UnpauseWithdraw() (*types.Transaction, error) {
	return _Fukuastake.Contract.UnpauseWithdraw(&_Fukuastake.TransactOpts)
}

// UnpauseWithdraw is a paid mutator transaction binding the contract method 0x5bb6d007.
//
// Solidity: function unpauseWithdraw() returns()
func (_Fukuastake *FukuastakeTransactorSession) UnpauseWithdraw() (*types.Transaction, error) {
	return _Fukuastake.Contract.UnpauseWithdraw(&_Fukuastake.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x9e2c8a5b.
//
// Solidity: function unstake(uint256 _pid, uint256 _amount) returns()
func (_Fukuastake *FukuastakeTransactor) Unstake(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "unstake", _pid, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x9e2c8a5b.
//
// Solidity: function unstake(uint256 _pid, uint256 _amount) returns()
func (_Fukuastake *FukuastakeSession) Unstake(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Fukuastake.Contract.Unstake(&_Fukuastake.TransactOpts, _pid, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x9e2c8a5b.
//
// Solidity: function unstake(uint256 _pid, uint256 _amount) returns()
func (_Fukuastake *FukuastakeTransactorSession) Unstake(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Fukuastake.Contract.Unstake(&_Fukuastake.TransactOpts, _pid, _amount)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Fukuastake *FukuastakeTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Fukuastake *FukuastakeSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Fukuastake.Contract.UpdatePool(&_Fukuastake.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Fukuastake *FukuastakeTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Fukuastake.Contract.UpdatePool(&_Fukuastake.TransactOpts, _pid)
}

// UpdatePool0 is a paid mutator transaction binding the contract method 0xd86c0444.
//
// Solidity: function updatePool(uint256 _pid, uint256 _minDepositAmount, uint256 _unstakeLockedBlocks) returns()
func (_Fukuastake *FukuastakeTransactor) UpdatePool0(opts *bind.TransactOpts, _pid *big.Int, _minDepositAmount *big.Int, _unstakeLockedBlocks *big.Int) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "updatePool0", _pid, _minDepositAmount, _unstakeLockedBlocks)
}

// UpdatePool0 is a paid mutator transaction binding the contract method 0xd86c0444.
//
// Solidity: function updatePool(uint256 _pid, uint256 _minDepositAmount, uint256 _unstakeLockedBlocks) returns()
func (_Fukuastake *FukuastakeSession) UpdatePool0(_pid *big.Int, _minDepositAmount *big.Int, _unstakeLockedBlocks *big.Int) (*types.Transaction, error) {
	return _Fukuastake.Contract.UpdatePool0(&_Fukuastake.TransactOpts, _pid, _minDepositAmount, _unstakeLockedBlocks)
}

// UpdatePool0 is a paid mutator transaction binding the contract method 0xd86c0444.
//
// Solidity: function updatePool(uint256 _pid, uint256 _minDepositAmount, uint256 _unstakeLockedBlocks) returns()
func (_Fukuastake *FukuastakeTransactorSession) UpdatePool0(_pid *big.Int, _minDepositAmount *big.Int, _unstakeLockedBlocks *big.Int) (*types.Transaction, error) {
	return _Fukuastake.Contract.UpdatePool0(&_Fukuastake.TransactOpts, _pid, _minDepositAmount, _unstakeLockedBlocks)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Fukuastake *FukuastakeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Fukuastake *FukuastakeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Fukuastake.Contract.UpgradeToAndCall(&_Fukuastake.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Fukuastake *FukuastakeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Fukuastake.Contract.UpgradeToAndCall(&_Fukuastake.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _pid) returns()
func (_Fukuastake *FukuastakeTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Fukuastake.contract.Transact(opts, "withdraw", _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _pid) returns()
func (_Fukuastake *FukuastakeSession) Withdraw(_pid *big.Int) (*types.Transaction, error) {
	return _Fukuastake.Contract.Withdraw(&_Fukuastake.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _pid) returns()
func (_Fukuastake *FukuastakeTransactorSession) Withdraw(_pid *big.Int) (*types.Transaction, error) {
	return _Fukuastake.Contract.Withdraw(&_Fukuastake.TransactOpts, _pid)
}

// FukuastakeAddPoolIterator is returned from FilterAddPool and is used to iterate over the raw logs and unpacked data for AddPool events raised by the Fukuastake contract.
type FukuastakeAddPoolIterator struct {
	Event *FukuastakeAddPool // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakeAddPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakeAddPool)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakeAddPool)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakeAddPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakeAddPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakeAddPool represents a AddPool event raised by the Fukuastake contract.
type FukuastakeAddPool struct {
	StTokenAddress      common.Address
	PoolWeight          *big.Int
	LastRewardBlock     *big.Int
	MinDepositAmount    *big.Int
	UnstakeLockedBlocks *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterAddPool is a free log retrieval operation binding the contract event 0x0fa296fce13e7a0e622b3a892e66220c248337289483a3cfa4130cde0caa1346.
//
// Solidity: event AddPool(address indexed stTokenAddress, uint256 indexed poolWeight, uint256 indexed lastRewardBlock, uint256 minDepositAmount, uint256 unstakeLockedBlocks)
func (_Fukuastake *FukuastakeFilterer) FilterAddPool(opts *bind.FilterOpts, stTokenAddress []common.Address, poolWeight []*big.Int, lastRewardBlock []*big.Int) (*FukuastakeAddPoolIterator, error) {

	var stTokenAddressRule []interface{}
	for _, stTokenAddressItem := range stTokenAddress {
		stTokenAddressRule = append(stTokenAddressRule, stTokenAddressItem)
	}
	var poolWeightRule []interface{}
	for _, poolWeightItem := range poolWeight {
		poolWeightRule = append(poolWeightRule, poolWeightItem)
	}
	var lastRewardBlockRule []interface{}
	for _, lastRewardBlockItem := range lastRewardBlock {
		lastRewardBlockRule = append(lastRewardBlockRule, lastRewardBlockItem)
	}

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "AddPool", stTokenAddressRule, poolWeightRule, lastRewardBlockRule)
	if err != nil {
		return nil, err
	}
	return &FukuastakeAddPoolIterator{contract: _Fukuastake.contract, event: "AddPool", logs: logs, sub: sub}, nil
}

// WatchAddPool is a free log subscription operation binding the contract event 0x0fa296fce13e7a0e622b3a892e66220c248337289483a3cfa4130cde0caa1346.
//
// Solidity: event AddPool(address indexed stTokenAddress, uint256 indexed poolWeight, uint256 indexed lastRewardBlock, uint256 minDepositAmount, uint256 unstakeLockedBlocks)
func (_Fukuastake *FukuastakeFilterer) WatchAddPool(opts *bind.WatchOpts, sink chan<- *FukuastakeAddPool, stTokenAddress []common.Address, poolWeight []*big.Int, lastRewardBlock []*big.Int) (event.Subscription, error) {

	var stTokenAddressRule []interface{}
	for _, stTokenAddressItem := range stTokenAddress {
		stTokenAddressRule = append(stTokenAddressRule, stTokenAddressItem)
	}
	var poolWeightRule []interface{}
	for _, poolWeightItem := range poolWeight {
		poolWeightRule = append(poolWeightRule, poolWeightItem)
	}
	var lastRewardBlockRule []interface{}
	for _, lastRewardBlockItem := range lastRewardBlock {
		lastRewardBlockRule = append(lastRewardBlockRule, lastRewardBlockItem)
	}

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "AddPool", stTokenAddressRule, poolWeightRule, lastRewardBlockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakeAddPool)
				if err := _Fukuastake.contract.UnpackLog(event, "AddPool", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddPool is a log parse operation binding the contract event 0x0fa296fce13e7a0e622b3a892e66220c248337289483a3cfa4130cde0caa1346.
//
// Solidity: event AddPool(address indexed stTokenAddress, uint256 indexed poolWeight, uint256 indexed lastRewardBlock, uint256 minDepositAmount, uint256 unstakeLockedBlocks)
func (_Fukuastake *FukuastakeFilterer) ParseAddPool(log types.Log) (*FukuastakeAddPool, error) {
	event := new(FukuastakeAddPool)
	if err := _Fukuastake.contract.UnpackLog(event, "AddPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakeClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Fukuastake contract.
type FukuastakeClaimIterator struct {
	Event *FukuastakeClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakeClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakeClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakeClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakeClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakeClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakeClaim represents a Claim event raised by the Fukuastake contract.
type FukuastakeClaim struct {
	User        common.Address
	PoolId      *big.Int
	FukuaReward *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7.
//
// Solidity: event Claim(address indexed user, uint256 indexed poolId, uint256 fukuaReward)
func (_Fukuastake *FukuastakeFilterer) FilterClaim(opts *bind.FilterOpts, user []common.Address, poolId []*big.Int) (*FukuastakeClaimIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "Claim", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return &FukuastakeClaimIterator{contract: _Fukuastake.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7.
//
// Solidity: event Claim(address indexed user, uint256 indexed poolId, uint256 fukuaReward)
func (_Fukuastake *FukuastakeFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *FukuastakeClaim, user []common.Address, poolId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "Claim", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakeClaim)
				if err := _Fukuastake.contract.UnpackLog(event, "Claim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaim is a log parse operation binding the contract event 0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7.
//
// Solidity: event Claim(address indexed user, uint256 indexed poolId, uint256 fukuaReward)
func (_Fukuastake *FukuastakeFilterer) ParseClaim(log types.Log) (*FukuastakeClaim, error) {
	event := new(FukuastakeClaim)
	if err := _Fukuastake.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Fukuastake contract.
type FukuastakeDepositIterator struct {
	Event *FukuastakeDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakeDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakeDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakeDeposit represents a Deposit event raised by the Fukuastake contract.
type FukuastakeDeposit struct {
	User   common.Address
	PoolId *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed poolId, uint256 amount)
func (_Fukuastake *FukuastakeFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, poolId []*big.Int) (*FukuastakeDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "Deposit", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return &FukuastakeDepositIterator{contract: _Fukuastake.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed poolId, uint256 amount)
func (_Fukuastake *FukuastakeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *FukuastakeDeposit, user []common.Address, poolId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "Deposit", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakeDeposit)
				if err := _Fukuastake.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed poolId, uint256 amount)
func (_Fukuastake *FukuastakeFilterer) ParseDeposit(log types.Log) (*FukuastakeDeposit, error) {
	event := new(FukuastakeDeposit)
	if err := _Fukuastake.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Fukuastake contract.
type FukuastakeInitializedIterator struct {
	Event *FukuastakeInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakeInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakeInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakeInitialized represents a Initialized event raised by the Fukuastake contract.
type FukuastakeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Fukuastake *FukuastakeFilterer) FilterInitialized(opts *bind.FilterOpts) (*FukuastakeInitializedIterator, error) {

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FukuastakeInitializedIterator{contract: _Fukuastake.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Fukuastake *FukuastakeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FukuastakeInitialized) (event.Subscription, error) {

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakeInitialized)
				if err := _Fukuastake.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Fukuastake *FukuastakeFilterer) ParseInitialized(log types.Log) (*FukuastakeInitialized, error) {
	event := new(FukuastakeInitialized)
	if err := _Fukuastake.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakePauseClaimIterator is returned from FilterPauseClaim and is used to iterate over the raw logs and unpacked data for PauseClaim events raised by the Fukuastake contract.
type FukuastakePauseClaimIterator struct {
	Event *FukuastakePauseClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakePauseClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakePauseClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakePauseClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakePauseClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakePauseClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakePauseClaim represents a PauseClaim event raised by the Fukuastake contract.
type FukuastakePauseClaim struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPauseClaim is a free log retrieval operation binding the contract event 0x6d73d6b34c378ab3bf6630206d60b7882801b91d03ee20d016ff0d5054db81e1.
//
// Solidity: event PauseClaim()
func (_Fukuastake *FukuastakeFilterer) FilterPauseClaim(opts *bind.FilterOpts) (*FukuastakePauseClaimIterator, error) {

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "PauseClaim")
	if err != nil {
		return nil, err
	}
	return &FukuastakePauseClaimIterator{contract: _Fukuastake.contract, event: "PauseClaim", logs: logs, sub: sub}, nil
}

// WatchPauseClaim is a free log subscription operation binding the contract event 0x6d73d6b34c378ab3bf6630206d60b7882801b91d03ee20d016ff0d5054db81e1.
//
// Solidity: event PauseClaim()
func (_Fukuastake *FukuastakeFilterer) WatchPauseClaim(opts *bind.WatchOpts, sink chan<- *FukuastakePauseClaim) (event.Subscription, error) {

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "PauseClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakePauseClaim)
				if err := _Fukuastake.contract.UnpackLog(event, "PauseClaim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauseClaim is a log parse operation binding the contract event 0x6d73d6b34c378ab3bf6630206d60b7882801b91d03ee20d016ff0d5054db81e1.
//
// Solidity: event PauseClaim()
func (_Fukuastake *FukuastakeFilterer) ParsePauseClaim(log types.Log) (*FukuastakePauseClaim, error) {
	event := new(FukuastakePauseClaim)
	if err := _Fukuastake.contract.UnpackLog(event, "PauseClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakePauseWithdrawIterator is returned from FilterPauseWithdraw and is used to iterate over the raw logs and unpacked data for PauseWithdraw events raised by the Fukuastake contract.
type FukuastakePauseWithdrawIterator struct {
	Event *FukuastakePauseWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakePauseWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakePauseWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakePauseWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakePauseWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakePauseWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakePauseWithdraw represents a PauseWithdraw event raised by the Fukuastake contract.
type FukuastakePauseWithdraw struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPauseWithdraw is a free log retrieval operation binding the contract event 0x8099f593a6aaecd68b6494933cd71f703376ac3975be83692e1b7d800abf6837.
//
// Solidity: event PauseWithdraw()
func (_Fukuastake *FukuastakeFilterer) FilterPauseWithdraw(opts *bind.FilterOpts) (*FukuastakePauseWithdrawIterator, error) {

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "PauseWithdraw")
	if err != nil {
		return nil, err
	}
	return &FukuastakePauseWithdrawIterator{contract: _Fukuastake.contract, event: "PauseWithdraw", logs: logs, sub: sub}, nil
}

// WatchPauseWithdraw is a free log subscription operation binding the contract event 0x8099f593a6aaecd68b6494933cd71f703376ac3975be83692e1b7d800abf6837.
//
// Solidity: event PauseWithdraw()
func (_Fukuastake *FukuastakeFilterer) WatchPauseWithdraw(opts *bind.WatchOpts, sink chan<- *FukuastakePauseWithdraw) (event.Subscription, error) {

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "PauseWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakePauseWithdraw)
				if err := _Fukuastake.contract.UnpackLog(event, "PauseWithdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauseWithdraw is a log parse operation binding the contract event 0x8099f593a6aaecd68b6494933cd71f703376ac3975be83692e1b7d800abf6837.
//
// Solidity: event PauseWithdraw()
func (_Fukuastake *FukuastakeFilterer) ParsePauseWithdraw(log types.Log) (*FukuastakePauseWithdraw, error) {
	event := new(FukuastakePauseWithdraw)
	if err := _Fukuastake.contract.UnpackLog(event, "PauseWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Fukuastake contract.
type FukuastakePausedIterator struct {
	Event *FukuastakePaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakePaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakePaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakePaused represents a Paused event raised by the Fukuastake contract.
type FukuastakePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Fukuastake *FukuastakeFilterer) FilterPaused(opts *bind.FilterOpts) (*FukuastakePausedIterator, error) {

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &FukuastakePausedIterator{contract: _Fukuastake.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Fukuastake *FukuastakeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *FukuastakePaused) (event.Subscription, error) {

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakePaused)
				if err := _Fukuastake.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Fukuastake *FukuastakeFilterer) ParsePaused(log types.Log) (*FukuastakePaused, error) {
	event := new(FukuastakePaused)
	if err := _Fukuastake.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakeRequestUnstakeIterator is returned from FilterRequestUnstake and is used to iterate over the raw logs and unpacked data for RequestUnstake events raised by the Fukuastake contract.
type FukuastakeRequestUnstakeIterator struct {
	Event *FukuastakeRequestUnstake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakeRequestUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakeRequestUnstake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakeRequestUnstake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakeRequestUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakeRequestUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakeRequestUnstake represents a RequestUnstake event raised by the Fukuastake contract.
type FukuastakeRequestUnstake struct {
	User   common.Address
	PoolId *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRequestUnstake is a free log retrieval operation binding the contract event 0xc80277265097707f6f12a4ac4c09d46c9926e2eea2536f63616cb04d9fcad7d6.
//
// Solidity: event RequestUnstake(address indexed user, uint256 indexed poolId, uint256 amount)
func (_Fukuastake *FukuastakeFilterer) FilterRequestUnstake(opts *bind.FilterOpts, user []common.Address, poolId []*big.Int) (*FukuastakeRequestUnstakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "RequestUnstake", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return &FukuastakeRequestUnstakeIterator{contract: _Fukuastake.contract, event: "RequestUnstake", logs: logs, sub: sub}, nil
}

// WatchRequestUnstake is a free log subscription operation binding the contract event 0xc80277265097707f6f12a4ac4c09d46c9926e2eea2536f63616cb04d9fcad7d6.
//
// Solidity: event RequestUnstake(address indexed user, uint256 indexed poolId, uint256 amount)
func (_Fukuastake *FukuastakeFilterer) WatchRequestUnstake(opts *bind.WatchOpts, sink chan<- *FukuastakeRequestUnstake, user []common.Address, poolId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "RequestUnstake", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakeRequestUnstake)
				if err := _Fukuastake.contract.UnpackLog(event, "RequestUnstake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestUnstake is a log parse operation binding the contract event 0xc80277265097707f6f12a4ac4c09d46c9926e2eea2536f63616cb04d9fcad7d6.
//
// Solidity: event RequestUnstake(address indexed user, uint256 indexed poolId, uint256 amount)
func (_Fukuastake *FukuastakeFilterer) ParseRequestUnstake(log types.Log) (*FukuastakeRequestUnstake, error) {
	event := new(FukuastakeRequestUnstake)
	if err := _Fukuastake.contract.UnpackLog(event, "RequestUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Fukuastake contract.
type FukuastakeRoleAdminChangedIterator struct {
	Event *FukuastakeRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakeRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakeRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakeRoleAdminChanged represents a RoleAdminChanged event raised by the Fukuastake contract.
type FukuastakeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Fukuastake *FukuastakeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*FukuastakeRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &FukuastakeRoleAdminChangedIterator{contract: _Fukuastake.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Fukuastake *FukuastakeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *FukuastakeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakeRoleAdminChanged)
				if err := _Fukuastake.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Fukuastake *FukuastakeFilterer) ParseRoleAdminChanged(log types.Log) (*FukuastakeRoleAdminChanged, error) {
	event := new(FukuastakeRoleAdminChanged)
	if err := _Fukuastake.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Fukuastake contract.
type FukuastakeRoleGrantedIterator struct {
	Event *FukuastakeRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakeRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakeRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakeRoleGranted represents a RoleGranted event raised by the Fukuastake contract.
type FukuastakeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Fukuastake *FukuastakeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FukuastakeRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FukuastakeRoleGrantedIterator{contract: _Fukuastake.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Fukuastake *FukuastakeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *FukuastakeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakeRoleGranted)
				if err := _Fukuastake.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Fukuastake *FukuastakeFilterer) ParseRoleGranted(log types.Log) (*FukuastakeRoleGranted, error) {
	event := new(FukuastakeRoleGranted)
	if err := _Fukuastake.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Fukuastake contract.
type FukuastakeRoleRevokedIterator struct {
	Event *FukuastakeRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakeRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakeRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakeRoleRevoked represents a RoleRevoked event raised by the Fukuastake contract.
type FukuastakeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Fukuastake *FukuastakeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FukuastakeRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FukuastakeRoleRevokedIterator{contract: _Fukuastake.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Fukuastake *FukuastakeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *FukuastakeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakeRoleRevoked)
				if err := _Fukuastake.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Fukuastake *FukuastakeFilterer) ParseRoleRevoked(log types.Log) (*FukuastakeRoleRevoked, error) {
	event := new(FukuastakeRoleRevoked)
	if err := _Fukuastake.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakeSetEndBlockIterator is returned from FilterSetEndBlock and is used to iterate over the raw logs and unpacked data for SetEndBlock events raised by the Fukuastake contract.
type FukuastakeSetEndBlockIterator struct {
	Event *FukuastakeSetEndBlock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakeSetEndBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakeSetEndBlock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakeSetEndBlock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakeSetEndBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakeSetEndBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakeSetEndBlock represents a SetEndBlock event raised by the Fukuastake contract.
type FukuastakeSetEndBlock struct {
	EndBlock *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetEndBlock is a free log retrieval operation binding the contract event 0x1132c5baccb51da3d049fabc819697dc845fa224ad59d9b555507d6446b40850.
//
// Solidity: event SetEndBlock(uint256 indexed endBlock)
func (_Fukuastake *FukuastakeFilterer) FilterSetEndBlock(opts *bind.FilterOpts, endBlock []*big.Int) (*FukuastakeSetEndBlockIterator, error) {

	var endBlockRule []interface{}
	for _, endBlockItem := range endBlock {
		endBlockRule = append(endBlockRule, endBlockItem)
	}

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "SetEndBlock", endBlockRule)
	if err != nil {
		return nil, err
	}
	return &FukuastakeSetEndBlockIterator{contract: _Fukuastake.contract, event: "SetEndBlock", logs: logs, sub: sub}, nil
}

// WatchSetEndBlock is a free log subscription operation binding the contract event 0x1132c5baccb51da3d049fabc819697dc845fa224ad59d9b555507d6446b40850.
//
// Solidity: event SetEndBlock(uint256 indexed endBlock)
func (_Fukuastake *FukuastakeFilterer) WatchSetEndBlock(opts *bind.WatchOpts, sink chan<- *FukuastakeSetEndBlock, endBlock []*big.Int) (event.Subscription, error) {

	var endBlockRule []interface{}
	for _, endBlockItem := range endBlock {
		endBlockRule = append(endBlockRule, endBlockItem)
	}

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "SetEndBlock", endBlockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakeSetEndBlock)
				if err := _Fukuastake.contract.UnpackLog(event, "SetEndBlock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetEndBlock is a log parse operation binding the contract event 0x1132c5baccb51da3d049fabc819697dc845fa224ad59d9b555507d6446b40850.
//
// Solidity: event SetEndBlock(uint256 indexed endBlock)
func (_Fukuastake *FukuastakeFilterer) ParseSetEndBlock(log types.Log) (*FukuastakeSetEndBlock, error) {
	event := new(FukuastakeSetEndBlock)
	if err := _Fukuastake.contract.UnpackLog(event, "SetEndBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakeSetFukuaPerBlockIterator is returned from FilterSetFukuaPerBlock and is used to iterate over the raw logs and unpacked data for SetFukuaPerBlock events raised by the Fukuastake contract.
type FukuastakeSetFukuaPerBlockIterator struct {
	Event *FukuastakeSetFukuaPerBlock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakeSetFukuaPerBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakeSetFukuaPerBlock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakeSetFukuaPerBlock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakeSetFukuaPerBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakeSetFukuaPerBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakeSetFukuaPerBlock represents a SetFukuaPerBlock event raised by the Fukuastake contract.
type FukuastakeSetFukuaPerBlock struct {
	NewFukuaPerBlock *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetFukuaPerBlock is a free log retrieval operation binding the contract event 0xda51cd4dfe7787695fefa2693693bfac4e20eb62914567a95f6672a110406440.
//
// Solidity: event SetFukuaPerBlock(uint256 indexed newFukuaPerBlock)
func (_Fukuastake *FukuastakeFilterer) FilterSetFukuaPerBlock(opts *bind.FilterOpts, newFukuaPerBlock []*big.Int) (*FukuastakeSetFukuaPerBlockIterator, error) {

	var newFukuaPerBlockRule []interface{}
	for _, newFukuaPerBlockItem := range newFukuaPerBlock {
		newFukuaPerBlockRule = append(newFukuaPerBlockRule, newFukuaPerBlockItem)
	}

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "SetFukuaPerBlock", newFukuaPerBlockRule)
	if err != nil {
		return nil, err
	}
	return &FukuastakeSetFukuaPerBlockIterator{contract: _Fukuastake.contract, event: "SetFukuaPerBlock", logs: logs, sub: sub}, nil
}

// WatchSetFukuaPerBlock is a free log subscription operation binding the contract event 0xda51cd4dfe7787695fefa2693693bfac4e20eb62914567a95f6672a110406440.
//
// Solidity: event SetFukuaPerBlock(uint256 indexed newFukuaPerBlock)
func (_Fukuastake *FukuastakeFilterer) WatchSetFukuaPerBlock(opts *bind.WatchOpts, sink chan<- *FukuastakeSetFukuaPerBlock, newFukuaPerBlock []*big.Int) (event.Subscription, error) {

	var newFukuaPerBlockRule []interface{}
	for _, newFukuaPerBlockItem := range newFukuaPerBlock {
		newFukuaPerBlockRule = append(newFukuaPerBlockRule, newFukuaPerBlockItem)
	}

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "SetFukuaPerBlock", newFukuaPerBlockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakeSetFukuaPerBlock)
				if err := _Fukuastake.contract.UnpackLog(event, "SetFukuaPerBlock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetFukuaPerBlock is a log parse operation binding the contract event 0xda51cd4dfe7787695fefa2693693bfac4e20eb62914567a95f6672a110406440.
//
// Solidity: event SetFukuaPerBlock(uint256 indexed newFukuaPerBlock)
func (_Fukuastake *FukuastakeFilterer) ParseSetFukuaPerBlock(log types.Log) (*FukuastakeSetFukuaPerBlock, error) {
	event := new(FukuastakeSetFukuaPerBlock)
	if err := _Fukuastake.contract.UnpackLog(event, "SetFukuaPerBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakeSetFukuaTokenIterator is returned from FilterSetFukuaToken and is used to iterate over the raw logs and unpacked data for SetFukuaToken events raised by the Fukuastake contract.
type FukuastakeSetFukuaTokenIterator struct {
	Event *FukuastakeSetFukuaToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakeSetFukuaTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakeSetFukuaToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakeSetFukuaToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakeSetFukuaTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakeSetFukuaTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakeSetFukuaToken represents a SetFukuaToken event raised by the Fukuastake contract.
type FukuastakeSetFukuaToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetFukuaToken is a free log retrieval operation binding the contract event 0x861966917ccb4b480494600b4105dff4d2cb14675074324df6acb1281b7a8ac0.
//
// Solidity: event SetFukuaToken(address indexed token)
func (_Fukuastake *FukuastakeFilterer) FilterSetFukuaToken(opts *bind.FilterOpts, token []common.Address) (*FukuastakeSetFukuaTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "SetFukuaToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &FukuastakeSetFukuaTokenIterator{contract: _Fukuastake.contract, event: "SetFukuaToken", logs: logs, sub: sub}, nil
}

// WatchSetFukuaToken is a free log subscription operation binding the contract event 0x861966917ccb4b480494600b4105dff4d2cb14675074324df6acb1281b7a8ac0.
//
// Solidity: event SetFukuaToken(address indexed token)
func (_Fukuastake *FukuastakeFilterer) WatchSetFukuaToken(opts *bind.WatchOpts, sink chan<- *FukuastakeSetFukuaToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "SetFukuaToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakeSetFukuaToken)
				if err := _Fukuastake.contract.UnpackLog(event, "SetFukuaToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetFukuaToken is a log parse operation binding the contract event 0x861966917ccb4b480494600b4105dff4d2cb14675074324df6acb1281b7a8ac0.
//
// Solidity: event SetFukuaToken(address indexed token)
func (_Fukuastake *FukuastakeFilterer) ParseSetFukuaToken(log types.Log) (*FukuastakeSetFukuaToken, error) {
	event := new(FukuastakeSetFukuaToken)
	if err := _Fukuastake.contract.UnpackLog(event, "SetFukuaToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakeSetPoolWeightIterator is returned from FilterSetPoolWeight and is used to iterate over the raw logs and unpacked data for SetPoolWeight events raised by the Fukuastake contract.
type FukuastakeSetPoolWeightIterator struct {
	Event *FukuastakeSetPoolWeight // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakeSetPoolWeightIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakeSetPoolWeight)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakeSetPoolWeight)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakeSetPoolWeightIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakeSetPoolWeightIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakeSetPoolWeight represents a SetPoolWeight event raised by the Fukuastake contract.
type FukuastakeSetPoolWeight struct {
	PoolId          *big.Int
	PoolWeight      *big.Int
	TotalPoolWeight *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetPoolWeight is a free log retrieval operation binding the contract event 0x4b8fa3d6a87cb21d1bf4978bf60628ae358a28ac7f39de1751a481c6dd957617.
//
// Solidity: event SetPoolWeight(uint256 indexed poolId, uint256 indexed poolWeight, uint256 totalPoolWeight)
func (_Fukuastake *FukuastakeFilterer) FilterSetPoolWeight(opts *bind.FilterOpts, poolId []*big.Int, poolWeight []*big.Int) (*FukuastakeSetPoolWeightIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var poolWeightRule []interface{}
	for _, poolWeightItem := range poolWeight {
		poolWeightRule = append(poolWeightRule, poolWeightItem)
	}

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "SetPoolWeight", poolIdRule, poolWeightRule)
	if err != nil {
		return nil, err
	}
	return &FukuastakeSetPoolWeightIterator{contract: _Fukuastake.contract, event: "SetPoolWeight", logs: logs, sub: sub}, nil
}

// WatchSetPoolWeight is a free log subscription operation binding the contract event 0x4b8fa3d6a87cb21d1bf4978bf60628ae358a28ac7f39de1751a481c6dd957617.
//
// Solidity: event SetPoolWeight(uint256 indexed poolId, uint256 indexed poolWeight, uint256 totalPoolWeight)
func (_Fukuastake *FukuastakeFilterer) WatchSetPoolWeight(opts *bind.WatchOpts, sink chan<- *FukuastakeSetPoolWeight, poolId []*big.Int, poolWeight []*big.Int) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var poolWeightRule []interface{}
	for _, poolWeightItem := range poolWeight {
		poolWeightRule = append(poolWeightRule, poolWeightItem)
	}

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "SetPoolWeight", poolIdRule, poolWeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakeSetPoolWeight)
				if err := _Fukuastake.contract.UnpackLog(event, "SetPoolWeight", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetPoolWeight is a log parse operation binding the contract event 0x4b8fa3d6a87cb21d1bf4978bf60628ae358a28ac7f39de1751a481c6dd957617.
//
// Solidity: event SetPoolWeight(uint256 indexed poolId, uint256 indexed poolWeight, uint256 totalPoolWeight)
func (_Fukuastake *FukuastakeFilterer) ParseSetPoolWeight(log types.Log) (*FukuastakeSetPoolWeight, error) {
	event := new(FukuastakeSetPoolWeight)
	if err := _Fukuastake.contract.UnpackLog(event, "SetPoolWeight", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakeSetStartBlockIterator is returned from FilterSetStartBlock and is used to iterate over the raw logs and unpacked data for SetStartBlock events raised by the Fukuastake contract.
type FukuastakeSetStartBlockIterator struct {
	Event *FukuastakeSetStartBlock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakeSetStartBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakeSetStartBlock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakeSetStartBlock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakeSetStartBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakeSetStartBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakeSetStartBlock represents a SetStartBlock event raised by the Fukuastake contract.
type FukuastakeSetStartBlock struct {
	StartBlock *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetStartBlock is a free log retrieval operation binding the contract event 0x63b90b79f11a0f132bcb2c4a4ddd44abda45c1308a83b2919318df7f5f8b7be4.
//
// Solidity: event SetStartBlock(uint256 indexed startBlock)
func (_Fukuastake *FukuastakeFilterer) FilterSetStartBlock(opts *bind.FilterOpts, startBlock []*big.Int) (*FukuastakeSetStartBlockIterator, error) {

	var startBlockRule []interface{}
	for _, startBlockItem := range startBlock {
		startBlockRule = append(startBlockRule, startBlockItem)
	}

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "SetStartBlock", startBlockRule)
	if err != nil {
		return nil, err
	}
	return &FukuastakeSetStartBlockIterator{contract: _Fukuastake.contract, event: "SetStartBlock", logs: logs, sub: sub}, nil
}

// WatchSetStartBlock is a free log subscription operation binding the contract event 0x63b90b79f11a0f132bcb2c4a4ddd44abda45c1308a83b2919318df7f5f8b7be4.
//
// Solidity: event SetStartBlock(uint256 indexed startBlock)
func (_Fukuastake *FukuastakeFilterer) WatchSetStartBlock(opts *bind.WatchOpts, sink chan<- *FukuastakeSetStartBlock, startBlock []*big.Int) (event.Subscription, error) {

	var startBlockRule []interface{}
	for _, startBlockItem := range startBlock {
		startBlockRule = append(startBlockRule, startBlockItem)
	}

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "SetStartBlock", startBlockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakeSetStartBlock)
				if err := _Fukuastake.contract.UnpackLog(event, "SetStartBlock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetStartBlock is a log parse operation binding the contract event 0x63b90b79f11a0f132bcb2c4a4ddd44abda45c1308a83b2919318df7f5f8b7be4.
//
// Solidity: event SetStartBlock(uint256 indexed startBlock)
func (_Fukuastake *FukuastakeFilterer) ParseSetStartBlock(log types.Log) (*FukuastakeSetStartBlock, error) {
	event := new(FukuastakeSetStartBlock)
	if err := _Fukuastake.contract.UnpackLog(event, "SetStartBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakeUnpauseClaimIterator is returned from FilterUnpauseClaim and is used to iterate over the raw logs and unpacked data for UnpauseClaim events raised by the Fukuastake contract.
type FukuastakeUnpauseClaimIterator struct {
	Event *FukuastakeUnpauseClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakeUnpauseClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakeUnpauseClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakeUnpauseClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakeUnpauseClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakeUnpauseClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakeUnpauseClaim represents a UnpauseClaim event raised by the Fukuastake contract.
type FukuastakeUnpauseClaim struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpauseClaim is a free log retrieval operation binding the contract event 0xe72cb12952f056e3e7496019725f20a13108ca420f67f1ee9c9cdab73fb8ce85.
//
// Solidity: event UnpauseClaim()
func (_Fukuastake *FukuastakeFilterer) FilterUnpauseClaim(opts *bind.FilterOpts) (*FukuastakeUnpauseClaimIterator, error) {

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "UnpauseClaim")
	if err != nil {
		return nil, err
	}
	return &FukuastakeUnpauseClaimIterator{contract: _Fukuastake.contract, event: "UnpauseClaim", logs: logs, sub: sub}, nil
}

// WatchUnpauseClaim is a free log subscription operation binding the contract event 0xe72cb12952f056e3e7496019725f20a13108ca420f67f1ee9c9cdab73fb8ce85.
//
// Solidity: event UnpauseClaim()
func (_Fukuastake *FukuastakeFilterer) WatchUnpauseClaim(opts *bind.WatchOpts, sink chan<- *FukuastakeUnpauseClaim) (event.Subscription, error) {

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "UnpauseClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakeUnpauseClaim)
				if err := _Fukuastake.contract.UnpackLog(event, "UnpauseClaim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpauseClaim is a log parse operation binding the contract event 0xe72cb12952f056e3e7496019725f20a13108ca420f67f1ee9c9cdab73fb8ce85.
//
// Solidity: event UnpauseClaim()
func (_Fukuastake *FukuastakeFilterer) ParseUnpauseClaim(log types.Log) (*FukuastakeUnpauseClaim, error) {
	event := new(FukuastakeUnpauseClaim)
	if err := _Fukuastake.contract.UnpackLog(event, "UnpauseClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakeUnpauseWithdrawIterator is returned from FilterUnpauseWithdraw and is used to iterate over the raw logs and unpacked data for UnpauseWithdraw events raised by the Fukuastake contract.
type FukuastakeUnpauseWithdrawIterator struct {
	Event *FukuastakeUnpauseWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakeUnpauseWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakeUnpauseWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakeUnpauseWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakeUnpauseWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakeUnpauseWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakeUnpauseWithdraw represents a UnpauseWithdraw event raised by the Fukuastake contract.
type FukuastakeUnpauseWithdraw struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpauseWithdraw is a free log retrieval operation binding the contract event 0x1c84bcaead48b692cc46b9b12e9a068951a59c99a2e2bf10b00b60b403cf12e2.
//
// Solidity: event UnpauseWithdraw()
func (_Fukuastake *FukuastakeFilterer) FilterUnpauseWithdraw(opts *bind.FilterOpts) (*FukuastakeUnpauseWithdrawIterator, error) {

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "UnpauseWithdraw")
	if err != nil {
		return nil, err
	}
	return &FukuastakeUnpauseWithdrawIterator{contract: _Fukuastake.contract, event: "UnpauseWithdraw", logs: logs, sub: sub}, nil
}

// WatchUnpauseWithdraw is a free log subscription operation binding the contract event 0x1c84bcaead48b692cc46b9b12e9a068951a59c99a2e2bf10b00b60b403cf12e2.
//
// Solidity: event UnpauseWithdraw()
func (_Fukuastake *FukuastakeFilterer) WatchUnpauseWithdraw(opts *bind.WatchOpts, sink chan<- *FukuastakeUnpauseWithdraw) (event.Subscription, error) {

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "UnpauseWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakeUnpauseWithdraw)
				if err := _Fukuastake.contract.UnpackLog(event, "UnpauseWithdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpauseWithdraw is a log parse operation binding the contract event 0x1c84bcaead48b692cc46b9b12e9a068951a59c99a2e2bf10b00b60b403cf12e2.
//
// Solidity: event UnpauseWithdraw()
func (_Fukuastake *FukuastakeFilterer) ParseUnpauseWithdraw(log types.Log) (*FukuastakeUnpauseWithdraw, error) {
	event := new(FukuastakeUnpauseWithdraw)
	if err := _Fukuastake.contract.UnpackLog(event, "UnpauseWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Fukuastake contract.
type FukuastakeUnpausedIterator struct {
	Event *FukuastakeUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakeUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakeUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakeUnpaused represents a Unpaused event raised by the Fukuastake contract.
type FukuastakeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Fukuastake *FukuastakeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*FukuastakeUnpausedIterator, error) {

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &FukuastakeUnpausedIterator{contract: _Fukuastake.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Fukuastake *FukuastakeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *FukuastakeUnpaused) (event.Subscription, error) {

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakeUnpaused)
				if err := _Fukuastake.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Fukuastake *FukuastakeFilterer) ParseUnpaused(log types.Log) (*FukuastakeUnpaused, error) {
	event := new(FukuastakeUnpaused)
	if err := _Fukuastake.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakeUpdatePoolIterator is returned from FilterUpdatePool and is used to iterate over the raw logs and unpacked data for UpdatePool events raised by the Fukuastake contract.
type FukuastakeUpdatePoolIterator struct {
	Event *FukuastakeUpdatePool // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakeUpdatePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakeUpdatePool)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakeUpdatePool)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakeUpdatePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakeUpdatePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakeUpdatePool represents a UpdatePool event raised by the Fukuastake contract.
type FukuastakeUpdatePool struct {
	PoolId           *big.Int
	LastRewardBlock  *big.Int
	TotalFukuaReward *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdatePool is a free log retrieval operation binding the contract event 0xf5d2d72d9b25d6853afd7d0554a113b705234b6a68bb36b7f143662994632411.
//
// Solidity: event UpdatePool(uint256 indexed poolId, uint256 indexed lastRewardBlock, uint256 totalFukuaReward)
func (_Fukuastake *FukuastakeFilterer) FilterUpdatePool(opts *bind.FilterOpts, poolId []*big.Int, lastRewardBlock []*big.Int) (*FukuastakeUpdatePoolIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var lastRewardBlockRule []interface{}
	for _, lastRewardBlockItem := range lastRewardBlock {
		lastRewardBlockRule = append(lastRewardBlockRule, lastRewardBlockItem)
	}

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "UpdatePool", poolIdRule, lastRewardBlockRule)
	if err != nil {
		return nil, err
	}
	return &FukuastakeUpdatePoolIterator{contract: _Fukuastake.contract, event: "UpdatePool", logs: logs, sub: sub}, nil
}

// WatchUpdatePool is a free log subscription operation binding the contract event 0xf5d2d72d9b25d6853afd7d0554a113b705234b6a68bb36b7f143662994632411.
//
// Solidity: event UpdatePool(uint256 indexed poolId, uint256 indexed lastRewardBlock, uint256 totalFukuaReward)
func (_Fukuastake *FukuastakeFilterer) WatchUpdatePool(opts *bind.WatchOpts, sink chan<- *FukuastakeUpdatePool, poolId []*big.Int, lastRewardBlock []*big.Int) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var lastRewardBlockRule []interface{}
	for _, lastRewardBlockItem := range lastRewardBlock {
		lastRewardBlockRule = append(lastRewardBlockRule, lastRewardBlockItem)
	}

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "UpdatePool", poolIdRule, lastRewardBlockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakeUpdatePool)
				if err := _Fukuastake.contract.UnpackLog(event, "UpdatePool", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatePool is a log parse operation binding the contract event 0xf5d2d72d9b25d6853afd7d0554a113b705234b6a68bb36b7f143662994632411.
//
// Solidity: event UpdatePool(uint256 indexed poolId, uint256 indexed lastRewardBlock, uint256 totalFukuaReward)
func (_Fukuastake *FukuastakeFilterer) ParseUpdatePool(log types.Log) (*FukuastakeUpdatePool, error) {
	event := new(FukuastakeUpdatePool)
	if err := _Fukuastake.contract.UnpackLog(event, "UpdatePool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakeUpdatePoolInfoIterator is returned from FilterUpdatePoolInfo and is used to iterate over the raw logs and unpacked data for UpdatePoolInfo events raised by the Fukuastake contract.
type FukuastakeUpdatePoolInfoIterator struct {
	Event *FukuastakeUpdatePoolInfo // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakeUpdatePoolInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakeUpdatePoolInfo)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakeUpdatePoolInfo)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakeUpdatePoolInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakeUpdatePoolInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakeUpdatePoolInfo represents a UpdatePoolInfo event raised by the Fukuastake contract.
type FukuastakeUpdatePoolInfo struct {
	PoolId              *big.Int
	MinDepositAmount    *big.Int
	UnstakeLockedBlocks *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUpdatePoolInfo is a free log retrieval operation binding the contract event 0x30dffdedaa3e3b4849298233f7cd71d229956e875ab09270498c96b7cf9181fd.
//
// Solidity: event UpdatePoolInfo(uint256 indexed poolId, uint256 indexed minDepositAmount, uint256 indexed unstakeLockedBlocks)
func (_Fukuastake *FukuastakeFilterer) FilterUpdatePoolInfo(opts *bind.FilterOpts, poolId []*big.Int, minDepositAmount []*big.Int, unstakeLockedBlocks []*big.Int) (*FukuastakeUpdatePoolInfoIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var minDepositAmountRule []interface{}
	for _, minDepositAmountItem := range minDepositAmount {
		minDepositAmountRule = append(minDepositAmountRule, minDepositAmountItem)
	}
	var unstakeLockedBlocksRule []interface{}
	for _, unstakeLockedBlocksItem := range unstakeLockedBlocks {
		unstakeLockedBlocksRule = append(unstakeLockedBlocksRule, unstakeLockedBlocksItem)
	}

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "UpdatePoolInfo", poolIdRule, minDepositAmountRule, unstakeLockedBlocksRule)
	if err != nil {
		return nil, err
	}
	return &FukuastakeUpdatePoolInfoIterator{contract: _Fukuastake.contract, event: "UpdatePoolInfo", logs: logs, sub: sub}, nil
}

// WatchUpdatePoolInfo is a free log subscription operation binding the contract event 0x30dffdedaa3e3b4849298233f7cd71d229956e875ab09270498c96b7cf9181fd.
//
// Solidity: event UpdatePoolInfo(uint256 indexed poolId, uint256 indexed minDepositAmount, uint256 indexed unstakeLockedBlocks)
func (_Fukuastake *FukuastakeFilterer) WatchUpdatePoolInfo(opts *bind.WatchOpts, sink chan<- *FukuastakeUpdatePoolInfo, poolId []*big.Int, minDepositAmount []*big.Int, unstakeLockedBlocks []*big.Int) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var minDepositAmountRule []interface{}
	for _, minDepositAmountItem := range minDepositAmount {
		minDepositAmountRule = append(minDepositAmountRule, minDepositAmountItem)
	}
	var unstakeLockedBlocksRule []interface{}
	for _, unstakeLockedBlocksItem := range unstakeLockedBlocks {
		unstakeLockedBlocksRule = append(unstakeLockedBlocksRule, unstakeLockedBlocksItem)
	}

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "UpdatePoolInfo", poolIdRule, minDepositAmountRule, unstakeLockedBlocksRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakeUpdatePoolInfo)
				if err := _Fukuastake.contract.UnpackLog(event, "UpdatePoolInfo", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatePoolInfo is a log parse operation binding the contract event 0x30dffdedaa3e3b4849298233f7cd71d229956e875ab09270498c96b7cf9181fd.
//
// Solidity: event UpdatePoolInfo(uint256 indexed poolId, uint256 indexed minDepositAmount, uint256 indexed unstakeLockedBlocks)
func (_Fukuastake *FukuastakeFilterer) ParseUpdatePoolInfo(log types.Log) (*FukuastakeUpdatePoolInfo, error) {
	event := new(FukuastakeUpdatePoolInfo)
	if err := _Fukuastake.contract.UnpackLog(event, "UpdatePoolInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Fukuastake contract.
type FukuastakeUpgradedIterator struct {
	Event *FukuastakeUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakeUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakeUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakeUpgraded represents a Upgraded event raised by the Fukuastake contract.
type FukuastakeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Fukuastake *FukuastakeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*FukuastakeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &FukuastakeUpgradedIterator{contract: _Fukuastake.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Fukuastake *FukuastakeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *FukuastakeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakeUpgraded)
				if err := _Fukuastake.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Fukuastake *FukuastakeFilterer) ParseUpgraded(log types.Log) (*FukuastakeUpgraded, error) {
	event := new(FukuastakeUpgraded)
	if err := _Fukuastake.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FukuastakeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Fukuastake contract.
type FukuastakeWithdrawIterator struct {
	Event *FukuastakeWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FukuastakeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FukuastakeWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FukuastakeWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FukuastakeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FukuastakeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FukuastakeWithdraw represents a Withdraw event raised by the Fukuastake contract.
type FukuastakeWithdraw struct {
	User        common.Address
	PoolId      *big.Int
	Amount      *big.Int
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed poolId, uint256 amount, uint256 indexed blockNumber)
func (_Fukuastake *FukuastakeFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, poolId []*big.Int, blockNumber []*big.Int) (*FukuastakeWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}

	logs, sub, err := _Fukuastake.contract.FilterLogs(opts, "Withdraw", userRule, poolIdRule, blockNumberRule)
	if err != nil {
		return nil, err
	}
	return &FukuastakeWithdrawIterator{contract: _Fukuastake.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed poolId, uint256 amount, uint256 indexed blockNumber)
func (_Fukuastake *FukuastakeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *FukuastakeWithdraw, user []common.Address, poolId []*big.Int, blockNumber []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}

	logs, sub, err := _Fukuastake.contract.WatchLogs(opts, "Withdraw", userRule, poolIdRule, blockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FukuastakeWithdraw)
				if err := _Fukuastake.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed poolId, uint256 amount, uint256 indexed blockNumber)
func (_Fukuastake *FukuastakeFilterer) ParseWithdraw(log types.Log) (*FukuastakeWithdraw, error) {
	event := new(FukuastakeWithdraw)
	if err := _Fukuastake.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
