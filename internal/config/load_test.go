package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestLoadAndExpand(t *testing.T) {
	t.Run("without alt path", func(t *testing.T) {
		sqlxgenFilename := "sqlxgen.yml"

		workDir := t.TempDir()

		err := withEnv(workDir, sqlxgenFilename, true)

		if err != nil {
			t.Fatal(err)
		}

		got, err := loadAndExpand(workDir, "")

		assert.Nil(t, err)

		cupaloy.SnapshotT(t, got)
	})

	t.Run("with alt path", func(t *testing.T) {
		workDir := t.TempDir()

		sqlxgenFilename := "configs/sqlxgen.yml"

		err := withEnv(workDir, sqlxgenFilename, true)

		if err != nil {
			t.Fatal(err)
		}

		absolutAltPath := path.Join(workDir, sqlxgenFilename)

		got, err := loadAndExpand(workDir, absolutAltPath)

		assert.Nil(t, err)

		cupaloy.SnapshotT(t, got)
	})
}

func TestGetSqlxGenPath(t *testing.T) {
	t.Parallel()

	t.Run("from env", func(t *testing.T) {
		sqlxgenFilename := "sqlxgen.yml"

		workDir := t.TempDir()

		err := withEnv(workDir, sqlxgenFilename, true)

		if err != nil {
			t.Fatal(err)
		}

		envFile := path.Join(workDir, ".env")

		env, err := godotenv.Read(envFile)

		if err != nil {
			t.Fatal(err)
		}

		cfgPath := getSqlxGenPath(workDir, env, "")

		want := path.Join(workDir, sqlxgenFilename)

		assert.Equal(t, want, cfgPath)
	})

	t.Run("without env", func(t *testing.T) {
		sqlxgenFilename := "sqlxgen.yml"

		workDir := t.TempDir()

		err := withEnv(workDir, sqlxgenFilename, false)

		if err != nil {
			t.Fatal(err)
		}

		envFile := path.Join(workDir, ".env")

		env, _ := godotenv.Read(envFile)

		cfgPath := getSqlxGenPath(workDir, env, "")

		want := path.Join(workDir, sqlxgenFilename)

		assert.Equal(t, want, cfgPath)
	})

	t.Run("with alt path", func(t *testing.T) {
		sqlxgenFilename := "configs/sqlxgen.yml"

		workDir := t.TempDir()

		err := withEnv(workDir, sqlxgenFilename, true)

		if err != nil {
			t.Fatal(err)
		}

		envFile := path.Join(workDir, ".env")

		env, err := godotenv.Read(envFile)

		if err != nil {
			t.Fatal(err)
		}

		altPath := path.Join(workDir, sqlxgenFilename)

		sqlxGenPath := getSqlxGenPath(workDir, env, altPath)

		assert.Equal(t, altPath, sqlxGenPath)
	})
}

func withEnv(
	workDir string,
	sqlxgenFilename string,
	writeEnv bool,
) error {
	// language=yaml
	sqlxContent := `### example sqlxgen.yml file
version: 1

log:
  level: info # debug, info, warn, error
  format: text # json, text

configs:
  - name: pnp
    engine: postgres # postgres, mysql
    database:
      url: "${POSTGRESQL_URL}"
      host: "${POSTGRESQL_HOST}"
      port: "${POSTGRESQL_PORT}"
      user: "${POSTGRESQL_USER}"
      password: "${POSTGRESQL_PASSWORD}"
      db: "${POSTGRESQL_DATABASE}"
      sslmode: "${POSTGRESQL_SSLMODE}"
    source:
      models:
        schemas:
          - public
        # array of go regex pattern, empty means all, e.g. ["^.+$"]
        include: []
        # array of go regex pattern, empty means none e.g. ["^public\.migrations*"]
        exclude:
          - "^public.migrations$"
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
      url: "${MYSQL_URL}"
      host: "${MYSQL_HOST}"
      port: "${MYSQL_PORT}"
      user: "${MYSQL_USER}"
      password: "${MYSQL_PASSWORD}"
      db: "${MYSQL_DATABASE}"
      sslmode: "${MYSQL_SSLMODE}"
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

	sqlxgenPath := path.Join(workDir, sqlxgenFilename)

	sqlxGenDir := filepath.Dir(sqlxgenPath)

	err := os.MkdirAll(sqlxGenDir, 0755)

	if err != nil {
		return err
	}

	err = os.WriteFile(sqlxgenPath, []byte(sqlxContent), 0644)

	if err != nil {
		return err
	}

	if !writeEnv {
		return nil
	}

	// language=dotenv
	envContent := `### example .env file
SQLXGEN_CONFIG_PATH=%s

POSTGRESQL_URL=postgres://app:app@localhost:5432/app
POSTGRESQL_HOST=localhost
POSTGRESQL_PORT=5432
POSTGRESQL_DATABASE=app
POSTGRESQL_USER=app
POSTGRESQL_PASSWORD=app
POSTGRESQL_SSLMODE=disable

MYSQL_URL=root:@localhost:3306/public
MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_DATABASE=public
MYSQL_USER=root
MYSQL_PASSWORD=""
MYSQL_SSLMODE=disable
`

	envFile := fmt.Sprintf(envContent, sqlxgenPath)

	err = os.WriteFile(path.Join(workDir, ".env"), []byte(envFile), 0644)

	if err != nil {
		return err
	}

	return nil
}
