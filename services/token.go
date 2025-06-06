package services

import (
	"github.com/janction/meme-launchpad/database"
	"github.com/janction/meme-launchpad/models"
)

func CreateToken(input models.Token) (models.Token, error) {
	result := database.DB.Create(&input)
	return input, result.Error
}

func GetTokens() ([]models.Token, error) {
	var tokens []models.Token
	result := database.DB.Where("liquidity_pool_address != ''").Find(&tokens)
	return tokens, result.Error
}
