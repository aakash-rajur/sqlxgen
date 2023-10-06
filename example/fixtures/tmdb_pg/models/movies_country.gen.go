package models

import (
	"fmt"
	"strings"
)

type MoviesCountry struct {
	MovieId   *int64  `db:"movie_id" json:"movie_id"`
	CountryId *string `db:"country_id" json:"country_id"`
}

func (moviesCountry MoviesCountry) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("MovieId: %v", *moviesCountry.MovieId),
			fmt.Sprintf("CountryId: %v", *moviesCountry.CountryId),
		},
		", ",
	)

	return fmt.Sprintf("MoviesCountry{%s}", content)
}

func (_ MoviesCountry) TableName() string {
	return "public.movies_countries"
}

func (_ MoviesCountry) PrimaryKey() []string {
	return []string{
		"movie_id",
		"country_id",
	}
}

func (_ MoviesCountry) InsertQuery() string {
	return moviesCountryInsertSql
}

func (_ MoviesCountry) UpdateQuery() string {
	return moviesCountryUpdateSql
}

func (_ MoviesCountry) FindQuery() string {
	return moviesCountryFindSql
}

func (_ MoviesCountry) FindAllQuery() string {
	return moviesCountryFindAllSql
}

func (_ MoviesCountry) DeleteQuery() string {
	return moviesCountryDeleteSql
}

// language=postgresql
var moviesCountryInsertSql = `
INSERT INTO public.movies_countries(
  movie_id,
  country_id
)
VALUES (
  :movie_id,
  :country_id
)
RETURNING
  movie_id,
  country_id;
`

// language=postgresql
var moviesCountryUpdateSql = `
UPDATE public.movies_countries
SET
  movie_id = :movie_id,
  country_id = :country_id
WHERE TRUE
  AND movie_id = :movie_id
  AND country_id = :country_id
RETURNING
  movie_id,
  country_id;
`

// language=postgresql
var moviesCountryFindSql = `
SELECT
  movie_id,
  country_id
FROM public.movies_countries
WHERE TRUE
  AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
  AND (CAST(:country_id AS TEXT) IS NULL or country_id = :country_id)
LIMIT 1;
`

// language=postgresql
var moviesCountryFindAllSql = `
SELECT
  movie_id,
  country_id
FROM public.movies_countries
WHERE TRUE
  AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
  AND (CAST(:country_id AS TEXT) IS NULL or country_id = :country_id);
`

// language=postgresql
var moviesCountryDeleteSql = `
DELETE FROM public.movies_countries
WHERE TRUE
  AND movie_id = :movie_id
  AND country_id = :country_id;
`
