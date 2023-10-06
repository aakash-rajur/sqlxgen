package generate

import (
	_ "embed"
	"os"
	"path"
	"testing"

	"github.com/aakash-rajur/sqlxgen/internal/generate/types"
	"github.com/aakash-rajur/sqlxgen/internal/introspect"
	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/aakash-rajur/sqlxgen/internal/utils/writer"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/assert"
)

func TestGenerate_generateStorePackage(t *testing.T) {
	tmpDir := t.TempDir()

	mw := writer.NewMemoryWriters()

	gen, err := createGen(mw.Creator, nil, tmpDir)

	if err != nil {
		t.Fatalf("unable to create generate object: %v", err)
	}

	_, err = gen.generateStorePackage("github.com/aakash-rajur/sqlxgen-example")

	assert.Nil(t, err)

	got := mw.Writers[0]

	storeFullPath := path.Join(tmpDir, "internal/store/store.gen.go")

	assert.Equal(t, storeFullPath, got.FullPath)

	cupaloy.SnapshotT(t, got.Content)
}

func TestGenerate_generateModelPackage(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()

	mw := writer.NewMemoryWriters()

	ft := types.NewFakeTranslate(
		`package {{ .PackageName }}
var content = `+"`"+`{
  "packageName": {{ .PackageName | ToJson }}, 
  "imports": {{ .Imports | ToJson }},
	"model": {{ .Model | ToJson }},
  "insertFields": {{ .InsertFields | ToJson }}
	"updateFields": {{ .UpdateFields | ToJson }}
	"selectFields": {{ .SelectFields | ToJson }}
}`+"`",
		`package {{ .PackageName }}
var content = `+"`"+`{
  "packageName": {{ .PackageName | ToJson }}, 
  "imports": {{ .Imports | ToJson }},
	"query": {{ .Query | ToJson }},
}`+"`",
	)

	gen, err := createGen(mw.Creator, ft, tmpDir)

	if err != nil {
		t.Fatalf("unable to create generate object: %v", err)
	}

	storePkg, err := gen.generateStorePackage("github.com/aakash-rajur/sqlxgen-example")

	assert.Nil(t, err)

	projectPackageName, err := utils.GetProjectPackageName(gen.ProjectDir)

	if err != nil {
		t.Fatalf("unable to get project package name: %v", err)
	}

	_, err = gen.generateModelPackage(storePkg, projectPackageName)

	assert.Nil(t, err)

	modelPens := mw.Writers[1:]

	modelPaths := []string{
		path.Join(tmpDir, "internal/models/actor.gen.go"),
		path.Join(tmpDir, "internal/models/movie.gen.go"),
	}

	for i, pen := range modelPens {
		testCaseName, _ := utils.SplitFilename(path.Base(pen.FullPath))

		t.Run(testCaseName, func(t *testing.T) {
			want := modelPaths[i]

			assert.Equal(t, want, pen.FullPath)

			cupaloy.SnapshotT(t, pen.Content)
		})
	}
}

func TestGenerate_generateQueryPackage(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()

	mw := writer.NewMemoryWriters()

	ft := types.NewFakeTranslate(
		`package {{ .PackageName }}
var content = `+"`"+`{
  "packageName": {{ .PackageName | ToJson }}, 
  "imports": {{ .Imports | ToJson }},
	"model": {{ .Model | ToJson }},
  "insertFields": {{ .InsertFields | ToJson }}
	"updateFields": {{ .UpdateFields | ToJson }}
	"selectFields": {{ .SelectFields | ToJson }}
}`+"`",
		`package {{ .PackageName }}
var content = `+"`"+`{
  "packageName": {{ .PackageName | ToJson }}, 
  "imports": {{ .Imports | ToJson }},
	"query": {{ .Query | ToJson }},
}`+"`",
	)

	gen, err := createGen(mw.Creator, ft, tmpDir)

	if err != nil {
		t.Fatalf("unable to create generate object: %v", err)
	}

	storePkg, err := gen.generateStorePackage("github.com/aakash-rajur/sqlxgen-example")

	assert.Nil(t, err)

	_, err = gen.generateQueryPackage(storePkg)

	assert.Nil(t, err)

	queryPens := mw.Writers[1:]

	queryPaths := []string{
		path.Join(tmpDir, "internal/api/list_actors.gen.go"),
		path.Join(tmpDir, "internal/api/get_actor.gen.go"),
		path.Join(tmpDir, "internal/api/list_movies.gen.go"),
		path.Join(tmpDir, "internal/api/get_movie.gen.go"),
	}

	for i, pen := range queryPens {
		testCaseName, _ := utils.SplitFilename(path.Base(pen.FullPath))

		t.Run(testCaseName, func(t *testing.T) {
			want := queryPaths[i]

			assert.Equal(t, want, pen.FullPath)

			cupaloy.SnapshotT(t, pen.Content)
		})
	}
}

