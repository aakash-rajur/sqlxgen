package mysql

import (
	_ "embed"

	"github.com/aakash-rajur/sqlxgen/internal/generate/types"
	"github.com/aakash-rajur/sqlxgen/internal/introspect"
)

type Mysql struct{}

func (mysql Mysql) Infer(column introspect.Column) (types.GoType, error) {
	return infer(column)
}

func (mysql Mysql) ModelTemplate() string {
	return modelTemplate
}

func (mysql Mysql) QueryTemplate() string {
	return queryTemplate
}

//go:embed model.go.tmpl
var modelTemplate string

//go:embed query.go.tmpl
var queryTemplate string

func NewTranslate() types.Translate {
	return Mysql{}
}
