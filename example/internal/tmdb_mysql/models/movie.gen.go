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

func (m *Movie) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Budget: %v", *m.Budget),
			fmt.Sprintf("Homepage: %v", *m.Homepage),
			fmt.Sprintf("Keywords: %v", *m.Keywords),
			fmt.Sprintf("OriginalLanguage: %v", *m.OriginalLanguage),
			fmt.Sprintf("OriginalTitle: %v", *m.OriginalTitle),
			fmt.Sprintf("Overview: %v", *m.Overview),
			fmt.Sprintf("Popularity: %v", *m.Popularity),
			fmt.Sprintf("ReleaseDate: %v", *m.ReleaseDate),
			fmt.Sprintf("Revenue: %v", *m.Revenue),
			fmt.Sprintf("Runtime: %v", *m.Runtime),
			fmt.Sprintf("Status: %v", *m.Status),
			fmt.Sprintf("Tagline: %v", *m.Tagline),
			fmt.Sprintf("Title: %v", *m.Title),
			fmt.Sprintf("VoteAverage: %v", *m.VoteAverage),
			fmt.Sprintf("VoteCount: %v", *m.VoteCount),
			fmt.Sprintf("Id: %v", *m.Id),
		},
		", ",
	)

	return fmt.Sprintf("Movie{%s}", content)
}

func (m *Movie) TableName() string {
	return "app.movies"
}

func (m *Movie) PrimaryKey() []string {
	return []string{
		"id",
	}
}

func (m *Movie) InsertQuery() string {
	return movieInsertSql
}

func (m *Movie) UpdateQuery() string {
	return movieUpdateSql
}

func (m *Movie) FindQuery() string {
	return movieFindSql
}

func (m *Movie) FindAllQuery() string {
	return movieFindAllSql
}

func (m *Movie) DeleteQuery() string {
	return movieDeleteSql
}

func (m *Movie) DeleteManyQuery() string {
	return movieDeleteManySql
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
  AND (:budget IS NULL or budget = :budget)
  AND (:homepage IS NULL or homepage = :homepage)
  AND (:keywords IS NULL or keywords = :keywords)
  AND (:original_language IS NULL or original_language = :original_language)
  AND (:original_title IS NULL or original_title = :original_title)
  AND (:overview IS NULL or overview = :overview)
  AND (:popularity IS NULL or popularity = :popularity)
  AND (:release_date IS NULL or release_date = :release_date)
  AND (:revenue IS NULL or revenue = :revenue)
  AND (:runtime IS NULL or runtime = :runtime)
  AND (:status IS NULL or status = :status)
  AND (:tagline IS NULL or tagline = :tagline)
  AND (:title IS NULL or title = :title)
  AND (:vote_average IS NULL or vote_average = :vote_average)
  AND (:vote_count IS NULL or vote_count = :vote_count)
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
  AND (:budget IS NULL or budget = :budget)
  AND (:homepage IS NULL or homepage = :homepage)
  AND (:keywords IS NULL or keywords = :keywords)
  AND (:original_language IS NULL or original_language = :original_language)
  AND (:original_title IS NULL or original_title = :original_title)
  AND (:overview IS NULL or overview = :overview)
  AND (:popularity IS NULL or popularity = :popularity)
  AND (:release_date IS NULL or release_date = :release_date)
  AND (:revenue IS NULL or revenue = :revenue)
  AND (:runtime IS NULL or runtime = :runtime)
  AND (:status IS NULL or status = :status)
  AND (:tagline IS NULL or tagline = :tagline)
  AND (:title IS NULL or title = :title)
  AND (:vote_average IS NULL or vote_average = :vote_average)
  AND (:vote_count IS NULL or vote_count = :vote_count)
  AND (:id IS NULL or id = :id);
`

// language=mysql
var movieDeleteSql = `
DELETE FROM app.movies
WHERE TRUE
  AND budget = :budget
  AND homepage = :homepage
  AND keywords = :keywords
  AND original_language = :original_language
  AND original_title = :original_title
  AND overview = :overview
  AND popularity = :popularity
  AND release_date = :release_date
  AND revenue = :revenue
  AND runtime = :runtime
  AND status = :status
  AND tagline = :tagline
  AND title = :title
  AND vote_average = :vote_average
  AND vote_count = :vote_count
  AND id = :id;
`

// language=mysql
var movieDeleteManySql = `
DELETE FROM app.movies
WHERE TRUE
  AND (:budget IS NULL or budget = :budget)
  AND (:homepage IS NULL or homepage = :homepage)
  AND (:keywords IS NULL or keywords = :keywords)
  AND (:original_language IS NULL or original_language = :original_language)
  AND (:original_title IS NULL or original_title = :original_title)
  AND (:overview IS NULL or overview = :overview)
  AND (:popularity IS NULL or popularity = :popularity)
  AND (:release_date IS NULL or release_date = :release_date)
  AND (:revenue IS NULL or revenue = :revenue)
  AND (:runtime IS NULL or runtime = :runtime)
  AND (:status IS NULL or status = :status)
  AND (:tagline IS NULL or tagline = :tagline)
  AND (:title IS NULL or title = :title)
  AND (:vote_average IS NULL or vote_average = :vote_average)
  AND (:vote_count IS NULL or vote_count = :vote_count)
  AND (:id IS NULL or id = :id);
`
