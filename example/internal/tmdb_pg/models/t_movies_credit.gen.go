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

func (tMoviesCredit TMoviesCredit) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Casting: %v", tMoviesCredit.Casting),
			fmt.Sprintf("Crew: %v", tMoviesCredit.Crew),
			fmt.Sprintf("MovieId: %v", *tMoviesCredit.MovieId),
			fmt.Sprintf("Title: %v", *tMoviesCredit.Title),
		},
		", ",
	)

	return fmt.Sprintf("TMoviesCredit{%s}", content)
}

func (_ TMoviesCredit) TableName() string {
	return "public.t_movies_credits"
}

func (_ TMoviesCredit) PrimaryKey() []string {
	return []string{
		"casting",
		"crew",
		"movie_id",
		"title",
	}
}

func (_ TMoviesCredit) InsertQuery() string {
	return tMoviesCreditInsertSql
}

func (_ TMoviesCredit) UpdateQuery() string {
	return tMoviesCreditUpdateSql
}

func (_ TMoviesCredit) FindQuery() string {
	return tMoviesCreditFindSql
}

func (_ TMoviesCredit) FindAllQuery() string {
	return tMoviesCreditFindAllSql
}

func (_ TMoviesCredit) DeleteQuery() string {
	return tMoviesCreditDeleteSql
}

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
var tMoviesCreditUpdateSql = `
UPDATE public.t_movies_credits
SET
  casting = :casting,
  crew = :crew,
  movie_id = :movie_id,
  title = :title
WHERE TRUE
  AND casting = :casting
  AND crew = :crew
  AND movie_id = :movie_id
  AND title = :title
RETURNING
  casting,
  crew,
  movie_id,
  title;
`

// language=postgresql
var tMoviesCreditFindSql = `
SELECT
  casting,
  crew,
  movie_id,
  title
FROM public.t_movies_credits
WHERE TRUE
  AND (CAST(:casting AS JSONB) IS NULL or casting = :casting)
  AND (CAST(:crew AS JSONB) IS NULL or crew = :crew)
  AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
  AND (CAST(:title AS TEXT) IS NULL or title = :title)
LIMIT 1;
`

// language=postgresql
var tMoviesCreditFindAllSql = `
SELECT
  casting,
  crew,
  movie_id,
  title
FROM public.t_movies_credits
WHERE TRUE
  AND (CAST(:casting AS JSONB) IS NULL or casting = :casting)
  AND (CAST(:crew AS JSONB) IS NULL or crew = :crew)
  AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
  AND (CAST(:title AS TEXT) IS NULL or title = :title);
`

// language=postgresql
var tMoviesCreditDeleteSql = `
DELETE FROM public.t_movies_credits
WHERE TRUE
  AND casting = :casting
  AND crew = :crew
  AND movie_id = :movie_id
  AND title = :title;
`
