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

// IdentityMetaData contains all meta data concerning the Identity contract.
var IdentityMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_collectionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dayMonthTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yearTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"categoryTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"craftingCardTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"CollectionCrafted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousCrafter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCrafter\",\"type\":\"address\"}],\"name\":\"CraftingContractChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"PartCrafted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"categoryContract\",\"outputs\":[{\"internalType\":\"contractIPart\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newCrafter\",\"type\":\"address\"}],\"name\":\"changeCrafter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"changeToCrafted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dayMonthContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_yearContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_categoryContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_craftingCardContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dayMonthTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"yearTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"categoryTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"craftingCardTokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"craftCollection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"crafted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"craftingCardContract\",\"outputs\":[{\"internalType\":\"contractIPart\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dayMonthContract\",\"outputs\":[{\"internalType\":\"contractIPart\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"}],\"name\":\"getCollectionTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isCrafted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yearContract\",\"outputs\":[{\"internalType\":\"contractIPart\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IdentityABI is the input ABI used to generate the binding from.
// Deprecated: Use IdentityMetaData.ABI instead.
var IdentityABI = IdentityMetaData.ABI

// Identity is an auto generated Go binding around an Ethereum contract.
type Identity struct {
	IdentityCaller     // Read-only binding to the contract
	IdentityTransactor // Write-only binding to the contract
	IdentityFilterer   // Log filterer for contract events
}

// IdentityCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentitySession struct {
	Contract     *Identity         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityCallerSession struct {
	Contract *IdentityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IdentityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityTransactorSession struct {
	Contract     *IdentityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IdentityRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityRaw struct {
	Contract *Identity // Generic contract binding to access the raw methods on
}

// IdentityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityCallerRaw struct {
	Contract *IdentityCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityTransactorRaw struct {
	Contract *IdentityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentity creates a new instance of Identity, bound to a specific deployed contract.
func NewIdentity(address common.Address, backend bind.ContractBackend) (*Identity, error) {
	contract, err := bindIdentity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Identity{IdentityCaller: IdentityCaller{contract: contract}, IdentityTransactor: IdentityTransactor{contract: contract}, IdentityFilterer: IdentityFilterer{contract: contract}}, nil
}

// NewIdentityCaller creates a new read-only instance of Identity, bound to a specific deployed contract.
func NewIdentityCaller(address common.Address, caller bind.ContractCaller) (*IdentityCaller, error) {
	contract, err := bindIdentity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityCaller{contract: contract}, nil
}

// NewIdentityTransactor creates a new write-only instance of Identity, bound to a specific deployed contract.
func NewIdentityTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityTransactor, error) {
	contract, err := bindIdentity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityTransactor{contract: contract}, nil
}

// NewIdentityFilterer creates a new log filterer instance of Identity, bound to a specific deployed contract.
func NewIdentityFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityFilterer, error) {
	contract, err := bindIdentity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityFilterer{contract: contract}, nil
}

// bindIdentity binds a generic wrapper to an already deployed contract.
func bindIdentity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IdentityMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identity *IdentityRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Identity.Contract.IdentityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identity *IdentityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.Contract.IdentityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identity *IdentityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identity.Contract.IdentityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identity *IdentityCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Identity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identity *IdentityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identity *IdentityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identity.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Identity *IdentityCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Identity *IdentitySession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Identity.Contract.BalanceOf(&_Identity.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Identity *IdentityCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Identity.Contract.BalanceOf(&_Identity.CallOpts, owner)
}

// CategoryContract is a free data retrieval call binding the contract method 0xe4f0bdcc.
//
// Solidity: function categoryContract() view returns(address)
func (_Identity *IdentityCaller) CategoryContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "categoryContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CategoryContract is a free data retrieval call binding the contract method 0xe4f0bdcc.
//
// Solidity: function categoryContract() view returns(address)
func (_Identity *IdentitySession) CategoryContract() (common.Address, error) {
	return _Identity.Contract.CategoryContract(&_Identity.CallOpts)
}

// CategoryContract is a free data retrieval call binding the contract method 0xe4f0bdcc.
//
// Solidity: function categoryContract() view returns(address)
func (_Identity *IdentityCallerSession) CategoryContract() (common.Address, error) {
	return _Identity.Contract.CategoryContract(&_Identity.CallOpts)
}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_Identity *IdentityCaller) CollectionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "collectionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_Identity *IdentitySession) CollectionId() (*big.Int, error) {
	return _Identity.Contract.CollectionId(&_Identity.CallOpts)
}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_Identity *IdentityCallerSession) CollectionId() (*big.Int, error) {
	return _Identity.Contract.CollectionId(&_Identity.CallOpts)
}

