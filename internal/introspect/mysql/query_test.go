package mysql

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/aakash-rajur/sqlxgen/internal/utils/array"
	"github.com/aakash-rajur/sqlxgen/internal/utils/fs"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestMsgWithFilename(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		filename string
		msg      string
		want     string
	}{
		{
			name:     "empty filename",
			filename: "",
			msg:      "failed to parse query params",
			want:     ": failed to parse query params",
		},
		{
			name:     "valid filename 1",
			filename: "foo.sql",
			msg:      "failed to parse query params",
			want:     "foo.sql: failed to parse query params",
		},
		{
			name:     "valid filename 2",
			filename: "foo/bar.sql",
			msg:      "failed to parse query params",
			want:     "foo/bar.sql: failed to parse query params",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := msgWithFilename(testCase.filename, testCase.msg)

			assert.Equal(t, testCase.want, got)
		})
	}
}

func TestGenerateIntrospectQuery(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		query string
		want  string
		err   error
	}{
		{
			name: "query 1",
			//language=mysql
			query: `select * from users`,
			//language=mysql
			want: `--
drop table if exists sample_query_introspection;
--
create table if not exists sample_query_introspection
select * from users;
--
select
c.column_name as column_name,
c.data_type as type,
'0' as type_id,
false as is_array,
false as is_sequence,
c.is_nullable = 'YES' as nullable,
c.generation_expression != '' as "generated"
from information_schema.columns c
left join information_schema.key_column_usage kc on (
  true
  and kc.table_schema = c.table_schema
  and kc.table_name = c.table_name
  and kc.column_name = c.column_name
  and kc.constraint_name = 'PRIMARY'
)
where true
and c.table_schema = 'public'
and c.table_name = 'sample_query_introspection'
order by c.column_name;
`,
			err: nil,
		},
		{
			name: "query 2",
			//language=mysql
			query: `select * from users where id = :id; -- :id type: int`,
			//language=mysql
			want: `--
drop table if exists sample_query_introspection;
--
create table if not exists sample_query_introspection
select * from users where false and (
id = :id
)
;;
--
select
c.column_name as column_name,
c.data_type as type,
'0' as type_id,
false as is_array,
false as is_sequence,
c.is_nullable = 'YES' as nullable,
c.generation_expression != '' as "generated"
from information_schema.columns c
left join information_schema.key_column_usage kc on (
  true
  and kc.table_schema = c.table_schema
  and kc.table_name = c.table_name
  and kc.column_name = c.column_name
  and kc.constraint_name = 'PRIMARY'
)
where true
and c.table_schema = 'public'
and c.table_name = 'sample_query_introspection'
order by c.column_name;
`,
			err: nil,
		},
		{
			name: "custom 1",
			//language=mysql
			query: `select
count(*) over () as "totalRecordsCount",
m.id as "id",
m.title as "title",
m.release_date as "releaseDate",
m.status as "status",
m.popularity as "popularity"
from movies m
where true
and (
  false
  or :search is null
  or match(m.title, m.original_title, m.keywords) against (:search in natural language mode)
) -- :search type: text
and (
  false
  or :genre_id is null
  or m.id in (
    select
    g.movie_id
    from movies_genres g
    where true
    and g.genre_id = :genre_id -- :genre_id type: text
    order by g.movie_id
  )
)
order by (case when :sort = 'desc' then m.id end) desc, m.id -- :sort type: text
limit :limit -- :limit type: int
offset :offset; -- :offset type: int`,
			//language=mysql
			want: `--
drop table if exists sample_query_introspection;
--
create table if not exists sample_query_introspection
select
count(*) over () as "totalRecordsCount",
m.id as "id",
m.title as "title",
m.release_date as "releaseDate",
m.status as "status",
m.popularity as "popularity"
from movies m
where true
and (
  false
  or :search is null
  or match(m.title, m.original_title, m.keywords) against (:search in natural language mode)
) 
and (
  false
  or :genre_id is null
  or m.id in (
    select
    g.movie_id
    from movies_genres g
    where false and (
true
    and g.genre_id = :genre_id
)
order by g.movie_id
  )
)
order by (case when :sort = 'desc' then m.id end) desc, m.id 
limit :limit 
offset :offset;;
--
select
c.column_name as column_name,
c.data_type as type,
'0' as type_id,
false as is_array,
false as is_sequence,
c.is_nullable = 'YES' as nullable,
c.generation_expression != '' as "generated"
from information_schema.columns c
left join information_schema.key_column_usage kc on (
  true
  and kc.table_schema = c.table_schema
  and kc.table_name = c.table_name
  and kc.column_name = c.column_name
  and kc.constraint_name = 'PRIMARY'
)
where true
and c.table_schema = 'public'
and c.table_name = 'sample_query_introspection'
order by c.column_name;
`,
			err: nil,
		},
		{
			name: "custom 2",
			//language=mysql
			query: `select
m.id as "id",
m.title as "title",
m.original_title as "originalTitle",
m.original_language as "originalLanguage",
m.overview as "overview",
m.runtime as "runtime",
m.release_date as "releaseDate",
m.tagline as "tagline",
m.status as "status",
m.homepage as "homepage",
m.popularity as "popularity",
m.vote_average as "voteAverage",
m.vote_count as "voteCount",
m.budget as "budget",
m.revenue as "revenue",
m.keywords as "keywords",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', g.genre_id,
        'name', hp.friendly_name
      )
    )
    from movies_genres g
    inner join hyper_parameters hp on (
      true
      and hp.type = 'genre'
      and hp.value = g.genre_id
    )
    where true
    and g.movie_id = m.id
    order by hp.friendly_name
  ),
  '[]'
) as "genres",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', c.country_id,
        'name', hp.friendly_name
      )
    )
    from movies_countries c
    inner join hyper_parameters hp on (
      true
      and hp.type = 'country'
      and hp.value = c.country_id
    )
    where true
    and c.movie_id = m.id
    order by hp.friendly_name
  ),
  '[]'
) as "countries",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', l.language_id,
        'name', hp.friendly_name
      )
    )
    from movies_languages l
    inner join hyper_parameters hp on (
      true
      and hp.type = 'language'
      and hp.value = l.language_id
    )
    where true
    and l.movie_id = m.id
    order by hp.friendly_name
  ),
  '[]'
) as "languages",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', mc.company_id,
        'name', c.name
      )
    )
    from movies_companies mc
    inner join companies c on mc.company_id = c.id
    where true
    and mc.movie_id = m.id
    order by c.name
  ),
  '[]'
) as "companies",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', ma.actor_id,
        'name', a.name,
        'character', ma.cast,
        'order', ma.cast_order
      )
    )
    from movies_actors ma
    inner join actors a on ma.actor_id = a.id
    where true
    and ma.movie_id = m.id
    order by ma.cast_order
  ),
  '[]'
) as "actors",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', mc.crew_id,
        'name', c.name,
        'job', j.friendly_name,
        'department', d.friendly_name
      )
    )
    from movies_crew mc
    inner join crew c on mc.crew_id = c.id
    inner join hyper_parameters d on (
      true
      and d.type = 'department'
      and d.value = mc.department_id
    )
    inner join hyper_parameters j on (
      true
      and j.type = 'job'
      and j.value = mc.job_id
    )
    where true
    and mc.movie_id = m.id
    order by j.friendly_name
  ),
  '[]'
) as "crews",
1
from movies m
where true
and m.id = :id; -- :id type: bigint`,
			//language=mysql
			want: `--
drop table if exists sample_query_introspection;
--
create table if not exists sample_query_introspection
select
m.id as "id",
m.title as "title",
m.original_title as "originalTitle",
m.original_language as "originalLanguage",
m.overview as "overview",
m.runtime as "runtime",
m.release_date as "releaseDate",
m.tagline as "tagline",
m.status as "status",
m.homepage as "homepage",
m.popularity as "popularity",
m.vote_average as "voteAverage",
m.vote_count as "voteCount",
m.budget as "budget",
m.revenue as "revenue",
m.keywords as "keywords",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', g.genre_id,
        'name', hp.friendly_name
      )
    )
    from movies_genres g
    inner join hyper_parameters hp on (
      true
      and hp.type = 'genre'
      and hp.value = g.genre_id
    )
    where false and (
true
    and g.movie_id = m.id
)
order by hp.friendly_name
  ),
  '[]'
) as "genres",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', c.country_id,
        'name', hp.friendly_name
      )
    )
    from movies_countries c
    inner join hyper_parameters hp on (
      true
      and hp.type = 'country'
      and hp.value = c.country_id
    )
    where false and (
true
    and c.movie_id = m.id
)
order by hp.friendly_name
  ),
  '[]'
) as "countries",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', l.language_id,
        'name', hp.friendly_name
      )
    )
    from movies_languages l
    inner join hyper_parameters hp on (
      true
      and hp.type = 'language'
      and hp.value = l.language_id
    )
    where false and (
true
    and l.movie_id = m.id
)
order by hp.friendly_name
  ),
  '[]'
) as "languages",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', mc.company_id,
        'name', c.name
      )
    )
    from movies_companies mc
    inner join companies c on mc.company_id = c.id
    where false and (
true
    and mc.movie_id = m.id
)
order by c.name
  ),
  '[]'
) as "companies",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', ma.actor_id,
        'name', a.name,
        'character', ma.cast,
        'order', ma.cast_order
      )
    )
    from movies_actors ma
    inner join actors a on ma.actor_id = a.id
    where false and (
true
    and ma.movie_id = m.id
)
order by ma.cast_order
  ),
  '[]'
) as "actors",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', mc.crew_id,
        'name', c.name,
        'job', j.friendly_name,
        'department', d.friendly_name
      )
    )
    from movies_crew mc
    inner join crew c on mc.crew_id = c.id
    inner join hyper_parameters d on (
      true
      and d.type = 'department'
      and d.value = mc.department_id
    )
    inner join hyper_parameters j on (
      true
      and j.type = 'job'
      and j.value = mc.job_id
    )
    where false and (
true
    and mc.movie_id = m.id
)
order by j.friendly_name
  ),
  '[]'
) as "crews",
1
from movies m
where true
and m.id = :id;;
--
select
c.column_name as column_name,
c.data_type as type,
'0' as type_id,
false as is_array,
false as is_sequence,
c.is_nullable = 'YES' as nullable,
c.generation_expression != '' as "generated"
from information_schema.columns c
left join information_schema.key_column_usage kc on (
  true
  and kc.table_schema = c.table_schema
  and kc.table_name = c.table_name
  and kc.column_name = c.column_name
  and kc.constraint_name = 'PRIMARY'
)
where true
and c.table_schema = 'public'
and c.table_name = 'sample_query_introspection'
order by c.column_name;
`,
			err: nil,
		},
		{
			name: "custom 3",
			//language=mysql
			query: `select
c.id as "id",
c.name as "name",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', mc.movie_id,
        'title', m.title,
        'releaseDate', m.release_date,
        'job', j.friendly_name,
        'department', d.friendly_name
      )
    )
    from movies_crew mc
    inner join movies m on mc.movie_id = m.id
    inner join hyper_parameters d on (
      true
      and d.type = 'department'
      and d.value = mc.department_id
    )
    inner join hyper_parameters j on (
      true
      and j.type = 'job'
      and j.value = mc.job_id
    )
    where true
    and mc.crew_id = c.id
    order by m.release_date desc
  ),
  '[]'
) as "movies"
from crew c
where c.id = :id; -- :id type: bigint`,
			//language=mysql
			want: `--
drop table if exists sample_query_introspection;
--
create table if not exists sample_query_introspection
select
c.id as "id",
c.name as "name",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', mc.movie_id,
        'title', m.title,
        'releaseDate', m.release_date,
        'job', j.friendly_name,
        'department', d.friendly_name
      )
    )
    from movies_crew mc
    inner join movies m on mc.movie_id = m.id
    inner join hyper_parameters d on (
      true
      and d.type = 'department'
      and d.value = mc.department_id
    )
    inner join hyper_parameters j on (
      true
      and j.type = 'job'
      and j.value = mc.job_id
    )
    where false and (
true
    and mc.crew_id = c.id
)
order by m.release_date desc
  ),
  '[]'
) as "movies"
from crew c
where c.id = :id;;
--
select
c.column_name as column_name,
c.data_type as type,
'0' as type_id,
false as is_array,
false as is_sequence,
c.is_nullable = 'YES' as nullable,
c.generation_expression != '' as "generated"
from information_schema.columns c
left join information_schema.key_column_usage kc on (
  true
  and kc.table_schema = c.table_schema
  and kc.table_name = c.table_name
  and kc.column_name = c.column_name
  and kc.constraint_name = 'PRIMARY'
)
where true
and c.table_schema = 'public'
and c.table_name = 'sample_query_introspection'
order by c.column_name;
`,
			err: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := generateIntrospectQuery(testCase.query)

			assert.Equal(t, testCase.err, err)

			assert.Equal(t, testCase.want, got)
		})
	}
}

