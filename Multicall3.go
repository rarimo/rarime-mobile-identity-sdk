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

// Multicall3Call is an auto generated low-level Go binding around an user-defined struct.
type Multicall3Call struct {
	Target   common.Address
	CallData []byte
}

// Multicall3Call3 is an auto generated low-level Go binding around an user-defined struct.
type Multicall3Call3 struct {
	Target       common.Address
	AllowFailure bool
	CallData     []byte
}

// Multicall3Call3Value is an auto generated low-level Go binding around an user-defined struct.
type Multicall3Call3Value struct {
	Target       common.Address
	AllowFailure bool
	Value        *big.Int
	CallData     []byte
}

// Multicall3Result is an auto generated low-level Go binding around an user-defined struct.
type Multicall3Result struct {
	Success    bool
	ReturnData []byte
}

// MulticallStorageMetaData contains all meta data concerning the MulticallStorage contract.
var MulticallStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call3[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate3\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call3Value[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate3Value\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"blockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBasefee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"basefee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryAggregate\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryBlockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MulticallStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use MulticallStorageMetaData.ABI instead.
var MulticallStorageABI = MulticallStorageMetaData.ABI

// MulticallStorage is an auto generated Go binding around an Ethereum contract.
type MulticallStorage struct {
	MulticallStorageCaller     // Read-only binding to the contract
	MulticallStorageTransactor // Write-only binding to the contract
	MulticallStorageFilterer   // Log filterer for contract events
}

// MulticallStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type MulticallStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MulticallStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MulticallStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MulticallStorageSession struct {
	Contract     *MulticallStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MulticallStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MulticallStorageCallerSession struct {
	Contract *MulticallStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MulticallStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MulticallStorageTransactorSession struct {
	Contract     *MulticallStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MulticallStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type MulticallStorageRaw struct {
	Contract *MulticallStorage // Generic contract binding to access the raw methods on
}

// MulticallStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MulticallStorageCallerRaw struct {
	Contract *MulticallStorageCaller // Generic read-only contract binding to access the raw methods on
}

// MulticallStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MulticallStorageTransactorRaw struct {
	Contract *MulticallStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMulticallStorage creates a new instance of MulticallStorage, bound to a specific deployed contract.
func NewMulticallStorage(address common.Address, backend bind.ContractBackend) (*MulticallStorage, error) {
	contract, err := bindMulticallStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MulticallStorage{MulticallStorageCaller: MulticallStorageCaller{contract: contract}, MulticallStorageTransactor: MulticallStorageTransactor{contract: contract}, MulticallStorageFilterer: MulticallStorageFilterer{contract: contract}}, nil
}

// NewMulticallStorageCaller creates a new read-only instance of MulticallStorage, bound to a specific deployed contract.
func NewMulticallStorageCaller(address common.Address, caller bind.ContractCaller) (*MulticallStorageCaller, error) {
	contract, err := bindMulticallStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallStorageCaller{contract: contract}, nil
}

// NewMulticallStorageTransactor creates a new write-only instance of MulticallStorage, bound to a specific deployed contract.
func NewMulticallStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*MulticallStorageTransactor, error) {
	contract, err := bindMulticallStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallStorageTransactor{contract: contract}, nil
}

// NewMulticallStorageFilterer creates a new log filterer instance of MulticallStorage, bound to a specific deployed contract.
func NewMulticallStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*MulticallStorageFilterer, error) {
	contract, err := bindMulticallStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MulticallStorageFilterer{contract: contract}, nil
}

// bindMulticallStorage binds a generic wrapper to an already deployed contract.
func bindMulticallStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MulticallStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MulticallStorage *MulticallStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MulticallStorage.Contract.MulticallStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MulticallStorage *MulticallStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MulticallStorage.Contract.MulticallStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MulticallStorage *MulticallStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MulticallStorage.Contract.MulticallStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MulticallStorage *MulticallStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MulticallStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MulticallStorage *MulticallStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MulticallStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MulticallStorage *MulticallStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MulticallStorage.Contract.contract.Transact(opts, method, params...)
}

