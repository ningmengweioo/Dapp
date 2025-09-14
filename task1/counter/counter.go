// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package counter

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

// CounterMetaData contains all meta data concerning the Counter contract.
var CounterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"CountChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decrement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b505f5f819055506103a7806100225f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c806306661abd146100595780632baeceb714610077578063a87d942c14610095578063d09de08a146100b3578063d826f88f146100d1575b5f5ffd5b6100616100ef565b60405161006e919061024d565b60405180910390f35b61007f6100f4565b60405161008c919061024d565b60405180910390f35b61009d61018f565b6040516100aa919061024d565b60405180910390f35b6100bb610197565b6040516100c8919061024d565b60405180910390f35b6100d96101ef565b6040516100e6919061024d565b60405180910390f35b5f5481565b5f5f5f5411610138576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161012f906102c0565b60405180910390fd5b60015f5f828254610149919061030b565b925050819055507f0ef4482aceb854636f33f9cd319f9e1cd6fe3aa2e60523f3583c287b893824455f54604051610180919061024d565b60405180910390a15f54905090565b5f5f54905090565b5f60015f5f8282546101a9919061033e565b925050819055507f0ef4482aceb854636f33f9cd319f9e1cd6fe3aa2e60523f3583c287b893824455f546040516101e0919061024d565b60405180910390a15f54905090565b5f5f5f819055507f0ef4482aceb854636f33f9cd319f9e1cd6fe3aa2e60523f3583c287b893824455f54604051610226919061024d565b60405180910390a15f54905090565b5f819050919050565b61024781610235565b82525050565b5f6020820190506102605f83018461023e565b92915050565b5f82825260208201905092915050565b7f436f756e742063616e6e6f74206265206e6567617469766500000000000000005f82015250565b5f6102aa601883610266565b91506102b582610276565b602082019050919050565b5f6020820190508181035f8301526102d78161029e565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61031582610235565b915061032083610235565b9250828203905081811115610338576103376102de565b5b92915050565b5f61034882610235565b915061035383610235565b925082820190508082111561036b5761036a6102de565b5b9291505056fea2646970667358221220f20db51d0bba01bf22e3f7403a22d7fe223cbb8c9f395d246f02bd479d5b294164736f6c634300081e0033",
}

// CounterABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterMetaData.ABI instead.
var CounterABI = CounterMetaData.ABI

// CounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterMetaData.Bin instead.
var CounterBin = CounterMetaData.Bin

// DeployCounter deploys a new Ethereum contract, binding an instance of Counter to it.
func DeployCounter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counter, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// Counter is an auto generated Go binding around an Ethereum contract.
type Counter struct {
	CounterCaller     // Read-only binding to the contract
	CounterTransactor // Write-only binding to the contract
	CounterFilterer   // Log filterer for contract events
}

// CounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterSession struct {
	Contract     *Counter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterCallerSession struct {
	Contract *CounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterTransactorSession struct {
	Contract     *CounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterRaw struct {
	Contract *Counter // Generic contract binding to access the raw methods on
}

// CounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterCallerRaw struct {
	Contract *CounterCaller // Generic read-only contract binding to access the raw methods on
}

// CounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterTransactorRaw struct {
	Contract *CounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounter creates a new instance of Counter, bound to a specific deployed contract.
func NewCounter(address common.Address, backend bind.ContractBackend) (*Counter, error) {
	contract, err := bindCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// NewCounterCaller creates a new read-only instance of Counter, bound to a specific deployed contract.
func NewCounterCaller(address common.Address, caller bind.ContractCaller) (*CounterCaller, error) {
	contract, err := bindCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterCaller{contract: contract}, nil
}

// NewCounterTransactor creates a new write-only instance of Counter, bound to a specific deployed contract.
func NewCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterTransactor, error) {
	contract, err := bindCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterTransactor{contract: contract}, nil
}

// NewCounterFilterer creates a new log filterer instance of Counter, bound to a specific deployed contract.
func NewCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterFilterer, error) {
	contract, err := bindCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterFilterer{contract: contract}, nil
}

