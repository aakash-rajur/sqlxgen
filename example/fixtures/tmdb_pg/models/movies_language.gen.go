package models

import (
	"fmt"
	"strings"
)

type MoviesLanguage struct {
	MovieId    *int64  `db:"movie_id" json:"movie_id"`
	LanguageId *string `db:"language_id" json:"language_id"`
}

func (moviesLanguage MoviesLanguage) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("MovieId: %v", *moviesLanguage.MovieId),
			fmt.Sprintf("LanguageId: %v", *moviesLanguage.LanguageId),
		},
		", ",
	)

	return fmt.Sprintf("MoviesLanguage{%s}", content)
}

func (_ MoviesLanguage) TableName() string {
	return "public.movies_languages"
}

func (_ MoviesLanguage) PrimaryKey() []string {
	return []string{
		"movie_id",
		"language_id",
	}
}

func (_ MoviesLanguage) InsertQuery() string {
	return moviesLanguageInsertSql
}

func (_ MoviesLanguage) UpdateQuery() string {
	return moviesLanguageUpdateSql
}

func (_ MoviesLanguage) FindQuery() string {
	return moviesLanguageFindSql
}

func (_ MoviesLanguage) FindAllQuery() string {
	return moviesLanguageFindAllSql
}

func (_ MoviesLanguage) DeleteQuery() string {
	return moviesLanguageDeleteSql
}

// language=postgresql
var moviesLanguageInsertSql = `
INSERT INTO public.movies_languages(
  movie_id,
  language_id
)
VALUES (
  :movie_id,
  :language_id
)
RETURNING
  movie_id,
  language_id;
`

// language=postgresql
var moviesLanguageUpdateSql = `
UPDATE public.movies_languages
SET
  movie_id = :movie_id,
  language_id = :language_id
WHERE TRUE
  AND movie_id = :movie_id
  AND language_id = :language_id
RETURNING
  movie_id,
  language_id;
`

// language=postgresql
var moviesLanguageFindSql = `
SELECT
  movie_id,
  language_id
FROM public.movies_languages
WHERE TRUE
  AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
  AND (CAST(:language_id AS TEXT) IS NULL or language_id = :language_id)
LIMIT 1;
`

// language=postgresql
var moviesLanguageFindAllSql = `
SELECT
  movie_id,
  language_id
FROM public.movies_languages
WHERE TRUE
  AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
  AND (CAST(:language_id AS TEXT) IS NULL or language_id = :language_id);
`

// language=postgresql
var moviesLanguageDeleteSql = `
DELETE FROM public.movies_languages
WHERE TRUE
  AND movie_id = :movie_id
  AND language_id = :language_id;
`
