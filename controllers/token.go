package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/janction/meme-launchpad/models"
	"github.com/janction/meme-launchpad/services"
)

func CreateToken(c *gin.Context) {
	var input models.Token

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	token, err := services.CreateToken(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusCreated, token)
}