// Aggregate is a free data retrieval call binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) view returns(uint256 blockNumber, bytes[] returnData)
func (_MulticallStorage *MulticallStorageCaller) Aggregate(opts *bind.CallOpts, calls []Multicall3Call) (struct {
	BlockNumber *big.Int
	ReturnData  [][]byte
}, error) {
	var out []interface{}
	err := _MulticallStorage.contract.Call(opts, &out, "aggregate", calls)

	outstruct := new(struct {
		BlockNumber *big.Int
		ReturnData  [][]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReturnData = *abi.ConvertType(out[1], new([][]byte)).(*[][]byte)

	return *outstruct, err

}

// Aggregate is a free data retrieval call binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) view returns(uint256 blockNumber, bytes[] returnData)
func (_MulticallStorage *MulticallStorageSession) Aggregate(calls []Multicall3Call) (struct {
	BlockNumber *big.Int
	ReturnData  [][]byte
}, error) {
	return _MulticallStorage.Contract.Aggregate(&_MulticallStorage.CallOpts, calls)
}

// Aggregate is a free data retrieval call binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) view returns(uint256 blockNumber, bytes[] returnData)
func (_MulticallStorage *MulticallStorageCallerSession) Aggregate(calls []Multicall3Call) (struct {
	BlockNumber *big.Int
	ReturnData  [][]byte
}, error) {
	return _MulticallStorage.Contract.Aggregate(&_MulticallStorage.CallOpts, calls)
}

// Aggregate3 is a free data retrieval call binding the contract method 0x82ad56cb.
//
// Solidity: function aggregate3((address,bool,bytes)[] calls) view returns((bool,bytes)[] returnData)
func (_MulticallStorage *MulticallStorageCaller) Aggregate3(opts *bind.CallOpts, calls []Multicall3Call3) ([]Multicall3Result, error) {
	var out []interface{}
	err := _MulticallStorage.contract.Call(opts, &out, "aggregate3", calls)

	if err != nil {
		return *new([]Multicall3Result), err
	}

	out0 := *abi.ConvertType(out[0], new([]Multicall3Result)).(*[]Multicall3Result)

	return out0, err

}

// Aggregate3 is a free data retrieval call binding the contract method 0x82ad56cb.
//
// Solidity: function aggregate3((address,bool,bytes)[] calls) view returns((bool,bytes)[] returnData)
func (_MulticallStorage *MulticallStorageSession) Aggregate3(calls []Multicall3Call3) ([]Multicall3Result, error) {
	return _MulticallStorage.Contract.Aggregate3(&_MulticallStorage.CallOpts, calls)
}

// Aggregate3 is a free data retrieval call binding the contract method 0x82ad56cb.
//
// Solidity: function aggregate3((address,bool,bytes)[] calls) view returns((bool,bytes)[] returnData)
func (_MulticallStorage *MulticallStorageCallerSession) Aggregate3(calls []Multicall3Call3) ([]Multicall3Result, error) {
	return _MulticallStorage.Contract.Aggregate3(&_MulticallStorage.CallOpts, calls)
}

// Aggregate3Value is a free data retrieval call binding the contract method 0x174dea71.
//
// Solidity: function aggregate3Value((address,bool,uint256,bytes)[] calls) view returns((bool,bytes)[] returnData)
func (_MulticallStorage *MulticallStorageCaller) Aggregate3Value(opts *bind.CallOpts, calls []Multicall3Call3Value) ([]Multicall3Result, error) {
	var out []interface{}
	err := _MulticallStorage.contract.Call(opts, &out, "aggregate3Value", calls)

	if err != nil {
		return *new([]Multicall3Result), err
	}

	out0 := *abi.ConvertType(out[0], new([]Multicall3Result)).(*[]Multicall3Result)

	return out0, err

}

// Aggregate3Value is a free data retrieval call binding the contract method 0x174dea71.
//
// Solidity: function aggregate3Value((address,bool,uint256,bytes)[] calls) view returns((bool,bytes)[] returnData)
func (_MulticallStorage *MulticallStorageSession) Aggregate3Value(calls []Multicall3Call3Value) ([]Multicall3Result, error) {
	return _MulticallStorage.Contract.Aggregate3Value(&_MulticallStorage.CallOpts, calls)
}

// Aggregate3Value is a free data retrieval call binding the contract method 0x174dea71.
//
// Solidity: function aggregate3Value((address,bool,uint256,bytes)[] calls) view returns((bool,bytes)[] returnData)
func (_MulticallStorage *MulticallStorageCallerSession) Aggregate3Value(calls []Multicall3Call3Value) ([]Multicall3Result, error) {
	return _MulticallStorage.Contract.Aggregate3Value(&_MulticallStorage.CallOpts, calls)
}

// BlockAndAggregate is a free data retrieval call binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) view returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MulticallStorage *MulticallStorageCaller) BlockAndAggregate(opts *bind.CallOpts, calls []Multicall3Call) (struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	ReturnData  []Multicall3Result
}, error) {
	var out []interface{}
	err := _MulticallStorage.contract.Call(opts, &out, "blockAndAggregate", calls)

	outstruct := new(struct {
		BlockNumber *big.Int
		BlockHash   [32]byte
		ReturnData  []Multicall3Result
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlockHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ReturnData = *abi.ConvertType(out[2], new([]Multicall3Result)).(*[]Multicall3Result)

	return *outstruct, err

}

// BlockAndAggregate is a free data retrieval call binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) view returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MulticallStorage *MulticallStorageSession) BlockAndAggregate(calls []Multicall3Call) (struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	ReturnData  []Multicall3Result
}, error) {
	return _MulticallStorage.Contract.BlockAndAggregate(&_MulticallStorage.CallOpts, calls)
}

