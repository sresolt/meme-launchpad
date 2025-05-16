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
	InitialSupply        string
	DiscordURL           string
	TelegramURL          string
	XURL                 string
}
