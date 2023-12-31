package models

import (
	"fmt"
	"strings"
)

type MoviesCrew struct {
	MovieId      *int64  `db:"movie_id" json:"movie_id"`
	CrewId       *int64  `db:"crew_id" json:"crew_id"`
	DepartmentId *string `db:"department_id" json:"department_id"`
	JobId        *string `db:"job_id" json:"job_id"`
}

func (moviesCrew MoviesCrew) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("MovieId: %v", *moviesCrew.MovieId),
			fmt.Sprintf("CrewId: %v", *moviesCrew.CrewId),
			fmt.Sprintf("DepartmentId: %v", *moviesCrew.DepartmentId),
			fmt.Sprintf("JobId: %v", *moviesCrew.JobId),
		},
		", ",
	)

	return fmt.Sprintf("MoviesCrew{%s}", content)
}

func (_ MoviesCrew) TableName() string {
	return "public.movies_crew"
}

func (_ MoviesCrew) PrimaryKey() []string {
	return []string{
		"movie_id",
		"crew_id",
	}
}

func (_ MoviesCrew) InsertQuery() string {
	return moviesCrewInsertSql
}

func (_ MoviesCrew) UpdateQuery() string {
	return moviesCrewUpdateSql
}

func (_ MoviesCrew) FindQuery() string {
	return moviesCrewFindSql
}

func (_ MoviesCrew) FindAllQuery() string {
	return moviesCrewFindAllSql
}

func (_ MoviesCrew) DeleteQuery() string {
	return moviesCrewDeleteSql
}

// language=postgresql
var moviesCrewInsertSql = `
INSERT INTO public.movies_crew(
  movie_id,
  crew_id,
  department_id,
  job_id
)
VALUES (
  :movie_id,
  :crew_id,
  :department_id,
  :job_id
)
RETURNING
  movie_id,
  crew_id,
  department_id,
  job_id;
`

// language=postgresql
var moviesCrewUpdateSql = `
UPDATE public.movies_crew
SET
  movie_id = :movie_id,
  crew_id = :crew_id,
  department_id = :department_id,
  job_id = :job_id
WHERE TRUE
  AND movie_id = :movie_id
  AND crew_id = :crew_id
RETURNING
  movie_id,
  crew_id,
  department_id,
  job_id;
`

// language=postgresql
var moviesCrewFindSql = `
SELECT
  movie_id,
  crew_id,
  department_id,
  job_id
FROM public.movies_crew
WHERE TRUE
  AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
  AND (CAST(:crew_id AS INT8) IS NULL or crew_id = :crew_id)
LIMIT 1;
`

// language=postgresql
var moviesCrewFindAllSql = `
SELECT
  movie_id,
  crew_id,
  department_id,
  job_id
FROM public.movies_crew
WHERE TRUE
  AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
  AND (CAST(:crew_id AS INT8) IS NULL or crew_id = :crew_id);
`

// language=postgresql
var moviesCrewDeleteSql = `
DELETE FROM public.movies_crew
WHERE TRUE
  AND movie_id = :movie_id
  AND crew_id = :crew_id;
`
