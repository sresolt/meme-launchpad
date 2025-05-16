package models

import (
	"time"
)

type Token struct {
	ID                   int64     `gorm:"primaryKey;autoIncrement"`
	CreatedAt            time.Time `gorm:"autoCreateTime"`
	ContractAddress      string
	UserAddress          string
	LiquidityPoolAddress string
	Name                 string
	Symbol               string
	InitialSupply        int64
	DiscordURL           string
	TelegramURL          string
	XURL                 string
}

type TokenInput struct {
	Name          string `json:"name" binding:"required"`
	Symbol        string `json:"symbol" binding:"required"`
	InitialSupply int64  `json:"initial_supply" binding:"required"`
	DiscordURL    string `json:"discord_url"`
	TelegramURL   string `json:"telegram_url"`
	XURL          string `json:"x_url"`
}
