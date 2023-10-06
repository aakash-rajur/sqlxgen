package api

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/aakash-rajur/example/fixtures/tmdb_pg/store"
)

type ListCompaniesArgs struct {
	Limit  *int32  `db:"limit" json:"limit"`
	Offset *int32  `db:"offset" json:"offset"`
	Search *string `db:"search" json:"search"`
	Sort   *string `db:"sort" json:"sort"`
}

func (args ListCompaniesArgs) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Limit: %v", *args.Limit),
			fmt.Sprintf("Offset: %v", *args.Offset),
			fmt.Sprintf("Search: %v", *args.Search),
			fmt.Sprintf("Sort: %v", *args.Sort),
		},
		", ",
	)

	return fmt.Sprintf("ListCompaniesArgs{%s}", content)
}

func (args ListCompaniesArgs) Query(db store.Database) ([]ListCompaniesResult, error) {
	return store.Query[ListCompaniesResult](db, args)
}

func (args ListCompaniesArgs) Sql() string {
	return listCompaniesSql
}

type ListCompaniesResult struct {
	TotalRecordsCount *int64  `db:"totalRecordsCount" json:"totalRecordsCount"`
	Id                *int64  `db:"id" json:"id"`
	Name              *string `db:"name" json:"name"`
}

func (result ListCompaniesResult) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("TotalRecordsCount: %v", *result.TotalRecordsCount),
			fmt.Sprintf("Id: %v", *result.Id),
			fmt.Sprintf("Name: %v", *result.Name),
		},
		", ",
	)

	return fmt.Sprintf("ListCompaniesResult{%s}", content)
}

//go:embed list-companies.sql
var listCompaniesSql string
