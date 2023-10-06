package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetJsonType(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		jsonFn string
		want   string
	}{
		{
			name:   "json_arrayagg",
			jsonFn: "json_arrayagg",
			want:   "array",
		},
		{
			name:   "json_objectagg",
			jsonFn: "json_objectagg",
			want:   "object",
		},
		{
			name:   "json_object",
			jsonFn: "json_object",
			want:   "object",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.want, getJsonType(testCase.jsonFn))
		})
	}
}

func TestParseJsonColumns(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		query string
		want  map[string]string
		err   error
	}{
		{
			name: "simple json",
			//language=mysql
			query: `
				select
					json_object(
						'foo', foo,
						'bar', bar
					) as "baz"
				from foo
			`,
			want: map[string]string{
				"baz": "json_object",
			},
			err: nil,
		},
		{
			name: "simple json with underscore",
			//language=mysql
			query: `
				select
					json_object(
						'foo', foo,
						'bar', bar
					) as "baz_bat"
				from foo
			`,
			want: map[string]string{
				"baz_bat": "json_object",
			},
			err: nil,
		},
		{
			name: "simple json with underscore and json",
			//language=mysql
			query: `
				select
					json_object(
						'foo', foo,
						'bar', bar
					) as "baz_bat_json"
				from foo
			`,
			want: map[string]string{
				"baz_bat_json": "json_object",
			},
			err: nil,
		},
		{
			name: "custom 1",
			//language=mysql
			query: `
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
			want: map[string]string{},
			err:  nil,
		},
		{
			name: "custom 2",
			//language=mysql
			query: `
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
    where true
    and g.movie_id = m.id
    order by hp.friendly_name
  ),
  cast('[]' as json)
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
  cast('[]' as json)
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
  cast('[]' as json)
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
  cast('[]' as json)
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
  cast('[]' as json)
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
  cast('[]' as json)
) as "crews",
1
from movies m
where true
and m.id = :id; -- :id type: bigint
`,
			want: map[string]string{
				"actors":    "json_arrayagg",
				"companies": "json_arrayagg",
				"countries": "json_arrayagg",
				"crews":     "json_arrayagg",
				"genres":    "json_arrayagg",
				"languages": "json_arrayagg",
			},
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
  cast('[]' as json)
) as "movies"
from crew c
where c.id = :id; -- :id type: bigint`,
			want: map[string]string{"movies": "json_arrayagg"},
			err:  nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseJsonColumns(tt.query)

			assert.Equal(t, tt.err, err)

			assert.Equal(t, tt.want, got)
		})
	}
}
