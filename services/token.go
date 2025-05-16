package services

import (
	"github.com/janction/meme-launchpad/database"
	"github.com/janction/meme-launchpad/models"
)

func CreateToken(input models.Token) (models.Token, error) {
	result := database.DB.Create(&input)
	return input, result.Error
}
