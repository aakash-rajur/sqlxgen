package models

import (
	"fmt"
	"strings"
)

type MoviesGenre struct {
	MovieId *int64  `db:"movie_id" json:"movie_id"`
	GenreId *string `db:"genre_id" json:"genre_id"`
}

func (moviesGenre MoviesGenre) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("MovieId: %v", *moviesGenre.MovieId),
			fmt.Sprintf("GenreId: %v", *moviesGenre.GenreId),
		},
		", ",
	)

	return fmt.Sprintf("MoviesGenre{%s}", content)
}

func (_ MoviesGenre) TableName() string {
	return "public.movies_genres"
}

func (_ MoviesGenre) PrimaryKey() []string {
	return []string{
		"movie_id",
		"genre_id",
	}
}

func (_ MoviesGenre) InsertQuery() string {
	return moviesGenreInsertSql
}

func (_ MoviesGenre) UpdateQuery() string {
	return moviesGenreUpdateSql
}

func (_ MoviesGenre) FindQuery() string {
	return moviesGenreFindSql
}

func (_ MoviesGenre) FindAllQuery() string {
	return moviesGenreFindAllSql
}

func (_ MoviesGenre) DeleteQuery() string {
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
  AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
  AND (CAST(:genre_id AS TEXT) IS NULL or genre_id = :genre_id)
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
