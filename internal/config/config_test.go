package config

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/aakash-rajur/sqlxgen/internal/utils/array"
	"github.com/aakash-rajur/sqlxgen/internal/utils/casing"
	"github.com/aakash-rajur/sqlxgen/internal/utils/fs"
	"github.com/aakash-rajur/sqlxgen/internal/utils/writer"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/jinzhu/inflection"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestConfig_String(t *testing.T) {
	t.Parallel()

	t.Run("pg config", func(t *testing.T) {
		cfg, err := withPgConfig()

		if err != nil {
			t.Fatal(err)
		}

		want := `Config{Name: pnp, Engine: postgres, Database: Database{Host: localhost, Port: 5432, Db: app, User: app, Password: app, SslMode: disable, Url: postgres://app:app@localhost:5432/app}, Source: Source{models: Model{schemas: [public], include: [], exclude: [^public.migrations$ ^public.t_movies$ ^public.t_movies_credits$]}, queries: Query{paths: [gen/pg], include: [], exclude: [^list-project-2.sql$]}}, Gen: Gen{store: GenPartial{path: gen/pg/store}, model: GenPartial{path: gen/pg/models}}}`

		got := cfg.String()

		assert.Equal(t, want, got)
	})

	t.Run("mysql config", func(t *testing.T) {
		cfg, err := withMysqlConfig()

		if err != nil {
			t.Fatal(err)
		}

		want := `Config{Name: mysql, Engine: mysql, Database: Database{Host: localhost, Port: 3306, Db: public, User: root, Password: , SslMode: disable, Url: root:@localhost:3306/public}, Source: Source{models: Model{schemas: [public], include: [], exclude: []}, queries: Query{paths: [gen/mysql], include: [], exclude: []}}, Gen: Gen{store: GenPartial{path: gen/mysql/store}, model: GenPartial{path: gen/mysql/models}}}`

		got := cfg.String()

		assert.Equal(t, want, got)
	})
}

func TestConfig_Merge(t *testing.T) {
	t.Parallel()

	t.Run("pg config nil", func(t *testing.T) {
		defaultPgCfg := defaultPgConfig()

		got := defaultPgCfg.Merge(nil)

		gotJson, err := json.MarshalIndent(got, "", "  ")

		if err != nil {
			t.Fatal(err)
		}

		cupaloy.SnapshotT(t, gotJson)
	})

	t.Run("pg config", func(t *testing.T) {
		defaultPgCfg := defaultPgConfig()

		pgCfg, err := withPgConfig()

		if err != nil {
			t.Fatal(err)
		}

		got := defaultPgCfg.Merge(pgCfg)

		gotJson, err := json.MarshalIndent(got, "", "  ")

		if err != nil {
			t.Fatal(err)
		}

		cupaloy.SnapshotT(t, gotJson)
	})

	t.Run("mysql config nil", func(t *testing.T) {
		defaultPgCfg := defaultPgConfig()

		mysqlCfg, err := withMysqlConfig()

		if err != nil {
			t.Fatal(err)
		}

		got := defaultPgCfg.Merge(mysqlCfg)

		gotJson, err := json.MarshalIndent(got, "", "  ")

		if err != nil {
			t.Fatal(err)
		}

		cupaloy.SnapshotT(t, gotJson)
	})

	t.Run("mysql config", func(t *testing.T) {
		defaultPgCfg := defaultPgConfig()

		mysqlCfg, err := withMysqlConfig()

		if err != nil {
			t.Fatal(err)
		}

		got := defaultPgCfg.Merge(mysqlCfg)

		gotJson, err := json.MarshalIndent(got, "", "  ")

		if err != nil {
			t.Fatal(err)
		}

		cupaloy.SnapshotT(t, gotJson)
	})
}

