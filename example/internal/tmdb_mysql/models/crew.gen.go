package models

import (
	"fmt"
	"strings"
)

type Crew struct {
	Name *string `db:"name" json:"name"`
	Id   *int64  `db:"id" json:"id"`
}

func (c *Crew) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Name: %v", *c.Name),
			fmt.Sprintf("Id: %v", *c.Id),
		},
		", ",
	)

	return fmt.Sprintf("Crew{%s}", content)
}

func (c *Crew) TableName() string {
	return "app.crew"
}

func (c *Crew) PrimaryKey() []string {
	return []string{
		"id",
	}
}

func (c *Crew) InsertQuery() string {
	return crewInsertSql
}

func (c *Crew) UpdateQuery() string {
	return crewUpdateSql
}

func (c *Crew) FindQuery() string {
	return crewFindSql
}

func (c *Crew) FindAllQuery() string {
	return crewFindAllSql
}

func (c *Crew) DeleteQuery() string {
	return crewDeleteSql
}

func (c *Crew) DeleteManyQuery() string {
	return crewDeleteManySql
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
  AND (:name IS NULL or name = :name)
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
  AND (:name IS NULL or name = :name)
  AND (:id IS NULL or id = :id);
`

// language=mysql
var crewDeleteSql = `
DELETE FROM app.crew
WHERE TRUE
  AND name = :name
  AND id = :id;
`

// language=mysql
var crewDeleteManySql = `
DELETE FROM app.crew
WHERE TRUE
  AND (:name IS NULL or name = :name)
  AND (:id IS NULL or id = :id);
`
