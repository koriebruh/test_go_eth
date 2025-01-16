// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lendingContract

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

// LendingContractMetaData contains all meta data concerning the LendingContract contract.
var LendingContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_interestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_loanAmountLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"LoanDefaulted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dueDate\",\"type\":\"uint256\"}],\"name\":\"LoanIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountPaid\",\"type\":\"uint256\"}],\"name\":\"LoanRepaid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"checkLoanStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"checkOverdueLoan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issueLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"loanAmountLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"loans\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dueDate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"repayLoan\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newInterestRate\",\"type\":\"uint256\"}],\"name\":\"updateInterestRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLoanLimit\",\"type\":\"uint256\"}],\"name\":\"updateLoanAmountLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260405161131f38038061131f83398181016040528101906100259190610142565b5f8211801561003657506127108211155b610075576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161006c90610200565b60405180910390fd5b5f81116100b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100ae9061028e565b60405180910390fd5b335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816001819055508060028190555050506102ac565b5f5ffd5b5f819050919050565b6101218161010f565b811461012b575f5ffd5b50565b5f8151905061013c81610118565b92915050565b5f5f604083850312156101585761015761010b565b5b5f6101658582860161012e565b92505060206101768582860161012e565b9150509250929050565b5f82825260208201905092915050565b7f496e7465726573742072617465206d757374206265206265747765656e2031205f8201527f616e6420313030303020626173697320706f696e747300000000000000000000602082015250565b5f6101ea603683610180565b91506101f582610190565b604082019050919050565b5f6020820190508181035f830152610217816101de565b9050919050565b7f4c6f616e20616d6f756e74206c696d6974206d757374206265206772656174655f8201527f72207468616e2030000000000000000000000000000000000000000000000000602082015250565b5f610278602883610180565b91506102838261021e565b604082019050919050565b5f6020820190508181035f8301526102a58161026c565b9050919050565b611066806102b95f395ff3fe60806040526004361061009b575f3560e01c8063752a50a611610063578063752a50a6146101975780637c3a00fd146101bf5780638da5cb5b146101e9578063b0ae23c114610213578063f4a2c1a11461023b578063f966ade7146102655761009b565b80631895687d1461009f5780633c050cf0146100c75780633ccfd60b14610103578063607fa5151461011957806373b4086b14610158575b5f5ffd5b3480156100aa575f5ffd5b506100c560048036038101906100c09190610b31565b61026f565b005b3480156100d2575f5ffd5b506100ed60048036038101906100e89190610b6f565b610452565b6040516100fa9190610bb4565b60405180910390f35b34801561010e575f5ffd5b50610117610503565b005b348015610124575f5ffd5b5061013f600480360381019061013a9190610b6f565b6105f7565b60405161014f9493929190610bdc565b60405180910390f35b348015610163575f5ffd5b5061017e60048036038101906101799190610b6f565b6106a0565b60405161018e9493929190610bdc565b60405180910390f35b3480156101a2575f5ffd5b506101bd60048036038101906101b89190610c1f565b6106d8565b005b3480156101ca575f5ffd5b506101d3610770565b6040516101e09190610c4a565b60405180910390f35b3480156101f4575f5ffd5b506101fd610776565b60405161020a9190610c72565b60405180910390f35b34801561021e575f5ffd5b5061023960048036038101906102349190610c1f565b61079a565b005b348015610246575f5ffd5b5061024f610832565b60405161025c9190610c4a565b60405180910390f35b61026d610838565b005b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f490610d0b565b60405180910390fd5b600254811115610342576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033990610d73565b60405180910390fd5b5f6064600154836103539190610dbe565b61035d9190610e2c565b90505f62278d004261036f9190610e5c565b905060405180608001604052808481526020018381526020018281526020015f151581525060035f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f015560208201518160010155604082015181600201556060820151816003015f6101000a81548160ff0219169083151502179055509050507f03cd57c87bbf542b3c5edcc7801148ea7d9a169cb16849e0a523ea18eadc58d3848484846040516104449493929190610e8f565b60405180910390a150505050565b5f5f60035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206040518060800160405290815f82015481526020016001820154815260200160028201548152602001600382015f9054906101000a900460ff16151515158152505090508060400151421180156104ea57508060600151155b156104f95760019150506104fe565b5f9150505b919050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610591576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058890610d0b565b60405180910390fd5b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc4790811502906040515f60405180830381858888f193505050501580156105f4573d5f5f3e3d5ffd5b50565b5f5f5f5f5f60035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206040518060800160405290815f82015481526020016001820154815260200160028201548152602001600382015f9054906101000a900460ff1615151515815250509050805f01518160200151826040015183606001519450945094509450509193509193565b6003602052805f5260405f205f91509050805f015490806001015490806002015490806003015f9054906101000a900460ff16905084565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610766576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075d90610d0b565b60405180910390fd5b8060018190555050565b60015481565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610828576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081f90610d0b565b60405180910390fd5b8060028190555050565b60025481565b335f60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0154116108ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b190610f1c565b60405180910390fd5b3360035f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003015f9054906101000a900460ff1615610948576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093f90610f5d565b60405180910390fd5b5f60035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f8160010154825f015461099c9190610e5c565b9050803410156109e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d890610feb565b60405180910390fd5b6001826003015f6101000a81548160ff0219169083151502179055505f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc3490811502906040515f60405180830381858888f19350505050158015610a60573d5f5f3e3d5ffd5b507fc200a1f31dd659e356e0f112c82558e25f49f7b0f84438691cd96f5cb35588233334604051610a92929190611009565b60405180910390a150505050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610acd82610aa4565b9050919050565b610add81610ac3565b8114610ae7575f5ffd5b50565b5f81359050610af881610ad4565b92915050565b5f819050919050565b610b1081610afe565b8114610b1a575f5ffd5b50565b5f81359050610b2b81610b07565b92915050565b5f5f60408385031215610b4757610b46610aa0565b5b5f610b5485828601610aea565b9250506020610b6585828601610b1d565b9150509250929050565b5f60208284031215610b8457610b83610aa0565b5b5f610b9184828501610aea565b91505092915050565b5f8115159050919050565b610bae81610b9a565b82525050565b5f602082019050610bc75f830184610ba5565b92915050565b610bd681610afe565b82525050565b5f608082019050610bef5f830187610bcd565b610bfc6020830186610bcd565b610c096040830185610bcd565b610c166060830184610ba5565b95945050505050565b5f60208284031215610c3457610c33610aa0565b5b5f610c4184828501610b1d565b91505092915050565b5f602082019050610c5d5f830184610bcd565b92915050565b610c6c81610ac3565b82525050565b5f602082019050610c855f830184610c63565b92915050565b5f82825260208201905092915050565b7f4f6e6c7920746865204f776e65722063616e20706572666f726d2074686973205f8201527f616374696f6e0000000000000000000000000000000000000000000000000000602082015250565b5f610cf5602683610c8b565b9150610d0082610c9b565b604082019050919050565b5f6020820190508181035f830152610d2281610ce9565b9050919050565b7f4c6f616e20616d6f756e742065786365656473206c696d6974000000000000005f82015250565b5f610d5d601983610c8b565b9150610d6882610d29565b602082019050919050565b5f6020820190508181035f830152610d8a81610d51565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610dc882610afe565b9150610dd383610afe565b9250828202610de181610afe565b91508282048414831517610df857610df7610d91565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f610e3682610afe565b9150610e4183610afe565b925082610e5157610e50610dff565b5b828204905092915050565b5f610e6682610afe565b9150610e7183610afe565b9250828201905080821115610e8957610e88610d91565b5b92915050565b5f608082019050610ea25f830187610c63565b610eaf6020830186610bcd565b610ebc6040830185610bcd565b610ec96060830184610bcd565b95945050505050565b7f4e6f20616374697665206c6f616e20666f7220746869732061646472657373005f82015250565b5f610f06601f83610c8b565b9150610f1182610ed2565b602082019050919050565b5f6020820190508181035f830152610f3381610efa565b9050919050565b50565b5f610f485f83610c8b565b9150610f5382610f3a565b5f82019050919050565b5f6020820190508181035f830152610f7481610f3d565b9050919050565b7f496e73756666696369656e742066756e647320746f20726570617920746865205f8201527f6c6f616e00000000000000000000000000000000000000000000000000000000602082015250565b5f610fd5602483610c8b565b9150610fe082610f7b565b604082019050919050565b5f6020820190508181035f83015261100281610fc9565b9050919050565b5f60408201905061101c5f830185610c63565b6110296020830184610bcd565b939250505056fea2646970667358221220773aa8dfb9539677cf88c5cfa84d0217e07d8ff071513cebe073ae594957eb6264736f6c634300081c0033",
}

