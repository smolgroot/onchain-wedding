// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package utils

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

// WeddingFiance is an auto generated low-level Go binding around an user-defined struct.
type WeddingFiance struct {
	Sender      common.Address
	Description string
	Timestamp   *big.Int
}

// WeddingMessage is an auto generated low-level Go binding around an user-defined struct.
type WeddingMessage struct {
	Sender    common.Address
	Content   string
	Timestamp *big.Int
}

// WeddingMetaData contains all meta data concerning the Wedding contract.
var WeddingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"NewMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"NewSignature\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getMessages\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structWedding.Message[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sayYes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structWedding.Fiance\",\"name\":\"fiance\",\"type\":\"tuple\"}],\"name\":\"updateFianceDescription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610e54806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063469c8110146100515780635ff6cbf31461006d5780639bee33651461008b578063d9b667c414610095575b600080fd5b61006b600480360381019061006691906104e2565b6100b1565b005b610075610210565b604051610082919061072b565b60405180910390f35b610093610361565b005b6100af60048036038101906100aa919061095f565b610466565b005b600760405180606001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200184848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200142815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190816101ad9190610bb4565b506040820151816002015550503373ffffffffffffffffffffffffffffffffffffffff167f0ff94fec3112de81726d79117e091c7c9d47da8a38c73e8cec7ee350a61898a642848460405161020493929190610cd3565b60405180910390a25050565b60606007805480602002602001604051908101604052809291908181526020016000905b8282101561035857838290600052602060002090600302016040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820180546102bd906109d7565b80601f01602080910402602001604051908101604052809291908181526020018280546102e9906109d7565b80156103365780601f1061030b57610100808354040283529160200191610336565b820191906000526020600020905b81548152906001019060200180831161031957829003601f168201915b5050505050815260200160028201548152505081526020019060010190610234565b50505050905090565b6006600090806001815401808255809150506001900390600052602060002090600302016000909190919091506000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600182018160010190816104099190610d1b565b506002820154816002015550503373ffffffffffffffffffffffffffffffffffffffff167ffb79a88b762f56bcd11bcc12554f81106e62770680055da0027dc8fe0694e96c4260405161045c9190610e03565b60405180910390a2565b50565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f8401126104a2576104a161047d565b5b8235905067ffffffffffffffff8111156104bf576104be610482565b5b6020830191508360018202830111156104db576104da610487565b5b9250929050565b600080602083850312156104f9576104f8610473565b5b600083013567ffffffffffffffff81111561051757610516610478565b5b6105238582860161048c565b92509250509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105868261055b565b9050919050565b6105968161057b565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156105d65780820151818401526020810190506105bb565b60008484015250505050565b6000601f19601f8301169050919050565b60006105fe8261059c565b61060881856105a7565b93506106188185602086016105b8565b610621816105e2565b840191505092915050565b6000819050919050565b61063f8161062c565b82525050565b600060608301600083015161065d600086018261058d565b506020830151848203602086015261067582826105f3565b915050604083015161068a6040860182610636565b508091505092915050565b60006106a18383610645565b905092915050565b6000602082019050919050565b60006106c18261052f565b6106cb818561053a565b9350836020820285016106dd8561054b565b8060005b8581101561071957848403895281516106fa8582610695565b9450610705836106a9565b925060208a019950506001810190506106e1565b50829750879550505050505092915050565b6000602082019050818103600083015261074581846106b6565b905092915050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61078a826105e2565b810181811067ffffffffffffffff821117156107a9576107a8610752565b5b80604052505050565b60006107bc610469565b90506107c88282610781565b919050565b600080fd5b6107db8161057b565b81146107e657600080fd5b50565b6000813590506107f8816107d2565b92915050565b600080fd5b600067ffffffffffffffff82111561081e5761081d610752565b5b610827826105e2565b9050602081019050919050565b82818337600083830152505050565b600061085661085184610803565b6107b2565b905082815260208101848484011115610872576108716107fe565b5b61087d848285610834565b509392505050565b600082601f83011261089a5761089961047d565b5b81356108aa848260208601610843565b91505092915050565b6108bc8161062c565b81146108c757600080fd5b50565b6000813590506108d9816108b3565b92915050565b6000606082840312156108f5576108f461074d565b5b6108ff60606107b2565b9050600061090f848285016107e9565b600083015250602082013567ffffffffffffffff811115610933576109326107cd565b5b61093f84828501610885565b6020830152506040610953848285016108ca565b60408301525092915050565b60006020828403121561097557610974610473565b5b600082013567ffffffffffffffff81111561099357610992610478565b5b61099f848285016108df565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806109ef57607f821691505b602082108103610a0257610a016109a8565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610a6a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610a2d565b610a748683610a2d565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610ab1610aac610aa78461062c565b610a8c565b61062c565b9050919050565b6000819050919050565b610acb83610a96565b610adf610ad782610ab8565b848454610a3a565b825550505050565b600090565b610af4610ae7565b610aff818484610ac2565b505050565b5b81811015610b2357610b18600082610aec565b600181019050610b05565b5050565b601f821115610b6857610b3981610a08565b610b4284610a1d565b81016020851015610b51578190505b610b65610b5d85610a1d565b830182610b04565b50505b505050565b600082821c905092915050565b6000610b8b60001984600802610b6d565b1980831691505092915050565b6000610ba48383610b7a565b9150826002028217905092915050565b610bbd8261059c565b67ffffffffffffffff811115610bd657610bd5610752565b5b610be082546109d7565b610beb828285610b27565b600060209050601f831160018114610c1e5760008415610c0c578287015190505b610c168582610b98565b865550610c7e565b601f198416610c2c86610a08565b60005b82811015610c5457848901518255600182019150602085019450602081019050610c2f565b86831015610c715784890151610c6d601f891682610b7a565b8355505b6001600288020188555050505b505050505050565b610c8f8161062c565b82525050565b600082825260208201905092915050565b6000610cb28385610c95565b9350610cbf838584610834565b610cc8836105e2565b840190509392505050565b6000604082019050610ce86000830186610c86565b8181036020830152610cfb818486610ca6565b9050949350505050565b600081549050610d14816109d7565b9050919050565b818103610d29575050610e01565b610d3282610d05565b67ffffffffffffffff811115610d4b57610d4a610752565b5b610d5582546109d7565b610d60828285610b27565b6000601f831160018114610d8f5760008415610d7d578287015490505b610d878582610b98565b865550610dfa565b601f198416610d9d87610a08565b9650610da886610a08565b60005b82811015610dd057848901548255600182019150600185019450602081019050610dab565b86831015610ded5784890154610de9601f891682610b7a565b8355505b6001600288020188555050505b5050505050505b565b6000602082019050610e186000830184610c86565b9291505056fea2646970667358221220a287a983aa32954bb8070c17392f4bc6dd65d2789e98747da940d560e488302664736f6c63430008100033",
}

