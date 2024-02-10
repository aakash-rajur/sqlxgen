package models

import (
	"fmt"
	"strings"
)

type Actor struct {
	Id         *int32  `db:"id" json:"id"`
	Name       *string `db:"name" json:"name"`
	NameSearch *string `db:"name_search" json:"name_search"`
}

func (a *Actor) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *a.Id),
			fmt.Sprintf("Name: %v", *a.Name),
			fmt.Sprintf("NameSearch: %v", *a.NameSearch),
		},
		", ",
	)

	return fmt.Sprintf("Actor{%s}", content)
}

func (a *Actor) TableName() string {
	return "public.actors"
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

func (a *Actor) FindQuery() string {
	return actorFindSql
}

func (a *Actor) FindAllQuery() string {
	return actorFindAllSql
}

func (a *Actor) DeleteQuery() string {
	return actorDeleteSql
}

// language=postgresql
var actorInsertSql = `
INSERT INTO public.actors(
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
var actorUpdateSql = `
UPDATE public.actors
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
var actorFindSql = `
SELECT
  id,
  name,
  name_search
FROM public.actors
WHERE TRUE
  AND (CAST(:id AS INT4) IS NULL or id = :id)
  AND (CAST(:name AS TEXT) IS NULL or name = :name)
  AND (CAST(:name_search AS TSVECTOR) IS NULL or name_search = :name_search)
LIMIT 1;
`

// language=postgresql
var actorFindAllSql = `
SELECT
  id,
  name,
  name_search
FROM public.actors
WHERE TRUE
  AND (CAST(:id AS INT4) IS NULL or id = :id)
  AND (CAST(:name AS TEXT) IS NULL or name = :name)
  AND (CAST(:name_search AS TSVECTOR) IS NULL or name_search = :name_search);
`

// language=postgresql
var actorDeleteSql = `
DELETE FROM public.actors
WHERE TRUE
  AND id = :id
  AND name = :name
  AND name_search = :name_search;
`
