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

func (actor Actor) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *actor.Id),
			fmt.Sprintf("Name: %v", *actor.Name),
			fmt.Sprintf("NameSearch: %v", *actor.NameSearch),
		},
		", ",
	)

	return fmt.Sprintf("Actor{%s}", content)
}

func (_ Actor) TableName() string {
	return "public.actors"
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
  AND (CAST(:id AS INT4) IS NULL or id = :id);
`

// language=postgresql
var actorDeleteSql = `
DELETE FROM public.actors
WHERE TRUE
  AND id = :id;
`
