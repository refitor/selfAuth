// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package demo

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
)

// ContextMetaData contains all meta data concerning the Context contract.
var ContextMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextMetaData.ABI instead.
var ContextABI = ContextMetaData.ABI

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Deprecated: Use OwnableMetaData.Sigs instead.
// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = OwnableMetaData.Sigs

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SADemoMetaData contains all meta data concerning the SADemo contract.
var SADemoMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"authRequest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"authResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authAddr\",\"type\":\"address\"}],\"name\":\"trigger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authAddr\",\"type\":\"address\"}],\"name\":\"verified\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d5ed0b90": "authResponse(address,bytes)",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
		"6f6ff405": "trigger(address)",
		"0db065f4": "verified(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6106138061007e6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630db065f4146100675780636f6ff4051461008d578063715018a6146100a25780638da5cb5b146100aa578063d5ed0b90146100c5578063f2fde38b146100d8575b600080fd5b61007a61007536600461040a565b6100eb565b6040519081526020015b60405180910390f35b6100a061009b36600461040a565b61013f565b005b6100a06101d7565b6000546040516001600160a01b039091168152602001610084565b6100a06100d3366004610442565b61020d565b6100a06100e636600461040a565b610282565b600080546001600160a01b0316331461011f5760405162461bcd60e51b815260040161011690610504565b60405180910390fd5b506001600160a01b0381166000908152600160205260409020545b919050565b6000546001600160a01b031633146101695760405162461bcd60e51b815260040161011690610504565b6001600160a01b038116600090815260016020908152604091829020548251918201527fb9d8fb4f681a60258cd8900c5d9a36b449e6e12a30d844a13311d560e85a0b4e9183910160408051601f19818403018152908290526101cc9291610539565b60405180910390a150565b6000546001600160a01b031633146102015760405162461bcd60e51b815260040161011690610504565b61020b600061031d565b565b6000546001600160a01b0316331461025b5760405162461bcd60e51b81526020600482015260116024820152701c195c9b5a5cdcda5bdb8819195b9a5959607a1b6044820152606401610116565b600081806020019051810190610271919061059e565b905061027d838261036d565b505050565b6000546001600160a01b031633146102ac5760405162461bcd60e51b815260040161011690610504565b6001600160a01b0381166103115760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610116565b61031a8161031d565b50565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03821660009081526001602052604090205481146103c85760405162461bcd60e51b815260206004820152601160248201527034b73b30b634b21037b832b930ba34b7b760791b6044820152606401610116565b6103d38160016105b7565b6001600160a01b0390921660009081526001602052604090209190915550565b80356001600160a01b038116811461013a57600080fd5b60006020828403121561041c57600080fd5b610425826103f3565b9392505050565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561045557600080fd5b61045e836103f3565b9150602083013567ffffffffffffffff8082111561047b57600080fd5b818501915085601f83011261048f57600080fd5b8135818111156104a1576104a161042c565b604051601f8201601f19908116603f011681019083821181831017156104c9576104c961042c565b816040528281528860208487010111156104e257600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60018060a01b038316815260006020604081840152835180604085015260005b8181101561057557858101830151858201606001528201610559565b81811115610587576000606083870101525b50601f01601f191692909201606001949350505050565b6000602082840312156105b057600080fd5b5051919050565b600082198211156105d857634e487b7160e01b600052601160045260246000fd5b50019056fea264697066735822122083a20d086ba58c6e96d0b9a12e2fc0069f7b3433907226799f527dc6d0fbcc3164736f6c634300080e0033",
}

// SADemoABI is the input ABI used to generate the binding from.
// Deprecated: Use SADemoMetaData.ABI instead.
var SADemoABI = SADemoMetaData.ABI

// Deprecated: Use SADemoMetaData.Sigs instead.
// SADemoFuncSigs maps the 4-byte function signature to its string representation.
var SADemoFuncSigs = SADemoMetaData.Sigs

// SADemoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SADemoMetaData.Bin instead.
var SADemoBin = SADemoMetaData.Bin

// DeploySADemo deploys a new Ethereum contract, binding an instance of SADemo to it.
func DeploySADemo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SADemo, error) {
	parsed, err := SADemoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SADemoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SADemo{SADemoCaller: SADemoCaller{contract: contract}, SADemoTransactor: SADemoTransactor{contract: contract}, SADemoFilterer: SADemoFilterer{contract: contract}}, nil
}

