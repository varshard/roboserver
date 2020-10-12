package roboserver

import (
	"github.com/varshard/roboserver/algorithm"
	"github.com/varshard/roboserver/services"
)

// RoboServer is an implementation of AssetManagerServer
type RoboServer struct {
	cryptodini services.CryptodiniService
	algorithm algorithm.Algorithm
}

// Deposit handle deposit request and produce BuyOrders
func (r RoboServer) Deposit(uid services.UserID, amount float32) services.BuyOrders {
	portfolio := r.cryptodini.GetPort(uid)
	return r.algorithm.CreateBuyOrders(portfolio, amount)
}

// Deposit handle withdraw request and produce SellOrders
func (r RoboServer) Withdraw(uid services.UserID, amount float32) services.SellOrders {
	portfolio := r.cryptodini.GetPort(uid)
	return r.algorithm.CreateSellOrders(portfolio, amount)
}

// GetPort return portfolio of a user
func (r RoboServer) GetPort(uid services.UserID) services.Portfolio {
	return *r.cryptodini.GetPort(uid)
}

// InitRoboServer create a RoboServer from the supplied CryptodiniServer and Algorithm with Dependency Injection
func InitRoboServer(cryptodini services.CryptodiniService, algo algorithm.Algorithm) RoboServer{
	return RoboServer{
		cryptodini,
		algo,
	}
}
