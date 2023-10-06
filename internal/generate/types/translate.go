package types

import (
	"strings"

	"github.com/aakash-rajur/sqlxgen/internal/introspect"
)

type Translate interface {
	Infer(column introspect.Column) (GoType, error)

	ModelTemplate() string

	QueryTemplate() string
}

type fakeTranslate struct {
	ModelTemplateContent string
	QueryTemplateContent string
}

func (t fakeTranslate) Infer(column introspect.Column) (GoType, error) {
	goType := GoType{
		DbType:    column.Type,
		GoType:    "interface{}",
		IsPointer: true,
	}

	columnType := column.Type

	if strings.Contains(columnType, "int") {
		goType.GoType = "*int"

		return goType, nil
	}

	if columnType == "date" {
		goType.Import = "time"

		goType.GoType = "*time.Time"

		return goType, nil
	}

	goType.GoType = "*string"

	return goType, nil
}

func (t fakeTranslate) ModelTemplate() string {
	return t.ModelTemplateContent
}

func (t fakeTranslate) QueryTemplate() string {
	return t.QueryTemplateContent
}

func NewFakeTranslate(
	modelTemplateContent string,
	queryTemplateContent string,
) Translate {
	return fakeTranslate{
		ModelTemplateContent: modelTemplateContent,
		QueryTemplateContent: queryTemplateContent,
	}
}
