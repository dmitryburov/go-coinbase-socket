package repository

import (
	"github.com/jmoiron/sqlx"
	"reflect"
	"testing"
)

func TestNewRepositories(t *testing.T) {
	type args struct {
		db *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want *Repositories
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRepositories(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepositories() = %v, want %v", got, tt.want)
			}
		})
	}
}