func TestGenerate_Generate(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()

	mw := writer.NewMemoryWriters()

	ft := types.NewFakeTranslate(
		`package {{ .PackageName }}
var content = `+"`"+`{
  "packageName": {{ .PackageName | ToJson }}, 
  "imports": {{ .Imports | ToJson }},
	"model": {{ .Model | ToJson }},
  "insertFields": {{ .InsertFields | ToJson }}
	"updateFields": {{ .UpdateFields | ToJson }}
	"selectFields": {{ .SelectFields | ToJson }}
}`+"`",
		`package {{ .PackageName }}
var content = `+"`"+`{
  "packageName": {{ .PackageName | ToJson }}, 
  "imports": {{ .Imports | ToJson }},
	"query": {{ .Query | ToJson }},
  "insertFields": {{ .InsertFields | ToJson }}
	"updateFields": {{ .UpdateFields | ToJson }}
	"selectFields": {{ .SelectFields | ToJson }}
}`+"`",
	)

	gen, err := createGen(mw.Creator, ft, tmpDir)

	if err != nil {
		t.Fatalf("unable to create generate object: %v", err)
	}

	err = gen.Generate()

	assert.Nil(t, err)

	paths := []string{
		path.Join(tmpDir, "internal/store/store.gen.go"),
		path.Join(tmpDir, "internal/models/actor.gen.go"),
		path.Join(tmpDir, "internal/models/movie.gen.go"),
		path.Join(tmpDir, "internal/api/list_actors.gen.go"),
		path.Join(tmpDir, "internal/api/get_actor.gen.go"),
		path.Join(tmpDir, "internal/api/list_movies.gen.go"),
		path.Join(tmpDir, "internal/api/get_movie.gen.go"),
	}

	for i, pen := range mw.Writers {
		testCaseName, _ := utils.SplitFilename(path.Base(pen.FullPath))

		t.Run(testCaseName, func(t *testing.T) {
			want := paths[i]

			assert.Equal(t, want, pen.FullPath)

			cupaloy.SnapshotT(t, pen.Content)
		})
	}
}

func createGen(
	writerCreator writer.Creator,
	translate types.Translate,
	workDir string,
) (Generate, error) {
	goMod := `module github.com/aakash-rajur/sqlxgen-example

go 1.21.1

require (
	github.com/deckarep/golang-set v1.8.0
	github.com/go-sql-driver/mysql v1.7.1
	github.com/jinzhu/inflection v1.0.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/joho/godotenv v1.5.1
	github.com/joomcode/errorx v1.1.1
	github.com/lib/pq v1.10.9
	golang.org/x/mod v0.12.0
	gopkg.in/yaml.v3 v3.0.1
)

require github.com/bradleyjkemp/cupaloy v2.3.0+incompatible // indirect

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.8.4
)
`

	err := os.WriteFile(
		path.Join(workDir, "go.mod"),
		[]byte(goMod),
		0644,
	)

	if err != nil {
		return Generate{}, err
	}

	tables, err := utils.FromJson[introspect.Table](
		[]string{actorTableJson, movieTableJson},
	)

	if err != nil {
		return Generate{}, err
	}

	queries, err := utils.FromJson[introspect.Query](
		[]string{listActorQueryJson, getActorQueryJson, listMoviesQueryJson, getMovieQueryJson},
	)

	if err != nil {
		return Generate{}, err
	}

	gen := Generate{
		WriterCreator:   writerCreator,
		Translate:       translate,
		ProjectDir:      workDir,
		StorePackageDir: "internal/store",
		ModelPackageDir: "internal/models",
		Tables:          tables,
		Queries:         queries,
	}

	return gen, nil
}

//go:embed models/fixtures/actor-table.json
var actorTableJson string

//go:embed models/fixtures/movie-table.json
var movieTableJson string

//go:embed queries/fixtures/list-actor-query.json
var listActorQueryJson string

//go:embed queries/fixtures/get-actor-query.json
var getActorQueryJson string

//go:embed queries/fixtures/list-movies-query.json
var listMoviesQueryJson string

//go:embed queries/fixtures/get-movie-query.json
var getMovieQueryJson string
