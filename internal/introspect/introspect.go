package introspect

import (
	"github.com/jmoiron/sqlx"
)

type Introspect interface {
	IntrospectSchema(tx *sqlx.Tx) ([]Table, error)

	IntrospectQueries(tx *sqlx.Tx) ([]Query, error)
}
