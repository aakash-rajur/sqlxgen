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
	Homepage            *string         `db:"homepage" json:"homepage"`
	Genre               json.RawMessage `db:"genre" json:"genre"`
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

func (tMovie TMovie) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *tMovie.Id),
			fmt.Sprintf("Budget: %v", *tMovie.Budget),
			fmt.Sprintf("Homepage: %v", *tMovie.Homepage),
			fmt.Sprintf("Genre: %v", tMovie.Genre),
			fmt.Sprintf("Keywords: %v", tMovie.Keywords),
			fmt.Sprintf("OriginalLanguage: %v", *tMovie.OriginalLanguage),
			fmt.Sprintf("OriginalTitle: %v", *tMovie.OriginalTitle),
			fmt.Sprintf("Overview: %v", *tMovie.Overview),
			fmt.Sprintf("Popularity: %v", *tMovie.Popularity),
			fmt.Sprintf("ProductionCompanies: %v", tMovie.ProductionCompanies),
			fmt.Sprintf("ProductionCountries: %v", tMovie.ProductionCountries),
			fmt.Sprintf("ReleaseDate: %v", *tMovie.ReleaseDate),
			fmt.Sprintf("Revenue: %v", *tMovie.Revenue),
			fmt.Sprintf("Runtime: %v", *tMovie.Runtime),
			fmt.Sprintf("SpokenLanguages: %v", tMovie.SpokenLanguages),
			fmt.Sprintf("Status: %v", *tMovie.Status),
			fmt.Sprintf("Tagline: %v", *tMovie.Tagline),
			fmt.Sprintf("Title: %v", *tMovie.Title),
			fmt.Sprintf("VoteAverage: %v", *tMovie.VoteAverage),
			fmt.Sprintf("VoteCount: %v", *tMovie.VoteCount),
		},
		", ",
	)

	return fmt.Sprintf("TMovie{%s}", content)
}

func (_ TMovie) TableName() string {
	return "public.t_movies"
}

func (_ TMovie) PrimaryKey() []string {
	return []string{
		"id",
	}
}

func (_ TMovie) InsertQuery() string {
	return tMovieInsertSql
}

func (_ TMovie) UpdateQuery() string {
	return tMovieUpdateSql
}

func (_ TMovie) FindQuery() string {
	return tMovieFindSql
}

func (_ TMovie) FindAllQuery() string {
	return tMovieFindAllSql
}

func (_ TMovie) DeleteQuery() string {
	return tMovieDeleteSql
}

// language=postgresql
var tMovieInsertSql = `
INSERT INTO public.t_movies(
  id,
  budget,
  homepage,
  genre,
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
  :homepage,
  :genre,
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
  homepage,
  genre,
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
var tMovieUpdateSql = `
UPDATE public.t_movies
SET
  id = :id,
  budget = :budget,
  homepage = :homepage,
  genre = :genre,
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
WHERE TRUE
  AND id = :id
RETURNING
  id,
  budget,
  homepage,
  genre,
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
var tMovieFindSql = `
SELECT
  id,
  budget,
  homepage,
  genre,
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
WHERE TRUE
  AND (CAST(:id AS INT8) IS NULL or id = :id)
LIMIT 1;
`

// language=postgresql
var tMovieFindAllSql = `
SELECT
  id,
  budget,
  homepage,
  genre,
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
WHERE TRUE
  AND (CAST(:id AS INT8) IS NULL or id = :id);
`

// language=postgresql
var tMovieDeleteSql = `
DELETE FROM public.t_movies
WHERE TRUE
  AND id = :id;
`
