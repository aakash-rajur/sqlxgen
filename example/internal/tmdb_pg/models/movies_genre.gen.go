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

func (m *MoviesGenre) FindQuery() string {
	return moviesGenreFindSql
}

func (m *MoviesGenre) FindAllQuery() string {
	return moviesGenreFindAllSql
}

func (m *MoviesGenre) DeleteQuery() string {
	return moviesGenreDeleteSql
}

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
var moviesGenreUpdateSql = `
UPDATE public.movies_genres
SET
  movie_id = :movie_id,
  genre_id = :genre_id
WHERE TRUE
  AND movie_id = :movie_id
  AND genre_id = :genre_id
RETURNING
  movie_id,
  genre_id;
`

// language=postgresql
var moviesGenreFindSql = `
SELECT
  movie_id,
  genre_id
FROM public.movies_genres
WHERE TRUE
  AND movie_id = :movie_id
  AND genre_id = :genre_id;
LIMIT 1;
`

// language=postgresql
var moviesGenreFindAllSql = `
SELECT
  movie_id,
  genre_id
FROM public.movies_genres
WHERE TRUE
  AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
  AND (CAST(:genre_id AS TEXT) IS NULL or genre_id = :genre_id);
`

// language=postgresql
var moviesGenreDeleteSql = `
DELETE FROM public.movies_genres
WHERE TRUE
  AND movie_id = :movie_id
  AND genre_id = :genre_id;
`
