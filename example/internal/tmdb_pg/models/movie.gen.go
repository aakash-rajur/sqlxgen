package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"
)

type Movie struct {
	Id                   *int32          `db:"id" json:"id"`
	Budget               *int64          `db:"budget" json:"budget"`
	ClientId             *string         `db:"client_id" json:"client_id"`
	CompletedCoordinates interface{}     `db:"completed_coordinates" json:"completed_coordinates"`
	DataSyncedAt         *time.Time      `db:"data_synced_at" json:"data_synced_at"`
	DistanceToPlace      *string         `db:"distance_to_place" json:"distance_to_place"`
	Homepage             *string         `db:"homepage" json:"homepage"`
	IsCompleted          *bool           `db:"is_completed" json:"is_completed"`
	Keywords             *pq.StringArray `db:"keywords" json:"keywords"`
	KeywordsSearch       *string         `db:"keywords_search" json:"keywords_search"`
	LocationAccuracy     *int32          `db:"location_accuracy" json:"location_accuracy"`
	OriginalLanguage     *string         `db:"original_language" json:"original_language"`
	OriginalTitle        *string         `db:"original_title" json:"original_title"`
	Overview             *string         `db:"overview" json:"overview"`
	Popularity           *float64        `db:"popularity" json:"popularity"`
	ReleaseDate          *time.Time      `db:"release_date" json:"release_date"`
	Revenue              *int64          `db:"revenue" json:"revenue"`
	Runtime              *int32          `db:"runtime" json:"runtime"`
	SearchVector         *string         `db:"search_vector" json:"search_vector"`
	Status               *string         `db:"status" json:"status"`
	Summary              *string         `db:"summary" json:"summary"`
	Synopsis             *string         `db:"synopsis" json:"synopsis"`
	Tagline              *string         `db:"tagline" json:"tagline"`
	Title                *string         `db:"title" json:"title"`
	TitleSearch          *string         `db:"title_search" json:"title_search"`
	VoteAverage          *float64        `db:"vote_average" json:"vote_average"`
	VoteCount            *int32          `db:"vote_count" json:"vote_count"`
}

func (m *Movie) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *m.Id),
			fmt.Sprintf("Budget: %v", *m.Budget),
			fmt.Sprintf("Homepage: %v", *m.Homepage),
			fmt.Sprintf("Keywords: %v", *m.Keywords),
			fmt.Sprintf("KeywordsSearch: %v", *m.KeywordsSearch),
			fmt.Sprintf("OriginalLanguage: %v", *m.OriginalLanguage),
			fmt.Sprintf("OriginalTitle: %v", *m.OriginalTitle),
			fmt.Sprintf("Overview: %v", *m.Overview),
			fmt.Sprintf("Popularity: %v", *m.Popularity),
			fmt.Sprintf("ReleaseDate: %v", *m.ReleaseDate),
			fmt.Sprintf("Revenue: %v", *m.Revenue),
			fmt.Sprintf("Runtime: %v", *m.Runtime),
			fmt.Sprintf("Status: %v", *m.Status),
			fmt.Sprintf("Tagline: %v", *m.Tagline),
			fmt.Sprintf("Title: %v", *m.Title),
			fmt.Sprintf("TitleSearch: %v", *m.TitleSearch),
			fmt.Sprintf("VoteAverage: %v", *m.VoteAverage),
			fmt.Sprintf("VoteCount: %v", *m.VoteCount),
		},
		", ",
	)

	return fmt.Sprintf("Movie{%s}", content)
}

func (m *Movie) TableName() string {
	return "public.movies"
}

func (m *Movie) PrimaryKey() []string {
	return []string{
		"id",
	}
}

func (m *Movie) InsertQuery() string {
	return movieInsertSql
}

func (m *Movie) UpdateQuery() string {
	return movieUpdateSql
}

func (m *Movie) FindQuery() string {
	return movieFindSql
}

