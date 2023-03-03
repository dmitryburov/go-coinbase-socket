package app

import (
	"context"
	"github.com/dmitryburov/go-coinbase-socket/config"
	"github.com/dmitryburov/go-coinbase-socket/internal/delivery/websocket"
	"github.com/dmitryburov/go-coinbase-socket/internal/repository"
	"github.com/dmitryburov/go-coinbase-socket/internal/usecase"
	"github.com/dmitryburov/go-coinbase-socket/pkg/database/mysql"
	"github.com/dmitryburov/go-coinbase-socket/pkg/exchange/coinbase"
	"github.com/dmitryburov/go-coinbase-socket/pkg/logger/zap"
	"os"
)

// Run started application
func Run(ctx context.Context, cfg *config.Config) {
	// logger
	loggerProvider := zap.NewZapLogger(cfg.Logger.Level, cfg.Logger.DisableCaller, cfg.Logger.DisableStacktrace)
	loggerProvider.InitLogger()

	// database
	dbClient, err := mysql.NewMysqlClient(cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.Base)
	if err != nil {
		loggerProvider.Fatal(err)
	}
	defer dbClient.CloseConnect()

	// exchange
	exchangeClient, err := coinbase.NewCoinbaseClient(cfg.Exchange.Url, cfg.Exchange.Protocol, cfg.Exchange.Origin)
	if err != nil {
		loggerProvider.Fatal(err)
	}
	defer exchangeClient.CloseConnection()

	// repositories & business logic
	repo := repository.NewRepositories(dbClient.DB)
	uc := usecase.NewUseCase(repo, &usecase.Packages{
		Logger: loggerProvider,
	})

	// init client
	client, err := websocket.NewSocketClient(exchangeClient, uc, loggerProvider, cfg.Exchange)
	if err != nil {
		loggerProvider.Fatal(err)
	}

	// run
	go func() {
		loggerProvider.Info("socket starting...")
		if err = client.Run(ctx); err != nil {
			loggerProvider.Fatal(err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()

	loggerProvider.Info("socket stopping...")
}
