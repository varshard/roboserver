package algorithm

import "github.com/varshard/roboserver/services"

// Algorithm is an abstraction of an algorithm.
type Algorithm interface {
	// CreateBuyOrders create BuyOrders based on the current portfolio and deposit amount
	CreateBuyOrders(portfolio *services.Portfolio, amount float32) services.BuyOrders
	// CreateSellOrders create SellOrders based on the current portfolio and withdraw amount
	CreateSellOrders(portfolio *services.Portfolio, amount float32) services.SellOrders
	// AdjustPort adjust the current portfolio
	// return true if the portfolio require an adjustment, and the desired portfolio
	AdjustPort(portfolio *services.Portfolio) (bool, *services.Portfolio)
}

type MockedAlgorithm struct {}

func (m MockedAlgorithm) CreateBuyOrders(portfolio *services.Portfolio, amount float32) services.BuyOrders {
	panic("implement me")
}

func (m MockedAlgorithm) CreateSellOrders(portfolio *services.Portfolio, amount float32) services.SellOrders {
	panic("implement me")
}

func (m MockedAlgorithm) AdjustPort(portfolio *services.Portfolio) (bool, *services.Portfolio) {
	panic("implement me")
}

