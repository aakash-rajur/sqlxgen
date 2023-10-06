package mysql

import (
	"testing"

	"github.com/aakash-rajur/sqlxgen/internal/generate/types"
	"github.com/aakash-rajur/sqlxgen/internal/introspect"
	"github.com/stretchr/testify/assert"
)

func TestMysql_ModelTemplate(t *testing.T) {
	translate := NewTranslate()

	assert.Equal(t, modelTemplate, translate.ModelTemplate())
}

func TestMysql_QueryTemplate(t *testing.T) {
	translate := NewTranslate()

	assert.Equal(t, queryTemplate, translate.QueryTemplate())
}

func TestMysql_Infer(t *testing.T) {
	t.Parallel()

	translate := NewTranslate()

	testCases := []struct {
		name   string
		column introspect.Column
		want   types.GoType
		err    error
	}{
		{
			name: "varchar",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "varchar",
			},
			want: types.GoType{
				DbType:    "varchar",
				GoType:    "*string",
				IsPointer: true,
			},
			err: nil,
		},
		{
			name: "text",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "text",
			},
			want: types.GoType{
				DbType:    "text",
				GoType:    "*string",
				IsPointer: true,
			},
			err: nil,
		},
		{
			name: "longtext",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "longtext",
			},
			want: types.GoType{
				DbType:    "longtext",
				GoType:    "*string",
				IsPointer: true,
			},
			err: nil,
		},
		{
			name: "mediumtext",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "mediumtext",
			},
			want: types.GoType{
				DbType:    "mediumtext",
				GoType:    "*string",
				IsPointer: true,
			},
			err: nil,
		},
		{
			name: "char",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "char",
			},
			want: types.GoType{
				DbType:    "char",
				GoType:    "*string",
				IsPointer: true,
			},
			err: nil,
		},
		{
			name: "tinyint",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "tinyint",
			},
			want: types.GoType{
				DbType:    "tinyint",
				GoType:    "*int16",
				IsPointer: true,
			},
			err: nil,
		},
		{
			name: "smallint",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "smallint",
			},
			want: types.GoType{
				DbType:    "smallint",
				GoType:    "*int16",
				IsPointer: true,
			},
			err: nil,
		},
		{
			name: "int",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "int",
			},
			want: types.GoType{
				DbType:    "int",
				GoType:    "*int32",
				IsPointer: true,
			},
			err: nil,
		},
		{
			name: "bigint",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "bigint",
			},
			want: types.GoType{
				DbType:    "bigint",
				GoType:    "*int64",
				IsPointer: true,
			},
			err: nil,
		},
		{
			name: "float",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "float",
			},
			want: types.GoType{
				DbType:    "float",
				GoType:    "*float32",
				IsPointer: true,
			},
			err: nil,
		},
		{
			name: "double",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "double",
			},
			want: types.GoType{
				DbType:    "double",
				GoType:    "*float64",
				IsPointer: true,
			},
			err: nil,
		},
		{
			name: "decimal",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "decimal",
			},
			want: types.GoType{
				DbType:    "decimal",
				GoType:    "*string",
				IsPointer: true,
			},
			err: nil,
		},
		{
			name: "time",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "time",
			},
			want: types.GoType{
				DbType:    "time",
				GoType:    "*time.Time",
				IsPointer: true,
				Import:    "time",
			},
			err: nil,
		},
		{
			name: "timestamp",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "timestamp",
			},
			want: types.GoType{
				DbType:    "timestamp",
				GoType:    "*time.Time",
				IsPointer: true,
				Import:    "time",
			},
			err: nil,
		},
		{
			name: "datetime",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "datetime",
			},
			want: types.GoType{
				DbType:    "datetime",
				GoType:    "*time.Time",
				IsPointer: true,
				Import:    "time",
			},
			err: nil,
		},
		{
			name: "json 1",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "json",
				JsonType:   "array",
			},
			want: types.GoType{
				DbType:    "json",
				GoType:    "[]map[string]interface{}",
				IsPointer: false,
			},
			err: nil,
		},
		{
			name: "json 2",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "json",
				JsonType:   "object",
			},
			want: types.GoType{
				DbType:    "json",
				GoType:    "map[string]interface{}",
				IsPointer: false,
			},
			err: nil,
		},
		{
			name: "json 3",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "json",
				JsonType:   "identity",
			},
			want: types.GoType{
				DbType:    "json",
				GoType:    "json.RawMessage",
				IsPointer: false,
				Import:    "encoding/json",
			},
			err: nil,
		},
		{
			name: "set",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "set",
			},
			want: types.GoType{
				DbType:    "set",
				GoType:    "*string",
				IsPointer: true,
			},
			err: nil,
		},
		{
			name: "binary",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "binary",
			},
			want: types.GoType{
				DbType:    "binary",
				GoType:    "*pq.ByteaArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
			err: nil,
		},
		{
			name: "varbinary",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "varbinary",
			},
			want: types.GoType{
				DbType:    "varbinary",
				GoType:    "*pq.ByteaArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
			err: nil,
		},
		{
			name: "blob",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "blob",
			},
			want: types.GoType{
				DbType:    "blob",
				GoType:    "*pq.ByteaArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
			err: nil,
		},
		{
			name: "longblob",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "longblob",
			},
			want: types.GoType{
				DbType:    "longblob",
				GoType:    "*pq.ByteaArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
			err: nil,
		},
		{
			name: "mediumblob",
			column: introspect.Column{
				ColumnName: "c1",
				Type:       "mediumblob",
			},
			want: types.GoType{
				DbType:    "mediumblob",
				GoType:    "*pq.ByteaArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
			err: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := translate.Infer(testCase.column)

			assert.Equal(t, testCase.err, err)

			assert.Equal(t, testCase.want, got)
		})
	}
}
