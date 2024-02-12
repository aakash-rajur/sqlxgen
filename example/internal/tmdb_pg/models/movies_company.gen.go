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

func (m *MoviesCompany) UpdateAllQuery() string {
	return moviesCompanyUpdateAllSql
}

func (m *MoviesCompany) UpdateByPkQuery() string {
	return moviesCompanyUpdateByPkSql
}

func (m *MoviesCompany) CountQuery() string {
	return moviesCompanyModelCountSql
}

func (m *MoviesCompany) FindAllQuery() string {
	return moviesCompanyFindAllSql
}

func (m *MoviesCompany) FindFirstQuery() string {
	return moviesCompanyFindFirstSql
}

func (m *MoviesCompany) FindByPkQuery() string {
	return moviesCompanyFindByPkSql
}

func (m *MoviesCompany) DeleteByPkQuery() string {
	return moviesCompanyDeleteByPkSql
}

func (m *MoviesCompany) DeleteAllQuery() string {
	return moviesCompanyDeleteAllSql
}

// language=postgresql
var moviesCompanyAllFieldsWhere = `
WHERE TRUE
    AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
    AND (CAST(:company_id AS INT8) IS NULL or company_id = :company_id)
`

// language=postgresql
var moviesCompanyPkFieldsWhere = `
WHERE movie_id = :movie_id
  AND company_id = :company_id
`

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
var moviesCompanyUpdateByPkSql = `
UPDATE public.movies_companies
SET
  movie_id = :movie_id,
  company_id = :company_id
` + moviesCompanyPkFieldsWhere + `
RETURNING
  movie_id,
  company_id;
`

// language=postgresql
var moviesCompanyUpdateAllSql = `
UPDATE public.movies_companies
SET
  movie_id = :movie_id,
  company_id = :company_id
` + moviesCompanyAllFieldsWhere + `
RETURNING
  movie_id,
  company_id;
`

// language=postgresql
var moviesCompanyModelCountSql = `
SELECT count(*) as count
FROM public.movies_companies
` + moviesCompanyAllFieldsWhere + ";"

// language=postgresql
var moviesCompanyFindAllSql = `
SELECT
  movie_id,
  company_id
FROM public.movies_companies
` + moviesCompanyAllFieldsWhere + ";"

// language=postgresql
var moviesCompanyFindFirstSql = strings.TrimRight(moviesCompanyFindAllSql, ";") + `
LIMIT 1;`

// language=postgresql
var moviesCompanyFindByPkSql = `
SELECT
  movie_id,
  company_id
FROM public.movies_companies
` + moviesCompanyPkFieldsWhere + `
LIMIT 1;`

// language=postgresql
var moviesCompanyDeleteByPkSql = `
DELETE FROM public.movies_companies
WHERE movie_id = :movie_id
  AND company_id = :company_id;
`

// language=postgresql
var moviesCompanyDeleteAllSql = `
DELETE FROM public.movies_companies
WHERE movie_id = :movie_id
  AND company_id = :company_id;
`
