// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lib

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Name   string
	Course string
	Grade  string
	Date   string
}

// CertMetaData contains all meta data concerning the Cert contract.
var CertMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"name\":\"__init__\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"issue_certificate\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"course\",\"type\":\"string\"},{\"name\":\"grade\",\"type\":\"string\"},{\"name\":\"date\",\"type\":\"string\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"fetch_certificate\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"course\",\"type\":\"string\"},{\"name\":\"grade\",\"type\":\"string\"},{\"name\":\"date\",\"type\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AccessDenied\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Issued\",\"inputs\":[{\"name\":\"course\",\"type\":\"string\",\"indexed\":true},{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"date\",\"type\":\"string\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Certificate\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false},{\"name\":\"course\",\"type\":\"string\",\"indexed\":false},{\"name\":\"grade\",\"type\":\"string\",\"indexed\":false},{\"name\":\"date\",\"type\":\"string\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Context\",\"inputs\":[],\"anonymous\":false}]",
	Bin: "0x61082380380390604051801561003a575b82810160405239600080546001600160601b03163360601b1781556107e090819061004390396000f35b50606061001056fe60003560e01c80632b9c2b77146100465763078aea041461001c57005b61003161002c3660031901610429565b6101c7565b61004361003c610253565b91826105b2565b90f35b366003190161012081106101a08211176100aa576100626102e6565b61007160a48294939401610336565b9091019061008160a48301610387565b9092019260a061009360a486016103d8565b90950101036100aa576100a8936004356100af565b005b610265565b94939291909460ff60005460601c33141516806001146101a557156100d657505050509050565b60405190811561019c575b60388201604052818760005b6001811061018657509088610158939261015d999a038060380361015f575b50505061012161011a61028f565b8095610436565b61013361012c6102ac565b8096610489565b61014561013e6102c9565b80976104dc565b60005260208052604060002060ff191690565b610747565b565b60016000199160380360031b610100031b019081199051169082511617905238808061010c565b895183526020998a0199909201916001016100ed565b606091506100e1565b60046101af610253565b80516001600160e01b0316634ca8886760e01b178152fd5b6000908152602080526040902060ff191660405190811561024a575b60e88201604052819060051c60005b6007811061023457508282038060e80361020d575b50505090565b60016000199160e80360031b610100031b0190811990541690825116179052388080610207565b91906020600182819354855201920192016101f2565b606091506101e3565b60405190811561025f57565b60609150565b602461026f610253565b80516001600160e01b0316632e36a70960e01b1781526101036004820152fd5b6040519081156102a3575b60328201604052565b6060915061029a565b6040519081156102c0575b60218201604052565b606091506102b7565b6040519081156102dd575b60288201604052565b606091506102d4565b60a43590601882116100aa57601f199081603f840116916084830135908484030160031b610100031b6100aa57604051801561032e575b603881016040529260200160a48437565b50606061031d565b803591601283116100aa57601f199182603f85011692808483010135908585030160031b610100031b6100aa5760405190811561037e575b6032820160405260208295019137565b6060915061036e565b803591600183116100aa57601f199182603f85011692808483010135908585030160031b610100031b6100aa576040519081156103cf575b6021820160405260208295019137565b606091506103bf565b803591600883116100aa57601f199182603f85011692808483010135908585030160031b610100031b6100aa57604051908115610420575b6028820160405260208295019137565b60609150610410565b6020036100aa5760043590565b8060005b60018110610473575081038060320361045257505050565b60016000199160320360031b610100031b0190811990511690825116179052565b825184526020938401939092019160010161043a565b8060005b600181106104c657508103806021036104a557505050565b60016000199160210360031b610100031b0190811990511690825116179052565b825184526020938401939092019160010161048d565b8060005b6001811061051957508103806028036104f857505050565b60016000199160280360031b610100031b0190811990511690825116179052565b82518452602093840193909201916001016104e0565b919091805192602093808501601f1995600087603f8501169788860101528493948260051c6000905b80821061059b5750508403809203610572575b5050505050565b60019160001993030160031b610100031b0190811990511690825116179052388080808061056b565b909695838082600193518a52019701970190610558565b61061960a0926020835260c08460808060208701526105d38287018561052f565b81810160408801526105ec83828901016040870161052f565b018181016060880152610605838289010183870161052f565b01958187928301908201520101910161052f565b010190565b919060051c91806000905b6001908183101561064a579060208260019351885501950191019093610629565b9291505092919281038060320361066057505050565b60016000199160320360031b610100031b0190811990511690825416179055565b919060051c91806000905b600190818310156106ad57906020826001935188550195019101909361068c565b929150509291928103806021036106c357505050565b60016000199160210360031b610100031b0190811990511690825416179055565b919060051c91806000905b600190818310156107105790602082600193518855019501910190936106ef565b9291505092919281038060280361072657505050565b60016000199160280360031b610100031b0190811990511690825416179055565b91939490929483928060051c966000945b600190818710156107795790602082600193518c5501990195019497610758565b60c0965061015d989150986107a8929395996107b295988103806038036107b9575b505050604085019061061e565b6080830190610681565b01906106e4565b60016000199160380360031b610100031b019081199051169082541617905538808061079b56",
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
// Solidity: function fetch_certificate(uint256 id) view returns((string,string,string,string))
func (_Cert *CertCaller) FetchCertificate(opts *bind.CallOpts, id *big.Int) (Struct0, error) {
	var out []interface{}
	err := _Cert.contract.Call(opts, &out, "fetch_certificate", id)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// FetchCertificate is a free data retrieval call binding the contract method 0x078aea04.
//
// Solidity: function fetch_certificate(uint256 id) view returns((string,string,string,string))
func (_Cert *CertSession) FetchCertificate(id *big.Int) (Struct0, error) {
	return _Cert.Contract.FetchCertificate(&_Cert.CallOpts, id)
}

// FetchCertificate is a free data retrieval call binding the contract method 0x078aea04.
//
// Solidity: function fetch_certificate(uint256 id) view returns((string,string,string,string))
func (_Cert *CertCallerSession) FetchCertificate(id *big.Int) (Struct0, error) {
	return _Cert.Contract.FetchCertificate(&_Cert.CallOpts, id)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0x2b9c2b77.
//
// Solidity: function issue_certificate(uint256 id, string name, string course, string grade, string date) payable returns()
func (_Cert *CertTransactor) IssueCertificate(opts *bind.TransactOpts, id *big.Int, name string, course string, grade string, date string) (*types.Transaction, error) {
	return _Cert.contract.Transact(opts, "issue_certificate", id, name, course, grade, date)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0x2b9c2b77.
//
// Solidity: function issue_certificate(uint256 id, string name, string course, string grade, string date) payable returns()
func (_Cert *CertSession) IssueCertificate(id *big.Int, name string, course string, grade string, date string) (*types.Transaction, error) {
	return _Cert.Contract.IssueCertificate(&_Cert.TransactOpts, id, name, course, grade, date)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0x2b9c2b77.
//
// Solidity: function issue_certificate(uint256 id, string name, string course, string grade, string date) payable returns()
func (_Cert *CertTransactorSession) IssueCertificate(id *big.Int, name string, course string, grade string, date string) (*types.Transaction, error) {
	return _Cert.Contract.IssueCertificate(&_Cert.TransactOpts, id, name, course, grade, date)
}

// CertAccessDeniedIterator is returned from FilterAccessDenied and is used to iterate over the raw logs and unpacked data for AccessDenied events raised by the Cert contract.
type CertAccessDeniedIterator struct {
	Event *CertAccessDenied // Event containing the contract specifics and raw log

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
func (it *CertAccessDeniedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertAccessDenied)
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
		it.Event = new(CertAccessDenied)
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
func (it *CertAccessDeniedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertAccessDeniedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertAccessDenied represents a AccessDenied event raised by the Cert contract.
type CertAccessDenied struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAccessDenied is a free log retrieval operation binding the contract event 0x4ca888678577e39e3ac9ac6cfe78050ee33f0c0eec2c686bc222b2cf4140a62c.
//
// Solidity: event AccessDenied()
func (_Cert *CertFilterer) FilterAccessDenied(opts *bind.FilterOpts) (*CertAccessDeniedIterator, error) {

	logs, sub, err := _Cert.contract.FilterLogs(opts, "AccessDenied")
	if err != nil {
		return nil, err
	}
	return &CertAccessDeniedIterator{contract: _Cert.contract, event: "AccessDenied", logs: logs, sub: sub}, nil
}

// WatchAccessDenied is a free log subscription operation binding the contract event 0x4ca888678577e39e3ac9ac6cfe78050ee33f0c0eec2c686bc222b2cf4140a62c.
//
// Solidity: event AccessDenied()
func (_Cert *CertFilterer) WatchAccessDenied(opts *bind.WatchOpts, sink chan<- *CertAccessDenied) (event.Subscription, error) {

	logs, sub, err := _Cert.contract.WatchLogs(opts, "AccessDenied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertAccessDenied)
				if err := _Cert.contract.UnpackLog(event, "AccessDenied", log); err != nil {
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

// ParseAccessDenied is a log parse operation binding the contract event 0x4ca888678577e39e3ac9ac6cfe78050ee33f0c0eec2c686bc222b2cf4140a62c.
//
// Solidity: event AccessDenied()
func (_Cert *CertFilterer) ParseAccessDenied(log types.Log) (*CertAccessDenied, error) {
	event := new(CertAccessDenied)
	if err := _Cert.contract.UnpackLog(event, "AccessDenied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CertCertificateIterator is returned from FilterCertificate and is used to iterate over the raw logs and unpacked data for Certificate events raised by the Cert contract.
type CertCertificateIterator struct {
	Event *CertCertificate // Event containing the contract specifics and raw log

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
func (it *CertCertificateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertCertificate)
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
		it.Event = new(CertCertificate)
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
func (it *CertCertificateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertCertificateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertCertificate represents a Certificate event raised by the Cert contract.
type CertCertificate struct {
	Name   string
	Course string
	Grade  string
	Date   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCertificate is a free log retrieval operation binding the contract event 0x4818a178390df554e2726b3a8850661069bdd2da8f62cdaa1f8f1380087b2b88.
//
// Solidity: event Certificate(string name, string course, string grade, string date)
func (_Cert *CertFilterer) FilterCertificate(opts *bind.FilterOpts) (*CertCertificateIterator, error) {

	logs, sub, err := _Cert.contract.FilterLogs(opts, "Certificate")
	if err != nil {
		return nil, err
	}
	return &CertCertificateIterator{contract: _Cert.contract, event: "Certificate", logs: logs, sub: sub}, nil
}

// WatchCertificate is a free log subscription operation binding the contract event 0x4818a178390df554e2726b3a8850661069bdd2da8f62cdaa1f8f1380087b2b88.
//
// Solidity: event Certificate(string name, string course, string grade, string date)
func (_Cert *CertFilterer) WatchCertificate(opts *bind.WatchOpts, sink chan<- *CertCertificate) (event.Subscription, error) {

	logs, sub, err := _Cert.contract.WatchLogs(opts, "Certificate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertCertificate)
				if err := _Cert.contract.UnpackLog(event, "Certificate", log); err != nil {
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

// ParseCertificate is a log parse operation binding the contract event 0x4818a178390df554e2726b3a8850661069bdd2da8f62cdaa1f8f1380087b2b88.
//
// Solidity: event Certificate(string name, string course, string grade, string date)
func (_Cert *CertFilterer) ParseCertificate(log types.Log) (*CertCertificate, error) {
	event := new(CertCertificate)
	if err := _Cert.contract.UnpackLog(event, "Certificate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CertContextIterator is returned from FilterContext and is used to iterate over the raw logs and unpacked data for Context events raised by the Cert contract.
type CertContextIterator struct {
	Event *CertContext // Event containing the contract specifics and raw log

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
func (it *CertContextIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertContext)
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
		it.Event = new(CertContext)
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
func (it *CertContextIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertContextIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertContext represents a Context event raised by the Cert contract.
type CertContext struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterContext is a free log retrieval operation binding the contract event 0x9aca9e90d8c35dda9a7220c7e3b36fe8fd9533ee4071250d23f571bbce60c246.
//
// Solidity: event Context()
func (_Cert *CertFilterer) FilterContext(opts *bind.FilterOpts) (*CertContextIterator, error) {

	logs, sub, err := _Cert.contract.FilterLogs(opts, "Context")
	if err != nil {
		return nil, err
	}
	return &CertContextIterator{contract: _Cert.contract, event: "Context", logs: logs, sub: sub}, nil
}

// WatchContext is a free log subscription operation binding the contract event 0x9aca9e90d8c35dda9a7220c7e3b36fe8fd9533ee4071250d23f571bbce60c246.
//
// Solidity: event Context()
func (_Cert *CertFilterer) WatchContext(opts *bind.WatchOpts, sink chan<- *CertContext) (event.Subscription, error) {

	logs, sub, err := _Cert.contract.WatchLogs(opts, "Context")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertContext)
				if err := _Cert.contract.UnpackLog(event, "Context", log); err != nil {
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

// ParseContext is a log parse operation binding the contract event 0x9aca9e90d8c35dda9a7220c7e3b36fe8fd9533ee4071250d23f571bbce60c246.
//
// Solidity: event Context()
func (_Cert *CertFilterer) ParseContext(log types.Log) (*CertContext, error) {
	event := new(CertContext)
	if err := _Cert.contract.UnpackLog(event, "Context", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Course common.Hash
	Id     *big.Int
	Date   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssued is a free log retrieval operation binding the contract event 0xc7fc4858662996e8fe091a9941384ab745bfc51199d7c0a5e45a791ff945993c.
//
// Solidity: event Issued(string indexed course, uint256 id, string date)
func (_Cert *CertFilterer) FilterIssued(opts *bind.FilterOpts, course []string) (*CertIssuedIterator, error) {

	var courseRule []interface{}
	for _, courseItem := range course {
		courseRule = append(courseRule, courseItem)
	}

	logs, sub, err := _Cert.contract.FilterLogs(opts, "Issued", courseRule)
	if err != nil {
		return nil, err
	}
	return &CertIssuedIterator{contract: _Cert.contract, event: "Issued", logs: logs, sub: sub}, nil
}

// WatchIssued is a free log subscription operation binding the contract event 0xc7fc4858662996e8fe091a9941384ab745bfc51199d7c0a5e45a791ff945993c.
//
// Solidity: event Issued(string indexed course, uint256 id, string date)
func (_Cert *CertFilterer) WatchIssued(opts *bind.WatchOpts, sink chan<- *CertIssued, course []string) (event.Subscription, error) {

	var courseRule []interface{}
	for _, courseItem := range course {
		courseRule = append(courseRule, courseItem)
	}

	logs, sub, err := _Cert.contract.WatchLogs(opts, "Issued", courseRule)
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
// Solidity: event Issued(string indexed course, uint256 id, string date)
func (_Cert *CertFilterer) ParseIssued(log types.Log) (*CertIssued, error) {
	event := new(CertIssued)
	if err := _Cert.contract.UnpackLog(event, "Issued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
