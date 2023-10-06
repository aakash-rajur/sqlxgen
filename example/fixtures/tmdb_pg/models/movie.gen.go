package models

import (
	"fmt"
	"github.com/lib/pq"
	"strings"
	"time"
)

type Movie struct {
	Id               *int32          `db:"id" json:"id"`
	Title            *string         `db:"title" json:"title"`
	OriginalTitle    *string         `db:"original_title" json:"original_title"`
	OriginalLanguage *string         `db:"original_language" json:"original_language"`
	Overview         *string         `db:"overview" json:"overview"`
	Runtime          *int32          `db:"runtime" json:"runtime"`
	ReleaseDate      *time.Time      `db:"release_date" json:"release_date"`
	Tagline          *string         `db:"tagline" json:"tagline"`
	Status           *string         `db:"status" json:"status"`
	Homepage         *string         `db:"homepage" json:"homepage"`
	Popularity       *float64        `db:"popularity" json:"popularity"`
	VoteAverage      *float64        `db:"vote_average" json:"vote_average"`
	VoteCount        *int32          `db:"vote_count" json:"vote_count"`
	Budget           *int64          `db:"budget" json:"budget"`
	Revenue          *int64          `db:"revenue" json:"revenue"`
	Keywords         *pq.StringArray `db:"keywords" json:"keywords"`
	TitleSearch      *string         `db:"title_search" json:"title_search"`
	KeywordsSearch   *string         `db:"keywords_search" json:"keywords_search"`
}

func (movie Movie) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *movie.Id),
			fmt.Sprintf("Title: %v", *movie.Title),
			fmt.Sprintf("OriginalTitle: %v", *movie.OriginalTitle),
			fmt.Sprintf("OriginalLanguage: %v", *movie.OriginalLanguage),
			fmt.Sprintf("Overview: %v", *movie.Overview),
			fmt.Sprintf("Runtime: %v", *movie.Runtime),
			fmt.Sprintf("ReleaseDate: %v", *movie.ReleaseDate),
			fmt.Sprintf("Tagline: %v", *movie.Tagline),
			fmt.Sprintf("Status: %v", *movie.Status),
			fmt.Sprintf("Homepage: %v", *movie.Homepage),
			fmt.Sprintf("Popularity: %v", *movie.Popularity),
			fmt.Sprintf("VoteAverage: %v", *movie.VoteAverage),
			fmt.Sprintf("VoteCount: %v", *movie.VoteCount),
			fmt.Sprintf("Budget: %v", *movie.Budget),
			fmt.Sprintf("Revenue: %v", *movie.Revenue),
			fmt.Sprintf("Keywords: %v", *movie.Keywords),
			fmt.Sprintf("TitleSearch: %v", *movie.TitleSearch),
			fmt.Sprintf("KeywordsSearch: %v", *movie.KeywordsSearch),
		},
		", ",
	)

	return fmt.Sprintf("Movie{%s}", content)
}

func (_ Movie) TableName() string {
	return "public.movies"
}

func (_ Movie) PrimaryKey() []string {
	return []string{
		"id",
	}
}

func (_ Movie) InsertQuery() string {
	return movieInsertSql
}

func (_ Movie) UpdateQuery() string {
	return movieUpdateSql
}

func (_ Movie) FindQuery() string {
	return movieFindSql
}

func (_ Movie) FindAllQuery() string {
	return movieFindAllSql
}

func (_ Movie) DeleteQuery() string {
	return movieDeleteSql
}

// language=postgresql
var movieInsertSql = `
INSERT INTO public.movies(
  title,
  original_title,
  original_language,
  overview,
  runtime,
  release_date,
  tagline,
  status,
  homepage,
  popularity,
  vote_average,
  vote_count,
  budget,
  revenue,
  keywords
)
VALUES (
  :title,
  :original_title,
  :original_language,
  :overview,
  :runtime,
  :release_date,
  :tagline,
  :status,
  :homepage,
  :popularity,
  :vote_average,
  :vote_count,
  :budget,
  :revenue,
  :keywords
)
RETURNING
  id,
  title,
  original_title,
  original_language,
  overview,
  runtime,
  release_date,
  tagline,
  status,
  homepage,
  popularity,
  vote_average,
  vote_count,
  budget,
  revenue,
  keywords,
  title_search,
  keywords_search;
`

// language=postgresql
var movieUpdateSql = `
UPDATE public.movies
SET
  id = :id,
  title = :title,
  original_title = :original_title,
  original_language = :original_language,
  overview = :overview,
  runtime = :runtime,
  release_date = :release_date,
  tagline = :tagline,
  status = :status,
  homepage = :homepage,
  popularity = :popularity,
  vote_average = :vote_average,
  vote_count = :vote_count,
  budget = :budget,
  revenue = :revenue,
  keywords = :keywords
WHERE TRUE
  AND id = :id
RETURNING
  id,
  title,
  original_title,
  original_language,
  overview,
  runtime,
  release_date,
  tagline,
  status,
  homepage,
  popularity,
  vote_average,
  vote_count,
  budget,
  revenue,
  keywords,
  title_search,
  keywords_search;
`

// language=postgresql
var movieFindSql = `
SELECT
  id,
  title,
  original_title,
  original_language,
  overview,
  runtime,
  release_date,
  tagline,
  status,
  homepage,
  popularity,
  vote_average,
  vote_count,
  budget,
  revenue,
  keywords,
  title_search,
  keywords_search
FROM public.movies
WHERE TRUE
  AND (CAST(:id AS INT4) IS NULL or id = :id)
LIMIT 1;
`

// language=postgresql
var movieFindAllSql = `
SELECT
  id,
  title,
  original_title,
  original_language,
  overview,
  runtime,
  release_date,
  tagline,
  status,
  homepage,
  popularity,
  vote_average,
  vote_count,
  budget,
  revenue,
  keywords,
  title_search,
  keywords_search
FROM public.movies
WHERE TRUE
  AND (CAST(:id AS INT4) IS NULL or id = :id);
`

// language=postgresql
var movieDeleteSql = `
DELETE FROM public.movies
WHERE TRUE
  AND id = :id;
`
