package models

import (
	"fmt"
	"strings"
)

type MoviesGenre struct {
	MovieId *int64  `db:"movie_id" json:"movie_id"`
	GenreId *string `db:"genre_id" json:"genre_id"`
}

func (m *MoviesGenre) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("MovieId: %v", *m.MovieId),
			fmt.Sprintf("GenreId: %v", *m.GenreId),
		},
		", ",
	)

	return fmt.Sprintf("MoviesGenre{%s}", content)
}

func (m *MoviesGenre) TableName() string {
	return "public.movies_genres"
}

func (m *MoviesGenre) PrimaryKey() []string {
	return []string{
		"movie_id",
		"genre_id",
	}
}

func (m *MoviesGenre) InsertQuery() string {
	return moviesGenreInsertSql
}

func (m *MoviesGenre) UpdateQuery() string {
	return moviesGenreUpdateSql
}

func (m *MoviesGenre) UpdateByPkQuery() string {
	return moviesGenreUpdateByPkSql
}

func (m *MoviesGenre) CountQuery() string {
	return moviesGenreModelCountSql
}

func (m *MoviesGenre) FindAllQuery() string {
	return moviesGenreFindAllSql
}

func (m *MoviesGenre) FindFirstQuery() string {
	return moviesGenreFindFirstSql
}

func (m *MoviesGenre) FindByPkQuery() string {
	return moviesGenreFindByPkSql
}

func (m *MoviesGenre) DeleteByPkQuery() string {
	return moviesGenreDeleteByPkSql
}

func (m *MoviesGenre) DeleteQuery() string {
	return moviesGenreDeleteSql
}

// language=postgresql
var moviesGenreAllFieldsWhere = `
WHERE TRUE
    AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
    AND (CAST(:genre_id AS TEXT) IS NULL or genre_id = :genre_id)
`

// language=postgresql
var moviesGenrePkFieldsWhere = `
WHERE movie_id = :movie_id
  AND genre_id = :genre_id
`

// language=postgresql
var moviesGenreInsertSql = `
INSERT INTO public.movies_genres(
  movie_id,
  genre_id
)
VALUES (
  :movie_id,
  :genre_id
)
RETURNING
  movie_id,
  genre_id;
`

// language=postgresql
var moviesGenreUpdateByPkSql = `
UPDATE public.movies_genres
SET
  movie_id = :movie_id,
  genre_id = :genre_id
` + moviesGenrePkFieldsWhere + `
RETURNING
  movie_id,
  genre_id;
`

// language=postgresql
var moviesGenreUpdateSql = `
UPDATE public.movies_genres
SET
  movie_id = :movie_id,
  genre_id = :genre_id
` + moviesGenreAllFieldsWhere + `
RETURNING
  movie_id,
  genre_id;
`

// language=postgresql
var moviesGenreModelCountSql = `
SELECT count(*) as count
FROM public.movies_genres
` + moviesGenreAllFieldsWhere + ";"

// language=postgresql
var moviesGenreFindAllSql = `
SELECT
  movie_id,
  genre_id
FROM public.movies_genres
` + moviesGenreAllFieldsWhere + ";"

// language=postgresql
var moviesGenreFindFirstSql = strings.TrimRight(moviesGenreFindAllSql, ";") + `
LIMIT 1;`

// language=postgresql
var moviesGenreFindByPkSql = `
SELECT
  movie_id,
  genre_id
FROM public.movies_genres
` + moviesGenrePkFieldsWhere + `
LIMIT 1;`

// language=postgresql
var moviesGenreDeleteByPkSql = `
DELETE FROM public.movies_genres
WHERE movie_id = :movie_id
  AND genre_id = :genre_id;
`

// language=postgresql
var moviesGenreDeleteSql = `
DELETE FROM public.movies_genres
WHERE movie_id = :movie_id
  AND genre_id = :genre_id;
`
