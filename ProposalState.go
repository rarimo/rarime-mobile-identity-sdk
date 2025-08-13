// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package identity

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

// ProposalsStateProposalConfig is an auto generated low-level Go binding around an user-defined struct.
type ProposalsStateProposalConfig struct {
	StartTimestamp      uint64
	Duration            uint64
	Multichoice         *big.Int
	AcceptedOptions     []*big.Int
	Description         string
	VotingWhitelist     []common.Address
	VotingWhitelistData [][]byte
}

// ProposalsStateProposalInfo is an auto generated low-level Go binding around an user-defined struct.
type ProposalsStateProposalInfo struct {
	ProposalSMT   common.Address
	Status        uint8
	Config        ProposalsStateProposalConfig
	VotingResults [][8]*big.Int
}

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposalSMT\",\"type\":\"address\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"hide\",\"type\":\"bool\"}],\"name\":\"ProposalHidden\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userNullifier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"vote\",\"type\":\"uint256[]\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAGIC_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_CHOICES_PER_OPTION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_OPTIONS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"P\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"proposalSMTImpl_\",\"type\":\"address\"}],\"name\":\"__ProposalsState_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"votingName_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"votingAddress_\",\"type\":\"address\"}],\"name\":\"addVoting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"multichoice\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"acceptedOptions\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"votingWhitelist\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"votingWhitelistData\",\"type\":\"bytes[]\"}],\"internalType\":\"structProposalsState.ProposalConfig\",\"name\":\"newProposalConfig_\",\"type\":\"tuple\"}],\"name\":\"changeProposalConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"newSignerPubKey_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"changeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"multichoice\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"acceptedOptions\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"votingWhitelist\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"votingWhitelistData\",\"type\":\"bytes[]\"}],\"internalType\":\"structProposalsState.ProposalConfig\",\"name\":\"proposalConfig_\",\"type\":\"tuple\"}],\"name\":\"createProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId_\",\"type\":\"uint256\"}],\"name\":\"getProposalConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"multichoice\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"acceptedOptions\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"votingWhitelist\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"votingWhitelistData\",\"type\":\"bytes[]\"}],\"internalType\":\"structProposalsState.ProposalConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId_\",\"type\":\"uint256\"}],\"name\":\"getProposalEventId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId_\",\"type\":\"uint256\"}],\"name\":\"getProposalInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"proposalSMT\",\"type\":\"address\"},{\"internalType\":\"enumProposalsState.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"multichoice\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"acceptedOptions\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"votingWhitelist\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"votingWhitelistData\",\"type\":\"bytes[]\"}],\"internalType\":\"structProposalsState.ProposalConfig\",\"name\":\"config\",\"type\":\"tuple\"},{\"internalType\":\"uint256[8][]\",\"name\":\"votingResults\",\"type\":\"uint256[8][]\"}],\"internalType\":\"structProposalsState.ProposalInfo\",\"name\":\"info_\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId_\",\"type\":\"uint256\"}],\"name\":\"getProposalStatus\",\"outputs\":[{\"internalType\":\"enumProposalsState.ProposalStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key_\",\"type\":\"string\"}],\"name\":\"getVotingByKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVotings\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"keys_\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"values_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hide_\",\"type\":\"bool\"}],\"name\":\"hideProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voting_\",\"type\":\"address\"}],\"name\":\"isVoting\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastProposalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalSMTImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"votingName_\",\"type\":\"string\"}],\"name\":\"removeVoting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCallWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"}],\"name\":\"upgradeToWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userNullifier_\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"vote_\",\"type\":\"uint256[]\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// MAGICID is a free data retrieval call binding the contract method 0xdf95574a.
//
// Solidity: function MAGIC_ID() view returns(uint8)
func (_Storage *StorageCaller) MAGICID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "MAGIC_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAGICID is a free data retrieval call binding the contract method 0xdf95574a.
//
// Solidity: function MAGIC_ID() view returns(uint8)
func (_Storage *StorageSession) MAGICID() (uint8, error) {
	return _Storage.Contract.MAGICID(&_Storage.CallOpts)
}

// MAGICID is a free data retrieval call binding the contract method 0xdf95574a.
//
// Solidity: function MAGIC_ID() view returns(uint8)
func (_Storage *StorageCallerSession) MAGICID() (uint8, error) {
	return _Storage.Contract.MAGICID(&_Storage.CallOpts)
}

// MAXIMUMCHOICESPEROPTION is a free data retrieval call binding the contract method 0x299a9b54.
//
// Solidity: function MAXIMUM_CHOICES_PER_OPTION() view returns(uint256)
func (_Storage *StorageCaller) MAXIMUMCHOICESPEROPTION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "MAXIMUM_CHOICES_PER_OPTION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXIMUMCHOICESPEROPTION is a free data retrieval call binding the contract method 0x299a9b54.
//
// Solidity: function MAXIMUM_CHOICES_PER_OPTION() view returns(uint256)
func (_Storage *StorageSession) MAXIMUMCHOICESPEROPTION() (*big.Int, error) {
	return _Storage.Contract.MAXIMUMCHOICESPEROPTION(&_Storage.CallOpts)
}

// MAXIMUMCHOICESPEROPTION is a free data retrieval call binding the contract method 0x299a9b54.
//
// Solidity: function MAXIMUM_CHOICES_PER_OPTION() view returns(uint256)
func (_Storage *StorageCallerSession) MAXIMUMCHOICESPEROPTION() (*big.Int, error) {
	return _Storage.Contract.MAXIMUMCHOICESPEROPTION(&_Storage.CallOpts)
}

