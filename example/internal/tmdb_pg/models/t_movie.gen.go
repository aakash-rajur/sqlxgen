package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TMovie struct {
	Id                  *int64          `db:"id" json:"id"`
	Budget              *float64        `db:"budget" json:"budget"`
	Genre               json.RawMessage `db:"genre" json:"genre"`
	Homepage            *string         `db:"homepage" json:"homepage"`
	Keywords            json.RawMessage `db:"keywords" json:"keywords"`
	OriginalLanguage    *string         `db:"original_language" json:"original_language"`
	OriginalTitle       *string         `db:"original_title" json:"original_title"`
	Overview            *string         `db:"overview" json:"overview"`
	Popularity          *float64        `db:"popularity" json:"popularity"`
	ProductionCompanies json.RawMessage `db:"production_companies" json:"production_companies"`
	ProductionCountries json.RawMessage `db:"production_countries" json:"production_countries"`
	ReleaseDate         *time.Time      `db:"release_date" json:"release_date"`
	Revenue             *float64        `db:"revenue" json:"revenue"`
	Runtime             *float64        `db:"runtime" json:"runtime"`
	SpokenLanguages     json.RawMessage `db:"spoken_languages" json:"spoken_languages"`
	Status              *string         `db:"status" json:"status"`
	Tagline             *string         `db:"tagline" json:"tagline"`
	Title               *string         `db:"title" json:"title"`
	VoteAverage         *float64        `db:"vote_average" json:"vote_average"`
	VoteCount           *int64          `db:"vote_count" json:"vote_count"`
}

func (t *TMovie) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *t.Id),
			fmt.Sprintf("Budget: %v", *t.Budget),
			fmt.Sprintf("Genre: %v", t.Genre),
			fmt.Sprintf("Homepage: %v", *t.Homepage),
			fmt.Sprintf("Keywords: %v", t.Keywords),
			fmt.Sprintf("OriginalLanguage: %v", *t.OriginalLanguage),
			fmt.Sprintf("OriginalTitle: %v", *t.OriginalTitle),
			fmt.Sprintf("Overview: %v", *t.Overview),
			fmt.Sprintf("Popularity: %v", *t.Popularity),
			fmt.Sprintf("ProductionCompanies: %v", t.ProductionCompanies),
			fmt.Sprintf("ProductionCountries: %v", t.ProductionCountries),
			fmt.Sprintf("ReleaseDate: %v", *t.ReleaseDate),
			fmt.Sprintf("Revenue: %v", *t.Revenue),
			fmt.Sprintf("Runtime: %v", *t.Runtime),
			fmt.Sprintf("SpokenLanguages: %v", t.SpokenLanguages),
			fmt.Sprintf("Status: %v", *t.Status),
			fmt.Sprintf("Tagline: %v", *t.Tagline),
			fmt.Sprintf("Title: %v", *t.Title),
			fmt.Sprintf("VoteAverage: %v", *t.VoteAverage),
			fmt.Sprintf("VoteCount: %v", *t.VoteCount),
		},
		", ",
	)

	return fmt.Sprintf("TMovie{%s}", content)
}

func (t *TMovie) TableName() string {
	return "public.t_movies"
}

func (t *TMovie) PrimaryKey() []string {
	return []string{
		"id",
	}
}

func (t *TMovie) InsertQuery() string {
	return tMovieInsertSql
}

func (t *TMovie) UpdateAllQuery() string {
	return tMovieUpdateAllSql
}

func (t *TMovie) UpdateByPkQuery() string {
	return tMovieUpdateByPkSql
}

func (t *TMovie) CountQuery() string {
	return tMovieModelCountSql
}

func (t *TMovie) FindAllQuery() string {
	return tMovieFindAllSql
}

func (t *TMovie) FindFirstQuery() string {
	return tMovieFindFirstSql
}

func (t *TMovie) FindByPkQuery() string {
	return tMovieFindByPkSql
}

func (t *TMovie) DeleteByPkQuery() string {
	return tMovieDeleteByPkSql
}

func (t *TMovie) DeleteAllQuery() string {
	return tMovieDeleteAllSql
}

// language=postgresql
var tMovieAllFieldsWhere = `
WHERE TRUE
    AND (CAST(:id AS INT8) IS NULL or id = :id)
    AND (CAST(:budget AS FLOAT8) IS NULL or budget = :budget)
    AND (CAST(:genre AS JSONB) IS NULL or genre = :genre)
    AND (CAST(:homepage AS TEXT) IS NULL or homepage = :homepage)
    AND (CAST(:keywords AS JSONB) IS NULL or keywords = :keywords)
    AND (CAST(:original_language AS TEXT) IS NULL or original_language = :original_language)
    AND (CAST(:original_title AS TEXT) IS NULL or original_title = :original_title)
    AND (CAST(:overview AS TEXT) IS NULL or overview = :overview)
    AND (CAST(:popularity AS FLOAT8) IS NULL or popularity = :popularity)
    AND (CAST(:production_companies AS JSONB) IS NULL or production_companies = :production_companies)
    AND (CAST(:production_countries AS JSONB) IS NULL or production_countries = :production_countries)
    AND (CAST(:release_date AS DATE) IS NULL or release_date = :release_date)
    AND (CAST(:revenue AS FLOAT8) IS NULL or revenue = :revenue)
    AND (CAST(:runtime AS FLOAT8) IS NULL or runtime = :runtime)
    AND (CAST(:spoken_languages AS JSONB) IS NULL or spoken_languages = :spoken_languages)
    AND (CAST(:status AS TEXT) IS NULL or status = :status)
    AND (CAST(:tagline AS TEXT) IS NULL or tagline = :tagline)
    AND (CAST(:title AS TEXT) IS NULL or title = :title)
    AND (CAST(:vote_average AS FLOAT8) IS NULL or vote_average = :vote_average)
    AND (CAST(:vote_count AS INT8) IS NULL or vote_count = :vote_count)
`

