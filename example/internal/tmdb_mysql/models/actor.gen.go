package models

import (
	"fmt"
	"strings"
)

type Actor struct {
	Name *string `db:"name" json:"name"`
	Id   *int64  `db:"id" json:"id"`
}

func (a *Actor) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Name: %v", *a.Name),
			fmt.Sprintf("Id: %v", *a.Id),
		},
		", ",
	)

	return fmt.Sprintf("Actor{%s}", content)
}

func (a *Actor) TableName() string {
	return "app.actors"
}

func (a *Actor) PrimaryKey() []string {
	return []string{
		"id",
	}
}

func (a *Actor) InsertQuery() string {
	return actorInsertSql
}

func (a *Actor) UpdateQuery() string {
	return actorUpdateSql
}

func (a *Actor) UpdateByPkQuery() string {
	return actorUpdateByPkSql
}

func (a *Actor) CountQuery() string {
	return actorCountSql
}

func (a *Actor) FindAllQuery() string {
	return actorFindAllSql
}

func (a *Actor) FindFirstQuery() string {
	return actorFindFirstSql
}

func (a *Actor) FindByPkQuery() string {
	return actorFindByPkSql
}

func (a *Actor) DeleteByPkQuery() string {
	return actorDeleteByPkSql
}

func (a *Actor) DeleteQuery() string {
	return actorDeleteSql
}

// language=mysql
var actorAllFieldsWhere = `
WHERE (CAST(:name AS TEXT) IS NULL or name = :name)
  AND (CAST(:id AS BIGINT) IS NULL or id = :id)
`

// language=mysql
var actorPkFieldsWhere = `
WHERE id = :id
`

// language=mysql
var actorInsertSql = `
INSERT INTO app.actors(
  name
)
VALUES (
  :name
)
RETURNING
  name,
  id;
`

// language=mysql
var actorUpdateByPkSql = `
UPDATE app.actors
SET
  name = :name,
  id = :id
` + actorPkFieldsWhere + `
RETURNING
  name,
  id;
`

// language=mysql
var actorUpdateSql = `
UPDATE app.actors
SET
  name = :name,
  id = :id
` + actorAllFieldsWhere + `
RETURNING
  name,
  id;
`

// language=mysql
var actorCountSql = `
SELECT count(*) as count
FROM app.actors
` + actorAllFieldsWhere + ";"

// language=mysql
var actorFindAllSql = `
SELECT
  name,
  id
FROM app.actors
` + actorAllFieldsWhere + ";"

// language=mysql
var actorFindFirstSql = strings.TrimRight(actorFindAllSql, ";") + `
LIMIT 1;`

// language=mysql
var actorFindByPkSql = `
SELECT
  name,
  id
FROM app.actors
` + actorPkFieldsWhere + `
LIMIT 1;`

// language=mysql
var actorDeleteByPkSql = `
DELETE FROM app.actors
WHERE id = :id;
`

// language=mysql
var actorDeleteSql = `
DELETE FROM app.actors
WHERE name = :name
  AND id = :id;
`
