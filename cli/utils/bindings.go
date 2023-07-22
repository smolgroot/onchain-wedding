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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MarriageDone\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"NewMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"NewSignature\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"NotAllowed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getFiance1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFiance2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMessages\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structWedding.Message[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSignersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sayYes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"firstname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsPhotoCid\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"alreadySaidYes\",\"type\":\"bool\"}],\"internalType\":\"structWedding.Fiance\",\"name\":\"fiance\",\"type\":\"tuple\"}],\"name\":\"updateFianceDescription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"firstname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsPhotoCid\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"alreadySaidYes\",\"type\":\"bool\"}],\"internalType\":\"structWedding.Fiance\",\"name\":\"fiance\",\"type\":\"tuple\"}],\"name\":\"updateFiancePhoto\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526040518060a0016040528073e42749aa40040929cd8db5ac5200df12f9b5254073ffffffffffffffffffffffffffffffffffffffff1681526020016040518060400160405280600681526020017f52616a6565760000000000000000000000000000000000000000000000000000815250815260200160405180606001604052806026815260200162001d4d6026913981526020016040518060600160405280602e815260200162001d1f602e9139815260200160001515815250600160008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190816200011f91906200056e565b5060408201518160020190816200013791906200056e565b5060608201518160030190816200014f91906200056e565b5060808201518160040160006101000a81548160ff02191690831515021790555050506040518060a001604052807323bc9eb2343125b91a7d5c59038e40e0ae45155173ffffffffffffffffffffffffffffffffffffffff1681526020016040518060400160405280600581526020017f4d61726961000000000000000000000000000000000000000000000000000000815250815260200160405180606001604052806029815260200162001cf66029913981526020016040518060600160405280602e815260200162001cc8602e9139815260200160001515815250600660008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190816200028d91906200056e565b506040820151816002019081620002a591906200056e565b506060820151816003019081620002bd91906200056e565b5060808201518160040160006101000a81548160ff0219169083151502179055505050348015620002ed57600080fd5b5062000655565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200037657607f821691505b6020821081036200038c576200038b6200032e565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620003f67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620003b7565b620004028683620003b7565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006200044f6200044962000443846200041a565b62000424565b6200041a565b9050919050565b6000819050919050565b6200046b836200042e565b620004836200047a8262000456565b848454620003c4565b825550505050565b600090565b6200049a6200048b565b620004a781848462000460565b505050565b5b81811015620004cf57620004c360008262000490565b600181019050620004ad565b5050565b601f8211156200051e57620004e88162000392565b620004f384620003a7565b8101602085101562000503578190505b6200051b6200051285620003a7565b830182620004ac565b50505b505050565b600082821c905092915050565b6000620005436000198460080262000523565b1980831691505092915050565b60006200055e838362000530565b9150826002028217905092915050565b6200057982620002f4565b67ffffffffffffffff811115620005955762000594620002ff565b5b620005a182546200035d565b620005ae828285620004d3565b600060209050601f831160018114620005e65760008415620005d1578287015190505b620005dd858262000550565b8655506200064d565b601f198416620005f68662000392565b60005b828110156200062057848901518255600182019150602085019450602081019050620005f9565b868310156200064057848901516200063c601f89168262000530565b8355505b6001600288020188555050505b505050505050565b61166380620006656000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80639bee3365116100665780639bee336514610134578063a0c1deb41461013e578063a76b326c1461015c578063bed85a731461017a578063e0cfc4fc146101985761009e565b80631a778375146100a35780631cfa141e146100bf578063469c8110146100de5780635ff6cbf3146100fa5780637fda2d2514610118575b600080fd5b6100bd60048036038101906100b89190610e24565b6101b7565b005b6100c76101ba565b6040516100d5929190610efb565b60405180910390f35b6100f860048036038101906100f39190610f8b565b61027d565b005b6101026103dc565b60405161010f919061115c565b60405180910390f35b610132600480360381019061012d9190610e24565b61052d565b005b61013c610530565b005b610146610a2e565b604051610153919061118d565b60405180910390f35b610164610a3b565b60405161017191906111a8565b60405180910390f35b610182610a64565b60405161018f91906111a8565b60405180910390f35b6101a0610a88565b6040516101ae929190610efb565b60405180910390f35b50565b60006060600660000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660066001018080546101f5906111f2565b80601f0160208091040260200160405190810160405280929190818152602001828054610221906111f2565b801561026e5780601f106102435761010080835404028352916020019161026e565b820191906000526020600020905b81548152906001019060200180831161025157829003601f168201915b50505050509050915091509091565b600c60405180606001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200184848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200142815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101908161037991906113cf565b506040820151816002015550503373ffffffffffffffffffffffffffffffffffffffff167f0ff94fec3112de81726d79117e091c7c9d47da8a38c73e8cec7ee350a61898a64284846040516103d0939291906114ce565b60405180910390a25050565b6060600c805480602002602001604051908101604052809291908181526020016000905b8282101561052457838290600052602060002090600302016040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054610489906111f2565b80601f01602080910402602001604051908101604052809291908181526020018280546104b5906111f2565b80156105025780601f106104d757610100808354040283529160200191610502565b820191906000526020600020905b8154815290600101906020018083116104e557829003601f168201915b5050505050815260200160028201548152505081526020019060010190610400565b50505050905090565b50565b600160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff160361071f5760018060040160006101000a81548160ff021916908315150217905550600b600190806001815401808255809150506001900390600052602060002090600502016000909190919091506000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001820181600101908161064d9190611516565b50600282018160020190816106629190611516565b50600382018160030190816106779190611516565b506004820160009054906101000a900460ff168160040160006101000a81548160ff0219169083151502179055505050600160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ffb79a88b762f56bcd11bcc12554f81106e62770680055da0027dc8fe0694e96c42604051610712919061118d565b60405180910390a261095f565b600660000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff160361090f576001600660040160006101000a81548160ff021916908315150217905550600b600690806001815401808255809150506001900390600052602060002090600502016000909190919091506000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001820181600101908161083d9190611516565b50600282018160020190816108529190611516565b50600382018160030190816108679190611516565b506004820160009054906101000a900460ff168160040160006101000a81548160ff0219169083151502179055505050600660000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ffb79a88b762f56bcd11bcc12554f81106e62770680055da0027dc8fe0694e96c42604051610902919061118d565b60405180910390a261095e565b3373ffffffffffffffffffffffffffffffffffffffff167f8626cc03991f4dbbb34a0077f5b4d9d6fbeb0aecdd7045d78597298e05f17e1c42604051610955919061118d565b60405180910390a25b5b6002610969610a2e565b03610a2c57600b600181548110610983576109826115fe565b5b906000526020600020906005020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f9ed866c38cbf8638bead8e7c73fbe96b0277adce255685cae1c4e05260472d4242604051610a23919061118d565b60405180910390a15b565b6000600b80549050905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006060600160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660018001808054610ac2906111f2565b80601f0160208091040260200160405190810160405280929190818152602001828054610aee906111f2565b8015610b3b5780601f10610b1057610100808354040283529160200191610b3b565b820191906000526020600020905b815481529060010190602001808311610b1e57829003601f168201915b50505050509050915091509091565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610bac82610b63565b810181811067ffffffffffffffff82111715610bcb57610bca610b74565b5b80604052505050565b6000610bde610b4a565b9050610bea8282610ba3565b919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610c1f82610bf4565b9050919050565b610c2f81610c14565b8114610c3a57600080fd5b50565b600081359050610c4c81610c26565b92915050565b600080fd5b600080fd5b600067ffffffffffffffff821115610c7757610c76610b74565b5b610c8082610b63565b9050602081019050919050565b82818337600083830152505050565b6000610caf610caa84610c5c565b610bd4565b905082815260208101848484011115610ccb57610cca610c57565b5b610cd6848285610c8d565b509392505050565b600082601f830112610cf357610cf2610c52565b5b8135610d03848260208601610c9c565b91505092915050565b60008115159050919050565b610d2181610d0c565b8114610d2c57600080fd5b50565b600081359050610d3e81610d18565b92915050565b600060a08284031215610d5a57610d59610b5e565b5b610d6460a0610bd4565b90506000610d7484828501610c3d565b600083015250602082013567ffffffffffffffff811115610d9857610d97610bef565b5b610da484828501610cde565b602083015250604082013567ffffffffffffffff811115610dc857610dc7610bef565b5b610dd484828501610cde565b604083015250606082013567ffffffffffffffff811115610df857610df7610bef565b5b610e0484828501610cde565b6060830152506080610e1884828501610d2f565b60808301525092915050565b600060208284031215610e3a57610e39610b54565b5b600082013567ffffffffffffffff811115610e5857610e57610b59565b5b610e6484828501610d44565b91505092915050565b610e7681610c14565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610eb6578082015181840152602081019050610e9b565b60008484015250505050565b6000610ecd82610e7c565b610ed78185610e87565b9350610ee7818560208601610e98565b610ef081610b63565b840191505092915050565b6000604082019050610f106000830185610e6d565b8181036020830152610f228184610ec2565b90509392505050565b600080fd5b600080fd5b60008083601f840112610f4b57610f4a610c52565b5b8235905067ffffffffffffffff811115610f6857610f67610f2b565b5b602083019150836001820283011115610f8457610f83610f30565b5b9250929050565b60008060208385031215610fa257610fa1610b54565b5b600083013567ffffffffffffffff811115610fc057610fbf610b59565b5b610fcc85828601610f35565b92509250509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61100d81610c14565b82525050565b600082825260208201905092915050565b600061102f82610e7c565b6110398185611013565b9350611049818560208601610e98565b61105281610b63565b840191505092915050565b6000819050919050565b6110708161105d565b82525050565b600060608301600083015161108e6000860182611004565b50602083015184820360208601526110a68282611024565b91505060408301516110bb6040860182611067565b508091505092915050565b60006110d28383611076565b905092915050565b6000602082019050919050565b60006110f282610fd8565b6110fc8185610fe3565b93508360208202850161110e85610ff4565b8060005b8581101561114a578484038952815161112b85826110c6565b9450611136836110da565b925060208a01995050600181019050611112565b50829750879550505050505092915050565b6000602082019050818103600083015261117681846110e7565b905092915050565b6111878161105d565b82525050565b60006020820190506111a2600083018461117e565b92915050565b60006020820190506111bd6000830184610e6d565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061120a57607f821691505b60208210810361121d5761121c6111c3565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026112857fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611248565b61128f8683611248565b95508019841693508086168417925050509392505050565b6000819050919050565b60006112cc6112c76112c28461105d565b6112a7565b61105d565b9050919050565b6000819050919050565b6112e6836112b1565b6112fa6112f2826112d3565b848454611255565b825550505050565b600090565b61130f611302565b61131a8184846112dd565b505050565b5b8181101561133e57611333600082611307565b600181019050611320565b5050565b601f8211156113835761135481611223565b61135d84611238565b8101602085101561136c578190505b61138061137885611238565b83018261131f565b50505b505050565b600082821c905092915050565b60006113a660001984600802611388565b1980831691505092915050565b60006113bf8383611395565b9150826002028217905092915050565b6113d882610e7c565b67ffffffffffffffff8111156113f1576113f0610b74565b5b6113fb82546111f2565b611406828285611342565b600060209050601f8311600181146114395760008415611427578287015190505b61143185826113b3565b865550611499565b601f19841661144786611223565b60005b8281101561146f5784890151825560018201915060208501945060208101905061144a565b8683101561148c5784890151611488601f891682611395565b8355505b6001600288020188555050505b505050505050565b60006114ad8385610e87565b93506114ba838584610c8d565b6114c383610b63565b840190509392505050565b60006040820190506114e3600083018661117e565b81810360208301526114f68184866114a1565b9050949350505050565b60008154905061150f816111f2565b9050919050565b8181036115245750506115fc565b61152d82611500565b67ffffffffffffffff81111561154657611545610b74565b5b61155082546111f2565b61155b828285611342565b6000601f83116001811461158a5760008415611578578287015490505b61158285826113b3565b8655506115f5565b601f19841661159887611223565b96506115a386611223565b60005b828110156115cb578489015482556001820191506001850194506020810190506115a6565b868310156115e857848901546115e4601f891682611395565b8355505b6001600288020188555050505b5050505050505b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea2646970667358221220aebdb6f7bb8abecf9dfb7769e690e3c2aec1872669fb72146229ea736caf7afc64736f6c63430008100033516d58346143486a637571704479734a6b3657444c6a3639735954744d4768795479704c753275774c474b4e59675468697320697320746865206d6f7374206578636974696e672064617920696e206d79206c69666521516d5351547a344378486d724e6170546a595a635964543436714a526a4a4e323958466a32507472764670364c7648656c6c6f2074686572652c2069276d2061626f75742074686520676574206d617272696564",
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

