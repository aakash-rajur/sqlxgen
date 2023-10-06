package models

import (
	"fmt"
	"strings"
)

type Company struct {
	Id   *int64  `db:"id" json:"id"`
	Name *string `db:"name" json:"name"`
}

func (company Company) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *company.Id),
			fmt.Sprintf("Name: %v", *company.Name),
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
  id,
  name
)
VALUES (
  :id,
  :name
)
RETURNING
  id,
  name;
`

// language=mysql
var companyUpdateSql = `
UPDATE app.companies
SET
  id = :id,
  name = :name
WHERE TRUE
  AND id = :id
RETURNING
  id,
  name;
`

// language=mysql
var companyFindSql = `
SELECT
  id,
  name
FROM app.companies
WHERE TRUE
  AND (:id IS NULL or id = :id)
LIMIT 1;
`

// language=mysql
var companyFindAllSql = `
SELECT
  id,
  name
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
