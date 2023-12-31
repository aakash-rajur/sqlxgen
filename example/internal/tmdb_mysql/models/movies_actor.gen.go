package models

import (
	"fmt"
	"strings"
)

type MoviesActor struct {
	Cast      *string `db:"cast" json:"cast"`
	CastOrder *int32  `db:"cast_order" json:"cast_order"`
	MovieId   *int64  `db:"movie_id" json:"movie_id"`
	ActorId   *int64  `db:"actor_id" json:"actor_id"`
}

func (moviesActor MoviesActor) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Cast: %v", *moviesActor.Cast),
			fmt.Sprintf("CastOrder: %v", *moviesActor.CastOrder),
			fmt.Sprintf("MovieId: %v", *moviesActor.MovieId),
			fmt.Sprintf("ActorId: %v", *moviesActor.ActorId),
		},
		", ",
	)

	return fmt.Sprintf("MoviesActor{%s}", content)
}

func (_ MoviesActor) TableName() string {
	return "app.movies_actors"
}

func (_ MoviesActor) PrimaryKey() []string {
	return []string{
		"movie_id",
		"actor_id",
	}
}

func (_ MoviesActor) InsertQuery() string {
	return moviesActorInsertSql
}

func (_ MoviesActor) UpdateQuery() string {
	return moviesActorUpdateSql
}

func (_ MoviesActor) FindQuery() string {
	return moviesActorFindSql
}

func (_ MoviesActor) FindAllQuery() string {
	return moviesActorFindAllSql
}

func (_ MoviesActor) DeleteQuery() string {
	return moviesActorDeleteSql
}

// language=mysql
var moviesActorInsertSql = `
INSERT INTO app.movies_actors(
  cast,
  cast_order,
  movie_id,
  actor_id
)
VALUES (
  :cast,
  :cast_order,
  :movie_id,
  :actor_id
)
RETURNING
  cast,
  cast_order,
  movie_id,
  actor_id;
`

// language=mysql
var moviesActorUpdateSql = `
UPDATE app.movies_actors
SET
  cast = :cast,
  cast_order = :cast_order,
  movie_id = :movie_id,
  actor_id = :actor_id
WHERE TRUE
  AND movie_id = :movie_id
  AND actor_id = :actor_id
RETURNING
  cast,
  cast_order,
  movie_id,
  actor_id;
`

// language=mysql
var moviesActorFindSql = `
SELECT
  cast,
  cast_order,
  movie_id,
  actor_id
FROM app.movies_actors
WHERE TRUE
  AND (:movie_id IS NULL or movie_id = :movie_id)
  AND (:actor_id IS NULL or actor_id = :actor_id)
LIMIT 1;
`

// language=mysql
var moviesActorFindAllSql = `
SELECT
  cast,
  cast_order,
  movie_id,
  actor_id
FROM app.movies_actors
WHERE TRUE
  AND (:movie_id IS NULL or movie_id = :movie_id)
  AND (:actor_id IS NULL or actor_id = :actor_id);
`

// language=mysql
var moviesActorDeleteSql = `
DELETE FROM app.movies_actors
WHERE TRUE
  AND movie_id = :movie_id
  AND actor_id = :actor_id;
`
