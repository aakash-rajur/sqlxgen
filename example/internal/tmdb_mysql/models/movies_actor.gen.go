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

func (m *MoviesActor) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Cast: %v", *m.Cast),
			fmt.Sprintf("CastOrder: %v", *m.CastOrder),
			fmt.Sprintf("MovieId: %v", *m.MovieId),
			fmt.Sprintf("ActorId: %v", *m.ActorId),
		},
		", ",
	)

	return fmt.Sprintf("MoviesActor{%s}", content)
}

func (m *MoviesActor) TableName() string {
	return "app.movies_actors"
}

func (m *MoviesActor) PrimaryKey() []string {
	return []string{
		"movie_id",
		"actor_id",
	}
}

func (m *MoviesActor) InsertQuery() string {
	return moviesActorInsertSql
}

func (m *MoviesActor) UpdateQuery() string {
	return moviesActorUpdateSql
}

func (m *MoviesActor) UpdateByPkQuery() string {
	return moviesActorUpdateByPkSql
}

func (m *MoviesActor) CountQuery() string {
	return moviesActorCountSql
}

func (m *MoviesActor) FindAllQuery() string {
	return moviesActorFindAllSql
}

func (m *MoviesActor) FindFirstQuery() string {
	return moviesActorFindFirstSql
}

func (m *MoviesActor) FindByPkQuery() string {
	return moviesActorFindByPkSql
}

func (m *MoviesActor) DeleteByPkQuery() string {
	return moviesActorDeleteByPkSql
}

func (m *MoviesActor) DeleteQuery() string {
	return moviesActorDeleteSql
}

// language=mysql
var moviesActorAllFieldsWhere = `
WHERE (CAST(:cast AS TEXT) IS NULL or cast = :cast)
  AND (CAST(:cast_order AS INT) IS NULL or cast_order = :cast_order)
  AND (CAST(:movie_id AS BIGINT) IS NULL or movie_id = :movie_id)
  AND (CAST(:actor_id AS BIGINT) IS NULL or actor_id = :actor_id)
`

// language=mysql
var moviesActorPkFieldsWhere = `
WHERE movie_id = :movie_id
  AND actor_id = :actor_id
`

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
var moviesActorUpdateByPkSql = `
UPDATE app.movies_actors
SET
  cast = :cast,
  cast_order = :cast_order,
  movie_id = :movie_id,
  actor_id = :actor_id
` + moviesActorPkFieldsWhere + `
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
` + moviesActorAllFieldsWhere + `
RETURNING
  cast,
  cast_order,
  movie_id,
  actor_id;
`

// language=mysql
var moviesActorCountSql = `
SELECT count(*) as count
FROM app.movies_actors
` + moviesActorAllFieldsWhere + ";"

// language=mysql
var moviesActorFindAllSql = `
SELECT
  cast,
  cast_order,
  movie_id,
  actor_id
FROM app.movies_actors
` + moviesActorAllFieldsWhere + ";"

// language=mysql
var moviesActorFindFirstSql = strings.TrimRight(moviesActorFindAllSql, ";") + `
LIMIT 1;`

// language=mysql
var moviesActorFindByPkSql = `
SELECT
  cast,
  cast_order,
  movie_id,
  actor_id
FROM app.movies_actors
` + moviesActorPkFieldsWhere + `
LIMIT 1;`

// language=mysql
var moviesActorDeleteByPkSql = `
DELETE FROM app.movies_actors
WHERE movie_id = :movie_id
  AND actor_id = :actor_id;
`

// language=mysql
var moviesActorDeleteSql = `
DELETE FROM app.movies_actors
WHERE cast = :cast
  AND cast_order = :cast_order
  AND movie_id = :movie_id
  AND actor_id = :actor_id;
`
