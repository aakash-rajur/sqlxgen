package api

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/aakash-rajur/example/fixtures/tmdb_pg/store"
)

type GetActorArgs struct {
	Id *int64 `db:"id" json:"id"`
}

func (args GetActorArgs) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *args.Id),
		},
		", ",
	)

	return fmt.Sprintf("GetActorArgs{%s}", content)
}

func (args GetActorArgs) Query(db store.Database) ([]GetActorResult, error) {
	return store.Query[GetActorResult](db, args)
}

func (args GetActorArgs) Sql() string {
	return getActorSql
}

type GetActorResult struct {
	Id     *int32                   `db:"id" json:"id"`
	Name   *string                  `db:"name" json:"name"`
	Movies []map[string]interface{} `db:"movies" json:"movies"`
}

func (result GetActorResult) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *result.Id),
			fmt.Sprintf("Name: %v", *result.Name),
			fmt.Sprintf("Movies: %v", result.Movies),
		},
		", ",
	)

	return fmt.Sprintf("GetActorResult{%s}", content)
}

//go:embed get-actor.sql
var getActorSql string
