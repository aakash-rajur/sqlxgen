package api

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/aakash-rajur/example/internal/tmdb_mysql/store"
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
}

func (result *GetCompanyResult) String() string {
	content := strings.Join(
		[]string{},
		", ",
	)

	return fmt.Sprintf("GetCompanyResult{%s}", content)
}

//go:embed get-company.sql
var getCompanySql string
