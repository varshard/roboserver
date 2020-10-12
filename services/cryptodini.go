package services

type CryptodiniService interface {
	Adjust(uid UserID, desiredPort *Portfolio)
	GetPort(uid UserID) *Portfolio
}

type MockedCryptodiniService struct {

}

func (c MockedCryptodiniService) Adjust(uid UserID, desiredPort *Portfolio) {
	// TODO: Send adjust to Cryptodini	SDK
}

func (c MockedCryptodiniService) GetPort(uid UserID) *Portfolio {
	// TODO: Send request to get a portfolio to Cryptodini SDK
	return &Portfolio{}
}

