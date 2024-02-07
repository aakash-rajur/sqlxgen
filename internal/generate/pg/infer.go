package go_type

import (
	"fmt"
	"path/filepath"

	"github.com/aakash-rajur/sqlxgen/internal/generate/types"
	"github.com/aakash-rajur/sqlxgen/internal/introspect"
)

func infer(
	storePackageDir string,
	column introspect.Column,
) (types.GoType, error) {
	if column.IsArray {
		return fromArray(column)
	}

	return fromSingle(storePackageDir, column)
}

// ref: https://github.com/sqlc-dev/sqlc/blob/main/internal/codegen/golang/postgresql_type.go#L36
func fromSingle(storePackageDir string, column introspect.Column) (types.GoType, error) {
	goType := types.GoType{
		DbType:    column.Type,
		GoType:    "interface{}",
		IsPointer: true,
	}

	switch column.Type {
	case "serial", "serial4", "pg_catalog.serial4":
		goType.GoType = "*int32"

		return goType, nil

	case "bigserial", "serial8", "pg_catalog.serial8":
		goType.GoType = "*int64"

		return goType, nil

	case "smallserial", "serial2", "pg_catalog.serial2":
		goType.GoType = "*int16"

		return goType, nil

	case "integer", "int", "int4", "pg_catalog.int4":
		goType.GoType = "*int32"

		return goType, nil

	case "bigint", "int8", "pg_catalog.int8":
		goType.GoType = "*int64"

		return goType, nil

	case "smallint", "int2", "pg_catalog.int2":
		goType.GoType = "*int16"

		return goType, nil

	case "float", "double precision", "float8", "pg_catalog.float8":
		goType.GoType = "*float64"

		return goType, nil

	case "real", "float4", "pg_catalog.float4":
		goType.GoType = "*float32"

		return goType, nil

	case "numeric", "pg_catalog.numeric", "money", "pg_catalog.money":
		// Since the Go standard library does not have a decimal type, lib/pq
		// returns numerics as strings.
		//
		// https://github.com/lib/pq/issues/648
		goType.GoType = "*string"

		return goType, nil

	case "boolean", "bool", "pg_catalog.bool":
		goType.GoType = "*bool"

		return goType, nil

	case "json", "jsonb", "pg_catalog.json", "pg_catalog.jsonb":
		goType.IsPointer = false

		goType.GoType = "json.RawMessage"

		goType.Import = "encoding/json"

		_, storePkg := filepath.Split(storePackageDir)

		if column.JsonType == "array" {
			goType.GoType = fmt.Sprintf("*%s.JsonArray", storePkg)

			goType.Import = storePackageDir
		}

		if column.JsonType == "object" {
			goType.GoType = fmt.Sprintf("*%s.JsonObject", storePkg)

			goType.Import = storePackageDir
		}

		return goType, nil

	case "bytea", "blob", "pg_catalog.bytea":
		goType.GoType = "[]byte"

		return goType, nil

	case "date", "time", "timetz", "timestamp", "timestamptz", "pg_catalog.time", "pg_catalog.timetz", "pg_catalog.date", "pg_catalog.timestamp", "pg_catalog.timestamptz":
		goType.Import = "time"

		goType.GoType = "*time.Time"

		return goType, nil

	case "text", "string", "citext", "name", "bpchar", "tsquery", "varchar", "pg_catalog.varchar", "pg_catalog.bpchar", "pg_catalog.tsquery":
		goType.GoType = "*string"

		return goType, nil

	case "inet", "cidr", "macaddr", "macaddr8", "pg_catalog.inet", "pg_catalog.cidr", "pg_catalog.macaddr", "pg_catalog.macaddr8":
		goType.Import = "net"

		goType.GoType = "*net.IP"

		return goType, nil

	case "ltree", "lquery", "ltxtquery":
		// This module implements a data type ltree for representing labels
		// of data stored in a hierarchical tree-like structure. Extensive
		// facilities for searching through label trees are provided.
		//
		// https://www.postgresql.org/docs/current/ltree.html
		goType.GoType = "*string"

		return goType, nil

	case "interval", "pg_catalog.interval":
		goType.Import = "time"

		goType.GoType = "*time.Duration"

		return goType, nil

	case "tsvector", "pg_catalog.tsvector":
		goType.GoType = "*string"

		return goType, nil

	case "uuid", "pg_catalog.uuid":
		goType.GoType = "*string"

		return goType, nil

	case "void", "pg_catalog.void":
		return goType, nil
	}

	return goType, nil
}

