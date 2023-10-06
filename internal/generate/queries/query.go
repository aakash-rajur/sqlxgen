package queries

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"go/format"
	"log/slog"
	"path"
	"text/template"

	"github.com/aakash-rajur/sqlxgen/internal/generate/types"
	"github.com/aakash-rajur/sqlxgen/internal/introspect"
	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/aakash-rajur/sqlxgen/internal/utils/array"
	"github.com/aakash-rajur/sqlxgen/internal/utils/casing"
	"github.com/aakash-rajur/sqlxgen/internal/utils/writer"
	mapset "github.com/deckarep/golang-set"
	"github.com/joomcode/errorx"
)

type queryModel struct {
	WriterCreator    writer.Creator `json:"-"`
	StorePackageDir  string         `json:"store_package_dir"`
	StorePackageName string         `json:"store_package_name"`
	Filename         string         `json:"filename"`
	PascalName       string         `json:"pascal_name"`
	CamelName        string         `json:"camel_name"`
	Fields           []types.Field  `json:"fields"`
	Params           []types.Field  `json:"params"`
	GenDir           string         `json:"-"`
}

func (qm *queryModel) getImports() []string {
	uniqueImports := mapset.NewSet()

	for _, f := range qm.Fields {
		if f.Type.Import == "" {
			continue
		}

		uniqueImports.Add(f.Type.Import)
	}

	for _, f := range qm.Params {
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

	imports = append(imports, qm.StorePackageDir)

	return imports
}

func (qm *queryModel) generate(queryTemplate string) error {
	slog.Debug("generating query", "query", qm.Filename)

	helpers := template.FuncMap{
		"ToJson": func(v interface{}) string {
			b, _ := json.MarshalIndent(v, "", " ")

			return string(b)
		},
	}

	tmpl, err := template.New("query").Funcs(helpers).Parse(queryTemplate)

	if err != nil {
		return errorx.IllegalFormat.Wrap(err, "unable to parse model template")
	}

	packageNameBaseDir := path.Base(qm.GenDir)

	packageName, err := casing.SnakeCase(packageNameBaseDir)

	imports := qm.getImports()

	var queryFileBuffer bytes.Buffer

	err = tmpl.Execute(
		&queryFileBuffer,
		map[string]interface{}{
			"PackageName": packageName,
			"Imports":     imports,
			"Query":       qm,
		},
	)

	if err != nil {
		return errorx.InternalError.Wrap(err, "unable to execute model template")
	}

	formatted, err := format.Source(queryFileBuffer.Bytes())

	if err != nil {
		return err
	}

	qfn, _ := utils.SplitFilename(qm.Filename)

	qfnSnakeName, err := casing.SnakeCase(qfn)

	if err != nil {
		return err
	}

	queryFileName := utils.FilenameWithGen(qfnSnakeName + ".go")

	queryFilePath := path.Join(qm.GenDir, queryFileName)

	pen := qm.WriterCreator(queryFilePath, string(formatted))

	err = pen.Write()

	if err != nil {
		return errorx.InternalError.Wrap(err, "unable to write generated query file")
	}

	slog.Debug("generated query", "query", qm.Filename)

	return nil
}

func newQueryModel(
	writerCreator writer.Creator,
	translate types.Translate,
	projectDir string,
	storePackageDir string,
	storePackageName string,
	query introspect.Query,
) (queryModel, error) {
	filename, _ := utils.SplitFilename(query.Filename)

	pascalName, err := casing.PascalCase(filename)

	if err != nil {
		return queryModel{}, errorx.InternalError.Wrap(err, "failed to convert filename to pascal case")
	}

	camelName, err := casing.CamelCase(filename)

	if err != nil {
		return queryModel{}, errorx.InternalError.Wrap(err, "failed to convert filename to camel case")
	}

	fields := make([]types.Field, len(query.Columns))

	for i, column := range query.Columns {
		f, err := types.NewField(column, translate)

		if err != nil {
			return queryModel{}, err
		}

		fields[i] = f
	}

	params := make([]types.Field, len(query.Params))

	for i, param := range query.Params {
		p, err := types.NewField(param, translate)

		if err != nil {
			return queryModel{}, err
		}

		params[i] = p
	}

	genDir := path.Join(projectDir, query.SourceDir)

	qm := queryModel{
		WriterCreator:    writerCreator,
		StorePackageDir:  storePackageDir,
		StorePackageName: storePackageName,
		Filename:         query.Filename,
		PascalName:       pascalName,
		CamelName:        camelName,
		Fields:           fields,
		Params:           params,
		GenDir:           genDir,
	}

	return qm, nil
}