// LendingContractABI is the input ABI used to generate the binding from.
// Deprecated: Use LendingContractMetaData.ABI instead.
var LendingContractABI = LendingContractMetaData.ABI

// LendingContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LendingContractMetaData.Bin instead.
var LendingContractBin = LendingContractMetaData.Bin

// DeployLendingContract deploys a new Ethereum contract, binding an instance of LendingContract to it.
func DeployLendingContract(auth *bind.TransactOpts, backend bind.ContractBackend, _interestRate *big.Int, _loanAmountLimit *big.Int) (common.Address, *types.Transaction, *LendingContract, error) {
	parsed, err := LendingContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LendingContractBin), backend, _interestRate, _loanAmountLimit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LendingContract{LendingContractCaller: LendingContractCaller{contract: contract}, LendingContractTransactor: LendingContractTransactor{contract: contract}, LendingContractFilterer: LendingContractFilterer{contract: contract}}, nil
}

// LendingContract is an auto generated Go binding around an Ethereum contract.
type LendingContract struct {
	LendingContractCaller     // Read-only binding to the contract
	LendingContractTransactor // Write-only binding to the contract
	LendingContractFilterer   // Log filterer for contract events
}

// LendingContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type LendingContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LendingContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LendingContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LendingContractSession struct {
	Contract     *LendingContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LendingContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LendingContractCallerSession struct {
	Contract *LendingContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// LendingContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LendingContractTransactorSession struct {
	Contract     *LendingContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// LendingContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type LendingContractRaw struct {
	Contract *LendingContract // Generic contract binding to access the raw methods on
}

// LendingContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LendingContractCallerRaw struct {
	Contract *LendingContractCaller // Generic read-only contract binding to access the raw methods on
}

// LendingContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LendingContractTransactorRaw struct {
	Contract *LendingContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLendingContract creates a new instance of LendingContract, bound to a specific deployed contract.
func NewLendingContract(address common.Address, backend bind.ContractBackend) (*LendingContract, error) {
	contract, err := bindLendingContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LendingContract{LendingContractCaller: LendingContractCaller{contract: contract}, LendingContractTransactor: LendingContractTransactor{contract: contract}, LendingContractFilterer: LendingContractFilterer{contract: contract}}, nil
}

// NewLendingContractCaller creates a new read-only instance of LendingContract, bound to a specific deployed contract.
func NewLendingContractCaller(address common.Address, caller bind.ContractCaller) (*LendingContractCaller, error) {
	contract, err := bindLendingContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LendingContractCaller{contract: contract}, nil
}

// NewLendingContractTransactor creates a new write-only instance of LendingContract, bound to a specific deployed contract.
func NewLendingContractTransactor(address common.Address, transactor bind.ContractTransactor) (*LendingContractTransactor, error) {
	contract, err := bindLendingContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LendingContractTransactor{contract: contract}, nil
}

