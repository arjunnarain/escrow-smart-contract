package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	client          *ethclient.Client
	contractAddress common.Address
	privateKey      *ecdsa.PrivateKey
	lastBlockNumber uint64
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
	Role string `json:"role"` // buyer, seller, trustee
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ganacheURL := os.Getenv("GANACHE_URL")
	privateKeyHex := os.Getenv("PRIVATE_KEY")

	// Remove "0x" prefix if it exists
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")

	client, err = ethclient.Dial(ganacheURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	privateKey, err = crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// Initialize lastBlockNumber
	lastBlockNumber, err = client.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("Failed to get latest block number: %v", err)
	}

	go pollForLogs()

	router := gin.Default()
	router.POST("/createEscrow", handleCreateEscrow)
	router.POST("/approvePayout", handleApprovePayout)
	router.POST("/disapprove", handleDisapprove)
	router.Run(":3000")
}

func pollForLogs() {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fetchNewLogs()
		}
	}
}

func fetchNewLogs() {
	ctx := context.Background()
	latestBlock, err := client.BlockNumber(ctx)
	if err != nil {
		log.Printf("Failed to get latest block number: %v", err)
		return
	}

	if latestBlock <= lastBlockNumber {
		return // No new blocks
	}

	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(lastBlockNumber + 1),
		ToBlock:   new(big.Int).SetUint64(latestBlock),
		Addresses: []common.Address{contractAddress},
	}

	logs, err := client.FilterLogs(ctx, query)
	if err != nil {
		log.Printf("Failed to filter logs: %v", err)
		return
	}

	for _, vLog := range logs {
		handleLog(vLog)
	}

	lastBlockNumber = latestBlock
}

func handleLog(vLog types.Log) {
	contractAbi, err := MainMetaData.GetAbi()
	if err != nil {
		log.Printf("Failed to get ABI: %v", err)
		return
	}

	event := struct {
		BuyerAccount  string
		SellerAccount string
		AmountINR     *big.Int
	}{}

	err = contractAbi.UnpackIntoInterface(&event, "PayoutRequested", vLog.Data)
	if err != nil {
		log.Printf("Failed to unpack log data: %v", err)
		return
	}

	processPayout(event)
}

func handleCreateEscrow(c *gin.Context) {
	var createRequest CreateEscrowRequest
	if err := c.ShouldBindJSON(&createRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(5777)) // Ganache's default network ID
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	buyer := common.HexToAddress(createRequest.BuyerAddress)
	seller := common.HexToAddress(createRequest.SellerAddress)
	trustee := common.HexToAddress(createRequest.TrusteeAddress)
	amountINR := new(big.Int).SetUint64(createRequest.AmountINR)

	instanceAddress, tx, _, err := DeployMain(
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
		log.Fatalf("Failed to deploy contract: %v", err)
	}

	contractAddress = instanceAddress

	log.Printf("Contract deployed at address: %s", instanceAddress.Hex())
	log.Printf("Transaction hash: %s", tx.Hash().Hex())

	c.JSON(http.StatusOK, gin.H{"status": "escrow created", "contractAddress": instanceAddress.Hex()})
}

func handleApprovePayout(c *gin.Context) {
	var approveRequest ApproveRequest
	if err := c.ShouldBindJSON(&approveRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(5777)) // Ganache's default network ID
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	instance, err := NewMain(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to load contract: %v", err)
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
		log.Fatalf("Failed to approve payout: %v", err)
	}

	log.Printf("Transaction hash: %s", tx.Hash().Hex())

	c.JSON(http.StatusOK, gin.H{"status": "payout approved"})
}

func handleDisapprove(c *gin.Context) {
	var approveRequest ApproveRequest
	if err := c.ShouldBindJSON(&approveRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(5777)) // Ganache's default network ID
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	instance, err := NewMain(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to load contract: %v", err)
	}

	var tx *types.Transaction
	switch approveRequest.Role {
	case "buyer":
		tx, err = instance.DisapproveByBuyer(auth)
	case "seller":
		tx, err = instance.DisapproveBySeller(auth)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role"})
		return
	}

	if err != nil {
		log.Fatalf("Failed to disapprove payout: %v", err)
	}

	log.Printf("Transaction hash: %s", tx.Hash().Hex())

	c.JSON(http.StatusOK, gin.H{"status": "disapproval registered"})
}

func processPayout(payoutRequest struct {
	BuyerAccount  string
	SellerAccount string
	AmountINR     *big.Int
}) {
	payoutData := fmt.Sprintf(`{
        "account_number": "%s",
        "amount": %s,
        "currency": "INR",
        "mode": "IMPS",
        "purpose": "payout",
        "queue_if_low_balance": true,
        "reference_id": "order_123",
        "narration": "Payout for order",
        "notes": {
            "sender_name": "Buyer Name",
            "receiver_name": "Seller Name"
        }
    }`, payoutRequest.SellerAccount, payoutRequest.AmountINR.String())

	req, err := http.NewRequest("POST", "https://api.razorpay.com/v1/payouts", strings.NewReader(payoutData))
	if err != nil {
		log.Printf("Failed to create payout request: %v", err)
		return
	}

	req.SetBasicAuth(os.Getenv("RAZORPAY_KEY_ID"), os.Getenv("RAZORPAY_KEY_SECRET"))
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Printf("Failed to execute payout request: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Payout request failed with status: %v", resp.Status)
	} else {
		log.Println("Payout request successful")
	}
}
