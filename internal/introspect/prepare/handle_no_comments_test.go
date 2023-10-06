package prepare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleNoComments(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		query string
		want  string
		err   error
	}{
		{
			name:  "no comments",
			query: `select * from users`,
			want:  `select * from users`,
			err:   nil,
		},
		{
			name:  "comment 1",
			query: `select * from users; -- this is a comment`,
			want:  `select * from users;`,
			err:   nil,
		},
		{
			name: "comment 2",
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
offset :offset; -- :offset type: int
`,
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
    where true
    and g.genre_id = :genre_id 
    order by g.movie_id
  )
)
order by (case when :sort = 'desc' then m.id end) desc, m.id 
limit :limit 
offset :offset;`,
			err: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := handleNoComments(testCase.query)

			assert.Nil(t, err)

			assert.Equal(t, testCase.want, got)
		})
	}
}
