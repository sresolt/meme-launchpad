package services

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/janction/meme-launchpad/contracts"
)

func DeployERC20Token(name, symbol string, initialSupply int64) (string, error) {
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to the network: %v", err)
		return "", err
	}
	defer client.Close()

	// Load private key
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
		return "", err
	}

	// Derive the public address from the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot cast public key to ECDSA")
		return "", err
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get the current account nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
		return "", err
	}

	// Suggest a gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get gas price: %v", err)
		return "", err
	}

	// Create the transaction signer
	chainIDStr := os.Getenv("CHAIN_ID")
	chainID, err := strconv.ParseInt(chainIDStr, 10, 64)
	if err != nil {
		log.Fatalf("Invalid CHAIN_ID: %v", err)
		return "", err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
		return "", err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // no ETH to transfer
	auth.GasLimit = uint64(3000000) // gas limit
	auth.GasPrice = gasPrice

	initialSupplyBigInt := new(big.Int).Mul(big.NewInt(initialSupply), big.NewInt(1e18))
	address, _, _, err := contracts.DeployToken(auth, client, name, symbol, initialSupplyBigInt)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
		return "", err
	}

	return address.String(), nil
}