// bindCounter binds a generic wrapper to an already deployed contract.
func bindCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.CounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Counter *CounterCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Counter *CounterSession) Count() (*big.Int, error) {
	return _Counter.Contract.Count(&_Counter.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Counter *CounterCallerSession) Count() (*big.Int, error) {
	return _Counter.Contract.Count(&_Counter.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Counter *CounterCaller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Counter *CounterSession) GetCount() (*big.Int, error) {
	return _Counter.Contract.GetCount(&_Counter.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Counter *CounterCallerSession) GetCount() (*big.Int, error) {
	return _Counter.Contract.GetCount(&_Counter.CallOpts)
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns(uint256)
func (_Counter *CounterTransactor) Decrement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "decrement")
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns(uint256)
func (_Counter *CounterSession) Decrement() (*types.Transaction, error) {
	return _Counter.Contract.Decrement(&_Counter.TransactOpts)
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns(uint256)
func (_Counter *CounterTransactorSession) Decrement() (*types.Transaction, error) {
	return _Counter.Contract.Decrement(&_Counter.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns(uint256)
func (_Counter *CounterTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns(uint256)
func (_Counter *CounterSession) Increment() (*types.Transaction, error) {
	return _Counter.Contract.Increment(&_Counter.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns(uint256)
func (_Counter *CounterTransactorSession) Increment() (*types.Transaction, error) {
	return _Counter.Contract.Increment(&_Counter.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns(uint256)
func (_Counter *CounterTransactor) Reset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "reset")
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns(uint256)
func (_Counter *CounterSession) Reset() (*types.Transaction, error) {
	return _Counter.Contract.Reset(&_Counter.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns(uint256)
func (_Counter *CounterTransactorSession) Reset() (*types.Transaction, error) {
	return _Counter.Contract.Reset(&_Counter.TransactOpts)
}

// CounterCountChangedIterator is returned from FilterCountChanged and is used to iterate over the raw logs and unpacked data for CountChanged events raised by the Counter contract.
type CounterCountChangedIterator struct {
	Event *CounterCountChanged // Event containing the contract specifics and raw log

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
func (it *CounterCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterCountChanged)
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
		it.Event = new(CounterCountChanged)
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
func (it *CounterCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterCountChanged represents a CountChanged event raised by the Counter contract.
type CounterCountChanged struct {
	NewCount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCountChanged is a free log retrieval operation binding the contract event 0x0ef4482aceb854636f33f9cd319f9e1cd6fe3aa2e60523f3583c287b89382445.
//
// Solidity: event CountChanged(uint256 newCount)
func (_Counter *CounterFilterer) FilterCountChanged(opts *bind.FilterOpts) (*CounterCountChangedIterator, error) {

	logs, sub, err := _Counter.contract.FilterLogs(opts, "CountChanged")
	if err != nil {
		return nil, err
	}
	return &CounterCountChangedIterator{contract: _Counter.contract, event: "CountChanged", logs: logs, sub: sub}, nil
}

// WatchCountChanged is a free log subscription operation binding the contract event 0x0ef4482aceb854636f33f9cd319f9e1cd6fe3aa2e60523f3583c287b89382445.
//
// Solidity: event CountChanged(uint256 newCount)
func (_Counter *CounterFilterer) WatchCountChanged(opts *bind.WatchOpts, sink chan<- *CounterCountChanged) (event.Subscription, error) {

	logs, sub, err := _Counter.contract.WatchLogs(opts, "CountChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterCountChanged)
				if err := _Counter.contract.UnpackLog(event, "CountChanged", log); err != nil {
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

// ParseCountChanged is a log parse operation binding the contract event 0x0ef4482aceb854636f33f9cd319f9e1cd6fe3aa2e60523f3583c287b89382445.
//
// Solidity: event CountChanged(uint256 newCount)
func (_Counter *CounterFilterer) ParseCountChanged(log types.Log) (*CounterCountChanged, error) {
	event := new(CounterCountChanged)
	if err := _Counter.contract.UnpackLog(event, "CountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