// GetFiance1 is a free data retrieval call binding the contract method 0xe0cfc4fc.
//
// Solidity: function getFiance1() view returns(address, string)
func (_Wedding *WeddingCaller) GetFiance1(opts *bind.CallOpts) (common.Address, string, error) {
	var out []interface{}
	err := _Wedding.contract.Call(opts, &out, "getFiance1")

	if err != nil {
		return *new(common.Address), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetFiance1 is a free data retrieval call binding the contract method 0xe0cfc4fc.
//
// Solidity: function getFiance1() view returns(address, string)
func (_Wedding *WeddingSession) GetFiance1() (common.Address, string, error) {
	return _Wedding.Contract.GetFiance1(&_Wedding.CallOpts)
}

// GetFiance1 is a free data retrieval call binding the contract method 0xe0cfc4fc.
//
// Solidity: function getFiance1() view returns(address, string)
func (_Wedding *WeddingCallerSession) GetFiance1() (common.Address, string, error) {
	return _Wedding.Contract.GetFiance1(&_Wedding.CallOpts)
}

// GetFiance2 is a free data retrieval call binding the contract method 0x1cfa141e.
//
// Solidity: function getFiance2() view returns(address, string)
func (_Wedding *WeddingCaller) GetFiance2(opts *bind.CallOpts) (common.Address, string, error) {
	var out []interface{}
	err := _Wedding.contract.Call(opts, &out, "getFiance2")

	if err != nil {
		return *new(common.Address), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetFiance2 is a free data retrieval call binding the contract method 0x1cfa141e.
//
// Solidity: function getFiance2() view returns(address, string)
func (_Wedding *WeddingSession) GetFiance2() (common.Address, string, error) {
	return _Wedding.Contract.GetFiance2(&_Wedding.CallOpts)
}

// GetFiance2 is a free data retrieval call binding the contract method 0x1cfa141e.
//
// Solidity: function getFiance2() view returns(address, string)
func (_Wedding *WeddingCallerSession) GetFiance2() (common.Address, string, error) {
	return _Wedding.Contract.GetFiance2(&_Wedding.CallOpts)
}

// GetLastSigner is a free data retrieval call binding the contract method 0xa76b326c.
//
// Solidity: function getLastSigner() view returns(address)
func (_Wedding *WeddingCaller) GetLastSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Wedding.contract.Call(opts, &out, "getLastSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLastSigner is a free data retrieval call binding the contract method 0xa76b326c.
//
// Solidity: function getLastSigner() view returns(address)
func (_Wedding *WeddingSession) GetLastSigner() (common.Address, error) {
	return _Wedding.Contract.GetLastSigner(&_Wedding.CallOpts)
}

// GetLastSigner is a free data retrieval call binding the contract method 0xa76b326c.
//
// Solidity: function getLastSigner() view returns(address)
func (_Wedding *WeddingCallerSession) GetLastSigner() (common.Address, error) {
	return _Wedding.Contract.GetLastSigner(&_Wedding.CallOpts)
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

// LastSigner is a free data retrieval call binding the contract method 0xbed85a73.
//
// Solidity: function lastSigner() view returns(address)
func (_Wedding *WeddingCaller) LastSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Wedding.contract.Call(opts, &out, "lastSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastSigner is a free data retrieval call binding the contract method 0xbed85a73.
//
// Solidity: function lastSigner() view returns(address)
func (_Wedding *WeddingSession) LastSigner() (common.Address, error) {
	return _Wedding.Contract.LastSigner(&_Wedding.CallOpts)
}

// LastSigner is a free data retrieval call binding the contract method 0xbed85a73.
//
// Solidity: function lastSigner() view returns(address)
func (_Wedding *WeddingCallerSession) LastSigner() (common.Address, error) {
	return _Wedding.Contract.LastSigner(&_Wedding.CallOpts)
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
