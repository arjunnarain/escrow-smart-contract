package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	client                      *ethclient.Client
	buyerPrivateKey             *ecdsa.PrivateKey
	trusteePrivateKey           *ecdsa.PrivateKey
	sellerPrivateKey            *ecdsa.PrivateKey
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
	PrivateKey      string `json:"buyerPrivateKey"`
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

	trusteePrivateKeyHex := getEnv("TRUSTEE_PRIVATE_KEY", "")
	if trusteePrivateKeyHex == "" {
		log.Fatal("TRUSTEE PRIVATE_KEY environment variable is not set")
	}

	sellerPrivateKeyHex := getEnv("SELLER_PRIVATE_KEY", "")
	if sellerPrivateKeyHex == "" {
		log.Fatal("SELLER PRIVATE_KEY environment variable is not set")
	}

	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	trusteePrivateKeyHex = strings.TrimPrefix(trusteePrivateKeyHex, "0x")
	sellerPrivateKeyHex = strings.TrimPrefix(sellerPrivateKeyHex, "0x")

	client, err = ethclient.Dial(ganacheURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	buyerPrivateKey, err = crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	trusteePrivateKey, err = crypto.HexToECDSA(trusteePrivateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	sellerPrivateKey, err = crypto.HexToECDSA(sellerPrivateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	router := gin.Default()
	router.POST("/createEscrow", handleCreateEscrow)
	router.POST("/approvePayout", handleApprovePayout)
	router.POST("/disapprove", handleDisapprove)
	router.GET("/getSellerContracts", handleGetSellerContracts)
	router.Run(":3000")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func updateEnvFile(key, value, filePath string) error {
	// Read the environment file
	envMap, err := godotenv.Read(filePath)
	if err != nil {
		return fmt.Errorf("error reading .env file: %w", err)
	}

	// Update the value
	envMap[key] = value

	// Create a new file to write the updated values
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating .env file: %w", err)
	}
	defer file.Close()

	// Write the updated values to the file
	writer := bufio.NewWriter(file)
	for k, v := range envMap {
		_, err := writer.WriteString(fmt.Sprintf("%s=%s\n", k, v))
		if err != nil {
			return fmt.Errorf("error writing to .env file: %w", err)
		}
	}
	writer.Flush()

	return nil
}

func handleCreateEscrow(c *gin.Context) {
	var createRequest CreateEscrowRequest
	if err := c.ShouldBindJSON(&createRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Run the truffle migrate command
	cmd := exec.Command("truffle", "migrate", "--network", "development")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Printf("Failed to deploy contract: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deploy contract"})
		return
	}

	// Extract the contract address from the output
	re := regexp.MustCompile(`contract address:\s+([0-9a-fA-Fx]+)`)
	matches := re.FindStringSubmatch(out.String())
	if len(matches) < 2 {
		log.Printf("Failed to find contract address in truffle output")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deploy contract"})
		return
	}

	contractAddress := matches[1]
	log.Printf("Contract deployed at address: %s", contractAddress)

	// Update the .env file with the new contract address
	err = updateEnvFile("CONTRACT_ADDRESS", contractAddress, ".env")
	if err != nil {
		log.Printf("Failed to update .env file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update .env file"})
		return
	}

	// Load the updated contract address
	lastDeployedContractAddress, err := loadUpdatedContractAddress()
	if err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Updated contract address: %s", lastDeployedContractAddress.Hex())

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
		"contractAddress": lastDeployedContractAddress,
	})
}

func loadUpdatedContractAddress() (common.Address, error) {
	// Reload the .env file to fetch the updated contract address
	err := godotenv.Load(".env")
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to reload .env file: %v", err)
	}

	// Fetch the updated contract address
	updatedContractAddress := getEnv("CONTRACT_ADDRESS", "")
	if updatedContractAddress == "" {
		return common.Address{}, fmt.Errorf("failed to fetch updated contract address from .env file")
	}

	return common.HexToAddress(updatedContractAddress), nil
}

