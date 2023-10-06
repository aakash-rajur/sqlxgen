package prepare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllParentheses(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		content string
		want    []parentheses
	}{
		{
			name:    "empty",
			content: "",
			want:    []parentheses{},
		},
		{
			name:    "no parentheses",
			content: "select * from table",
			want:    []parentheses{},
		},
		{
			name:    "single parentheses",
			content: "select * from table where (id = 1)",
			want:    []parentheses{{Start: 26, End: 33}}},
		{
			name:    "multiple parentheses",
			content: "select * from table where (id = 1) and (name = 'test')",
			want:    []parentheses{{Start: 26, End: 33}, {Start: 39, End: 53}}},
		{
			name:    "nested parentheses",
			content: "select * from table where (id = 1 and (name = 'test'))",
			want:    []parentheses{{Start: 38, End: 52}, {Start: 26, End: 53}}},
		{
			name:    "multiple nested parentheses",
			content: "select * from table where (id = 1 and (name = 'test')) and (id = 2 and (name = 'test2'))",
			want: []parentheses{
				{Start: 38, End: 52},
				{Start: 26, End: 53},
				{Start: 71, End: 86},
				{Start: 59, End: 87},
			},
		},
		{
			name:    "multiple nested parentheses with multiple levels",
			content: "select * from table where (id = 1 and (name = 'test' and (id = 2 and (name = 'test2')))) and (id = 2 and (name = 'test2'))",
			want: []parentheses{
				{Start: 69, End: 84},
				{Start: 57, End: 85},
				{Start: 38, End: 86},
				{Start: 26, End: 87},
				{Start: 105, End: 120},
				{Start: 93, End: 121},
			},
		},
		{
			name:    "complex 1",
			content: "select * from table where (id = 1 and (name = 'test' and (id = 2 and (name = 'test2')))) and (id = 2 and (name = 'test2')) and (id = 1 and (name = 'test' and (id = 2 and (name = 'test2')))) and (id = 2 and (name = 'test2'))",
			want: []parentheses{
				{Start: 69, End: 84},
				{Start: 57, End: 85},
				{Start: 38, End: 86},
				{Start: 26, End: 87},
				{Start: 105, End: 120},
				{Start: 93, End: 121},
				{Start: 170, End: 185},
				{Start: 158, End: 186},
				{Start: 139, End: 187},
				{Start: 127, End: 188},
				{Start: 206, End: 221},
				{Start: 194, End: 222},
			},
		},
		{
			name:    "complex 2",
			content: "select * from table where (id = 1 and (name = 'test' and (id = 2 and (name = 'test2')))) and (id = 2 and (name = 'test2')) and (id = 1 and (name = 'test' and (id = 2 and (name = 'test2')))) and (id = 2 and (name = 'test2')) and (id = 1 and (name = 'test' and (id = 2 and (name = 'test2')))) and (id = 2 and (name = 'test2'))",
			want: []parentheses{
				{Start: 69, End: 84},
				{Start: 57, End: 85},
				{Start: 38, End: 86},
				{Start: 26, End: 87},
				{Start: 105, End: 120},
				{Start: 93, End: 121},
				{Start: 170, End: 185},
				{Start: 158, End: 186},
				{Start: 139, End: 187},
				{Start: 127, End: 188},
				{Start: 206, End: 221},
				{Start: 194, End: 222},
				{Start: 271, End: 286},
				{Start: 259, End: 287},
				{Start: 240, End: 288},
				{Start: 228, End: 289},
				{Start: 307, End: 322},
				{Start: 295, End: 323},
			},
		},
		{
			name:    "complex 3",
			content: "select * from table where (id = 1 and (name = 'test' and (id = 2 and (name = 'test2')))) and (id = 2 and (name = 'test2')) and (id = 1 and (name = 'test' and (id = 2 and (name = 'test2')))) and (id = 2 and (name = 'test2')) and (id = 1 and (name = 'test' and (id = 2 and (name = 'test2')))) and (id = 2 and (name = 'test2')) and (id = 1 and (name = 'test' and (id = 2 and (name = 'test2')))) and (id = 2 and (name = 'test2'))",
			want: []parentheses{
				{Start: 69, End: 84},
				{Start: 57, End: 85},
				{Start: 38, End: 86},
				{Start: 26, End: 87},
				{Start: 105, End: 120},
				{Start: 93, End: 121},
				{Start: 170, End: 185},
				{Start: 158, End: 186},
				{Start: 139, End: 187},
				{Start: 127, End: 188},
				{Start: 206, End: 221},
				{Start: 194, End: 222},
				{Start: 271, End: 286},
				{Start: 259, End: 287},
				{Start: 240, End: 288},
				{Start: 228, End: 289},
				{Start: 307, End: 322},
				{Start: 295, End: 323},
				{Start: 372, End: 387},
				{Start: 360, End: 388},
				{Start: 341, End: 389},
				{Start: 329, End: 390},
				{Start: 408, End: 423},
				{Start: 396, End: 424},
			},
		},
		{
			name: "custom 1",
			content: `
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
			want: []parentheses{
				{Start: 13, End: 15},
				{Start: 22, End: 23},
				{Start: 215, End: 231},
				{Start: 274, End: 282},
				{Start: 320, End: 328},
				{Start: 196, End: 330},
				{Start: 377, End: 395},
				{Start: 418, End: 567},
				{Start: 358, End: 569},
				{Start: 580, End: 619},
			},
		},
		{
			name: "custom 2",
			content: `
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
			want: []parentheses{
				{Start: 537, End: 604},
				{Start: 511, End: 636},
				{Start: 700, End: 777},
				{Start: 485, End: 822},
				{Start: 481, End: 832},
				{Start: 911, End: 980},
				{Start: 885, End: 1012},
				{Start: 1079, End: 1160},
				{Start: 859, End: 1205},
				{Start: 855, End: 1215},
				{Start: 1297, End: 1367},
				{Start: 1271, End: 1399},
				{Start: 1466, End: 1549},
				{Start: 1245, End: 1594},
				{Start: 1241, End: 1604},
				{Start: 1686, End: 1746},
				{Start: 1660, End: 1768},
				{Start: 1634, End: 1894},
				{Start: 1630, End: 1904},
				{Start: 1986, End: 2111},
				{Start: 1960, End: 2140},
				{Start: 1934, End: 2258},
				{Start: 1930, End: 2268},
				{Start: 2347, End: 2475},
				{Start: 2321, End: 2497},
				{Start: 2602, End: 2688},
				{Start: 2726, End: 2798},
				{Start: 2295, End: 2844},
				{Start: 2291, End: 2854},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := findAllParentheses(testCase.content)

			assert.Equal(t, testCase.want, actual)
		})
	}
}

