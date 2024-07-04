package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
)

var (
	client                      *ethclient.Client
	privateKey                  *ecdsa.PrivateKey
	lastDeployedContractAddress common.Address
)

type CreateEscrowRequest struct {
	BuyerAddress   string            `json:"buyerAddress"`
	SellerAddress  string            `json:"sellerAddress"`
	TrusteeAddress string            `json:"trusteeAddress"`
	AmountINR      uint64            `json:"amountINR"`
	BuyerBank      EscrowBankDetails `json:"buyerBank"`
	SellerBank     EscrowBankDetails `json:"sellerBank"`
	PayoutAuth     EscrowPayoutAuth  `json:"payoutAuth"`
}

type ApproveRequest struct {
	ContractAddress string `json:"contractAddress"`
	Role            string `json:"role"`
	PrivateKey      string `json:"privateKey"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file. Using environment variables.")
	}

	ganacheURL := getEnv("GANACHE_URL", "http://localhost:8545")
	privateKeyHex := getEnv("PRIVATE_KEY", "")
	if privateKeyHex == "" {
		log.Fatal("PRIVATE_KEY environment variable is not set")
	}

	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")

	client, err = ethclient.Dial(ganacheURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	privateKey, err = crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	router := gin.Default()
	router.POST("/createEscrow", handleCreateEscrow)
	router.POST("/approvePayout", handleApprovePayout)
	router.POST("/disapprove", handleDisapprove)
	router.Run(":3000")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func handleCreateEscrow(c *gin.Context) {
	var createRequest CreateEscrowRequest
	if err := c.ShouldBindJSON(&createRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the chain ID from the connected client
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Printf("Failed to get chain ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get chain ID"})
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Printf("Failed to create transactor: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transactor"})
		return
	}

	auth.GasLimit = 30000000 // Adjust this value as needed

	buyer := common.HexToAddress(createRequest.BuyerAddress)
	seller := common.HexToAddress(createRequest.SellerAddress)
	trustee := common.HexToAddress(createRequest.TrusteeAddress)
	amountINR := new(big.Int).SetUint64(createRequest.AmountINR)

	log.Printf("Deploying contract with parameters:\n")
	log.Printf("Chain ID: %d\n", chainID)
	log.Printf("Buyer: %s\n", buyer.Hex())
	log.Printf("Seller: %s\n", seller.Hex())
	log.Printf("Trustee: %s\n", trustee.Hex())
	log.Printf("Amount INR: %s\n", amountINR.String())
	log.Printf("Buyer Bank Details: %+v\n", createRequest.BuyerBank)
	log.Printf("Seller Bank Details: %+v\n", createRequest.SellerBank)
	log.Printf("Payout Auth: %+v\n", createRequest.PayoutAuth)

	// Ensure all parameters are correct and properly formatted
	if createRequest.BuyerAddress == "" || createRequest.SellerAddress == "" || createRequest.TrusteeAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address provided"})
		return
	}
	if createRequest.AmountINR == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Amount INR must be greater than zero"})
		return
	}

	address, tx, _, err := DeployMain(
		auth,
		client,
		buyer,
		seller,
		trustee,
		amountINR,
		createRequest.BuyerBank,
		createRequest.SellerBank,
		createRequest.PayoutAuth,
	)

	if err != nil {
		log.Printf("Failed to deploy contract: %v", err)
		if tx != nil {
			receipt, receiptErr := client.TransactionReceipt(context.Background(), tx.Hash())
			if receiptErr != nil {
				log.Printf("Failed to get transaction receipt: %v", receiptErr)
			} else {
				log.Printf("Transaction receipt: Status: %d, Gas Used: %d", receipt.Status, receipt.GasUsed)
			}
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deploy contract"})
		return
	}

	log.Printf("Contract deployed at address: %s", address.Hex())
	log.Printf("Transaction hash: %s", tx.Hash().Hex())

	// Store the contract address
	lastDeployedContractAddress = common.HexToAddress("0xAeDC40bf670996912bc3E2e7a9592984eC78467e")
	// Verify the deployed contract
	err = verifyEscrowContract(lastDeployedContractAddress)
	if err != nil {
		log.Printf("Contract verification failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Contract verification failed"})
		return
	}

	log.Printf("Contract verified successfully")

	c.JSON(http.StatusOK, gin.H{
		"status":          "escrow contract deployed and verified",
		"contractAddress": address.Hex(),
		"transactionHash": tx.Hash().Hex(),
	})
}

func verifyEscrowContract(contractAddress common.Address) error {
	instance, err := NewMain(contractAddress, client)
	if err != nil {
		log.Printf("Failed to instantiate contract: %v", err)
		return fmt.Errorf("failed to instantiate contract: %v", err)
	}

	buyer, err := instance.Buyer(&bind.CallOpts{})
	if err != nil {
		log.Printf("Failed to get buyer: %v", err)
		return fmt.Errorf("failed to get buyer: %v", err)
	}
	log.Printf("Verified Buyer address: %s", buyer.Hex())

	seller, err := instance.Seller(&bind.CallOpts{})
	if err != nil {
		log.Printf("Failed to get seller: %v", err)
		return fmt.Errorf("failed to get seller: %v", err)
	}
	log.Printf("Verified Seller address: %s", seller.Hex())

	trustee, err := instance.Trustee(&bind.CallOpts{})
	if err != nil {
		log.Printf("Failed to get trustee: %v", err)
		return fmt.Errorf("failed to get trustee: %v", err)
	}
	log.Printf("Verified Trustee address: %s", trustee.Hex())

	amount, err := instance.AmountINR(&bind.CallOpts{})
	if err != nil {
		log.Printf("Failed to get amount: %v", err)
		return fmt.Errorf("failed to get amount: %v", err)
	}
	log.Printf("Verified Amount INR: %s", amount.String())

	return nil
}

func logDeploymentDetails(auth *bind.TransactOpts, buyer, seller, trustee common.Address, amountINR *big.Int) {
	log.Printf("Deployment details:")
	log.Printf("From: %s", auth.From.Hex())
	log.Printf("Buyer: %s", buyer.Hex())
	log.Printf("Seller: %s", seller.Hex())
	log.Printf("Trustee: %s", trustee.Hex())
	log.Printf("Amount INR: %s", amountINR.String())
	log.Printf("Gas Price: %s", auth.GasPrice.String())
	log.Printf("Gas Limit: %d", auth.GasLimit)
	log.Printf("Nonce: %d", auth.Nonce.Uint64())
}

func handleApprovePayout(c *gin.Context) {
	var approveRequest ApproveRequest
	if err := c.ShouldBindJSON(&approveRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert the private key from string to ecdsa.PrivateKey
	privateKey, err := crypto.HexToECDSA(approveRequest.PrivateKey)
	if err != nil {
		log.Printf("Failed to load private key: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid private key"})
		return
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Printf("Failed to get chain ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get chain ID"})
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Printf("Failed to create transactor: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transactor"})
		return
	}

	contractAddress := common.HexToAddress(approveRequest.ContractAddress)
	instance, err := NewMain(contractAddress, client)
	if err != nil {
		log.Printf("Failed to instantiate contract: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to instantiate contract"})
		return
	}

	var tx *types.Transaction
	switch approveRequest.Role {
	case "buyer":
		tx, err = instance.ApproveByBuyer(auth)
	case "seller":
		tx, err = instance.ApproveBySeller(auth)
	case "trustee":
		tx, err = instance.ApproveByTrustee(auth)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role"})
		return
	}

	if err != nil {
		log.Printf("Failed to approve payout: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to approve payout"})
		return
	}

	log.Printf("Approval transaction hash: %s", tx.Hash().Hex())

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Printf("Failed to wait for transaction to be mined: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to confirm approval"})
		return
	}

	// Check for PayoutRequested event
	defer checkForPayoutRequestedEvent(contractAddress, client, receipt.BlockNumber)
	c.JSON(http.StatusOK, gin.H{"status": "payout approved", "transactionHash": tx.Hash().Hex()})
}

func handleDisapprove(c *gin.Context) {
	var disapproveRequest ApproveRequest
	if err := c.ShouldBindJSON(&disapproveRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert the private key from string to ecdsa.PrivateKey
	privateKey, err := crypto.HexToECDSA(disapproveRequest.PrivateKey)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Printf("Failed to get chain ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get chain ID"})
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Printf("Failed to create transactor: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transactor"})
		return
	}

	contractAddress := common.HexToAddress(disapproveRequest.ContractAddress)
	instance, err := NewMain(contractAddress, client)
	if err != nil {
		log.Printf("Failed to instantiate contract: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to instantiate contract"})
		return
	}

	var tx *types.Transaction
	switch disapproveRequest.Role {
	case "buyer":
		tx, err = instance.DisapproveByBuyer(auth)
	case "seller":
		tx, err = instance.DisapproveBySeller(auth)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role"})
		return
	}

	if err != nil {
		log.Printf("Failed to disapprove: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to disapprove"})
		return
	}

	log.Printf("Disapproval transaction hash: %s", tx.Hash().Hex())

	c.JSON(http.StatusOK, gin.H{"status": "disapproval registered", "transactionHash": tx.Hash().Hex()})
}

func checkForPayoutRequestedEvent(contractAddress common.Address, client *ethclient.Client, fromBlock *big.Int) error {
	instance, err := NewMain(contractAddress, client)
	if err != nil {
		return fmt.Errorf("failed to instantiate contract: %v", err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		FromBlock: fromBlock,
		ToBlock:   nil, // nil means the latest block
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return fmt.Errorf("failed to filter logs: %v", err)
	}

	for _, vLog := range logs {
		event, err := instance.ParsePayoutRequested(vLog)
		if err != nil {
			log.Printf("Failed to parse PayoutRequested event: %v", err)
			continue
		}

		log.Printf("PayoutRequested event received!")
		log.Printf("Buyer Account: %s", event.BuyerAccount)
		log.Printf("Seller Account: %s", event.SellerAccount)
		log.Printf("Amount INR: %s", event.AmountINR.String())

		processPayout(event.BuyerAccount, event.SellerAccount, event.AmountINR)
	}

	return nil
}

func processPayout(buyerAccount, sellerAccount string, amountINR *big.Int) {
	log.Printf("Processing payout: From %s to %s, Amount: %s INR", buyerAccount, sellerAccount, amountINR.String())

	// Convert big.Int to float64 for Razorpay
	//amountFloat := new(big.Float).SetInt(amountINR)
	//amount, _ := amountFloat.Float64()

	// Initialize Razorpay client
	// client := razorpay.NewClient(os.Getenv("RAZORPAY_KEY_ID"), os.Getenv("RAZORPAY_SECRET_KEY"))

	//data := map[string]interface{}{
	//	"account_number": sellerAccount,
	//	"amount":         amount * 100, // Razorpay expects amount in paise
	//	"currency":       "INR",
	//	"mode":           "IMPS",
	//	"purpose":        "payout",
	//	"queue_if_low_balance": true,
	//	"reference_id":   fmt.Sprintf("order_%s", amountINR.String()),
	//	"narration":      "Escrow Payout",
	//}

	//body, err := client.Payout.Create(data, nil)
	//if err != nil {
	//	log.Printf("Failed to create payout: %v", err)
	//	// Here you might want to implement retry logic or notify an admin
	//	return
	//}

	log.Printf("Payout successful. Payout ID")

	// You might want to update your database or the smart contract here to record the successful payout
	//updatePayoutStatus(buyerAccount, sellerAccount, body["id"].(string))
}

func updatePayoutStatus(buyerAccount, sellerAccount, payoutID string) {
	// This function would update your database or smart contract with the payout status
	// Implementation depends on your specific requirements
	log.Printf("Updating payout status for buyer %s, seller %s, payout ID %s", buyerAccount, sellerAccount, payoutID)
	// Add your implementation here
}
