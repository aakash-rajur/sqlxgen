package models

import (
	"fmt"
	"strings"
	"time"
)

type Movie struct {
	Revenue          *int64     `db:"revenue" json:"revenue"`
	Keywords         *string    `db:"keywords" json:"keywords"`
	Budget           *int64     `db:"budget" json:"budget"`
	VoteCount        *int32     `db:"vote_count" json:"vote_count"`
	VoteAverage      *float32   `db:"vote_average" json:"vote_average"`
	Popularity       *float32   `db:"popularity" json:"popularity"`
	Homepage         *string    `db:"homepage" json:"homepage"`
	Status           *string    `db:"status" json:"status"`
	ReleaseDate      *time.Time `db:"release_date" json:"release_date"`
	Runtime          *int32     `db:"runtime" json:"runtime"`
	Overview         *string    `db:"overview" json:"overview"`
	OriginalLanguage *string    `db:"original_language" json:"original_language"`
	OriginalTitle    *string    `db:"original_title" json:"original_title"`
	Title            *string    `db:"title" json:"title"`
	Id               *int64     `db:"id" json:"id"`
	Tagline          *string    `db:"tagline" json:"tagline"`
}

func (movie Movie) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Revenue: %v", *movie.Revenue),
			fmt.Sprintf("Keywords: %v", *movie.Keywords),
			fmt.Sprintf("Budget: %v", *movie.Budget),
			fmt.Sprintf("VoteCount: %v", *movie.VoteCount),
			fmt.Sprintf("VoteAverage: %v", *movie.VoteAverage),
			fmt.Sprintf("Popularity: %v", *movie.Popularity),
			fmt.Sprintf("Homepage: %v", *movie.Homepage),
			fmt.Sprintf("Status: %v", *movie.Status),
			fmt.Sprintf("ReleaseDate: %v", *movie.ReleaseDate),
			fmt.Sprintf("Runtime: %v", *movie.Runtime),
			fmt.Sprintf("Overview: %v", *movie.Overview),
			fmt.Sprintf("OriginalLanguage: %v", *movie.OriginalLanguage),
			fmt.Sprintf("OriginalTitle: %v", *movie.OriginalTitle),
			fmt.Sprintf("Title: %v", *movie.Title),
			fmt.Sprintf("Id: %v", *movie.Id),
			fmt.Sprintf("Tagline: %v", *movie.Tagline),
		},
		", ",
	)

	return fmt.Sprintf("Movie{%s}", content)
}

func (_ Movie) TableName() string {
	return "app.movies"
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

// language=mysql
var movieInsertSql = `
INSERT INTO app.movies(
  revenue,
  keywords,
  budget,
  vote_count,
  vote_average,
  popularity,
  homepage,
  status,
  release_date,
  runtime,
  overview,
  original_language,
  original_title,
  title,
  tagline
)
VALUES (
  :revenue,
  :keywords,
  :budget,
  :vote_count,
  :vote_average,
  :popularity,
  :homepage,
  :status,
  :release_date,
  :runtime,
  :overview,
  :original_language,
  :original_title,
  :title,
  :tagline
)
RETURNING
  revenue,
  keywords,
  budget,
  vote_count,
  vote_average,
  popularity,
  homepage,
  status,
  release_date,
  runtime,
  overview,
  original_language,
  original_title,
  title,
  id,
  tagline;
`

// language=mysql
var movieUpdateSql = `
UPDATE app.movies
SET
  revenue = :revenue,
  keywords = :keywords,
  budget = :budget,
  vote_count = :vote_count,
  vote_average = :vote_average,
  popularity = :popularity,
  homepage = :homepage,
  status = :status,
  release_date = :release_date,
  runtime = :runtime,
  overview = :overview,
  original_language = :original_language,
  original_title = :original_title,
  title = :title,
  id = :id,
  tagline = :tagline
WHERE TRUE
  AND id = :id
RETURNING
  revenue,
  keywords,
  budget,
  vote_count,
  vote_average,
  popularity,
  homepage,
  status,
  release_date,
  runtime,
  overview,
  original_language,
  original_title,
  title,
  id,
  tagline;
`

// language=mysql
var movieFindSql = `
SELECT
  revenue,
  keywords,
  budget,
  vote_count,
  vote_average,
  popularity,
  homepage,
  status,
  release_date,
  runtime,
  overview,
  original_language,
  original_title,
  title,
  id,
  tagline
FROM app.movies
WHERE TRUE
  AND (:id IS NULL or id = :id)
LIMIT 1;
`

// language=mysql
var movieFindAllSql = `
SELECT
  revenue,
  keywords,
  budget,
  vote_count,
  vote_average,
  popularity,
  homepage,
  status,
  release_date,
  runtime,
  overview,
  original_language,
  original_title,
  title,
  id,
  tagline
FROM app.movies
WHERE TRUE
  AND (:id IS NULL or id = :id);
`

// language=mysql
var movieDeleteSql = `
DELETE FROM app.movies
WHERE TRUE
  AND id = :id;
`