func verifyEscrowContract(contractAddress common.Address) error {
	instance, err := NewMain(contractAddress, client)
	if err != nil {
		return fmt.Errorf("failed to instantiate contract: %v", err)
	}

	buyer, err := instance.Buyer(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get buyer: %v", err)
	}
	log.Printf("Verified Buyer address: %s", buyer.Hex())

	seller, err := instance.Seller(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get seller: %v", err)
	}
	log.Printf("Verified Seller address: %s", seller.Hex())

	trustee, err := instance.Trustee(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get trustee: %v", err)
	}
	log.Printf("Verified Trustee address: %s", trustee.Hex())

	amount, err := instance.AmountINR(&bind.CallOpts{})
	if err != nil {
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

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Printf("Failed to get chain ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get chain ID"})
		return
	}

	contractAddress, err := loadUpdatedContractAddress()
	if err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Updated contract address: %s", lastDeployedContractAddress)

	instance, err := NewMain(contractAddress, client)
	if err != nil {
		log.Printf("Failed to instantiate contract: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to instantiate contract"})
		return
	}

	var tx *types.Transaction
	switch approveRequest.Role {
	case "buyer":
		auth, err := bind.NewKeyedTransactorWithChainID(buyerPrivateKey, chainID)
		if err != nil {
			log.Printf("Failed to create transactor: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transactor"})
			return
		}
		tx, err = instance.ApproveByBuyer(auth)
	case "seller":
		auth, err := bind.NewKeyedTransactorWithChainID(sellerPrivateKey, chainID)
		if err != nil {
			log.Printf("Failed to create transactor: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transactor"})
			return
		}
		tx, err = instance.ApproveBySeller(auth)
	case "trustee":
		auth, err := bind.NewKeyedTransactorWithChainID(trusteePrivateKey, chainID)
		if err != nil {
			log.Printf("Failed to create transactor: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transactor"})
			return
		}
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

	if tx == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Contract is completed already."})
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

