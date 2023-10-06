package models

import (
	"fmt"
	"strings"
)

type MoviesCountry struct {
	CountryId *string `db:"country_id" json:"country_id"`
	MovieId   *int64  `db:"movie_id" json:"movie_id"`
}

func (moviesCountry MoviesCountry) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("CountryId: %v", *moviesCountry.CountryId),
			fmt.Sprintf("MovieId: %v", *moviesCountry.MovieId),
		},
		", ",
	)

	return fmt.Sprintf("MoviesCountry{%s}", content)
}

func (_ MoviesCountry) TableName() string {
	return "app.movies_countries"
}

func (_ MoviesCountry) PrimaryKey() []string {
	return []string{
		"country_id",
		"movie_id",
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

// language=mysql
var moviesCountryInsertSql = `
INSERT INTO app.movies_countries(
  country_id,
  movie_id
)
VALUES (
  :country_id,
  :movie_id
)
RETURNING
  country_id,
  movie_id;
`

// language=mysql
var moviesCountryUpdateSql = `
UPDATE app.movies_countries
SET
  country_id = :country_id,
  movie_id = :movie_id
WHERE TRUE
  AND country_id = :country_id
  AND movie_id = :movie_id
RETURNING
  country_id,
  movie_id;
`

// language=mysql
var moviesCountryFindSql = `
SELECT
  country_id,
  movie_id
FROM app.movies_countries
WHERE TRUE
  AND (:country_id IS NULL or country_id = :country_id)
  AND (:movie_id IS NULL or movie_id = :movie_id)
LIMIT 1;
`

// language=mysql
var moviesCountryFindAllSql = `
SELECT
  country_id,
  movie_id
FROM app.movies_countries
WHERE TRUE
  AND (:country_id IS NULL or country_id = :country_id)
  AND (:movie_id IS NULL or movie_id = :movie_id);
`

// language=mysql
var moviesCountryDeleteSql = `
DELETE FROM app.movies_countries
WHERE TRUE
  AND country_id = :country_id
  AND movie_id = :movie_id;
`