// NewLendingContractFilterer creates a new log filterer instance of LendingContract, bound to a specific deployed contract.
func NewLendingContractFilterer(address common.Address, filterer bind.ContractFilterer) (*LendingContractFilterer, error) {
	contract, err := bindLendingContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LendingContractFilterer{contract: contract}, nil
}

// bindLendingContract binds a generic wrapper to an already deployed contract.
func bindLendingContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LendingContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingContract *LendingContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendingContract.Contract.LendingContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingContract *LendingContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingContract.Contract.LendingContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingContract *LendingContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingContract.Contract.LendingContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingContract *LendingContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendingContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingContract *LendingContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingContract *LendingContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingContract.Contract.contract.Transact(opts, method, params...)
}

// CheckLoanStatus is a free data retrieval call binding the contract method 0x607fa515.
//
// Solidity: function checkLoanStatus(address borrower) view returns(uint256, uint256, uint256, bool)
func (_LendingContract *LendingContractCaller) CheckLoanStatus(opts *bind.CallOpts, borrower common.Address) (*big.Int, *big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _LendingContract.contract.Call(opts, &out, "checkLoanStatus", borrower)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)

	return out0, out1, out2, out3, err

}

// CheckLoanStatus is a free data retrieval call binding the contract method 0x607fa515.
//
// Solidity: function checkLoanStatus(address borrower) view returns(uint256, uint256, uint256, bool)
func (_LendingContract *LendingContractSession) CheckLoanStatus(borrower common.Address) (*big.Int, *big.Int, *big.Int, bool, error) {
	return _LendingContract.Contract.CheckLoanStatus(&_LendingContract.CallOpts, borrower)
}

// CheckLoanStatus is a free data retrieval call binding the contract method 0x607fa515.
//
// Solidity: function checkLoanStatus(address borrower) view returns(uint256, uint256, uint256, bool)
func (_LendingContract *LendingContractCallerSession) CheckLoanStatus(borrower common.Address) (*big.Int, *big.Int, *big.Int, bool, error) {
	return _LendingContract.Contract.CheckLoanStatus(&_LendingContract.CallOpts, borrower)
}

