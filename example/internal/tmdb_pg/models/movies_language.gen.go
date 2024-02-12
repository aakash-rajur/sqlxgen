package models

import (
	"fmt"
	"strings"
)

type MoviesLanguage struct {
	MovieId    *int64  `db:"movie_id" json:"movie_id"`
	LanguageId *string `db:"language_id" json:"language_id"`
}

func (m *MoviesLanguage) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("MovieId: %v", *m.MovieId),
			fmt.Sprintf("LanguageId: %v", *m.LanguageId),
		},
		", ",
	)

	return fmt.Sprintf("MoviesLanguage{%s}", content)
}

func (m *MoviesLanguage) TableName() string {
	return "public.movies_languages"
}

func (m *MoviesLanguage) PrimaryKey() []string {
	return []string{
		"movie_id",
		"language_id",
	}
}

func (m *MoviesLanguage) InsertQuery() string {
	return moviesLanguageInsertSql
}

func (m *MoviesLanguage) UpdateAllQuery() string {
	return moviesLanguageUpdateAllSql
}

func (m *MoviesLanguage) UpdateByPkQuery() string {
	return moviesLanguageUpdateByPkSql
}

func (m *MoviesLanguage) CountQuery() string {
	return moviesLanguageModelCountSql
}

func (m *MoviesLanguage) FindAllQuery() string {
	return moviesLanguageFindAllSql
}

func (m *MoviesLanguage) FindFirstQuery() string {
	return moviesLanguageFindFirstSql
}

func (m *MoviesLanguage) FindByPkQuery() string {
	return moviesLanguageFindByPkSql
}

func (m *MoviesLanguage) DeleteByPkQuery() string {
	return moviesLanguageDeleteByPkSql
}

func (m *MoviesLanguage) DeleteAllQuery() string {
	return moviesLanguageDeleteAllSql
}

// language=postgresql
var moviesLanguageAllFieldsWhere = `
WHERE TRUE
    AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
    AND (CAST(:language_id AS TEXT) IS NULL or language_id = :language_id)
`

// language=postgresql
var moviesLanguagePkFieldsWhere = `
WHERE movie_id = :movie_id
  AND language_id = :language_id
`

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
var moviesLanguageUpdateByPkSql = `
UPDATE public.movies_languages
SET
  movie_id = :movie_id,
  language_id = :language_id
` + moviesLanguagePkFieldsWhere + `
RETURNING
  movie_id,
  language_id;
`

// language=postgresql
var moviesLanguageUpdateAllSql = `
UPDATE public.movies_languages
SET
  movie_id = :movie_id,
  language_id = :language_id
` + moviesLanguageAllFieldsWhere + `
RETURNING
  movie_id,
  language_id;
`

// language=postgresql
var moviesLanguageModelCountSql = `
SELECT count(*) as count
FROM public.movies_languages
` + moviesLanguageAllFieldsWhere + ";"

// language=postgresql
var moviesLanguageFindAllSql = `
SELECT
  movie_id,
  language_id
FROM public.movies_languages
` + moviesLanguageAllFieldsWhere + ";"

// language=postgresql
var moviesLanguageFindFirstSql = strings.TrimRight(moviesLanguageFindAllSql, ";") + `
LIMIT 1;`

// language=postgresql
var moviesLanguageFindByPkSql = `
SELECT
  movie_id,
  language_id
FROM public.movies_languages
` + moviesLanguagePkFieldsWhere + `
LIMIT 1;`

// language=postgresql
var moviesLanguageDeleteByPkSql = `
DELETE FROM public.movies_languages
WHERE movie_id = :movie_id
  AND language_id = :language_id;
`

// language=postgresql
var moviesLanguageDeleteAllSql = `
DELETE FROM public.movies_languages
WHERE movie_id = :movie_id
  AND language_id = :language_id;
`
