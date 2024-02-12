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

// language=postgresql
var actorAllFieldsWhere = `
WHERE TRUE
    AND (CAST(:id AS INT4) IS NULL or id = :id)
    AND (CAST(:name AS TEXT) IS NULL or name = :name)
    AND (CAST(:name_search AS TSVECTOR) IS NULL or name_search = :name_search)
`

// language=postgresql
var actorPkFieldsWhere = `
WHERE id = :id
`

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
var actorUpdateByPkSql = `
UPDATE public.actors
SET
  id = :id,
  name = :name
` + actorPkFieldsWhere + `
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
` + actorAllFieldsWhere + `
RETURNING
  id,
  name,
  name_search;
`

// language=postgresql
var actorCountSql = `
SELECT count(*) as count
FROM public.actors
` + actorAllFieldsWhere + ";"

// language=postgresql
var actorFindAllSql = `
SELECT
  id,
  name,
  name_search
FROM public.actors
` + actorAllFieldsWhere + ";"

// language=postgresql
var actorFindFirstSql = strings.TrimRight(actorFindAllSql, ";") + `
LIMIT 1;`

// language=postgresql
var actorFindByPkSql = `
SELECT
  id,
  name,
  name_search
FROM public.actors
` + actorPkFieldsWhere + `
LIMIT 1;`

// language=postgresql
var actorDeleteByPkSql = `
DELETE FROM public.actors
WHERE id = :id;
`

// language=postgresql
var actorDeleteSql = `
DELETE FROM public.actors
WHERE id = :id
  AND name = :name
  AND name_search = :name_search;
`