func (m *Movie) FindAllQuery() string {
	return movieFindAllSql
}

func (m *Movie) DeleteQuery() string {
	return movieDeleteSql
}

// language=postgresql
var movieInsertSql = `
INSERT INTO public.movies(
  budget,
  client_id,
  completed_coordinates,
  data_synced_at,
  distance_to_place,
  homepage,
  is_completed,
  keywords,
  location_accuracy,
  original_language,
  original_title,
  overview,
  popularity,
  release_date,
  revenue,
  runtime,
  search_vector,
  status,
  summary,
  synopsis,
  tagline,
  title,
  vote_average,
  vote_count
)
VALUES (
  :budget,
  :client_id,
  :completed_coordinates,
  :data_synced_at,
  :distance_to_place,
  :homepage,
  :is_completed,
  :keywords,
  :location_accuracy,
  :original_language,
  :original_title,
  :overview,
  :popularity,
  :release_date,
  :revenue,
  :runtime,
  :search_vector,
  :status,
  :summary,
  :synopsis,
  :tagline,
  :title,
  :vote_average,
  :vote_count
)
RETURNING
  id,
  budget,
  client_id,
  completed_coordinates,
  data_synced_at,
  distance_to_place,
  homepage,
  is_completed,
  keywords,
  keywords_search,
  location_accuracy,
  original_language,
  original_title,
  overview,
  popularity,
  release_date,
  revenue,
  runtime,
  search_vector,
  status,
  summary,
  synopsis,
  tagline,
  title,
  title_search,
  vote_average,
  vote_count;
`

// language=postgresql
var movieUpdateSql = `
UPDATE public.movies
SET
  id = :id,
  budget = :budget,
  client_id = :client_id,
  completed_coordinates = :completed_coordinates,
  data_synced_at = :data_synced_at,
  distance_to_place = :distance_to_place,
  homepage = :homepage,
  is_completed = :is_completed,
  keywords = :keywords,
  location_accuracy = :location_accuracy,
  original_language = :original_language,
  original_title = :original_title,
  overview = :overview,
  popularity = :popularity,
  release_date = :release_date,
  revenue = :revenue,
  runtime = :runtime,
  search_vector = :search_vector,
  status = :status,
  summary = :summary,
  synopsis = :synopsis,
  tagline = :tagline,
  title = :title,
  vote_average = :vote_average,
  vote_count = :vote_count
WHERE TRUE
  AND id = :id
RETURNING
  id,
  budget,
  client_id,
  completed_coordinates,
  data_synced_at,
  distance_to_place,
  homepage,
  is_completed,
  keywords,
  keywords_search,
  location_accuracy,
  original_language,
  original_title,
  overview,
  popularity,
  release_date,
  revenue,
  runtime,
  search_vector,
  status,
  summary,
  synopsis,
  tagline,
  title,
  title_search,
  vote_average,
  vote_count;
`

