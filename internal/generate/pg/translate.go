package go_type

import (
	_ "embed"

	"github.com/aakash-rajur/sqlxgen/internal/generate/types"
	"github.com/aakash-rajur/sqlxgen/internal/introspect"
)

type Pg struct{}

func (pg Pg) Infer(column introspect.Column) (types.GoType, error) {
	return infer(column)
}

func (pg Pg) ModelTemplate() string {
	return modelTemplate
}

func (pg Pg) QueryTemplate() string {
	return queryTemplate
}

//go:embed model.go.tmpl
var modelTemplate string

//go:embed query.go.tmpl
var queryTemplate string

func NewTranslate() types.Translate {
	return Pg{}
}
