package services

import "time"

type Order struct {
	Currency string `json:"currency"`
	Amount float32 `json:"amount"`
	Price float32 `json:"price"`
}

type BuyOrders struct {
	Orders []Order `json:"orders"`
	Time time.Time `json:"time"`
}

type SellOrders struct {
	Orders []Order `json:"orders"`
	Time time.Time `json:"time"`
}

type Asset struct {
	Currency string `json:"currency"`
	Amount float32 `json:"amount"`
	Price float32 `json:"price"`
	MarketPrice float32 `json:"market_price"`
}

type Portfolio struct {
	Assets []Asset
}

type UserID string

type AssetManagerService interface {
	Deposit(uid UserID, amount float32) BuyOrders
	Withdraw(uid UserID, amount float32) SellOrders
	GetPort(uid UserID) Portfolio
}
