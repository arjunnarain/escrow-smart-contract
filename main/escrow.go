// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// EscrowBankDetails is an auto generated low-level Go binding around an user-defined struct.
type EscrowBankDetails struct {
	AccountHolderName string
	AccountNumber     string
	BankName          string
	IfscCode          string
}

// EscrowPayoutAuth is an auto generated low-level Go binding around an user-defined struct.
type EscrowPayoutAuth struct {
	RazorpayApiKey    string
	RazorpayApiSecret string
}

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountINR\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"accountHolderName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"accountNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bankName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ifscCode\",\"type\":\"string\"}],\"internalType\":\"structEscrow.BankDetails\",\"name\":\"_buyerBankDetails\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"accountHolderName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"accountNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bankName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ifscCode\",\"type\":\"string\"}],\"internalType\":\"structEscrow.BankDetails\",\"name\":\"_sellerBankDetails\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"razorpayApiKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"razorpayApiSecret\",\"type\":\"string\"}],\"internalType\":\"structEscrow.PayoutAuth\",\"name\":\"_payoutAuth\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ApprovalReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"buyerAccount\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sellerAccount\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountINR\",\"type\":\"uint256\"}],\"name\":\"PayoutRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"amountINR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approveByBuyer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approveBySeller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approveByTrustee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyerApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyerBankDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"accountHolderName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"accountNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bankName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ifscCode\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disapproveByBuyer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disapproveBySeller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payoutAuth\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"razorpayApiKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"razorpayApiSecret\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellerApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellerBankDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"accountHolderName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"accountNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bankName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ifscCode\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustee\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trusteeApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051611c3b380380611c3b83398181016040528101906100319190610537565b865f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508560015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508460025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550836003819055508260045f820151815f01908161010d919061082c565b506020820151816001019081610123919061082c565b506040820151816002019081610139919061082c565b50606082015181600301908161014f919061082c565b509050508160085f820151815f019081610169919061082c565b50602082015181600101908161017f919061082c565b506040820151816002019081610195919061082c565b5060608201518160030190816101ab919061082c565b5090505080600c5f820151815f0190816101c5919061082c565b5060208201518160010190816101db919061082c565b50905050505050505050506108fb565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610225826101fc565b9050919050565b6102358161021b565b811461023f575f80fd5b50565b5f815190506102508161022c565b92915050565b5f819050919050565b61026881610256565b8114610272575f80fd5b50565b5f815190506102838161025f565b92915050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6102d38261028d565b810181811067ffffffffffffffff821117156102f2576102f161029d565b5b80604052505050565b5f6103046101eb565b905061031082826102ca565b919050565b5f80fd5b5f80fd5b5f80fd5b5f67ffffffffffffffff82111561033b5761033a61029d565b5b6103448261028d565b9050602081019050919050565b8281835e5f83830152505050565b5f61037161036c84610321565b6102fb565b90508281526020810184848401111561038d5761038c61031d565b5b610398848285610351565b509392505050565b5f82601f8301126103b4576103b3610319565b5b81516103c484826020860161035f565b91505092915050565b5f608082840312156103e2576103e1610289565b5b6103ec60806102fb565b90505f82015167ffffffffffffffff81111561040b5761040a610315565b5b610417848285016103a0565b5f83015250602082015167ffffffffffffffff81111561043a57610439610315565b5b610446848285016103a0565b602083015250604082015167ffffffffffffffff81111561046a57610469610315565b5b610476848285016103a0565b604083015250606082015167ffffffffffffffff81111561049a57610499610315565b5b6104a6848285016103a0565b60608301525092915050565b5f604082840312156104c7576104c6610289565b5b6104d160406102fb565b90505f82015167ffffffffffffffff8111156104f0576104ef610315565b5b6104fc848285016103a0565b5f83015250602082015167ffffffffffffffff81111561051f5761051e610315565b5b61052b848285016103a0565b60208301525092915050565b5f805f805f805f60e0888a031215610552576105516101f4565b5b5f61055f8a828b01610242565b97505060206105708a828b01610242565b96505060406105818a828b01610242565b95505060606105928a828b01610275565b945050608088015167ffffffffffffffff8111156105b3576105b26101f8565b5b6105bf8a828b016103cd565b93505060a088015167ffffffffffffffff8111156105e0576105df6101f8565b5b6105ec8a828b016103cd565b92505060c088015167ffffffffffffffff81111561060d5761060c6101f8565b5b6106198a828b016104b2565b91505092959891949750929550565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061067657607f821691505b60208210810361068957610688610632565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026106eb7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826106b0565b6106f586836106b0565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61073061072b61072684610256565b61070d565b610256565b9050919050565b5f819050919050565b61074983610716565b61075d61075582610737565b8484546106bc565b825550505050565b5f90565b610771610765565b61077c818484610740565b505050565b5b8181101561079f576107945f82610769565b600181019050610782565b5050565b601f8211156107e4576107b58161068f565b6107be846106a1565b810160208510156107cd578190505b6107e16107d9856106a1565b830182610781565b50505b505050565b5f82821c905092915050565b5f6108045f19846008026107e9565b1980831691505092915050565b5f61081c83836107f5565b9150826002028217905092915050565b61083582610628565b67ffffffffffffffff81111561084e5761084d61029d565b5b610858825461065f565b6108638282856107a3565b5f60209050601f831160018114610894575f8415610882578287015190505b61088c8582610811565b8655506108f3565b601f1984166108a28661068f565b5f5b828110156108c9578489015182556001820191506020850194506020810190506108a4565b868310156108e657848901516108e2601f8916826107f5565b8355505b6001600288020188555050505b505050505050565b611333806109085f395ff3fe608060405234801561000f575f80fd5b50600436106100f3575f3560e01c80637150d8ae11610095578063f41d2d7111610064578063f41d2d711461020c578063fa65600a14610216578063fdf97cb214610234578063ff80a38914610252576100f3565b80637150d8ae146101a4578063b1084e7f146101c2578063da849476146101e3578063dd257bde14610202576100f3565b80631afcb321116100d15780631afcb3211461015157806332fc28ec1461016f5780636b566475146101795780636cc332ce1461019a576100f3565b806308551a53146100f75780630a1792cc14610115578063183f66b614610133575b5f80fd5b6100ff61025c565b60405161010c9190610e41565b60405180910390f35b61011d610281565b60405161012a9190610e74565b60405180910390f35b61013b610293565b6040516101489190610e74565b60405180910390f35b6101596102a6565b6040516101669190610e74565b60405180910390f35b6101776102b9565b005b6101816103a7565b6040516101919493929190610efd565b60405180910390f35b6101a26105dc565b005b6101ac6106d3565b6040516101b99190610e41565b60405180910390f35b6101ca6106f6565b6040516101da9493929190610efd565b60405180910390f35b6101eb61092b565b6040516101f9929190610f5c565b60405180910390f35b61020a610a48565b005b610214610b33565b005b61021e610c2a565b60405161022b9190610fa9565b60405180910390f35b61023c610c30565b6040516102499190610e41565b60405180910390f35b61025a610c55565b005b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600e5f9054906101000a900460ff1681565b600e60019054906101000a900460ff1681565b600e60029054906101000a900460ff1681565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610348576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033f9061100c565b60405180910390fd5b5f600e60016101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f135b662e7c460f01a8962dabbc952859694282ca1f1bfbacf24ebecc00160d3560405160405180910390a2565b6008805f0180546103b790611057565b80601f01602080910402602001604051908101604052809291908181526020018280546103e390611057565b801561042e5780601f106104055761010080835404028352916020019161042e565b820191905f5260205f20905b81548152906001019060200180831161041157829003601f168201915b50505050509080600101805461044390611057565b80601f016020809104026020016040519081016040528092919081815260200182805461046f90611057565b80156104ba5780601f10610491576101008083540402835291602001916104ba565b820191905f5260205f20905b81548152906001019060200180831161049d57829003601f168201915b5050505050908060020180546104cf90611057565b80601f01602080910402602001604051908101604052809291908181526020018280546104fb90611057565b80156105465780601f1061051d57610100808354040283529160200191610546565b820191905f5260205f20905b81548152906001019060200180831161052957829003601f168201915b50505050509080600301805461055b90611057565b80601f016020809104026020016040519081016040528092919081815260200182805461058790611057565b80156105d25780601f106105a9576101008083540402835291602001916105d2565b820191905f5260205f20905b8154815290600101906020018083116105b557829003601f168201915b5050505050905084565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461066b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610662906110d1565b60405180910390fd5b6001600e60016101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f135b662e7c460f01a8962dabbc952859694282ca1f1bfbacf24ebecc00160d3560405160405180910390a26106d1610d49565b565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6004805f01805461070690611057565b80601f016020809104026020016040519081016040528092919081815260200182805461073290611057565b801561077d5780601f106107545761010080835404028352916020019161077d565b820191905f5260205f20905b81548152906001019060200180831161076057829003601f168201915b50505050509080600101805461079290611057565b80601f01602080910402602001604051908101604052809291908181526020018280546107be90611057565b80156108095780601f106107e057610100808354040283529160200191610809565b820191905f5260205f20905b8154815290600101906020018083116107ec57829003601f168201915b50505050509080600201805461081e90611057565b80601f016020809104026020016040519081016040528092919081815260200182805461084a90611057565b80156108955780601f1061086c57610100808354040283529160200191610895565b820191905f5260205f20905b81548152906001019060200180831161087857829003601f168201915b5050505050908060030180546108aa90611057565b80601f01602080910402602001604051908101604052809291908181526020018280546108d690611057565b80156109215780601f106108f857610100808354040283529160200191610921565b820191905f5260205f20905b81548152906001019060200180831161090457829003601f168201915b5050505050905084565b600c805f01805461093b90611057565b80601f016020809104026020016040519081016040528092919081815260200182805461096790611057565b80156109b25780601f10610989576101008083540402835291602001916109b2565b820191905f5260205f20905b81548152906001019060200180831161099557829003601f168201915b5050505050908060010180546109c790611057565b80601f01602080910402602001604051908101604052809291908181526020018280546109f390611057565b8015610a3e5780601f10610a1557610100808354040283529160200191610a3e565b820191905f5260205f20905b815481529060010190602001808311610a2157829003601f168201915b5050505050905082565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ad5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610acc90611139565b60405180910390fd5b5f600e5f6101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f135b662e7c460f01a8962dabbc952859694282ca1f1bfbacf24ebecc00160d3560405160405180910390a2565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bc2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bb9906111a1565b60405180910390fd5b6001600e60026101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f135b662e7c460f01a8962dabbc952859694282ca1f1bfbacf24ebecc00160d3560405160405180910390a2610c28610d49565b565b60035481565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ce2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cd990611209565b60405180910390fd5b6001600e5f6101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f135b662e7c460f01a8962dabbc952859694282ca1f1bfbacf24ebecc00160d3560405160405180910390a2610d47610d49565b565b600e5f9054906101000a900460ff168015610d705750600e60019054906101000a900460ff165b80610db55750600e60029054906101000a900460ff168015610db45750600e5f9054906101000a900460ff1680610db35750600e60019054906101000a900460ff165b5b5b15610e00577f4423fec52feed3aa8bd2a99c8211cfcc2672119920cb1d206da40301c641d16d60046001016008600101600354604051610df7939291906112ba565b60405180910390a15b565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610e2b82610e02565b9050919050565b610e3b81610e21565b82525050565b5f602082019050610e545f830184610e32565b92915050565b5f8115159050919050565b610e6e81610e5a565b82525050565b5f602082019050610e875f830184610e65565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610ecf82610e8d565b610ed98185610e97565b9350610ee9818560208601610ea7565b610ef281610eb5565b840191505092915050565b5f6080820190508181035f830152610f158187610ec5565b90508181036020830152610f298186610ec5565b90508181036040830152610f3d8185610ec5565b90508181036060830152610f518184610ec5565b905095945050505050565b5f6040820190508181035f830152610f748185610ec5565b90508181036020830152610f888184610ec5565b90509392505050565b5f819050919050565b610fa381610f91565b82525050565b5f602082019050610fbc5f830184610f9a565b92915050565b7f4f6e6c792073656c6c65722063616e20646973617070726f76650000000000005f82015250565b5f610ff6601a83610e97565b915061100182610fc2565b602082019050919050565b5f6020820190508181035f83015261102381610fea565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061106e57607f821691505b6020821081036110815761108061102a565b5b50919050565b7f4f6e6c792073656c6c65722063616e20617070726f76650000000000000000005f82015250565b5f6110bb601783610e97565b91506110c682611087565b602082019050919050565b5f6020820190508181035f8301526110e8816110af565b9050919050565b7f4f6e6c792062757965722063616e20646973617070726f7665000000000000005f82015250565b5f611123601983610e97565b915061112e826110ef565b602082019050919050565b5f6020820190508181035f83015261115081611117565b9050919050565b7f4f6e6c7920747275737465652063616e20617070726f766500000000000000005f82015250565b5f61118b601883610e97565b915061119682611157565b602082019050919050565b5f6020820190508181035f8301526111b88161117f565b9050919050565b7f4f6e6c792062757965722063616e20617070726f7665000000000000000000005f82015250565b5f6111f3601683610e97565b91506111fe826111bf565b602082019050919050565b5f6020820190508181035f830152611220816111e7565b9050919050565b5f819050815f5260205f209050919050565b5f815461124581611057565b61124f8186610e97565b9450600182165f8114611269576001811461127f576112b1565b60ff1983168652811515602002860193506112b1565b61128885611227565b5f5b838110156112a95781548189015260018201915060208101905061128a565b808801955050505b50505092915050565b5f6060820190508181035f8301526112d28186611239565b905081810360208301526112e68185611239565b90506112f56040830184610f9a565b94935050505056fea26469706673582212205bbde23ca448b964a7339cc354bfbf3051cd6babd9f58d5e4f79382c83aec9fa64736f6c634300081a0033",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// MainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MainMetaData.Bin instead.
var MainBin = MainMetaData.Bin

