package models

import (
	"fmt"
	"strings"
	"time"
)

type SampleQueryIntrospection struct {
	Id                *int64     `db:"id" json:"id"`
	Popularity        *float32   `db:"popularity" json:"popularity"`
	ReleaseDate       *time.Time `db:"releaseDate" json:"releaseDate"`
	Status            *string    `db:"status" json:"status"`
	Title             *string    `db:"title" json:"title"`
	TotalRecordsCount *int64     `db:"totalRecordsCount" json:"totalRecordsCount"`
}

func (sampleQueryIntrospection SampleQueryIntrospection) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *sampleQueryIntrospection.Id),
			fmt.Sprintf("Popularity: %v", *sampleQueryIntrospection.Popularity),
			fmt.Sprintf("ReleaseDate: %v", *sampleQueryIntrospection.ReleaseDate),
			fmt.Sprintf("Status: %v", *sampleQueryIntrospection.Status),
			fmt.Sprintf("Title: %v", *sampleQueryIntrospection.Title),
			fmt.Sprintf("TotalRecordsCount: %v", *sampleQueryIntrospection.TotalRecordsCount),
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
		"id",
		"popularity",
		"releaseDate",
		"status",
		"title",
		"totalRecordsCount",
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
  id,
  popularity,
  releaseDate,
  status,
  title,
  totalRecordsCount
)
VALUES (
  :id,
  :popularity,
  :releaseDate,
  :status,
  :title,
  :totalRecordsCount
)
RETURNING
  id,
  popularity,
  releaseDate,
  status,
  title,
  totalRecordsCount;
`

// language=mysql
var sampleQueryIntrospectionUpdateSql = `
UPDATE app.sample_query_introspection
SET
  id = :id,
  popularity = :popularity,
  releaseDate = :releaseDate,
  status = :status,
  title = :title,
  totalRecordsCount = :totalRecordsCount
WHERE TRUE
  AND id = :id
  AND popularity = :popularity
  AND releaseDate = :releaseDate
  AND status = :status
  AND title = :title
  AND totalRecordsCount = :totalRecordsCount
RETURNING
  id,
  popularity,
  releaseDate,
  status,
  title,
  totalRecordsCount;
`

// language=mysql
var sampleQueryIntrospectionFindSql = `
SELECT
  id,
  popularity,
  releaseDate,
  status,
  title,
  totalRecordsCount
FROM app.sample_query_introspection
WHERE TRUE
  AND (:id IS NULL or id = :id)
  AND (:popularity IS NULL or popularity = :popularity)
  AND (:releaseDate IS NULL or releaseDate = :releaseDate)
  AND (:status IS NULL or status = :status)
  AND (:title IS NULL or title = :title)
  AND (:totalRecordsCount IS NULL or totalRecordsCount = :totalRecordsCount)
LIMIT 1;
`

// language=mysql
var sampleQueryIntrospectionFindAllSql = `
SELECT
  id,
  popularity,
  releaseDate,
  status,
  title,
  totalRecordsCount
FROM app.sample_query_introspection
WHERE TRUE
  AND (:id IS NULL or id = :id)
  AND (:popularity IS NULL or popularity = :popularity)
  AND (:releaseDate IS NULL or releaseDate = :releaseDate)
  AND (:status IS NULL or status = :status)
  AND (:title IS NULL or title = :title)
  AND (:totalRecordsCount IS NULL or totalRecordsCount = :totalRecordsCount);
`

// language=mysql
var sampleQueryIntrospectionDeleteSql = `
DELETE FROM app.sample_query_introspection
WHERE TRUE
  AND id = :id
  AND popularity = :popularity
  AND releaseDate = :releaseDate
  AND status = :status
  AND title = :title
  AND totalRecordsCount = :totalRecordsCount;
`
