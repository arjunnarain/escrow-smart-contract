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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountINR\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"accountHolderName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"accountNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bankName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ifscCode\",\"type\":\"string\"}],\"internalType\":\"structEscrow.BankDetails\",\"name\":\"_buyerBankDetails\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"accountHolderName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"accountNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bankName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ifscCode\",\"type\":\"string\"}],\"internalType\":\"structEscrow.BankDetails\",\"name\":\"_sellerBankDetails\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"razorpayApiKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"razorpayApiSecret\",\"type\":\"string\"}],\"internalType\":\"structEscrow.PayoutAuth\",\"name\":\"_payoutAuth\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ApprovalReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ContractCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"buyerAccountHolderName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"buyerAccountNumber\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"buyerBankName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"buyerIfscCode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sellerAccountHolderName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sellerAccountNumber\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sellerBankName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sellerIfscCode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountINR\",\"type\":\"uint256\"}],\"name\":\"PayoutRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"amountINR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approveByBuyer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approveBySeller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approveByTrustee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyerApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyerBankDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"accountHolderName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"accountNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bankName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ifscCode\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"completeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentState\",\"outputs\":[{\"internalType\":\"enumEscrow.State\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disapproveByBuyer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disapproveBySeller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payoutAuth\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"razorpayApiKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"razorpayApiSecret\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellerApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellerBankDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"accountHolderName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"accountNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bankName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ifscCode\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustee\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trusteeApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516122a43803806122a483398181016040528101906100319190610560565b865f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508560015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508460025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550836003819055508260055f820151815f01908161010d9190610855565b5060208201518160010190816101239190610855565b5060408201518160020190816101399190610855565b50606082015181600301908161014f9190610855565b509050508160095f820151815f0190816101699190610855565b50602082015181600101908161017f9190610855565b5060408201518160020190816101959190610855565b5060608201518160030190816101ab9190610855565b5090505080600d5f820151815f0190816101c59190610855565b5060208201518160010190816101db9190610855565b509050505f60045f6101000a81548160ff0219169083600181111561020357610202610924565b5b021790555050505050505050610951565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61024e82610225565b9050919050565b61025e81610244565b8114610268575f80fd5b50565b5f8151905061027981610255565b92915050565b5f819050919050565b6102918161027f565b811461029b575f80fd5b50565b5f815190506102ac81610288565b92915050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6102fc826102b6565b810181811067ffffffffffffffff8211171561031b5761031a6102c6565b5b80604052505050565b5f61032d610214565b905061033982826102f3565b919050565b5f80fd5b5f80fd5b5f80fd5b5f67ffffffffffffffff821115610364576103636102c6565b5b61036d826102b6565b9050602081019050919050565b8281835e5f83830152505050565b5f61039a6103958461034a565b610324565b9050828152602081018484840111156103b6576103b5610346565b5b6103c184828561037a565b509392505050565b5f82601f8301126103dd576103dc610342565b5b81516103ed848260208601610388565b91505092915050565b5f6080828403121561040b5761040a6102b2565b5b6104156080610324565b90505f82015167ffffffffffffffff8111156104345761043361033e565b5b610440848285016103c9565b5f83015250602082015167ffffffffffffffff8111156104635761046261033e565b5b61046f848285016103c9565b602083015250604082015167ffffffffffffffff8111156104935761049261033e565b5b61049f848285016103c9565b604083015250606082015167ffffffffffffffff8111156104c3576104c261033e565b5b6104cf848285016103c9565b60608301525092915050565b5f604082840312156104f0576104ef6102b2565b5b6104fa6040610324565b90505f82015167ffffffffffffffff8111156105195761051861033e565b5b610525848285016103c9565b5f83015250602082015167ffffffffffffffff8111156105485761054761033e565b5b610554848285016103c9565b60208301525092915050565b5f805f805f805f60e0888a03121561057b5761057a61021d565b5b5f6105888a828b0161026b565b97505060206105998a828b0161026b565b96505060406105aa8a828b0161026b565b95505060606105bb8a828b0161029e565b945050608088015167ffffffffffffffff8111156105dc576105db610221565b5b6105e88a828b016103f6565b93505060a088015167ffffffffffffffff81111561060957610608610221565b5b6106158a828b016103f6565b92505060c088015167ffffffffffffffff81111561063657610635610221565b5b6106428a828b016104db565b91505092959891949750929550565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061069f57607f821691505b6020821081036106b2576106b161065b565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026107147fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826106d9565b61071e86836106d9565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61075961075461074f8461027f565b610736565b61027f565b9050919050565b5f819050919050565b6107728361073f565b61078661077e82610760565b8484546106e5565b825550505050565b5f90565b61079a61078e565b6107a5818484610769565b505050565b5b818110156107c8576107bd5f82610792565b6001810190506107ab565b5050565b601f82111561080d576107de816106b8565b6107e7846106ca565b810160208510156107f6578190505b61080a610802856106ca565b8301826107aa565b50505b505050565b5f82821c905092915050565b5f61082d5f1984600802610812565b1980831691505092915050565b5f610845838361081e565b9150826002028217905092915050565b61085e82610651565b67ffffffffffffffff811115610877576108766102c6565b5b6108818254610688565b61088c8282856107cc565b5f60209050601f8311600181146108bd575f84156108ab578287015190505b6108b5858261083a565b86555061091c565b601f1984166108cb866106b8565b5f5b828110156108f2578489015182556001820191506020850194506020810190506108cd565b8683101561090f578489015161090b601f89168261081e565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6119468061095e5f395ff3fe608060405234801561000f575f80fd5b5060043610610109575f3560e01c80637150d8ae116100a0578063dd257bde1161006f578063dd257bde14610240578063f41d2d711461024a578063fa65600a14610254578063fdf97cb214610272578063ff80a3891461029057610109565b80637150d8ae146101d8578063b1084e7f146101f6578063b4fedc1114610217578063da8494761461022157610109565b80631afcb321116100dc5780631afcb3211461018557806332fc28ec146101a35780636b566475146101ad5780636cc332ce146101ce57610109565b806308551a531461010d5780630a1792cc1461012b5780630c3f6acf14610149578063183f66b614610167575b5f80fd5b61011561029a565b6040516101229190611252565b60405180910390f35b6101336102bf565b6040516101409190611285565b60405180910390f35b6101516102d1565b60405161015e9190611311565b60405180910390f35b61016f6102e3565b60405161017c9190611285565b60405180910390f35b61018d6102f6565b60405161019a9190611285565b60405180910390f35b6101ab610309565b005b6101b561046b565b6040516101c5949392919061139a565b60405180910390f35b6101d66106a0565b005b6101e061080b565b6040516101ed9190611252565b60405180910390f35b6101fe61082e565b60405161020e949392919061139a565b60405180910390f35b61021f610a63565b005b610229610bbe565b6040516102379291906113f9565b60405180910390f35b610248610cdb565b005b610252610e3a565b005b61025c610fa5565b6040516102699190611446565b60405180910390f35b61027a610fab565b6040516102879190611252565b60405180910390f35b610298610fd0565b005b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600f5f9054906101000a900460ff1681565b60045f9054906101000a900460ff1681565b600f60019054906101000a900460ff1681565b600f60029054906101000a900460ff1681565b5f600181111561031c5761031b61129e565b5b60045f9054906101000a900460ff16600181111561033d5761033c61129e565b5b1461037d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610374906114a9565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461040c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040390611511565b60405180910390fd5b5f600f60016101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f135b662e7c460f01a8962dabbc952859694282ca1f1bfbacf24ebecc00160d3560405160405180910390a2565b6009805f01805461047b9061155c565b80601f01602080910402602001604051908101604052809291908181526020018280546104a79061155c565b80156104f25780601f106104c9576101008083540402835291602001916104f2565b820191905f5260205f20905b8154815290600101906020018083116104d557829003601f168201915b5050505050908060010180546105079061155c565b80601f01602080910402602001604051908101604052809291908181526020018280546105339061155c565b801561057e5780601f106105555761010080835404028352916020019161057e565b820191905f5260205f20905b81548152906001019060200180831161056157829003601f168201915b5050505050908060020180546105939061155c565b80601f01602080910402602001604051908101604052809291908181526020018280546105bf9061155c565b801561060a5780601f106105e15761010080835404028352916020019161060a565b820191905f5260205f20905b8154815290600101906020018083116105ed57829003601f168201915b50505050509080600301805461061f9061155c565b80601f016020809104026020016040519081016040528092919081815260200182805461064b9061155c565b80156106965780601f1061066d57610100808354040283529160200191610696565b820191905f5260205f20905b81548152906001019060200180831161067957829003601f168201915b5050505050905084565b5f60018111156106b3576106b261129e565b5b60045f9054906101000a900460ff1660018111156106d4576106d361129e565b5b14610714576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070b906114a9565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079a906115d6565b60405180910390fd5b6001600f60016101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f135b662e7c460f01a8962dabbc952859694282ca1f1bfbacf24ebecc00160d3560405160405180910390a2610809611138565b565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6005805f01805461083e9061155c565b80601f016020809104026020016040519081016040528092919081815260200182805461086a9061155c565b80156108b55780601f1061088c576101008083540402835291602001916108b5565b820191905f5260205f20905b81548152906001019060200180831161089857829003601f168201915b5050505050908060010180546108ca9061155c565b80601f01602080910402602001604051908101604052809291908181526020018280546108f69061155c565b80156109415780601f1061091857610100808354040283529160200191610941565b820191905f5260205f20905b81548152906001019060200180831161092457829003601f168201915b5050505050908060020180546109569061155c565b80601f01602080910402602001604051908101604052809291908181526020018280546109829061155c565b80156109cd5780601f106109a4576101008083540402835291602001916109cd565b820191905f5260205f20905b8154815290600101906020018083116109b057829003601f168201915b5050505050908060030180546109e29061155c565b80601f0160208091040260200160405190810160405280929190818152602001828054610a0e9061155c565b8015610a595780601f10610a3057610100808354040283529160200191610a59565b820191905f5260205f20905b815481529060010190602001808311610a3c57829003601f168201915b5050505050905084565b5f6001811115610a7657610a7561129e565b5b60045f9054906101000a900460ff166001811115610a9757610a9661129e565b5b14610ad7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ace906114a9565b60405180910390fd5b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b66576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b5d90611664565b60405180910390fd5b600160045f6101000a81548160ff02191690836001811115610b8b57610b8a61129e565b5b02179055507f21af1b35312e092d631654105d02329ce4aa3a7aa58f89d2e61e3fa80874dd8c60405160405180910390a1565b600d805f018054610bce9061155c565b80601f0160208091040260200160405190810160405280929190818152602001828054610bfa9061155c565b8015610c455780601f10610c1c57610100808354040283529160200191610c45565b820191905f5260205f20905b815481529060010190602001808311610c2857829003601f168201915b505050505090806001018054610c5a9061155c565b80601f0160208091040260200160405190810160405280929190818152602001828054610c869061155c565b8015610cd15780601f10610ca857610100808354040283529160200191610cd1565b820191905f5260205f20905b815481529060010190602001808311610cb457829003601f168201915b5050505050905082565b5f6001811115610cee57610ced61129e565b5b60045f9054906101000a900460ff166001811115610d0f57610d0e61129e565b5b14610d4f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d46906114a9565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ddc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dd3906116cc565b60405180910390fd5b5f600f5f6101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f135b662e7c460f01a8962dabbc952859694282ca1f1bfbacf24ebecc00160d3560405160405180910390a2565b5f6001811115610e4d57610e4c61129e565b5b60045f9054906101000a900460ff166001811115610e6e57610e6d61129e565b5b14610eae576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ea5906114a9565b60405180910390fd5b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f3d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f3490611734565b60405180910390fd5b6001600f60026101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f135b662e7c460f01a8962dabbc952859694282ca1f1bfbacf24ebecc00160d3560405160405180910390a2610fa3611138565b565b60035481565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f6001811115610fe357610fe261129e565b5b60045f9054906101000a900460ff1660018111156110045761100361129e565b5b14611044576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103b906114a9565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146110d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110c89061179c565b60405180910390fd5b6001600f5f6101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f135b662e7c460f01a8962dabbc952859694282ca1f1bfbacf24ebecc00160d3560405160405180910390a2611136611138565b565b600f5f9054906101000a900460ff16801561115f5750600f60019054906101000a900460ff165b806111a45750600f60029054906101000a900460ff1680156111a35750600f5f9054906101000a900460ff16806111a25750600f60019054906101000a900460ff165b5b5b15611211577fcb88be8b613994e3054133de733e8b7af2e0a3613877928da103d77275bdcb5e60055f0160056001016005600201600560030160095f016009600101600960020160096003016003546040516112089998979695949392919061184d565b60405180910390a15b565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61123c82611213565b9050919050565b61124c81611232565b82525050565b5f6020820190506112655f830184611243565b92915050565b5f8115159050919050565b61127f8161126b565b82525050565b5f6020820190506112985f830184611276565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600281106112dc576112db61129e565b5b50565b5f8190506112ec826112cb565b919050565b5f6112fb826112df565b9050919050565b61130b816112f1565b82525050565b5f6020820190506113245f830184611302565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61136c8261132a565b6113768185611334565b9350611386818560208601611344565b61138f81611352565b840191505092915050565b5f6080820190508181035f8301526113b28187611362565b905081810360208301526113c68186611362565b905081810360408301526113da8185611362565b905081810360608301526113ee8184611362565b905095945050505050565b5f6040820190508181035f8301526114118185611362565b905081810360208301526114258184611362565b90509392505050565b5f819050919050565b6114408161142e565b82525050565b5f6020820190506114595f830184611437565b92915050565b7f436f6e7472616374206973206e6f206c6f6e67657220616374697665000000005f82015250565b5f611493601c83611334565b915061149e8261145f565b602082019050919050565b5f6020820190508181035f8301526114c081611487565b9050919050565b7f4f6e6c792073656c6c65722063616e20646973617070726f76650000000000005f82015250565b5f6114fb601a83611334565b9150611506826114c7565b602082019050919050565b5f6020820190508181035f830152611528816114ef565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061157357607f821691505b6020821081036115865761158561152f565b5b50919050565b7f4f6e6c792073656c6c65722063616e20617070726f76650000000000000000005f82015250565b5f6115c0601783611334565b91506115cb8261158c565b602082019050919050565b5f6020820190508181035f8301526115ed816115b4565b9050919050565b7f4f6e6c7920747275737465652063616e20636f6d706c6574652074686520636f5f8201527f6e74726163740000000000000000000000000000000000000000000000000000602082015250565b5f61164e602683611334565b9150611659826115f4565b604082019050919050565b5f6020820190508181035f83015261167b81611642565b9050919050565b7f4f6e6c792062757965722063616e20646973617070726f7665000000000000005f82015250565b5f6116b6601983611334565b91506116c182611682565b602082019050919050565b5f6020820190508181035f8301526116e3816116aa565b9050919050565b7f4f6e6c7920747275737465652063616e20617070726f766500000000000000005f82015250565b5f61171e601883611334565b9150611729826116ea565b602082019050919050565b5f6020820190508181035f83015261174b81611712565b9050919050565b7f4f6e6c792062757965722063616e20617070726f7665000000000000000000005f82015250565b5f611786601683611334565b915061179182611752565b602082019050919050565b5f6020820190508181035f8301526117b38161177a565b9050919050565b5f819050815f5260205f209050919050565b5f81546117d88161155c565b6117e28186611334565b9450600182165f81146117fc576001811461181257611844565b60ff198316865281151560200286019350611844565b61181b856117ba565b5f5b8381101561183c5781548189015260018201915060208101905061181d565b808801955050505b50505092915050565b5f610120820190508181035f830152611866818c6117cc565b9050818103602083015261187a818b6117cc565b9050818103604083015261188e818a6117cc565b905081810360608301526118a281896117cc565b905081810360808301526118b681886117cc565b905081810360a08301526118ca81876117cc565b905081810360c08301526118de81866117cc565b905081810360e08301526118f281856117cc565b9050611902610100830184611437565b9a995050505050505050505056fea26469706673582212209fb1682f8ced157dc3a2bce1e420fe231097cf745846f460424d1c4c537af9bd64736f6c634300081a0033",
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

