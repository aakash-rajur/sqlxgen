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
	return "app.movies_genres"
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

func (m *MoviesGenre) DeleteManyQuery() string {
	return moviesGenreDeleteManySql
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

// language=mysql
var moviesGenreDeleteManySql = `
DELETE FROM app.movies_genres
WHERE TRUE
  AND (:movie_id IS NULL or movie_id = :movie_id)
  AND (:genre_id IS NULL or genre_id = :genre_id);
`
