package api

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/aakash-rajur/example/fixtures/tmdb_pg/store"
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
	TotalRecordsCount *int64     `db:"totalRecordsCount" json:"totalRecordsCount"`
	Id                *int32     `db:"id" json:"id"`
	Title             *string    `db:"title" json:"title"`
	ReleaseDate       *time.Time `db:"releaseDate" json:"releaseDate"`
	Status            *string    `db:"status" json:"status"`
	Popularity        *float64   `db:"popularity" json:"popularity"`
}

func (result ListMoviesResult) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("TotalRecordsCount: %v", *result.TotalRecordsCount),
			fmt.Sprintf("Id: %v", *result.Id),
			fmt.Sprintf("Title: %v", *result.Title),
			fmt.Sprintf("ReleaseDate: %v", *result.ReleaseDate),
			fmt.Sprintf("Status: %v", *result.Status),
			fmt.Sprintf("Popularity: %v", *result.Popularity),
		},
		", ",
	)

	return fmt.Sprintf("ListMoviesResult{%s}", content)
}

//go:embed list-movies.sql
var listMoviesSql string
