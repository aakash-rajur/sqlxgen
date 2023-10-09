package config

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/aakash-rajur/sqlxgen/internal/logger"
	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/aakash-rajur/sqlxgen/internal/utils/casing"
	"github.com/aakash-rajur/sqlxgen/internal/utils/fs"
	"github.com/aakash-rajur/sqlxgen/internal/utils/writer"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/jinzhu/inflection"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestNewSqlxConfig(t *testing.T) {
	sqlxgenFilename := "sqlxgen.yml"

	workDir := t.TempDir()

	err := withEnv(workDir, sqlxgenFilename, true)

	if err != nil {
		t.Fatal(err)
	}

	got, err := NewSqlxGen(
		SqlxGenArgs{
			WorkingDir:    workDir,
			SqlxAltPath:   "",
			Connect:       nil,
			Fd:            nil,
			WriterCreator: nil,
		},
	)

	assert.Nil(t, err)

	// hack because new path is created on every run
	*got.ProjectDir = "/playground"

	gotJson, err := json.MarshalIndent(got, "", "  ")

	if err != nil {
		t.Fatalf("failed to marshal sqlxgen: %v", err)
	}

	cupaloy.SnapshotT(t, gotJson)
}

func TestWithDefaults(t *testing.T) {
	sqlxgenFilename := "sqlxgen.yml"

	workDir := t.TempDir()

	err := withEnv(workDir, sqlxgenFilename, true)

	if err != nil {
		t.Fatal(err)
	}

	cfgContent, err := loadAndExpand(workDir, "")

	if err != nil {
		t.Fatal(err)
	}

	genCfg := &SqlxGen{ProjectDir: &workDir}

	err = yaml.Unmarshal([]byte(cfgContent), genCfg)

	if err != nil {
		t.Fatal(err)
	}

	got, err := withDefaults(
		SqlxGenArgs{
			WorkingDir:    workDir,
			SqlxAltPath:   "",
			Connect:       nil,
			Fd:            nil,
			WriterCreator: nil,
		},
		genCfg,
	)

	assert.Nil(t, err)

	*got.ProjectDir = "/playground"

	gotJson, err := json.MarshalIndent(got, "", "  ")

	if err != nil {
		t.Fatalf("failed to marshal sqlxgen: %v", err)
	}

	cupaloy.SnapshotT(t, gotJson)
}

func TestSqlxGen_InitLogger(t *testing.T) {
	t.Parallel()

	t.Run("text logger", func(t *testing.T) {
		content := &bytes.Buffer{}

		w := bufio.NewWriter(content)

		gen := &SqlxGen{
			Version: utils.PointerTo("1"),
			LogArgs: &logger.Args{
				LogLevel: "debug",
				Format:   "text",
				Writer:   w,
			},
		}

		gen.InitLogger()

		err := w.Flush()

		if err != nil {
			t.Error(err)
		}

		want := `level=INFO msg="logger initialized" version=1`

		assert.Contains(t, content.String(), want)
	})

	t.Run("json logger", func(t *testing.T) {
		content := &bytes.Buffer{}

		w := bufio.NewWriter(content)

		gen := &SqlxGen{
			Version: utils.PointerTo("1"),
			LogArgs: &logger.Args{
				LogLevel: "debug",
				Format:   "json",
				Writer:   w,
			},
		}

		gen.InitLogger()

		err := w.Flush()

		if err != nil {
			t.Error(err)
		}

		want := make(map[string]interface{})

		err = json.Unmarshal(content.Bytes(), &want)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "1", want["version"])

		assert.Equal(t, "INFO", want["level"])

		assert.Equal(t, "logger initialized", want["msg"])
	})
}

