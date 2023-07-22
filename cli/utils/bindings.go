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
	WalletAddr     common.Address
	Firstname      string
	Description    string
	IpfsPhotoCid   string
	AlreadySaidYes bool
}

// WeddingMessage is an auto generated low-level Go binding around an user-defined struct.
type WeddingMessage struct {
	Sender    common.Address
	Content   string
	Timestamp *big.Int
}

// WeddingMetaData contains all meta data concerning the Wedding contract.
var WeddingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MarriageDone\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"NewMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"NewSignature\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"NotAllowed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getMessages\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structWedding.Message[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSignersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sayYes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"firstname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsPhotoCid\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"alreadySaidYes\",\"type\":\"bool\"}],\"internalType\":\"structWedding.Fiance\",\"name\":\"fiance\",\"type\":\"tuple\"}],\"name\":\"updateFianceDescription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"firstname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsPhotoCid\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"alreadySaidYes\",\"type\":\"bool\"}],\"internalType\":\"structWedding.Fiance\",\"name\":\"fiance\",\"type\":\"tuple\"}],\"name\":\"updateFiancePhoto\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526040518060a0016040528073e42749aa40040929cd8db5ac5200df12f9b5254073ffffffffffffffffffffffffffffffffffffffff1681526020016040518060400160405280600681526020017f52616a656576000000000000000000000000000000000000000000000000000081525081526020016040518060600160405280602681526020016200197b6026913981526020016040518060600160405280602e81526020016200194d602e91398152602001600015158152506000808201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190816200011e91906200056d565b5060408201518160020190816200013691906200056d565b5060608201518160030190816200014e91906200056d565b5060808201518160040160006101000a81548160ff02191690831515021790555050506040518060a001604052807323bc9eb2343125b91a7d5c59038e40e0ae45155173ffffffffffffffffffffffffffffffffffffffff1681526020016040518060400160405280600581526020017f4d617269610000000000000000000000000000000000000000000000000000008152508152602001604051806060016040528060298152602001620019246029913981526020016040518060600160405280602e8152602001620018f6602e9139815260200160001515815250600560008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190816200028c91906200056d565b506040820151816002019081620002a491906200056d565b506060820151816003019081620002bc91906200056d565b5060808201518160040160006101000a81548160ff0219169083151502179055505050348015620002ec57600080fd5b5062000654565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200037557607f821691505b6020821081036200038b576200038a6200032d565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620003f57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620003b6565b620004018683620003b6565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006200044e62000448620004428462000419565b62000423565b62000419565b9050919050565b6000819050919050565b6200046a836200042d565b62000482620004798262000455565b848454620003c3565b825550505050565b600090565b620004996200048a565b620004a68184846200045f565b505050565b5b81811015620004ce57620004c26000826200048f565b600181019050620004ac565b5050565b601f8211156200051d57620004e78162000391565b620004f284620003a6565b8101602085101562000502578190505b6200051a6200051185620003a6565b830182620004ab565b50505b505050565b600082821c905092915050565b6000620005426000198460080262000522565b1980831691505092915050565b60006200055d83836200052f565b9150826002028217905092915050565b6200057882620002f3565b67ffffffffffffffff811115620005945762000593620002fe565b5b620005a082546200035c565b620005ad828285620004d2565b600060209050601f831160018114620005e55760008415620005d0578287015190505b620005dc85826200054f565b8655506200064c565b601f198416620005f58662000391565b60005b828110156200061f57848901518255600182019150602085019450602081019050620005f8565b868310156200063f57848901516200063b601f8916826200052f565b8355505b6001600288020188555050505b505050505050565b61129280620006646000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80631a77837514610067578063469c8110146100835780635ff6cbf31461009f5780637fda2d25146100bd5780639bee3365146100d9578063a0c1deb4146100e3575b600080fd5b610081600480360381019061007c9190610b15565b610101565b005b61009d60048036038101906100989190610bbe565b610104565b005b6100a7610263565b6040516100b49190610dc4565b60405180910390f35b6100d760048036038101906100d29190610b15565b6103b4565b005b6100e16103b7565b005b6100eb61082e565b6040516100f89190610df5565b60405180910390f35b50565b600b60405180606001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200184848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200142815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019081610200919061101c565b506040820151816002015550503373ffffffffffffffffffffffffffffffffffffffff167f0ff94fec3112de81726d79117e091c7c9d47da8a38c73e8cec7ee350a61898a64284846040516102579392919061112c565b60405180910390a25050565b6060600b805480602002602001604051908101604052809291908181526020016000905b828210156103ab57838290600052602060002090600302016040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461031090610e3f565b80601f016020809104026020016040519081016040528092919081815260200182805461033c90610e3f565b80156103895780601f1061035e57610100808354040283529160200191610389565b820191906000526020600020905b81548152906001019060200180831161036c57829003601f168201915b5050505050815260200160028201548152505081526020019060010190610287565b50505050905090565b50565b6000800160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16036105a5576001600060040160006101000a81548160ff021916908315150217905550600a600090806001815401808255809150506001900390600052602060002090600502016000909190919091506000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600182018160010190816104d49190611174565b50600282018160020190816104e99190611174565b50600382018160030190816104fe9190611174565b506004820160009054906101000a900460ff168160040160006101000a81548160ff02191690831515021790555050506000800160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ffb79a88b762f56bcd11bcc12554f81106e62770680055da0027dc8fe0694e96c426040516105989190610df5565b60405180910390a26107e5565b600560000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1603610795576001600560040160006101000a81548160ff021916908315150217905550600a600590806001815401808255809150506001900390600052602060002090600502016000909190919091506000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600182018160010190816106c39190611174565b50600282018160020190816106d89190611174565b50600382018160030190816106ed9190611174565b506004820160009054906101000a900460ff168160040160006101000a81548160ff0219169083151502179055505050600560000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ffb79a88b762f56bcd11bcc12554f81106e62770680055da0027dc8fe0694e96c426040516107889190610df5565b60405180910390a26107e4565b3373ffffffffffffffffffffffffffffffffffffffff167f8626cc03991f4dbbb34a0077f5b4d9d6fbeb0aecdd7045d78597298e05f17e1c426040516107db9190610df5565b60405180910390a25b5b60026107ef61082e565b0361082c577f9ed866c38cbf8638bead8e7c73fbe96b0277adce255685cae1c4e05260472d42426040516108239190610df5565b60405180910390a15b565b6000600a80549050905090565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61089d82610854565b810181811067ffffffffffffffff821117156108bc576108bb610865565b5b80604052505050565b60006108cf61083b565b90506108db8282610894565b919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610910826108e5565b9050919050565b61092081610905565b811461092b57600080fd5b50565b60008135905061093d81610917565b92915050565b600080fd5b600080fd5b600067ffffffffffffffff82111561096857610967610865565b5b61097182610854565b9050602081019050919050565b82818337600083830152505050565b60006109a061099b8461094d565b6108c5565b9050828152602081018484840111156109bc576109bb610948565b5b6109c784828561097e565b509392505050565b600082601f8301126109e4576109e3610943565b5b81356109f484826020860161098d565b91505092915050565b60008115159050919050565b610a12816109fd565b8114610a1d57600080fd5b50565b600081359050610a2f81610a09565b92915050565b600060a08284031215610a4b57610a4a61084f565b5b610a5560a06108c5565b90506000610a658482850161092e565b600083015250602082013567ffffffffffffffff811115610a8957610a886108e0565b5b610a95848285016109cf565b602083015250604082013567ffffffffffffffff811115610ab957610ab86108e0565b5b610ac5848285016109cf565b604083015250606082013567ffffffffffffffff811115610ae957610ae86108e0565b5b610af5848285016109cf565b6060830152506080610b0984828501610a20565b60808301525092915050565b600060208284031215610b2b57610b2a610845565b5b600082013567ffffffffffffffff811115610b4957610b4861084a565b5b610b5584828501610a35565b91505092915050565b600080fd5b600080fd5b60008083601f840112610b7e57610b7d610943565b5b8235905067ffffffffffffffff811115610b9b57610b9a610b5e565b5b602083019150836001820283011115610bb757610bb6610b63565b5b9250929050565b60008060208385031215610bd557610bd4610845565b5b600083013567ffffffffffffffff811115610bf357610bf261084a565b5b610bff85828601610b68565b92509250509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610c4081610905565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610c80578082015181840152602081019050610c65565b60008484015250505050565b6000610c9782610c46565b610ca18185610c51565b9350610cb1818560208601610c62565b610cba81610854565b840191505092915050565b6000819050919050565b610cd881610cc5565b82525050565b6000606083016000830151610cf66000860182610c37565b5060208301518482036020860152610d0e8282610c8c565b9150506040830151610d236040860182610ccf565b508091505092915050565b6000610d3a8383610cde565b905092915050565b6000602082019050919050565b6000610d5a82610c0b565b610d648185610c16565b935083602082028501610d7685610c27565b8060005b85811015610db25784840389528151610d938582610d2e565b9450610d9e83610d42565b925060208a01995050600181019050610d7a565b50829750879550505050505092915050565b60006020820190508181036000830152610dde8184610d4f565b905092915050565b610def81610cc5565b82525050565b6000602082019050610e0a6000830184610de6565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610e5757607f821691505b602082108103610e6a57610e69610e10565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610ed27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610e95565b610edc8683610e95565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610f19610f14610f0f84610cc5565b610ef4565b610cc5565b9050919050565b6000819050919050565b610f3383610efe565b610f47610f3f82610f20565b848454610ea2565b825550505050565b600090565b610f5c610f4f565b610f67818484610f2a565b505050565b5b81811015610f8b57610f80600082610f54565b600181019050610f6d565b5050565b601f821115610fd057610fa181610e70565b610faa84610e85565b81016020851015610fb9578190505b610fcd610fc585610e85565b830182610f6c565b50505b505050565b600082821c905092915050565b6000610ff360001984600802610fd5565b1980831691505092915050565b600061100c8383610fe2565b9150826002028217905092915050565b61102582610c46565b67ffffffffffffffff81111561103e5761103d610865565b5b6110488254610e3f565b611053828285610f8f565b600060209050601f8311600181146110865760008415611074578287015190505b61107e8582611000565b8655506110e6565b601f19841661109486610e70565b60005b828110156110bc57848901518255600182019150602085019450602081019050611097565b868310156110d957848901516110d5601f891682610fe2565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b600061110b83856110ee565b935061111883858461097e565b61112183610854565b840190509392505050565b60006040820190506111416000830186610de6565b81810360208301526111548184866110ff565b9050949350505050565b60008154905061116d81610e3f565b9050919050565b81810361118257505061125a565b61118b8261115e565b67ffffffffffffffff8111156111a4576111a3610865565b5b6111ae8254610e3f565b6111b9828285610f8f565b6000601f8311600181146111e857600084156111d6578287015490505b6111e08582611000565b865550611253565b601f1984166111f687610e70565b965061120186610e70565b60005b8281101561122957848901548255600182019150600185019450602081019050611204565b868310156112465784890154611242601f891682610fe2565b8355505b6001600288020188555050505b5050505050505b56fea2646970667358221220d899e54a41c854a1f9e90b20e675e0f7b1215683efe79282ab2e24b0a92ecbbe64736f6c63430008100033516d58346143486a637571704479734a6b3657444c6a3639735954744d4768795479704c753275774c474b4e59675468697320697320746865206d6f7374206578636974696e672064617920696e206d79206c69666521516d5351547a344378486d724e6170546a595a635964543436714a526a4a4e323958466a32507472764670364c7648656c6c6f2074686572652c2069276d2061626f75742074686520676574206d617272696564",
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

