package models

import (
	"fmt"
	"strings"
)

type Crew struct {
	Id   *int64  `db:"id" json:"id"`
	Name *string `db:"name" json:"name"`
}

func (crew Crew) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *crew.Id),
			fmt.Sprintf("Name: %v", *crew.Name),
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
  id,
  name;
`

// language=mysql
var crewUpdateSql = `
UPDATE app.crew
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
var crewFindSql = `
SELECT
  id,
  name
FROM app.crew
WHERE TRUE
  AND (:id IS NULL or id = :id)
LIMIT 1;
`

// language=mysql
var crewFindAllSql = `
SELECT
  id,
  name
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
