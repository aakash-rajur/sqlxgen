package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

type TMoviesCredit struct {
	MovieId *int64          `db:"movie_id" json:"movie_id"`
	Title   *string         `db:"title" json:"title"`
	Casting json.RawMessage `db:"casting" json:"casting"`
	Crew    json.RawMessage `db:"crew" json:"crew"`
}

func (tMoviesCredit TMoviesCredit) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("MovieId: %v", *tMoviesCredit.MovieId),
			fmt.Sprintf("Title: %v", *tMoviesCredit.Title),
			fmt.Sprintf("Casting: %v", tMoviesCredit.Casting),
			fmt.Sprintf("Crew: %v", tMoviesCredit.Crew),
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
		"movie_id",
		"title",
		"casting",
		"crew",
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
  movie_id,
  title,
  casting,
  crew
)
VALUES (
  :movie_id,
  :title,
  :casting,
  :crew
)
RETURNING
  movie_id,
  title,
  casting,
  crew;
`

// language=postgresql
var tMoviesCreditUpdateSql = `
UPDATE public.t_movies_credits
SET
  movie_id = :movie_id,
  title = :title,
  casting = :casting,
  crew = :crew
WHERE TRUE
  AND movie_id = :movie_id
  AND title = :title
  AND casting = :casting
  AND crew = :crew
RETURNING
  movie_id,
  title,
  casting,
  crew;
`

// language=postgresql
var tMoviesCreditFindSql = `
SELECT
  movie_id,
  title,
  casting,
  crew
FROM public.t_movies_credits
WHERE TRUE
  AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
  AND (CAST(:title AS TEXT) IS NULL or title = :title)
  AND (CAST(:casting AS JSONB) IS NULL or casting = :casting)
  AND (CAST(:crew AS JSONB) IS NULL or crew = :crew)
LIMIT 1;
`

// language=postgresql
var tMoviesCreditFindAllSql = `
SELECT
  movie_id,
  title,
  casting,
  crew
FROM public.t_movies_credits
WHERE TRUE
  AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
  AND (CAST(:title AS TEXT) IS NULL or title = :title)
  AND (CAST(:casting AS JSONB) IS NULL or casting = :casting)
  AND (CAST(:crew AS JSONB) IS NULL or crew = :crew);
`

// language=postgresql
var tMoviesCreditDeleteSql = `
DELETE FROM public.t_movies_credits
WHERE TRUE
  AND movie_id = :movie_id
  AND title = :title
  AND casting = :casting
  AND crew = :crew;
`
