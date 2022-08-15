package types

type Coin struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Amount int64  `json:"amount"`
}

type UpdateCoinInput struct {
	Name   string `json:"name"`
	Amount int64  `json:"amount"`
}
