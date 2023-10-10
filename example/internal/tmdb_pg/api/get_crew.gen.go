package api

import (
	_ "embed"
	"fmt"
	"github.com/aakash-rajur/example/internal/tmdb_pg/store"
	"strings"
)

type GetCrewArgs struct {
	Id *int64 `db:"id" json:"id"`
}

func (args GetCrewArgs) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *args.Id),
		},
		", ",
	)

	return fmt.Sprintf("GetCrewArgs{%s}", content)
}

func (args GetCrewArgs) Query(db store.Database) ([]GetCrewResult, error) {
	return store.Query[GetCrewResult](db, args)
}

func (args GetCrewArgs) Sql() string {
	return getCrewSql
}

type GetCrewResult struct {
	Id     *int32           `db:"id" json:"id"`
	Movies *store.JsonArray `db:"movies" json:"movies"`
	Name   *string          `db:"name" json:"name"`
}

func (result GetCrewResult) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *result.Id),
			fmt.Sprintf("Movies: %v", result.Movies),
			fmt.Sprintf("Name: %v", *result.Name),
		},
		", ",
	)

	return fmt.Sprintf("GetCrewResult{%s}", content)
}

//go:embed get-crew.sql
var getCrewSql string