// WeddingABI is the input ABI used to generate the binding from.
// Deprecated: Use WeddingMetaData.ABI instead.
var WeddingABI = WeddingMetaData.ABI

// WeddingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WeddingMetaData.Bin instead.
var WeddingBin = WeddingMetaData.Bin

// DeployWedding deploys a new Ethereum contract, binding an instance of Wedding to it.
func DeployWedding(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Wedding, error) {
	parsed, err := WeddingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WeddingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Wedding{WeddingCaller: WeddingCaller{contract: contract}, WeddingTransactor: WeddingTransactor{contract: contract}, WeddingFilterer: WeddingFilterer{contract: contract}}, nil
}

// Wedding is an auto generated Go binding around an Ethereum contract.
type Wedding struct {
	WeddingCaller     // Read-only binding to the contract
	WeddingTransactor // Write-only binding to the contract
	WeddingFilterer   // Log filterer for contract events
}

// WeddingCaller is an auto generated read-only Go binding around an Ethereum contract.
type WeddingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WeddingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WeddingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WeddingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WeddingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WeddingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WeddingSession struct {
	Contract     *Wedding          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WeddingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WeddingCallerSession struct {
	Contract *WeddingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// WeddingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WeddingTransactorSession struct {
	Contract     *WeddingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// WeddingRaw is an auto generated low-level Go binding around an Ethereum contract.
type WeddingRaw struct {
	Contract *Wedding // Generic contract binding to access the raw methods on
}

// WeddingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WeddingCallerRaw struct {
	Contract *WeddingCaller // Generic read-only contract binding to access the raw methods on
}

// WeddingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WeddingTransactorRaw struct {
	Contract *WeddingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWedding creates a new instance of Wedding, bound to a specific deployed contract.
func NewWedding(address common.Address, backend bind.ContractBackend) (*Wedding, error) {
	contract, err := bindWedding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wedding{WeddingCaller: WeddingCaller{contract: contract}, WeddingTransactor: WeddingTransactor{contract: contract}, WeddingFilterer: WeddingFilterer{contract: contract}}, nil
}

// NewWeddingCaller creates a new read-only instance of Wedding, bound to a specific deployed contract.
func NewWeddingCaller(address common.Address, caller bind.ContractCaller) (*WeddingCaller, error) {
	contract, err := bindWedding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WeddingCaller{contract: contract}, nil
}

// NewWeddingTransactor creates a new write-only instance of Wedding, bound to a specific deployed contract.
func NewWeddingTransactor(address common.Address, transactor bind.ContractTransactor) (*WeddingTransactor, error) {
	contract, err := bindWedding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WeddingTransactor{contract: contract}, nil
}

// NewWeddingFilterer creates a new log filterer instance of Wedding, bound to a specific deployed contract.
func NewWeddingFilterer(address common.Address, filterer bind.ContractFilterer) (*WeddingFilterer, error) {
	contract, err := bindWedding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WeddingFilterer{contract: contract}, nil
}

// bindWedding binds a generic wrapper to an already deployed contract.
func bindWedding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WeddingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wedding *WeddingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wedding.Contract.WeddingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wedding *WeddingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wedding.Contract.WeddingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wedding *WeddingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wedding.Contract.WeddingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wedding *WeddingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wedding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wedding *WeddingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wedding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wedding *WeddingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wedding.Contract.contract.Transact(opts, method, params...)
}