func TestIntrospectQuery(t *testing.T) {
	testCases := []struct {
		name       string
		args       QueryArgs
		resultsCsv string
	}{
		{
			name: "list actors",
			args: QueryArgs{
				Query:    listActorsQuery,
				Filename: "list-actors.sql",
				GenDir:   "fixtures",
			},
			resultsCsv: listActorsResultCsv,
		},
		{
			name: "get actor",
			args: QueryArgs{
				Query:    getActorQuery,
				Filename: "get-actor.sql",
				GenDir:   "fixtures",
			},
			resultsCsv: getActorResultCsv,
		},
		{
			name: "list movies",
			args: QueryArgs{
				Query:    listMoviesQuery,
				Filename: "list-movies.sql",
				GenDir:   "fixtures",
			},
			resultsCsv: listMoviesResultCsv,
		},
		{
			name: "get movie",
			args: QueryArgs{
				Query:    getMovieQuery,
				Filename: "get-movie.sql",
				GenDir:   "fixtures",
			},
			resultsCsv: getMoviesResultCsv,
		},
	}

	db, mock, err := utils.NewMockSqlx()

	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()

		if err != nil {
			t.Fatalf("failed to close mock db: %v", err)
		}
	}(db)

	mock.ExpectBegin()

	for _, testCase := range testCases {
		mock.ExpectExec("drop table if exists sample_query_introspection").
			WillReturnResult(
				sqlmock.NewResult(0, 0),
			)

		mock.ExpectExec("create table if not exists sample_query_introspection (.+)").
			WillReturnResult(
				sqlmock.NewResult(0, 0),
			)

		mock.ExpectQuery("select (.+) from information_schema.columns c (.+)").
			WillReturnRows(
				sqlmock.NewRows([]string{"column_name", "type", "type_id", "is_array", "is_sequence", "nullable", "generated"}).
					FromCSVString(testCase.resultsCsv),
			)
	}

	mock.ExpectRollback()

	mock.ExpectClose()

	tx, err := db.Beginx()

	if err != nil {
		t.Fatalf("failed to create transaction: %v", err)
	}

	defer func(tx *sqlx.Tx) {
		err := tx.Rollback()

		if err != nil {
			t.Fatalf("failed to rollback transaction: %v", err)
		}
	}(tx)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := introspectQuery(tx, testCase.args)

			assert.NoError(t, err)

			assert.Equal(t, testCase.args.Filename, got.Filename)

			gotJson, err := json.MarshalIndent(got, "", "  ")

			if err != nil {
				t.Fatalf("failed to marshal query: %v", err)
			}

			cupaloy.SnapshotT(t, gotJson)
		})
	}
}

