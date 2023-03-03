package mysql

import (
	"context"
	"github.com/dmitryburov/go-coinbase-socket/internal/entity"
	"github.com/jmoiron/sqlx"
	"time"
)

type exchangeRepo struct {
	db *sqlx.DB
}

func NewExchangeRepository(db *sqlx.DB) *exchangeRepo {
	return &exchangeRepo{db}
}

func (e *exchangeRepo) CreateTick(ctx context.Context, ticker entity.Ticker) error {
	ctxReq, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if _, err := e.db.NamedExecContext(
		ctxReq,
		"INSERT INTO ticks (symbol, timestamp, bid, ask) VALUES (:symbol, :timestamp, :bid, :ask)",
		ticker,
	); err != nil {
		return err
	}

	return nil
}
