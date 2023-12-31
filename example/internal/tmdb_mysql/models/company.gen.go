package models

import (
	"fmt"
	"strings"
)

type Company struct {
	Name *string `db:"name" json:"name"`
	Id   *int64  `db:"id" json:"id"`
}

func (company Company) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Name: %v", *company.Name),
			fmt.Sprintf("Id: %v", *company.Id),
		},
		", ",
	)

	return fmt.Sprintf("Company{%s}", content)
}

func (_ Company) TableName() string {
	return "app.companies"
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

// language=mysql
var companyInsertSql = `
INSERT INTO app.companies(
  name,
  id
)
VALUES (
  :name,
  :id
)
RETURNING
  name,
  id;
`

// language=mysql
var companyUpdateSql = `
UPDATE app.companies
SET
  name = :name,
  id = :id
WHERE TRUE
  AND id = :id
RETURNING
  name,
  id;
`

// language=mysql
var companyFindSql = `
SELECT
  name,
  id
FROM app.companies
WHERE TRUE
  AND (:id IS NULL or id = :id)
LIMIT 1;
`

// language=mysql
var companyFindAllSql = `
SELECT
  name,
  id
FROM app.companies
WHERE TRUE
  AND (:id IS NULL or id = :id);
`

// language=mysql
var companyDeleteSql = `
DELETE FROM app.companies
WHERE TRUE
  AND id = :id;
`
