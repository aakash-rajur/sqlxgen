package go_type

import (
	"testing"

	"github.com/aakash-rajur/sqlxgen/internal/generate/types"
	"github.com/aakash-rajur/sqlxgen/internal/introspect"
	"github.com/stretchr/testify/assert"
)

func TestFromSingle(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		column introspect.Column
		want   types.GoType
	}{
		// generate for all pg types
		{
			name:   "int2",
			column: introspect.Column{Type: "int2"},
			want: types.GoType{
				DbType:    "int2",
				GoType:    "*int16",
				IsPointer: true,
			},
		},
		{
			name:   "int4",
			column: introspect.Column{Type: "int4"},
			want: types.GoType{
				DbType:    "int4",
				GoType:    "*int32",
				IsPointer: true,
			},
		},
		{
			name:   "int8",
			column: introspect.Column{Type: "int8"},
			want: types.GoType{
				DbType:    "int8",
				GoType:    "*int64",
				IsPointer: true,
			},
		},
		{
			name:   "float4",
			column: introspect.Column{Type: "float4"},
			want: types.GoType{
				DbType:    "float4",
				GoType:    "*float32",
				IsPointer: true,
			},
		},
		{
			name:   "float8",
			column: introspect.Column{Type: "float8"},
			want: types.GoType{
				DbType:    "float8",
				GoType:    "*float64",
				IsPointer: true,
			},
		},
		{
			name:   "serial2",
			column: introspect.Column{Type: "serial2"},
			want: types.GoType{
				DbType:    "serial2",
				GoType:    "*int16",
				IsPointer: true,
			},
		},
		{
			name:   "serial4",
			column: introspect.Column{Type: "serial4"},
			want: types.GoType{
				DbType:    "serial4",
				GoType:    "*int32",
				IsPointer: true,
			},
		},
		{
			name:   "serial8",
			column: introspect.Column{Type: "serial8"},
			want: types.GoType{
				DbType:    "serial8",
				GoType:    "*int64",
				IsPointer: true,
			},
		},
		{
			name:   "numeric",
			column: introspect.Column{Type: "numeric"},
			want: types.GoType{
				DbType:    "numeric",
				GoType:    "*string",
				IsPointer: true,
			},
		},
		{
			name:   "money",
			column: introspect.Column{Type: "money"},
			want: types.GoType{
				DbType:    "money",
				GoType:    "*string",
				IsPointer: true,
			},
		},
		{
			name:   "text",
			column: introspect.Column{Type: "text"},
			want: types.GoType{
				DbType:    "text",
				GoType:    "*string",
				IsPointer: true,
			},
		},
		{
			name:   "bytea",
			column: introspect.Column{Type: "bytea"},
			want: types.GoType{
				DbType:    "bytea",
				GoType:    "[]byte",
				IsPointer: true,
			},
		},
		{
			name:   "bool",
			column: introspect.Column{Type: "bool"},
			want: types.GoType{
				DbType:    "bool",
				GoType:    "*bool",
				IsPointer: true,
			},
		},
		{
			name:   "date",
			column: introspect.Column{Type: "date"},
			want: types.GoType{
				DbType:    "date",
				GoType:    "*time.Time",
				IsPointer: true,
				Import:    "time",
			},
		},
		{
			name:   "time",
			column: introspect.Column{Type: "time"},
			want: types.GoType{
				DbType:    "time",
				GoType:    "*time.Time",
				IsPointer: true,
				Import:    "time",
			},
		},
		{
			name:   "timetz",
			column: introspect.Column{Type: "timetz"},
			want: types.GoType{
				DbType:    "timetz",
				GoType:    "*time.Time",
				IsPointer: true,
				Import:    "time",
			},
		},
		{
			name:   "timestamp",
			column: introspect.Column{Type: "timestamp"},
			want: types.GoType{
				DbType:    "timestamp",
				GoType:    "*time.Time",
				IsPointer: true,
				Import:    "time",
			},
		},
		{
			name:   "timestamptz",
			column: introspect.Column{Type: "timestamptz"},
			want: types.GoType{
				DbType:    "timestamptz",
				GoType:    "*time.Time",
				IsPointer: true,
				Import:    "time",
			},
		},
		{
			name:   "interval",
			column: introspect.Column{Type: "interval"},
			want: types.GoType{
				DbType:    "interval",
				GoType:    "*time.Duration",
				IsPointer: true,
				Import:    "time",
			},
		},
		{
			name:   "inet",
			column: introspect.Column{Type: "inet"},
			want: types.GoType{
				DbType:    "inet",
				GoType:    "*net.IP",
				IsPointer: true,
				Import:    "net",
			},
		},
		{
			name:   "cidr",
			column: introspect.Column{Type: "cidr"},
			want: types.GoType{
				DbType:    "cidr",
				GoType:    "*net.IP",
				IsPointer: true,
				Import:    "net",
			},
		},
		{
			name:   "macaddr",
			column: introspect.Column{Type: "macaddr"},
			want: types.GoType{
				DbType:    "macaddr",
				GoType:    "*net.IP",
				IsPointer: true,
				Import:    "net",
			},
		},
		{
			name:   "macaddr8",
			column: introspect.Column{Type: "macaddr8"},
			want: types.GoType{
				DbType:    "macaddr8",
				GoType:    "*net.IP",
				IsPointer: true,
				Import:    "net",
			},
		},
		{
			name:   "uuid",
			column: introspect.Column{Type: "uuid"},
			want: types.GoType{
				DbType:    "uuid",
				GoType:    "*string",
				IsPointer: true,
			},
		},
		{
			name:   "void",
			column: introspect.Column{Type: "void"},
			want: types.GoType{
				DbType:    "void",
				GoType:    "interface{}",
				IsPointer: true,
			},
		},
		{
			name: "json object",
			column: introspect.Column{
				Type:     "json",
				JsonType: "object",
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
			},
			want: types.GoType{
				DbType:    "jsonb",
				GoType:    "*store.JsonObject",
				Import:    "github.com/john-doe/gen/store",
				IsPointer: false,
			},
		},
		{
			name: "json array",
			column: introspect.Column{
				Type:     "json",
				JsonType: "array",
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
			},
			want: types.GoType{
				DbType:    "jsonb",
				GoType:    "*store.JsonArray",
				Import:    "github.com/john-doe/gen/store",
				IsPointer: false,
			},
		},
		{
			name: "json identity",
			column: introspect.Column{
				Type:     "json",
				JsonType: "identity",
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
			},
			want: types.GoType{
				DbType:    "jsonb",
				GoType:    "json.RawMessage",
				IsPointer: false,
				Import:    "encoding/json",
			},
		},
		{
			name:   "xml",
			column: introspect.Column{Type: "xml"},
			want: types.GoType{
				DbType:    "xml",
				GoType:    "interface{}",
				IsPointer: true,
			},
		},
		{
			name:   "tsvector",
			column: introspect.Column{Type: "tsvector"},
			want: types.GoType{
				DbType:    "tsvector",
				GoType:    "*string",
				IsPointer: true,
			},
		},
		{
			name:   "tsquery",
			column: introspect.Column{Type: "tsquery"},
			want: types.GoType{
				DbType:    "tsquery",
				GoType:    "*string",
				IsPointer: true,
			},
		},
		{
			name:   "pg_catalog.int2",
			column: introspect.Column{Type: "pg_catalog.int2"},
			want: types.GoType{
				DbType:    "pg_catalog.int2",
				GoType:    "*int16",
				IsPointer: true,
			},
		},
		{
			name:   "pg_catalog.int4",
			column: introspect.Column{Type: "pg_catalog.int4"},
			want: types.GoType{
				DbType:    "pg_catalog.int4",
				GoType:    "*int32",
				IsPointer: true,
			},
		},
		{
			name:   "pg_catalog.int8",
			column: introspect.Column{Type: "pg_catalog.int8"},
			want: types.GoType{
				DbType:    "pg_catalog.int8",
				GoType:    "*int64",
				IsPointer: true,
			},
		},
		{
			name:   "pg_catalog.float4",
			column: introspect.Column{Type: "pg_catalog.float4"},
			want: types.GoType{
				DbType:    "pg_catalog.float4",
				GoType:    "*float32",
				IsPointer: true,
			},
		},
		{
			name:   "pg_catalog.float8",
			column: introspect.Column{Type: "pg_catalog.float8"},
			want: types.GoType{
				DbType:    "pg_catalog.float8",
				GoType:    "*float64",
				IsPointer: true,
			},
		},
		{
			name:   "pg_catalog.serial2",
			column: introspect.Column{Type: "pg_catalog.serial2"},
			want: types.GoType{
				DbType:    "pg_catalog.serial2",
				GoType:    "*int16",
				IsPointer: true,
			},
		},
		{
			name:   "pg_catalog.serial4",
			column: introspect.Column{Type: "pg_catalog.serial4"},
			want: types.GoType{
				DbType:    "pg_catalog.serial4",
				GoType:    "*int32",
				IsPointer: true,
			},
		},
		{
			name:   "pg_catalog.serial8",
			column: introspect.Column{Type: "pg_catalog.serial8"},
			want: types.GoType{
				DbType:    "pg_catalog.serial8",
				GoType:    "*int64",
				IsPointer: true,
			},
		},
		{
			name:   "pg_catalog.numeric",
			column: introspect.Column{Type: "pg_catalog.numeric"},
			want: types.GoType{
				DbType:    "pg_catalog.numeric",
				GoType:    "*string",
				IsPointer: true,
			},
		},
		{
			name:   "pg_catalog.money",
			column: introspect.Column{Type: "pg_catalog.money"},
			want: types.GoType{
				DbType:    "pg_catalog.money",
				GoType:    "*string",
				IsPointer: true,
			},
		},
		{
			name:   "pg_catalog.bpchar",
			column: introspect.Column{Type: "pg_catalog.bpchar"},
			want: types.GoType{
				DbType:    "pg_catalog.bpchar",
				GoType:    "*string",
				IsPointer: true,
			},
		},
		{
			name:   "pg_catalog.varchar",
			column: introspect.Column{Type: "pg_catalog.varchar"},
			want: types.GoType{
				DbType:    "pg_catalog.varchar",
				GoType:    "*string",
				IsPointer: true,
			},
		},
		{
			name:   "pg_catalog.text",
			column: introspect.Column{Type: "pg_catalog.text"},
			want: types.GoType{
				DbType:    "pg_catalog.text",
				GoType:    "interface{}",
				IsPointer: true,
			},
		},
		{
			name:   "pg_catalog.bytea",
			column: introspect.Column{Type: "pg_catalog.bytea"},
			want: types.GoType{
				DbType:    "pg_catalog.bytea",
				GoType:    "[]byte",
				IsPointer: true,
			},
		},
		{
			name:   "pg_catalog.bool",
			column: introspect.Column{Type: "pg_catalog.bool"},
			want: types.GoType{
				DbType:    "pg_catalog.bool",
				GoType:    "*bool",
				IsPointer: true,
			},
		},
		{
			name:   "pg_catalog.date",
			column: introspect.Column{Type: "pg_catalog.date"},
			want: types.GoType{
				DbType:    "pg_catalog.date",
				GoType:    "*time.Time",
				IsPointer: true,
				Import:    "time",
			},
		},
		{
			name:   "pg_catalog.time",
			column: introspect.Column{Type: "pg_catalog.time"},
			want: types.GoType{
				DbType:    "pg_catalog.time",
				GoType:    "*time.Time",
				IsPointer: true,
				Import:    "time",
			},
		},
		{
			name:   "pg_catalog.timetz",
			column: introspect.Column{Type: "pg_catalog.timetz"},
			want: types.GoType{
				DbType:    "pg_catalog.timetz",
				GoType:    "*time.Time",
				IsPointer: true,
				Import:    "time",
			},
		},
		{
			name:   "pg_catalog.timestamp",
			column: introspect.Column{Type: "pg_catalog.timestamp"},
			want: types.GoType{
				DbType:    "pg_catalog.timestamp",
				GoType:    "*time.Time",
				IsPointer: true,
				Import:    "time",
			},
		},
		{
			name:   "pg_catalog.timestamptz",
			column: introspect.Column{Type: "pg_catalog.timestamptz"},
			want: types.GoType{
				DbType:    "pg_catalog.timestamptz",
				GoType:    "*time.Time",
				IsPointer: true,
				Import:    "time",
			},
		},
		{
			name:   "pg_catalog.interval",
			column: introspect.Column{Type: "pg_catalog.interval"},
			want: types.GoType{
				DbType:    "pg_catalog.interval",
				GoType:    "*time.Duration",
				IsPointer: true,
				Import:    "time",
			},
		},
		{
			name:   "pg_catalog.inet",
			column: introspect.Column{Type: "pg_catalog.inet"},
			want: types.GoType{
				DbType:    "pg_catalog.inet",
				GoType:    "*net.IP",
				IsPointer: true,
				Import:    "net",
			},
		},
		{
			name:   "pg_catalog.cidr",
			column: introspect.Column{Type: "pg_catalog.cidr"},
			want: types.GoType{
				DbType:    "pg_catalog.cidr",
				GoType:    "*net.IP",
				IsPointer: true,
				Import:    "net",
			},
		},
		{
			name:   "pg_catalog.macaddr",
			column: introspect.Column{Type: "pg_catalog.macaddr"},
			want: types.GoType{
				DbType:    "pg_catalog.macaddr",
				GoType:    "*net.IP",
				IsPointer: true,
				Import:    "net",
			},
		},
		{
			name:   "pg_catalog.macaddr8",
			column: introspect.Column{Type: "pg_catalog.macaddr8"},
			want: types.GoType{
				DbType:    "pg_catalog.macaddr8",
				GoType:    "*net.IP",
				IsPointer: true,
				Import:    "net",
			},
		},
		{
			name:   "pg_catalog.uuid",
			column: introspect.Column{Type: "pg_catalog.uuid"},
			want: types.GoType{
				DbType:    "pg_catalog.uuid",
				GoType:    "*string",
				IsPointer: true,
			},
		},
		{
			name:   "pg_catalog.void",
			column: introspect.Column{Type: "pg_catalog.void"},
			want: types.GoType{
				DbType:    "pg_catalog.void",
				GoType:    "interface{}",
				IsPointer: true,
			},
		},
		{
			name: "pg_catalog.json object",
			column: introspect.Column{
				Type:     "pg_catalog.json",
				JsonType: "object",
			},
			want: types.GoType{
				DbType:    "pg_catalog.json",
				GoType:    "*store.JsonObject",
				Import:    "github.com/john-doe/gen/store",
				IsPointer: false,
			},
		},
		{
			name: "pg_catalog.jsonb object",
			column: introspect.Column{
				Type:     "pg_catalog.jsonb",
				JsonType: "object",
			},
			want: types.GoType{
				DbType:    "pg_catalog.jsonb",
				GoType:    "*store.JsonObject",
				Import:    "github.com/john-doe/gen/store",
				IsPointer: false,
			},
		},
		{
			name: "pg_catalog.json array",
			column: introspect.Column{
				Type:     "pg_catalog.json",
				JsonType: "array",
			},
			want: types.GoType{
				DbType:    "pg_catalog.json",
				GoType:    "*store.JsonArray",
				Import:    "github.com/john-doe/gen/store",
				IsPointer: false,
			},
		},
		{
			name: "pg_catalog.jsonb array",
			column: introspect.Column{
				Type:     "pg_catalog.jsonb",
				JsonType: "array",
			},
			want: types.GoType{
				DbType:    "pg_catalog.jsonb",
				GoType:    "*store.JsonArray",
				Import:    "github.com/john-doe/gen/store",
				IsPointer: false,
			},
		},
		{
			name: "pg_catalog.json identity",
			column: introspect.Column{
				Type:     "pg_catalog.json",
				JsonType: "identity",
			},
			want: types.GoType{
				DbType:    "pg_catalog.json",
				GoType:    "json.RawMessage",
				IsPointer: false,
				Import:    "encoding/json",
			},
		},
		{
			name: "pg_catalog.jsonb identity",
			column: introspect.Column{
				Type:     "pg_catalog.jsonb",
				JsonType: "identity",
			},
			want: types.GoType{
				DbType:    "pg_catalog.jsonb",
				GoType:    "json.RawMessage",
				IsPointer: false,
				Import:    "encoding/json",
			},
		},
		{
			name:   "pg_catalog.xml",
			column: introspect.Column{Type: "pg_catalog.xml"},
			want: types.GoType{
				DbType:    "pg_catalog.xml",
				GoType:    "interface{}",
				IsPointer: true,
			},
		},
		{
			name:   "pg_catalog.tsvector",
			column: introspect.Column{Type: "pg_catalog.tsvector"},
			want: types.GoType{
				DbType:    "pg_catalog.tsvector",
				GoType:    "*string",
				IsPointer: true,
			},
		},
		{
			name:   "pg_catalog.tsquery",
			column: introspect.Column{Type: "pg_catalog.tsquery"},
			want: types.GoType{
				DbType:    "pg_catalog.tsquery",
				GoType:    "*string",
				IsPointer: true,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := fromSingle("github.com/john-doe/gen/store", testCase.column)

			assert.Nil(t, err)

			assert.Equal(t, testCase.want, got)
		})
	}
}