// CheckOverdueLoan is a free data retrieval call binding the contract method 0x3c050cf0.
//
// Solidity: function checkOverdueLoan(address borrower) view returns(bool)
func (_LendingContract *LendingContractCaller) CheckOverdueLoan(opts *bind.CallOpts, borrower common.Address) (bool, error) {
	var out []interface{}
	err := _LendingContract.contract.Call(opts, &out, "checkOverdueLoan", borrower)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckOverdueLoan is a free data retrieval call binding the contract method 0x3c050cf0.
//
// Solidity: function checkOverdueLoan(address borrower) view returns(bool)
func (_LendingContract *LendingContractSession) CheckOverdueLoan(borrower common.Address) (bool, error) {
	return _LendingContract.Contract.CheckOverdueLoan(&_LendingContract.CallOpts, borrower)
}

// CheckOverdueLoan is a free data retrieval call binding the contract method 0x3c050cf0.
//
// Solidity: function checkOverdueLoan(address borrower) view returns(bool)
func (_LendingContract *LendingContractCallerSession) CheckOverdueLoan(borrower common.Address) (bool, error) {
	return _LendingContract.Contract.CheckOverdueLoan(&_LendingContract.CallOpts, borrower)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_LendingContract *LendingContractCaller) InterestRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LendingContract.contract.Call(opts, &out, "interestRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_LendingContract *LendingContractSession) InterestRate() (*big.Int, error) {
	return _LendingContract.Contract.InterestRate(&_LendingContract.CallOpts)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_LendingContract *LendingContractCallerSession) InterestRate() (*big.Int, error) {
	return _LendingContract.Contract.InterestRate(&_LendingContract.CallOpts)
}

// LoanAmountLimit is a free data retrieval call binding the contract method 0xf4a2c1a1.
//
// Solidity: function loanAmountLimit() view returns(uint256)
func (_LendingContract *LendingContractCaller) LoanAmountLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LendingContract.contract.Call(opts, &out, "loanAmountLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LoanAmountLimit is a free data retrieval call binding the contract method 0xf4a2c1a1.
//
// Solidity: function loanAmountLimit() view returns(uint256)
func (_LendingContract *LendingContractSession) LoanAmountLimit() (*big.Int, error) {
	return _LendingContract.Contract.LoanAmountLimit(&_LendingContract.CallOpts)
}

// LoanAmountLimit is a free data retrieval call binding the contract method 0xf4a2c1a1.
//
// Solidity: function loanAmountLimit() view returns(uint256)
func (_LendingContract *LendingContractCallerSession) LoanAmountLimit() (*big.Int, error) {
	return _LendingContract.Contract.LoanAmountLimit(&_LendingContract.CallOpts)
}

// Loans is a free data retrieval call binding the contract method 0x73b4086b.
//
// Solidity: function loans(address ) view returns(uint256 amount, uint256 interest, uint256 dueDate, bool isPaid)
func (_LendingContract *LendingContractCaller) Loans(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount   *big.Int
	Interest *big.Int
	DueDate  *big.Int
	IsPaid   bool
}, error) {
	var out []interface{}
	err := _LendingContract.contract.Call(opts, &out, "loans", arg0)

	outstruct := new(struct {
		Amount   *big.Int
		Interest *big.Int
		DueDate  *big.Int
		IsPaid   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Interest = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DueDate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsPaid = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Loans is a free data retrieval call binding the contract method 0x73b4086b.
//
// Solidity: function loans(address ) view returns(uint256 amount, uint256 interest, uint256 dueDate, bool isPaid)
func (_LendingContract *LendingContractSession) Loans(arg0 common.Address) (struct {
	Amount   *big.Int
	Interest *big.Int
	DueDate  *big.Int
	IsPaid   bool
}, error) {
	return _LendingContract.Contract.Loans(&_LendingContract.CallOpts, arg0)
}

// Loans is a free data retrieval call binding the contract method 0x73b4086b.
//
// Solidity: function loans(address ) view returns(uint256 amount, uint256 interest, uint256 dueDate, bool isPaid)
func (_LendingContract *LendingContractCallerSession) Loans(arg0 common.Address) (struct {
	Amount   *big.Int
	Interest *big.Int
	DueDate  *big.Int
	IsPaid   bool
}, error) {
	return _LendingContract.Contract.Loans(&_LendingContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingContract *LendingContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingContract *LendingContractSession) Owner() (common.Address, error) {
	return _LendingContract.Contract.Owner(&_LendingContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingContract *LendingContractCallerSession) Owner() (common.Address, error) {
	return _LendingContract.Contract.Owner(&_LendingContract.CallOpts)
}

// IssueLoan is a paid mutator transaction binding the contract method 0x1895687d.
//
// Solidity: function issueLoan(address borrower, uint256 amount) returns()
func (_LendingContract *LendingContractTransactor) IssueLoan(opts *bind.TransactOpts, borrower common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LendingContract.contract.Transact(opts, "issueLoan", borrower, amount)
}

// IssueLoan is a paid mutator transaction binding the contract method 0x1895687d.
//
// Solidity: function issueLoan(address borrower, uint256 amount) returns()
func (_LendingContract *LendingContractSession) IssueLoan(borrower common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LendingContract.Contract.IssueLoan(&_LendingContract.TransactOpts, borrower, amount)
}

// IssueLoan is a paid mutator transaction binding the contract method 0x1895687d.
//
// Solidity: function issueLoan(address borrower, uint256 amount) returns()
func (_LendingContract *LendingContractTransactorSession) IssueLoan(borrower common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LendingContract.Contract.IssueLoan(&_LendingContract.TransactOpts, borrower, amount)
}

// RepayLoan is a paid mutator transaction binding the contract method 0xf966ade7.
//
// Solidity: function repayLoan() payable returns()
func (_LendingContract *LendingContractTransactor) RepayLoan(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingContract.contract.Transact(opts, "repayLoan")
}

// RepayLoan is a paid mutator transaction binding the contract method 0xf966ade7.
//
// Solidity: function repayLoan() payable returns()
func (_LendingContract *LendingContractSession) RepayLoan() (*types.Transaction, error) {
	return _LendingContract.Contract.RepayLoan(&_LendingContract.TransactOpts)
}

// RepayLoan is a paid mutator transaction binding the contract method 0xf966ade7.
//
// Solidity: function repayLoan() payable returns()
func (_LendingContract *LendingContractTransactorSession) RepayLoan() (*types.Transaction, error) {
	return _LendingContract.Contract.RepayLoan(&_LendingContract.TransactOpts)
}

// UpdateInterestRate is a paid mutator transaction binding the contract method 0x752a50a6.
//
// Solidity: function updateInterestRate(uint256 newInterestRate) returns()
func (_LendingContract *LendingContractTransactor) UpdateInterestRate(opts *bind.TransactOpts, newInterestRate *big.Int) (*types.Transaction, error) {
	return _LendingContract.contract.Transact(opts, "updateInterestRate", newInterestRate)
}

// UpdateInterestRate is a paid mutator transaction binding the contract method 0x752a50a6.
//
// Solidity: function updateInterestRate(uint256 newInterestRate) returns()
func (_LendingContract *LendingContractSession) UpdateInterestRate(newInterestRate *big.Int) (*types.Transaction, error) {
	return _LendingContract.Contract.UpdateInterestRate(&_LendingContract.TransactOpts, newInterestRate)
}

// UpdateInterestRate is a paid mutator transaction binding the contract method 0x752a50a6.
//
// Solidity: function updateInterestRate(uint256 newInterestRate) returns()
func (_LendingContract *LendingContractTransactorSession) UpdateInterestRate(newInterestRate *big.Int) (*types.Transaction, error) {
	return _LendingContract.Contract.UpdateInterestRate(&_LendingContract.TransactOpts, newInterestRate)
}

// UpdateLoanAmountLimit is a paid mutator transaction binding the contract method 0xb0ae23c1.
//
// Solidity: function updateLoanAmountLimit(uint256 newLoanLimit) returns()
func (_LendingContract *LendingContractTransactor) UpdateLoanAmountLimit(opts *bind.TransactOpts, newLoanLimit *big.Int) (*types.Transaction, error) {
	return _LendingContract.contract.Transact(opts, "updateLoanAmountLimit", newLoanLimit)
}

// UpdateLoanAmountLimit is a paid mutator transaction binding the contract method 0xb0ae23c1.
//
// Solidity: function updateLoanAmountLimit(uint256 newLoanLimit) returns()
func (_LendingContract *LendingContractSession) UpdateLoanAmountLimit(newLoanLimit *big.Int) (*types.Transaction, error) {
	return _LendingContract.Contract.UpdateLoanAmountLimit(&_LendingContract.TransactOpts, newLoanLimit)
}

// UpdateLoanAmountLimit is a paid mutator transaction binding the contract method 0xb0ae23c1.
//
// Solidity: function updateLoanAmountLimit(uint256 newLoanLimit) returns()
func (_LendingContract *LendingContractTransactorSession) UpdateLoanAmountLimit(newLoanLimit *big.Int) (*types.Transaction, error) {
	return _LendingContract.Contract.UpdateLoanAmountLimit(&_LendingContract.TransactOpts, newLoanLimit)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_LendingContract *LendingContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingContract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_LendingContract *LendingContractSession) Withdraw() (*types.Transaction, error) {
	return _LendingContract.Contract.Withdraw(&_LendingContract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_LendingContract *LendingContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _LendingContract.Contract.Withdraw(&_LendingContract.TransactOpts)
}

// LendingContractLoanDefaultedIterator is returned from FilterLoanDefaulted and is used to iterate over the raw logs and unpacked data for LoanDefaulted events raised by the LendingContract contract.
type LendingContractLoanDefaultedIterator struct {
	Event *LendingContractLoanDefaulted // Event containing the contract specifics and raw log

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
func (it *LendingContractLoanDefaultedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingContractLoanDefaulted)
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
		it.Event = new(LendingContractLoanDefaulted)
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
func (it *LendingContractLoanDefaultedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingContractLoanDefaultedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingContractLoanDefaulted represents a LoanDefaulted event raised by the LendingContract contract.
type LendingContractLoanDefaulted struct {
	Borrower common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLoanDefaulted is a free log retrieval operation binding the contract event 0xf4e5cf42cbb1d0c6c91e16d925349feb6c5d970652f6ec77b1ff1262f7284cd5.
//
// Solidity: event LoanDefaulted(address borrower)
func (_LendingContract *LendingContractFilterer) FilterLoanDefaulted(opts *bind.FilterOpts) (*LendingContractLoanDefaultedIterator, error) {

	logs, sub, err := _LendingContract.contract.FilterLogs(opts, "LoanDefaulted")
	if err != nil {
		return nil, err
	}
	return &LendingContractLoanDefaultedIterator{contract: _LendingContract.contract, event: "LoanDefaulted", logs: logs, sub: sub}, nil
}

// WatchLoanDefaulted is a free log subscription operation binding the contract event 0xf4e5cf42cbb1d0c6c91e16d925349feb6c5d970652f6ec77b1ff1262f7284cd5.
//
// Solidity: event LoanDefaulted(address borrower)
func (_LendingContract *LendingContractFilterer) WatchLoanDefaulted(opts *bind.WatchOpts, sink chan<- *LendingContractLoanDefaulted) (event.Subscription, error) {

	logs, sub, err := _LendingContract.contract.WatchLogs(opts, "LoanDefaulted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingContractLoanDefaulted)
				if err := _LendingContract.contract.UnpackLog(event, "LoanDefaulted", log); err != nil {
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

// ParseLoanDefaulted is a log parse operation binding the contract event 0xf4e5cf42cbb1d0c6c91e16d925349feb6c5d970652f6ec77b1ff1262f7284cd5.
//
// Solidity: event LoanDefaulted(address borrower)
func (_LendingContract *LendingContractFilterer) ParseLoanDefaulted(log types.Log) (*LendingContractLoanDefaulted, error) {
	event := new(LendingContractLoanDefaulted)
	if err := _LendingContract.contract.UnpackLog(event, "LoanDefaulted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingContractLoanIssuedIterator is returned from FilterLoanIssued and is used to iterate over the raw logs and unpacked data for LoanIssued events raised by the LendingContract contract.
type LendingContractLoanIssuedIterator struct {
	Event *LendingContractLoanIssued // Event containing the contract specifics and raw log

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
func (it *LendingContractLoanIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingContractLoanIssued)
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
		it.Event = new(LendingContractLoanIssued)
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
func (it *LendingContractLoanIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingContractLoanIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingContractLoanIssued represents a LoanIssued event raised by the LendingContract contract.
type LendingContractLoanIssued struct {
	Borrower common.Address
	Amount   *big.Int
	Interest *big.Int
	DueDate  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLoanIssued is a free log retrieval operation binding the contract event 0x03cd57c87bbf542b3c5edcc7801148ea7d9a169cb16849e0a523ea18eadc58d3.
//
// Solidity: event LoanIssued(address borrower, uint256 amount, uint256 interest, uint256 dueDate)
func (_LendingContract *LendingContractFilterer) FilterLoanIssued(opts *bind.FilterOpts) (*LendingContractLoanIssuedIterator, error) {

	logs, sub, err := _LendingContract.contract.FilterLogs(opts, "LoanIssued")
	if err != nil {
		return nil, err
	}
	return &LendingContractLoanIssuedIterator{contract: _LendingContract.contract, event: "LoanIssued", logs: logs, sub: sub}, nil
}

// WatchLoanIssued is a free log subscription operation binding the contract event 0x03cd57c87bbf542b3c5edcc7801148ea7d9a169cb16849e0a523ea18eadc58d3.
//
// Solidity: event LoanIssued(address borrower, uint256 amount, uint256 interest, uint256 dueDate)
func (_LendingContract *LendingContractFilterer) WatchLoanIssued(opts *bind.WatchOpts, sink chan<- *LendingContractLoanIssued) (event.Subscription, error) {

	logs, sub, err := _LendingContract.contract.WatchLogs(opts, "LoanIssued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingContractLoanIssued)
				if err := _LendingContract.contract.UnpackLog(event, "LoanIssued", log); err != nil {
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

// ParseLoanIssued is a log parse operation binding the contract event 0x03cd57c87bbf542b3c5edcc7801148ea7d9a169cb16849e0a523ea18eadc58d3.
//
// Solidity: event LoanIssued(address borrower, uint256 amount, uint256 interest, uint256 dueDate)
func (_LendingContract *LendingContractFilterer) ParseLoanIssued(log types.Log) (*LendingContractLoanIssued, error) {
	event := new(LendingContractLoanIssued)
	if err := _LendingContract.contract.UnpackLog(event, "LoanIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingContractLoanRepaidIterator is returned from FilterLoanRepaid and is used to iterate over the raw logs and unpacked data for LoanRepaid events raised by the LendingContract contract.
type LendingContractLoanRepaidIterator struct {
	Event *LendingContractLoanRepaid // Event containing the contract specifics and raw log

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
func (it *LendingContractLoanRepaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingContractLoanRepaid)
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
		it.Event = new(LendingContractLoanRepaid)
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
func (it *LendingContractLoanRepaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingContractLoanRepaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingContractLoanRepaid represents a LoanRepaid event raised by the LendingContract contract.
type LendingContractLoanRepaid struct {
	Borrower   common.Address
	AmountPaid *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLoanRepaid is a free log retrieval operation binding the contract event 0xc200a1f31dd659e356e0f112c82558e25f49f7b0f84438691cd96f5cb3558823.
//
// Solidity: event LoanRepaid(address borrower, uint256 amountPaid)
func (_LendingContract *LendingContractFilterer) FilterLoanRepaid(opts *bind.FilterOpts) (*LendingContractLoanRepaidIterator, error) {

	logs, sub, err := _LendingContract.contract.FilterLogs(opts, "LoanRepaid")
	if err != nil {
		return nil, err
	}
	return &LendingContractLoanRepaidIterator{contract: _LendingContract.contract, event: "LoanRepaid", logs: logs, sub: sub}, nil
}

// WatchLoanRepaid is a free log subscription operation binding the contract event 0xc200a1f31dd659e356e0f112c82558e25f49f7b0f84438691cd96f5cb3558823.
//
// Solidity: event LoanRepaid(address borrower, uint256 amountPaid)
func (_LendingContract *LendingContractFilterer) WatchLoanRepaid(opts *bind.WatchOpts, sink chan<- *LendingContractLoanRepaid) (event.Subscription, error) {

	logs, sub, err := _LendingContract.contract.WatchLogs(opts, "LoanRepaid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingContractLoanRepaid)
				if err := _LendingContract.contract.UnpackLog(event, "LoanRepaid", log); err != nil {
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

// ParseLoanRepaid is a log parse operation binding the contract event 0xc200a1f31dd659e356e0f112c82558e25f49f7b0f84438691cd96f5cb3558823.
//
// Solidity: event LoanRepaid(address borrower, uint256 amountPaid)
func (_LendingContract *LendingContractFilterer) ParseLoanRepaid(log types.Log) (*LendingContractLoanRepaid, error) {
	event := new(LendingContractLoanRepaid)
	if err := _LendingContract.contract.UnpackLog(event, "LoanRepaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
