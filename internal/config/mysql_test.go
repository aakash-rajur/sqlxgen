package config

import (
	"testing"

	"github.com/aakash-rajur/sqlxgen/internal/config/types"
	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestDefaultMysqlConfig(t *testing.T) {
	got := defaultMysqlConfig()

	want := &Config{
		Name:   utils.PointerTo("default"),
		Engine: utils.PointerTo("mysql"),
		Database: &types.Database{
			Host:     utils.PointerTo("localhost"),
			Port:     utils.PointerTo("3306"),
			Db:       utils.PointerTo("mysql"),
			User:     utils.PointerTo("root"),
			Password: utils.PointerTo(""),
			SslMode:  nil,
			Url:      utils.PointerTo("mysql://root:@localhost:3306/mysql"),
		},
		Source: &types.Source{
			Models: &types.Model{
				Schemas: []string{"public"},
				Include: []string{"^.+$"},
				Exclude: []string{},
			},
			Queries: &types.Query{
				Paths:   []string{},
				Include: []string{"^.+$"},
				Exclude: []string{},
			},
		},
		Gen: &types.Gen{
			Store: &types.GenPartial{
				Path: "gen/store",
			},
			Model: &types.GenPartial{
				Path: "gen/mysql",
			},
		},
	}

	assert.Equal(t, want, got)
}
