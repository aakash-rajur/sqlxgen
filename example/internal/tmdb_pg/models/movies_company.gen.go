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
	return "public.movies_companies"
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

// language=postgresql
var moviesCompanyInsertSql = `
INSERT INTO public.movies_companies(
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

// language=postgresql
var moviesCompanyUpdateSql = `
UPDATE public.movies_companies
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

// language=postgresql
var moviesCompanyFindSql = `
SELECT
  movie_id,
  company_id
FROM public.movies_companies
WHERE TRUE
  AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
  AND (CAST(:company_id AS INT8) IS NULL or company_id = :company_id)
LIMIT 1;
`

// language=postgresql
var moviesCompanyFindAllSql = `
SELECT
  movie_id,
  company_id
FROM public.movies_companies
WHERE TRUE
  AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
  AND (CAST(:company_id AS INT8) IS NULL or company_id = :company_id);
`

// language=postgresql
var moviesCompanyDeleteSql = `
DELETE FROM public.movies_companies
WHERE TRUE
  AND movie_id = :movie_id
  AND company_id = :company_id;
`
