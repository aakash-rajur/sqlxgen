package mysql

import (
	"github.com/aakash-rajur/sqlxgen/internal/generate/types"
	"github.com/aakash-rajur/sqlxgen/internal/introspect"
)

func infer(column introspect.Column) (types.GoType, error) {
	goType := types.GoType{
		DbType:    column.Type,
		GoType:    "interface{}",
		IsPointer: true,
	}

	switch column.Type {
	case "varchar", "text", "longtext", "mediumtext", "char":
		goType.GoType = "*string"

		return goType, nil

	case "tinyint", "smallint":
		goType.GoType = "*int16"

		return goType, nil

	case "int":
		goType.GoType = "*int32"

		return goType, nil

	case "bigint":
		goType.GoType = "*int64"

		return goType, nil

	case "float":
		goType.GoType = "*float32"

		return goType, nil

	case "double":
		goType.GoType = "*float64"

		return goType, nil

	case "decimal":
		goType.GoType = "*string"

		return goType, nil

	case "time", "timestamp", "datetime", "date":
		goType.Import = "time"

		goType.GoType = "*time.Time"

		return goType, nil

	case "json":
		goType.IsPointer = false

		goType.GoType = "json.RawMessage"

		goType.Import = "encoding/json"

		if column.JsonType == "array" {
			goType.GoType = "[]map[string]interface{}"

			goType.Import = ""
		}

		if column.JsonType == "object" {
			goType.GoType = "map[string]interface{}"

			goType.Import = ""
		}

		return goType, nil

	case "set":
		goType.GoType = "*string"

		return goType, nil

	case "binary", "varbinary", "blob", "longblob", "mediumblob":
		goType.GoType = "*pq.ByteaArray"

		goType.Import = "github.com/lib/pq"

		return goType, nil
	}

	return goType, nil
}
