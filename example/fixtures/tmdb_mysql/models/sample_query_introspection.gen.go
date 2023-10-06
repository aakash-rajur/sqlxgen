package models

import (
	"fmt"
	"strings"
	"time"
)

type SampleQueryIntrospection struct {
	TotalRecordsCount *int64     `db:"totalRecordsCount" json:"totalRecordsCount"`
	Id                *int64     `db:"id" json:"id"`
	Title             *string    `db:"title" json:"title"`
	ReleaseDate       *time.Time `db:"releaseDate" json:"releaseDate"`
	Status            *string    `db:"status" json:"status"`
	Popularity        *float32   `db:"popularity" json:"popularity"`
}

func (sampleQueryIntrospection SampleQueryIntrospection) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("TotalRecordsCount: %v", *sampleQueryIntrospection.TotalRecordsCount),
			fmt.Sprintf("Id: %v", *sampleQueryIntrospection.Id),
			fmt.Sprintf("Title: %v", *sampleQueryIntrospection.Title),
			fmt.Sprintf("ReleaseDate: %v", *sampleQueryIntrospection.ReleaseDate),
			fmt.Sprintf("Status: %v", *sampleQueryIntrospection.Status),
			fmt.Sprintf("Popularity: %v", *sampleQueryIntrospection.Popularity),
		},
		", ",
	)

	return fmt.Sprintf("SampleQueryIntrospection{%s}", content)
}

func (_ SampleQueryIntrospection) TableName() string {
	return "app.sample_query_introspection"
}

func (_ SampleQueryIntrospection) PrimaryKey() []string {
	return []string{
		"totalRecordsCount",
		"id",
		"title",
		"releaseDate",
		"status",
		"popularity",
	}
}

func (_ SampleQueryIntrospection) InsertQuery() string {
	return sampleQueryIntrospectionInsertSql
}

func (_ SampleQueryIntrospection) UpdateQuery() string {
	return sampleQueryIntrospectionUpdateSql
}

func (_ SampleQueryIntrospection) FindQuery() string {
	return sampleQueryIntrospectionFindSql
}

func (_ SampleQueryIntrospection) FindAllQuery() string {
	return sampleQueryIntrospectionFindAllSql
}

func (_ SampleQueryIntrospection) DeleteQuery() string {
	return sampleQueryIntrospectionDeleteSql
}

// language=mysql
var sampleQueryIntrospectionInsertSql = `
INSERT INTO app.sample_query_introspection(
  totalRecordsCount,
  id,
  title,
  releaseDate,
  status,
  popularity
)
VALUES (
  :totalRecordsCount,
  :id,
  :title,
  :releaseDate,
  :status,
  :popularity
)
RETURNING
  totalRecordsCount,
  id,
  title,
  releaseDate,
  status,
  popularity;
`

// language=mysql
var sampleQueryIntrospectionUpdateSql = `
UPDATE app.sample_query_introspection
SET
  totalRecordsCount = :totalRecordsCount,
  id = :id,
  title = :title,
  releaseDate = :releaseDate,
  status = :status,
  popularity = :popularity
WHERE TRUE
  AND totalRecordsCount = :totalRecordsCount
  AND id = :id
  AND title = :title
  AND releaseDate = :releaseDate
  AND status = :status
  AND popularity = :popularity
RETURNING
  totalRecordsCount,
  id,
  title,
  releaseDate,
  status,
  popularity;
`

// language=mysql
var sampleQueryIntrospectionFindSql = `
SELECT
  totalRecordsCount,
  id,
  title,
  releaseDate,
  status,
  popularity
FROM app.sample_query_introspection
WHERE TRUE
  AND (:totalRecordsCount IS NULL or totalRecordsCount = :totalRecordsCount)
  AND (:id IS NULL or id = :id)
  AND (:title IS NULL or title = :title)
  AND (:releaseDate IS NULL or releaseDate = :releaseDate)
  AND (:status IS NULL or status = :status)
  AND (:popularity IS NULL or popularity = :popularity)
LIMIT 1;
`

// language=mysql
var sampleQueryIntrospectionFindAllSql = `
SELECT
  totalRecordsCount,
  id,
  title,
  releaseDate,
  status,
  popularity
FROM app.sample_query_introspection
WHERE TRUE
  AND (:totalRecordsCount IS NULL or totalRecordsCount = :totalRecordsCount)
  AND (:id IS NULL or id = :id)
  AND (:title IS NULL or title = :title)
  AND (:releaseDate IS NULL or releaseDate = :releaseDate)
  AND (:status IS NULL or status = :status)
  AND (:popularity IS NULL or popularity = :popularity);
`

// language=mysql
var sampleQueryIntrospectionDeleteSql = `
DELETE FROM app.sample_query_introspection
WHERE TRUE
  AND totalRecordsCount = :totalRecordsCount
  AND id = :id
  AND title = :title
  AND releaseDate = :releaseDate
  AND status = :status
  AND popularity = :popularity;
`
