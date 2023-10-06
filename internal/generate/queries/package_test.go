package queries

import (
	"encoding/json"
	"path"
	"testing"

	"github.com/aakash-rajur/sqlxgen/internal/generate/types"
	"github.com/aakash-rajur/sqlxgen/internal/introspect"
	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/aakash-rajur/sqlxgen/internal/utils/writer"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/assert"
)

func TestNewPackage(t *testing.T) {
	tmpDir := t.TempDir()

	queries, err := utils.FromJson[introspect.Query](
		[]string{listActorQueryJson, getActorQueryJson, listMoviesQueryJson, getMovieQueryJson},
	)

	if err != nil {
		t.Fatalf("failed to parse queries: %v", err)
	}

	ft := types.NewFakeTranslate(
		"",
		`package {{ .PackageName }}
var content = `+"`"+`{
  "packageName": {{ .PackageName | ToJson }}, 
  "imports": {{ .Imports | ToJson }},
	"query": {{ .Query | ToJson }},
}`+"`",
	)

	got, err := NewPackage(
		nil,
		ft,
		tmpDir,
		"internal/store",
		"internal/store",
		queries,
	)

	assert.Nil(t, err)

	assert.Equal(t, "internal/store", got.StorePackageDir)

	gotJson, err := json.MarshalIndent(got, "", "  ")

	if err != nil {
		t.Fatalf("failed to marshal package: %v", err)
	}

	cupaloy.SnapshotT(t, gotJson)
}

func TestPackage_Generate(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()

	mw := writer.NewMemoryWriters()

	queries, err := utils.FromJson[introspect.Query](
		[]string{listActorQueryJson, getActorQueryJson, listMoviesQueryJson, getMovieQueryJson},
	)

	if err != nil {
		t.Fatalf("failed to parse queries: %v", err)
	}

	ft := types.NewFakeTranslate(
		"",
		`package {{ .PackageName }}
var content = `+"`"+`{
  "packageName": {{ .PackageName | ToJson }}, 
  "imports": {{ .Imports | ToJson }},
	"query": {{ .Query | ToJson }},
}`+"`",
	)

	pkg, err := NewPackage(
		mw.Creator,
		ft,
		tmpDir,
		"internal/store",
		"internal/store",
		queries,
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = pkg.Generate()

	assert.Nil(t, err)

	paths := []string{
		"/internal/api/list_actors.gen.go",
		"/internal/api/get_actor.gen.go",
		"/internal/api/list_movies.gen.go",
		"/internal/api/get_movie.gen.go",
	}

	for i, p := range paths {
		testName, _ := utils.SplitFilename(path.Base(p))

		t.Run(testName, func(t *testing.T) {
			got := mw.Writers[i]

			assert.Contains(t, got.FullPath, p)

			cupaloy.SnapshotT(t, got.Content)
		})
	}
}
