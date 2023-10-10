package api

import (
	_ "embed"
	"fmt"
	"github.com/aakash-rajur/example/internal/tmdb_mysql/store"
	"strings"
)

type ListCrewArgs struct {
	Limit  *int32  `db:"limit" json:"limit"`
	Offset *int32  `db:"offset" json:"offset"`
	Search *string `db:"search" json:"search"`
	Sort   *string `db:"sort" json:"sort"`
}

func (args ListCrewArgs) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Limit: %v", *args.Limit),
			fmt.Sprintf("Offset: %v", *args.Offset),
			fmt.Sprintf("Search: %v", *args.Search),
			fmt.Sprintf("Sort: %v", *args.Sort),
		},
		", ",
	)

	return fmt.Sprintf("ListCrewArgs{%s}", content)
}

func (args ListCrewArgs) Query(db store.Database) ([]ListCrewResult, error) {
	return store.Query[ListCrewResult](db, args)
}

func (args ListCrewArgs) Sql() string {
	return listCrewSql
}

type ListCrewResult struct {
}

func (result ListCrewResult) String() string {
	content := strings.Join(
		[]string{},
		", ",
	)

	return fmt.Sprintf("ListCrewResult{%s}", content)
}

//go:embed list-crew.sql
var listCrewSql string
