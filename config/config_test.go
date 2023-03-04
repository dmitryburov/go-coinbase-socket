package config

import (
	"context"
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{name: testing.CoverMode(), args: args{ctx: context.Background()}, want: &Config{
			Exchange: ExchangeConfig{
				Url:      "wss://ws-feed.exchange.coinbase.com",
				Origin:   "https://coinbase.com",
				Protocol: "",
				Symbols:  []string{"ETH-BTC", "BTC-USD", "BTC-EUR"},
				Channels: []string{"ticker"},
			},
			Database: DatabaseConfig{
				Host:     "localhost:3306",
				User:     "test_mysql",
				Password: "a2s_kjlasjd",
				Base:     "test",
			},
			Logger: LoggerConfig{
				DisableCaller:     false,
				DisableStacktrace: true,
				Level:             "debug",
			},
		}, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConfig(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