// GetMessages is a free data retrieval call binding the contract method 0x5ff6cbf3.
//
// Solidity: function getMessages() view returns((address,string,uint256)[])
func (_Wedding *WeddingCaller) GetMessages(opts *bind.CallOpts) ([]WeddingMessage, error) {
	var out []interface{}
	err := _Wedding.contract.Call(opts, &out, "getMessages")

	if err != nil {
		return *new([]WeddingMessage), err
	}

	out0 := *abi.ConvertType(out[0], new([]WeddingMessage)).(*[]WeddingMessage)

	return out0, err

}

// GetMessages is a free data retrieval call binding the contract method 0x5ff6cbf3.
//
// Solidity: function getMessages() view returns((address,string,uint256)[])
func (_Wedding *WeddingSession) GetMessages() ([]WeddingMessage, error) {
	return _Wedding.Contract.GetMessages(&_Wedding.CallOpts)
}

// GetMessages is a free data retrieval call binding the contract method 0x5ff6cbf3.
//
// Solidity: function getMessages() view returns((address,string,uint256)[])
func (_Wedding *WeddingCallerSession) GetMessages() ([]WeddingMessage, error) {
	return _Wedding.Contract.GetMessages(&_Wedding.CallOpts)
}

// SayYes is a paid mutator transaction binding the contract method 0x9bee3365.
//
// Solidity: function sayYes() returns()
func (_Wedding *WeddingTransactor) SayYes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wedding.contract.Transact(opts, "sayYes")
}