func TestConfig_GeneratePg(t *testing.T) {
	t.Parallel()

	models := []string{
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

	queryMetas := []queryMeta{
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

	db, mock, err := utils.NewMockSqlx()

	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}

	preparePgMock(queryMetas, &mock)

	defer func(db *sqlx.DB) {
		err := db.Close()

		if err != nil {
			t.Fatalf("failed to close mock db: %v", err)
		}
	}(db)

	tx, err := db.Beginx()

	if err != nil {
		t.Fatalf("failed to create transaction: %v", err)
	}

	defer func(tx *sqlx.Tx) {
		err := tx.Rollback()

		if err != nil {
			t.Fatalf("failed to rollback transaction: %v", err)
		}
	}(tx)

	fd := withFileDiscovery(queryMetas)

	mw := writer.NewMemoryWriters()

	wc := mw.Creator

	cfg, err := withPgConfig()

	if err != nil {
		t.Fatalf("failed to create config: %v", err)
	}

	workDir := t.TempDir()

	err = withProjectContext(workDir)

	err = cfg.generatePg(fd, wc, tx, workDir)

	assert.NoError(t, err)

	storePen, modelPens, queryPens := mw.Writers[0], mw.Writers[1:1+len(models)], mw.Writers[1+len(models):]

	t.Run("store", func(t *testing.T) {
		pen := storePen

		fullPath := path.Join(workDir, "gen/pg/store/store.gen.go")

		assert.Equal(t, fullPath, pen.FullPath)

		got := pen.Content

		cupaloy.SnapshotT(t, got)
	})

	for i, pen := range modelPens {
		modelName := inflection.Singular(models[i])

		t.Run(modelName, func(t *testing.T) {
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

	for i, pen := range queryPens {
		qm := queryMetas[i]

		t.Run(qm.name, func(t *testing.T) {
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

func TestConfig_GenerateMysql(t *testing.T) {
	t.Parallel()

	models := []string{
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

	queryMetas := []queryMeta{
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

	db, mock, err := utils.NewMockSqlx()

	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()

		if err != nil {
			t.Fatalf("failed to close mock db: %v", err)
		}
	}(db)

	prepareMysqlMock(queryMetas, &mock)

	tx, err := db.Beginx()

	if err != nil {
		t.Fatalf("failed to create transaction: %v", err)
	}

	defer func(tx *sqlx.Tx) {
		err := tx.Rollback()

		if err != nil {
			t.Fatalf("failed to rollback transaction: %v", err)
		}
	}(tx)

	fd := withFileDiscovery(queryMetas)

	mw := writer.NewMemoryWriters()

	wc := mw.Creator

	cfg, err := withMysqlConfig()

	if err != nil {
		t.Fatalf("failed to create config: %v", err)
	}

	workDir := t.TempDir()

	err = withProjectContext(workDir)

	err = cfg.generateMysql(fd, wc, tx, workDir)

	assert.NoError(t, err)

	storePen, modelPens, queryPens := mw.Writers[0], mw.Writers[1:1+len(models)], mw.Writers[1+len(models):]

	t.Run("store", func(t *testing.T) {
		pen := storePen

		fullPath := path.Join(workDir, "gen/mysql/store/store.gen.go")

		assert.Equal(t, fullPath, pen.FullPath)

		got := pen.Content

		cupaloy.SnapshotT(t, got)
	})

	for i, pen := range modelPens {
		modelName := inflection.Singular(models[i])

		t.Run(modelName, func(t *testing.T) {
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

	for i, pen := range queryPens {
		qm := queryMetas[i]

		t.Run(qm.name, func(t *testing.T) {
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

func TestConfig_Generate(t *testing.T) {
	t.Parallel()

	t.Run("pg", func(t *testing.T) {
		t.Parallel()

		models := []string{
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

		queryMetas := []queryMeta{
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

		db, mock, err := utils.NewMockSqlx()

		if err != nil {
			t.Fatalf("failed to create mock db: %v", err)
		}

		preparePgMock(queryMetas, &mock)

		connect := func(_ string, _ string) (*sqlx.DB, error) {
			return db, nil
		}

		fd := withFileDiscovery(queryMetas)

		mw := writer.NewMemoryWriters()

		wc := mw.Creator

		cfg, err := withPgConfig()

		if err != nil {
			t.Fatalf("failed to create config: %v", err)
		}

		workDir := t.TempDir()

		err = withProjectContext(workDir)

		err = cfg.Generate(connect, fd, wc, workDir)

		assert.NoError(t, err)

		storePen, modelPens, queryPens := mw.Writers[0], mw.Writers[1:1+len(models)], mw.Writers[1+len(models):]

		t.Run("store", func(t *testing.T) {
			pen := storePen

			fullPath := path.Join(workDir, "gen/pg/store/store.gen.go")

			assert.Equal(t, fullPath, pen.FullPath)

			got := pen.Content

			cupaloy.SnapshotT(t, got)
		})

		for i, pen := range modelPens {
			modelName := inflection.Singular(models[i])

			t.Run(modelName, func(t *testing.T) {
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

		for i, pen := range queryPens {
			qm := queryMetas[i]

			t.Run(qm.name, func(t *testing.T) {
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
	})

	t.Run("mysql", func(t *testing.T) {
		t.Parallel()

		models := []string{
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

		queryMetas := []queryMeta{
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

		db, mock, err := utils.NewMockSqlx()

		if err != nil {
			t.Fatalf("failed to create mock db: %v", err)
		}

		prepareMysqlMock(queryMetas, &mock)

		connect := func(_ string, _ string) (*sqlx.DB, error) {
			return db, nil
		}

		fd := withFileDiscovery(queryMetas)

		mw := writer.NewMemoryWriters()

		wc := mw.Creator

		cfg, err := withMysqlConfig()

		if err != nil {
			t.Fatalf("failed to create config: %v", err)
		}

		workDir := t.TempDir()

		err = withProjectContext(workDir)

		err = cfg.Generate(connect, fd, wc, workDir)

		assert.NoError(t, err)

		assert.NoError(t, err)

		storePen, modelPens, queryPens := mw.Writers[0], mw.Writers[1:1+len(models)], mw.Writers[1+len(models):]

		t.Run("store", func(t *testing.T) {
			pen := storePen

			fullPath := path.Join(workDir, "gen/mysql/store/store.gen.go")

			assert.Equal(t, fullPath, pen.FullPath)

			got := pen.Content

			cupaloy.SnapshotT(t, got)
		})

		for i, pen := range modelPens {
			modelName := inflection.Singular(models[i])

			t.Run(modelName, func(t *testing.T) {
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

		for i, pen := range queryPens {
			qm := queryMetas[i]

			t.Run(qm.name, func(t *testing.T) {
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
	})
}

func withPgConfig() (*Config, error) {
	// language=yaml
	pgCfgContent := `
name: pnp
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
`

	cfg := &Config{}

	err := yaml.Unmarshal([]byte(pgCfgContent), cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func withMysqlConfig() (*Config, error) {
	// language=yaml
	mysqlCfgContent := `
name: mysql
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

	cfg := &Config{}

	err := yaml.Unmarshal([]byte(mysqlCfgContent), cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func preparePgMock(
	qms []queryMeta,
	mock *sqlmock.Sqlmock,
) {
	m := *mock

	m.ExpectBegin()

	m.
		ExpectQuery("select (.+) from pg_catalog.pg_attribute attr (.+) where").
		WithArgs("public").
		WillReturnRows(
			sqlmock.NewRows([]string{"schema_name", "table_name", "columns"}).
				FromCSVString(
					withArtifact("../introspect/pg/fixtures/tables.csv"),
				),
		)

	for _, qm := range qms {
		m.ExpectExec("drop table if exists sample_query_introspection").
			WillReturnResult(
				sqlmock.NewResult(0, 0),
			)

		m.ExpectExec("create temp table if not exists sample_query_introspection (.+)").
			WillReturnResult(
				sqlmock.NewResult(0, 0),
			)

		m.ExpectQuery("select (.+) from pg_attribute attr (.+)").
			WillReturnRows(
				sqlmock.NewRows([]string{"column_name", "type", "type_id", "is_array", "is_sequence", "nullable", "generated"}).
					FromCSVString(qm.result),
			)
	}

	m.ExpectRollback()

	m.ExpectClose()
}

func prepareMysqlMock(
	qms []queryMeta,
	mock *sqlmock.Sqlmock,
) {
	m := *mock

	m.ExpectBegin()

	m.
		ExpectQuery("select (.+) from information_schema.columns c (.+) where").
		WithArgs("public").
		WillReturnRows(
			sqlmock.NewRows([]string{"schema_name", "table_name", "columns"}).
				FromCSVString(
					withArtifact("../introspect/mysql/fixtures/tables.csv"),
				),
		)

	for _, qm := range qms {
		m.ExpectExec("create table if not exists sample_query_introspection (.+)").
			WillReturnResult(
				sqlmock.NewResult(0, 0),
			)

		m.ExpectQuery("select (.+) from information_schema.columns c (.+)").
			WillReturnRows(
				sqlmock.NewRows([]string{"column_name", "type", "type_id", "is_array", "is_sequence", "nullable", "generated"}).
					FromCSVString(qm.result),
			)

		m.ExpectExec("drop table if exists sample_query_introspection").
			WillReturnResult(
				sqlmock.NewResult(0, 0),
			)
	}

	m.ExpectRollback()

	m.ExpectClose()
}

func withFileDiscovery(qms []queryMeta) fs.FileDiscovery {
	fd := fs.NewFakeFileDiscovery(
		array.Map(
			qms,
			func(qm queryMeta, i int) fs.FakeDiscover {
				return fs.FakeDiscover{
					Content:  qm.query,
					Dir:      "fixtures",
					Filename: qm.filename,
					FullPath: "fixtures/" + qm.filename,
				}
			},
		),
	)

	return fd
}

func withProjectContext(workDir string) error {
	filepath := path.Join(workDir, "go.mod")

	return os.WriteFile(filepath, []byte(fakeGoMod), 0644)
}

func withArtifact(relPath string) string {
	content, _ := os.ReadFile(relPath)

	return string(content)
}

type queryMeta struct {
	name     string
	query    string
	result   string
	filename string
}

var fakeGoMod string = `module github.com/aakash-rajur/sqlxgen

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

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0 // indirect
	github.com/bradleyjkemp/cupaloy v2.3.0+incompatible // indirect
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.8.4
)
`
