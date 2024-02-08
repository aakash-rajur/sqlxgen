package api

import (
	_ "embed"
	"fmt"
	"github.com/aakash-rajur/example/internal/tmdb_pg/store"
	"strings"
)

type GetCompanyArgs struct {
	Id *int64 `db:"id" json:"id"`
}

func (args *GetCompanyArgs) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *args.Id),
		},
		", ",
	)

	return fmt.Sprintf("GetCompanyArgs{%s}", content)
}

func (args *GetCompanyArgs) Query(db store.Database) ([]*GetCompanyResult, error) {
	return store.Query[*GetCompanyResult](db, args)
}

func (args *GetCompanyArgs) Sql() string {
	return getCompanySql
}

type GetCompanyResult struct {
	Id     *int64           `db:"id" json:"id"`
	Movies *store.JsonArray `db:"movies" json:"movies"`
	Name   *string          `db:"name" json:"name"`
}

func (result *GetCompanyResult) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *result.Id),
			fmt.Sprintf("Movies: %v", result.Movies),
			fmt.Sprintf("Name: %v", *result.Name),
		},
		", ",
	)

	return fmt.Sprintf("GetCompanyResult{%s}", content)
}

//go:embed get-company.sql
var getCompanySql string
