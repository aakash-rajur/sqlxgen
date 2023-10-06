package config

import (
	"github.com/aakash-rajur/sqlxgen/internal/config/types"
	"github.com/aakash-rajur/sqlxgen/internal/utils"
)

func defaultPgConfig() *Config {
	cfg := &Config{
		Name:   utils.PointerTo("default"),
		Engine: utils.PointerTo("postgresql"),
		Database: &types.Database{
			Host:     utils.PointerTo("localhost"),
			Port:     utils.PointerTo("5432"),
			Db:       utils.PointerTo("postgres"),
			User:     utils.PointerTo("postgres"),
			Password: utils.PointerTo("postgres"),
			SslMode:  utils.PointerTo("disable"),
			Url:      utils.PointerTo("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"),
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
				Path: "gen/pg",
			},
		},
	}

	return cfg
}
