package models

import (
	"fmt"
	"strings"
)

type MoviesCompany struct {
	MovieId   *int64 `db:"movie_id" json:"movie_id"`
	CompanyId *int64 `db:"company_id" json:"company_id"`
}

func (moviesCompany MoviesCompany) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("MovieId: %v", *moviesCompany.MovieId),
			fmt.Sprintf("CompanyId: %v", *moviesCompany.CompanyId),
		},
		", ",
	)

	return fmt.Sprintf("MoviesCompany{%s}", content)
}

func (_ MoviesCompany) TableName() string {
	return "app.movies_companies"
}

func (_ MoviesCompany) PrimaryKey() []string {
	return []string{
		"movie_id",
		"company_id",
	}
}

func (_ MoviesCompany) InsertQuery() string {
	return moviesCompanyInsertSql
}

func (_ MoviesCompany) UpdateQuery() string {
	return moviesCompanyUpdateSql
}

func (_ MoviesCompany) FindQuery() string {
	return moviesCompanyFindSql
}

func (_ MoviesCompany) FindAllQuery() string {
	return moviesCompanyFindAllSql
}

func (_ MoviesCompany) DeleteQuery() string {
	return moviesCompanyDeleteSql
}

// language=mysql
var moviesCompanyInsertSql = `
INSERT INTO app.movies_companies(
  movie_id,
  company_id
)
VALUES (
  :movie_id,
  :company_id
)
RETURNING
  movie_id,
  company_id;
`

// language=mysql
var moviesCompanyUpdateSql = `
UPDATE app.movies_companies
SET
  movie_id = :movie_id,
  company_id = :company_id
WHERE TRUE
  AND movie_id = :movie_id
  AND company_id = :company_id
RETURNING
  movie_id,
  company_id;
`

// language=mysql
var moviesCompanyFindSql = `
SELECT
  movie_id,
  company_id
FROM app.movies_companies
WHERE TRUE
  AND (:movie_id IS NULL or movie_id = :movie_id)
  AND (:company_id IS NULL or company_id = :company_id)
LIMIT 1;
`

// language=mysql
var moviesCompanyFindAllSql = `
SELECT
  movie_id,
  company_id
FROM app.movies_companies
WHERE TRUE
  AND (:movie_id IS NULL or movie_id = :movie_id)
  AND (:company_id IS NULL or company_id = :company_id);
`

// language=mysql
var moviesCompanyDeleteSql = `
DELETE FROM app.movies_companies
WHERE TRUE
  AND movie_id = :movie_id
  AND company_id = :company_id;
`
