package models

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"go/format"
	"log/slog"
	"path"
	"reflect"
	"strings"
	"text/template"

	"github.com/aakash-rajur/sqlxgen/internal/generate/types"
	"github.com/aakash-rajur/sqlxgen/internal/introspect"
	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/aakash-rajur/sqlxgen/internal/utils/array"
	"github.com/aakash-rajur/sqlxgen/internal/utils/casing"
	"github.com/aakash-rajur/sqlxgen/internal/utils/writer"
	mapset "github.com/deckarep/golang-set"
	"github.com/jinzhu/inflection"
	"github.com/joomcode/errorx"
)

type model struct {
	WriterCreator    writer.Creator   `json:"-"`
	StorePackageDir  string           `json:"store_package_dir"`
	StorePackageName string           `json:"store_package_name"`
	ModelTemplate    string           `json:"model_template"`
	FileName         string           `json:"file_name"`
	PascalName       string           `json:"pascal_name"`
	CamelName        string           `json:"camel_name"`
	Fields           []types.Field    `json:"fields"`
	PkFields         []types.Field    `json:"pk_fields"`
	Table            introspect.Table `json:"table"`
}

func (m model) getImports() []string {
	uniqueImports := mapset.NewSet()

	for _, f := range m.Fields {
		if f.Type.Import == "" {
			continue
		}

		uniqueImports.Add(f.Type.Import)
	}

	importSlice := uniqueImports.ToSlice()

	imports := array.Map(
		importSlice,
		func(each interface{}, index int) string {
			return each.(string)
		},
	)

	return imports
}

func (m model) generate(
	modelTemplate string,
	packageName string,
	genDir string,
) error {
	slog.Debug("generating model", "table", m.Table.TableName, "model", m.PascalName)

	helpers := template.FuncMap{
		"isLast": func(index int, array interface{}) bool {
			return index == reflect.ValueOf(array).Len()-1
		},
		"isFirst": func(index int) bool {
			return index == 0
		},
		"ToUpper": strings.ToUpper,
		"ToJson": func(v interface{}) string {
			b, _ := json.Marshal(v)

			return string(b)
		},
	}

	tmpl, err := template.New("model").Funcs(helpers).Parse(modelTemplate)

	if err != nil {
		return errorx.IllegalFormat.Wrap(err, "unable to parse model template")
	}

	imports := m.getImports()

	insertFields, updateFields, selectFields := distinguishFields(m.Fields)

	var modelFileBuffer bytes.Buffer

	err = tmpl.Execute(
		&modelFileBuffer,
		map[string]interface{}{
			"PackageName":  packageName,
			"Imports":      imports,
			"Model":        m,
			"InsertFields": insertFields,
			"UpdateFields": updateFields,
			"SelectFields": selectFields,
		},
	)

	if err != nil {
		return errorx.InternalError.Wrap(err, "unable to execute model template")
	}

	formatted, err := format.Source(modelFileBuffer.Bytes())

	if err != nil {
		return err
	}

	genFileName := utils.FilenameWithGen(m.FileName + ".go")

	modelFilePath := path.Join(genDir, genFileName)

	pen := m.WriterCreator(modelFilePath, string(formatted))

	err = pen.Write()

	if err != nil {
		return errorx.InternalError.Wrap(err, "unable to write model file")
	}

	slog.Debug("generated model", "table", m.Table.TableName, "model", m.PascalName)

	return nil
}

func newModel(
	writerCreator writer.Creator,
	translate types.Translate,
	storePackageDir string,
	storePackageName string,
	table introspect.Table,
) (model, error) {
	singular := inflection.Singular(table.TableName)

	fileName, err := casing.SnakeCase(singular)

	if err != nil {
		return model{}, err
	}

	pascalName, err := casing.PascalCase(singular)

	if err != nil {
		return model{}, err
	}

	camelName, err := casing.CamelCase(singular)

	if err != nil {
		return model{}, err
	}

	fields := make([]types.Field, len(table.Columns))

	pkFields := make([]types.Field, 0)

	for index, column := range table.Columns {
		f, err := types.NewField(column, translate, storePackageDir, storePackageName)

		if err != nil {
			return model{}, err
		}

		fields[index] = f

		if column.PkOrdinalPosition > 0 {
			pkFields = append(pkFields, f)
		}
	}

	if len(pkFields) == 0 {
		slog.Warn("no primary key found for table", "table", table.TableName)
		pkFields = array.Filter(
			fields,
			func(each types.Field, index int) bool {
				return !each.Column.Generated
			},
		)
	}

	m := model{
		WriterCreator:    writerCreator,
		StorePackageDir:  storePackageDir,
		StorePackageName: storePackageName,
		FileName:         fileName,
		PascalName:       pascalName,
		CamelName:        camelName,
		Fields:           fields,
		PkFields:         pkFields,
		Table:            table,
	}

	return m, nil
}

func distinguishFields(fields []types.Field) ([]types.Field, []types.Field, []types.Field) {
	insertFields := make([]types.Field, 0)

	updateFields := make([]types.Field, 0)

	selectFields := make([]types.Field, 0)

	for _, field := range fields {
		isSequence := field.Column.IsSequence

		isGenerated := field.Column.Generated

		if !isSequence && !isGenerated {
			insertFields = append(insertFields, field)
		}

		if !isGenerated {
			updateFields = append(updateFields, field)
		}

		selectFields = append(selectFields, field)
	}

	return insertFields, updateFields, selectFields
}
