package introspect

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/aakash-rajur/sqlxgen/internal/utils/array"
)

type Column struct {
	ColumnName        string `db:"column_name" json:"column_name"`
	Type              string `db:"type" json:"type"`
	TypeId            string `db:"type_id" json:"type_id"`
	IsArray           bool   `db:"is_array" json:"is_array"`
	IsSequence        bool   `db:"is_sequence" json:"is_sequence"`
	Nullable          bool   `db:"nullable" json:"nullable"`
	Generated         bool   `db:"generated" json:"generated"`
	PkName            string `db:"pk_name" json:"pk_name"`
	PkOrdinalPosition int    `db:"pk_ordinal_position" json:"pk_ordinal_position"`
	JsonType          string `db:"json_type" json:"json_type"`
}

func (column *Column) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("ColumnName: %v", column.ColumnName),
			fmt.Sprintf("Type: %v", column.Type),
			fmt.Sprintf("TypeId: %v", column.TypeId),
			fmt.Sprintf("IsArray: %v", column.IsArray),
			fmt.Sprintf("Nullable: %v", column.Nullable),
			fmt.Sprintf("Generated: %v", column.Generated),
			fmt.Sprintf("PkName: %v", column.PkName),
			fmt.Sprintf("PkOrdinalPosition: %v", column.PkOrdinalPosition),
			fmt.Sprintf("JsonType: %v", column.JsonType),
		},
		", ",
	)

	return fmt.Sprintf("Column{%s}", content)
}

func (column *Column) Value() (driver.Value, error) {
	return json.Marshal(column)
}

func (column *Column) Scan(value interface{}) error {
	buffer, ok := value.([]byte)

	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(buffer, column)
}

type Columns []Column

func (columns *Columns) String() string {
	content := strings.Join(
		array.Map(
			*columns,
			func(column Column, _ int) string {
				return (&column).String()
			},
		),
		", ",
	)

	return fmt.Sprintf("[%s]", content)
}

func (columns *Columns) Value() (driver.Value, error) {
	return json.Marshal(columns)
}

func (columns *Columns) Scan(value interface{}) error {
	buffer, ok := value.([]byte)

	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(buffer, columns)
}
