package models

import (
	_ "embed"
	"testing"

	"github.com/aakash-rajur/sqlxgen/internal/generate/types"
	"github.com/aakash-rajur/sqlxgen/internal/introspect"
	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/aakash-rajur/sqlxgen/internal/utils/writer"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/assert"
)

func TestNewModel(t *testing.T) {
	t.Parallel()

	ft := types.NewFakeTranslate("", "")

	tables, err := utils.FromJson[introspect.Table](
		[]string{actorTableJson, movieTableJson},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	testCases := []struct {
		name  string
		table introspect.Table
		err   error
	}{
		{
			name:  "actor",
			table: tables[0],
		},
		{
			name:  "movie",
			table: tables[1],
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := newModel(nil, ft, testCase.table)

			if testCase.err != nil {
				errMsgLeft := testCase.err.Error()

				errMsgRight := err.Error()

				assert.Contains(t, errMsgRight, errMsgLeft)
			} else {
				assert.Nil(t, err)

				cupaloy.SnapshotT(t, got)
			}
		})
	}
}

func TestModel_GetImports(t *testing.T) {
	t.Parallel()

	tables, err := utils.FromJson[introspect.Table](
		[]string{actorTableJson, movieTableJson},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	ft := types.NewFakeTranslate("", "")

	testCases := []struct {
		name string
		t    introspect.Table
		want []string
	}{
		{
			name: "actor",
			t:    tables[0],
			want: []string{},
		},
		{
			name: "movie",
			t:    tables[1],
			want: []string{"time"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			m, err := newModel(nil, ft, testCase.t)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			got := m.getImports()

			assert.Equal(t, testCase.want, got)
		})
	}
}

func TestModel_Generate(t *testing.T) {
	t.Parallel()

	tables, err := utils.FromJson[introspect.Table](
		[]string{actorTableJson, movieTableJson},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

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
		"",
	)

	testCases := []struct {
		name     string
		t        introspect.Table
		filepath string
	}{
		{
			name:     "actor",
			t:        tables[0],
			filepath: "test/actor.gen.go",
		},
		{
			name:     "movie",
			t:        tables[1],
			filepath: "test/movie.gen.go",
		},
	}

	for i, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			m, err := newModel(mw.Creator, ft, testCase.t)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			err = m.generate(ft.ModelTemplate(), "test", "test")

			assert.Nil(t, err)

			got := mw.Writers[i]

			cupaloy.SnapshotT(t, got.Content)

			assert.Equal(t, testCase.filepath, got.FullPath)
		})
	}
}

