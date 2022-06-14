// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rscore

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

// SelfAuthMetaData contains all meta data concerning the SelfAuth contract.
var SelfAuthMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"authRequest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"authResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SelfAuthABI is the input ABI used to generate the binding from.
// Deprecated: Use SelfAuthMetaData.ABI instead.
var SelfAuthABI = SelfAuthMetaData.ABI

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
