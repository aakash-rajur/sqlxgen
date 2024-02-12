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

func (m *MoviesCrew) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("MovieId: %v", *m.MovieId),
			fmt.Sprintf("CrewId: %v", *m.CrewId),
			fmt.Sprintf("DepartmentId: %v", *m.DepartmentId),
			fmt.Sprintf("JobId: %v", *m.JobId),
		},
		", ",
	)

	return fmt.Sprintf("MoviesCrew{%s}", content)
}

func (m *MoviesCrew) TableName() string {
	return "public.movies_crew"
}

func (m *MoviesCrew) PrimaryKey() []string {
	return []string{
		"movie_id",
		"crew_id",
	}
}

func (m *MoviesCrew) InsertQuery() string {
	return moviesCrewInsertSql
}

func (m *MoviesCrew) UpdateAllQuery() string {
	return moviesCrewUpdateAllSql
}

func (m *MoviesCrew) UpdateByPkQuery() string {
	return moviesCrewUpdateByPkSql
}

func (m *MoviesCrew) CountQuery() string {
	return moviesCrewModelCountSql
}

func (m *MoviesCrew) FindAllQuery() string {
	return moviesCrewFindAllSql
}

func (m *MoviesCrew) FindFirstQuery() string {
	return moviesCrewFindFirstSql
}

func (m *MoviesCrew) FindByPkQuery() string {
	return moviesCrewFindByPkSql
}

func (m *MoviesCrew) DeleteByPkQuery() string {
	return moviesCrewDeleteByPkSql
}

func (m *MoviesCrew) DeleteAllQuery() string {
	return moviesCrewDeleteAllSql
}

// language=postgresql
var moviesCrewAllFieldsWhere = `
WHERE TRUE
    AND (CAST(:movie_id AS INT8) IS NULL or movie_id = :movie_id)
    AND (CAST(:crew_id AS INT8) IS NULL or crew_id = :crew_id)
    AND (CAST(:department_id AS TEXT) IS NULL or department_id = :department_id)
    AND (CAST(:job_id AS TEXT) IS NULL or job_id = :job_id)
`

// language=postgresql
var moviesCrewPkFieldsWhere = `
WHERE movie_id = :movie_id
  AND crew_id = :crew_id
`

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
var moviesCrewUpdateByPkSql = `
UPDATE public.movies_crew
SET
  movie_id = :movie_id,
  crew_id = :crew_id,
  department_id = :department_id,
  job_id = :job_id
` + moviesCrewPkFieldsWhere + `
RETURNING
  movie_id,
  crew_id,
  department_id,
  job_id;
`

// language=postgresql
var moviesCrewUpdateAllSql = `
UPDATE public.movies_crew
SET
  movie_id = :movie_id,
  crew_id = :crew_id,
  department_id = :department_id,
  job_id = :job_id
` + moviesCrewAllFieldsWhere + `
RETURNING
  movie_id,
  crew_id,
  department_id,
  job_id;
`

// language=postgresql
var moviesCrewModelCountSql = `
SELECT count(*) as count
FROM public.movies_crew
` + moviesCrewAllFieldsWhere + ";"

// language=postgresql
var moviesCrewFindAllSql = `
SELECT
  movie_id,
  crew_id,
  department_id,
  job_id
FROM public.movies_crew
` + moviesCrewAllFieldsWhere + ";"

// language=postgresql
var moviesCrewFindFirstSql = strings.TrimRight(moviesCrewFindAllSql, ";") + `
LIMIT 1;`

// language=postgresql
var moviesCrewFindByPkSql = `
SELECT
  movie_id,
  crew_id,
  department_id,
  job_id
FROM public.movies_crew
` + moviesCrewPkFieldsWhere + `
LIMIT 1;`

// language=postgresql
var moviesCrewDeleteByPkSql = `
DELETE FROM public.movies_crew
WHERE movie_id = :movie_id
  AND crew_id = :crew_id;
`

// language=postgresql
var moviesCrewDeleteAllSql = `
DELETE FROM public.movies_crew
WHERE movie_id = :movie_id
  AND crew_id = :crew_id
  AND department_id = :department_id
  AND job_id = :job_id;
`
