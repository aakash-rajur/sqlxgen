package models

import (
	"fmt"
	"strings"
)

type MoviesCompany struct {
	MovieId   *int64 `db:"movie_id" json:"movie_id"`
	CompanyId *int64 `db:"company_id" json:"company_id"`
}

func (m *MoviesCompany) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("MovieId: %v", *m.MovieId),
			fmt.Sprintf("CompanyId: %v", *m.CompanyId),
		},
		", ",
	)

	return fmt.Sprintf("MoviesCompany{%s}", content)
}

func (m *MoviesCompany) TableName() string {
	return "public.movies_companies"
}

func (m *MoviesCompany) PrimaryKey() []string {
	return []string{
		"movie_id",
		"company_id",
	}
}

func (m *MoviesCompany) InsertQuery() string {
	return moviesCompanyInsertSql
}

func (m *MoviesCompany) UpdateQuery() string {
	return moviesCompanyUpdateSql
}

func (m *MoviesCompany) FindQuery() string {
	return moviesCompanyFindSql
}

func (m *MoviesCompany) FindAllQuery() string {
	return moviesCompanyFindAllSql
}

func (m *MoviesCompany) DeleteQuery() string {
	return moviesCompanyDeleteSql
}

func (m *MoviesCompany) DeleteManyQuery() string {
	return moviesCompanyDeleteManySql
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

// language=postgresql
var moviesCompanyDeleteManySql = `
DELETE FROM public.movies_companies
WHERE TRUE
  AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
  AND (CAST(:company_id AS INT8) IS NULL or company_id = :company_id);
`