func TestFindAllWhereMatches(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		query string
		want  [][]int
		err   error
	}{
		{
			name:  "empty",
			query: "",
			want:  [][]int{},
			err:   nil,
		},
		{
			name:  "no where",
			query: "select * from table",
			want:  [][]int{},
			err:   nil,
		},
		{
			name:  "single where",
			query: "select * from table where id = 1",
			want:  [][]int{{19, 26, 19, 20, 20, 25, 25, 26}},
			err:   nil,
		},
		{
			name:  "multiple where",
			query: "select * from table where id = 1 and name = 'test'",
			want:  [][]int{{19, 26, 19, 20, 20, 25, 25, 26}},
			err:   nil,
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
			want: [][]int{{180, 187, 180, 181, 181, 186, 186, 187}, {474, 481, 474, 475, 475, 480, 480, 481}},
			err:  nil,
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
			want: [][]int{{782, 789, 782, 783, 783, 788, 788, 789}, {1165, 1172, 1165, 1166, 1166, 1171, 1171, 1172}, {1554, 1561, 1554, 1555, 1555, 1560, 1560, 1561}, {1853, 1860, 1853, 1854, 1854, 1859, 1859, 1860}, {2217, 2224, 2217, 2218, 2218, 2223, 2223, 2224}, {2803, 2810, 2803, 2804, 2804, 2809, 2809, 2810}, {2883, 2890, 2883, 2884, 2884, 2889, 2889, 2890}},
			err:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := findAllWhereMatches(testCase.query)

			assert.Equal(t, testCase.want, actual)

			assert.Equal(t, testCase.err, err)
		})
	}
}

func TestSplitWhereContexts(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		query string
		want  []string
		err   error
	}{
		{
			name:  "empty",
			query: "",
			want:  []string{},
			err:   nil,
		},
		{
			name:  "no where",
			query: "select * from table",
			want:  []string{"select * from table"},
			err:   nil,
		},
		{
			name:  "single where",
			query: "select * from table where id = 1",
			want:  []string{"select * from table where id = 1"},
			err:   nil,
		},
		{
			name:  "multiple where",
			query: "select * from table where id = 1 and name = 'test'",
			want:  []string{"select * from table where id = 1 and name = 'test'"},
			err:   nil,
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
			want: []string{
				`where true
    and g.genre_id = :genre_id -- :genre_id type: text
    order by g.movie_id
  `,
			},
			err: nil,
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
			want: []string{
				`where true
    and g.movie_id = m.id
  `,
				`where true
    and c.movie_id = m.id
  `,
				`where true
    and l.movie_id = m.id
  `,
				`where true
    and mc.movie_id = m.id
  `,
				`where true
    and ma.movie_id = m.id
  `,
				`where true
    and mc.movie_id = m.id
  `,
			},
			err: nil,
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
			want: []string{
				`where true
    and mc.crew_id = c.id
  `,
			},
			err: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := splitWhereContexts(testCase.query)

			assert.Equal(t, testCase.want, actual)

			assert.Equal(t, testCase.err, err)
		})
	}
}
