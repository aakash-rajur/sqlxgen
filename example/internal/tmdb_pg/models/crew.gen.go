package models

import (
	"fmt"
	"strings"
)

type Crew struct {
	Id         *int32  `db:"id" json:"id"`
	Name       *string `db:"name" json:"name"`
	NameSearch *string `db:"name_search" json:"name_search"`
}

func (c *Crew) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *c.Id),
			fmt.Sprintf("Name: %v", *c.Name),
			fmt.Sprintf("NameSearch: %v", *c.NameSearch),
		},
		", ",
	)

	return fmt.Sprintf("Crew{%s}", content)
}

func (c *Crew) TableName() string {
	return "public.crew"
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

// language=postgresql
var crewInsertSql = `
INSERT INTO public.crew(
  name
)
VALUES (
  :name
)
RETURNING
  id,
  name,
  name_search;
`

// language=postgresql
var crewUpdateSql = `
UPDATE public.crew
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
var crewFindSql = `
SELECT
  id,
  name,
  name_search
FROM public.crew
WHERE TRUE
  AND id = :id;
LIMIT 1;
`

// language=postgresql
var crewFindAllSql = `
SELECT
  id,
  name,
  name_search
FROM public.crew
WHERE TRUE
  AND (CAST(:id AS INT4) IS NULL or id = :id)
  AND (CAST(:name AS TEXT) IS NULL or name = :name)
  AND (CAST(:name_search AS TSVECTOR) IS NULL or name_search = :name_search);
`

// language=postgresql
var crewDeleteSql = `
DELETE FROM public.crew
WHERE TRUE
  AND id = :id;
`
