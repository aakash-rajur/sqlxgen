package integration

import (
	"cmp"
	"os"
	"path"
	"slices"
	"testing"

	"github.com/aakash-rajur/sqlxgen/internal/config"
	"github.com/aakash-rajur/sqlxgen/internal/utils/fs"
	"github.com/aakash-rajur/sqlxgen/internal/utils/writer"
	"github.com/bradleyjkemp/cupaloy"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func setupTest(workDir string, t *testing.T) func(t *testing.T) {
	cleanup, err := withGoMod(workDir)

	if err != nil {
		t.Fatal(err)

		return func(_ *testing.T) {}
	}

	return func(t *testing.T) {
		err := cleanup()

		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestIntegration(t *testing.T) {
	t.Parallel()

	workingDir, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	cleanup := setupTest(workingDir, t)

	defer cleanup(t)

	fd := fs.NewFileDiscovery()

	mw := writer.NewMemoryWriters()

	connect := sqlx.Connect

	sqlxGenCfg, err := config.NewSqlxGen(
		config.SqlxGenArgs{
			Connect:       connect,
			Fd:            fd,
			WriterCreator: mw.Creator,
			WorkingDir:    workingDir,
			SqlxAltPath:   "",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	sqlxGenCfg.InitLogger()

	err = sqlxGenCfg.Generate()

	assert.Nil(t, err)

	slices.SortStableFunc(
		mw.Writers,
		func(a, b *writer.MemoryWriter) int {
			return cmp.Compare(a.FullPath, b.FullPath)
		},
	)

	testCases := []string{
		"fixtures/tmdb_mysql/get_actor.gen.go",
		"fixtures/tmdb_mysql/get_company.gen.go",
		"fixtures/tmdb_mysql/get_crew.gen.go",
		"fixtures/tmdb_mysql/get_movie.gen.go",
		"fixtures/tmdb_mysql/list_actors.gen.go",
		"fixtures/tmdb_mysql/list_companies.gen.go",
		"fixtures/tmdb_mysql/list_crew.gen.go",
		"fixtures/tmdb_mysql/list_hyper_parameters.gen.go",
		"fixtures/tmdb_mysql/list_movies.gen.go",
		"fixtures/tmdb_mysql/models/actor.gen.go",
		"fixtures/tmdb_mysql/models/company.gen.go",
		"fixtures/tmdb_mysql/models/crew.gen.go",
		"fixtures/tmdb_mysql/models/hyper_parameter.gen.go",
		"fixtures/tmdb_mysql/models/movie.gen.go",
		"fixtures/tmdb_mysql/models/movies_actor.gen.go",
		"fixtures/tmdb_mysql/models/movies_company.gen.go",
		"fixtures/tmdb_mysql/models/movies_country.gen.go",
		"fixtures/tmdb_mysql/models/movies_crew.gen.go",
		"fixtures/tmdb_mysql/models/movies_genre.gen.go",
		"fixtures/tmdb_mysql/models/movies_language.gen.go",
		"fixtures/tmdb_mysql/store/store.gen.go",
		"fixtures/tmdb_pg/get_actor.gen.go",
		"fixtures/tmdb_pg/get_company.gen.go",
		"fixtures/tmdb_pg/get_crew.gen.go",
		"fixtures/tmdb_pg/get_movie.gen.go",
		"fixtures/tmdb_pg/list_actors.gen.go",
		"fixtures/tmdb_pg/list_companies.gen.go",
		"fixtures/tmdb_pg/list_crew.gen.go",
		"fixtures/tmdb_pg/list_hyper_parameters.gen.go",
		"fixtures/tmdb_pg/list_movies.gen.go",
		"fixtures/tmdb_pg/models/actor.gen.go",
		"fixtures/tmdb_pg/models/company.gen.go",
		"fixtures/tmdb_pg/models/crew.gen.go",
		"fixtures/tmdb_pg/models/hyper_parameter.gen.go",
		"fixtures/tmdb_pg/models/movie.gen.go",
		"fixtures/tmdb_pg/models/movies_actor.gen.go",
		"fixtures/tmdb_pg/models/movies_company.gen.go",
		"fixtures/tmdb_pg/models/movies_country.gen.go",
		"fixtures/tmdb_pg/models/movies_crew.gen.go",
		"fixtures/tmdb_pg/models/movies_genre.gen.go",
		"fixtures/tmdb_pg/models/movies_language.gen.go",
		"fixtures/tmdb_pg/models/t_movie.gen.go",
		"fixtures/tmdb_pg/models/t_movies_credit.gen.go",
		"fixtures/tmdb_pg/store/store.gen.go",
	}

	assert.Len(t, mw.Writers, len(testCases), "number of generated files should match number of test cases")

	for i, testCase := range testCases {
		w := mw.Writers[i]

		testCaseName := testCase[:len(testCase)-3]

		t.Run(testCaseName, func(t *testing.T) {
			wantPath := testCase

			gotPath := w.FullPath

			assert.Contains(t, gotPath, wantPath)

			cupaloy.SnapshotT(t, w.Content)
		})
	}
}

func withGoMod(workDir string) (func() error, error) {
	modFile := path.Join(workDir, "go.mod")

	content := `module github.com/aakash-rajur/example

go 1.21.1
`

	err := os.WriteFile(modFile, []byte(content), 0644)

	if err != nil {
		return nil, err
	}

	cleanup := func() error {
		return os.Remove(modFile)
	}

	return cleanup, nil
}
