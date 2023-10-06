package introspect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseParams(t *testing.T) {
	t.Parallel()

	type args struct {
		query string
	}

	testCases := []struct {
		name string
		args args
		want []Column
		err  error
	}{
		{
			name: "should parse params 1",
			args: args{
				query: `SELECT * FROM users WHERE id = :id; -- :id type: int`,
			},
			want: []Column{
				{
					ColumnName:        "id",
					Type:              "int",
					TypeId:            "0",
					Nullable:          true,
					PkName:            "",
					PkOrdinalPosition: 0,
					JsonType:          "identity",
				},
			},
		},
		{
			name: "should parse params 2",
			args: args{
				query: `
select
count(*) over (partition by hp.type) as "totalRecordsCount",
hp.value as "value",
hp.friendly_name as "friendlyName"
from hyper_parameter hp
where true
and hp.type = :type -- :type type: text
and (cast(:value as text) is null or hp.value = :value) -- :value type: text
and (cast(:search as text) is null or hp.friendly_name_search @@ to_tsquery(:search)) -- :search type: text
order by hp.friendly_name
limit :limit -- :limit type: int
offset :offset; -- :offset type: int
`,
			},
			want: []Column{
				{
					ColumnName:        "limit",
					Type:              "int",
					TypeId:            "0",
					Nullable:          true,
					PkName:            "",
					PkOrdinalPosition: 0,
					JsonType:          "identity",
				},
				{
					ColumnName:        "offset",
					Type:              "int",
					TypeId:            "0",
					Nullable:          true,
					PkName:            "",
					PkOrdinalPosition: 0,
					JsonType:          "identity",
				},
				{
					ColumnName:        "search",
					Type:              "text",
					TypeId:            "0",
					Nullable:          true,
					PkName:            "",
					PkOrdinalPosition: 0,
					JsonType:          "identity",
				},
				{
					ColumnName:        "type",
					Type:              "text",
					TypeId:            "0",
					Nullable:          true,
					PkName:            "",
					PkOrdinalPosition: 0,
					JsonType:          "identity",
				},
				{
					ColumnName:        "value",
					Type:              "text",
					TypeId:            "0",
					Nullable:          true,
					PkName:            "",
					PkOrdinalPosition: 0,
					JsonType:          "identity",
				},
			},
		},
		{
			name: "should parse params 3",
			args: args{
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
			},
			want: []Column{
				{
					ColumnName:        "genre_id",
					Type:              "text",
					TypeId:            "0",
					Nullable:          true,
					PkName:            "",
					PkOrdinalPosition: 0,
					JsonType:          "identity",
				},
				{
					ColumnName:        "limit",
					Type:              "int",
					TypeId:            "0",
					Nullable:          true,
					PkName:            "",
					PkOrdinalPosition: 0,
					JsonType:          "identity",
				},
				{
					ColumnName:        "offset",
					Type:              "int",
					TypeId:            "0",
					Nullable:          true,
					PkName:            "",
					PkOrdinalPosition: 0,
					JsonType:          "identity",
				},
				{
					ColumnName:        "search",
					Type:              "text",
					TypeId:            "0",
					Nullable:          true,
					PkName:            "",
					PkOrdinalPosition: 0,
					JsonType:          "identity",
				},
				{
					ColumnName:        "sort",
					Type:              "text",
					TypeId:            "0",
					Nullable:          true,
					PkName:            "",
					PkOrdinalPosition: 0,
					JsonType:          "identity",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseParams(testCase.args.query)

			assert.Nil(t, err)

			assert.Equal(t, testCase.want, got)
		})
	}
}

func TestParseParamsAsNil(t *testing.T) {
	t.Parallel()

	type args struct {
		query string
	}

	testCases := []struct {
		name string
		args args
		want map[string]any
		err  error
	}{
		{
			name: "should parse params as nil 1",
			args: args{
				query: `SELECT * FROM users WHERE id = :id; -- :id type: int`,
			},
			want: map[string]any{
				"id": nil,
			},
			err: nil,
		},
		{
			name: "should parse params as nil 2",
			args: args{
				query: `
select
count(*) over (partition by hp.type) as "totalRecordsCount",
hp.value as "value",
hp.friendly_name as "friendlyName"
from hyper_parameter hp
where true
and hp.type = :type -- :type type: text
and (cast(:value as text) is null or hp.value = :value) -- :value type: text
and (cast(:search as text) is null or hp.friendly_name_search @@ to_tsquery(:search)) -- :search type: text
order by hp.friendly_name
limit :limit -- :limit type: int
offset :offset; -- :offset type: int
`,
			},
			want: map[string]any{
				"limit":  nil,
				"offset": nil,
				"search": nil,
				"type":   nil,
				"value":  nil,
			},
			err: nil,
		},
		{
			name: "should parse params as nil 3",
			args: args{
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
			},
			want: map[string]any{
				"genre_id": nil,
				"limit":    nil,
				"offset":   nil,
				"search":   nil,
				"sort":     nil,
			},
			err: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseParamsAsNil(testCase.args.query)

			assert.Nil(t, err)

			assert.Equal(t, testCase.want, got)
		})
	}
}
