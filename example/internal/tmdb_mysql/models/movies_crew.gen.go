package models

import (
	"fmt"
	"strings"
)

type MoviesCrew struct {
	DepartmentId *string `db:"department_id" json:"department_id"`
	JobId        *string `db:"job_id" json:"job_id"`
	MovieId      *int64  `db:"movie_id" json:"movie_id"`
	CrewId       *int64  `db:"crew_id" json:"crew_id"`
}

func (moviesCrew MoviesCrew) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("DepartmentId: %v", *moviesCrew.DepartmentId),
			fmt.Sprintf("JobId: %v", *moviesCrew.JobId),
			fmt.Sprintf("MovieId: %v", *moviesCrew.MovieId),
			fmt.Sprintf("CrewId: %v", *moviesCrew.CrewId),
		},
		", ",
	)

	return fmt.Sprintf("MoviesCrew{%s}", content)
}

func (_ MoviesCrew) TableName() string {
	return "app.movies_crew"
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

// language=mysql
var moviesCrewInsertSql = `
INSERT INTO app.movies_crew(
  department_id,
  job_id,
  movie_id,
  crew_id
)
VALUES (
  :department_id,
  :job_id,
  :movie_id,
  :crew_id
)
RETURNING
  department_id,
  job_id,
  movie_id,
  crew_id;
`

// language=mysql
var moviesCrewUpdateSql = `
UPDATE app.movies_crew
SET
  department_id = :department_id,
  job_id = :job_id,
  movie_id = :movie_id,
  crew_id = :crew_id
WHERE TRUE
  AND movie_id = :movie_id
  AND crew_id = :crew_id
RETURNING
  department_id,
  job_id,
  movie_id,
  crew_id;
`

// language=mysql
var moviesCrewFindSql = `
SELECT
  department_id,
  job_id,
  movie_id,
  crew_id
FROM app.movies_crew
WHERE TRUE
  AND (:movie_id IS NULL or movie_id = :movie_id)
  AND (:crew_id IS NULL or crew_id = :crew_id)
LIMIT 1;
`

// language=mysql
var moviesCrewFindAllSql = `
SELECT
  department_id,
  job_id,
  movie_id,
  crew_id
FROM app.movies_crew
WHERE TRUE
  AND (:movie_id IS NULL or movie_id = :movie_id)
  AND (:crew_id IS NULL or crew_id = :crew_id);
`

// language=mysql
var moviesCrewDeleteSql = `
DELETE FROM app.movies_crew
WHERE TRUE
  AND movie_id = :movie_id
  AND crew_id = :crew_id;
`