// language=postgresql
var tMoviePkFieldsWhere = `
WHERE id = :id
`

// language=postgresql
var tMovieInsertSql = `
INSERT INTO public.t_movies(
  id,
  budget,
  genre,
  homepage,
  keywords,
  original_language,
  original_title,
  overview,
  popularity,
  production_companies,
  production_countries,
  release_date,
  revenue,
  runtime,
  spoken_languages,
  status,
  tagline,
  title,
  vote_average,
  vote_count
)
VALUES (
  :id,
  :budget,
  :genre,
  :homepage,
  :keywords,
  :original_language,
  :original_title,
  :overview,
  :popularity,
  :production_companies,
  :production_countries,
  :release_date,
  :revenue,
  :runtime,
  :spoken_languages,
  :status,
  :tagline,
  :title,
  :vote_average,
  :vote_count
)
RETURNING
  id,
  budget,
  genre,
  homepage,
  keywords,
  original_language,
  original_title,
  overview,
  popularity,
  production_companies,
  production_countries,
  release_date,
  revenue,
  runtime,
  spoken_languages,
  status,
  tagline,
  title,
  vote_average,
  vote_count;
`

// language=postgresql
var tMovieUpdateByPkSql = `
UPDATE public.t_movies
SET
  id = :id,
  budget = :budget,
  genre = :genre,
  homepage = :homepage,
  keywords = :keywords,
  original_language = :original_language,
  original_title = :original_title,
  overview = :overview,
  popularity = :popularity,
  production_companies = :production_companies,
  production_countries = :production_countries,
  release_date = :release_date,
  revenue = :revenue,
  runtime = :runtime,
  spoken_languages = :spoken_languages,
  status = :status,
  tagline = :tagline,
  title = :title,
  vote_average = :vote_average,
  vote_count = :vote_count
` + tMoviePkFieldsWhere + `
RETURNING
  id,
  budget,
  genre,
  homepage,
  keywords,
  original_language,
  original_title,
  overview,
  popularity,
  production_companies,
  production_countries,
  release_date,
  revenue,
  runtime,
  spoken_languages,
  status,
  tagline,
  title,
  vote_average,
  vote_count;
`

// language=postgresql
var tMovieUpdateAllSql = `
UPDATE public.t_movies
SET
  id = :id,
  budget = :budget,
  genre = :genre,
  homepage = :homepage,
  keywords = :keywords,
  original_language = :original_language,
  original_title = :original_title,
  overview = :overview,
  popularity = :popularity,
  production_companies = :production_companies,
  production_countries = :production_countries,
  release_date = :release_date,
  revenue = :revenue,
  runtime = :runtime,
  spoken_languages = :spoken_languages,
  status = :status,
  tagline = :tagline,
  title = :title,
  vote_average = :vote_average,
  vote_count = :vote_count
` + tMovieAllFieldsWhere + `
RETURNING
  id,
  budget,
  genre,
  homepage,
  keywords,
  original_language,
  original_title,
  overview,
  popularity,
  production_companies,
  production_countries,
  release_date,
  revenue,
  runtime,
  spoken_languages,
  status,
  tagline,
  title,
  vote_average,
  vote_count;
`

// language=postgresql
var tMovieModelCountSql = `
SELECT count(*) as count
FROM public.t_movies
` + tMovieAllFieldsWhere + ";"

// language=postgresql
var tMovieFindAllSql = `
SELECT
  id,
  budget,
  genre,
  homepage,
  keywords,
  original_language,
  original_title,
  overview,
  popularity,
  production_companies,
  production_countries,
  release_date,
  revenue,
  runtime,
  spoken_languages,
  status,
  tagline,
  title,
  vote_average,
  vote_count
FROM public.t_movies
` + tMovieAllFieldsWhere + ";"

// language=postgresql
var tMovieFindFirstSql = strings.TrimRight(tMovieFindAllSql, ";") + `
LIMIT 1;`

// language=postgresql
var tMovieFindByPkSql = `
SELECT
  id,
  budget,
  genre,
  homepage,
  keywords,
  original_language,
  original_title,
  overview,
  popularity,
  production_companies,
  production_countries,
  release_date,
  revenue,
  runtime,
  spoken_languages,
  status,
  tagline,
  title,
  vote_average,
  vote_count
FROM public.t_movies
` + tMoviePkFieldsWhere + `
LIMIT 1;`

// language=postgresql
var tMovieDeleteByPkSql = `
DELETE FROM public.t_movies
WHERE id = :id;
`

// language=postgresql
var tMovieDeleteAllSql = `
DELETE FROM public.t_movies
WHERE id = :id
  AND budget = :budget
  AND genre = :genre
  AND homepage = :homepage
  AND keywords = :keywords
  AND original_language = :original_language
  AND original_title = :original_title
  AND overview = :overview
  AND popularity = :popularity
  AND production_companies = :production_companies
  AND production_countries = :production_countries
  AND release_date = :release_date
  AND revenue = :revenue
  AND runtime = :runtime
  AND spoken_languages = :spoken_languages
  AND status = :status
  AND tagline = :tagline
  AND title = :title
  AND vote_average = :vote_average
  AND vote_count = :vote_count;
`
