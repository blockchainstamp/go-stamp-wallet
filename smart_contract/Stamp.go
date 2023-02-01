// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StampRecord is an auto generated low-level Go binding around an user-defined struct.
type StampRecord struct {
	Value *big.Int
	Nonce *big.Int
}

// StampABI is the input ABI used to generate the binding from.
const StampABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mailboxName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"consummable\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"AdminOP\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Revoke\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"BalanceOf\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structStamp.Record\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"adminOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"conf\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"MailBoxName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"IsConsummable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Stamp is an auto generated Go binding around an Ethereum contract.
type Stamp struct {
	StampCaller     // Read-only binding to the contract
	StampTransactor // Write-only binding to the contract
	StampFilterer   // Log filterer for contract events
}

// StampCaller is an auto generated read-only Go binding around an Ethereum contract.
type StampCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StampTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StampFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StampSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StampSession struct {
	Contract     *Stamp            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StampCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StampCallerSession struct {
	Contract *StampCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StampTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StampTransactorSession struct {
	Contract     *StampTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StampRaw is an auto generated low-level Go binding around an Ethereum contract.
type StampRaw struct {
	Contract *Stamp // Generic contract binding to access the raw methods on
}

// StampCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StampCallerRaw struct {
	Contract *StampCaller // Generic read-only contract binding to access the raw methods on
}

// StampTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StampTransactorRaw struct {
	Contract *StampTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStamp creates a new instance of Stamp, bound to a specific deployed contract.
func NewStamp(address common.Address, backend bind.ContractBackend) (*Stamp, error) {
	contract, err := bindStamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stamp{StampCaller: StampCaller{contract: contract}, StampTransactor: StampTransactor{contract: contract}, StampFilterer: StampFilterer{contract: contract}}, nil
}

// NewStampCaller creates a new read-only instance of Stamp, bound to a specific deployed contract.
func NewStampCaller(address common.Address, caller bind.ContractCaller) (*StampCaller, error) {
	contract, err := bindStamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StampCaller{contract: contract}, nil
}

// NewStampTransactor creates a new write-only instance of Stamp, bound to a specific deployed contract.
func NewStampTransactor(address common.Address, transactor bind.ContractTransactor) (*StampTransactor, error) {
	contract, err := bindStamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StampTransactor{contract: contract}, nil
}

// NewStampFilterer creates a new log filterer instance of Stamp, bound to a specific deployed contract.
func NewStampFilterer(address common.Address, filterer bind.ContractFilterer) (*StampFilterer, error) {
	contract, err := bindStamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StampFilterer{contract: contract}, nil
}

// bindStamp binds a generic wrapper to an already deployed contract.
func bindStamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stamp *StampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stamp.Contract.StampCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stamp *StampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stamp.Contract.StampTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stamp *StampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stamp.Contract.StampTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stamp *StampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stamp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stamp *StampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stamp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stamp *StampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stamp.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xd9967889.
//
// Solidity: function BalanceOf(address user) view returns((uint256,uint256))
func (_Stamp *StampCaller) BalanceOf(opts *bind.CallOpts, user common.Address) (StampRecord, error) {
	var out []interface{}
	err := _Stamp.contract.Call(opts, &out, "BalanceOf", user)

	if err != nil {
		return *new(StampRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(StampRecord)).(*StampRecord)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xd9967889.
//
// Solidity: function BalanceOf(address user) view returns((uint256,uint256))
func (_Stamp *StampSession) BalanceOf(user common.Address) (StampRecord, error) {
	return _Stamp.Contract.BalanceOf(&_Stamp.CallOpts, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0xd9967889.
//
// Solidity: function BalanceOf(address user) view returns((uint256,uint256))
func (_Stamp *StampCallerSession) BalanceOf(user common.Address) (StampRecord, error) {
	return _Stamp.Contract.BalanceOf(&_Stamp.CallOpts, user)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_Stamp *StampCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Stamp.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_Stamp *StampSession) Admins(arg0 common.Address) (bool, error) {
	return _Stamp.Contract.Admins(&_Stamp.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_Stamp *StampCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _Stamp.Contract.Admins(&_Stamp.CallOpts, arg0)
}

// Conf is a free data retrieval call binding the contract method 0x17792729.
//
// Solidity: function conf() view returns(string MailBoxName, bool IsConsummable)
func (_Stamp *StampCaller) Conf(opts *bind.CallOpts) (struct {
	MailBoxName   string
	IsConsummable bool
}, error) {
	var out []interface{}
	err := _Stamp.contract.Call(opts, &out, "conf")

	outstruct := new(struct {
		MailBoxName   string
		IsConsummable bool
	})

	outstruct.MailBoxName = out[0].(string)
	outstruct.IsConsummable = out[1].(bool)

	return *outstruct, err

}

// Conf is a free data retrieval call binding the contract method 0x17792729.
//
// Solidity: function conf() view returns(string MailBoxName, bool IsConsummable)
func (_Stamp *StampSession) Conf() (struct {
	MailBoxName   string
	IsConsummable bool
}, error) {
	return _Stamp.Contract.Conf(&_Stamp.CallOpts)
}

// Conf is a free data retrieval call binding the contract method 0x17792729.
//
// Solidity: function conf() view returns(string MailBoxName, bool IsConsummable)
func (_Stamp *StampCallerSession) Conf() (struct {
	MailBoxName   string
	IsConsummable bool
}, error) {
	return _Stamp.Contract.Conf(&_Stamp.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Stamp *StampCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Stamp.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Stamp *StampSession) Decimals() (uint8, error) {
	return _Stamp.Contract.Decimals(&_Stamp.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Stamp *StampCallerSession) Decimals() (uint8, error) {
	return _Stamp.Contract.Decimals(&_Stamp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stamp *StampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stamp *StampSession) Owner() (common.Address, error) {
	return _Stamp.Contract.Owner(&_Stamp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stamp *StampCallerSession) Owner() (common.Address, error) {
	return _Stamp.Contract.Owner(&_Stamp.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Stamp *StampCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stamp.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Stamp *StampSession) TotalSupply() (*big.Int, error) {
	return _Stamp.Contract.TotalSupply(&_Stamp.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Stamp *StampCallerSession) TotalSupply() (*big.Int, error) {
	return _Stamp.Contract.TotalSupply(&_Stamp.CallOpts)
}

// AdminOp is a paid mutator transaction binding the contract method 0x809b4767.
//
// Solidity: function adminOp(address addr, bool isAdd) returns()
func (_Stamp *StampTransactor) AdminOp(opts *bind.TransactOpts, addr common.Address, isAdd bool) (*types.Transaction, error) {
	return _Stamp.contract.Transact(opts, "adminOp", addr, isAdd)
}

// AdminOp is a paid mutator transaction binding the contract method 0x809b4767.
//
// Solidity: function adminOp(address addr, bool isAdd) returns()
func (_Stamp *StampSession) AdminOp(addr common.Address, isAdd bool) (*types.Transaction, error) {
	return _Stamp.Contract.AdminOp(&_Stamp.TransactOpts, addr, isAdd)
}

// AdminOp is a paid mutator transaction binding the contract method 0x809b4767.
//
// Solidity: function adminOp(address addr, bool isAdd) returns()
func (_Stamp *StampTransactorSession) AdminOp(addr common.Address, isAdd bool) (*types.Transaction, error) {
	return _Stamp.Contract.AdminOp(&_Stamp.TransactOpts, addr, isAdd)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Stamp *StampTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Stamp.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Stamp *StampSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Stamp.Contract.ChangeOwner(&_Stamp.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Stamp *StampTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Stamp.Contract.ChangeOwner(&_Stamp.TransactOpts, newOwner)
}

// Revoke is a paid mutator transaction binding the contract method 0xeac449d9.
//
// Solidity: function revoke(address to, uint256 value) returns()
func (_Stamp *StampTransactor) Revoke(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Stamp.contract.Transact(opts, "revoke", to, value)
}

// Revoke is a paid mutator transaction binding the contract method 0xeac449d9.
//
// Solidity: function revoke(address to, uint256 value) returns()
func (_Stamp *StampSession) Revoke(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Stamp.Contract.Revoke(&_Stamp.TransactOpts, to, value)
}

// Revoke is a paid mutator transaction binding the contract method 0xeac449d9.
//
// Solidity: function revoke(address to, uint256 value) returns()
func (_Stamp *StampTransactorSession) Revoke(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Stamp.Contract.Revoke(&_Stamp.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns()
func (_Stamp *StampTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Stamp.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns()
func (_Stamp *StampSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Stamp.Contract.Transfer(&_Stamp.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns()
func (_Stamp *StampTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Stamp.Contract.Transfer(&_Stamp.TransactOpts, to, value)
}

// StampAdminOPIterator is returned from FilterAdminOP and is used to iterate over the raw logs and unpacked data for AdminOP events raised by the Stamp contract.
type StampAdminOPIterator struct {
	Event *StampAdminOP // Event containing the contract specifics and raw log

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
func (it *StampAdminOPIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampAdminOP)
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
		it.Event = new(StampAdminOP)
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
func (it *StampAdminOPIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampAdminOPIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampAdminOP represents a AdminOP event raised by the Stamp contract.
type StampAdminOP struct {
	NewAdmin common.Address
	Add      bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminOP is a free log retrieval operation binding the contract event 0x1994fc2eb43ee6b946efd903e39930076e6b3241d31ff0ffe9e583b5b913470d.
//
// Solidity: event AdminOP(address indexed newAdmin, bool add)
func (_Stamp *StampFilterer) FilterAdminOP(opts *bind.FilterOpts, newAdmin []common.Address) (*StampAdminOPIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Stamp.contract.FilterLogs(opts, "AdminOP", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &StampAdminOPIterator{contract: _Stamp.contract, event: "AdminOP", logs: logs, sub: sub}, nil
}

// WatchAdminOP is a free log subscription operation binding the contract event 0x1994fc2eb43ee6b946efd903e39930076e6b3241d31ff0ffe9e583b5b913470d.
//
// Solidity: event AdminOP(address indexed newAdmin, bool add)
func (_Stamp *StampFilterer) WatchAdminOP(opts *bind.WatchOpts, sink chan<- *StampAdminOP, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Stamp.contract.WatchLogs(opts, "AdminOP", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampAdminOP)
				if err := _Stamp.contract.UnpackLog(event, "AdminOP", log); err != nil {
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

// ParseAdminOP is a log parse operation binding the contract event 0x1994fc2eb43ee6b946efd903e39930076e6b3241d31ff0ffe9e583b5b913470d.
//
// Solidity: event AdminOP(address indexed newAdmin, bool add)
func (_Stamp *StampFilterer) ParseAdminOP(log types.Log) (*StampAdminOP, error) {
	event := new(StampAdminOP)
	if err := _Stamp.contract.UnpackLog(event, "AdminOP", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StampOwnerSetIterator is returned from FilterOwnerSet and is used to iterate over the raw logs and unpacked data for OwnerSet events raised by the Stamp contract.
type StampOwnerSetIterator struct {
	Event *StampOwnerSet // Event containing the contract specifics and raw log

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
func (it *StampOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampOwnerSet)
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
		it.Event = new(StampOwnerSet)
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
func (it *StampOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampOwnerSet represents a OwnerSet event raised by the Stamp contract.
type StampOwnerSet struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerSet is a free log retrieval operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_Stamp *StampFilterer) FilterOwnerSet(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*StampOwnerSetIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stamp.contract.FilterLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StampOwnerSetIterator{contract: _Stamp.contract, event: "OwnerSet", logs: logs, sub: sub}, nil
}

// WatchOwnerSet is a free log subscription operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_Stamp *StampFilterer) WatchOwnerSet(opts *bind.WatchOpts, sink chan<- *StampOwnerSet, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stamp.contract.WatchLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampOwnerSet)
				if err := _Stamp.contract.UnpackLog(event, "OwnerSet", log); err != nil {
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

// ParseOwnerSet is a log parse operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_Stamp *StampFilterer) ParseOwnerSet(log types.Log) (*StampOwnerSet, error) {
	event := new(StampOwnerSet)
	if err := _Stamp.contract.UnpackLog(event, "OwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StampRevokeIterator is returned from FilterRevoke and is used to iterate over the raw logs and unpacked data for Revoke events raised by the Stamp contract.
type StampRevokeIterator struct {
	Event *StampRevoke // Event containing the contract specifics and raw log

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
func (it *StampRevokeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampRevoke)
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
		it.Event = new(StampRevoke)
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
func (it *StampRevokeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampRevokeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampRevoke represents a Revoke event raised by the Stamp contract.
type StampRevoke struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRevoke is a free log retrieval operation binding the contract event 0xec9ab91322523c899ede7830ec9bfc992b5981cdcc27b91162fb23de5791117b.
//
// Solidity: event Revoke(address indexed to, uint256 value)
func (_Stamp *StampFilterer) FilterRevoke(opts *bind.FilterOpts, to []common.Address) (*StampRevokeIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Stamp.contract.FilterLogs(opts, "Revoke", toRule)
	if err != nil {
		return nil, err
	}
	return &StampRevokeIterator{contract: _Stamp.contract, event: "Revoke", logs: logs, sub: sub}, nil
}

// WatchRevoke is a free log subscription operation binding the contract event 0xec9ab91322523c899ede7830ec9bfc992b5981cdcc27b91162fb23de5791117b.
//
// Solidity: event Revoke(address indexed to, uint256 value)
func (_Stamp *StampFilterer) WatchRevoke(opts *bind.WatchOpts, sink chan<- *StampRevoke, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Stamp.contract.WatchLogs(opts, "Revoke", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampRevoke)
				if err := _Stamp.contract.UnpackLog(event, "Revoke", log); err != nil {
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

// ParseRevoke is a log parse operation binding the contract event 0xec9ab91322523c899ede7830ec9bfc992b5981cdcc27b91162fb23de5791117b.
//
// Solidity: event Revoke(address indexed to, uint256 value)
func (_Stamp *StampFilterer) ParseRevoke(log types.Log) (*StampRevoke, error) {
	event := new(StampRevoke)
	if err := _Stamp.contract.UnpackLog(event, "Revoke", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StampTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Stamp contract.
type StampTransferIterator struct {
	Event *StampTransfer // Event containing the contract specifics and raw log

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
func (it *StampTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StampTransfer)
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
		it.Event = new(StampTransfer)
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
func (it *StampTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StampTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StampTransfer represents a Transfer event raised by the Stamp contract.
type StampTransfer struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address indexed to, uint256 value)
func (_Stamp *StampFilterer) FilterTransfer(opts *bind.FilterOpts, to []common.Address) (*StampTransferIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Stamp.contract.FilterLogs(opts, "Transfer", toRule)
	if err != nil {
		return nil, err
	}
	return &StampTransferIterator{contract: _Stamp.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address indexed to, uint256 value)
func (_Stamp *StampFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StampTransfer, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Stamp.contract.WatchLogs(opts, "Transfer", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StampTransfer)
				if err := _Stamp.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address indexed to, uint256 value)
func (_Stamp *StampFilterer) ParseTransfer(log types.Log) (*StampTransfer, error) {
	event := new(StampTransfer)
	if err := _Stamp.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