func fromArray(column introspect.Column) (types.GoType, error) {
	goType := types.GoType{
		DbType:    column.Type,
		GoType:    "*pq.GenericArray",
		IsPointer: true,
		Import:    "github.com/lib/pq",
	}

	switch column.Type {
	case "serial", "serial4", "pg_catalog.serial4":
		goType.GoType = "*pg.Int32Array"

		return goType, nil

	case "bigserial", "serial8", "pg_catalog.serial8":
		goType.GoType = "*pq.Int64Array"

		return goType, nil

	case "smallserial", "serial2", "pg_catalog.serial2":
		goType.GoType = "*pq.Int32Array"

		return goType, nil

	case "integer", "int", "int4", "pg_catalog.int4":
		goType.GoType = "*pq.Int32Array"

		return goType, nil

	case "bigint", "int8", "pg_catalog.int8":
		goType.GoType = "*pq.Int64Array"

		return goType, nil

	case "smallint", "int2", "pg_catalog.int2":
		goType.GoType = "*pq.Int32Array"

		return goType, nil

	case "float", "double precision", "float8", "pg_catalog.float8":
		goType.GoType = "*pq.Float64Array"

		return goType, nil

	case "real", "float4", "pg_catalog.float4":
		goType.GoType = "*pq.Float32Array"

		return goType, nil

	case "numeric", "pg_catalog.numeric", "money", "pg_catalog.money":
		// Since the Go standard library does not have a decimal type, lib/pq
		// returns numerics as strings.
		//
		// https://github.com/lib/pq/issues/648
		goType.GoType = "*pq.StringArray"

		return goType, nil

	case "boolean", "bool", "pg_catalog.bool":
		goType.GoType = "*pq.BoolArray"

		return goType, nil

	case "json", "jsonb", "pg_catalog.json", "pg_catalog.jsonb":
		goType.IsPointer = false

		goType.GoType = "json.RawMessage"

		goType.Import = "encoding/json"

		return goType, nil

	case "bytea", "blob", "pg_catalog.bytea":
		goType.GoType = "*pq.ByteaArray"

		return goType, nil

	case "date", "time", "timetz", "timestamp", "timestamptz", "pg_catalog.time", "pg_catalog.timetz", "pg_catalog.date", "pg_catalog.timestamp", "pg_catalog.timestamptz":
		goType.GoType = "*pq.GenericArray"

		return goType, nil

	case "text", "string", "citext", "name", "tsquery", "pg_catalog.text", "pg_catalog.varchar", "pg_catalog.bpchar", "pg_catalog.tsquery":
		goType.GoType = "*pq.StringArray"

		return goType, nil

	case "inet", "cidr", "macaddr", "macaddr8", "pg_catalog.inet", "pg_catalog.cidr", "pg_catalog.macaddr", "pg_catalog.macaddr8":
		goType.Import = "github.com/lib/pq"

		goType.GoType = "*pq.GenericArray"

		return goType, nil

	case "ltree", "lquery", "ltxtquery":
		// This module implements a data type ltree for representing labels
		// of data stored in a hierarchical tree-like structure. Extensive
		// facilities for searching through label trees are provided.
		//
		// https://www.postgresql.org/docs/current/ltree.html
		goType.GoType = "*pq.StringArray"

		return goType, nil

	case "interval", "pg_catalog.interval":
		goType.Import = "github.com/lib/pq"

		goType.GoType = "*pq.GenericArray"

		return goType, nil
	}

	return goType, nil
}