// SayYes is a paid mutator transaction binding the contract method 0x9bee3365.
//
// Solidity: function sayYes() returns()
func (_Wedding *WeddingSession) SayYes() (*types.Transaction, error) {
	return _Wedding.Contract.SayYes(&_Wedding.TransactOpts)
}

// SayYes is a paid mutator transaction binding the contract method 0x9bee3365.
//
// Solidity: function sayYes() returns()
func (_Wedding *WeddingTransactorSession) SayYes() (*types.Transaction, error) {
	return _Wedding.Contract.SayYes(&_Wedding.TransactOpts)
}

// SendMessage is a paid mutator transaction binding the contract method 0x469c8110.
//
// Solidity: function sendMessage(string _content) returns()
func (_Wedding *WeddingTransactor) SendMessage(opts *bind.TransactOpts, _content string) (*types.Transaction, error) {
	return _Wedding.contract.Transact(opts, "sendMessage", _content)
}

// SendMessage is a paid mutator transaction binding the contract method 0x469c8110.
//
// Solidity: function sendMessage(string _content) returns()
func (_Wedding *WeddingSession) SendMessage(_content string) (*types.Transaction, error) {
	return _Wedding.Contract.SendMessage(&_Wedding.TransactOpts, _content)
}

// SendMessage is a paid mutator transaction binding the contract method 0x469c8110.
//
// Solidity: function sendMessage(string _content) returns()
func (_Wedding *WeddingTransactorSession) SendMessage(_content string) (*types.Transaction, error) {
	return _Wedding.Contract.SendMessage(&_Wedding.TransactOpts, _content)
}

// UpdateFianceDescription is a paid mutator transaction binding the contract method 0xd9b667c4.
//
// Solidity: function updateFianceDescription((address,string,uint256) fiance) returns()
func (_Wedding *WeddingTransactor) UpdateFianceDescription(opts *bind.TransactOpts, fiance WeddingFiance) (*types.Transaction, error) {
	return _Wedding.contract.Transact(opts, "updateFianceDescription", fiance)
}

// UpdateFianceDescription is a paid mutator transaction binding the contract method 0xd9b667c4.
//
// Solidity: function updateFianceDescription((address,string,uint256) fiance) returns()
func (_Wedding *WeddingSession) UpdateFianceDescription(fiance WeddingFiance) (*types.Transaction, error) {
	return _Wedding.Contract.UpdateFianceDescription(&_Wedding.TransactOpts, fiance)
}

// UpdateFianceDescription is a paid mutator transaction binding the contract method 0xd9b667c4.
//
// Solidity: function updateFianceDescription((address,string,uint256) fiance) returns()
func (_Wedding *WeddingTransactorSession) UpdateFianceDescription(fiance WeddingFiance) (*types.Transaction, error) {
	return _Wedding.Contract.UpdateFianceDescription(&_Wedding.TransactOpts, fiance)
}

