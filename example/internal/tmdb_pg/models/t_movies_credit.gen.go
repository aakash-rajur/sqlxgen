package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

type TMoviesCredit struct {
	Casting json.RawMessage `db:"casting" json:"casting"`
	Crew    json.RawMessage `db:"crew" json:"crew"`
	MovieId *int64          `db:"movie_id" json:"movie_id"`
	Title   *string         `db:"title" json:"title"`
}

func (t *TMoviesCredit) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Casting: %v", t.Casting),
			fmt.Sprintf("Crew: %v", t.Crew),
			fmt.Sprintf("MovieId: %v", *t.MovieId),
			fmt.Sprintf("Title: %v", *t.Title),
		},
		", ",
	)

	return fmt.Sprintf("TMoviesCredit{%s}", content)
}

func (t *TMoviesCredit) TableName() string {
	return "public.t_movies_credits"
}

func (t *TMoviesCredit) PrimaryKey() []string {
	return []string{
		"casting",
		"crew",
		"movie_id",
		"title",
	}
}

func (t *TMoviesCredit) InsertQuery() string {
	return tMoviesCreditInsertSql
}

func (t *TMoviesCredit) UpdateAllQuery() string {
	return tMoviesCreditUpdateAllSql
}

func (t *TMoviesCredit) UpdateByPkQuery() string {
	return tMoviesCreditUpdateByPkSql
}

func (t *TMoviesCredit) CountQuery() string {
	return tMoviesCreditModelCountSql
}

func (t *TMoviesCredit) FindAllQuery() string {
	return tMoviesCreditFindAllSql
}

func (t *TMoviesCredit) FindFirstQuery() string {
	return tMoviesCreditFindFirstSql
}

func (t *TMoviesCredit) FindByPkQuery() string {
	return tMoviesCreditFindByPkSql
}

func (t *TMoviesCredit) DeleteByPkQuery() string {
	return tMoviesCreditDeleteByPkSql
}

func (t *TMoviesCredit) DeleteAllQuery() string {
	return tMoviesCreditDeleteAllSql
}

// language=postgresql
var tMoviesCreditAllFieldsWhere = `
WHERE TRUE
    AND (CAST(:casting AS JSONB) IS NULL or casting = :casting)
    AND (CAST(:crew AS JSONB) IS NULL or crew = :crew)
    AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
    AND (CAST(:title AS TEXT) IS NULL or title = :title)
`

// language=postgresql
var tMoviesCreditPkFieldsWhere = `
WHERE casting = :casting
  AND crew = :crew
  AND movie_id = :movie_id
  AND title = :title
`

// language=postgresql
var tMoviesCreditInsertSql = `
INSERT INTO public.t_movies_credits(
  casting,
  crew,
  movie_id,
  title
)
VALUES (
  :casting,
  :crew,
  :movie_id,
  :title
)
RETURNING
  casting,
  crew,
  movie_id,
  title;
`

// language=postgresql
var tMoviesCreditUpdateByPkSql = `
UPDATE public.t_movies_credits
SET
  casting = :casting,
  crew = :crew,
  movie_id = :movie_id,
  title = :title
` + tMoviesCreditPkFieldsWhere + `
RETURNING
  casting,
  crew,
  movie_id,
  title;
`

// language=postgresql
var tMoviesCreditUpdateAllSql = `
UPDATE public.t_movies_credits
SET
  casting = :casting,
  crew = :crew,
  movie_id = :movie_id,
  title = :title
` + tMoviesCreditAllFieldsWhere + `
RETURNING
  casting,
  crew,
  movie_id,
  title;
`

// language=postgresql
var tMoviesCreditModelCountSql = `
SELECT count(*) as count
FROM public.t_movies_credits
` + tMoviesCreditAllFieldsWhere + ";"

// language=postgresql
var tMoviesCreditFindAllSql = `
SELECT
  casting,
  crew,
  movie_id,
  title
FROM public.t_movies_credits
` + tMoviesCreditAllFieldsWhere + ";"

// language=postgresql
var tMoviesCreditFindFirstSql = strings.TrimRight(tMoviesCreditFindAllSql, ";") + `
LIMIT 1;`

// language=postgresql
var tMoviesCreditFindByPkSql = `
SELECT
  casting,
  crew,
  movie_id,
  title
FROM public.t_movies_credits
` + tMoviesCreditPkFieldsWhere + `
LIMIT 1;`

// language=postgresql
var tMoviesCreditDeleteByPkSql = `
DELETE FROM public.t_movies_credits
WHERE casting = :casting
  AND crew = :crew
  AND movie_id = :movie_id
  AND title = :title;
`

// language=postgresql
var tMoviesCreditDeleteAllSql = `
DELETE FROM public.t_movies_credits
WHERE casting = :casting
  AND crew = :crew
  AND movie_id = :movie_id
  AND title = :title;
`