// language=postgresql
var movieFindSql = `
SELECT
  id,
  budget,
  client_id,
  completed_coordinates,
  data_synced_at,
  distance_to_place,
  homepage,
  is_completed,
  keywords,
  keywords_search,
  location_accuracy,
  original_language,
  original_title,
  overview,
  popularity,
  release_date,
  revenue,
  runtime,
  search_vector,
  status,
  summary,
  synopsis,
  tagline,
  title,
  title_search,
  vote_average,
  vote_count
FROM public.movies
WHERE TRUE
  AND (CAST(:id AS INT4) IS NULL or id = :id)
  AND (CAST(:budget AS INT8) IS NULL or budget = :budget)
  AND (CAST(:homepage AS TEXT) IS NULL or homepage = :homepage)
  AND (CAST(:keywords AS TEXT) IS NULL or keywords = :keywords)
  AND (CAST(:keywords_search AS TSVECTOR) IS NULL or keywords_search = :keywords_search)
  AND (CAST(:original_language AS TEXT) IS NULL or original_language = :original_language)
  AND (CAST(:original_title AS TEXT) IS NULL or original_title = :original_title)
  AND (CAST(:overview AS TEXT) IS NULL or overview = :overview)
  AND (CAST(:popularity AS FLOAT8) IS NULL or popularity = :popularity)
  AND (CAST(:release_date AS DATE) IS NULL or release_date = :release_date)
  AND (CAST(:revenue AS INT8) IS NULL or revenue = :revenue)
  AND (CAST(:runtime AS INT4) IS NULL or runtime = :runtime)
  AND (CAST(:status AS TEXT) IS NULL or status = :status)
  AND (CAST(:tagline AS TEXT) IS NULL or tagline = :tagline)
  AND (CAST(:title AS TEXT) IS NULL or title = :title)
  AND (CAST(:title_search AS TSVECTOR) IS NULL or title_search = :title_search)
  AND (CAST(:vote_average AS FLOAT8) IS NULL or vote_average = :vote_average)
  AND (CAST(:vote_count AS INT4) IS NULL or vote_count = :vote_count)
LIMIT 1;
`

// language=postgresql
var movieFindAllSql = `
SELECT
  id,
  budget,
  client_id,
  completed_coordinates,
  data_synced_at,
  distance_to_place,
  homepage,
  is_completed,
  keywords,
  keywords_search,
  location_accuracy,
  original_language,
  original_title,
  overview,
  popularity,
  release_date,
  revenue,
  runtime,
  search_vector,
  status,
  summary,
  synopsis,
  tagline,
  title,
  title_search,
  vote_average,
  vote_count
FROM public.movies
WHERE TRUE
  AND (CAST(:id AS INT4) IS NULL or id = :id)
  AND (CAST(:budget AS INT8) IS NULL or budget = :budget)
  AND (CAST(:homepage AS TEXT) IS NULL or homepage = :homepage)
  AND (CAST(:keywords AS TEXT) IS NULL or keywords = :keywords)
  AND (CAST(:keywords_search AS TSVECTOR) IS NULL or keywords_search = :keywords_search)
  AND (CAST(:original_language AS TEXT) IS NULL or original_language = :original_language)
  AND (CAST(:original_title AS TEXT) IS NULL or original_title = :original_title)
  AND (CAST(:overview AS TEXT) IS NULL or overview = :overview)
  AND (CAST(:popularity AS FLOAT8) IS NULL or popularity = :popularity)
  AND (CAST(:release_date AS DATE) IS NULL or release_date = :release_date)
  AND (CAST(:revenue AS INT8) IS NULL or revenue = :revenue)
  AND (CAST(:runtime AS INT4) IS NULL or runtime = :runtime)
  AND (CAST(:status AS TEXT) IS NULL or status = :status)
  AND (CAST(:tagline AS TEXT) IS NULL or tagline = :tagline)
  AND (CAST(:title AS TEXT) IS NULL or title = :title)
  AND (CAST(:title_search AS TSVECTOR) IS NULL or title_search = :title_search)
  AND (CAST(:vote_average AS FLOAT8) IS NULL or vote_average = :vote_average)
  AND (CAST(:vote_count AS INT4) IS NULL or vote_count = :vote_count);
`

// language=postgresql
var movieDeleteSql = `
DELETE FROM public.movies
WHERE TRUE
  AND id = :id
  AND budget = :budget
  AND homepage = :homepage
  AND keywords = :keywords
  AND keywords_search = :keywords_search
  AND original_language = :original_language
  AND original_title = :original_title
  AND overview = :overview
  AND popularity = :popularity
  AND release_date = :release_date
  AND revenue = :revenue
  AND runtime = :runtime
  AND status = :status
  AND tagline = :tagline
  AND title = :title
  AND title_search = :title_search
  AND vote_average = :vote_average
  AND vote_count = :vote_count;
`
