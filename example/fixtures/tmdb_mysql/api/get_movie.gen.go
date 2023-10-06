package api

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/aakash-rajur/example/fixtures/tmdb_mysql/store"
)

type GetMovieArgs struct {
	Id *int64 `db:"id" json:"id"`
}

func (args GetMovieArgs) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *args.Id),
		},
		", ",
	)

	return fmt.Sprintf("GetMovieArgs{%s}", content)
}

func (args GetMovieArgs) Query(db store.Database) ([]GetMovieResult, error) {
	return store.Query[GetMovieResult](db, args)
}

func (args GetMovieArgs) Sql() string {
	return getMovieSql
}

type GetMovieResult struct {
}

func (result GetMovieResult) String() string {
	content := strings.Join(
		[]string{},
		", ",
	)

	return fmt.Sprintf("GetMovieResult{%s}", content)
}

//go:embed get-movie.sql
var getMovieSql string
