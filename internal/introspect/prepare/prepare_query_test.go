package prepare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareQuery(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		query string
		want  string
		err   error
	}{
		{
			name:  "no comments",
			query: "select * from users",
			want:  "select * from users",
			err:   nil,
		},
		{
			name:  "where false clause",
			query: "select * from users where false",
			want: `select * from users where false and (
false
)
`,
			err: nil,
		},
		{
			name:  "where false clause with comments",
			query: "select * from users where false -- comment",
			want: `select * from users where false and (
false
)
`,
			err: nil,
		},
		{
			name: "custom 1",
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
  or cast(:search as text) is null
  or m.title_search @@ to_tsquery(:search)
  or m.keywords_search @@ to_tsquery(:search)
) -- :search type: text
and (
  false
  or cast(:genre_id as text) is null
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
			want: `select
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
  or cast(:search as text) is null
  or m.title_search @@ to_tsquery(:search)
  or m.keywords_search @@ to_tsquery(:search)
) 
and (
  false
  or cast(:genre_id as text) is null
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
offset :offset;`,
		},
		{
			name: "custom 2",
			query: `
select
m."id" as "id",
m."title" as "title",
m."original_title" as "originalTitle",
m."original_language" as "originalLanguage",
m."overview" as "overview",
m."runtime" as "runtime",
m."release_date" as "releaseDate",
m."tagline" as "tagline",
m."status" as "status",
m."homepage" as "homepage",
m."popularity" as "popularity",
m."vote_average" as "voteAverage",
m."vote_count" as "voteCount",
m."budget" as "budget",
m."revenue" as "revenue",
m."keywords" as "keywords",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', g.genre_id,
        'name', hp.friendly_name
      ) order by hp.friendly_name
    )
    from movies_genres g
    inner join hyper_parameter hp on (
      true
      and hp.type = 'genre'
      and hp.value = g.genre_id
    )
    where true
    and g.movie_id = m.id
  ),
  '[]'
) as "genres",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', c.country_id,
        'name', hp.friendly_name
      ) order by hp.friendly_name
    )
    from movies_countries c
    inner join hyper_parameter hp on (
      true
      and hp.type = 'country'
      and hp.value = c.country_id
    )
    where true
    and c.movie_id = m.id
  ),
  '[]'
) as "countries",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', l.language_id,
        'name', hp.friendly_name
      ) order by hp.friendly_name
    )
    from movies_languages l
    inner join hyper_parameter hp on (
      true
      and hp.type = 'language'
      and hp.value = l.language_id
    )
    where true
    and l.movie_id = m.id
  ),
  '[]'
) as "languages",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', mc.company_id,
        'name', c.name
      ) order by c.name
    )
    from movies_companies mc
    inner join companies c on mc.company_id = c.id
    where true
    and mc.movie_id = m.id
  ),
  '[]'
) as "companies",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', ma.actor_id,
        'name', a.name,
        'character', ma.character,
        'order', ma.cast_order
      ) order by ma.cast_order
    )
    from movies_actors ma
    inner join actors a on ma.actor_id = a.id
    where true
    and ma.movie_id = m.id
  ),
  '[]'
) as "actors",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', mc.crew_id,
        'name', c.name,
        'job', j.friendly_name,
        'department', d.friendly_name
      ) order by j.name
    )
    from movies_crew mc
    inner join crew c on mc.crew_id = c.id
    inner join hyper_parameter d on (
      true
      and d.type = 'department'
      and d.value = mc.department_id
    )
    inner join hyper_parameter j on (
      true
      and j.type = 'job'
      and j.value = mc.job_id
    )
    where true
    and mc.movie_id = m.id
  ),
  '[]'
) as "crews",
1
from movies m
where true
and m.id = :id; -- :id type: bigint
`,
			want: `select
m."id" as "id",
m."title" as "title",
m."original_title" as "originalTitle",
m."original_language" as "originalLanguage",
m."overview" as "overview",
m."runtime" as "runtime",
m."release_date" as "releaseDate",
m."tagline" as "tagline",
m."status" as "status",
m."homepage" as "homepage",
m."popularity" as "popularity",
m."vote_average" as "voteAverage",
m."vote_count" as "voteCount",
m."budget" as "budget",
m."revenue" as "revenue",
m."keywords" as "keywords",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', g.genre_id,
        'name', hp.friendly_name
      ) order by hp.friendly_name
    )
    from movies_genres g
    inner join hyper_parameter hp on (
      true
      and hp.type = 'genre'
      and hp.value = g.genre_id
    )
    where false and (
true
    and g.movie_id = m.id
)
),
  '[]'
) as "genres",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', c.country_id,
        'name', hp.friendly_name
      ) order by hp.friendly_name
    )
    from movies_countries c
    inner join hyper_parameter hp on (
      true
      and hp.type = 'country'
      and hp.value = c.country_id
    )
    where false and (
true
    and c.movie_id = m.id
)
),
  '[]'
) as "countries",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', l.language_id,
        'name', hp.friendly_name
      ) order by hp.friendly_name
    )
    from movies_languages l
    inner join hyper_parameter hp on (
      true
      and hp.type = 'language'
      and hp.value = l.language_id
    )
    where false and (
true
    and l.movie_id = m.id
)
),
  '[]'
) as "languages",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', mc.company_id,
        'name', c.name
      ) order by c.name
    )
    from movies_companies mc
    inner join companies c on mc.company_id = c.id
    where false and (
true
    and mc.movie_id = m.id
)
),
  '[]'
) as "companies",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', ma.actor_id,
        'name', a.name,
        'character', ma.character,
        'order', ma.cast_order
      ) order by ma.cast_order
    )
    from movies_actors ma
    inner join actors a on ma.actor_id = a.id
    where false and (
true
    and ma.movie_id = m.id
)
),
  '[]'
) as "actors",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', mc.crew_id,
        'name', c.name,
        'job', j.friendly_name,
        'department', d.friendly_name
      ) order by j.name
    )
    from movies_crew mc
    inner join crew c on mc.crew_id = c.id
    inner join hyper_parameter d on (
      true
      and d.type = 'department'
      and d.value = mc.department_id
    )
    inner join hyper_parameter j on (
      true
      and j.type = 'job'
      and j.value = mc.job_id
    )
    where false and (
true
    and mc.movie_id = m.id
)
),
  '[]'
) as "crews",
1
from movies m
where true
and m.id = :id;`,
		},
		{
			name: "custom 3",
			query: `
select
c."id" as "id",
c."name" as "name",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', mc.movie_id,
        'title', m.title,
        'releaseDate', m.release_date,
        'job', j.friendly_name,
        'department', d.friendly_name
      ) order by m.release_date desc
    )
    from movies_crew mc
    inner join movies m on mc.movie_id = m.id
    inner join hyper_parameter d on (
      true
      and d.type = 'department'
      and d.value = mc.department_id
    )
    inner join hyper_parameter j on (
      true
      and j.type = 'job'
      and j.value = mc.job_id
    )
    where true
    and mc.crew_id = c.id
  ),
  '[]'
) as "movies"
from crew c
where c.id = :id; -- :id type: bigint`,
			want: `select
c."id" as "id",
c."name" as "name",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', mc.movie_id,
        'title', m.title,
        'releaseDate', m.release_date,
        'job', j.friendly_name,
        'department', d.friendly_name
      ) order by m.release_date desc
    )
    from movies_crew mc
    inner join movies m on mc.movie_id = m.id
    inner join hyper_parameter d on (
      true
      and d.type = 'department'
      and d.value = mc.department_id
    )
    inner join hyper_parameter j on (
      true
      and j.type = 'job'
      and j.value = mc.job_id
    )
    where false and (
true
    and mc.crew_id = c.id
)
),
  '[]'
) as "movies"
from crew c
where c.id = :id;`,
			err: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := PrepareQuery(tc.query)

			assert.Nil(t, err)

			assert.Equal(t, tc.want, got)
		})
	}
}
