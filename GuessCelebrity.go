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

// Groth16VerifierHelperProofPoints is an auto generated low-level Go binding around an user-defined struct.
type Groth16VerifierHelperProofPoints struct {
	A [2]*big.Int
	B [2][2]*big.Int
	C [2]*big.Int
}

// GuessCelebrityABI is the input ABI used to generate the binding from.
// Deprecated: Use GuessCelebrityMetaData.ABI instead.
var GuessCelebrityABI = GuessCelebrityMetaData.ABI

// GuessCelebrity is an auto generated Go binding around an Ethereum contract.
type GuessCelebrity struct {
	GuessCelebrityCaller     // Read-only binding to the contract
	GuessCelebrityTransactor // Write-only binding to the contract
	GuessCelebrityFilterer   // Log filterer for contract events
}

// GuessCelebrityCaller is an auto generated read-only Go binding around an Ethereum contract.
type GuessCelebrityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuessCelebrityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GuessCelebrityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuessCelebrityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GuessCelebrityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuessCelebritySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GuessCelebritySession struct {
	Contract     *GuessCelebrity   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GuessCelebrityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GuessCelebrityCallerSession struct {
	Contract *GuessCelebrityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GuessCelebrityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GuessCelebrityTransactorSession struct {
	Contract     *GuessCelebrityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GuessCelebrityRaw is an auto generated low-level Go binding around an Ethereum contract.
type GuessCelebrityRaw struct {
	Contract *GuessCelebrity // Generic contract binding to access the raw methods on
}

// GuessCelebrityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GuessCelebrityCallerRaw struct {
	Contract *GuessCelebrityCaller // Generic read-only contract binding to access the raw methods on
}

// GuessCelebrityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GuessCelebrityTransactorRaw struct {
	Contract *GuessCelebrityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGuessCelebrity creates a new instance of GuessCelebrity, bound to a specific deployed contract.
func NewGuessCelebrity(address common.Address, backend bind.ContractBackend) (*GuessCelebrity, error) {
	contract, err := bindGuessCelebrity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GuessCelebrity{GuessCelebrityCaller: GuessCelebrityCaller{contract: contract}, GuessCelebrityTransactor: GuessCelebrityTransactor{contract: contract}, GuessCelebrityFilterer: GuessCelebrityFilterer{contract: contract}}, nil
}

// NewGuessCelebrityCaller creates a new read-only instance of GuessCelebrity, bound to a specific deployed contract.
func NewGuessCelebrityCaller(address common.Address, caller bind.ContractCaller) (*GuessCelebrityCaller, error) {
	contract, err := bindGuessCelebrity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GuessCelebrityCaller{contract: contract}, nil
}

// NewGuessCelebrityTransactor creates a new write-only instance of GuessCelebrity, bound to a specific deployed contract.
func NewGuessCelebrityTransactor(address common.Address, transactor bind.ContractTransactor) (*GuessCelebrityTransactor, error) {
	contract, err := bindGuessCelebrity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GuessCelebrityTransactor{contract: contract}, nil
}

// NewGuessCelebrityFilterer creates a new log filterer instance of GuessCelebrity, bound to a specific deployed contract.
func NewGuessCelebrityFilterer(address common.Address, filterer bind.ContractFilterer) (*GuessCelebrityFilterer, error) {
	contract, err := bindGuessCelebrity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GuessCelebrityFilterer{contract: contract}, nil
}

// bindGuessCelebrity binds a generic wrapper to an already deployed contract.
func bindGuessCelebrity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GuessCelebrityMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GuessCelebrity *GuessCelebrityRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GuessCelebrity.Contract.GuessCelebrityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GuessCelebrity *GuessCelebrityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GuessCelebrity.Contract.GuessCelebrityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GuessCelebrity *GuessCelebrityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GuessCelebrity.Contract.GuessCelebrityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GuessCelebrity *GuessCelebrityCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GuessCelebrity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GuessCelebrity *GuessCelebrityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GuessCelebrity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GuessCelebrity *GuessCelebrityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GuessCelebrity.Contract.contract.Transact(opts, method, params...)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xbd942515.
//
// Solidity: function claimReward(address recipient_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_GuessCelebrity *GuessCelebrityTransactor) ClaimReward(opts *bind.TransactOpts, recipient_ common.Address, zkPoints_ Groth16VerifierHelperProofPoints) (*types.Transaction, error) {
	return _GuessCelebrity.contract.Transact(opts, "claimReward", recipient_, zkPoints_)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xbd942515.
//
// Solidity: function claimReward(address recipient_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_GuessCelebrity *GuessCelebritySession) ClaimReward(recipient_ common.Address, zkPoints_ Groth16VerifierHelperProofPoints) (*types.Transaction, error) {
	return _GuessCelebrity.Contract.ClaimReward(&_GuessCelebrity.TransactOpts, recipient_, zkPoints_)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xbd942515.
//
// Solidity: function claimReward(address recipient_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_GuessCelebrity *GuessCelebrityTransactorSession) ClaimReward(recipient_ common.Address, zkPoints_ Groth16VerifierHelperProofPoints) (*types.Transaction, error) {
	return _GuessCelebrity.Contract.ClaimReward(&_GuessCelebrity.TransactOpts, recipient_, zkPoints_)
}
