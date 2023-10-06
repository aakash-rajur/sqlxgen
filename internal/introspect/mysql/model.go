package mysql

import (
	_ "embed"
	"fmt"

	i "github.com/aakash-rajur/sqlxgen/internal/introspect"
	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/jmoiron/sqlx"
	"github.com/joomcode/errorx"
)

func (s source) IntrospectSchema(tx *sqlx.Tx) ([]i.Table, error) {
	tables := make([]i.Table, 0)

	validateTableName, err := utils.CreateValidateEntityNames(s.args.TableInclusions, s.args.TableExclusions)

	if err != nil {
		return tables, err
	}

	for _, schema := range s.args.Schemas {
		rows, err := tx.NamedQuery(
			introspectSchemaSql,
			map[string]interface{}{
				"schema": schema,
			},
		)

		if err != nil {
			msg := msgWithSchema(schema, "failed to introspect schema")

			return tables, errorx.Decorate(err, msg)
		}

		for rows.Next() {
			table := i.Table{}

			err = rows.StructScan(&table)

			if err != nil {
				msg := msgWithSchema(schema, "failed to scan table")

				return tables, errorx.Decorate(err, msg)
			}

			fullTableName := table.TableName

			if !validateTableName(fullTableName) {
				continue
			}

			tables = append(tables, table)
		}
	}

	return tables, nil
}

//go:embed model.sql
var introspectSchemaSql string

func msgWithSchema(schema string, msg string) string {
	return fmt.Sprintf("%s: %s", schema, msg)
}
