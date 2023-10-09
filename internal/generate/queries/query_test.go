package queries

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/aakash-rajur/sqlxgen/internal/generate/types"
	"github.com/aakash-rajur/sqlxgen/internal/introspect"
	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/aakash-rajur/sqlxgen/internal/utils/writer"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/assert"
)

func TestNewQueryModel(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()

	ft := types.NewFakeTranslate("", "")

	queries, err := utils.FromJson[introspect.Query](
		[]string{listActorQueryJson, getActorQueryJson, listMoviesQueryJson, getMovieQueryJson},
	)

	if err != nil {
		t.Fatalf("failed to parse queries: %v", err)
	}

	testCases := []struct {
		name  string
		query introspect.Query
	}{
		{
			name:  "list-actor-query",
			query: queries[0],
		},
		{
			name:  "get-actor-query",
			query: queries[1],
		},
		{
			name:  "list-movies-query",
			query: queries[2],
		},
		{
			name:  "get-movie-query",
			query: queries[3],
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := newQueryModel(
				nil,
				ft,
				tmpDir,
				"gen/store",
				"github.com/john-doe/gen/store",
				testCase.query,
			)

			assert.Nil(t, err)

			gotJson, err := json.MarshalIndent(got, "", "  ")

			cupaloy.SnapshotT(t, gotJson)
		})
	}
}

func TestQueryModel_GetImports(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()

	ft := types.NewFakeTranslate("", "")

	queries, err := utils.FromJson[introspect.Query](
		[]string{listActorQueryJson, getActorQueryJson, listMoviesQueryJson, getMovieQueryJson},
	)

	if err != nil {
		t.Fatalf("failed to parse queries: %v", err)
	}

	testCases := []struct {
		name  string
		query introspect.Query
		want  []string
	}{
		{
			name:  "list-actor-query",
			query: queries[0],
			want:  []string{"gen/store"},
		},
		{
			name:  "get-actor-query",
			query: queries[1],
			want:  []string{"gen/store"},
		},
		{
			name:  "list-movies-query",
			query: queries[2],
			want:  []string{"time", "gen/store"},
		},
		{
			name:  "get-movie-query",
			query: queries[3],
			want:  []string{"time", "gen/store"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := newQueryModel(
				nil,
				ft,
				tmpDir,
				"gen/store",
				"github.com/john-doe/gen/store",
				testCase.query,
			)

			assert.Nil(t, err)

			assert.Equal(t, testCase.want, got.getImports())
		})
	}
}

func TestQueryModel_Generate(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()

	queries, err := utils.FromJson[introspect.Query](
		[]string{listActorQueryJson, getActorQueryJson, listMoviesQueryJson, getMovieQueryJson},
	)

	if err != nil {
		t.Fatalf("failed to parse queries: %v", err)
	}

	mw := writer.NewMemoryWriters()

	ft := types.NewFakeTranslate(
		"",
		`package {{ .PackageName }}
var content = `+"`"+`{
  "packageName": {{ .PackageName | ToJson }}, 
  "imports": {{ .Imports | ToJson }},
	"query": {{ .Query | ToJson }},
}`+"`",
	)

	testCases := []struct {
		name     string
		query    introspect.Query
		filepath string
	}{
		{
			name:     "list-actor-query",
			query:    queries[0],
			filepath: "internal/api/list_actors.gen.go",
		},
		{
			name:     "get-actor-query",
			query:    queries[1],
			filepath: "internal/api/get_actor.gen.go",
		},
		{
			name:     "list-movies-query",
			query:    queries[2],
			filepath: "internal/api/list_movies.gen.go",
		},
		{
			name:     "get-movie-query",
			query:    queries[3],
			filepath: "internal/api/get_movie.gen.go",
		},
	}

	for i, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			qm, err := newQueryModel(
				mw.Creator,
				ft,
				tmpDir,
				"gen/store",
				"github.com/john-doe/gen/store",
				testCase.query,
			)

			assert.Nil(t, err)

			err = qm.generate(ft.QueryTemplate())

			assert.Nil(t, err)

			got := mw.Writers[i]

			assert.Contains(t, got.FullPath, testCase.filepath)

			cupaloy.SnapshotT(t, got.Content)
		})
	}
}

//go:embed fixtures/list-actor-query.json
var listActorQueryJson string

//go:embed fixtures/get-actor-query.json
var getActorQueryJson string

//go:embed fixtures/list-movies-query.json
var listMoviesQueryJson string

//go:embed fixtures/get-movie-query.json
var getMovieQueryJson string
