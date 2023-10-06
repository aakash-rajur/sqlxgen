package models

import (
	"fmt"
	"strings"
)

type Actor struct {
	Name *string `db:"name" json:"name"`
	Id   *int64  `db:"id" json:"id"`
}

func (actor Actor) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Name: %v", *actor.Name),
			fmt.Sprintf("Id: %v", *actor.Id),
		},
		", ",
	)

	return fmt.Sprintf("Actor{%s}", content)
}

func (_ Actor) TableName() string {
	return "app.actors"
}

func (_ Actor) PrimaryKey() []string {
	return []string{
		"id",
	}
}

func (_ Actor) InsertQuery() string {
	return actorInsertSql
}

func (_ Actor) UpdateQuery() string {
	return actorUpdateSql
}

func (_ Actor) FindQuery() string {
	return actorFindSql
}

func (_ Actor) FindAllQuery() string {
	return actorFindAllSql
}

func (_ Actor) DeleteQuery() string {
	return actorDeleteSql
}

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
var actorUpdateSql = `
UPDATE app.actors
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
var actorFindSql = `
SELECT
  name,
  id
FROM app.actors
WHERE TRUE
  AND (:id IS NULL or id = :id)
LIMIT 1;
`

// language=mysql
var actorFindAllSql = `
SELECT
  name,
  id
FROM app.actors
WHERE TRUE
  AND (:id IS NULL or id = :id);
`

// language=mysql
var actorDeleteSql = `
DELETE FROM app.actors
WHERE TRUE
  AND id = :id;
`
