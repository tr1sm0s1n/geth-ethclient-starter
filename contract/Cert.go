// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// CertCertificate is an auto generated low-level Go binding around an user-defined struct.
type CertCertificate struct {
	Name   string
	Course string
	Grade  string
	Date   string
}

// CertMetaData contains all meta data concerning the Cert contract.
var CertMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"course\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"}],\"name\":\"Issued\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"fetch_certificate\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"course\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"grade\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"}],\"internalType\":\"structCert.Certificate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_course\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_grade\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"issue_certificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610bfc8061005b5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c8063078aea04146100385780632b9c2b7714610068575b5f5ffd5b610052600480360381019061004d91906104b5565b610084565b60405161005f91906105c5565b60405180910390f35b610082600480360381019061007d9190610711565b6102ee565b005b61008c610449565b60015f8381526020019081526020015f206040518060800160405290815f820180546100b790610825565b80601f01602080910402602001604051908101604052809291908181526020018280546100e390610825565b801561012e5780601f106101055761010080835404028352916020019161012e565b820191905f5260205f20905b81548152906001019060200180831161011157829003601f168201915b5050505050815260200160018201805461014790610825565b80601f016020809104026020016040519081016040528092919081815260200182805461017390610825565b80156101be5780601f10610195576101008083540402835291602001916101be565b820191905f5260205f20905b8154815290600101906020018083116101a157829003601f168201915b505050505081526020016002820180546101d790610825565b80601f016020809104026020016040519081016040528092919081815260200182805461020390610825565b801561024e5780601f106102255761010080835404028352916020019161024e565b820191905f5260205f20905b81548152906001019060200180831161023157829003601f168201915b5050505050815260200160038201805461026790610825565b80601f016020809104026020016040519081016040528092919081815260200182805461029390610825565b80156102de5780601f106102b5576101008083540402835291602001916102de565b820191905f5260205f20905b8154815290600101906020018083116102c157829003601f168201915b5050505050815250509050919050565b3373ffffffffffffffffffffffffffffffffffffffff165f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461037c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610373906108af565b60405180910390fd5b60405180608001604052808581526020018481526020018381526020018281525060015f8781526020019081526020015f205f820151815f0190816103c19190610a6d565b5060208201518160010190816103d79190610a6d565b5060408201518160020190816103ed9190610a6d565b5060608201518160030190816104039190610a6d565b509050507fc7fc4858662996e8fe091a9941384ab745bfc51199d7c0a5e45a791ff945993c83868460405161043a93929190610b83565b60405180910390a15050505050565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b61049481610482565b811461049e575f5ffd5b50565b5f813590506104af8161048b565b92915050565b5f602082840312156104ca576104c961047a565b5b5f6104d7848285016104a1565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610522826104e0565b61052c81856104ea565b935061053c8185602086016104fa565b61054581610508565b840191505092915050565b5f608083015f8301518482035f86015261056a8282610518565b915050602083015184820360208601526105848282610518565b9150506040830151848203604086015261059e8282610518565b915050606083015184820360608601526105b88282610518565b9150508091505092915050565b5f6020820190508181035f8301526105dd8184610550565b905092915050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61062382610508565b810181811067ffffffffffffffff82111715610642576106416105ed565b5b80604052505050565b5f610654610471565b9050610660828261061a565b919050565b5f67ffffffffffffffff82111561067f5761067e6105ed565b5b61068882610508565b9050602081019050919050565b828183375f83830152505050565b5f6106b56106b084610665565b61064b565b9050828152602081018484840111156106d1576106d06105e9565b5b6106dc848285610695565b509392505050565b5f82601f8301126106f8576106f76105e5565b5b81356107088482602086016106a3565b91505092915050565b5f5f5f5f5f60a0868803121561072a5761072961047a565b5b5f610737888289016104a1565b955050602086013567ffffffffffffffff8111156107585761075761047e565b5b610764888289016106e4565b945050604086013567ffffffffffffffff8111156107855761078461047e565b5b610791888289016106e4565b935050606086013567ffffffffffffffff8111156107b2576107b161047e565b5b6107be888289016106e4565b925050608086013567ffffffffffffffff8111156107df576107de61047e565b5b6107eb888289016106e4565b9150509295509295909350565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061083c57607f821691505b60208210810361084f5761084e6107f8565b5b50919050565b5f82825260208201905092915050565b7f4163636573732044656e696564000000000000000000000000000000000000005f82015250565b5f610899600d83610855565b91506108a482610865565b602082019050919050565b5f6020820190508181035f8301526108c68161088d565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026109297fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826108ee565b61093386836108ee565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61096e61096961096484610482565b61094b565b610482565b9050919050565b5f819050919050565b61098783610954565b61099b61099382610975565b8484546108fa565b825550505050565b5f5f905090565b6109b26109a3565b6109bd81848461097e565b505050565b5b818110156109e0576109d55f826109aa565b6001810190506109c3565b5050565b601f821115610a25576109f6816108cd565b6109ff846108df565b81016020851015610a0e578190505b610a22610a1a856108df565b8301826109c2565b50505b505050565b5f82821c905092915050565b5f610a455f1984600802610a2a565b1980831691505092915050565b5f610a5d8383610a36565b9150826002028217905092915050565b610a76826104e0565b67ffffffffffffffff811115610a8f57610a8e6105ed565b5b610a998254610825565b610aa48282856109e4565b5f60209050601f831160018114610ad5575f8415610ac3578287015190505b610acd8582610a52565b865550610b34565b601f198416610ae3866108cd565b5f5b82811015610b0a57848901518255600182019150602085019450602081019050610ae5565b86831015610b275784890151610b23601f891682610a36565b8355505b6001600288020188555050505b505050505050565b5f610b46826104e0565b610b508185610855565b9350610b608185602086016104fa565b610b6981610508565b840191505092915050565b610b7d81610482565b82525050565b5f6060820190508181035f830152610b9b8186610b3c565b9050610baa6020830185610b74565b8181036040830152610bbc8184610b3c565b905094935050505056fea2646970667358221220ed42e4b4f687c0616b40873d13a22d46ca8bd1e9f3b8e94ea4205722440ab5c964736f6c634300081c0033",
}