// CurrentState is a free data retrieval call binding the contract method 0x0c3f6acf.
//
// Solidity: function currentState() view returns(uint8)
func (_Main *MainCaller) CurrentState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "currentState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CurrentState is a free data retrieval call binding the contract method 0x0c3f6acf.
//
// Solidity: function currentState() view returns(uint8)
func (_Main *MainSession) CurrentState() (uint8, error) {
	return _Main.Contract.CurrentState(&_Main.CallOpts)
}

// CurrentState is a free data retrieval call binding the contract method 0x0c3f6acf.
//
// Solidity: function currentState() view returns(uint8)
func (_Main *MainCallerSession) CurrentState() (uint8, error) {
	return _Main.Contract.CurrentState(&_Main.CallOpts)
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

// CompleteContract is a paid mutator transaction binding the contract method 0xb4fedc11.
//
// Solidity: function completeContract() returns()
func (_Main *MainTransactor) CompleteContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "completeContract")
}

// CompleteContract is a paid mutator transaction binding the contract method 0xb4fedc11.
//
// Solidity: function completeContract() returns()
func (_Main *MainSession) CompleteContract() (*types.Transaction, error) {
	return _Main.Contract.CompleteContract(&_Main.TransactOpts)
}

// CompleteContract is a paid mutator transaction binding the contract method 0xb4fedc11.
//
// Solidity: function completeContract() returns()
func (_Main *MainTransactorSession) CompleteContract() (*types.Transaction, error) {
	return _Main.Contract.CompleteContract(&_Main.TransactOpts)
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

// MainContractCompletedIterator is returned from FilterContractCompleted and is used to iterate over the raw logs and unpacked data for ContractCompleted events raised by the Main contract.
type MainContractCompletedIterator struct {
	Event *MainContractCompleted // Event containing the contract specifics and raw log

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
func (it *MainContractCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainContractCompleted)
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
		it.Event = new(MainContractCompleted)
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
func (it *MainContractCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainContractCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainContractCompleted represents a ContractCompleted event raised by the Main contract.
type MainContractCompleted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterContractCompleted is a free log retrieval operation binding the contract event 0x21af1b35312e092d631654105d02329ce4aa3a7aa58f89d2e61e3fa80874dd8c.
//
// Solidity: event ContractCompleted()
func (_Main *MainFilterer) FilterContractCompleted(opts *bind.FilterOpts) (*MainContractCompletedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "ContractCompleted")
	if err != nil {
		return nil, err
	}
	return &MainContractCompletedIterator{contract: _Main.contract, event: "ContractCompleted", logs: logs, sub: sub}, nil
}

// WatchContractCompleted is a free log subscription operation binding the contract event 0x21af1b35312e092d631654105d02329ce4aa3a7aa58f89d2e61e3fa80874dd8c.
//
// Solidity: event ContractCompleted()
func (_Main *MainFilterer) WatchContractCompleted(opts *bind.WatchOpts, sink chan<- *MainContractCompleted) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "ContractCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainContractCompleted)
				if err := _Main.contract.UnpackLog(event, "ContractCompleted", log); err != nil {
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

// ParseContractCompleted is a log parse operation binding the contract event 0x21af1b35312e092d631654105d02329ce4aa3a7aa58f89d2e61e3fa80874dd8c.
//
// Solidity: event ContractCompleted()
func (_Main *MainFilterer) ParseContractCompleted(log types.Log) (*MainContractCompleted, error) {
	event := new(MainContractCompleted)
	if err := _Main.contract.UnpackLog(event, "ContractCompleted", log); err != nil {
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
	BuyerAccountHolderName  string
	BuyerAccountNumber      string
	BuyerBankName           string
	BuyerIfscCode           string
	SellerAccountHolderName string
	SellerAccountNumber     string
	SellerBankName          string
	SellerIfscCode          string
	AmountINR               *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterPayoutRequested is a free log retrieval operation binding the contract event 0xcb88be8b613994e3054133de733e8b7af2e0a3613877928da103d77275bdcb5e.
//
// Solidity: event PayoutRequested(string buyerAccountHolderName, string buyerAccountNumber, string buyerBankName, string buyerIfscCode, string sellerAccountHolderName, string sellerAccountNumber, string sellerBankName, string sellerIfscCode, uint256 amountINR)
func (_Main *MainFilterer) FilterPayoutRequested(opts *bind.FilterOpts) (*MainPayoutRequestedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "PayoutRequested")
	if err != nil {
		return nil, err
	}
	return &MainPayoutRequestedIterator{contract: _Main.contract, event: "PayoutRequested", logs: logs, sub: sub}, nil
}

// WatchPayoutRequested is a free log subscription operation binding the contract event 0xcb88be8b613994e3054133de733e8b7af2e0a3613877928da103d77275bdcb5e.
//
// Solidity: event PayoutRequested(string buyerAccountHolderName, string buyerAccountNumber, string buyerBankName, string buyerIfscCode, string sellerAccountHolderName, string sellerAccountNumber, string sellerBankName, string sellerIfscCode, uint256 amountINR)
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

// ParsePayoutRequested is a log parse operation binding the contract event 0xcb88be8b613994e3054133de733e8b7af2e0a3613877928da103d77275bdcb5e.
//
// Solidity: event PayoutRequested(string buyerAccountHolderName, string buyerAccountNumber, string buyerBankName, string buyerIfscCode, string sellerAccountHolderName, string sellerAccountNumber, string sellerBankName, string sellerIfscCode, uint256 amountINR)
func (_Main *MainFilterer) ParsePayoutRequested(log types.Log) (*MainPayoutRequested, error) {
	event := new(MainPayoutRequested)
	if err := _Main.contract.UnpackLog(event, "PayoutRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
