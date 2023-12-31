package models

import (
	"fmt"
	"strings"
)

type Company struct {
	Id         *int64  `db:"id" json:"id"`
	Name       *string `db:"name" json:"name"`
	NameSearch *string `db:"name_search" json:"name_search"`
}

func (company Company) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *company.Id),
			fmt.Sprintf("Name: %v", *company.Name),
			fmt.Sprintf("NameSearch: %v", *company.NameSearch),
		},
		", ",
	)

	return fmt.Sprintf("Company{%s}", content)
}

func (_ Company) TableName() string {
	return "public.companies"
}

func (_ Company) PrimaryKey() []string {
	return []string{
		"id",
	}
}

func (_ Company) InsertQuery() string {
	return companyInsertSql
}

func (_ Company) UpdateQuery() string {
	return companyUpdateSql
}

func (_ Company) FindQuery() string {
	return companyFindSql
}

func (_ Company) FindAllQuery() string {
	return companyFindAllSql
}

func (_ Company) DeleteQuery() string {
	return companyDeleteSql
}

// language=postgresql
var companyInsertSql = `
INSERT INTO public.companies(
  id,
  name
)
VALUES (
  :id,
  :name
)
RETURNING
  id,
  name,
  name_search;
`

// language=postgresql
var companyUpdateSql = `
UPDATE public.companies
SET
  id = :id,
  name = :name
WHERE TRUE
  AND id = :id
RETURNING
  id,
  name,
  name_search;
`

// language=postgresql
var companyFindSql = `
SELECT
  id,
  name,
  name_search
FROM public.companies
WHERE TRUE
  AND (CAST(:id AS INT8) IS NULL or id = :id)
LIMIT 1;
`

// language=postgresql
var companyFindAllSql = `
SELECT
  id,
  name,
  name_search
FROM public.companies
WHERE TRUE
  AND (CAST(:id AS INT8) IS NULL or id = :id);
`

// language=postgresql
var companyDeleteSql = `
DELETE FROM public.companies
WHERE TRUE
  AND id = :id;
`
