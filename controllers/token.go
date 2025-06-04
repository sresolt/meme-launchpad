package controllers

import (
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/janction/meme-launchpad/models"
	"github.com/janction/meme-launchpad/services"
)

func CreateToken(c *gin.Context) {
	var input models.TokenInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	auth, client, err := services.PrepareAuth()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare auth"})
		return
	}
	defer client.Close()

	contractAddress, err := services.DeployERC20Token(auth, client, input.Name, input.Symbol, input.InitialSupply)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deploy token"})
		return
	}

	auth.Nonce.Add(auth.Nonce, big.NewInt(1))
	liquidityPoolAddress, err := services.DeployAMM(auth, client, contractAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deploy liquidity pool"})
		return
	}

	auth.Nonce.Add(auth.Nonce, big.NewInt(1))
	err = services.AddLiquidity(auth, client, liquidityPoolAddress, contractAddress, 10000, 0.005)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add liquidity"})
		return
	}

	token := models.Token{
		Name:                 input.Name,
		Symbol:               input.Symbol,
		InitialSupply:        input.InitialSupply,
		ContractAddress:      contractAddress,
		LiquidityPoolAddress: liquidityPoolAddress,
		DiscordURL:           input.DiscordURL,
		TelegramURL:          input.TelegramURL,
		XURL:                 input.XURL,
	}

	createdToken, err := services.CreateToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusCreated, createdToken)
}
