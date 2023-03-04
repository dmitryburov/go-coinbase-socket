package usecase

import (
	"context"
	"github.com/dmitryburov/go-coinbase-socket/internal/entity"
	"github.com/dmitryburov/go-coinbase-socket/internal/repository"
	"github.com/dmitryburov/go-coinbase-socket/pkg/logger"
	"reflect"
	"testing"
)

func TestNewExchangeService(t *testing.T) {
	type args struct {
		exchange repository.Exchange
		logger   logger.Logger
	}
	tests := []struct {
		name string
		args args
		want *exchangeService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewExchangeService(tt.args.exchange, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExchangeService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_exchangeService_Tick(t *testing.T) {
	type fields struct {
		exchange repository.Exchange
		logger   logger.Logger
	}
	type args struct {
		ctx context.Context
		ch  <-chan entity.Ticker
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &exchangeService{
				exchange: tt.fields.exchange,
				logger:   tt.fields.logger,
			}
			if err := e.Tick(tt.args.ctx, tt.args.ch); (err != nil) != tt.wantErr {
				t.Errorf("Tick() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
