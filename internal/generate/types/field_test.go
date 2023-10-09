package types

import (
	_ "embed"
	"encoding/json"
	"strings"
	"testing"

	"github.com/aakash-rajur/sqlxgen/internal/introspect"
	"github.com/stretchr/testify/assert"
)

func TestNewField(t *testing.T) {
	t.Parallel()

	actorsColumns, err := withActors()

	if err != nil {
		t.Error(err)

		return
	}

	movieColumns, err := withMovies()

	if err != nil {
		t.Error(err)

		return
	}

	columns := append(actorsColumns, movieColumns...)

	fields, err := withFields()

	if err != nil {
		t.Error(err)

		return
	}

	ft := NewFakeTranslate("", "")

	for i, column := range columns {
		testCaseName := strings.Join([]string{column.ColumnName, column.Type}, "-")

		t.Run(testCaseName, func(t *testing.T) {
			t.Parallel()

			want := fields[i]

			got, err := NewField(column, ft, "gen/store", "github.com/john-doe/gen/store")

			assert.Nil(t, err)

			assert.Equal(t, want, got)
		})
	}
}

func withActors() ([]introspect.Column, error) {
	columns := make([]introspect.Column, 0)

	err := json.Unmarshal([]byte(actorsColumnsJson), &columns)

	if err != nil {
		return nil, err
	}

	return columns, nil
}

//go:embed fixtures/actors.json
var actorsColumnsJson string

func withMovies() ([]introspect.Column, error) {
	columns := make([]introspect.Column, 0)

	err := json.Unmarshal([]byte(moviesColumnsJson), &columns)

	if err != nil {
		return nil, err
	}

	return columns, nil
}

//go:embed fixtures/movies.json
var moviesColumnsJson string

func withFields() ([]Field, error) {
	fields := make([]Field, 0)

	err := json.Unmarshal([]byte(fieldsJson), &fields)

	if err != nil {
		return nil, err
	}

	return fields, nil
}

//go:embed fixtures/fields.json
var fieldsJson string
