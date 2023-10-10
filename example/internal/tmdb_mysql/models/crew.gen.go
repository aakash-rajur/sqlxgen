package models

import (
	"fmt"
	"strings"
)

type Crew struct {
	Name *string `db:"name" json:"name"`
	Id   *int64  `db:"id" json:"id"`
}

func (crew Crew) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Name: %v", *crew.Name),
			fmt.Sprintf("Id: %v", *crew.Id),
		},
		", ",
	)

	return fmt.Sprintf("Crew{%s}", content)
}

func (_ Crew) TableName() string {
	return "app.crew"
}

func (_ Crew) PrimaryKey() []string {
	return []string{
		"id",
	}
}

func (_ Crew) InsertQuery() string {
	return crewInsertSql
}

func (_ Crew) UpdateQuery() string {
	return crewUpdateSql
}

func (_ Crew) FindQuery() string {
	return crewFindSql
}

func (_ Crew) FindAllQuery() string {
	return crewFindAllSql
}

func (_ Crew) DeleteQuery() string {
	return crewDeleteSql
}

// language=mysql
var crewInsertSql = `
INSERT INTO app.crew(
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
var crewUpdateSql = `
UPDATE app.crew
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
var crewFindSql = `
SELECT
  name,
  id
FROM app.crew
WHERE TRUE
  AND (:id IS NULL or id = :id)
LIMIT 1;
`

// language=mysql
var crewFindAllSql = `
SELECT
  name,
  id
FROM app.crew
WHERE TRUE
  AND (:id IS NULL or id = :id);
`

// language=mysql
var crewDeleteSql = `
DELETE FROM app.crew
WHERE TRUE
  AND id = :id;
`