// MAXIMUMOPTIONS is a free data retrieval call binding the contract method 0x881e3447.
//
// Solidity: function MAXIMUM_OPTIONS() view returns(uint256)
func (_Storage *StorageCaller) MAXIMUMOPTIONS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "MAXIMUM_OPTIONS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXIMUMOPTIONS is a free data retrieval call binding the contract method 0x881e3447.
//
// Solidity: function MAXIMUM_OPTIONS() view returns(uint256)
func (_Storage *StorageSession) MAXIMUMOPTIONS() (*big.Int, error) {
	return _Storage.Contract.MAXIMUMOPTIONS(&_Storage.CallOpts)
}

// MAXIMUMOPTIONS is a free data retrieval call binding the contract method 0x881e3447.
//
// Solidity: function MAXIMUM_OPTIONS() view returns(uint256)
func (_Storage *StorageCallerSession) MAXIMUMOPTIONS() (*big.Int, error) {
	return _Storage.Contract.MAXIMUMOPTIONS(&_Storage.CallOpts)
}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_Storage *StorageCaller) P(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "P")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_Storage *StorageSession) P() (*big.Int, error) {
	return _Storage.Contract.P(&_Storage.CallOpts)
}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_Storage *StorageCallerSession) P() (*big.Int, error) {
	return _Storage.Contract.P(&_Storage.CallOpts)
}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_Storage *StorageCaller) ChainName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "chainName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_Storage *StorageSession) ChainName() (string, error) {
	return _Storage.Contract.ChainName(&_Storage.CallOpts)
}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_Storage *StorageCallerSession) ChainName() (string, error) {
	return _Storage.Contract.ChainName(&_Storage.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xf4fc6341.
//
// Solidity: function getNonce(uint8 methodId_) view returns(uint256)
func (_Storage *StorageCaller) GetNonce(opts *bind.CallOpts, methodId_ uint8) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getNonce", methodId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xf4fc6341.
//
// Solidity: function getNonce(uint8 methodId_) view returns(uint256)
func (_Storage *StorageSession) GetNonce(methodId_ uint8) (*big.Int, error) {
	return _Storage.Contract.GetNonce(&_Storage.CallOpts, methodId_)
}

// GetNonce is a free data retrieval call binding the contract method 0xf4fc6341.
//
// Solidity: function getNonce(uint8 methodId_) view returns(uint256)
func (_Storage *StorageCallerSession) GetNonce(methodId_ uint8) (*big.Int, error) {
	return _Storage.Contract.GetNonce(&_Storage.CallOpts, methodId_)
}

// GetProposalConfig is a free data retrieval call binding the contract method 0x7d5d687f.
//
// Solidity: function getProposalConfig(uint256 proposalId_) view returns((uint64,uint64,uint256,uint256[],string,address[],bytes[]))
func (_Storage *StorageCaller) GetProposalConfig(opts *bind.CallOpts, proposalId_ *big.Int) (ProposalsStateProposalConfig, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getProposalConfig", proposalId_)

	if err != nil {
		return *new(ProposalsStateProposalConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ProposalsStateProposalConfig)).(*ProposalsStateProposalConfig)

	return out0, err

}

// GetProposalConfig is a free data retrieval call binding the contract method 0x7d5d687f.
//
// Solidity: function getProposalConfig(uint256 proposalId_) view returns((uint64,uint64,uint256,uint256[],string,address[],bytes[]))
func (_Storage *StorageSession) GetProposalConfig(proposalId_ *big.Int) (ProposalsStateProposalConfig, error) {
	return _Storage.Contract.GetProposalConfig(&_Storage.CallOpts, proposalId_)
}

// GetProposalConfig is a free data retrieval call binding the contract method 0x7d5d687f.
//
// Solidity: function getProposalConfig(uint256 proposalId_) view returns((uint64,uint64,uint256,uint256[],string,address[],bytes[]))
func (_Storage *StorageCallerSession) GetProposalConfig(proposalId_ *big.Int) (ProposalsStateProposalConfig, error) {
	return _Storage.Contract.GetProposalConfig(&_Storage.CallOpts, proposalId_)
}

// GetProposalEventId is a free data retrieval call binding the contract method 0x31e181c5.
//
// Solidity: function getProposalEventId(uint256 proposalId_) view returns(uint256)
func (_Storage *StorageCaller) GetProposalEventId(opts *bind.CallOpts, proposalId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getProposalEventId", proposalId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProposalEventId is a free data retrieval call binding the contract method 0x31e181c5.
//
// Solidity: function getProposalEventId(uint256 proposalId_) view returns(uint256)
func (_Storage *StorageSession) GetProposalEventId(proposalId_ *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetProposalEventId(&_Storage.CallOpts, proposalId_)
}

// GetProposalEventId is a free data retrieval call binding the contract method 0x31e181c5.
//
// Solidity: function getProposalEventId(uint256 proposalId_) view returns(uint256)
func (_Storage *StorageCallerSession) GetProposalEventId(proposalId_ *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetProposalEventId(&_Storage.CallOpts, proposalId_)
}

// GetProposalInfo is a free data retrieval call binding the contract method 0xbc903cb8.
//
// Solidity: function getProposalInfo(uint256 proposalId_) view returns((address,uint8,(uint64,uint64,uint256,uint256[],string,address[],bytes[]),uint256[8][]) info_)
func (_Storage *StorageCaller) GetProposalInfo(opts *bind.CallOpts, proposalId_ *big.Int) (ProposalsStateProposalInfo, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getProposalInfo", proposalId_)

	if err != nil {
		return *new(ProposalsStateProposalInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ProposalsStateProposalInfo)).(*ProposalsStateProposalInfo)

	return out0, err

}

// GetProposalInfo is a free data retrieval call binding the contract method 0xbc903cb8.
//
// Solidity: function getProposalInfo(uint256 proposalId_) view returns((address,uint8,(uint64,uint64,uint256,uint256[],string,address[],bytes[]),uint256[8][]) info_)
func (_Storage *StorageSession) GetProposalInfo(proposalId_ *big.Int) (ProposalsStateProposalInfo, error) {
	return _Storage.Contract.GetProposalInfo(&_Storage.CallOpts, proposalId_)
}

// GetProposalInfo is a free data retrieval call binding the contract method 0xbc903cb8.
//
// Solidity: function getProposalInfo(uint256 proposalId_) view returns((address,uint8,(uint64,uint64,uint256,uint256[],string,address[],bytes[]),uint256[8][]) info_)
func (_Storage *StorageCallerSession) GetProposalInfo(proposalId_ *big.Int) (ProposalsStateProposalInfo, error) {
	return _Storage.Contract.GetProposalInfo(&_Storage.CallOpts, proposalId_)
}

// GetProposalStatus is a free data retrieval call binding the contract method 0x401853b7.
//
// Solidity: function getProposalStatus(uint256 proposalId_) view returns(uint8)
func (_Storage *StorageCaller) GetProposalStatus(opts *bind.CallOpts, proposalId_ *big.Int) (uint8, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getProposalStatus", proposalId_)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetProposalStatus is a free data retrieval call binding the contract method 0x401853b7.
//
// Solidity: function getProposalStatus(uint256 proposalId_) view returns(uint8)
func (_Storage *StorageSession) GetProposalStatus(proposalId_ *big.Int) (uint8, error) {
	return _Storage.Contract.GetProposalStatus(&_Storage.CallOpts, proposalId_)
}

// GetProposalStatus is a free data retrieval call binding the contract method 0x401853b7.
//
// Solidity: function getProposalStatus(uint256 proposalId_) view returns(uint8)
func (_Storage *StorageCallerSession) GetProposalStatus(proposalId_ *big.Int) (uint8, error) {
	return _Storage.Contract.GetProposalStatus(&_Storage.CallOpts, proposalId_)
}

// GetVotingByKey is a free data retrieval call binding the contract method 0xd8720106.
//
// Solidity: function getVotingByKey(string key_) view returns(address)
func (_Storage *StorageCaller) GetVotingByKey(opts *bind.CallOpts, key_ string) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getVotingByKey", key_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVotingByKey is a free data retrieval call binding the contract method 0xd8720106.
//
// Solidity: function getVotingByKey(string key_) view returns(address)
func (_Storage *StorageSession) GetVotingByKey(key_ string) (common.Address, error) {
	return _Storage.Contract.GetVotingByKey(&_Storage.CallOpts, key_)
}

// GetVotingByKey is a free data retrieval call binding the contract method 0xd8720106.
//
// Solidity: function getVotingByKey(string key_) view returns(address)
func (_Storage *StorageCallerSession) GetVotingByKey(key_ string) (common.Address, error) {
	return _Storage.Contract.GetVotingByKey(&_Storage.CallOpts, key_)
}

// GetVotings is a free data retrieval call binding the contract method 0x05c1112d.
//
// Solidity: function getVotings() view returns(string[] keys_, address[] values_)
func (_Storage *StorageCaller) GetVotings(opts *bind.CallOpts) (struct {
	Keys   []string
	Values []common.Address
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getVotings")

	outstruct := new(struct {
		Keys   []string
		Values []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Keys = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.Values = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetVotings is a free data retrieval call binding the contract method 0x05c1112d.
//
// Solidity: function getVotings() view returns(string[] keys_, address[] values_)
func (_Storage *StorageSession) GetVotings() (struct {
	Keys   []string
	Values []common.Address
}, error) {
	return _Storage.Contract.GetVotings(&_Storage.CallOpts)
}

// GetVotings is a free data retrieval call binding the contract method 0x05c1112d.
//
// Solidity: function getVotings() view returns(string[] keys_, address[] values_)
func (_Storage *StorageCallerSession) GetVotings() (struct {
	Keys   []string
	Values []common.Address
}, error) {
	return _Storage.Contract.GetVotings(&_Storage.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Storage *StorageCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Storage *StorageSession) Implementation() (common.Address, error) {
	return _Storage.Contract.Implementation(&_Storage.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Storage *StorageCallerSession) Implementation() (common.Address, error) {
	return _Storage.Contract.Implementation(&_Storage.CallOpts)
}

// IsVoting is a free data retrieval call binding the contract method 0x5f8dd649.
//
// Solidity: function isVoting(address voting_) view returns(bool)
func (_Storage *StorageCaller) IsVoting(opts *bind.CallOpts, voting_ common.Address) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "isVoting", voting_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoting is a free data retrieval call binding the contract method 0x5f8dd649.
//
// Solidity: function isVoting(address voting_) view returns(bool)
func (_Storage *StorageSession) IsVoting(voting_ common.Address) (bool, error) {
	return _Storage.Contract.IsVoting(&_Storage.CallOpts, voting_)
}

// IsVoting is a free data retrieval call binding the contract method 0x5f8dd649.
//
// Solidity: function isVoting(address voting_) view returns(bool)
func (_Storage *StorageCallerSession) IsVoting(voting_ common.Address) (bool, error) {
	return _Storage.Contract.IsVoting(&_Storage.CallOpts, voting_)
}

// LastProposalId is a free data retrieval call binding the contract method 0x74cb3041.
//
// Solidity: function lastProposalId() view returns(uint256)
func (_Storage *StorageCaller) LastProposalId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "lastProposalId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastProposalId is a free data retrieval call binding the contract method 0x74cb3041.
//
// Solidity: function lastProposalId() view returns(uint256)
func (_Storage *StorageSession) LastProposalId() (*big.Int, error) {
	return _Storage.Contract.LastProposalId(&_Storage.CallOpts)
}

// LastProposalId is a free data retrieval call binding the contract method 0x74cb3041.
//
// Solidity: function lastProposalId() view returns(uint256)
func (_Storage *StorageCallerSession) LastProposalId() (*big.Int, error) {
	return _Storage.Contract.LastProposalId(&_Storage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageCallerSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// ProposalSMTImpl is a free data retrieval call binding the contract method 0x59917f46.
//
// Solidity: function proposalSMTImpl() view returns(address)
func (_Storage *StorageCaller) ProposalSMTImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "proposalSMTImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalSMTImpl is a free data retrieval call binding the contract method 0x59917f46.
//
// Solidity: function proposalSMTImpl() view returns(address)
func (_Storage *StorageSession) ProposalSMTImpl() (common.Address, error) {
	return _Storage.Contract.ProposalSMTImpl(&_Storage.CallOpts)
}

// ProposalSMTImpl is a free data retrieval call binding the contract method 0x59917f46.
//
// Solidity: function proposalSMTImpl() view returns(address)
func (_Storage *StorageCallerSession) ProposalSMTImpl() (common.Address, error) {
	return _Storage.Contract.ProposalSMTImpl(&_Storage.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Storage *StorageCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Storage *StorageSession) ProxiableUUID() ([32]byte, error) {
	return _Storage.Contract.ProxiableUUID(&_Storage.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Storage *StorageCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Storage.Contract.ProxiableUUID(&_Storage.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Storage *StorageCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Storage *StorageSession) Signer() (common.Address, error) {
	return _Storage.Contract.Signer(&_Storage.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Storage *StorageCallerSession) Signer() (common.Address, error) {
	return _Storage.Contract.Signer(&_Storage.CallOpts)
}

// ProposalsStateInit is a paid mutator transaction binding the contract method 0xb5697e6a.
//
// Solidity: function __ProposalsState_init(address signer_, string chainName_, address proposalSMTImpl_) returns()
func (_Storage *StorageTransactor) ProposalsStateInit(opts *bind.TransactOpts, signer_ common.Address, chainName_ string, proposalSMTImpl_ common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "__ProposalsState_init", signer_, chainName_, proposalSMTImpl_)
}

// ProposalsStateInit is a paid mutator transaction binding the contract method 0xb5697e6a.
//
// Solidity: function __ProposalsState_init(address signer_, string chainName_, address proposalSMTImpl_) returns()
func (_Storage *StorageSession) ProposalsStateInit(signer_ common.Address, chainName_ string, proposalSMTImpl_ common.Address) (*types.Transaction, error) {
	return _Storage.Contract.ProposalsStateInit(&_Storage.TransactOpts, signer_, chainName_, proposalSMTImpl_)
}

// ProposalsStateInit is a paid mutator transaction binding the contract method 0xb5697e6a.
//
// Solidity: function __ProposalsState_init(address signer_, string chainName_, address proposalSMTImpl_) returns()
func (_Storage *StorageTransactorSession) ProposalsStateInit(signer_ common.Address, chainName_ string, proposalSMTImpl_ common.Address) (*types.Transaction, error) {
	return _Storage.Contract.ProposalsStateInit(&_Storage.TransactOpts, signer_, chainName_, proposalSMTImpl_)
}

// AddVoting is a paid mutator transaction binding the contract method 0xde947541.
//
// Solidity: function addVoting(string votingName_, address votingAddress_) returns()
func (_Storage *StorageTransactor) AddVoting(opts *bind.TransactOpts, votingName_ string, votingAddress_ common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addVoting", votingName_, votingAddress_)
}

// AddVoting is a paid mutator transaction binding the contract method 0xde947541.
//
// Solidity: function addVoting(string votingName_, address votingAddress_) returns()
func (_Storage *StorageSession) AddVoting(votingName_ string, votingAddress_ common.Address) (*types.Transaction, error) {
	return _Storage.Contract.AddVoting(&_Storage.TransactOpts, votingName_, votingAddress_)
}

// AddVoting is a paid mutator transaction binding the contract method 0xde947541.
//
// Solidity: function addVoting(string votingName_, address votingAddress_) returns()
func (_Storage *StorageTransactorSession) AddVoting(votingName_ string, votingAddress_ common.Address) (*types.Transaction, error) {
	return _Storage.Contract.AddVoting(&_Storage.TransactOpts, votingName_, votingAddress_)
}

// ChangeProposalConfig is a paid mutator transaction binding the contract method 0x1cbdbbc0.
//
// Solidity: function changeProposalConfig(uint256 proposalId_, (uint64,uint64,uint256,uint256[],string,address[],bytes[]) newProposalConfig_) returns()
func (_Storage *StorageTransactor) ChangeProposalConfig(opts *bind.TransactOpts, proposalId_ *big.Int, newProposalConfig_ ProposalsStateProposalConfig) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "changeProposalConfig", proposalId_, newProposalConfig_)
}

// ChangeProposalConfig is a paid mutator transaction binding the contract method 0x1cbdbbc0.
//
// Solidity: function changeProposalConfig(uint256 proposalId_, (uint64,uint64,uint256,uint256[],string,address[],bytes[]) newProposalConfig_) returns()
func (_Storage *StorageSession) ChangeProposalConfig(proposalId_ *big.Int, newProposalConfig_ ProposalsStateProposalConfig) (*types.Transaction, error) {
	return _Storage.Contract.ChangeProposalConfig(&_Storage.TransactOpts, proposalId_, newProposalConfig_)
}

// ChangeProposalConfig is a paid mutator transaction binding the contract method 0x1cbdbbc0.
//
// Solidity: function changeProposalConfig(uint256 proposalId_, (uint64,uint64,uint256,uint256[],string,address[],bytes[]) newProposalConfig_) returns()
func (_Storage *StorageTransactorSession) ChangeProposalConfig(proposalId_ *big.Int, newProposalConfig_ ProposalsStateProposalConfig) (*types.Transaction, error) {
	return _Storage.Contract.ChangeProposalConfig(&_Storage.TransactOpts, proposalId_, newProposalConfig_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_Storage *StorageTransactor) ChangeSigner(opts *bind.TransactOpts, newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "changeSigner", newSignerPubKey_, signature_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_Storage *StorageSession) ChangeSigner(newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _Storage.Contract.ChangeSigner(&_Storage.TransactOpts, newSignerPubKey_, signature_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_Storage *StorageTransactorSession) ChangeSigner(newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _Storage.Contract.ChangeSigner(&_Storage.TransactOpts, newSignerPubKey_, signature_)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x9151b81f.
//
// Solidity: function createProposal((uint64,uint64,uint256,uint256[],string,address[],bytes[]) proposalConfig_) returns()
func (_Storage *StorageTransactor) CreateProposal(opts *bind.TransactOpts, proposalConfig_ ProposalsStateProposalConfig) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "createProposal", proposalConfig_)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x9151b81f.
//
// Solidity: function createProposal((uint64,uint64,uint256,uint256[],string,address[],bytes[]) proposalConfig_) returns()
func (_Storage *StorageSession) CreateProposal(proposalConfig_ ProposalsStateProposalConfig) (*types.Transaction, error) {
	return _Storage.Contract.CreateProposal(&_Storage.TransactOpts, proposalConfig_)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x9151b81f.
//
// Solidity: function createProposal((uint64,uint64,uint256,uint256[],string,address[],bytes[]) proposalConfig_) returns()
func (_Storage *StorageTransactorSession) CreateProposal(proposalConfig_ ProposalsStateProposalConfig) (*types.Transaction, error) {
	return _Storage.Contract.CreateProposal(&_Storage.TransactOpts, proposalConfig_)
}

// HideProposal is a paid mutator transaction binding the contract method 0x50df86a3.
//
// Solidity: function hideProposal(uint256 proposalId_, bool hide_) returns()
func (_Storage *StorageTransactor) HideProposal(opts *bind.TransactOpts, proposalId_ *big.Int, hide_ bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "hideProposal", proposalId_, hide_)
}

// HideProposal is a paid mutator transaction binding the contract method 0x50df86a3.
//
// Solidity: function hideProposal(uint256 proposalId_, bool hide_) returns()
func (_Storage *StorageSession) HideProposal(proposalId_ *big.Int, hide_ bool) (*types.Transaction, error) {
	return _Storage.Contract.HideProposal(&_Storage.TransactOpts, proposalId_, hide_)
}

// HideProposal is a paid mutator transaction binding the contract method 0x50df86a3.
//
// Solidity: function hideProposal(uint256 proposalId_, bool hide_) returns()
func (_Storage *StorageTransactorSession) HideProposal(proposalId_ *big.Int, hide_ bool) (*types.Transaction, error) {
	return _Storage.Contract.HideProposal(&_Storage.TransactOpts, proposalId_, hide_)
}

// RemoveVoting is a paid mutator transaction binding the contract method 0x4fcfccd7.
//
// Solidity: function removeVoting(string votingName_) returns()
func (_Storage *StorageTransactor) RemoveVoting(opts *bind.TransactOpts, votingName_ string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "removeVoting", votingName_)
}

// RemoveVoting is a paid mutator transaction binding the contract method 0x4fcfccd7.
//
// Solidity: function removeVoting(string votingName_) returns()
func (_Storage *StorageSession) RemoveVoting(votingName_ string) (*types.Transaction, error) {
	return _Storage.Contract.RemoveVoting(&_Storage.TransactOpts, votingName_)
}

// RemoveVoting is a paid mutator transaction binding the contract method 0x4fcfccd7.
//
// Solidity: function removeVoting(string votingName_) returns()
func (_Storage *StorageTransactorSession) RemoveVoting(votingName_ string) (*types.Transaction, error) {
	return _Storage.Contract.RemoveVoting(&_Storage.TransactOpts, votingName_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _Storage.Contract.RenounceOwnership(&_Storage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Storage.Contract.RenounceOwnership(&_Storage.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.TransferOwnership(&_Storage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.TransferOwnership(&_Storage.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Storage *StorageTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Storage *StorageSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Storage.Contract.UpgradeTo(&_Storage.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Storage *StorageTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Storage.Contract.UpgradeTo(&_Storage.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Storage *StorageTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Storage *StorageSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.UpgradeToAndCall(&_Storage.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Storage *StorageTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.UpgradeToAndCall(&_Storage.TransactOpts, newImplementation, data)
}

// UpgradeToAndCallWithProof is a paid mutator transaction binding the contract method 0xbf2c6db7.
//
// Solidity: function upgradeToAndCallWithProof(address newImplementation_, bytes proof_, bytes data_) returns()
func (_Storage *StorageTransactor) UpgradeToAndCallWithProof(opts *bind.TransactOpts, newImplementation_ common.Address, proof_ []byte, data_ []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "upgradeToAndCallWithProof", newImplementation_, proof_, data_)
}

// UpgradeToAndCallWithProof is a paid mutator transaction binding the contract method 0xbf2c6db7.
//
// Solidity: function upgradeToAndCallWithProof(address newImplementation_, bytes proof_, bytes data_) returns()
func (_Storage *StorageSession) UpgradeToAndCallWithProof(newImplementation_ common.Address, proof_ []byte, data_ []byte) (*types.Transaction, error) {
	return _Storage.Contract.UpgradeToAndCallWithProof(&_Storage.TransactOpts, newImplementation_, proof_, data_)
}

// UpgradeToAndCallWithProof is a paid mutator transaction binding the contract method 0xbf2c6db7.
//
// Solidity: function upgradeToAndCallWithProof(address newImplementation_, bytes proof_, bytes data_) returns()
func (_Storage *StorageTransactorSession) UpgradeToAndCallWithProof(newImplementation_ common.Address, proof_ []byte, data_ []byte) (*types.Transaction, error) {
	return _Storage.Contract.UpgradeToAndCallWithProof(&_Storage.TransactOpts, newImplementation_, proof_, data_)
}

// UpgradeToWithProof is a paid mutator transaction binding the contract method 0x628543ab.
//
// Solidity: function upgradeToWithProof(address newImplementation_, bytes proof_) returns()
func (_Storage *StorageTransactor) UpgradeToWithProof(opts *bind.TransactOpts, newImplementation_ common.Address, proof_ []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "upgradeToWithProof", newImplementation_, proof_)
}

// UpgradeToWithProof is a paid mutator transaction binding the contract method 0x628543ab.
//
// Solidity: function upgradeToWithProof(address newImplementation_, bytes proof_) returns()
func (_Storage *StorageSession) UpgradeToWithProof(newImplementation_ common.Address, proof_ []byte) (*types.Transaction, error) {
	return _Storage.Contract.UpgradeToWithProof(&_Storage.TransactOpts, newImplementation_, proof_)
}

// UpgradeToWithProof is a paid mutator transaction binding the contract method 0x628543ab.
//
// Solidity: function upgradeToWithProof(address newImplementation_, bytes proof_) returns()
func (_Storage *StorageTransactorSession) UpgradeToWithProof(newImplementation_ common.Address, proof_ []byte) (*types.Transaction, error) {
	return _Storage.Contract.UpgradeToWithProof(&_Storage.TransactOpts, newImplementation_, proof_)
}

// Vote is a paid mutator transaction binding the contract method 0xe1349bd7.
//
// Solidity: function vote(uint256 proposalId_, uint256 userNullifier_, uint256[] vote_) returns()
func (_Storage *StorageTransactor) Vote(opts *bind.TransactOpts, proposalId_ *big.Int, userNullifier_ *big.Int, vote_ []*big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "vote", proposalId_, userNullifier_, vote_)
}

// Vote is a paid mutator transaction binding the contract method 0xe1349bd7.
//
// Solidity: function vote(uint256 proposalId_, uint256 userNullifier_, uint256[] vote_) returns()
func (_Storage *StorageSession) Vote(proposalId_ *big.Int, userNullifier_ *big.Int, vote_ []*big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Vote(&_Storage.TransactOpts, proposalId_, userNullifier_, vote_)
}

// Vote is a paid mutator transaction binding the contract method 0xe1349bd7.
//
// Solidity: function vote(uint256 proposalId_, uint256 userNullifier_, uint256[] vote_) returns()
func (_Storage *StorageTransactorSession) Vote(proposalId_ *big.Int, userNullifier_ *big.Int, vote_ []*big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Vote(&_Storage.TransactOpts, proposalId_, userNullifier_, vote_)
}

// StorageAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Storage contract.
type StorageAdminChangedIterator struct {
	Event *StorageAdminChanged // Event containing the contract specifics and raw log

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
func (it *StorageAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageAdminChanged)
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
		it.Event = new(StorageAdminChanged)
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
func (it *StorageAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageAdminChanged represents a AdminChanged event raised by the Storage contract.
type StorageAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Storage *StorageFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*StorageAdminChangedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &StorageAdminChangedIterator{contract: _Storage.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Storage *StorageFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *StorageAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageAdminChanged)
				if err := _Storage.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Storage *StorageFilterer) ParseAdminChanged(log types.Log) (*StorageAdminChanged, error) {
	event := new(StorageAdminChanged)
	if err := _Storage.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Storage contract.
type StorageBeaconUpgradedIterator struct {
	Event *StorageBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *StorageBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageBeaconUpgraded)
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
		it.Event = new(StorageBeaconUpgraded)
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
func (it *StorageBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageBeaconUpgraded represents a BeaconUpgraded event raised by the Storage contract.
type StorageBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Storage *StorageFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*StorageBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &StorageBeaconUpgradedIterator{contract: _Storage.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Storage *StorageFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *StorageBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageBeaconUpgraded)
				if err := _Storage.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Storage *StorageFilterer) ParseBeaconUpgraded(log types.Log) (*StorageBeaconUpgraded, error) {
	event := new(StorageBeaconUpgraded)
	if err := _Storage.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Storage contract.
type StorageInitializedIterator struct {
	Event *StorageInitialized // Event containing the contract specifics and raw log

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
func (it *StorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageInitialized)
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
		it.Event = new(StorageInitialized)
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
func (it *StorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageInitialized represents a Initialized event raised by the Storage contract.
type StorageInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Storage *StorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*StorageInitializedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StorageInitializedIterator{contract: _Storage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Storage *StorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StorageInitialized) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageInitialized)
				if err := _Storage.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Storage *StorageFilterer) ParseInitialized(log types.Log) (*StorageInitialized, error) {
	event := new(StorageInitialized)
	if err := _Storage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Storage contract.
type StorageOwnershipTransferredIterator struct {
	Event *StorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageOwnershipTransferred)
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
		it.Event = new(StorageOwnershipTransferred)
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
func (it *StorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageOwnershipTransferred represents a OwnershipTransferred event raised by the Storage contract.
type StorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Storage *StorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StorageOwnershipTransferredIterator{contract: _Storage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Storage *StorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageOwnershipTransferred)
				if err := _Storage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Storage *StorageFilterer) ParseOwnershipTransferred(log types.Log) (*StorageOwnershipTransferred, error) {
	event := new(StorageOwnershipTransferred)
	if err := _Storage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProposalConfigChangedIterator is returned from FilterProposalConfigChanged and is used to iterate over the raw logs and unpacked data for ProposalConfigChanged events raised by the Storage contract.
type StorageProposalConfigChangedIterator struct {
	Event *StorageProposalConfigChanged // Event containing the contract specifics and raw log

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
func (it *StorageProposalConfigChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProposalConfigChanged)
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
		it.Event = new(StorageProposalConfigChanged)
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
func (it *StorageProposalConfigChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProposalConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProposalConfigChanged represents a ProposalConfigChanged event raised by the Storage contract.
type StorageProposalConfigChanged struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalConfigChanged is a free log retrieval operation binding the contract event 0xa9cc646240fc6ba1b4b124e96765839b67cd0e2e698942d5d5948a36c7b998d5.
//
// Solidity: event ProposalConfigChanged(uint256 indexed proposalId)
func (_Storage *StorageFilterer) FilterProposalConfigChanged(opts *bind.FilterOpts, proposalId []*big.Int) (*StorageProposalConfigChangedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "ProposalConfigChanged", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &StorageProposalConfigChangedIterator{contract: _Storage.contract, event: "ProposalConfigChanged", logs: logs, sub: sub}, nil
}

// WatchProposalConfigChanged is a free log subscription operation binding the contract event 0xa9cc646240fc6ba1b4b124e96765839b67cd0e2e698942d5d5948a36c7b998d5.
//
// Solidity: event ProposalConfigChanged(uint256 indexed proposalId)
func (_Storage *StorageFilterer) WatchProposalConfigChanged(opts *bind.WatchOpts, sink chan<- *StorageProposalConfigChanged, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "ProposalConfigChanged", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProposalConfigChanged)
				if err := _Storage.contract.UnpackLog(event, "ProposalConfigChanged", log); err != nil {
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

// ParseProposalConfigChanged is a log parse operation binding the contract event 0xa9cc646240fc6ba1b4b124e96765839b67cd0e2e698942d5d5948a36c7b998d5.
//
// Solidity: event ProposalConfigChanged(uint256 indexed proposalId)
func (_Storage *StorageFilterer) ParseProposalConfigChanged(log types.Log) (*StorageProposalConfigChanged, error) {
	event := new(StorageProposalConfigChanged)
	if err := _Storage.contract.UnpackLog(event, "ProposalConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the Storage contract.
type StorageProposalCreatedIterator struct {
	Event *StorageProposalCreated // Event containing the contract specifics and raw log

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
func (it *StorageProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProposalCreated)
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
		it.Event = new(StorageProposalCreated)
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
func (it *StorageProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProposalCreated represents a ProposalCreated event raised by the Storage contract.
type StorageProposalCreated struct {
	ProposalId  *big.Int
	ProposalSMT common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0xcd423cc1203c0af96b9b3d68d73b3064a69de2d14450bb7181c5e5df2132b358.
//
// Solidity: event ProposalCreated(uint256 indexed proposalId, address proposalSMT)
func (_Storage *StorageFilterer) FilterProposalCreated(opts *bind.FilterOpts, proposalId []*big.Int) (*StorageProposalCreatedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "ProposalCreated", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &StorageProposalCreatedIterator{contract: _Storage.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xcd423cc1203c0af96b9b3d68d73b3064a69de2d14450bb7181c5e5df2132b358.
//
// Solidity: event ProposalCreated(uint256 indexed proposalId, address proposalSMT)
func (_Storage *StorageFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *StorageProposalCreated, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "ProposalCreated", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProposalCreated)
				if err := _Storage.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0xcd423cc1203c0af96b9b3d68d73b3064a69de2d14450bb7181c5e5df2132b358.
//
// Solidity: event ProposalCreated(uint256 indexed proposalId, address proposalSMT)
func (_Storage *StorageFilterer) ParseProposalCreated(log types.Log) (*StorageProposalCreated, error) {
	event := new(StorageProposalCreated)
	if err := _Storage.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProposalHiddenIterator is returned from FilterProposalHidden and is used to iterate over the raw logs and unpacked data for ProposalHidden events raised by the Storage contract.
type StorageProposalHiddenIterator struct {
	Event *StorageProposalHidden // Event containing the contract specifics and raw log

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
func (it *StorageProposalHiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProposalHidden)
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
		it.Event = new(StorageProposalHidden)
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
func (it *StorageProposalHiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProposalHiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProposalHidden represents a ProposalHidden event raised by the Storage contract.
type StorageProposalHidden struct {
	ProposalId *big.Int
	Hide       bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalHidden is a free log retrieval operation binding the contract event 0x19f289534dc0123bc632723043c2eaf220105017950314fd948774c69282dc93.
//
// Solidity: event ProposalHidden(uint256 indexed proposalId, bool hide)
func (_Storage *StorageFilterer) FilterProposalHidden(opts *bind.FilterOpts, proposalId []*big.Int) (*StorageProposalHiddenIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "ProposalHidden", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &StorageProposalHiddenIterator{contract: _Storage.contract, event: "ProposalHidden", logs: logs, sub: sub}, nil
}

// WatchProposalHidden is a free log subscription operation binding the contract event 0x19f289534dc0123bc632723043c2eaf220105017950314fd948774c69282dc93.
//
// Solidity: event ProposalHidden(uint256 indexed proposalId, bool hide)
func (_Storage *StorageFilterer) WatchProposalHidden(opts *bind.WatchOpts, sink chan<- *StorageProposalHidden, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "ProposalHidden", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProposalHidden)
				if err := _Storage.contract.UnpackLog(event, "ProposalHidden", log); err != nil {
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

// ParseProposalHidden is a log parse operation binding the contract event 0x19f289534dc0123bc632723043c2eaf220105017950314fd948774c69282dc93.
//
// Solidity: event ProposalHidden(uint256 indexed proposalId, bool hide)
func (_Storage *StorageFilterer) ParseProposalHidden(log types.Log) (*StorageProposalHidden, error) {
	event := new(StorageProposalHidden)
	if err := _Storage.contract.UnpackLog(event, "ProposalHidden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Storage contract.
type StorageUpgradedIterator struct {
	Event *StorageUpgraded // Event containing the contract specifics and raw log

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
func (it *StorageUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageUpgraded)
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
		it.Event = new(StorageUpgraded)
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
func (it *StorageUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageUpgraded represents a Upgraded event raised by the Storage contract.
type StorageUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Storage *StorageFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StorageUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StorageUpgradedIterator{contract: _Storage.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Storage *StorageFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StorageUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageUpgraded)
				if err := _Storage.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Storage *StorageFilterer) ParseUpgraded(log types.Log) (*StorageUpgraded, error) {
	event := new(StorageUpgraded)
	if err := _Storage.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the Storage contract.
type StorageVoteCastIterator struct {
	Event *StorageVoteCast // Event containing the contract specifics and raw log

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
func (it *StorageVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageVoteCast)
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
		it.Event = new(StorageVoteCast)
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
func (it *StorageVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageVoteCast represents a VoteCast event raised by the Storage contract.
type StorageVoteCast struct {
	ProposalId    *big.Int
	UserNullifier *big.Int
	Vote          []*big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0x82e882ecc8f666d65e8120d1fb3859261652f808e5001ae2f169e5ea1bf5035c.
//
// Solidity: event VoteCast(uint256 indexed proposalId, uint256 userNullifier, uint256[] vote)
func (_Storage *StorageFilterer) FilterVoteCast(opts *bind.FilterOpts, proposalId []*big.Int) (*StorageVoteCastIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "VoteCast", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &StorageVoteCastIterator{contract: _Storage.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0x82e882ecc8f666d65e8120d1fb3859261652f808e5001ae2f169e5ea1bf5035c.
//
// Solidity: event VoteCast(uint256 indexed proposalId, uint256 userNullifier, uint256[] vote)
func (_Storage *StorageFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *StorageVoteCast, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "VoteCast", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageVoteCast)
				if err := _Storage.contract.UnpackLog(event, "VoteCast", log); err != nil {
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

// ParseVoteCast is a log parse operation binding the contract event 0x82e882ecc8f666d65e8120d1fb3859261652f808e5001ae2f169e5ea1bf5035c.
//
// Solidity: event VoteCast(uint256 indexed proposalId, uint256 userNullifier, uint256[] vote)
func (_Storage *StorageFilterer) ParseVoteCast(log types.Log) (*StorageVoteCast, error) {
	event := new(StorageVoteCast)
	if err := _Storage.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