// DeployMain deploys a new Ethereum contract, binding an instance of Main to it.
func DeployMain(auth *bind.TransactOpts, backend bind.ContractBackend, _buyer common.Address, _seller common.Address, _trustee common.Address, _amountINR *big.Int, _buyerBankDetails EscrowBankDetails, _sellerBankDetails EscrowBankDetails, _payoutAuth EscrowPayoutAuth) (common.Address, *types.Transaction, *Main, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MainBin), backend, _buyer, _seller, _trustee, _amountINR, _buyerBankDetails, _sellerBankDetails, _payoutAuth)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// AmountINR is a free data retrieval call binding the contract method 0xfa65600a.
//
// Solidity: function amountINR() view returns(uint256)
func (_Main *MainCaller) AmountINR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "amountINR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountINR is a free data retrieval call binding the contract method 0xfa65600a.
//
// Solidity: function amountINR() view returns(uint256)
func (_Main *MainSession) AmountINR() (*big.Int, error) {
	return _Main.Contract.AmountINR(&_Main.CallOpts)
}

// AmountINR is a free data retrieval call binding the contract method 0xfa65600a.
//
// Solidity: function amountINR() view returns(uint256)
func (_Main *MainCallerSession) AmountINR() (*big.Int, error) {
	return _Main.Contract.AmountINR(&_Main.CallOpts)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Main *MainCaller) Buyer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "buyer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Main *MainSession) Buyer() (common.Address, error) {
	return _Main.Contract.Buyer(&_Main.CallOpts)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Main *MainCallerSession) Buyer() (common.Address, error) {
	return _Main.Contract.Buyer(&_Main.CallOpts)
}