// GetSignersCount is a free data retrieval call binding the contract method 0xa0c1deb4.
//
// Solidity: function getSignersCount() view returns(uint256 count)
func (_Wedding *WeddingCaller) GetSignersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Wedding.contract.Call(opts, &out, "getSignersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSignersCount is a free data retrieval call binding the contract method 0xa0c1deb4.
//
// Solidity: function getSignersCount() view returns(uint256 count)
func (_Wedding *WeddingSession) GetSignersCount() (*big.Int, error) {
	return _Wedding.Contract.GetSignersCount(&_Wedding.CallOpts)
}

// GetSignersCount is a free data retrieval call binding the contract method 0xa0c1deb4.
//
// Solidity: function getSignersCount() view returns(uint256 count)
func (_Wedding *WeddingCallerSession) GetSignersCount() (*big.Int, error) {
	return _Wedding.Contract.GetSignersCount(&_Wedding.CallOpts)
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

// UpdateFianceDescription is a paid mutator transaction binding the contract method 0x7fda2d25.
//
// Solidity: function updateFianceDescription((address,string,string,string,bool) fiance) returns()
func (_Wedding *WeddingTransactor) UpdateFianceDescription(opts *bind.TransactOpts, fiance WeddingFiance) (*types.Transaction, error) {
	return _Wedding.contract.Transact(opts, "updateFianceDescription", fiance)
}

// UpdateFianceDescription is a paid mutator transaction binding the contract method 0x7fda2d25.
//
// Solidity: function updateFianceDescription((address,string,string,string,bool) fiance) returns()
func (_Wedding *WeddingSession) UpdateFianceDescription(fiance WeddingFiance) (*types.Transaction, error) {
	return _Wedding.Contract.UpdateFianceDescription(&_Wedding.TransactOpts, fiance)
}

// UpdateFianceDescription is a paid mutator transaction binding the contract method 0x7fda2d25.
//
// Solidity: function updateFianceDescription((address,string,string,string,bool) fiance) returns()
func (_Wedding *WeddingTransactorSession) UpdateFianceDescription(fiance WeddingFiance) (*types.Transaction, error) {
	return _Wedding.Contract.UpdateFianceDescription(&_Wedding.TransactOpts, fiance)
}

// UpdateFiancePhoto is a paid mutator transaction binding the contract method 0x1a778375.
//
// Solidity: function updateFiancePhoto((address,string,string,string,bool) fiance) returns()
func (_Wedding *WeddingTransactor) UpdateFiancePhoto(opts *bind.TransactOpts, fiance WeddingFiance) (*types.Transaction, error) {
	return _Wedding.contract.Transact(opts, "updateFiancePhoto", fiance)
}

// UpdateFiancePhoto is a paid mutator transaction binding the contract method 0x1a778375.
//
// Solidity: function updateFiancePhoto((address,string,string,string,bool) fiance) returns()
func (_Wedding *WeddingSession) UpdateFiancePhoto(fiance WeddingFiance) (*types.Transaction, error) {
	return _Wedding.Contract.UpdateFiancePhoto(&_Wedding.TransactOpts, fiance)
}

// UpdateFiancePhoto is a paid mutator transaction binding the contract method 0x1a778375.
//
// Solidity: function updateFiancePhoto((address,string,string,string,bool) fiance) returns()
func (_Wedding *WeddingTransactorSession) UpdateFiancePhoto(fiance WeddingFiance) (*types.Transaction, error) {
	return _Wedding.Contract.UpdateFiancePhoto(&_Wedding.TransactOpts, fiance)
}

// WeddingMarriageDoneIterator is returned from FilterMarriageDone and is used to iterate over the raw logs and unpacked data for MarriageDone events raised by the Wedding contract.
type WeddingMarriageDoneIterator struct {
	Event *WeddingMarriageDone // Event containing the contract specifics and raw log

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
func (it *WeddingMarriageDoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WeddingMarriageDone)
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
		it.Event = new(WeddingMarriageDone)
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
func (it *WeddingMarriageDoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WeddingMarriageDoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WeddingMarriageDone represents a MarriageDone event raised by the Wedding contract.
type WeddingMarriageDone struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMarriageDone is a free log retrieval operation binding the contract event 0x9ed866c38cbf8638bead8e7c73fbe96b0277adce255685cae1c4e05260472d42.
//
// Solidity: event MarriageDone(uint256 timestamp)
func (_Wedding *WeddingFilterer) FilterMarriageDone(opts *bind.FilterOpts) (*WeddingMarriageDoneIterator, error) {

	logs, sub, err := _Wedding.contract.FilterLogs(opts, "MarriageDone")
	if err != nil {
		return nil, err
	}
	return &WeddingMarriageDoneIterator{contract: _Wedding.contract, event: "MarriageDone", logs: logs, sub: sub}, nil
}

// WatchMarriageDone is a free log subscription operation binding the contract event 0x9ed866c38cbf8638bead8e7c73fbe96b0277adce255685cae1c4e05260472d42.
//
// Solidity: event MarriageDone(uint256 timestamp)
func (_Wedding *WeddingFilterer) WatchMarriageDone(opts *bind.WatchOpts, sink chan<- *WeddingMarriageDone) (event.Subscription, error) {

	logs, sub, err := _Wedding.contract.WatchLogs(opts, "MarriageDone")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WeddingMarriageDone)
				if err := _Wedding.contract.UnpackLog(event, "MarriageDone", log); err != nil {
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

// ParseMarriageDone is a log parse operation binding the contract event 0x9ed866c38cbf8638bead8e7c73fbe96b0277adce255685cae1c4e05260472d42.
//
// Solidity: event MarriageDone(uint256 timestamp)
func (_Wedding *WeddingFilterer) ParseMarriageDone(log types.Log) (*WeddingMarriageDone, error) {
	event := new(WeddingMarriageDone)
	if err := _Wedding.contract.UnpackLog(event, "MarriageDone", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// WeddingNotAllowedIterator is returned from FilterNotAllowed and is used to iterate over the raw logs and unpacked data for NotAllowed events raised by the Wedding contract.
type WeddingNotAllowedIterator struct {
	Event *WeddingNotAllowed // Event containing the contract specifics and raw log

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
func (it *WeddingNotAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WeddingNotAllowed)
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
		it.Event = new(WeddingNotAllowed)
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
func (it *WeddingNotAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WeddingNotAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WeddingNotAllowed represents a NotAllowed event raised by the Wedding contract.
type WeddingNotAllowed struct {
	From      common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNotAllowed is a free log retrieval operation binding the contract event 0x8626cc03991f4dbbb34a0077f5b4d9d6fbeb0aecdd7045d78597298e05f17e1c.
//
// Solidity: event NotAllowed(address indexed from, uint256 timestamp)
func (_Wedding *WeddingFilterer) FilterNotAllowed(opts *bind.FilterOpts, from []common.Address) (*WeddingNotAllowedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Wedding.contract.FilterLogs(opts, "NotAllowed", fromRule)
	if err != nil {
		return nil, err
	}
	return &WeddingNotAllowedIterator{contract: _Wedding.contract, event: "NotAllowed", logs: logs, sub: sub}, nil
}

// WatchNotAllowed is a free log subscription operation binding the contract event 0x8626cc03991f4dbbb34a0077f5b4d9d6fbeb0aecdd7045d78597298e05f17e1c.
//
// Solidity: event NotAllowed(address indexed from, uint256 timestamp)
func (_Wedding *WeddingFilterer) WatchNotAllowed(opts *bind.WatchOpts, sink chan<- *WeddingNotAllowed, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Wedding.contract.WatchLogs(opts, "NotAllowed", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WeddingNotAllowed)
				if err := _Wedding.contract.UnpackLog(event, "NotAllowed", log); err != nil {
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

// ParseNotAllowed is a log parse operation binding the contract event 0x8626cc03991f4dbbb34a0077f5b4d9d6fbeb0aecdd7045d78597298e05f17e1c.
//
// Solidity: event NotAllowed(address indexed from, uint256 timestamp)
func (_Wedding *WeddingFilterer) ParseNotAllowed(log types.Log) (*WeddingNotAllowed, error) {
	event := new(WeddingNotAllowed)
	if err := _Wedding.contract.UnpackLog(event, "NotAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
