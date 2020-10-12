package roboserver

import (
	"github.com/varshard/roboserver/algorithm"
	"github.com/varshard/roboserver/services"
)

type Adjustment struct {
	cryptodini services.CryptodiniService
	algo algorithm.Algorithm
}

// AdjustPorts will try to adjust portfolios of investors that it's investing
func (a Adjustment) AdjustPorts() {
	// TODO: Retrieve uID of users that the Robo Serer is managing
	uIDs := []services.UserID{}

	for _, uID := range uIDs {
		portfolio := a.cryptodini.GetPort(uID)

		// Make an adjustment to the received portfolio
		needAdjustment, desiredPortfolio := a.algo.AdjustPort(portfolio)
		if needAdjustment == true {
			a.cryptodini.Adjust(uID, desiredPortfolio)
		}
	}
}
