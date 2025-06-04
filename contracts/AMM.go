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
	_ = abi.ConvertType
)

// AmmMetaData contains all meta data concerning the Amm contract.
var AmmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapExactETHForToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"swapExactTokenForETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051610bf9380380610bf9833981810160405281019061003191906100d4565b805f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506100ff565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100a38261007a565b9050919050565b6100b381610099565b81146100bd575f5ffd5b50565b5f815190506100ce816100aa565b92915050565b5f602082840312156100e9576100e8610076565b5b5f6100f6848285016100c0565b91505092915050565b610aed8061010c5f395ff3fe60806040526004361061006f575f3560e01c8063b62aafd51161004d578063b62aafd5146100f5578063dcb4b4a11461011d578063f4325d6714610127578063fc0c546a146101515761006f565b8063054d50d414610073578063100af203146100af57806351c6590a146100d9575b5f5ffd5b34801561007e575f5ffd5b506100996004803603810190610094919061061f565b61017b565b6040516100a6919061067e565b60405180910390f35b3480156100ba575f5ffd5b506100c36101d0565b6040516100d0919061067e565b60405180910390f35b6100f360048036038101906100ee9190610697565b6101d6565b005b348015610100575f5ffd5b5061011b60048036038101906101169190610697565b6102e9565b005b610125610456565b005b348015610132575f5ffd5b5061013b6105be565b604051610148919061067e565b60405180910390f35b34801561015c575f5ffd5b506101656105c4565b6040516101729190610701565b60405180910390f35b5f5f6103e58561018b9190610747565b90505f838261019a9190610747565b90505f826103e8876101ac9190610747565b6101b69190610788565b905080826101c491906107e8565b93505050509392505050565b60025481565b5f3411610218576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161020f90610872565b60405180910390fd5b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b815260040161027593929190610890565b6020604051808303815f875af1158015610291573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102b591906108fa565b508060015f8282546102c79190610788565b925050819055503460025f8282546102df9190610788565b9250508190555050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b815260040161034693929190610890565b6020604051808303815f875af1158015610362573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061038691906108fa565b505f6103978260015460025461017b565b90506002548111156103de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d59061096f565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050158015610421573d5f5f3e3d5ffd5b508160015f8282546104339190610788565b925050819055508060025f82825461044b919061098d565b925050819055505050565b5f3411610498576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048f90610a0a565b60405180910390fd5b5f6104a83460025460015461017b565b90506001548111156104ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104e690610a72565b60405180910390fd5b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b815260040161054a929190610a90565b6020604051808303815f875af1158015610566573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061058a91906108fa565b503460025f82825461059c9190610788565b925050819055508060015f8282546105b4919061098d565b9250508190555050565b60015481565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5ffd5b5f819050919050565b6105fe816105ec565b8114610608575f5ffd5b50565b5f81359050610619816105f5565b92915050565b5f5f5f60608486031215610636576106356105e8565b5b5f6106438682870161060b565b93505060206106548682870161060b565b92505060406106658682870161060b565b9150509250925092565b610678816105ec565b82525050565b5f6020820190506106915f83018461066f565b92915050565b5f602082840312156106ac576106ab6105e8565b5b5f6106b98482850161060b565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6106eb826106c2565b9050919050565b6106fb816106e1565b82525050565b5f6020820190506107145f8301846106f2565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610751826105ec565b915061075c836105ec565b925082820261076a816105ec565b915082820484148315176107815761078061071a565b5b5092915050565b5f610792826105ec565b915061079d836105ec565b92508282019050808211156107b5576107b461071a565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6107f2826105ec565b91506107fd836105ec565b92508261080d5761080c6107bb565b5b828204905092915050565b5f82825260208201905092915050565b7f4e65656420746f2073656e6420455448000000000000000000000000000000005f82015250565b5f61085c601083610818565b915061086782610828565b602082019050919050565b5f6020820190508181035f83015261088981610850565b9050919050565b5f6060820190506108a35f8301866106f2565b6108b060208301856106f2565b6108bd604083018461066f565b949350505050565b5f8115159050919050565b6108d9816108c5565b81146108e3575f5ffd5b50565b5f815190506108f4816108d0565b92915050565b5f6020828403121561090f5761090e6105e8565b5b5f61091c848285016108e6565b91505092915050565b7f4e6f7420656e6f7567682045544820696e2072657365727665000000000000005f82015250565b5f610959601983610818565b915061096482610925565b602082019050919050565b5f6020820190508181035f8301526109868161094d565b9050919050565b5f610997826105ec565b91506109a2836105ec565b92508282039050818111156109ba576109b961071a565b5b92915050565b7f4e6f204554482073656e740000000000000000000000000000000000000000005f82015250565b5f6109f4600b83610818565b91506109ff826109c0565b602082019050919050565b5f6020820190508181035f830152610a21816109e8565b9050919050565b7f4e6f7420656e6f75676820746f6b656e20696e207265736572766500000000005f82015250565b5f610a5c601b83610818565b9150610a6782610a28565b602082019050919050565b5f6020820190508181035f830152610a8981610a50565b9050919050565b5f604082019050610aa35f8301856106f2565b610ab0602083018461066f565b939250505056fea26469706673582212208b031f35386aad395fdd91ebab4478b57e40d40eda2b302c7e7b347401b09bbc64736f6c634300081e0033",
}

