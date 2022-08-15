package service

import (
	"Skavengerr/coins-crud/internal/types"
)

type CoinsRepository interface {
	GetAllCoins() ([]types.Coin, error)
	CreateCoin(coin types.Coin) error
	GetCoinByID(id int64) (types.Coin, error)
	DeleteCoin(id int64) error
}

type Coins struct {
	repo CoinsRepository
}

func NewCoins(repo CoinsRepository) *Coins {
	return &Coins{
		repo: repo,
	}
}

func (b *Coins) GetAllCoins() ([]types.Coin, error) {
	return b.repo.GetAllCoins()
}

func (b *Coins) CreateCoin(coin types.Coin) error {
	return b.repo.CreateCoin(coin)
}

func (b *Coins) GetCoinByID(id int64) (types.Coin, error) {
	return b.repo.GetCoinByID(id)
}

func (b *Coins) DeleteCoin(id int64) error {
	return b.repo.DeleteCoin(id)
}