// BlockAndAggregate is a free data retrieval call binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) view returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MulticallStorage *MulticallStorageCallerSession) BlockAndAggregate(calls []Multicall3Call) (struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	ReturnData  []Multicall3Result
}, error) {
	return _MulticallStorage.Contract.BlockAndAggregate(&_MulticallStorage.CallOpts, calls)
}

// GetBasefee is a free data retrieval call binding the contract method 0x3e64a696.
//
// Solidity: function getBasefee() view returns(uint256 basefee)
func (_MulticallStorage *MulticallStorageCaller) GetBasefee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MulticallStorage.contract.Call(opts, &out, "getBasefee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBasefee is a free data retrieval call binding the contract method 0x3e64a696.
//
// Solidity: function getBasefee() view returns(uint256 basefee)
func (_MulticallStorage *MulticallStorageSession) GetBasefee() (*big.Int, error) {
	return _MulticallStorage.Contract.GetBasefee(&_MulticallStorage.CallOpts)
}

// GetBasefee is a free data retrieval call binding the contract method 0x3e64a696.
//
// Solidity: function getBasefee() view returns(uint256 basefee)
func (_MulticallStorage *MulticallStorageCallerSession) GetBasefee() (*big.Int, error) {
	return _MulticallStorage.Contract.GetBasefee(&_MulticallStorage.CallOpts)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MulticallStorage *MulticallStorageCaller) GetBlockHash(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MulticallStorage.contract.Call(opts, &out, "getBlockHash", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MulticallStorage *MulticallStorageSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _MulticallStorage.Contract.GetBlockHash(&_MulticallStorage.CallOpts, blockNumber)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MulticallStorage *MulticallStorageCallerSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _MulticallStorage.Contract.GetBlockHash(&_MulticallStorage.CallOpts, blockNumber)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_MulticallStorage *MulticallStorageCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MulticallStorage.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_MulticallStorage *MulticallStorageSession) GetBlockNumber() (*big.Int, error) {
	return _MulticallStorage.Contract.GetBlockNumber(&_MulticallStorage.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_MulticallStorage *MulticallStorageCallerSession) GetBlockNumber() (*big.Int, error) {
	return _MulticallStorage.Contract.GetBlockNumber(&_MulticallStorage.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainid)
func (_MulticallStorage *MulticallStorageCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MulticallStorage.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainid)
func (_MulticallStorage *MulticallStorageSession) GetChainId() (*big.Int, error) {
	return _MulticallStorage.Contract.GetChainId(&_MulticallStorage.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainid)
func (_MulticallStorage *MulticallStorageCallerSession) GetChainId() (*big.Int, error) {
	return _MulticallStorage.Contract.GetChainId(&_MulticallStorage.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MulticallStorage *MulticallStorageCaller) GetCurrentBlockCoinbase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MulticallStorage.contract.Call(opts, &out, "getCurrentBlockCoinbase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MulticallStorage *MulticallStorageSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _MulticallStorage.Contract.GetCurrentBlockCoinbase(&_MulticallStorage.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MulticallStorage *MulticallStorageCallerSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _MulticallStorage.Contract.GetCurrentBlockCoinbase(&_MulticallStorage.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MulticallStorage *MulticallStorageCaller) GetCurrentBlockDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MulticallStorage.contract.Call(opts, &out, "getCurrentBlockDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MulticallStorage *MulticallStorageSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _MulticallStorage.Contract.GetCurrentBlockDifficulty(&_MulticallStorage.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MulticallStorage *MulticallStorageCallerSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _MulticallStorage.Contract.GetCurrentBlockDifficulty(&_MulticallStorage.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MulticallStorage *MulticallStorageCaller) GetCurrentBlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MulticallStorage.contract.Call(opts, &out, "getCurrentBlockGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MulticallStorage *MulticallStorageSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _MulticallStorage.Contract.GetCurrentBlockGasLimit(&_MulticallStorage.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MulticallStorage *MulticallStorageCallerSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _MulticallStorage.Contract.GetCurrentBlockGasLimit(&_MulticallStorage.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MulticallStorage *MulticallStorageCaller) GetCurrentBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MulticallStorage.contract.Call(opts, &out, "getCurrentBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MulticallStorage *MulticallStorageSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _MulticallStorage.Contract.GetCurrentBlockTimestamp(&_MulticallStorage.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MulticallStorage *MulticallStorageCallerSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _MulticallStorage.Contract.GetCurrentBlockTimestamp(&_MulticallStorage.CallOpts)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MulticallStorage *MulticallStorageCaller) GetEthBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MulticallStorage.contract.Call(opts, &out, "getEthBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MulticallStorage *MulticallStorageSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _MulticallStorage.Contract.GetEthBalance(&_MulticallStorage.CallOpts, addr)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MulticallStorage *MulticallStorageCallerSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _MulticallStorage.Contract.GetEthBalance(&_MulticallStorage.CallOpts, addr)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MulticallStorage *MulticallStorageCaller) GetLastBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MulticallStorage.contract.Call(opts, &out, "getLastBlockHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MulticallStorage *MulticallStorageSession) GetLastBlockHash() ([32]byte, error) {
	return _MulticallStorage.Contract.GetLastBlockHash(&_MulticallStorage.CallOpts)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MulticallStorage *MulticallStorageCallerSession) GetLastBlockHash() ([32]byte, error) {
	return _MulticallStorage.Contract.GetLastBlockHash(&_MulticallStorage.CallOpts)
}

// TryAggregate is a free data retrieval call binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) view returns((bool,bytes)[] returnData)
func (_MulticallStorage *MulticallStorageCaller) TryAggregate(opts *bind.CallOpts, requireSuccess bool, calls []Multicall3Call) ([]Multicall3Result, error) {
	var out []interface{}
	err := _MulticallStorage.contract.Call(opts, &out, "tryAggregate", requireSuccess, calls)

	if err != nil {
		return *new([]Multicall3Result), err
	}

	out0 := *abi.ConvertType(out[0], new([]Multicall3Result)).(*[]Multicall3Result)

	return out0, err

}

// TryAggregate is a free data retrieval call binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) view returns((bool,bytes)[] returnData)
func (_MulticallStorage *MulticallStorageSession) TryAggregate(requireSuccess bool, calls []Multicall3Call) ([]Multicall3Result, error) {
	return _MulticallStorage.Contract.TryAggregate(&_MulticallStorage.CallOpts, requireSuccess, calls)
}

// TryAggregate is a free data retrieval call binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) view returns((bool,bytes)[] returnData)
func (_MulticallStorage *MulticallStorageCallerSession) TryAggregate(requireSuccess bool, calls []Multicall3Call) ([]Multicall3Result, error) {
	return _MulticallStorage.Contract.TryAggregate(&_MulticallStorage.CallOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a free data retrieval call binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) view returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MulticallStorage *MulticallStorageCaller) TryBlockAndAggregate(opts *bind.CallOpts, requireSuccess bool, calls []Multicall3Call) (struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	ReturnData  []Multicall3Result
}, error) {
	var out []interface{}
	err := _MulticallStorage.contract.Call(opts, &out, "tryBlockAndAggregate", requireSuccess, calls)

	outstruct := new(struct {
		BlockNumber *big.Int
		BlockHash   [32]byte
		ReturnData  []Multicall3Result
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlockHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ReturnData = *abi.ConvertType(out[2], new([]Multicall3Result)).(*[]Multicall3Result)

	return *outstruct, err

}

// TryBlockAndAggregate is a free data retrieval call binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) view returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MulticallStorage *MulticallStorageSession) TryBlockAndAggregate(requireSuccess bool, calls []Multicall3Call) (struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	ReturnData  []Multicall3Result
}, error) {
	return _MulticallStorage.Contract.TryBlockAndAggregate(&_MulticallStorage.CallOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a free data retrieval call binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) view returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MulticallStorage *MulticallStorageCallerSession) TryBlockAndAggregate(requireSuccess bool, calls []Multicall3Call) (struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	ReturnData  []Multicall3Result
}, error) {
	return _MulticallStorage.Contract.TryBlockAndAggregate(&_MulticallStorage.CallOpts, requireSuccess, calls)
}