func TestFromArray(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		column introspect.Column
		want   types.GoType
	}{
		{
			name: "pg_catalog.serial4",
			column: introspect.Column{
				Type:    "pg_catalog.serial4",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.serial4",
				GoType:    "*pg.Int32Array",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "pg_catalog.serial8",
			column: introspect.Column{
				Type:    "pg_catalog.serial8",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.serial8",
				GoType:    "*pq.Int64Array",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "pg_catalog.serial2",
			column: introspect.Column{
				Type:    "pg_catalog.serial2",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.serial2",
				GoType:    "*pq.Int32Array",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "pg_catalog.int4",
			column: introspect.Column{
				Type:    "pg_catalog.int4",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.int4",
				GoType:    "*pq.Int32Array",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "pg_catalog.int8",
			column: introspect.Column{
				Type:    "pg_catalog.int8",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.int8",
				GoType:    "*pq.Int64Array",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "pg_catalog.int2",
			column: introspect.Column{
				Type:    "pg_catalog.int2",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.int2",
				GoType:    "*pq.Int32Array",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "pg_catalog.float4",
			column: introspect.Column{
				Type:    "pg_catalog.float4",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.float4",
				GoType:    "*pq.Float32Array",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "pg_catalog.float8",
			column: introspect.Column{
				Type:    "pg_catalog.float8",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.float8",
				GoType:    "*pq.Float64Array",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "pg_catalog.numeric",
			column: introspect.Column{
				Type:    "pg_catalog.numeric",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.numeric",
				GoType:    "*pq.StringArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "pg_catalog.money",
			column: introspect.Column{
				Type:    "pg_catalog.money",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.money",
				GoType:    "*pq.StringArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "pg_catalog.json object",
			column: introspect.Column{
				Type:     "pg_catalog.json",
				JsonType: "object",
				IsArray:  true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.json",
				GoType:    "json.RawMessage",
				IsPointer: false,
				Import:    "encoding/json",
			},
		},
		{
			name: "pg_catalog.jsonb object",
			column: introspect.Column{
				Type:     "pg_catalog.jsonb",
				JsonType: "object",
				IsArray:  true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.jsonb",
				GoType:    "json.RawMessage",
				IsPointer: false,
				Import:    "encoding/json",
			},
		},
		{
			name: "pg_catalog.json array",
			column: introspect.Column{
				Type:     "pg_catalog.json",
				JsonType: "array",
				IsArray:  true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.json",
				GoType:    "json.RawMessage",
				IsPointer: false,
				Import:    "encoding/json",
			},
		},
		{
			name: "pg_catalog.jsonb array",
			column: introspect.Column{
				Type:     "pg_catalog.jsonb",
				JsonType: "array",
				IsArray:  true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.jsonb",
				GoType:    "json.RawMessage",
				IsPointer: false,
				Import:    "encoding/json",
			},
		},
		{
			name: "pg_catalog.json identity",
			column: introspect.Column{
				Type:     "pg_catalog.json",
				JsonType: "identity",
				IsArray:  true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.json",
				GoType:    "json.RawMessage",
				IsPointer: false,
				Import:    "encoding/json",
			},
		},
		{
			name: "pg_catalog.jsonb identity",
			column: introspect.Column{
				Type:     "pg_catalog.jsonb",
				JsonType: "identity",
				IsArray:  true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.jsonb",
				GoType:    "json.RawMessage",
				IsPointer: false,
				Import:    "encoding/json",
			},
		},
		{
			name: "pg_catalog.text",
			column: introspect.Column{
				Type:    "pg_catalog.text",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.text",
				GoType:    "*pq.StringArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "pg_catalog.varchar",
			column: introspect.Column{
				Type:    "pg_catalog.varchar",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.varchar",
				GoType:    "*pq.StringArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "pg_catalog.bpchar",
			column: introspect.Column{
				Type:    "pg_catalog.bpchar",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.bpchar",
				GoType:    "*pq.StringArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "pg_catalog.uuid",
			column: introspect.Column{
				Type:    "pg_catalog.uuid",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.uuid",
				GoType:    "*pq.GenericArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "pg_catalog.inet",
			column: introspect.Column{
				Type:    "pg_catalog.inet",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.inet",
				GoType:    "*pq.GenericArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "pg_catalog.cidr",
			column: introspect.Column{
				Type:    "pg_catalog.cidr",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.cidr",
				GoType:    "*pq.GenericArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "pg_catalog.macaddr",
			column: introspect.Column{
				Type:    "pg_catalog.macaddr",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.macaddr",
				GoType:    "*pq.GenericArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "pg_catalog.macaddr8",
			column: introspect.Column{
				Type:    "pg_catalog.macaddr8",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.macaddr8",
				GoType:    "*pq.GenericArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "pg_catalog.interval",
			column: introspect.Column{
				Type:    "pg_catalog.interval",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.interval",
				GoType:    "*pq.GenericArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},
		{
			name: "pg_catalog.bytea",
			column: introspect.Column{
				Type:    "pg_catalog.interval",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "pg_catalog.interval",
				GoType:    "*pq.GenericArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
			},
		},

		{
			name: "serial4",
			column: introspect.Column{
				Type:    "serial4",
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "serial4",
				GoType:    "*pg.Int32Array",
				IsPointer: true,
				Import:    "github.com/lib/pq",
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
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "serial2",
				GoType:    "*pq.Int32Array",
				IsPointer: true,
				Import:    "github.com/lib/pq",
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
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "int8",
				GoType:    "*pq.Int64Array",
				IsPointer: true,
				Import:    "github.com/lib/pq",
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
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "float4",
				GoType:    "*pq.Float32Array",
				IsPointer: true,
				Import:    "github.com/lib/pq",
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
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "numeric",
				GoType:    "*pq.StringArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
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
				IsArray:  true,
			},
			want: types.GoType{
				DbType:    "json",
				GoType:    "json.RawMessage",
				IsPointer: false,
				Import:    "encoding/json",
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
				IsArray:  true,
			},
			want: types.GoType{
				DbType:    "json",
				GoType:    "json.RawMessage",
				IsPointer: false,
				Import:    "encoding/json",
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
				IsArray:  true,
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
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "text",
				GoType:    "*pq.StringArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
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
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "bpchar",
				GoType:    "*pq.GenericArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
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
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "inet",
				GoType:    "*pq.GenericArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
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
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "macaddr",
				GoType:    "*pq.GenericArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
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
				IsArray: true,
			},
			want: types.GoType{
				DbType:    "interval",
				GoType:    "*pq.GenericArray",
				IsPointer: true,
				Import:    "github.com/lib/pq",
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
			got, err := fromArray(testCase.column)

			assert.Nil(t, err)

			assert.Equal(t, testCase.want, got)
		})
	}
}

func TestInfer(t *testing.T) {
	t.Parallel()

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
			got, err := infer("github.com/john-doe/gen/store", testCase.column)

			assert.Nil(t, err)

			assert.Equal(t, testCase.want, got)
		})
	}
}
