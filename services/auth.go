package services

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func PrepareAuth() (*bind.TransactOpts, *ethclient.Client, error) {
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		return nil, nil, err
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	chainID, _ := strconv.ParseInt(os.Getenv("CHAIN_ID"), 10, 64)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	return auth, client, nil
}