// CertABI is the input ABI used to generate the binding from.
// Deprecated: Use CertMetaData.ABI instead.
var CertABI = CertMetaData.ABI

// CertBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CertMetaData.Bin instead.
var CertBin = CertMetaData.Bin

// DeployCert deploys a new Ethereum contract, binding an instance of Cert to it.
func DeployCert(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Cert, error) {
	parsed, err := CertMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CertBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cert{CertCaller: CertCaller{contract: contract}, CertTransactor: CertTransactor{contract: contract}, CertFilterer: CertFilterer{contract: contract}}, nil
}

// Cert is an auto generated Go binding around an Ethereum contract.
type Cert struct {
	CertCaller     // Read-only binding to the contract
	CertTransactor // Write-only binding to the contract
	CertFilterer   // Log filterer for contract events
}

// CertCaller is an auto generated read-only Go binding around an Ethereum contract.
type CertCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CertTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CertFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CertSession struct {
	Contract     *Cert             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CertCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CertCallerSession struct {
	Contract *CertCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CertTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CertTransactorSession struct {
	Contract     *CertTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CertRaw is an auto generated low-level Go binding around an Ethereum contract.
type CertRaw struct {
	Contract *Cert // Generic contract binding to access the raw methods on
}

// CertCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CertCallerRaw struct {
	Contract *CertCaller // Generic read-only contract binding to access the raw methods on
}

// CertTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CertTransactorRaw struct {
	Contract *CertTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCert creates a new instance of Cert, bound to a specific deployed contract.
func NewCert(address common.Address, backend bind.ContractBackend) (*Cert, error) {
	contract, err := bindCert(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cert{CertCaller: CertCaller{contract: contract}, CertTransactor: CertTransactor{contract: contract}, CertFilterer: CertFilterer{contract: contract}}, nil
}

// NewCertCaller creates a new read-only instance of Cert, bound to a specific deployed contract.
func NewCertCaller(address common.Address, caller bind.ContractCaller) (*CertCaller, error) {
	contract, err := bindCert(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CertCaller{contract: contract}, nil
}

// NewCertTransactor creates a new write-only instance of Cert, bound to a specific deployed contract.
func NewCertTransactor(address common.Address, transactor bind.ContractTransactor) (*CertTransactor, error) {
	contract, err := bindCert(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CertTransactor{contract: contract}, nil
}

// NewCertFilterer creates a new log filterer instance of Cert, bound to a specific deployed contract.
func NewCertFilterer(address common.Address, filterer bind.ContractFilterer) (*CertFilterer, error) {
	contract, err := bindCert(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CertFilterer{contract: contract}, nil
}

// bindCert binds a generic wrapper to an already deployed contract.
func bindCert(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CertMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cert *CertRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cert.Contract.CertCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cert *CertRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cert.Contract.CertTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cert *CertRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cert.Contract.CertTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cert *CertCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cert.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cert *CertTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cert.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cert *CertTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cert.Contract.contract.Transact(opts, method, params...)
}

// FetchCertificate is a free data retrieval call binding the contract method 0x078aea04.
//
// Solidity: function fetch_certificate(uint256 _id) view returns((string,string,string,string))
func (_Cert *CertCaller) FetchCertificate(opts *bind.CallOpts, _id *big.Int) (CertCertificate, error) {
	var out []interface{}
	err := _Cert.contract.Call(opts, &out, "fetch_certificate", _id)

	if err != nil {
		return *new(CertCertificate), err
	}

	out0 := *abi.ConvertType(out[0], new(CertCertificate)).(*CertCertificate)

	return out0, err

}

// FetchCertificate is a free data retrieval call binding the contract method 0x078aea04.
//
// Solidity: function fetch_certificate(uint256 _id) view returns((string,string,string,string))
func (_Cert *CertSession) FetchCertificate(_id *big.Int) (CertCertificate, error) {
	return _Cert.Contract.FetchCertificate(&_Cert.CallOpts, _id)
}

// FetchCertificate is a free data retrieval call binding the contract method 0x078aea04.
//
// Solidity: function fetch_certificate(uint256 _id) view returns((string,string,string,string))
func (_Cert *CertCallerSession) FetchCertificate(_id *big.Int) (CertCertificate, error) {
	return _Cert.Contract.FetchCertificate(&_Cert.CallOpts, _id)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0x2b9c2b77.
//
// Solidity: function issue_certificate(uint256 _id, string _name, string _course, string _grade, string _date) returns()
func (_Cert *CertTransactor) IssueCertificate(opts *bind.TransactOpts, _id *big.Int, _name string, _course string, _grade string, _date string) (*types.Transaction, error) {
	return _Cert.contract.Transact(opts, "issue_certificate", _id, _name, _course, _grade, _date)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0x2b9c2b77.
//
// Solidity: function issue_certificate(uint256 _id, string _name, string _course, string _grade, string _date) returns()
func (_Cert *CertSession) IssueCertificate(_id *big.Int, _name string, _course string, _grade string, _date string) (*types.Transaction, error) {
	return _Cert.Contract.IssueCertificate(&_Cert.TransactOpts, _id, _name, _course, _grade, _date)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0x2b9c2b77.
//
// Solidity: function issue_certificate(uint256 _id, string _name, string _course, string _grade, string _date) returns()
func (_Cert *CertTransactorSession) IssueCertificate(_id *big.Int, _name string, _course string, _grade string, _date string) (*types.Transaction, error) {
	return _Cert.Contract.IssueCertificate(&_Cert.TransactOpts, _id, _name, _course, _grade, _date)
}

// CertIssuedIterator is returned from FilterIssued and is used to iterate over the raw logs and unpacked data for Issued events raised by the Cert contract.
type CertIssuedIterator struct {
	Event *CertIssued // Event containing the contract specifics and raw log

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
func (it *CertIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertIssued)
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
		it.Event = new(CertIssued)
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
func (it *CertIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertIssued represents a Issued event raised by the Cert contract.
type CertIssued struct {
	Course string
	Id     *big.Int
	Date   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssued is a free log retrieval operation binding the contract event 0xc7fc4858662996e8fe091a9941384ab745bfc51199d7c0a5e45a791ff945993c.
//
// Solidity: event Issued(string course, uint256 id, string date)
func (_Cert *CertFilterer) FilterIssued(opts *bind.FilterOpts) (*CertIssuedIterator, error) {

	logs, sub, err := _Cert.contract.FilterLogs(opts, "Issued")
	if err != nil {
		return nil, err
	}
	return &CertIssuedIterator{contract: _Cert.contract, event: "Issued", logs: logs, sub: sub}, nil
}

// WatchIssued is a free log subscription operation binding the contract event 0xc7fc4858662996e8fe091a9941384ab745bfc51199d7c0a5e45a791ff945993c.
//
// Solidity: event Issued(string course, uint256 id, string date)
func (_Cert *CertFilterer) WatchIssued(opts *bind.WatchOpts, sink chan<- *CertIssued) (event.Subscription, error) {

	logs, sub, err := _Cert.contract.WatchLogs(opts, "Issued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertIssued)
				if err := _Cert.contract.UnpackLog(event, "Issued", log); err != nil {
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

// ParseIssued is a log parse operation binding the contract event 0xc7fc4858662996e8fe091a9941384ab745bfc51199d7c0a5e45a791ff945993c.
//
// Solidity: event Issued(string course, uint256 id, string date)
func (_Cert *CertFilterer) ParseIssued(log types.Log) (*CertIssued, error) {
	event := new(CertIssued)
	if err := _Cert.contract.UnpackLog(event, "Issued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