func TestSqlxGen_Generate(t *testing.T) {
	t.Parallel()

	pgModels := []string{
		"actors",
		"companies",
		"crew",
		"hyper_parameters",
		"movies",
		"movies_actors",
		"movies_companies",
		"movies_countries",
		"movies_crew",
		"movies_genres",
		"movies_languages",
	}

	pgQueryMetas := []queryMeta{
		{
			name:     "list actors",
			filename: "list-actors.sql",
			query:    withArtifact("../introspect/pg/fixtures/list-actors.sql"),
			result:   withArtifact("../introspect/pg/fixtures/list-actors.csv"),
		},
		{
			name:     "get actor",
			filename: "get-actor.sql",
			query:    withArtifact("../introspect/pg/fixtures/get-actor.sql"),
			result:   withArtifact("../introspect/pg/fixtures/get-actor.csv"),
		},
		{
			name:     "list movies",
			filename: "list-movies.sql",
			query:    withArtifact("../introspect/pg/fixtures/list-movies.sql"),
			result:   withArtifact("../introspect/pg/fixtures/list-movies.csv"),
		},
		{
			name:     "get movie",
			filename: "get-movie.sql",
			query:    withArtifact("../introspect/pg/fixtures/get-movie.sql"),
			result:   withArtifact("../introspect/pg/fixtures/get-movie.csv"),
		},
	}

	mysqlModels := []string{
		"actors",
		"companies",
		"crew",
		"hyper_parameters",
		"movies",
		"movies_actors",
		"movies_companies",
		"movies_countries",
		"movies_crew",
		"movies_genres",
		"movies_languages",
	}

	mysqlQueryMetas := []queryMeta{
		{
			name:     "list actors",
			filename: "list-actors.sql",
			query:    withArtifact("../introspect/mysql/fixtures/list-actors.sql"),
			result:   withArtifact("../introspect/mysql/fixtures/list-actors.csv"),
		},
		{
			name:     "get actor",
			filename: "get-actor.sql",
			query:    withArtifact("../introspect/mysql/fixtures/get-actor.sql"),
			result:   withArtifact("../introspect/mysql/fixtures/get-actor.csv"),
		},
		{
			name:     "list movies",
			filename: "list-movies.sql",
			query:    withArtifact("../introspect/mysql/fixtures/list-movies.sql"),
			result:   withArtifact("../introspect/mysql/fixtures/list-movies.csv"),
		},
		{
			name:     "get movie",
			filename: "get-movie.sql",
			query:    withArtifact("../introspect/pg/fixtures/get-movie.sql"),
			result:   withArtifact("../introspect/pg/fixtures/get-movie.csv"),
		},
	}

	connect := func(engine string, _ string) (*sqlx.DB, error) {
		if engine == "postgres" {
			db, mock, err := utils.NewMockSqlx()

			if err != nil {
				t.Fatalf("failed to create mock db: %v", err)
			}

			preparePgMock(pgQueryMetas, &mock)

			return db, nil
		}

		if engine == "mysql" {
			db, mock, err := utils.NewMockSqlx()

			if err != nil {
				t.Fatalf("failed to create mock db: %v", err)
			}

			prepareMysqlMock(mysqlQueryMetas, &mock)

			return db, nil
		}

		return nil, errors.New("invalid engine")
	}

	fd := ffdRouter{
		"gen/pg":    withFileDiscovery(pgQueryMetas),
		"gen/mysql": withFileDiscovery(mysqlQueryMetas),
	}

	mw := writer.NewMemoryWriters()

	wc := mw.Creator

	workDir := t.TempDir()

	err := withProjectContext(workDir)

	err = withSqlxContext(workDir)

	sqlxGen, err := NewSqlxGen(
		SqlxGenArgs{
			WorkingDir:    workDir,
			SqlxAltPath:   "",
			Connect:       connect,
			Fd:            fd,
			WriterCreator: wc,
		},
	)

	if err != nil {
		t.Fatalf("failed to create sqlxgen: %v", err)
	}

	err = sqlxGen.Generate()

	assert.Nil(t, err)

	pgStorePen := mw.Writers[0]

	pgModelPens := mw.Writers[1 : 1+len(pgModels)]

	pgQueryPens := mw.Writers[1+len(pgModels) : 1+len(pgModels)+len(pgQueryMetas)]

	t.Run("pg-store", func(t *testing.T) {
		pen := pgStorePen

		fullPath := path.Join(workDir, "gen/pg/store/store.gen.go")

		assert.Equal(t, fullPath, pen.FullPath)

		got := pen.Content

		cupaloy.SnapshotT(t, got)
	})

	for i, pen := range pgModelPens {
		modelName := inflection.Singular(pgModels[i])

		testName := fmt.Sprintf("pg-%s", modelName)

		t.Run(testName, func(t *testing.T) {
			modelSnake, err := casing.SnakeCase(modelName)

			if err != nil {
				t.Fatal(err)
			}

			genFileName := fmt.Sprintf("%s.gen.go", modelSnake)

			fullPath := path.Join(workDir, "gen/pg/models", genFileName)

			assert.Equal(t, fullPath, pen.FullPath)

			got := pen.Content

			cupaloy.SnapshotT(t, got)
		})
	}

	for i, pen := range pgQueryPens {
		qm := pgQueryMetas[i]

		testName := fmt.Sprintf("pg-%s", qm.name)

		t.Run(testName, func(t *testing.T) {
			filename, _ := utils.SplitFilename(qm.filename)

			filenameSnake, err := casing.SnakeCase(filename)

			if err != nil {
				t.Fatal(err)
			}

			genFileName := fmt.Sprintf("%s.gen.go", filenameSnake)

			fullPath := path.Join(workDir, "fixtures", genFileName)

			assert.Equal(t, fullPath, pen.FullPath)

			got := pen.Content

			cupaloy.SnapshotT(t, got)
		})
	}

	mysqlOffset := 1 + len(pgModels) + len(pgQueryMetas)

	mysqlStorePen := mw.Writers[mysqlOffset]

	mysqlModelPens := mw.Writers[mysqlOffset+1 : 1+mysqlOffset+len(mysqlModels)]

	mysqlQueryPens := mw.Writers[1+mysqlOffset+len(mysqlModels):]

	t.Run("mysql-store", func(t *testing.T) {
		pen := mysqlStorePen

		fullPath := path.Join(workDir, "gen/mysql/store/store.gen.go")

		assert.Equal(t, fullPath, pen.FullPath)

		got := pen.Content

		cupaloy.SnapshotT(t, got)
	})

	for i, pen := range mysqlModelPens {
		modelName := inflection.Singular(mysqlModels[i])

		testName := fmt.Sprintf("mysql-%s", modelName)

		t.Run(testName, func(t *testing.T) {
			modelSnake, err := casing.SnakeCase(modelName)

			if err != nil {
				t.Fatal(err)
			}

			genFileName := fmt.Sprintf("%s.gen.go", modelSnake)

			fullPath := path.Join(workDir, "gen/mysql/models", genFileName)

			assert.Equal(t, fullPath, pen.FullPath)

			got := pen.Content

			cupaloy.SnapshotT(t, got)
		})
	}

	for i, pen := range mysqlQueryPens {
		qm := mysqlQueryMetas[i]

		testName := fmt.Sprintf("mysql-%s", qm.name)

		t.Run(testName, func(t *testing.T) {
			filename, _ := utils.SplitFilename(qm.filename)

			filenameSnake, err := casing.SnakeCase(filename)

			if err != nil {
				t.Fatal(err)
			}

			genFileName := fmt.Sprintf("%s.gen.go", filenameSnake)

			fullPath := path.Join(workDir, "fixtures", genFileName)

			assert.Equal(t, fullPath, pen.FullPath)

			got := pen.Content

			cupaloy.SnapshotT(t, got)
		})
	}
}

