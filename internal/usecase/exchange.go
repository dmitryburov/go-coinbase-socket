package usecase

import (
	"context"
	"fmt"
	"github.com/dmitryburov/go-coinbase-socket/internal/entity"
	"github.com/dmitryburov/go-coinbase-socket/internal/repository"
	"github.com/dmitryburov/go-coinbase-socket/pkg/logger"
)

type exchangeService struct {
	exchange repository.Exchange
	logger   logger.Logger
}

func NewExchangeService(
	exchange repository.Exchange,
	logger logger.Logger,
) *exchangeService {
	return &exchangeService{exchange, logger}
}

func (e *exchangeService) Tick(ctx context.Context, ch <-chan entity.Ticker) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case v, ok := <-ch:
			if ok {
				if err := e.exchange.CreateTick(v); err != nil {
					return err
				}

				e.logger.Info(fmt.Sprintf("write ticker [%s] > time:%d, bid:%f, ask:%f", v.Symbol, v.Timestamp, v.Bid, v.Ask))
			} else {
				return nil
			}
		}
	}
}