func handleGetSellerContracts(c *gin.Context) {
	sellerAddress := os.Getenv("SELLER_ADDRESS")
	if sellerAddress == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "SELLER_ADDRESS not set in environment"})
		return
	}

	seller := common.HexToAddress(sellerAddress)

	// Get the latest block number
	latestBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get latest block number"})
		return
	}

	// Get contracts
	contracts, err := getEscrowContracts(client, nil, &seller, nil, 0, big.NewInt(int64(latestBlock)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get seller contracts"})
		return
	}

	// Collect contract info
	var contractInfos []gin.H
	for _, contractAddress := range contracts {
		instance, err := NewMain(contractAddress, client)
		if err != nil {
			continue // Skip this contract if we can't instantiate it
		}

		buyer, _ := instance.Buyer(&bind.CallOpts{})
		seller, _ := instance.Seller(&bind.CallOpts{})
		trustee, _ := instance.Trustee(&bind.CallOpts{})
		amount, _ := instance.AmountINR(&bind.CallOpts{})
		state, _ := instance.CurrentState(&bind.CallOpts{})

		// Convert state to string based on its value
		var stateStr string
		switch state {
		case 0:
			stateStr = "Active"
		case 1:
			stateStr = "Completed"
		default:
			stateStr = fmt.Sprintf("Unknown (%d)", state)
		}

		contractInfos = append(contractInfos, gin.H{
			"address": contractAddress.Hex(),
			"buyer":   buyer.Hex(),
			"seller":  seller.Hex(),
			"trustee": trustee.Hex(),
			"amount":  amount.String(),
			"state":   stateStr,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"sellerAddress": sellerAddress,
		"contracts":     contractInfos,
	})
}

func handleDisapprove(c *gin.Context) {
	var disapproveRequest ApproveRequest
	if err := c.ShouldBindJSON(&disapproveRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Printf("Failed to get chain ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get chain ID"})
		return
	}

	contractAddress, err := loadUpdatedContractAddress()
	if err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Updated contract address: %s", lastDeployedContractAddress.Hex())

	instance, err := NewMain(contractAddress, client)
	if err != nil {
		log.Printf("Failed to instantiate contract: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to instantiate contract"})
		return
	}

	var tx *types.Transaction
	switch disapproveRequest.Role {
	case "buyer":
		auth, err := bind.NewKeyedTransactorWithChainID(buyerPrivateKey, chainID)
		if err != nil {
			log.Printf("Failed to create transactor: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transactor"})
			return
		}
		tx, err = instance.DisapproveByBuyer(auth)
	case "seller":
		auth, err := bind.NewKeyedTransactorWithChainID(sellerPrivateKey, chainID)
		if err != nil {
			log.Printf("Failed to create transactor: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transactor"})
			return
		}
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

	// Check if the contract is still active
	state, err := instance.CurrentState(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get contract state: %v", err)
	}
	if state != 0 { // 0 is Active, 1 is Completed
		log.Println("Contract is no longer active. Stopping event check.")
		return nil
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
		log.Printf("Buyer Bank Details:")
		log.Printf("  Account Holder Name: %s", event.BuyerAccountHolderName)
		log.Printf("  Account Number: %s", event.BuyerAccountNumber)
		log.Printf("  Bank Name: %s", event.BuyerBankName)
		log.Printf("  IFSC Code: %s", event.BuyerIfscCode)
		log.Printf("Seller Bank Details:")
		log.Printf("  Account Holder Name: %s", event.SellerAccountHolderName)
		log.Printf("  Account Number: %s", event.SellerAccountNumber)
		log.Printf("  Bank Name: %s", event.SellerBankName)
		log.Printf("  IFSC Code: %s", event.SellerIfscCode)
		log.Printf("Amount INR: %s", event.AmountINR.String())

		// Check contract state again before processing payout
		state, err = instance.CurrentState(&bind.CallOpts{})
		if err != nil {
			log.Printf("Failed to get contract state before processing payout: %v", err)
			continue
		}
		if state != 0 {
			log.Println("Contract is no longer active. Skipping payout processing.")
			continue
		}

		processPayout(event, contractAddress, client, trusteePrivateKey)
	}

	return nil
}

