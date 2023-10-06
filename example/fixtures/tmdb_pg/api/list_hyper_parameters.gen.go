package api

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/aakash-rajur/example/fixtures/tmdb_pg/store"
)

type ListHyperParametersArgs struct {
	Limit  *int32  `db:"limit" json:"limit"`
	Offset *int32  `db:"offset" json:"offset"`
	Search *string `db:"search" json:"search"`
	Type   *string `db:"type" json:"type"`
	Value  *string `db:"value" json:"value"`
}

func (args ListHyperParametersArgs) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Limit: %v", *args.Limit),
			fmt.Sprintf("Offset: %v", *args.Offset),
			fmt.Sprintf("Search: %v", *args.Search),
			fmt.Sprintf("Type: %v", *args.Type),
			fmt.Sprintf("Value: %v", *args.Value),
		},
		", ",
	)

	return fmt.Sprintf("ListHyperParametersArgs{%s}", content)
}

func (args ListHyperParametersArgs) Query(db store.Database) ([]ListHyperParametersResult, error) {
	return store.Query[ListHyperParametersResult](db, args)
}

func (args ListHyperParametersArgs) Sql() string {
	return listHyperParametersSql
}

type ListHyperParametersResult struct {
	TotalRecordsCount *int64  `db:"totalRecordsCount" json:"totalRecordsCount"`
	Value             *string `db:"value" json:"value"`
	FriendlyName      *string `db:"friendlyName" json:"friendlyName"`
}

func (result ListHyperParametersResult) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("TotalRecordsCount: %v", *result.TotalRecordsCount),
			fmt.Sprintf("Value: %v", *result.Value),
			fmt.Sprintf("FriendlyName: %v", *result.FriendlyName),
		},
		", ",
	)

	return fmt.Sprintf("ListHyperParametersResult{%s}", content)
}

//go:embed list-hyper-parameters.sql
var listHyperParametersSql string
