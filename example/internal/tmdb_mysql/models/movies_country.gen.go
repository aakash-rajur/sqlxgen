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
	return "app.movies_countries"
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

func (m *MoviesCountry) FindQuery() string {
	return moviesCountryFindSql
}

func (m *MoviesCountry) FindAllQuery() string {
	return moviesCountryFindAllSql
}

func (m *MoviesCountry) DeleteQuery() string {
	return moviesCountryDeleteSql
}

// language=mysql
var moviesCountryInsertSql = `
INSERT INTO app.movies_countries(
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

// language=mysql
var moviesCountryUpdateSql = `
UPDATE app.movies_countries
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

// language=mysql
var moviesCountryFindSql = `
SELECT
  movie_id,
  country_id
FROM app.movies_countries
WHERE TRUE
  AND movie_id = :movie_id
  AND country_id = :country_id;
LIMIT 1;
`

// language=mysql
var moviesCountryFindAllSql = `
SELECT
  movie_id,
  country_id
FROM app.movies_countries
WHERE TRUE
  AND (:movie_id IS NULL or movie_id = :movie_id)
  AND (:country_id IS NULL or country_id = :country_id);
`

// language=mysql
var moviesCountryDeleteSql = `
DELETE FROM app.movies_countries
WHERE TRUE
  AND movie_id = :movie_id
  AND country_id = :country_id;
`