// WeddingNewMessageIterator is returned from FilterNewMessage and is used to iterate over the raw logs and unpacked data for NewMessage events raised by the Wedding contract.
type WeddingNewMessageIterator struct {
	Event *WeddingNewMessage // Event containing the contract specifics and raw log

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
func (it *WeddingNewMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WeddingNewMessage)
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
		it.Event = new(WeddingNewMessage)
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
func (it *WeddingNewMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WeddingNewMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WeddingNewMessage represents a NewMessage event raised by the Wedding contract.
type WeddingNewMessage struct {
	From      common.Address
	Timestamp *big.Int
	Message   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewMessage is a free log retrieval operation binding the contract event 0x0ff94fec3112de81726d79117e091c7c9d47da8a38c73e8cec7ee350a61898a6.
//
// Solidity: event NewMessage(address indexed from, uint256 timestamp, string message)
func (_Wedding *WeddingFilterer) FilterNewMessage(opts *bind.FilterOpts, from []common.Address) (*WeddingNewMessageIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Wedding.contract.FilterLogs(opts, "NewMessage", fromRule)
	if err != nil {
		return nil, err
	}
	return &WeddingNewMessageIterator{contract: _Wedding.contract, event: "NewMessage", logs: logs, sub: sub}, nil
}

// WatchNewMessage is a free log subscription operation binding the contract event 0x0ff94fec3112de81726d79117e091c7c9d47da8a38c73e8cec7ee350a61898a6.
//
// Solidity: event NewMessage(address indexed from, uint256 timestamp, string message)
func (_Wedding *WeddingFilterer) WatchNewMessage(opts *bind.WatchOpts, sink chan<- *WeddingNewMessage, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Wedding.contract.WatchLogs(opts, "NewMessage", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WeddingNewMessage)
				if err := _Wedding.contract.UnpackLog(event, "NewMessage", log); err != nil {
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

// ParseNewMessage is a log parse operation binding the contract event 0x0ff94fec3112de81726d79117e091c7c9d47da8a38c73e8cec7ee350a61898a6.
//
// Solidity: event NewMessage(address indexed from, uint256 timestamp, string message)
func (_Wedding *WeddingFilterer) ParseNewMessage(log types.Log) (*WeddingNewMessage, error) {
	event := new(WeddingNewMessage)
	if err := _Wedding.contract.UnpackLog(event, "NewMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WeddingNewSignatureIterator is returned from FilterNewSignature and is used to iterate over the raw logs and unpacked data for NewSignature events raised by the Wedding contract.
type WeddingNewSignatureIterator struct {
	Event *WeddingNewSignature // Event containing the contract specifics and raw log

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
func (it *WeddingNewSignatureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WeddingNewSignature)
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
		it.Event = new(WeddingNewSignature)
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
func (it *WeddingNewSignatureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WeddingNewSignatureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WeddingNewSignature represents a NewSignature event raised by the Wedding contract.
type WeddingNewSignature struct {
	From      common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewSignature is a free log retrieval operation binding the contract event 0xfb79a88b762f56bcd11bcc12554f81106e62770680055da0027dc8fe0694e96c.
//
// Solidity: event NewSignature(address indexed from, uint256 timestamp)
func (_Wedding *WeddingFilterer) FilterNewSignature(opts *bind.FilterOpts, from []common.Address) (*WeddingNewSignatureIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Wedding.contract.FilterLogs(opts, "NewSignature", fromRule)
	if err != nil {
		return nil, err
	}
	return &WeddingNewSignatureIterator{contract: _Wedding.contract, event: "NewSignature", logs: logs, sub: sub}, nil
}

// WatchNewSignature is a free log subscription operation binding the contract event 0xfb79a88b762f56bcd11bcc12554f81106e62770680055da0027dc8fe0694e96c.
//
// Solidity: event NewSignature(address indexed from, uint256 timestamp)
func (_Wedding *WeddingFilterer) WatchNewSignature(opts *bind.WatchOpts, sink chan<- *WeddingNewSignature, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Wedding.contract.WatchLogs(opts, "NewSignature", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WeddingNewSignature)
				if err := _Wedding.contract.UnpackLog(event, "NewSignature", log); err != nil {
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

// ParseNewSignature is a log parse operation binding the contract event 0xfb79a88b762f56bcd11bcc12554f81106e62770680055da0027dc8fe0694e96c.
//
// Solidity: event NewSignature(address indexed from, uint256 timestamp)
func (_Wedding *WeddingFilterer) ParseNewSignature(log types.Log) (*WeddingNewSignature, error) {
	event := new(WeddingNewSignature)
	if err := _Wedding.contract.UnpackLog(event, "NewSignature", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