// SADemo is an auto generated Go binding around an Ethereum contract.
type SADemo struct {
	SADemoCaller     // Read-only binding to the contract
	SADemoTransactor // Write-only binding to the contract
	SADemoFilterer   // Log filterer for contract events
}

// SADemoCaller is an auto generated read-only Go binding around an Ethereum contract.
type SADemoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SADemoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SADemoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SADemoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SADemoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SADemoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SADemoSession struct {
	Contract     *SADemo           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SADemoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SADemoCallerSession struct {
	Contract *SADemoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SADemoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SADemoTransactorSession struct {
	Contract     *SADemoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SADemoRaw is an auto generated low-level Go binding around an Ethereum contract.
type SADemoRaw struct {
	Contract *SADemo // Generic contract binding to access the raw methods on
}

// SADemoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SADemoCallerRaw struct {
	Contract *SADemoCaller // Generic read-only contract binding to access the raw methods on
}

// SADemoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SADemoTransactorRaw struct {
	Contract *SADemoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSADemo creates a new instance of SADemo, bound to a specific deployed contract.
func NewSADemo(address common.Address, backend bind.ContractBackend) (*SADemo, error) {
	contract, err := bindSADemo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SADemo{SADemoCaller: SADemoCaller{contract: contract}, SADemoTransactor: SADemoTransactor{contract: contract}, SADemoFilterer: SADemoFilterer{contract: contract}}, nil
}

// NewSADemoCaller creates a new read-only instance of SADemo, bound to a specific deployed contract.
func NewSADemoCaller(address common.Address, caller bind.ContractCaller) (*SADemoCaller, error) {
	contract, err := bindSADemo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SADemoCaller{contract: contract}, nil
}

// NewSADemoTransactor creates a new write-only instance of SADemo, bound to a specific deployed contract.
func NewSADemoTransactor(address common.Address, transactor bind.ContractTransactor) (*SADemoTransactor, error) {
	contract, err := bindSADemo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SADemoTransactor{contract: contract}, nil
}

// NewSADemoFilterer creates a new log filterer instance of SADemo, bound to a specific deployed contract.
func NewSADemoFilterer(address common.Address, filterer bind.ContractFilterer) (*SADemoFilterer, error) {
	contract, err := bindSADemo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SADemoFilterer{contract: contract}, nil
}

// bindSADemo binds a generic wrapper to an already deployed contract.
func bindSADemo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SADemoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SADemo *SADemoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SADemo.Contract.SADemoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SADemo *SADemoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SADemo.Contract.SADemoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SADemo *SADemoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SADemo.Contract.SADemoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SADemo *SADemoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SADemo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SADemo *SADemoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SADemo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SADemo *SADemoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SADemo.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SADemo *SADemoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SADemo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SADemo *SADemoSession) Owner() (common.Address, error) {
	return _SADemo.Contract.Owner(&_SADemo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SADemo *SADemoCallerSession) Owner() (common.Address, error) {
	return _SADemo.Contract.Owner(&_SADemo.CallOpts)
}

// Verified is a free data retrieval call binding the contract method 0x0db065f4.
//
// Solidity: function verified(address authAddr) view returns(uint256)
func (_SADemo *SADemoCaller) Verified(opts *bind.CallOpts, authAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SADemo.contract.Call(opts, &out, "verified", authAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Verified is a free data retrieval call binding the contract method 0x0db065f4.
//
// Solidity: function verified(address authAddr) view returns(uint256)
func (_SADemo *SADemoSession) Verified(authAddr common.Address) (*big.Int, error) {
	return _SADemo.Contract.Verified(&_SADemo.CallOpts, authAddr)
}

// Verified is a free data retrieval call binding the contract method 0x0db065f4.
//
// Solidity: function verified(address authAddr) view returns(uint256)
func (_SADemo *SADemoCallerSession) Verified(authAddr common.Address) (*big.Int, error) {
	return _SADemo.Contract.Verified(&_SADemo.CallOpts, authAddr)
}

// AuthResponse is a paid mutator transaction binding the contract method 0xd5ed0b90.
//
// Solidity: function authResponse(address authAddr, bytes params) returns()
func (_SADemo *SADemoTransactor) AuthResponse(opts *bind.TransactOpts, authAddr common.Address, params []byte) (*types.Transaction, error) {
	return _SADemo.contract.Transact(opts, "authResponse", authAddr, params)
}

// AuthResponse is a paid mutator transaction binding the contract method 0xd5ed0b90.
//
// Solidity: function authResponse(address authAddr, bytes params) returns()
func (_SADemo *SADemoSession) AuthResponse(authAddr common.Address, params []byte) (*types.Transaction, error) {
	return _SADemo.Contract.AuthResponse(&_SADemo.TransactOpts, authAddr, params)
}

// AuthResponse is a paid mutator transaction binding the contract method 0xd5ed0b90.
//
// Solidity: function authResponse(address authAddr, bytes params) returns()
func (_SADemo *SADemoTransactorSession) AuthResponse(authAddr common.Address, params []byte) (*types.Transaction, error) {
	return _SADemo.Contract.AuthResponse(&_SADemo.TransactOpts, authAddr, params)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SADemo *SADemoTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SADemo.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SADemo *SADemoSession) RenounceOwnership() (*types.Transaction, error) {
	return _SADemo.Contract.RenounceOwnership(&_SADemo.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SADemo *SADemoTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SADemo.Contract.RenounceOwnership(&_SADemo.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SADemo *SADemoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SADemo.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SADemo *SADemoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SADemo.Contract.TransferOwnership(&_SADemo.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SADemo *SADemoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SADemo.Contract.TransferOwnership(&_SADemo.TransactOpts, newOwner)
}

// Trigger is a paid mutator transaction binding the contract method 0x6f6ff405.
//
// Solidity: function trigger(address authAddr) returns()
func (_SADemo *SADemoTransactor) Trigger(opts *bind.TransactOpts, authAddr common.Address) (*types.Transaction, error) {
	return _SADemo.contract.Transact(opts, "trigger", authAddr)
}

// Trigger is a paid mutator transaction binding the contract method 0x6f6ff405.
//
// Solidity: function trigger(address authAddr) returns()
func (_SADemo *SADemoSession) Trigger(authAddr common.Address) (*types.Transaction, error) {
	return _SADemo.Contract.Trigger(&_SADemo.TransactOpts, authAddr)
}

// Trigger is a paid mutator transaction binding the contract method 0x6f6ff405.
//
// Solidity: function trigger(address authAddr) returns()
func (_SADemo *SADemoTransactorSession) Trigger(authAddr common.Address) (*types.Transaction, error) {
	return _SADemo.Contract.Trigger(&_SADemo.TransactOpts, authAddr)
}

// SADemoOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SADemo contract.
type SADemoOwnershipTransferredIterator struct {
	Event *SADemoOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SADemoOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SADemoOwnershipTransferred)
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
		it.Event = new(SADemoOwnershipTransferred)
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
func (it *SADemoOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SADemoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SADemoOwnershipTransferred represents a OwnershipTransferred event raised by the SADemo contract.
type SADemoOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SADemo *SADemoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SADemoOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SADemo.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SADemoOwnershipTransferredIterator{contract: _SADemo.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SADemo *SADemoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SADemoOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SADemo.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SADemoOwnershipTransferred)
				if err := _SADemo.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SADemo *SADemoFilterer) ParseOwnershipTransferred(log types.Log) (*SADemoOwnershipTransferred, error) {
	event := new(SADemoOwnershipTransferred)
	if err := _SADemo.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SADemoAuthRequestIterator is returned from FilterAuthRequest and is used to iterate over the raw logs and unpacked data for AuthRequest events raised by the SADemo contract.
type SADemoAuthRequestIterator struct {
	Event *SADemoAuthRequest // Event containing the contract specifics and raw log

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
func (it *SADemoAuthRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SADemoAuthRequest)
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
		it.Event = new(SADemoAuthRequest)
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
func (it *SADemoAuthRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SADemoAuthRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SADemoAuthRequest represents a AuthRequest event raised by the SADemo contract.
type SADemoAuthRequest struct {
	AuthAddr common.Address
	Params   []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAuthRequest is a free log retrieval operation binding the contract event 0xb9d8fb4f681a60258cd8900c5d9a36b449e6e12a30d844a13311d560e85a0b4e.
//
// Solidity: event authRequest(address authAddr, bytes params)
func (_SADemo *SADemoFilterer) FilterAuthRequest(opts *bind.FilterOpts) (*SADemoAuthRequestIterator, error) {

	logs, sub, err := _SADemo.contract.FilterLogs(opts, "authRequest")
	if err != nil {
		return nil, err
	}
	return &SADemoAuthRequestIterator{contract: _SADemo.contract, event: "authRequest", logs: logs, sub: sub}, nil
}

// WatchAuthRequest is a free log subscription operation binding the contract event 0xb9d8fb4f681a60258cd8900c5d9a36b449e6e12a30d844a13311d560e85a0b4e.
//
// Solidity: event authRequest(address authAddr, bytes params)
func (_SADemo *SADemoFilterer) WatchAuthRequest(opts *bind.WatchOpts, sink chan<- *SADemoAuthRequest) (event.Subscription, error) {

	logs, sub, err := _SADemo.contract.WatchLogs(opts, "authRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SADemoAuthRequest)
				if err := _SADemo.contract.UnpackLog(event, "authRequest", log); err != nil {
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

// ParseAuthRequest is a log parse operation binding the contract event 0xb9d8fb4f681a60258cd8900c5d9a36b449e6e12a30d844a13311d560e85a0b4e.
//
// Solidity: event authRequest(address authAddr, bytes params)
func (_SADemo *SADemoFilterer) ParseAuthRequest(log types.Log) (*SADemoAuthRequest, error) {
	event := new(SADemoAuthRequest)
	if err := _SADemo.contract.UnpackLog(event, "authRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d6ec644c29adc79940e9052cff36d66cc082133797ae0999bc444f21bcd9e4fc64736f6c634300080e0033",
}

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SelfAuthMetaData contains all meta data concerning the SelfAuth contract.
var SelfAuthMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"authRequest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"authResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d5ed0b90": "authResponse(address,bytes)",
	},
}

// SelfAuthABI is the input ABI used to generate the binding from.
// Deprecated: Use SelfAuthMetaData.ABI instead.
var SelfAuthABI = SelfAuthMetaData.ABI

// Deprecated: Use SelfAuthMetaData.Sigs instead.
// SelfAuthFuncSigs maps the 4-byte function signature to its string representation.
var SelfAuthFuncSigs = SelfAuthMetaData.Sigs

// SelfAuth is an auto generated Go binding around an Ethereum contract.
type SelfAuth struct {
	SelfAuthCaller     // Read-only binding to the contract
	SelfAuthTransactor // Write-only binding to the contract
	SelfAuthFilterer   // Log filterer for contract events
}

// SelfAuthCaller is an auto generated read-only Go binding around an Ethereum contract.
type SelfAuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelfAuthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SelfAuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelfAuthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SelfAuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelfAuthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SelfAuthSession struct {
	Contract     *SelfAuth         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SelfAuthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SelfAuthCallerSession struct {
	Contract *SelfAuthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SelfAuthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SelfAuthTransactorSession struct {
	Contract     *SelfAuthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SelfAuthRaw is an auto generated low-level Go binding around an Ethereum contract.
type SelfAuthRaw struct {
	Contract *SelfAuth // Generic contract binding to access the raw methods on
}

// SelfAuthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SelfAuthCallerRaw struct {
	Contract *SelfAuthCaller // Generic read-only contract binding to access the raw methods on
}

// SelfAuthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SelfAuthTransactorRaw struct {
	Contract *SelfAuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSelfAuth creates a new instance of SelfAuth, bound to a specific deployed contract.
func NewSelfAuth(address common.Address, backend bind.ContractBackend) (*SelfAuth, error) {
	contract, err := bindSelfAuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SelfAuth{SelfAuthCaller: SelfAuthCaller{contract: contract}, SelfAuthTransactor: SelfAuthTransactor{contract: contract}, SelfAuthFilterer: SelfAuthFilterer{contract: contract}}, nil
}

// NewSelfAuthCaller creates a new read-only instance of SelfAuth, bound to a specific deployed contract.
func NewSelfAuthCaller(address common.Address, caller bind.ContractCaller) (*SelfAuthCaller, error) {
	contract, err := bindSelfAuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SelfAuthCaller{contract: contract}, nil
}

// NewSelfAuthTransactor creates a new write-only instance of SelfAuth, bound to a specific deployed contract.
func NewSelfAuthTransactor(address common.Address, transactor bind.ContractTransactor) (*SelfAuthTransactor, error) {
	contract, err := bindSelfAuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SelfAuthTransactor{contract: contract}, nil
}

// NewSelfAuthFilterer creates a new log filterer instance of SelfAuth, bound to a specific deployed contract.
func NewSelfAuthFilterer(address common.Address, filterer bind.ContractFilterer) (*SelfAuthFilterer, error) {
	contract, err := bindSelfAuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SelfAuthFilterer{contract: contract}, nil
}

// bindSelfAuth binds a generic wrapper to an already deployed contract.
func bindSelfAuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SelfAuthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SelfAuth *SelfAuthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SelfAuth.Contract.SelfAuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SelfAuth *SelfAuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfAuth.Contract.SelfAuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SelfAuth *SelfAuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SelfAuth.Contract.SelfAuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SelfAuth *SelfAuthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SelfAuth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SelfAuth *SelfAuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfAuth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SelfAuth *SelfAuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SelfAuth.Contract.contract.Transact(opts, method, params...)
}

// AuthResponse is a paid mutator transaction binding the contract method 0xd5ed0b90.
//
// Solidity: function authResponse(address authAddr, bytes params) returns()
func (_SelfAuth *SelfAuthTransactor) AuthResponse(opts *bind.TransactOpts, authAddr common.Address, params []byte) (*types.Transaction, error) {
	return _SelfAuth.contract.Transact(opts, "authResponse", authAddr, params)
}

// AuthResponse is a paid mutator transaction binding the contract method 0xd5ed0b90.
//
// Solidity: function authResponse(address authAddr, bytes params) returns()
func (_SelfAuth *SelfAuthSession) AuthResponse(authAddr common.Address, params []byte) (*types.Transaction, error) {
	return _SelfAuth.Contract.AuthResponse(&_SelfAuth.TransactOpts, authAddr, params)
}

// AuthResponse is a paid mutator transaction binding the contract method 0xd5ed0b90.
//
// Solidity: function authResponse(address authAddr, bytes params) returns()
func (_SelfAuth *SelfAuthTransactorSession) AuthResponse(authAddr common.Address, params []byte) (*types.Transaction, error) {
	return _SelfAuth.Contract.AuthResponse(&_SelfAuth.TransactOpts, authAddr, params)
}

// SelfAuthAuthRequestIterator is returned from FilterAuthRequest and is used to iterate over the raw logs and unpacked data for AuthRequest events raised by the SelfAuth contract.
type SelfAuthAuthRequestIterator struct {
	Event *SelfAuthAuthRequest // Event containing the contract specifics and raw log

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
func (it *SelfAuthAuthRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SelfAuthAuthRequest)
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
		it.Event = new(SelfAuthAuthRequest)
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
func (it *SelfAuthAuthRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SelfAuthAuthRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SelfAuthAuthRequest represents a AuthRequest event raised by the SelfAuth contract.
type SelfAuthAuthRequest struct {
	AuthAddr common.Address
	Params   []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAuthRequest is a free log retrieval operation binding the contract event 0xb9d8fb4f681a60258cd8900c5d9a36b449e6e12a30d844a13311d560e85a0b4e.
//
// Solidity: event authRequest(address authAddr, bytes params)
func (_SelfAuth *SelfAuthFilterer) FilterAuthRequest(opts *bind.FilterOpts) (*SelfAuthAuthRequestIterator, error) {

	logs, sub, err := _SelfAuth.contract.FilterLogs(opts, "authRequest")
	if err != nil {
		return nil, err
	}
	return &SelfAuthAuthRequestIterator{contract: _SelfAuth.contract, event: "authRequest", logs: logs, sub: sub}, nil
}

// WatchAuthRequest is a free log subscription operation binding the contract event 0xb9d8fb4f681a60258cd8900c5d9a36b449e6e12a30d844a13311d560e85a0b4e.
//
// Solidity: event authRequest(address authAddr, bytes params)
func (_SelfAuth *SelfAuthFilterer) WatchAuthRequest(opts *bind.WatchOpts, sink chan<- *SelfAuthAuthRequest) (event.Subscription, error) {

	logs, sub, err := _SelfAuth.contract.WatchLogs(opts, "authRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SelfAuthAuthRequest)
				if err := _SelfAuth.contract.UnpackLog(event, "authRequest", log); err != nil {
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

// ParseAuthRequest is a log parse operation binding the contract event 0xb9d8fb4f681a60258cd8900c5d9a36b449e6e12a30d844a13311d560e85a0b4e.
//
// Solidity: event authRequest(address authAddr, bytes params)
func (_SelfAuth *SelfAuthFilterer) ParseAuthRequest(log types.Log) (*SelfAuthAuthRequest, error) {
	event := new(SelfAuthAuthRequest)
	if err := _SelfAuth.contract.UnpackLog(event, "authRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
