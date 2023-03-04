package usecase

import (
	"context"
	"github.com/dmitryburov/go-coinbase-socket/internal/entity"
	"github.com/dmitryburov/go-coinbase-socket/internal/repository"
	"github.com/dmitryburov/go-coinbase-socket/pkg/logger"
)

// Exchange usecase
type Exchange interface {
	// Tick exchange tick data of stream
	Tick(ctx context.Context, ch <-chan entity.Ticker) error
}

// Services struct of usecase layout
type Services struct {
	Exchange
}

// Packages struct of usecase packages
type Packages struct {
	Logger logger.Logger
}

// NewUseCase create usecase layout
func NewUseCase(repos *repository.Repositories, pkg *Packages) *Services {
	return &Services{
		Exchange: NewExchangeService(repos.Exchange, pkg.Logger),
	}
}
