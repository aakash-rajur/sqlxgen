package api

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/aakash-rajur/example/internal/tmdb_pg/store"
	"github.com/lib/pq"
)

type GetMovieArgs struct {
	Id *int64 `db:"id" json:"id"`
}

func (args *GetMovieArgs) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *args.Id),
		},
		", ",
	)

	return fmt.Sprintf("GetMovieArgs{%s}", content)
}

func (args *GetMovieArgs) Query(db store.Database) ([]*GetMovieResult, error) {
	return store.Query[*GetMovieResult](db, args)
}

func (args *GetMovieArgs) Sql() string {
	return getMovieSql
}

type GetMovieResult struct {
	Actors           *store.JsonArray `db:"actors" json:"actors"`
	Budget           *int64           `db:"budget" json:"budget"`
	Companies        *store.JsonArray `db:"companies" json:"companies"`
	Countries        *store.JsonArray `db:"countries" json:"countries"`
	Crews            *store.JsonArray `db:"crews" json:"crews"`
	Genres           *store.JsonArray `db:"genres" json:"genres"`
	Homepage         *string          `db:"homepage" json:"homepage"`
	Id               *int32           `db:"id" json:"id"`
	Keywords         *pq.StringArray  `db:"keywords" json:"keywords"`
	Languages        *store.JsonArray `db:"languages" json:"languages"`
	OriginalLanguage *string          `db:"originalLanguage" json:"originalLanguage"`
	OriginalTitle    *string          `db:"originalTitle" json:"originalTitle"`
	Overview         *string          `db:"overview" json:"overview"`
	Popularity       *float64         `db:"popularity" json:"popularity"`
	ReleaseDate      *time.Time       `db:"releaseDate" json:"releaseDate"`
	Revenue          *int64           `db:"revenue" json:"revenue"`
	Runtime          *int32           `db:"runtime" json:"runtime"`
	Status           *string          `db:"status" json:"status"`
	Tagline          *string          `db:"tagline" json:"tagline"`
	Title            *string          `db:"title" json:"title"`
	VoteAverage      *float64         `db:"voteAverage" json:"voteAverage"`
	VoteCount        *int32           `db:"voteCount" json:"voteCount"`
}

func (result *GetMovieResult) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Actors: %v", result.Actors),
			fmt.Sprintf("Budget: %v", *result.Budget),
			fmt.Sprintf("Companies: %v", result.Companies),
			fmt.Sprintf("Countries: %v", result.Countries),
			fmt.Sprintf("Crews: %v", result.Crews),
			fmt.Sprintf("Genres: %v", result.Genres),
			fmt.Sprintf("Homepage: %v", *result.Homepage),
			fmt.Sprintf("Id: %v", *result.Id),
			fmt.Sprintf("Keywords: %v", *result.Keywords),
			fmt.Sprintf("Languages: %v", result.Languages),
			fmt.Sprintf("OriginalLanguage: %v", *result.OriginalLanguage),
			fmt.Sprintf("OriginalTitle: %v", *result.OriginalTitle),
			fmt.Sprintf("Overview: %v", *result.Overview),
			fmt.Sprintf("Popularity: %v", *result.Popularity),
			fmt.Sprintf("ReleaseDate: %v", *result.ReleaseDate),
			fmt.Sprintf("Revenue: %v", *result.Revenue),
			fmt.Sprintf("Runtime: %v", *result.Runtime),
			fmt.Sprintf("Status: %v", *result.Status),
			fmt.Sprintf("Tagline: %v", *result.Tagline),
			fmt.Sprintf("Title: %v", *result.Title),
			fmt.Sprintf("VoteAverage: %v", *result.VoteAverage),
			fmt.Sprintf("VoteCount: %v", *result.VoteCount),
		},
		", ",
	)

	return fmt.Sprintf("GetMovieResult{%s}", content)
}

//go:embed get-movie.sql
var getMovieSql string