func TestDistinguishFields(t *testing.T) {
	t.Parallel()

	ft := types.NewFakeTranslate("", "")

	tables, err := utils.FromJson[introspect.Table](
		[]string{actorTableJson, movieTableJson},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	testCases := []struct {
		name         string
		table        introspect.Table
		insertFields []types.Field
		updateFields []types.Field
		selectFields []types.Field
	}{
		{
			name:  "actor",
			table: tables[0],
			insertFields: []types.Field{
				{
					Name: "Name",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "name",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
			},
			updateFields: []types.Field{
				{
					Name: "Id",
					Type: types.GoType{
						DbType:    "int8",
						GoType:    "*int",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "id",
						Type:              "int8",
						TypeId:            "20",
						IsArray:           false,
						IsSequence:        true,
						Nullable:          false,
						Generated:         false,
						PkName:            "actors_pkey",
						PkOrdinalPosition: 1,
						JsonType:          "",
					},
				},
				{
					Name: "Name",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "name",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
			},
			selectFields: []types.Field{
				{
					Name: "Id",
					Type: types.GoType{
						DbType:    "int8",
						GoType:    "*int",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "id",
						Type:              "int8",
						TypeId:            "20",
						IsArray:           false,
						IsSequence:        true,
						Nullable:          false,
						Generated:         false,
						PkName:            "actors_pkey",
						PkOrdinalPosition: 1,
						JsonType:          "",
					},
				},
				{
					Name: "Name",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "name",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "NameSearch",
					Type: types.GoType{
						DbType:    "tsvector",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "name_search",
						Type:              "tsvector",
						TypeId:            "3614",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          true,
						Generated:         true,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
			},
		},
		{
			name:  "movie",
			table: tables[1],
			insertFields: []types.Field{
				{
					Name: "Title",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "title",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "OriginalTitle",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "original_title",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "OriginalLanguage",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "original_language",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Overview",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "overview",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Runtime",
					Type: types.GoType{
						DbType:    "int4",
						GoType:    "*int",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "runtime",
						Type:              "int4",
						TypeId:            "23",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "ReleaseDate",
					Type: types.GoType{
						DbType:    "date",
						GoType:    "*time.Time",
						Import:    "time",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "release_date",
						Type:              "date",
						TypeId:            "1082",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Tagline",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "tagline",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Status",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "status",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Homepage",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "homepage",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Popularity",
					Type: types.GoType{
						DbType:    "float8",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "popularity",
						Type:              "float8",
						TypeId:            "701",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "VoteAverage",
					Type: types.GoType{
						DbType:    "float8",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "vote_average",
						Type:              "float8",
						TypeId:            "701",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "VoteCount",
					Type: types.GoType{
						DbType:    "int4",
						GoType:    "*int",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "vote_count",
						Type:              "int4",
						TypeId:            "23",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Budget",
					Type: types.GoType{
						DbType:    "int8",
						GoType:    "*int",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "budget",
						Type:              "int8",
						TypeId:            "20",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Revenue",
					Type: types.GoType{
						DbType:    "int8",
						GoType:    "*int",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "revenue",
						Type:              "int8",
						TypeId:            "20",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Keywords",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "keywords",
						Type:              "text",
						TypeId:            "1009",
						IsArray:           true,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
			},
			updateFields: []types.Field{
				{
					Name: "Id",
					Type: types.GoType{
						DbType:    "int4",
						GoType:    "*int",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "id",
						Type:              "int4",
						TypeId:            "23",
						IsArray:           false,
						IsSequence:        true,
						Nullable:          false,
						Generated:         false,
						PkName:            "movies_pkey",
						PkOrdinalPosition: 1,
						JsonType:          "",
					},
				},
				{
					Name: "Title",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "title",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "OriginalTitle",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "original_title",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "OriginalLanguage",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "original_language",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Overview",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "overview",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Runtime",
					Type: types.GoType{
						DbType:    "int4",
						GoType:    "*int",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "runtime",
						Type:              "int4",
						TypeId:            "23",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "ReleaseDate",
					Type: types.GoType{
						DbType:    "date",
						GoType:    "*time.Time",
						Import:    "time",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "release_date",
						Type:              "date",
						TypeId:            "1082",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Tagline",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "tagline",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Status",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "status",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Homepage",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "homepage",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Popularity",
					Type: types.GoType{
						DbType:    "float8",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "popularity",
						Type:              "float8",
						TypeId:            "701",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "VoteAverage",
					Type: types.GoType{
						DbType:    "float8",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "vote_average",
						Type:              "float8",
						TypeId:            "701",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "VoteCount",
					Type: types.GoType{
						DbType:    "int4",
						GoType:    "*int",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "vote_count",
						Type:              "int4",
						TypeId:            "23",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Budget",
					Type: types.GoType{
						DbType:    "int8",
						GoType:    "*int",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "budget",
						Type:              "int8",
						TypeId:            "20",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Revenue",
					Type: types.GoType{
						DbType:    "int8",
						GoType:    "*int",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "revenue",
						Type:              "int8",
						TypeId:            "20",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Keywords",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "keywords",
						Type:              "text",
						TypeId:            "1009",
						IsArray:           true,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
			},
			selectFields: []types.Field{
				{
					Name: "Id",
					Type: types.GoType{
						DbType:    "int4",
						GoType:    "*int",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "id",
						Type:              "int4",
						TypeId:            "23",
						IsArray:           false,
						IsSequence:        true,
						Nullable:          false,
						Generated:         false,
						PkName:            "movies_pkey",
						PkOrdinalPosition: 1,
						JsonType:          "",
					},
				},
				{
					Name: "Title",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "title",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "OriginalTitle",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "original_title",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "OriginalLanguage",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "original_language",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Overview",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "overview",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Runtime",
					Type: types.GoType{
						DbType:    "int4",
						GoType:    "*int",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "runtime",
						Type:              "int4",
						TypeId:            "23",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "ReleaseDate",
					Type: types.GoType{
						DbType:    "date",
						GoType:    "*time.Time",
						Import:    "time",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "release_date",
						Type:              "date",
						TypeId:            "1082",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Tagline",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "tagline",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Status",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "status",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Homepage",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "homepage",
						Type:              "text",
						TypeId:            "25",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Popularity",
					Type: types.GoType{
						DbType:    "float8",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "popularity",
						Type:              "float8",
						TypeId:            "701",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "VoteAverage",
					Type: types.GoType{
						DbType:    "float8",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "vote_average",
						Type:              "float8",
						TypeId:            "701",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "VoteCount",
					Type: types.GoType{
						DbType:    "int4",
						GoType:    "*int",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "vote_count",
						Type:              "int4",
						TypeId:            "23",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Budget",
					Type: types.GoType{
						DbType:    "int8",
						GoType:    "*int",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "budget",
						Type:              "int8",
						TypeId:            "20",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Revenue",
					Type: types.GoType{
						DbType:    "int8",
						GoType:    "*int",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "revenue",
						Type:              "int8",
						TypeId:            "20",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "Keywords",
					Type: types.GoType{
						DbType:    "text",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "keywords",
						Type:              "text",
						TypeId:            "1009",
						IsArray:           true,
						IsSequence:        false,
						Nullable:          false,
						Generated:         false,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "TitleSearch",
					Type: types.GoType{
						DbType:    "tsvector",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "title_search",
						Type:              "tsvector",
						TypeId:            "3614",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          true,
						Generated:         true,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
				{
					Name: "KeywordsSearch",
					Type: types.GoType{
						DbType:    "tsvector",
						GoType:    "*string",
						Import:    "",
						IsPointer: true,
					},
					Column: introspect.Column{
						ColumnName:        "keywords_search",
						Type:              "tsvector",
						TypeId:            "3614",
						IsArray:           false,
						IsSequence:        false,
						Nullable:          true,
						Generated:         true,
						PkName:            "NONE",
						PkOrdinalPosition: -1,
						JsonType:          "",
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			m, err := newModel(nil, ft, testCase.table)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			insertFields, updateFields, selectFields := distinguishFields(m.Fields)

			assert.Nil(t, err)

			assert.Equal(t, testCase.insertFields, insertFields)

			assert.Equal(t, testCase.updateFields, updateFields)

			assert.Equal(t, testCase.selectFields, selectFields)
		})
	}
}

//go:embed fixtures/actor-table.json
var actorTableJson string

//go:embed fixtures/movie-table.json
var movieTableJson string
