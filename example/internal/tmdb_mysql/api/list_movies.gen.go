package api

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/aakash-rajur/example/internal/tmdb_mysql/store"
)

type ListMoviesArgs struct {
	GenreId *string `db:"genre_id" json:"genre_id"`
	Limit   *int32  `db:"limit" json:"limit"`
	Offset  *int32  `db:"offset" json:"offset"`
	Search  *string `db:"search" json:"search"`
	Sort    *string `db:"sort" json:"sort"`
}

func (args ListMoviesArgs) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("GenreId: %v", *args.GenreId),
			fmt.Sprintf("Limit: %v", *args.Limit),
			fmt.Sprintf("Offset: %v", *args.Offset),
			fmt.Sprintf("Search: %v", *args.Search),
			fmt.Sprintf("Sort: %v", *args.Sort),
		},
		", ",
	)

	return fmt.Sprintf("ListMoviesArgs{%s}", content)
}

func (args ListMoviesArgs) Query(db store.Database) ([]ListMoviesResult, error) {
	return store.Query[ListMoviesResult](db, args)
}

func (args ListMoviesArgs) Sql() string {
	return listMoviesSql
}

type ListMoviesResult struct {
}

func (result ListMoviesResult) String() string {
	content := strings.Join(
		[]string{},
		", ",
	)

	return fmt.Sprintf("ListMoviesResult{%s}", content)
}

//go:embed list-movies.sql
var listMoviesSql string
