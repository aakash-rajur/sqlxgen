package models

import (
	"fmt"
	"strings"
	"time"
)

type Movie struct {
	Budget           *int64     `db:"budget" json:"budget"`
	Homepage         *string    `db:"homepage" json:"homepage"`
	Keywords         *string    `db:"keywords" json:"keywords"`
	OriginalLanguage *string    `db:"original_language" json:"original_language"`
	OriginalTitle    *string    `db:"original_title" json:"original_title"`
	Overview         *string    `db:"overview" json:"overview"`
	Popularity       *float32   `db:"popularity" json:"popularity"`
	ReleaseDate      *time.Time `db:"release_date" json:"release_date"`
	Revenue          *int64     `db:"revenue" json:"revenue"`
	Runtime          *int32     `db:"runtime" json:"runtime"`
	Status           *string    `db:"status" json:"status"`
	Tagline          *string    `db:"tagline" json:"tagline"`
	Title            *string    `db:"title" json:"title"`
	VoteAverage      *float32   `db:"vote_average" json:"vote_average"`
	VoteCount        *int32     `db:"vote_count" json:"vote_count"`
	Id               *int64     `db:"id" json:"id"`
}

func (movie Movie) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Budget: %v", *movie.Budget),
			fmt.Sprintf("Homepage: %v", *movie.Homepage),
			fmt.Sprintf("Keywords: %v", *movie.Keywords),
			fmt.Sprintf("OriginalLanguage: %v", *movie.OriginalLanguage),
			fmt.Sprintf("OriginalTitle: %v", *movie.OriginalTitle),
			fmt.Sprintf("Overview: %v", *movie.Overview),
			fmt.Sprintf("Popularity: %v", *movie.Popularity),
			fmt.Sprintf("ReleaseDate: %v", *movie.ReleaseDate),
			fmt.Sprintf("Revenue: %v", *movie.Revenue),
			fmt.Sprintf("Runtime: %v", *movie.Runtime),
			fmt.Sprintf("Status: %v", *movie.Status),
			fmt.Sprintf("Tagline: %v", *movie.Tagline),
			fmt.Sprintf("Title: %v", *movie.Title),
			fmt.Sprintf("VoteAverage: %v", *movie.VoteAverage),
			fmt.Sprintf("VoteCount: %v", *movie.VoteCount),
			fmt.Sprintf("Id: %v", *movie.Id),
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
  budget,
  homepage,
  keywords,
  original_language,
  original_title,
  overview,
  popularity,
  release_date,
  revenue,
  runtime,
  status,
  tagline,
  title,
  vote_average,
  vote_count
)
VALUES (
  :budget,
  :homepage,
  :keywords,
  :original_language,
  :original_title,
  :overview,
  :popularity,
  :release_date,
  :revenue,
  :runtime,
  :status,
  :tagline,
  :title,
  :vote_average,
  :vote_count
)
RETURNING
  budget,
  homepage,
  keywords,
  original_language,
  original_title,
  overview,
  popularity,
  release_date,
  revenue,
  runtime,
  status,
  tagline,
  title,
  vote_average,
  vote_count,
  id;
`

// language=mysql
var movieUpdateSql = `
UPDATE app.movies
SET
  budget = :budget,
  homepage = :homepage,
  keywords = :keywords,
  original_language = :original_language,
  original_title = :original_title,
  overview = :overview,
  popularity = :popularity,
  release_date = :release_date,
  revenue = :revenue,
  runtime = :runtime,
  status = :status,
  tagline = :tagline,
  title = :title,
  vote_average = :vote_average,
  vote_count = :vote_count,
  id = :id
WHERE TRUE
  AND id = :id
RETURNING
  budget,
  homepage,
  keywords,
  original_language,
  original_title,
  overview,
  popularity,
  release_date,
  revenue,
  runtime,
  status,
  tagline,
  title,
  vote_average,
  vote_count,
  id;
`

// language=mysql
var movieFindSql = `
SELECT
  budget,
  homepage,
  keywords,
  original_language,
  original_title,
  overview,
  popularity,
  release_date,
  revenue,
  runtime,
  status,
  tagline,
  title,
  vote_average,
  vote_count,
  id
FROM app.movies
WHERE TRUE
  AND (:id IS NULL or id = :id)
LIMIT 1;
`

// language=mysql
var movieFindAllSql = `
SELECT
  budget,
  homepage,
  keywords,
  original_language,
  original_title,
  overview,
  popularity,
  release_date,
  revenue,
  runtime,
  status,
  tagline,
  title,
  vote_average,
  vote_count,
  id
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
