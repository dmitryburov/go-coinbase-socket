package mysql

import (
	"context"
	"github.com/dmitryburov/go-coinbase-socket/internal/entity"
	"github.com/jmoiron/sqlx"
	"reflect"
	"testing"
)

func TestNewExchangeRepository(t *testing.T) {
	type args struct {
		db *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want *exchangeRepo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewExchangeRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExchangeRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_exchangeRepo_CreateTick(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		ctx    context.Context
		ticker entity.Ticker
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
			e := &exchangeRepo{
				db: tt.fields.db,
			}
			if err := e.CreateTick(tt.args.ctx, tt.args.ticker); (err != nil) != tt.wantErr {
				t.Errorf("CreateTick() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
