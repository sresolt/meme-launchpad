package services

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/janction/meme-launchpad/contracts"
)

func DeployERC20Token(auth *bind.TransactOpts, client *ethclient.Client, name, symbol string, initialSupply int64) (string, error) {
	initialSupplyBigInt := new(big.Int).Mul(big.NewInt(initialSupply), big.NewInt(1e18))

	address, _, _, err := contracts.DeployToken(auth, client, name, symbol, initialSupplyBigInt)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
		return "", err
	}

	return address.String(), nil
}

func DeployAMM(auth *bind.TransactOpts, client *ethclient.Client, tokenAddress string) (string, error) {
	// Convert tokenAddress string to common.Address
	tokenAddr := common.HexToAddress(tokenAddress)

	address, _, _, err := contracts.DeployAmm(auth, client, tokenAddr)
	if err != nil {
		log.Fatalf("Failed to deploy AMM contract: %v", err)
		return "", err
	}

	return address.Hex(), nil
}

func AddLiquidity(auth *bind.TransactOpts, client *ethclient.Client, ammAddress, tokenAddress string, amountToken, amountETH float64) error {
	// Bind the token and AMM contracts
	token, err := contracts.NewToken(common.HexToAddress(tokenAddress), client)
	if err != nil {
		log.Fatalf("Failed to bind token: %v", err)
		return err
	}
	amm, err := contracts.NewAmm(common.HexToAddress(ammAddress), client)
	if err != nil {
		log.Fatalf("Failed to bind AMM: %v", err)
		return err
	}

	// Approve the AMM contract to transfer tokens on behalf of the sender
	txApprove, err := token.Approve(auth, common.HexToAddress(ammAddress), toWei(amountToken))
	if err != nil {
		log.Fatalf("Failed to approve: %v", err)
		return err
	}
	log.Printf("Approve TX: %s", txApprove.Hash().Hex())

	auth2 := *auth
	auth2.Nonce = new(big.Int).Add(auth.Nonce, big.NewInt(1))
	auth2.Value = toWei(amountETH)

	txAdd, err := amm.AddLiquidity(&auth2, toWei(amountToken))
	if err != nil {
		log.Fatalf("Failed to add liquidity: %v", err)
		return err
	}
	log.Printf("AddLiquidity TX: %s", txAdd.Hash().Hex())

	return nil
}

// toWei converts a float64 ETH/token amount to *big.Int with 18 decimals
func toWei(amount float64) *big.Int {
	amountFloat := new(big.Float).Mul(big.NewFloat(amount), big.NewFloat(1e18))
	wei := new(big.Int)
	amountFloat.Int(wei)
	return wei
}