func processPayout(event *MainPayoutRequested, contractAddress common.Address, client *ethclient.Client, privateKey *ecdsa.PrivateKey) {
	log.Printf("Processing payout:")
	log.Printf("From: %s (Account: %s, Bank: %s, IFSC: %s)",
		event.BuyerAccountHolderName,
		event.BuyerAccountNumber,
		event.BuyerBankName,
		event.BuyerIfscCode)
	log.Printf("To: %s (Account: %s, Bank: %s, IFSC: %s)",
		event.SellerAccountHolderName,
		event.SellerAccountNumber,
		event.SellerBankName,
		event.SellerIfscCode)
	log.Printf("Amount: %s INR", event.AmountINR.String())

	// Retrieve Razorpay credentials from the contract
	instance, err := NewMain(contractAddress, client)
	if err != nil {
		log.Printf("Failed to instantiate contract: %v", err)
		return
	}

	payoutAuth, err := instance.PayoutAuth(&bind.CallOpts{})
	if err != nil {
		log.Printf("Failed to get payout auth from contract: %v", err)
		return
	}
	// Prepare the payout request
	payoutData := map[string]interface{}{
		"account_number": event.BuyerAccountNumber,
		"amount":         event.AmountINR.Int64() * 100, // Convert to paise
		"currency":       "INR",
		"mode":           "NEFT",
		"purpose":        "refund",
		"fund_account": map[string]interface{}{
			"account_type": "bank_account",
			"bank_account": map[string]string{
				"name":           event.SellerAccountHolderName,
				"ifsc":           event.SellerIfscCode,
				"account_number": event.SellerAccountNumber,
			},
			"contact": map[string]interface{}{
				"name":         event.SellerAccountHolderName, // Assuming event has SellerAccountHolderName
				"email":        "gaurav.kumar@example.com",    // Assuming event has SellerEmail
				"contact":      "9876543210",
				"type":         "vendor",
				"reference_id": "Acme Contact ID 12345",
				"notes": map[string]string{
					"notes_key_1": "Tea, Earl Grey, Hot",    // Example note, replace with actual data if available
					"notes_key_2": "Tea, Earl Grey… decaf.", // Example note, replace with actual data if available
				},
			},
		},
		"queue_if_low_balance": true,
		"reference_id":         fmt.Sprintf("payout_%s", event.AmountINR.String()),
		"narration":            "Escrow Payout",
		"notes": map[string]string{
			"notes_key_1": "Beam me up Scotty", // Example note, replace with actual data if available
			"notes_key_2": "Engage",            // Example note, replace with actual data if available
		},
	}

	payoutJSON, err := json.Marshal(payoutData)
	if err != nil {
		log.Printf("Failed to marshal payout data: %v", err)
		return
	}

	// Make the API call to Razorpay
	url := "https://api-web-smart-contract.dev.razorpay.in/v1/payouts"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payoutJSON))
	if err != nil {
		log.Printf("Failed to create HTTP request: %v", err)
		return
	}

	req.SetBasicAuth(payoutAuth.RazorpayApiKey, payoutAuth.RazorpayApiSecret)
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Printf("Failed to send payout request: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Payout request failed. Status: %d, Response: %s", resp.StatusCode, string(body))
		return
	}

	log.Printf("Payout request successful. Response: %s", string(body))

	// Complete the contract after successful payout
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Printf("Failed to get chain ID: %v", err)
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Printf("Failed to create auth transactor: %v", err)
		return
	}

	err = completeContract(contractAddress, auth, client)
	if err != nil {
		log.Printf("Failed to complete contract: %v", err)
		return
	}

	log.Println("Contract completed successfully after payout")
}

func completeContract(contractAddress common.Address, auth *bind.TransactOpts, client *ethclient.Client) error {
	instance, err := NewMain(contractAddress, client)
	if err != nil {
		return fmt.Errorf("failed to instantiate contract: %v", err)
	}

	tx, err := instance.CompleteContract(auth)
	if err != nil {
		return fmt.Errorf("failed to complete contract: %v", err)
	}

	log.Printf("Contract completion transaction sent. Hash: %s", tx.Hash().Hex())

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction to be mined: %v", err)
	}

	if receipt.Status == types.ReceiptStatusFailed {
		return fmt.Errorf("contract completion transaction failed")
	}

	log.Println("Contract successfully marked as completed")
	return nil
}

func getEscrowContracts(client *ethclient.Client, buyerAddress, sellerAddress, trusteeAddress *common.Address, fromBlock uint64, toBlock *big.Int) ([]common.Address, error) {
	escrowCreatedSig := []byte("EscrowCreated(address,address,address,uint256)")
	escrowCreatedTopic := crypto.Keccak256Hash(escrowCreatedSig)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   toBlock,
		Topics:    [][]common.Hash{{escrowCreatedTopic}},
	}

	// Add buyer filter if provided
	if buyerAddress != nil {
		query.Topics = append(query.Topics, []common.Hash{common.BytesToHash(buyerAddress.Bytes())})
	} else {
		query.Topics = append(query.Topics, nil)
	}

	// Add seller filter if provided
	if sellerAddress != nil {
		query.Topics = append(query.Topics, []common.Hash{common.BytesToHash(sellerAddress.Bytes())})
	} else {
		query.Topics = append(query.Topics, nil)
	}

	// Add trustee filter if provided
	if trusteeAddress != nil {
		query.Topics = append(query.Topics, []common.Hash{common.BytesToHash(trusteeAddress.Bytes())})
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to filter logs: %v", err)
	}

	var contractAddresses []common.Address
	for _, vLog := range logs {
		contractAddresses = append(contractAddresses, vLog.Address)
	}

	return contractAddresses, nil
}
