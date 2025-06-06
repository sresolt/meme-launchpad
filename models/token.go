package models

import (
	"time"
)

type Token struct {
	ID                   int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt            time.Time `gorm:"autoCreateTime" json:"created_at"`
	ContractAddress      string    `json:"contract_address"`
	UserAddress          string    `json:"user_address"`
	LiquidityPoolAddress string    `json:"liquidity_pool_address"`
	Name                 string    `json:"name"`
	Symbol               string    `json:"symbol"`
	InitialSupply        int64     `json:"initial_supply"`
	DiscordURL           string    `json:"discord_url"`
	TelegramURL          string    `json:"telegram_url"`
	XURL                 string    `json:"x_url"`
}

type TokenInput struct {
	Name          string `json:"name" binding:"required"`
	Symbol        string `json:"symbol" binding:"required"`
	InitialSupply int64  `json:"initial_supply" binding:"required"`
	DiscordURL    string `json:"discord_url"`
	TelegramURL   string `json:"telegram_url"`
	XURL          string `json:"x_url"`
}