func withSqlxContext(workDir string) error {
	//language=yaml
	content := `version: 1

log:
  level: info # debug, info, warn, error
  format: text # json, text

configs:
  - name: pg
    engine: postgres # postgres, mysql
    database:
      url: "postgres://app:app@localhost:5432/app"
      host: "localhost"
      port: "5432"
      user: "app"
      password: "app"
      db: "app"
      sslmode: "disable"
    source:
      models:
        schemas:
          - public
        # array of go regex pattern, empty means all, e.g. ["^.+$"]
        include: []
        # array of go regex pattern, empty means none e.g. ["^public\.migrations*"]
        exclude:
          - "^public.migrations$"
          - "^public.t_movies$"
          - "^public.t_movies_credits$"
      queries:
        paths:
          - gen/pg
        # array of go regex pattern, empty means all e.g. ["^[a-zA-Z0-9_]*.sql$"]
        include: []
        # array of go regex pattern, empty means none e.g. ["^migrations*.sql$"]
        exclude:
          - "^list-project-2.sql$"
    gen:
      store:
        path: gen/pg/store
      models:
        path: gen/pg/models
  - name: mysql
    engine: mysql # postgres, mysql
    database:
      url: "root:@localhost:3306/public"
      host: "localhost"
      port: "3306"
      user: "root"
      password: ""
      db: "public"
      sslmode: "disable"
    source:
      models:
        schemas:
          - public
        include: []
        exclude: []
      queries:
        paths:
          - gen/mysql
        include: []
        exclude: []
    gen:
      store:
        path: gen/mysql/store
      models:
        path: gen/mysql/models
`

	filepath := path.Join(workDir, "sqlxgen.yml")

	return os.WriteFile(filepath, []byte(content), 0644)
}

type ffdRouter map[string]fs.FileDiscovery

func (f ffdRouter) Find(dir string, pattern string, shallow bool) ([]fs.Discovery, error) {
	fd, ok := f[dir]

	if !ok {
		return nil, errors.New("not found")
	}

	return fd.Find(dir, pattern, shallow)
}
