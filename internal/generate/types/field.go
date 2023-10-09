package types

import (
	"github.com/aakash-rajur/sqlxgen/internal/introspect"
	"github.com/aakash-rajur/sqlxgen/internal/utils/casing"
)

type Field struct {
	Name   string            `json:"name"`
	Type   GoType            `json:"type"`
	Column introspect.Column `json:"column"`
}

func NewField(
	column introspect.Column,
	translate Translate,
	storePackageDir string,
	storePackageName string,
) (Field, error) {
	fieldName, err := casing.PascalCase(column.ColumnName)

	if err != nil {
		return Field{}, err
	}

	goType, err := translate.Infer(storePackageDir, storePackageName, column)

	if err != nil {
		return Field{}, err
	}

	field := Field{
		Name:   fieldName,
		Type:   goType,
		Column: column,
	}

	return field, nil
}

type GoType struct {
	DbType    string
	GoType    string
	Import    string
	IsPointer bool
}
