package models

import (
	"fmt"
	"strings"
)

type MoviesActor struct {
	MovieId         *int64  `db:"movie_id" json:"movie_id"`
	ActorId         *int64  `db:"actor_id" json:"actor_id"`
	CastOrder       *int32  `db:"cast_order" json:"cast_order"`
	Character       *string `db:"character" json:"character"`
	CharacterSearch *string `db:"character_search" json:"character_search"`
}

func (moviesActor MoviesActor) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("MovieId: %v", *moviesActor.MovieId),
			fmt.Sprintf("ActorId: %v", *moviesActor.ActorId),
			fmt.Sprintf("CastOrder: %v", *moviesActor.CastOrder),
			fmt.Sprintf("Character: %v", *moviesActor.Character),
			fmt.Sprintf("CharacterSearch: %v", *moviesActor.CharacterSearch),
		},
		", ",
	)

	return fmt.Sprintf("MoviesActor{%s}", content)
}

func (_ MoviesActor) TableName() string {
	return "public.movies_actors"
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

// language=postgresql
var moviesActorInsertSql = `
INSERT INTO public.movies_actors(
  movie_id,
  actor_id,
  cast_order,
  character
)
VALUES (
  :movie_id,
  :actor_id,
  :cast_order,
  :character
)
RETURNING
  movie_id,
  actor_id,
  cast_order,
  character,
  character_search;
`

// language=postgresql
var moviesActorUpdateSql = `
UPDATE public.movies_actors
SET
  movie_id = :movie_id,
  actor_id = :actor_id,
  cast_order = :cast_order,
  character = :character
WHERE TRUE
  AND movie_id = :movie_id
  AND actor_id = :actor_id
RETURNING
  movie_id,
  actor_id,
  cast_order,
  character,
  character_search;
`

// language=postgresql
var moviesActorFindSql = `
SELECT
  movie_id,
  actor_id,
  cast_order,
  character,
  character_search
FROM public.movies_actors
WHERE TRUE
  AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
  AND (CAST(:actor_id AS INT8) IS NULL or actor_id = :actor_id)
LIMIT 1;
`

// language=postgresql
var moviesActorFindAllSql = `
SELECT
  movie_id,
  actor_id,
  cast_order,
  character,
  character_search
FROM public.movies_actors
WHERE TRUE
  AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
  AND (CAST(:actor_id AS INT8) IS NULL or actor_id = :actor_id);
`

// language=postgresql
var moviesActorDeleteSql = `
DELETE FROM public.movies_actors
WHERE TRUE
  AND movie_id = :movie_id
  AND actor_id = :actor_id;
`
