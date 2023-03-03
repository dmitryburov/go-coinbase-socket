package usecase

import (
	"context"
	"github.com/dmitryburov/go-coinbase-socket/internal/entity"
	"github.com/dmitryburov/go-coinbase-socket/internal/repository"
	"github.com/dmitryburov/go-coinbase-socket/pkg/logger"
)

type Exchange interface {
	Tick(ctx context.Context, ch <-chan entity.Ticker) error
}

type Services struct {
	Exchange
}

type Packages struct {
	Logger logger.Logger
}

func NewUseCase(repos *repository.Repositories, pkg *Packages) *Services {
	return &Services{
		Exchange: NewExchangeService(repos.Exchange, pkg.Logger),
	}
}
