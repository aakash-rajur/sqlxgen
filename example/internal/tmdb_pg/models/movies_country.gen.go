package models

import (
	"fmt"
	"strings"
)

type MoviesCountry struct {
	MovieId   *int64  `db:"movie_id" json:"movie_id"`
	CountryId *string `db:"country_id" json:"country_id"`
}

func (m *MoviesCountry) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("MovieId: %v", *m.MovieId),
			fmt.Sprintf("CountryId: %v", *m.CountryId),
		},
		", ",
	)

	return fmt.Sprintf("MoviesCountry{%s}", content)
}

func (m *MoviesCountry) TableName() string {
	return "public.movies_countries"
}

func (m *MoviesCountry) PrimaryKey() []string {
	return []string{
		"movie_id",
		"country_id",
	}
}

func (m *MoviesCountry) InsertQuery() string {
	return moviesCountryInsertSql
}

func (m *MoviesCountry) UpdateQuery() string {
	return moviesCountryUpdateSql
}

func (m *MoviesCountry) FindFirstQuery() string {
	return moviesCountryFindFirstSql
}

func (m *MoviesCountry) FindByPkQuery() string {
	return moviesCountryFindByPkSql
}

func (m *MoviesCountry) CountQuery() string {
	return moviesCountryCountSql
}

func (m *MoviesCountry) FindAllQuery() string {
	return moviesCountryFindAllSql
}

func (m *MoviesCountry) DeleteByPkQuery() string {
	return moviesCountryDeleteByPkSql
}

func (m *MoviesCountry) DeleteQuery() string {
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
var moviesCountryAllFieldsWhere = `
WHERE TRUE
  AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
  AND (CAST(:country_id AS TEXT) IS NULL or country_id = :country_id)
`

// language=postgresql
var moviesCountryPkFieldsWhere = `
WHERE TRUE
  AND movie_id = :movie_id
  AND country_id = :country_id
`

// language=postgresql
var moviesCountryFindFirstSql = `
SELECT
  movie_id,
  country_id
FROM public.movies_countries
` + moviesCountryAllFieldsWhere + " LIMIT 1;"

// language=postgresql
var moviesCountryFindByPkSql = `
SELECT
  movie_id,
  country_id
FROM public.movies_countries
` + moviesCountryPkFieldsWhere + " LIMIT 1;"

// language=postgresql
var moviesCountryCountSql = `
SELECT count(*) as count
FROM public.movies_countries
` + moviesCountryAllFieldsWhere + ";"

// language=postgresql
var moviesCountryFindAllSql = `
SELECT
  movie_id,
  country_id
FROM public.movies_countries
` + moviesCountryAllFieldsWhere + ";"

// language=postgresql
var moviesCountryDeleteByPkSql = `
DELETE FROM public.movies_countries
WHERE TRUE
  AND movie_id = :movie_id
  AND country_id = :country_id;
`

// language=postgresql
var moviesCountryDeleteSql = `
DELETE FROM public.movies_countries
WHERE TRUE
  AND movie_id = :movie_id
  AND country_id = :country_id;
`
