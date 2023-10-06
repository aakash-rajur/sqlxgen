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

func (crew Crew) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *crew.Id),
			fmt.Sprintf("Name: %v", *crew.Name),
			fmt.Sprintf("NameSearch: %v", *crew.NameSearch),
		},
		", ",
	)

	return fmt.Sprintf("Crew{%s}", content)
}

func (_ Crew) TableName() string {
	return "public.crew"
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
  AND (CAST(:id AS INT4) IS NULL or id = :id)
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
  AND (CAST(:id AS INT4) IS NULL or id = :id);
`

// language=postgresql
var crewDeleteSql = `
DELETE FROM public.crew
WHERE TRUE
  AND id = :id;
`
