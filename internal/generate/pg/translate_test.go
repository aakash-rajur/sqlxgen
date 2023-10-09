package go_type

import (
	"testing"

	"github.com/aakash-rajur/sqlxgen/internal/generate/types"
	"github.com/aakash-rajur/sqlxgen/internal/introspect"
	"github.com/stretchr/testify/assert"
)

func TestPg_ModelTemplate(t *testing.T) {
	translate := NewTranslate()

	assert.Equal(t, modelTemplate, translate.ModelTemplate())
}

func TestPg_QueryTemplate(t *testing.T) {
	translate := NewTranslate()

	assert.Equal(t, queryTemplate, translate.QueryTemplate())
}

func TestPg_Infer(t *testing.T) {
	t.Parallel()

	translate := NewTranslate()

	testCases := []struct {
		name   string
		column introspect.Column
		want   types.GoType
	}{
		{
			name: "serial4",
			column: introspect.Column{
				Type:    "serial4",
				IsArray: false,
			},
			want: types.GoType{
				DbType:    "serial4",
				GoType:    "*int32",
				IsPointer: true,
				Import:    "",
			},
		},
		{
			name: "serial8",
			column: introspect.Column{
				Type:    "serial8",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "serial8",
				GoType:    "*pq.Int64Array",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "serial2",
			column: introspect.Column{
				Type:    "serial2",
				IsArray: false,
			},
			want: types.GoType{
				DbType:    "serial2",
				GoType:    "*int16",
				IsPointer: true,
				Import:    "",
			},
		},
		{
			name: "int4",
			column: introspect.Column{
				Type:    "int4",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "int4",
				GoType:    "*pq.Int32Array",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "int8",
			column: introspect.Column{
				Type:    "int8",
				IsArray: false,
			},
			want: types.GoType{
				DbType:    "int8",
				GoType:    "*int64",
				IsPointer: true,
				Import:    "",
			},
		},
		{
			name: "int2",
			column: introspect.Column{
				Type:    "int2",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "int2",
				GoType:    "*pq.Int32Array",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "float4",
			column: introspect.Column{
				Type:    "float4",
				IsArray: false,
			},
			want: types.GoType{
				DbType:    "float4",
				GoType:    "*float32",
				IsPointer: true,
				Import:    "",
			},
		},
		{
			name: "float8",
			column: introspect.Column{
				Type:    "float8",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "float8",
				GoType:    "*pq.Float64Array",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "numeric",
			column: introspect.Column{
				Type:    "numeric",
				IsArray: false,
			},
			want: types.GoType{
				DbType:    "numeric",
				GoType:    "*string",
				IsPointer: true,
				Import:    "",
			},
		},
		{
			name: "money",
			column: introspect.Column{
				Type:    "money",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "money",
				GoType:    "*pq.StringArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "json object",
			column: introspect.Column{
				Type:     "json",
				JsonType: "object",
				IsArray:  false,
			},
			want: types.GoType{
				DbType:    "json",
				GoType:    "*store.JsonObject",
				Import:    "github.com/john-doe/gen/store",
				IsPointer: false,
			},
		},
		{
			name: "jsonb object",
			column: introspect.Column{
				Type:     "jsonb",
				JsonType: "object",
				IsArray:  true,
			},
			want: types.GoType{
				DbType:    "jsonb",
				GoType:    "json.RawMessage",
				IsPointer: false,
				Import:    "encoding/json",
			},
		},
		{
			name: "json array",
			column: introspect.Column{
				Type:     "json",
				JsonType: "array",
				IsArray:  false,
			},
			want: types.GoType{
				DbType:    "json",
				GoType:    "*store.JsonArray",
				Import:    "github.com/john-doe/gen/store",
				IsPointer: false,
			},
		},
		{
			name: "jsonb array",
			column: introspect.Column{
				Type:     "jsonb",
				JsonType: "array",
				IsArray:  true,
			},
			want: types.GoType{
				DbType:    "jsonb",
				GoType:    "json.RawMessage",
				IsPointer: false,
				Import:    "encoding/json",
			},
		},
		{
			name: "json identity",
			column: introspect.Column{
				Type:     "json",
				JsonType: "identity",
				IsArray:  false,
			},
			want: types.GoType{
				DbType:    "json",
				GoType:    "json.RawMessage",
				IsPointer: false,
				Import:    "encoding/json",
			},
		},
		{
			name: "jsonb identity",
			column: introspect.Column{
				Type:     "jsonb",
				JsonType: "identity",
				IsArray:  true,
			},
			want: types.GoType{
				DbType:    "jsonb",
				GoType:    "json.RawMessage",
				IsPointer: false,
				Import:    "encoding/json",
			},
		},
		{
			name: "text",
			column: introspect.Column{
				Type:    "text",
				IsArray: false,
			},
			want: types.GoType{
				DbType:    "text",
				GoType:    "*string",
				IsPointer: true,
				Import:    "",
			},
		},
		{
			name: "varchar",
			column: introspect.Column{
				Type:    "varchar",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "varchar",
				GoType:    "*pq.GenericArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "bpchar",
			column: introspect.Column{
				Type:    "bpchar",
				IsArray: false,
			},
			want: types.GoType{
				DbType:    "bpchar",
				GoType:    "*string",
				IsPointer: true,
				Import:    "",
			},
		},
		{
			name: "uuid",
			column: introspect.Column{
				Type:    "uuid",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "uuid",
				GoType:    "*pq.GenericArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "inet",
			column: introspect.Column{
				Type:    "inet",
				IsArray: false,
			},
			want: types.GoType{
				DbType:    "inet",
				GoType:    "*net.IP",
				IsPointer: true,
				Import:    "net",
			},
		},
		{
			name: "cidr",
			column: introspect.Column{
				Type:    "cidr",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "cidr",
				GoType:    "*pq.GenericArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "macaddr",
			column: introspect.Column{
				Type:    "macaddr",
				IsArray: false,
			},
			want: types.GoType{
				DbType:    "macaddr",
				GoType:    "*net.IP",
				IsPointer: true,
				Import:    "net",
			},
		},
		{
			name: "macaddr8",
			column: introspect.Column{
				Type:    "macaddr8",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "macaddr8",
				GoType:    "*pq.GenericArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "interval",
			column: introspect.Column{
				Type:    "interval",
				IsArray: false,
			},
			want: types.GoType{
				DbType:    "interval",
				GoType:    "*time.Duration",
				IsPointer: true,
				Import:    "time",
			},
		},
		{
			name: "bytea",
			column: introspect.Column{
				Type:    "interval",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "interval",
				GoType:    "*pq.GenericArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := translate.Infer(
				"github.com/john-doe/gen/store",
				"store",
				testCase.column,
			)

			assert.NoError(t, err)

			assert.Equal(t, testCase.want, got)
		})
	}
}