// Crafted is a free data retrieval call binding the contract method 0xcac4cc29.
//
// Solidity: function crafted(uint256 ) view returns(bool)
func (_Identity *IdentityCaller) Crafted(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "crafted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Crafted is a free data retrieval call binding the contract method 0xcac4cc29.
//
// Solidity: function crafted(uint256 ) view returns(bool)
func (_Identity *IdentitySession) Crafted(arg0 *big.Int) (bool, error) {
	return _Identity.Contract.Crafted(&_Identity.CallOpts, arg0)
}

// Crafted is a free data retrieval call binding the contract method 0xcac4cc29.
//
// Solidity: function crafted(uint256 ) view returns(bool)
func (_Identity *IdentityCallerSession) Crafted(arg0 *big.Int) (bool, error) {
	return _Identity.Contract.Crafted(&_Identity.CallOpts, arg0)
}

// CraftingCardContract is a free data retrieval call binding the contract method 0x2c948296.
//
// Solidity: function craftingCardContract() view returns(address)
func (_Identity *IdentityCaller) CraftingCardContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "craftingCardContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CraftingCardContract is a free data retrieval call binding the contract method 0x2c948296.
//
// Solidity: function craftingCardContract() view returns(address)
func (_Identity *IdentitySession) CraftingCardContract() (common.Address, error) {
	return _Identity.Contract.CraftingCardContract(&_Identity.CallOpts)
}

// CraftingCardContract is a free data retrieval call binding the contract method 0x2c948296.
//
// Solidity: function craftingCardContract() view returns(address)
func (_Identity *IdentityCallerSession) CraftingCardContract() (common.Address, error) {
	return _Identity.Contract.CraftingCardContract(&_Identity.CallOpts)
}

// DayMonthContract is a free data retrieval call binding the contract method 0x9829ba1c.
//
// Solidity: function dayMonthContract() view returns(address)
func (_Identity *IdentityCaller) DayMonthContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "dayMonthContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DayMonthContract is a free data retrieval call binding the contract method 0x9829ba1c.
//
// Solidity: function dayMonthContract() view returns(address)
func (_Identity *IdentitySession) DayMonthContract() (common.Address, error) {
	return _Identity.Contract.DayMonthContract(&_Identity.CallOpts)
}

// DayMonthContract is a free data retrieval call binding the contract method 0x9829ba1c.
//
// Solidity: function dayMonthContract() view returns(address)
func (_Identity *IdentityCallerSession) DayMonthContract() (common.Address, error) {
	return _Identity.Contract.DayMonthContract(&_Identity.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Identity *IdentityCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Identity *IdentitySession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Identity.Contract.GetApproved(&_Identity.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Identity *IdentityCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Identity.Contract.GetApproved(&_Identity.CallOpts, tokenId)
}

// GetCollectionTokens is a free data retrieval call binding the contract method 0x2bb88056.
//
// Solidity: function getCollectionTokens(uint256 collectionId) view returns(uint256, uint256, uint256, uint256)
func (_Identity *IdentityCaller) GetCollectionTokens(opts *bind.CallOpts, collectionId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getCollectionTokens", collectionId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetCollectionTokens is a free data retrieval call binding the contract method 0x2bb88056.
//
// Solidity: function getCollectionTokens(uint256 collectionId) view returns(uint256, uint256, uint256, uint256)
func (_Identity *IdentitySession) GetCollectionTokens(collectionId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Identity.Contract.GetCollectionTokens(&_Identity.CallOpts, collectionId)
}

// GetCollectionTokens is a free data retrieval call binding the contract method 0x2bb88056.
//
// Solidity: function getCollectionTokens(uint256 collectionId) view returns(uint256, uint256, uint256, uint256)
func (_Identity *IdentityCallerSession) GetCollectionTokens(collectionId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Identity.Contract.GetCollectionTokens(&_Identity.CallOpts, collectionId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Identity *IdentityCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Identity *IdentitySession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Identity.Contract.IsApprovedForAll(&_Identity.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Identity *IdentityCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Identity.Contract.IsApprovedForAll(&_Identity.CallOpts, owner, operator)
}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_Identity *IdentityCaller) IsCrafted(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "isCrafted", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_Identity *IdentitySession) IsCrafted(tokenId *big.Int) (bool, error) {
	return _Identity.Contract.IsCrafted(&_Identity.CallOpts, tokenId)
}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_Identity *IdentityCallerSession) IsCrafted(tokenId *big.Int) (bool, error) {
	return _Identity.Contract.IsCrafted(&_Identity.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Identity *IdentityCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Identity *IdentitySession) Name() (string, error) {
	return _Identity.Contract.Name(&_Identity.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Identity *IdentityCallerSession) Name() (string, error) {
	return _Identity.Contract.Name(&_Identity.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Identity *IdentityCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Identity *IdentitySession) Owner() (common.Address, error) {
	return _Identity.Contract.Owner(&_Identity.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Identity *IdentityCallerSession) Owner() (common.Address, error) {
	return _Identity.Contract.Owner(&_Identity.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Identity *IdentityCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Identity *IdentitySession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Identity.Contract.OwnerOf(&_Identity.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Identity *IdentityCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Identity.Contract.OwnerOf(&_Identity.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Identity *IdentityCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Identity *IdentitySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Identity.Contract.SupportsInterface(&_Identity.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Identity *IdentityCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Identity.Contract.SupportsInterface(&_Identity.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Identity *IdentityCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Identity *IdentitySession) Symbol() (string, error) {
	return _Identity.Contract.Symbol(&_Identity.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Identity *IdentityCallerSession) Symbol() (string, error) {
	return _Identity.Contract.Symbol(&_Identity.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Identity *IdentityCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Identity *IdentitySession) TokenURI(tokenId *big.Int) (string, error) {
	return _Identity.Contract.TokenURI(&_Identity.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Identity *IdentityCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Identity.Contract.TokenURI(&_Identity.CallOpts, tokenId)
}

// YearContract is a free data retrieval call binding the contract method 0x7230f8d4.
//
// Solidity: function yearContract() view returns(address)
func (_Identity *IdentityCaller) YearContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "yearContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YearContract is a free data retrieval call binding the contract method 0x7230f8d4.
//
// Solidity: function yearContract() view returns(address)
func (_Identity *IdentitySession) YearContract() (common.Address, error) {
	return _Identity.Contract.YearContract(&_Identity.CallOpts)
}

// YearContract is a free data retrieval call binding the contract method 0x7230f8d4.
//
// Solidity: function yearContract() view returns(address)
func (_Identity *IdentityCallerSession) YearContract() (common.Address, error) {
	return _Identity.Contract.YearContract(&_Identity.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Identity *IdentityTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Identity *IdentitySession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.Approve(&_Identity.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Identity *IdentityTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.Approve(&_Identity.TransactOpts, to, tokenId)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Identity *IdentityTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Identity *IdentitySession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Identity.Contract.ChangeAdmin(&_Identity.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Identity *IdentityTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Identity.Contract.ChangeAdmin(&_Identity.TransactOpts, _newAdmin)
}

// ChangeCrafter is a paid mutator transaction binding the contract method 0xe7d3b069.
//
// Solidity: function changeCrafter(address _newCrafter) returns()
func (_Identity *IdentityTransactor) ChangeCrafter(opts *bind.TransactOpts, _newCrafter common.Address) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "changeCrafter", _newCrafter)
}

// ChangeCrafter is a paid mutator transaction binding the contract method 0xe7d3b069.
//
// Solidity: function changeCrafter(address _newCrafter) returns()
func (_Identity *IdentitySession) ChangeCrafter(_newCrafter common.Address) (*types.Transaction, error) {
	return _Identity.Contract.ChangeCrafter(&_Identity.TransactOpts, _newCrafter)
}

// ChangeCrafter is a paid mutator transaction binding the contract method 0xe7d3b069.
//
// Solidity: function changeCrafter(address _newCrafter) returns()
func (_Identity *IdentityTransactorSession) ChangeCrafter(_newCrafter common.Address) (*types.Transaction, error) {
	return _Identity.Contract.ChangeCrafter(&_Identity.TransactOpts, _newCrafter)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_Identity *IdentityTransactor) ChangeToCrafted(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "changeToCrafted", tokenId)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_Identity *IdentitySession) ChangeToCrafted(tokenId *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.ChangeToCrafted(&_Identity.TransactOpts, tokenId)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_Identity *IdentityTransactorSession) ChangeToCrafted(tokenId *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.ChangeToCrafted(&_Identity.TransactOpts, tokenId)
}

// CraftCollection is a paid mutator transaction binding the contract method 0xbf0837fe.
//
// Solidity: function craftCollection(address _dayMonthContract, address _yearContract, address _categoryContract, address _craftingCardContract, uint256 dayMonthTokenId, uint256 yearTokenId, uint256 categoryTokenId, uint256 craftingCardTokenId, string _tokenURI) returns()
func (_Identity *IdentityTransactor) CraftCollection(opts *bind.TransactOpts, _dayMonthContract common.Address, _yearContract common.Address, _categoryContract common.Address, _craftingCardContract common.Address, dayMonthTokenId *big.Int, yearTokenId *big.Int, categoryTokenId *big.Int, craftingCardTokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "craftCollection", _dayMonthContract, _yearContract, _categoryContract, _craftingCardContract, dayMonthTokenId, yearTokenId, categoryTokenId, craftingCardTokenId, _tokenURI)
}

// CraftCollection is a paid mutator transaction binding the contract method 0xbf0837fe.
//
// Solidity: function craftCollection(address _dayMonthContract, address _yearContract, address _categoryContract, address _craftingCardContract, uint256 dayMonthTokenId, uint256 yearTokenId, uint256 categoryTokenId, uint256 craftingCardTokenId, string _tokenURI) returns()
func (_Identity *IdentitySession) CraftCollection(_dayMonthContract common.Address, _yearContract common.Address, _categoryContract common.Address, _craftingCardContract common.Address, dayMonthTokenId *big.Int, yearTokenId *big.Int, categoryTokenId *big.Int, craftingCardTokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Identity.Contract.CraftCollection(&_Identity.TransactOpts, _dayMonthContract, _yearContract, _categoryContract, _craftingCardContract, dayMonthTokenId, yearTokenId, categoryTokenId, craftingCardTokenId, _tokenURI)
}

// CraftCollection is a paid mutator transaction binding the contract method 0xbf0837fe.
//
// Solidity: function craftCollection(address _dayMonthContract, address _yearContract, address _categoryContract, address _craftingCardContract, uint256 dayMonthTokenId, uint256 yearTokenId, uint256 categoryTokenId, uint256 craftingCardTokenId, string _tokenURI) returns()
func (_Identity *IdentityTransactorSession) CraftCollection(_dayMonthContract common.Address, _yearContract common.Address, _categoryContract common.Address, _craftingCardContract common.Address, dayMonthTokenId *big.Int, yearTokenId *big.Int, categoryTokenId *big.Int, craftingCardTokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Identity.Contract.CraftCollection(&_Identity.TransactOpts, _dayMonthContract, _yearContract, _categoryContract, _craftingCardContract, dayMonthTokenId, yearTokenId, categoryTokenId, craftingCardTokenId, _tokenURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Identity *IdentityTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Identity *IdentitySession) RenounceOwnership() (*types.Transaction, error) {
	return _Identity.Contract.RenounceOwnership(&_Identity.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Identity *IdentityTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Identity.Contract.RenounceOwnership(&_Identity.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Identity *IdentityTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Identity *IdentitySession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.SafeTransferFrom(&_Identity.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Identity *IdentityTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.SafeTransferFrom(&_Identity.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Identity *IdentityTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Identity *IdentitySession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Identity.Contract.SafeTransferFrom0(&_Identity.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Identity *IdentityTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Identity.Contract.SafeTransferFrom0(&_Identity.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Identity *IdentityTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Identity *IdentitySession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Identity.Contract.SetApprovalForAll(&_Identity.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Identity *IdentityTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Identity.Contract.SetApprovalForAll(&_Identity.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Identity *IdentityTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Identity *IdentitySession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.TransferFrom(&_Identity.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Identity *IdentityTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.TransferFrom(&_Identity.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Identity *IdentityTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Identity *IdentitySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Identity.Contract.TransferOwnership(&_Identity.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Identity *IdentityTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Identity.Contract.TransferOwnership(&_Identity.TransactOpts, newOwner)
}

// IdentityApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Identity contract.
type IdentityApprovalIterator struct {
	Event *IdentityApproval // Event containing the contract specifics and raw log

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
func (it *IdentityApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityApproval)
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
		it.Event = new(IdentityApproval)
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
func (it *IdentityApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityApproval represents a Approval event raised by the Identity contract.
type IdentityApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Identity *IdentityFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IdentityApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Identity.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IdentityApprovalIterator{contract: _Identity.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Identity *IdentityFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IdentityApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Identity.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityApproval)
				if err := _Identity.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Identity *IdentityFilterer) ParseApproval(log types.Log) (*IdentityApproval, error) {
	event := new(IdentityApproval)
	if err := _Identity.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Identity contract.
type IdentityApprovalForAllIterator struct {
	Event *IdentityApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IdentityApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityApprovalForAll)
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
		it.Event = new(IdentityApprovalForAll)
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
func (it *IdentityApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityApprovalForAll represents a ApprovalForAll event raised by the Identity contract.
type IdentityApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Identity *IdentityFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IdentityApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Identity.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IdentityApprovalForAllIterator{contract: _Identity.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Identity *IdentityFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IdentityApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Identity.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityApprovalForAll)
				if err := _Identity.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Identity *IdentityFilterer) ParseApprovalForAll(log types.Log) (*IdentityApprovalForAll, error) {
	event := new(IdentityApprovalForAll)
	if err := _Identity.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityCollectionCraftedIterator is returned from FilterCollectionCrafted and is used to iterate over the raw logs and unpacked data for CollectionCrafted events raised by the Identity contract.
type IdentityCollectionCraftedIterator struct {
	Event *IdentityCollectionCrafted // Event containing the contract specifics and raw log

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
func (it *IdentityCollectionCraftedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityCollectionCrafted)
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
		it.Event = new(IdentityCollectionCrafted)
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
func (it *IdentityCollectionCraftedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityCollectionCraftedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityCollectionCrafted represents a CollectionCrafted event raised by the Identity contract.
type IdentityCollectionCrafted struct {
	CollectionId        *big.Int
	DayMonthTokenId     *big.Int
	YearTokenId         *big.Int
	CategoryTokenId     *big.Int
	CraftingCardTokenId *big.Int
	TokenURI            string
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCollectionCrafted is a free log retrieval operation binding the contract event 0xe5979afd3a6266ac0cf01741228fe2b459458306a8d7a2c3152da07ee0c2f2be.
//
// Solidity: event CollectionCrafted(uint256 indexed collectionId, uint256 dayMonthTokenId, uint256 yearTokenId, uint256 categoryTokenId, uint256 craftingCardTokenId, string tokenURI)
func (_Identity *IdentityFilterer) FilterCollectionCrafted(opts *bind.FilterOpts, collectionId []*big.Int) (*IdentityCollectionCraftedIterator, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}

	logs, sub, err := _Identity.contract.FilterLogs(opts, "CollectionCrafted", collectionIdRule)
	if err != nil {
		return nil, err
	}
	return &IdentityCollectionCraftedIterator{contract: _Identity.contract, event: "CollectionCrafted", logs: logs, sub: sub}, nil
}

// WatchCollectionCrafted is a free log subscription operation binding the contract event 0xe5979afd3a6266ac0cf01741228fe2b459458306a8d7a2c3152da07ee0c2f2be.
//
// Solidity: event CollectionCrafted(uint256 indexed collectionId, uint256 dayMonthTokenId, uint256 yearTokenId, uint256 categoryTokenId, uint256 craftingCardTokenId, string tokenURI)
func (_Identity *IdentityFilterer) WatchCollectionCrafted(opts *bind.WatchOpts, sink chan<- *IdentityCollectionCrafted, collectionId []*big.Int) (event.Subscription, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}

	logs, sub, err := _Identity.contract.WatchLogs(opts, "CollectionCrafted", collectionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityCollectionCrafted)
				if err := _Identity.contract.UnpackLog(event, "CollectionCrafted", log); err != nil {
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

// ParseCollectionCrafted is a log parse operation binding the contract event 0xe5979afd3a6266ac0cf01741228fe2b459458306a8d7a2c3152da07ee0c2f2be.
//
// Solidity: event CollectionCrafted(uint256 indexed collectionId, uint256 dayMonthTokenId, uint256 yearTokenId, uint256 categoryTokenId, uint256 craftingCardTokenId, string tokenURI)
func (_Identity *IdentityFilterer) ParseCollectionCrafted(log types.Log) (*IdentityCollectionCrafted, error) {
	event := new(IdentityCollectionCrafted)
	if err := _Identity.contract.UnpackLog(event, "CollectionCrafted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityCraftingContractChangedIterator is returned from FilterCraftingContractChanged and is used to iterate over the raw logs and unpacked data for CraftingContractChanged events raised by the Identity contract.
type IdentityCraftingContractChangedIterator struct {
	Event *IdentityCraftingContractChanged // Event containing the contract specifics and raw log

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
func (it *IdentityCraftingContractChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityCraftingContractChanged)
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
		it.Event = new(IdentityCraftingContractChanged)
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
func (it *IdentityCraftingContractChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityCraftingContractChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityCraftingContractChanged represents a CraftingContractChanged event raised by the Identity contract.
type IdentityCraftingContractChanged struct {
	PreviousCrafter common.Address
	NewCrafter      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCraftingContractChanged is a free log retrieval operation binding the contract event 0x83149e90b909cde0e8099cb6998203f8d8dce5fa29151688382650f5637d7d88.
//
// Solidity: event CraftingContractChanged(address indexed previousCrafter, address indexed newCrafter)
func (_Identity *IdentityFilterer) FilterCraftingContractChanged(opts *bind.FilterOpts, previousCrafter []common.Address, newCrafter []common.Address) (*IdentityCraftingContractChangedIterator, error) {

	var previousCrafterRule []interface{}
	for _, previousCrafterItem := range previousCrafter {
		previousCrafterRule = append(previousCrafterRule, previousCrafterItem)
	}
	var newCrafterRule []interface{}
	for _, newCrafterItem := range newCrafter {
		newCrafterRule = append(newCrafterRule, newCrafterItem)
	}

	logs, sub, err := _Identity.contract.FilterLogs(opts, "CraftingContractChanged", previousCrafterRule, newCrafterRule)
	if err != nil {
		return nil, err
	}
	return &IdentityCraftingContractChangedIterator{contract: _Identity.contract, event: "CraftingContractChanged", logs: logs, sub: sub}, nil
}

// WatchCraftingContractChanged is a free log subscription operation binding the contract event 0x83149e90b909cde0e8099cb6998203f8d8dce5fa29151688382650f5637d7d88.
//
// Solidity: event CraftingContractChanged(address indexed previousCrafter, address indexed newCrafter)
func (_Identity *IdentityFilterer) WatchCraftingContractChanged(opts *bind.WatchOpts, sink chan<- *IdentityCraftingContractChanged, previousCrafter []common.Address, newCrafter []common.Address) (event.Subscription, error) {

	var previousCrafterRule []interface{}
	for _, previousCrafterItem := range previousCrafter {
		previousCrafterRule = append(previousCrafterRule, previousCrafterItem)
	}
	var newCrafterRule []interface{}
	for _, newCrafterItem := range newCrafter {
		newCrafterRule = append(newCrafterRule, newCrafterItem)
	}

	logs, sub, err := _Identity.contract.WatchLogs(opts, "CraftingContractChanged", previousCrafterRule, newCrafterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityCraftingContractChanged)
				if err := _Identity.contract.UnpackLog(event, "CraftingContractChanged", log); err != nil {
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

// ParseCraftingContractChanged is a log parse operation binding the contract event 0x83149e90b909cde0e8099cb6998203f8d8dce5fa29151688382650f5637d7d88.
//
// Solidity: event CraftingContractChanged(address indexed previousCrafter, address indexed newCrafter)
func (_Identity *IdentityFilterer) ParseCraftingContractChanged(log types.Log) (*IdentityCraftingContractChanged, error) {
	event := new(IdentityCraftingContractChanged)
	if err := _Identity.contract.UnpackLog(event, "CraftingContractChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Identity contract.
type IdentityOwnershipTransferredIterator struct {
	Event *IdentityOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IdentityOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityOwnershipTransferred)
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
		it.Event = new(IdentityOwnershipTransferred)
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
func (it *IdentityOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityOwnershipTransferred represents a OwnershipTransferred event raised by the Identity contract.
type IdentityOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Identity *IdentityFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdentityOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Identity.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdentityOwnershipTransferredIterator{contract: _Identity.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Identity *IdentityFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IdentityOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Identity.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityOwnershipTransferred)
				if err := _Identity.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Identity *IdentityFilterer) ParseOwnershipTransferred(log types.Log) (*IdentityOwnershipTransferred, error) {
	event := new(IdentityOwnershipTransferred)
	if err := _Identity.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityPartCraftedIterator is returned from FilterPartCrafted and is used to iterate over the raw logs and unpacked data for PartCrafted events raised by the Identity contract.
type IdentityPartCraftedIterator struct {
	Event *IdentityPartCrafted // Event containing the contract specifics and raw log

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
func (it *IdentityPartCraftedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityPartCrafted)
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
		it.Event = new(IdentityPartCrafted)
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
func (it *IdentityPartCraftedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityPartCraftedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityPartCrafted represents a PartCrafted event raised by the Identity contract.
type IdentityPartCrafted struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPartCrafted is a free log retrieval operation binding the contract event 0x71094059908fc1a5d676895c3ab6ec0e8957a09dc7a09914b791a2be4dc7fb15.
//
// Solidity: event PartCrafted(uint256 indexed tokenId)
func (_Identity *IdentityFilterer) FilterPartCrafted(opts *bind.FilterOpts, tokenId []*big.Int) (*IdentityPartCraftedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Identity.contract.FilterLogs(opts, "PartCrafted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IdentityPartCraftedIterator{contract: _Identity.contract, event: "PartCrafted", logs: logs, sub: sub}, nil
}

// WatchPartCrafted is a free log subscription operation binding the contract event 0x71094059908fc1a5d676895c3ab6ec0e8957a09dc7a09914b791a2be4dc7fb15.
//
// Solidity: event PartCrafted(uint256 indexed tokenId)
func (_Identity *IdentityFilterer) WatchPartCrafted(opts *bind.WatchOpts, sink chan<- *IdentityPartCrafted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Identity.contract.WatchLogs(opts, "PartCrafted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityPartCrafted)
				if err := _Identity.contract.UnpackLog(event, "PartCrafted", log); err != nil {
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

// ParsePartCrafted is a log parse operation binding the contract event 0x71094059908fc1a5d676895c3ab6ec0e8957a09dc7a09914b791a2be4dc7fb15.
//
// Solidity: event PartCrafted(uint256 indexed tokenId)
func (_Identity *IdentityFilterer) ParsePartCrafted(log types.Log) (*IdentityPartCrafted, error) {
	event := new(IdentityPartCrafted)
	if err := _Identity.contract.UnpackLog(event, "PartCrafted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Identity contract.
type IdentityTransferIterator struct {
	Event *IdentityTransfer // Event containing the contract specifics and raw log

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
func (it *IdentityTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityTransfer)
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
		it.Event = new(IdentityTransfer)
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
func (it *IdentityTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityTransfer represents a Transfer event raised by the Identity contract.
type IdentityTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Identity *IdentityFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IdentityTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Identity.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IdentityTransferIterator{contract: _Identity.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Identity *IdentityFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IdentityTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Identity.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityTransfer)
				if err := _Identity.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Identity *IdentityFilterer) ParseTransfer(log types.Log) (*IdentityTransfer, error) {
	event := new(IdentityTransfer)
	if err := _Identity.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
