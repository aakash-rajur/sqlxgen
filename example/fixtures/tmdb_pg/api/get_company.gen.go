package api

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/aakash-rajur/example/fixtures/tmdb_pg/store"
)

type GetCompanyArgs struct {
	Id *int64 `db:"id" json:"id"`
}

func (args GetCompanyArgs) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *args.Id),
		},
		", ",
	)

	return fmt.Sprintf("GetCompanyArgs{%s}", content)
}

func (args GetCompanyArgs) Query(db store.Database) ([]GetCompanyResult, error) {
	return store.Query[GetCompanyResult](db, args)
}

func (args GetCompanyArgs) Sql() string {
	return getCompanySql
}

type GetCompanyResult struct {
	Id     *int64                   `db:"id" json:"id"`
	Name   *string                  `db:"name" json:"name"`
	Movies []map[string]interface{} `db:"movies" json:"movies"`
}

func (result GetCompanyResult) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *result.Id),
			fmt.Sprintf("Name: %v", *result.Name),
			fmt.Sprintf("Movies: %v", result.Movies),
		},
		", ",
	)

	return fmt.Sprintf("GetCompanyResult{%s}", content)
}

//go:embed get-company.sql
var getCompanySql string