// AmmABI is the input ABI used to generate the binding from.
// Deprecated: Use AmmMetaData.ABI instead.
var AmmABI = AmmMetaData.ABI

// AmmBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AmmMetaData.Bin instead.
var AmmBin = AmmMetaData.Bin

// DeployAmm deploys a new Ethereum contract, binding an instance of Amm to it.
func DeployAmm(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *Amm, error) {
	parsed, err := AmmMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AmmBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Amm{AmmCaller: AmmCaller{contract: contract}, AmmTransactor: AmmTransactor{contract: contract}, AmmFilterer: AmmFilterer{contract: contract}}, nil
}

// Amm is an auto generated Go binding around an Ethereum contract.
type Amm struct {
	AmmCaller     // Read-only binding to the contract
	AmmTransactor // Write-only binding to the contract
	AmmFilterer   // Log filterer for contract events
}

// AmmCaller is an auto generated read-only Go binding around an Ethereum contract.
type AmmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AmmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AmmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AmmSession struct {
	Contract     *Amm              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AmmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AmmCallerSession struct {
	Contract *AmmCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AmmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AmmTransactorSession struct {
	Contract     *AmmTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AmmRaw is an auto generated low-level Go binding around an Ethereum contract.
type AmmRaw struct {
	Contract *Amm // Generic contract binding to access the raw methods on
}

// AmmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AmmCallerRaw struct {
	Contract *AmmCaller // Generic read-only contract binding to access the raw methods on
}

// AmmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AmmTransactorRaw struct {
	Contract *AmmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAmm creates a new instance of Amm, bound to a specific deployed contract.
func NewAmm(address common.Address, backend bind.ContractBackend) (*Amm, error) {
	contract, err := bindAmm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Amm{AmmCaller: AmmCaller{contract: contract}, AmmTransactor: AmmTransactor{contract: contract}, AmmFilterer: AmmFilterer{contract: contract}}, nil
}

// NewAmmCaller creates a new read-only instance of Amm, bound to a specific deployed contract.
func NewAmmCaller(address common.Address, caller bind.ContractCaller) (*AmmCaller, error) {
	contract, err := bindAmm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AmmCaller{contract: contract}, nil
}

// NewAmmTransactor creates a new write-only instance of Amm, bound to a specific deployed contract.
func NewAmmTransactor(address common.Address, transactor bind.ContractTransactor) (*AmmTransactor, error) {
	contract, err := bindAmm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AmmTransactor{contract: contract}, nil
}

// NewAmmFilterer creates a new log filterer instance of Amm, bound to a specific deployed contract.
func NewAmmFilterer(address common.Address, filterer bind.ContractFilterer) (*AmmFilterer, error) {
	contract, err := bindAmm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AmmFilterer{contract: contract}, nil
}

// bindAmm binds a generic wrapper to an already deployed contract.
func bindAmm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AmmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Amm *AmmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Amm.Contract.AmmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Amm *AmmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Amm.Contract.AmmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Amm *AmmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Amm.Contract.AmmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Amm *AmmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Amm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Amm *AmmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Amm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Amm *AmmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Amm.Contract.contract.Transact(opts, method, params...)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256)
func (_Amm *AmmCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256)
func (_Amm *AmmSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Amm.Contract.GetAmountOut(&_Amm.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256)
func (_Amm *AmmCallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Amm.Contract.GetAmountOut(&_Amm.CallOpts, amountIn, reserveIn, reserveOut)
}

// ReserveETH is a free data retrieval call binding the contract method 0x100af203.
//
// Solidity: function reserveETH() view returns(uint256)
func (_Amm *AmmCaller) ReserveETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "reserveETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveETH is a free data retrieval call binding the contract method 0x100af203.
//
// Solidity: function reserveETH() view returns(uint256)
func (_Amm *AmmSession) ReserveETH() (*big.Int, error) {
	return _Amm.Contract.ReserveETH(&_Amm.CallOpts)
}

// ReserveETH is a free data retrieval call binding the contract method 0x100af203.
//
// Solidity: function reserveETH() view returns(uint256)
func (_Amm *AmmCallerSession) ReserveETH() (*big.Int, error) {
	return _Amm.Contract.ReserveETH(&_Amm.CallOpts)
}

// ReserveToken is a free data retrieval call binding the contract method 0xf4325d67.
//
// Solidity: function reserveToken() view returns(uint256)
func (_Amm *AmmCaller) ReserveToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "reserveToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveToken is a free data retrieval call binding the contract method 0xf4325d67.
//
// Solidity: function reserveToken() view returns(uint256)
func (_Amm *AmmSession) ReserveToken() (*big.Int, error) {
	return _Amm.Contract.ReserveToken(&_Amm.CallOpts)
}

// ReserveToken is a free data retrieval call binding the contract method 0xf4325d67.
//
// Solidity: function reserveToken() view returns(uint256)
func (_Amm *AmmCallerSession) ReserveToken() (*big.Int, error) {
	return _Amm.Contract.ReserveToken(&_Amm.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Amm *AmmCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Amm.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Amm *AmmSession) Token() (common.Address, error) {
	return _Amm.Contract.Token(&_Amm.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Amm *AmmCallerSession) Token() (common.Address, error) {
	return _Amm.Contract.Token(&_Amm.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 amountToken) payable returns()
func (_Amm *AmmTransactor) AddLiquidity(opts *bind.TransactOpts, amountToken *big.Int) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "addLiquidity", amountToken)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 amountToken) payable returns()
func (_Amm *AmmSession) AddLiquidity(amountToken *big.Int) (*types.Transaction, error) {
	return _Amm.Contract.AddLiquidity(&_Amm.TransactOpts, amountToken)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 amountToken) payable returns()
func (_Amm *AmmTransactorSession) AddLiquidity(amountToken *big.Int) (*types.Transaction, error) {
	return _Amm.Contract.AddLiquidity(&_Amm.TransactOpts, amountToken)
}

// SwapExactETHForToken is a paid mutator transaction binding the contract method 0xdcb4b4a1.
//
// Solidity: function swapExactETHForToken() payable returns()
func (_Amm *AmmTransactor) SwapExactETHForToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "swapExactETHForToken")
}

// SwapExactETHForToken is a paid mutator transaction binding the contract method 0xdcb4b4a1.
//
// Solidity: function swapExactETHForToken() payable returns()
func (_Amm *AmmSession) SwapExactETHForToken() (*types.Transaction, error) {
	return _Amm.Contract.SwapExactETHForToken(&_Amm.TransactOpts)
}

// SwapExactETHForToken is a paid mutator transaction binding the contract method 0xdcb4b4a1.
//
// Solidity: function swapExactETHForToken() payable returns()
func (_Amm *AmmTransactorSession) SwapExactETHForToken() (*types.Transaction, error) {
	return _Amm.Contract.SwapExactETHForToken(&_Amm.TransactOpts)
}

// SwapExactTokenForETH is a paid mutator transaction binding the contract method 0xb62aafd5.
//
// Solidity: function swapExactTokenForETH(uint256 amountIn) returns()
func (_Amm *AmmTransactor) SwapExactTokenForETH(opts *bind.TransactOpts, amountIn *big.Int) (*types.Transaction, error) {
	return _Amm.contract.Transact(opts, "swapExactTokenForETH", amountIn)
}

// SwapExactTokenForETH is a paid mutator transaction binding the contract method 0xb62aafd5.
//
// Solidity: function swapExactTokenForETH(uint256 amountIn) returns()
func (_Amm *AmmSession) SwapExactTokenForETH(amountIn *big.Int) (*types.Transaction, error) {
	return _Amm.Contract.SwapExactTokenForETH(&_Amm.TransactOpts, amountIn)
}

// SwapExactTokenForETH is a paid mutator transaction binding the contract method 0xb62aafd5.
//
// Solidity: function swapExactTokenForETH(uint256 amountIn) returns()
func (_Amm *AmmTransactorSession) SwapExactTokenForETH(amountIn *big.Int) (*types.Transaction, error) {
	return _Amm.Contract.SwapExactTokenForETH(&_Amm.TransactOpts, amountIn)
}
