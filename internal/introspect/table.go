package introspect

import (
	"fmt"
	"strings"

	"github.com/aakash-rajur/sqlxgen/internal/utils/array"
)

type Table struct {
	SchemaName string  `db:"schema_name" json:"schema_name"`
	TableName  string  `db:"table_name" json:"table_name"`
	Columns    Columns `db:"columns" json:"columns"`
}

func (table *Table) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("SchemaName: %v", table.SchemaName),
			fmt.Sprintf("TableName: %v", table.TableName),
			fmt.Sprintf("Columns: %v", &table.Columns),
		},
		", ",
	)

	return fmt.Sprintf("Table{%s}", content)
}

func (table *Table) PrimaryKey() Columns {
	primaryKeyColumns := array.Filter(
		table.Columns,
		func(column Column, _ int) bool {
			return column.PkOrdinalPosition > 0
		},
	)

	if len(primaryKeyColumns) > 0 {
		return primaryKeyColumns
	}

	return table.Columns
}
