// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ISelfAuthMetaData contains all meta data concerning the ISelfAuth contract.
var ISelfAuthMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"authRequest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"authResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ISelfAuthABI is the input ABI used to generate the binding from.
// Deprecated: Use ISelfAuthMetaData.ABI instead.
var ISelfAuthABI = ISelfAuthMetaData.ABI

// ISelfAuth is an auto generated Go binding around an Ethereum contract.
type ISelfAuth struct {
	ISelfAuthCaller     // Read-only binding to the contract
	ISelfAuthTransactor // Write-only binding to the contract
	ISelfAuthFilterer   // Log filterer for contract events
}

// ISelfAuthCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISelfAuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISelfAuthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISelfAuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISelfAuthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISelfAuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISelfAuthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISelfAuthSession struct {
	Contract     *ISelfAuth        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISelfAuthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISelfAuthCallerSession struct {
	Contract *ISelfAuthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ISelfAuthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISelfAuthTransactorSession struct {
	Contract     *ISelfAuthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ISelfAuthRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISelfAuthRaw struct {
	Contract *ISelfAuth // Generic contract binding to access the raw methods on
}

// ISelfAuthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISelfAuthCallerRaw struct {
	Contract *ISelfAuthCaller // Generic read-only contract binding to access the raw methods on
}

// ISelfAuthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISelfAuthTransactorRaw struct {
	Contract *ISelfAuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISelfAuth creates a new instance of ISelfAuth, bound to a specific deployed contract.
func NewISelfAuth(address common.Address, backend bind.ContractBackend) (*ISelfAuth, error) {
	contract, err := bindISelfAuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISelfAuth{ISelfAuthCaller: ISelfAuthCaller{contract: contract}, ISelfAuthTransactor: ISelfAuthTransactor{contract: contract}, ISelfAuthFilterer: ISelfAuthFilterer{contract: contract}}, nil
}

// NewISelfAuthCaller creates a new read-only instance of ISelfAuth, bound to a specific deployed contract.
func NewISelfAuthCaller(address common.Address, caller bind.ContractCaller) (*ISelfAuthCaller, error) {
	contract, err := bindISelfAuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISelfAuthCaller{contract: contract}, nil
}

// NewISelfAuthTransactor creates a new write-only instance of ISelfAuth, bound to a specific deployed contract.
func NewISelfAuthTransactor(address common.Address, transactor bind.ContractTransactor) (*ISelfAuthTransactor, error) {
	contract, err := bindISelfAuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISelfAuthTransactor{contract: contract}, nil
}

// NewISelfAuthFilterer creates a new log filterer instance of ISelfAuth, bound to a specific deployed contract.
func NewISelfAuthFilterer(address common.Address, filterer bind.ContractFilterer) (*ISelfAuthFilterer, error) {
	contract, err := bindISelfAuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISelfAuthFilterer{contract: contract}, nil
}

// bindISelfAuth binds a generic wrapper to an already deployed contract.
func bindISelfAuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISelfAuthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISelfAuth *ISelfAuthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISelfAuth.Contract.ISelfAuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISelfAuth *ISelfAuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISelfAuth.Contract.ISelfAuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISelfAuth *ISelfAuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISelfAuth.Contract.ISelfAuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISelfAuth *ISelfAuthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISelfAuth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISelfAuth *ISelfAuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISelfAuth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISelfAuth *ISelfAuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISelfAuth.Contract.contract.Transact(opts, method, params...)
}

// AuthResponse is a paid mutator transaction binding the contract method 0xd5ed0b90.
//
// Solidity: function authResponse(address authAddr, bytes params) returns()
func (_ISelfAuth *ISelfAuthTransactor) AuthResponse(opts *bind.TransactOpts, authAddr common.Address, params []byte) (*types.Transaction, error) {
	return _ISelfAuth.contract.Transact(opts, "authResponse", authAddr, params)
}

// AuthResponse is a paid mutator transaction binding the contract method 0xd5ed0b90.
//
// Solidity: function authResponse(address authAddr, bytes params) returns()
func (_ISelfAuth *ISelfAuthSession) AuthResponse(authAddr common.Address, params []byte) (*types.Transaction, error) {
	return _ISelfAuth.Contract.AuthResponse(&_ISelfAuth.TransactOpts, authAddr, params)
}

// AuthResponse is a paid mutator transaction binding the contract method 0xd5ed0b90.
//
// Solidity: function authResponse(address authAddr, bytes params) returns()
func (_ISelfAuth *ISelfAuthTransactorSession) AuthResponse(authAddr common.Address, params []byte) (*types.Transaction, error) {
	return _ISelfAuth.Contract.AuthResponse(&_ISelfAuth.TransactOpts, authAddr, params)
}

// ISelfAuthAuthRequestIterator is returned from FilterAuthRequest and is used to iterate over the raw logs and unpacked data for AuthRequest events raised by the ISelfAuth contract.
type ISelfAuthAuthRequestIterator struct {
	Event *ISelfAuthAuthRequest // Event containing the contract specifics and raw log

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
func (it *ISelfAuthAuthRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISelfAuthAuthRequest)
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
		it.Event = new(ISelfAuthAuthRequest)
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
func (it *ISelfAuthAuthRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISelfAuthAuthRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISelfAuthAuthRequest represents a AuthRequest event raised by the ISelfAuth contract.
type ISelfAuthAuthRequest struct {
	AuthAddr common.Address
	Params   []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAuthRequest is a free log retrieval operation binding the contract event 0xb9d8fb4f681a60258cd8900c5d9a36b449e6e12a30d844a13311d560e85a0b4e.
//
// Solidity: event authRequest(address authAddr, bytes params)
func (_ISelfAuth *ISelfAuthFilterer) FilterAuthRequest(opts *bind.FilterOpts) (*ISelfAuthAuthRequestIterator, error) {

	logs, sub, err := _ISelfAuth.contract.FilterLogs(opts, "authRequest")
	if err != nil {
		return nil, err
	}
	return &ISelfAuthAuthRequestIterator{contract: _ISelfAuth.contract, event: "authRequest", logs: logs, sub: sub}, nil
}

// WatchAuthRequest is a free log subscription operation binding the contract event 0xb9d8fb4f681a60258cd8900c5d9a36b449e6e12a30d844a13311d560e85a0b4e.
//
// Solidity: event authRequest(address authAddr, bytes params)
func (_ISelfAuth *ISelfAuthFilterer) WatchAuthRequest(opts *bind.WatchOpts, sink chan<- *ISelfAuthAuthRequest) (event.Subscription, error) {

	logs, sub, err := _ISelfAuth.contract.WatchLogs(opts, "authRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISelfAuthAuthRequest)
				if err := _ISelfAuth.contract.UnpackLog(event, "authRequest", log); err != nil {
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
func (_ISelfAuth *ISelfAuthFilterer) ParseAuthRequest(log types.Log) (*ISelfAuthAuthRequest, error) {
	event := new(ISelfAuthAuthRequest)
	if err := _ISelfAuth.contract.UnpackLog(event, "authRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
