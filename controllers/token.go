package controllers

import (
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

	contractAddress, err := services.DeployERC20Token(input.Name, input.Symbol, input.InitialSupply)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deploy token"})
		return
	}

	token := models.Token{
		Name:            input.Name,
		Symbol:          input.Symbol,
		InitialSupply:   input.InitialSupply,
		ContractAddress: contractAddress,
		DiscordURL:      input.DiscordURL,
		TelegramURL:     input.TelegramURL,
		XURL:            input.XURL,
	}

	createdToken, err := services.CreateToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusCreated, createdToken)
}
