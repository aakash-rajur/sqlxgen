package models

import (
	"fmt"
	"strings"
)

type Company struct {
	Name *string `db:"name" json:"name"`
	Id   *int64  `db:"id" json:"id"`
}

func (c *Company) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Name: %v", *c.Name),
			fmt.Sprintf("Id: %v", *c.Id),
		},
		", ",
	)

	return fmt.Sprintf("Company{%s}", content)
}

func (c *Company) TableName() string {
	return "app.companies"
}

func (c *Company) PrimaryKey() []string {
	return []string{
		"id",
	}
}

func (c *Company) InsertQuery() string {
	return companyInsertSql
}

func (c *Company) UpdateQuery() string {
	return companyUpdateSql
}

func (c *Company) UpdateByPkQuery() string {
	return companyUpdateByPkSql
}

func (c *Company) CountQuery() string {
	return companyCountSql
}

func (c *Company) FindAllQuery() string {
	return companyFindAllSql
}

func (c *Company) FindFirstQuery() string {
	return companyFindFirstSql
}

func (c *Company) FindByPkQuery() string {
	return companyFindByPkSql
}

func (c *Company) DeleteByPkQuery() string {
	return companyDeleteByPkSql
}

func (c *Company) DeleteQuery() string {
	return companyDeleteSql
}

// language=mysql
var companyAllFieldsWhere = `
WHERE (CAST(:name AS TEXT) IS NULL or name = :name)
  AND (CAST(:id AS BIGINT) IS NULL or id = :id)
`

// language=mysql
var companyPkFieldsWhere = `
WHERE id = :id
`

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
var companyUpdateByPkSql = `
UPDATE app.companies
SET
  name = :name,
  id = :id
` + companyPkFieldsWhere + `
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
` + companyAllFieldsWhere + `
RETURNING
  name,
  id;
`

// language=mysql
var companyCountSql = `
SELECT count(*) as count
FROM app.companies
` + companyAllFieldsWhere + ";"

// language=mysql
var companyFindAllSql = `
SELECT
  name,
  id
FROM app.companies
` + companyAllFieldsWhere + ";"

// language=mysql
var companyFindFirstSql = strings.TrimRight(companyFindAllSql, ";") + `
LIMIT 1;`

// language=mysql
var companyFindByPkSql = `
SELECT
  name,
  id
FROM app.companies
` + companyPkFieldsWhere + `
LIMIT 1;`

// language=mysql
var companyDeleteByPkSql = `
DELETE FROM app.companies
WHERE id = :id;
`

// language=mysql
var companyDeleteSql = `
DELETE FROM app.companies
WHERE name = :name
  AND id = :id;
`