// BuyerApproval is a free data retrieval call binding the contract method 0x0a1792cc.
//
// Solidity: function buyerApproval() view returns(bool)
func (_Main *MainCaller) BuyerApproval(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "buyerApproval")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BuyerApproval is a free data retrieval call binding the contract method 0x0a1792cc.
//
// Solidity: function buyerApproval() view returns(bool)
func (_Main *MainSession) BuyerApproval() (bool, error) {
	return _Main.Contract.BuyerApproval(&_Main.CallOpts)
}

// BuyerApproval is a free data retrieval call binding the contract method 0x0a1792cc.
//
// Solidity: function buyerApproval() view returns(bool)
func (_Main *MainCallerSession) BuyerApproval() (bool, error) {
	return _Main.Contract.BuyerApproval(&_Main.CallOpts)
}

// BuyerBankDetails is a free data retrieval call binding the contract method 0xb1084e7f.
//
// Solidity: function buyerBankDetails() view returns(string accountHolderName, string accountNumber, string bankName, string ifscCode)
func (_Main *MainCaller) BuyerBankDetails(opts *bind.CallOpts) (struct {
	AccountHolderName string
	AccountNumber     string
	BankName          string
	IfscCode          string
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "buyerBankDetails")

	outstruct := new(struct {
		AccountHolderName string
		AccountNumber     string
		BankName          string
		IfscCode          string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AccountHolderName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.AccountNumber = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.BankName = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.IfscCode = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// BuyerBankDetails is a free data retrieval call binding the contract method 0xb1084e7f.
//
// Solidity: function buyerBankDetails() view returns(string accountHolderName, string accountNumber, string bankName, string ifscCode)
func (_Main *MainSession) BuyerBankDetails() (struct {
	AccountHolderName string
	AccountNumber     string
	BankName          string
	IfscCode          string
}, error) {
	return _Main.Contract.BuyerBankDetails(&_Main.CallOpts)
}

// BuyerBankDetails is a free data retrieval call binding the contract method 0xb1084e7f.
//
// Solidity: function buyerBankDetails() view returns(string accountHolderName, string accountNumber, string bankName, string ifscCode)
func (_Main *MainCallerSession) BuyerBankDetails() (struct {
	AccountHolderName string
	AccountNumber     string
	BankName          string
	IfscCode          string
}, error) {
	return _Main.Contract.BuyerBankDetails(&_Main.CallOpts)
}

// PayoutAuth is a free data retrieval call binding the contract method 0xda849476.
//
// Solidity: function payoutAuth() view returns(string razorpayApiKey, string razorpayApiSecret)
func (_Main *MainCaller) PayoutAuth(opts *bind.CallOpts) (struct {
	RazorpayApiKey    string
	RazorpayApiSecret string
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "payoutAuth")

	outstruct := new(struct {
		RazorpayApiKey    string
		RazorpayApiSecret string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RazorpayApiKey = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.RazorpayApiSecret = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// PayoutAuth is a free data retrieval call binding the contract method 0xda849476.
//
// Solidity: function payoutAuth() view returns(string razorpayApiKey, string razorpayApiSecret)
func (_Main *MainSession) PayoutAuth() (struct {
	RazorpayApiKey    string
	RazorpayApiSecret string
}, error) {
	return _Main.Contract.PayoutAuth(&_Main.CallOpts)
}

// PayoutAuth is a free data retrieval call binding the contract method 0xda849476.
//
// Solidity: function payoutAuth() view returns(string razorpayApiKey, string razorpayApiSecret)
func (_Main *MainCallerSession) PayoutAuth() (struct {
	RazorpayApiKey    string
	RazorpayApiSecret string
}, error) {
	return _Main.Contract.PayoutAuth(&_Main.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Main *MainCaller) Seller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "seller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Main *MainSession) Seller() (common.Address, error) {
	return _Main.Contract.Seller(&_Main.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Main *MainCallerSession) Seller() (common.Address, error) {
	return _Main.Contract.Seller(&_Main.CallOpts)
}

// SellerApproval is a free data retrieval call binding the contract method 0x183f66b6.
//
// Solidity: function sellerApproval() view returns(bool)
func (_Main *MainCaller) SellerApproval(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "sellerApproval")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SellerApproval is a free data retrieval call binding the contract method 0x183f66b6.
//
// Solidity: function sellerApproval() view returns(bool)
func (_Main *MainSession) SellerApproval() (bool, error) {
	return _Main.Contract.SellerApproval(&_Main.CallOpts)
}

// SellerApproval is a free data retrieval call binding the contract method 0x183f66b6.
//
// Solidity: function sellerApproval() view returns(bool)
func (_Main *MainCallerSession) SellerApproval() (bool, error) {
	return _Main.Contract.SellerApproval(&_Main.CallOpts)
}

// SellerBankDetails is a free data retrieval call binding the contract method 0x6b566475.
//
// Solidity: function sellerBankDetails() view returns(string accountHolderName, string accountNumber, string bankName, string ifscCode)
func (_Main *MainCaller) SellerBankDetails(opts *bind.CallOpts) (struct {
	AccountHolderName string
	AccountNumber     string
	BankName          string
	IfscCode          string
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "sellerBankDetails")

	outstruct := new(struct {
		AccountHolderName string
		AccountNumber     string
		BankName          string
		IfscCode          string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AccountHolderName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.AccountNumber = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.BankName = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.IfscCode = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// SellerBankDetails is a free data retrieval call binding the contract method 0x6b566475.
//
// Solidity: function sellerBankDetails() view returns(string accountHolderName, string accountNumber, string bankName, string ifscCode)
func (_Main *MainSession) SellerBankDetails() (struct {
	AccountHolderName string
	AccountNumber     string
	BankName          string
	IfscCode          string
}, error) {
	return _Main.Contract.SellerBankDetails(&_Main.CallOpts)
}

// SellerBankDetails is a free data retrieval call binding the contract method 0x6b566475.
//
// Solidity: function sellerBankDetails() view returns(string accountHolderName, string accountNumber, string bankName, string ifscCode)
func (_Main *MainCallerSession) SellerBankDetails() (struct {
	AccountHolderName string
	AccountNumber     string
	BankName          string
	IfscCode          string
}, error) {
	return _Main.Contract.SellerBankDetails(&_Main.CallOpts)
}

// Trustee is a free data retrieval call binding the contract method 0xfdf97cb2.
//
// Solidity: function trustee() view returns(address)
func (_Main *MainCaller) Trustee(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "trustee")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Trustee is a free data retrieval call binding the contract method 0xfdf97cb2.
//
// Solidity: function trustee() view returns(address)
func (_Main *MainSession) Trustee() (common.Address, error) {
	return _Main.Contract.Trustee(&_Main.CallOpts)
}

// Trustee is a free data retrieval call binding the contract method 0xfdf97cb2.
//
// Solidity: function trustee() view returns(address)
func (_Main *MainCallerSession) Trustee() (common.Address, error) {
	return _Main.Contract.Trustee(&_Main.CallOpts)
}

// TrusteeApproval is a free data retrieval call binding the contract method 0x1afcb321.
//
// Solidity: function trusteeApproval() view returns(bool)
func (_Main *MainCaller) TrusteeApproval(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "trusteeApproval")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TrusteeApproval is a free data retrieval call binding the contract method 0x1afcb321.
//
// Solidity: function trusteeApproval() view returns(bool)
func (_Main *MainSession) TrusteeApproval() (bool, error) {
	return _Main.Contract.TrusteeApproval(&_Main.CallOpts)
}

// TrusteeApproval is a free data retrieval call binding the contract method 0x1afcb321.
//
// Solidity: function trusteeApproval() view returns(bool)
func (_Main *MainCallerSession) TrusteeApproval() (bool, error) {
	return _Main.Contract.TrusteeApproval(&_Main.CallOpts)
}

// ApproveByBuyer is a paid mutator transaction binding the contract method 0xff80a389.
//
// Solidity: function approveByBuyer() returns()
func (_Main *MainTransactor) ApproveByBuyer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "approveByBuyer")
}

// ApproveByBuyer is a paid mutator transaction binding the contract method 0xff80a389.
//
// Solidity: function approveByBuyer() returns()
func (_Main *MainSession) ApproveByBuyer() (*types.Transaction, error) {
	return _Main.Contract.ApproveByBuyer(&_Main.TransactOpts)
}

// ApproveByBuyer is a paid mutator transaction binding the contract method 0xff80a389.
//
// Solidity: function approveByBuyer() returns()
func (_Main *MainTransactorSession) ApproveByBuyer() (*types.Transaction, error) {
	return _Main.Contract.ApproveByBuyer(&_Main.TransactOpts)
}

// ApproveBySeller is a paid mutator transaction binding the contract method 0x6cc332ce.
//
// Solidity: function approveBySeller() returns()
func (_Main *MainTransactor) ApproveBySeller(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "approveBySeller")
}

// ApproveBySeller is a paid mutator transaction binding the contract method 0x6cc332ce.
//
// Solidity: function approveBySeller() returns()
func (_Main *MainSession) ApproveBySeller() (*types.Transaction, error) {
	return _Main.Contract.ApproveBySeller(&_Main.TransactOpts)
}

// ApproveBySeller is a paid mutator transaction binding the contract method 0x6cc332ce.
//
// Solidity: function approveBySeller() returns()
func (_Main *MainTransactorSession) ApproveBySeller() (*types.Transaction, error) {
	return _Main.Contract.ApproveBySeller(&_Main.TransactOpts)
}

// ApproveByTrustee is a paid mutator transaction binding the contract method 0xf41d2d71.
//
// Solidity: function approveByTrustee() returns()
func (_Main *MainTransactor) ApproveByTrustee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "approveByTrustee")
}

// ApproveByTrustee is a paid mutator transaction binding the contract method 0xf41d2d71.
//
// Solidity: function approveByTrustee() returns()
func (_Main *MainSession) ApproveByTrustee() (*types.Transaction, error) {
	return _Main.Contract.ApproveByTrustee(&_Main.TransactOpts)
}

// ApproveByTrustee is a paid mutator transaction binding the contract method 0xf41d2d71.
//
// Solidity: function approveByTrustee() returns()
func (_Main *MainTransactorSession) ApproveByTrustee() (*types.Transaction, error) {
	return _Main.Contract.ApproveByTrustee(&_Main.TransactOpts)
}

// DisapproveByBuyer is a paid mutator transaction binding the contract method 0xdd257bde.
//
// Solidity: function disapproveByBuyer() returns()
func (_Main *MainTransactor) DisapproveByBuyer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "disapproveByBuyer")
}

// DisapproveByBuyer is a paid mutator transaction binding the contract method 0xdd257bde.
//
// Solidity: function disapproveByBuyer() returns()
func (_Main *MainSession) DisapproveByBuyer() (*types.Transaction, error) {
	return _Main.Contract.DisapproveByBuyer(&_Main.TransactOpts)
}

// DisapproveByBuyer is a paid mutator transaction binding the contract method 0xdd257bde.
//
// Solidity: function disapproveByBuyer() returns()
func (_Main *MainTransactorSession) DisapproveByBuyer() (*types.Transaction, error) {
	return _Main.Contract.DisapproveByBuyer(&_Main.TransactOpts)
}

// DisapproveBySeller is a paid mutator transaction binding the contract method 0x32fc28ec.
//
// Solidity: function disapproveBySeller() returns()
func (_Main *MainTransactor) DisapproveBySeller(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "disapproveBySeller")
}

// DisapproveBySeller is a paid mutator transaction binding the contract method 0x32fc28ec.
//
// Solidity: function disapproveBySeller() returns()
func (_Main *MainSession) DisapproveBySeller() (*types.Transaction, error) {
	return _Main.Contract.DisapproveBySeller(&_Main.TransactOpts)
}

// DisapproveBySeller is a paid mutator transaction binding the contract method 0x32fc28ec.
//
// Solidity: function disapproveBySeller() returns()
func (_Main *MainTransactorSession) DisapproveBySeller() (*types.Transaction, error) {
	return _Main.Contract.DisapproveBySeller(&_Main.TransactOpts)
}

// MainApprovalReceivedIterator is returned from FilterApprovalReceived and is used to iterate over the raw logs and unpacked data for ApprovalReceived events raised by the Main contract.
type MainApprovalReceivedIterator struct {
	Event *MainApprovalReceived // Event containing the contract specifics and raw log

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
func (it *MainApprovalReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainApprovalReceived)
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
		it.Event = new(MainApprovalReceived)
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
func (it *MainApprovalReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainApprovalReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainApprovalReceived represents a ApprovalReceived event raised by the Main contract.
type MainApprovalReceived struct {
	Approver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalReceived is a free log retrieval operation binding the contract event 0x135b662e7c460f01a8962dabbc952859694282ca1f1bfbacf24ebecc00160d35.
//
// Solidity: event ApprovalReceived(address indexed approver)
func (_Main *MainFilterer) FilterApprovalReceived(opts *bind.FilterOpts, approver []common.Address) (*MainApprovalReceivedIterator, error) {

	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "ApprovalReceived", approverRule)
	if err != nil {
		return nil, err
	}
	return &MainApprovalReceivedIterator{contract: _Main.contract, event: "ApprovalReceived", logs: logs, sub: sub}, nil
}

// WatchApprovalReceived is a free log subscription operation binding the contract event 0x135b662e7c460f01a8962dabbc952859694282ca1f1bfbacf24ebecc00160d35.
//
// Solidity: event ApprovalReceived(address indexed approver)
func (_Main *MainFilterer) WatchApprovalReceived(opts *bind.WatchOpts, sink chan<- *MainApprovalReceived, approver []common.Address) (event.Subscription, error) {

	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "ApprovalReceived", approverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainApprovalReceived)
				if err := _Main.contract.UnpackLog(event, "ApprovalReceived", log); err != nil {
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

// ParseApprovalReceived is a log parse operation binding the contract event 0x135b662e7c460f01a8962dabbc952859694282ca1f1bfbacf24ebecc00160d35.
//
// Solidity: event ApprovalReceived(address indexed approver)
func (_Main *MainFilterer) ParseApprovalReceived(log types.Log) (*MainApprovalReceived, error) {
	event := new(MainApprovalReceived)
	if err := _Main.contract.UnpackLog(event, "ApprovalReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainPayoutRequestedIterator is returned from FilterPayoutRequested and is used to iterate over the raw logs and unpacked data for PayoutRequested events raised by the Main contract.
type MainPayoutRequestedIterator struct {
	Event *MainPayoutRequested // Event containing the contract specifics and raw log

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
func (it *MainPayoutRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainPayoutRequested)
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
		it.Event = new(MainPayoutRequested)
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
func (it *MainPayoutRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainPayoutRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainPayoutRequested represents a PayoutRequested event raised by the Main contract.
type MainPayoutRequested struct {
	BuyerAccount  string
	SellerAccount string
	AmountINR     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPayoutRequested is a free log retrieval operation binding the contract event 0x4423fec52feed3aa8bd2a99c8211cfcc2672119920cb1d206da40301c641d16d.
//
// Solidity: event PayoutRequested(string buyerAccount, string sellerAccount, uint256 amountINR)
func (_Main *MainFilterer) FilterPayoutRequested(opts *bind.FilterOpts) (*MainPayoutRequestedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "PayoutRequested")
	if err != nil {
		return nil, err
	}
	return &MainPayoutRequestedIterator{contract: _Main.contract, event: "PayoutRequested", logs: logs, sub: sub}, nil
}

// WatchPayoutRequested is a free log subscription operation binding the contract event 0x4423fec52feed3aa8bd2a99c8211cfcc2672119920cb1d206da40301c641d16d.
//
// Solidity: event PayoutRequested(string buyerAccount, string sellerAccount, uint256 amountINR)
func (_Main *MainFilterer) WatchPayoutRequested(opts *bind.WatchOpts, sink chan<- *MainPayoutRequested) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "PayoutRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainPayoutRequested)
				if err := _Main.contract.UnpackLog(event, "PayoutRequested", log); err != nil {
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

// ParsePayoutRequested is a log parse operation binding the contract event 0x4423fec52feed3aa8bd2a99c8211cfcc2672119920cb1d206da40301c641d16d.
//
// Solidity: event PayoutRequested(string buyerAccount, string sellerAccount, uint256 amountINR)
func (_Main *MainFilterer) ParsePayoutRequested(log types.Log) (*MainPayoutRequested, error) {
	event := new(MainPayoutRequested)
	if err := _Main.contract.UnpackLog(event, "PayoutRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
