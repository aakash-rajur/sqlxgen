package pg

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestMsgWithSchema(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		schema string
		msg    string
		want   string
	}{
		{
			name:   "empty schema",
			schema: "",
			msg:    "failed to parse query params",
			want:   ": failed to parse query params",
		},
		{
			name:   "valid schema 1",
			schema: "foo",
			msg:    "failed to parse query params",
			want:   "foo: failed to parse query params",
		},
		{
			name:   "valid schema 2",
			schema: "bar",
			msg:    "failed to parse query params",
			want:   "bar: failed to parse query params",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := msgWithFilename(testCase.schema, testCase.msg)

			assert.Equal(t, testCase.want, got)
		})
	}
}

func TestIntrospectSchema(t *testing.T) {
	t.Parallel()

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

	mock.ExpectBegin()

	mock.
		ExpectQuery("select (.+) from pg_catalog.pg_attribute attr (.+) where").
		WithArgs("public").
		WillReturnRows(
			sqlmock.NewRows([]string{"schema_name", "table_name", "columns"}).
				FromCSVString(introspectSchemaResultCsv),
		)

	mock.ExpectRollback()

	mock.ExpectClose()

	args := IntrospectArgs{
		Schemas:         []string{"public"},
		TableInclusions: []string{},
		TableExclusions: []string{
			"^public.t_movies$",
			"^public.t_movies_credits$",
		},
		QueryDirs:       []string{},
		QueryInclusions: []string{},
		QueryExclusions: []string{},
	}

	source := NewIntrospect(nil, args)

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

	tables, err := source.IntrospectSchema(tx)

	assert.Nil(t, err)

	for _, table := range tables {
		t.Run(table.TableName, func(t *testing.T) {
			tableJson, err := json.MarshalIndent(table, "", "  ")

			if err != nil {
				t.Fatalf("failed to marshal table: %v", err)
			}

			cupaloy.SnapshotT(t, tableJson)
		})
	}
}

//go:embed fixtures/tables.csv
var introspectSchemaResultCsv string
