package models

import (
	"fmt"
	"strings"
)

type MoviesActor struct {
	MovieId   *int64  `db:"movie_id" json:"movie_id"`
	ActorId   *int64  `db:"actor_id" json:"actor_id"`
	Cast      *string `db:"cast" json:"cast"`
	CastOrder *int32  `db:"cast_order" json:"cast_order"`
}

func (moviesActor MoviesActor) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("MovieId: %v", *moviesActor.MovieId),
			fmt.Sprintf("ActorId: %v", *moviesActor.ActorId),
			fmt.Sprintf("Cast: %v", *moviesActor.Cast),
			fmt.Sprintf("CastOrder: %v", *moviesActor.CastOrder),
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
  movie_id,
  actor_id,
  cast,
  cast_order
)
VALUES (
  :movie_id,
  :actor_id,
  :cast,
  :cast_order
)
RETURNING
  movie_id,
  actor_id,
  cast,
  cast_order;
`

// language=mysql
var moviesActorUpdateSql = `
UPDATE app.movies_actors
SET
  movie_id = :movie_id,
  actor_id = :actor_id,
  cast = :cast,
  cast_order = :cast_order
WHERE TRUE
  AND movie_id = :movie_id
  AND actor_id = :actor_id
RETURNING
  movie_id,
  actor_id,
  cast,
  cast_order;
`

// language=mysql
var moviesActorFindSql = `
SELECT
  movie_id,
  actor_id,
  cast,
  cast_order
FROM app.movies_actors
WHERE TRUE
  AND (:movie_id IS NULL or movie_id = :movie_id)
  AND (:actor_id IS NULL or actor_id = :actor_id)
LIMIT 1;
`

// language=mysql
var moviesActorFindAllSql = `
SELECT
  movie_id,
  actor_id,
  cast,
  cast_order
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
