package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName    string       `gorm:"unique;not null"`
	Wallet      Wallet       `gorm:"foreignKey:UserID"`
	Investments []Investment `gorm:"foreignKey:UserID"`
	Sells       []Sell       `gorm:"foreignKey:UserID"`
	Balance     float64      `gorm:"not null"`
	Watchlist   []Watchlist  `gorm:"foreignKey:UserID"`
}

type Wallet struct {
	gorm.Model
	UserID     uint
	Deposit    float64 `gorm:"not null"`
	Investment float64 `gorm:"not null"`
}

type Investment struct {
	gorm.Model
	UserID      uint
	StockSymbol string  `gorm:"not null"`
	Amount      float64 `gorm:"not null"`
	Quantity    int     `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	CompanyName string  `gorm:"not null"`
	LogoUrl     string  `gorm:"not null"`
	PriceChange string  `gorm:"not null"`
	Type        string  `gorm:"not null"`
	Date        string  `gorm:"not null"`
}

type Sell struct {
	gorm.Model
	UserID      uint
	StockSymbol string  `gorm:"not null"`
	Amount      float64 `gorm:"not null"`
	Date        string  `gorm:"not null"`
}

type Watchlist struct {
	gorm.Model
	UserID      uint
	Symbol      string `gorm:"not null"`
	CompanyName string `gorm:"not null"`
	LogoUrl     string `gorm:"not null"`
}
