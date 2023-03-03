package mysql

import (
	"github.com/dmitryburov/go-coinbase-socket/internal/entity"
	"github.com/jmoiron/sqlx"
)

type exchangeRepo struct {
	db *sqlx.DB
}

func NewExchangeRepository(db *sqlx.DB) *exchangeRepo {
	return &exchangeRepo{db}
}

func (e *exchangeRepo) CreateTick(ticker entity.Ticker) error {
	return nil
}
