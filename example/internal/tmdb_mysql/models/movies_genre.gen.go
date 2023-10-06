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
	return "app.movies_genres"
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

// language=mysql
var moviesGenreInsertSql = `
INSERT INTO app.movies_genres(
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

// language=mysql
var moviesGenreUpdateSql = `
UPDATE app.movies_genres
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

// language=mysql
var moviesGenreFindSql = `
SELECT
  movie_id,
  genre_id
FROM app.movies_genres
WHERE TRUE
  AND (:movie_id IS NULL or movie_id = :movie_id)
  AND (:genre_id IS NULL or genre_id = :genre_id)
LIMIT 1;
`

// language=mysql
var moviesGenreFindAllSql = `
SELECT
  movie_id,
  genre_id
FROM app.movies_genres
WHERE TRUE
  AND (:movie_id IS NULL or movie_id = :movie_id)
  AND (:genre_id IS NULL or genre_id = :genre_id);
`

// language=mysql
var moviesGenreDeleteSql = `
DELETE FROM app.movies_genres
WHERE TRUE
  AND movie_id = :movie_id
  AND genre_id = :genre_id;
`
