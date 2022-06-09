// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts_demo

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
	Sigs: map[string]string{
		"d5ed0b90": "authResponse(address,bytes)",
	},
}

// ISelfAuthABI is the input ABI used to generate the binding from.
// Deprecated: Use ISelfAuthMetaData.ABI instead.
var ISelfAuthABI = ISelfAuthMetaData.ABI

// Deprecated: Use ISelfAuthMetaData.Sigs instead.
// ISelfAuthFuncSigs maps the 4-byte function signature to its string representation.
var ISelfAuthFuncSigs = ISelfAuthMetaData.Sigs

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

// SADemoMetaData contains all meta data concerning the SADemo contract.
var SADemoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"authRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SELFAUTH_OFFICIAL_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"authResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authAddr\",\"type\":\"address\"}],\"name\":\"trigger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authAddr\",\"type\":\"address\"}],\"name\":\"verified\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"4c6deef4": "SELFAUTH_OFFICIAL_ADDRESS()",
		"d5ed0b90": "authResponse(address,bytes)",
		"8da5cb5b": "owner()",
		"6f6ff405": "trigger(address)",
		"0db065f4": "verified(address)",
	},
	Bin: "0x6080604052600280546001600160a01b0319167361e1ceb3e8bd42d2be9ad0d4fd58fc0e10c7bb5e17905534801561003657600080fd5b50600080546001600160a01b0319163317905561045c806100586000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630db065f41461005c5780634c6deef4146100985780636f6ff405146100c35780638da5cb5b146100d8578063d5ed0b90146100eb575b600080fd5b61008561006a366004610288565b6001600160a01b031660009081526001602052604090205490565b6040519081526020015b60405180910390f35b6002546100ab906001600160a01b031681565b6040516001600160a01b03909116815260200161008f565b6100d66100d1366004610288565b6100fe565b005b6000546100ab906001600160a01b031681565b6100d66100f93660046102c0565b61016c565b6001600160a01b038116600090815260016020908152604091829020548251918201527fb9d8fb4f681a60258cd8900c5d9a36b449e6e12a30d844a13311d560e85a0b4e9183910160408051601f19818403018152908290526101619291610382565b60405180910390a150565b6002546001600160a01b031633146101bf5760405162461bcd60e51b81526020600482015260116024820152701c195c9b5a5cdcda5bdb8819195b9a5959607a1b60448201526064015b60405180910390fd5b6000818060200190518101906101d591906103e7565b90506101e183826101e6565b505050565b6001600160a01b03821660009081526001602052604090205481146102415760405162461bcd60e51b815260206004820152601160248201527034b73b30b634b21037b832b930ba34b7b760791b60448201526064016101b6565b61024c816001610400565b6001600160a01b0390921660009081526001602052604090209190915550565b80356001600160a01b038116811461028357600080fd5b919050565b60006020828403121561029a57600080fd5b6102a38261026c565b9392505050565b634e487b7160e01b600052604160045260246000fd5b600080604083850312156102d357600080fd5b6102dc8361026c565b9150602083013567ffffffffffffffff808211156102f957600080fd5b818501915085601f83011261030d57600080fd5b81358181111561031f5761031f6102aa565b604051601f8201601f19908116603f01168101908382118183101715610347576103476102aa565b8160405282815288602084870101111561036057600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b60018060a01b038316815260006020604081840152835180604085015260005b818110156103be578581018301518582016060015282016103a2565b818111156103d0576000606083870101525b50601f01601f191692909201606001949350505050565b6000602082840312156103f957600080fd5b5051919050565b6000821982111561042157634e487b7160e01b600052601160045260246000fd5b50019056fea2646970667358221220d2852624d79a122ac26187aae50738637913ed938c5f55ca40d8d93c1d806b7364736f6c634300080e0033",
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

// SELFAUTHOFFICIALADDRESS is a free data retrieval call binding the contract method 0x4c6deef4.
//
// Solidity: function SELFAUTH_OFFICIAL_ADDRESS() view returns(address)
func (_SADemo *SADemoCaller) SELFAUTHOFFICIALADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SADemo.contract.Call(opts, &out, "SELFAUTH_OFFICIAL_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SELFAUTHOFFICIALADDRESS is a free data retrieval call binding the contract method 0x4c6deef4.
//
// Solidity: function SELFAUTH_OFFICIAL_ADDRESS() view returns(address)
func (_SADemo *SADemoSession) SELFAUTHOFFICIALADDRESS() (common.Address, error) {
	return _SADemo.Contract.SELFAUTHOFFICIALADDRESS(&_SADemo.CallOpts)
}

// SELFAUTHOFFICIALADDRESS is a free data retrieval call binding the contract method 0x4c6deef4.
//
// Solidity: function SELFAUTH_OFFICIAL_ADDRESS() view returns(address)
func (_SADemo *SADemoCallerSession) SELFAUTHOFFICIALADDRESS() (common.Address, error) {
	return _SADemo.Contract.SELFAUTHOFFICIALADDRESS(&_SADemo.CallOpts)
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