func TestIntrospectQueries(t *testing.T) {
	type testCase struct {
		name     string
		query    string
		result   string
		filename string
	}

	testCases := []testCase{
		{
			name:     "list actors",
			filename: "list-actors.sql",
			query:    listActorsQuery,
			result:   listActorsResultCsv,
		},
		{
			name:     "get actor",
			filename: "get-actor.sql",
			query:    getActorQuery,
			result:   getActorResultCsv,
		},
		{
			name:     "list movies",
			filename: "list-movies.sql",
			query:    listMoviesQuery,
			result:   listMoviesResultCsv,
		},
		{
			name:     "get movie",
			filename: "get-movie.sql",
			query:    getMovieQuery,
			result:   getMoviesResultCsv,
		},
	}

	fd := fs.NewFakeFileDiscovery(
		array.Map(
			testCases,
			func(testCase testCase, i int) fs.FakeDiscover {
				return fs.FakeDiscover{
					Content:  testCase.query,
					Dir:      "fixtures",
					Filename: testCase.filename,
					FullPath: "fixtures/" + testCase.filename,
				}
			},
		),
	)

	s := NewIntrospect(fd, IntrospectArgs{QueryDirs: []string{"fixtures"}})

	db, mock, err := utils.NewMockSqlx()

	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()

		if err != nil {
			t.Fatalf("failed to close mock db: %v", err)
		}
	}(db)

	mock.ExpectBegin()

	for _, testCase := range testCases {
		mock.ExpectExec("drop table if exists sample_query_introspection").
			WillReturnResult(
				sqlmock.NewResult(0, 0),
			)

		mock.ExpectExec("create table if not exists sample_query_introspection (.+)").
			WillReturnResult(
				sqlmock.NewResult(0, 0),
			)

		mock.ExpectQuery("select (.+) from information_schema.columns c (.+)").
			WillReturnRows(
				sqlmock.NewRows([]string{"column_name", "type", "type_id", "is_array", "is_sequence", "nullable", "generated"}).
					FromCSVString(testCase.result),
			)
	}

	mock.ExpectExec("drop table if exists sample_query_introspection").
		WillReturnResult(
			sqlmock.NewResult(0, 0),
		)

	mock.ExpectRollback()

	mock.ExpectClose()

	tx, err := db.Beginx()

	if err != nil {
		t.Fatalf("failed to create transaction: %v", err)
	}

	defer func(tx *sqlx.Tx) {
		err := tx.Rollback()

		if err != nil {
			t.Fatalf("failed to rollback transaction: %v", err)
		}
	}(tx)

	queries, err := s.IntrospectQueries(tx)

	assert.NoError(t, err)

	for i, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := queries[i]

			assert.Equal(t, testCase.filename, got.Filename)

			gotJson, err := json.MarshalIndent(got, "", "  ")

			if err != nil {
				t.Fatalf("failed to marshal query: %v", err)
			}

			cupaloy.SnapshotT(t, gotJson)
		})
	}
}

//go:embed fixtures/list-actors.sql
var listActorsQuery string

//go:embed fixtures/list-actors.csv
var listActorsResultCsv string

//go:embed fixtures/get-actor.sql
var getActorQuery string

//go:embed fixtures/get-actor.csv
var getActorResultCsv string

//go:embed fixtures/list-movies.sql
var listMoviesQuery string

//go:embed fixtures/list-movies.csv
var listMoviesResultCsv string

//go:embed fixtures/get-movie.sql
var getMovieQuery string

//go:embed fixtures/get-movie.csv
var getMoviesResultCsv string
