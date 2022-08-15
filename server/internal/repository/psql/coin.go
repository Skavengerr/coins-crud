package psql

import (
	"database/sql"

	"Skavengerr/coins-crud/internal/types"
)

type Coins struct {
	db *sql.DB
}

func NewCoins(db *sql.DB) *Coins {
	return &Coins{db}
}

func (b *Coins) CreateCoin(coin types.Coin) error {
	_, err := b.db.Exec("INSERT INTO coins (name, amount) values ($1, $2)", coin.Name, coin.Amount)

	return err
}

func (b *Coins) GetCoinByID(id int64) (types.Coin, error) {
	var coin types.Coin
	err := b.db.QueryRow("SELECT id, name, amount FROM coins WHERE id=$1", id).
		Scan(&coin.ID, &coin.Name, &coin.Amount)
	if err == sql.ErrNoRows {
		return coin, err
	}

	return coin, err
}

func (b *Coins) GetAllCoins() ([]types.Coin, error) {
	rows, err := b.db.Query("SELECT id, name, amount FROM coins")
	if err != nil {
		return nil, err
	}

	coins := make([]types.Coin, 0)
	for rows.Next() {
		var coin types.Coin
		if err := rows.Scan(&coin.ID, &coin.Name, &coin.Amount); err != nil {
			return nil, err
		}

		coins = append(coins, coin)
	}

	return coins, rows.Err()
}

func (b *Coins) DeleteCoin(id int64) error {
	_, err := b.db.Exec("DELETE FROM coins WHERE id=$1", id)

	return err
}
