package mysql

import (
	"bytes"
	_ "embed"
	"fmt"
	"regexp"
	"strings"
	"text/template"

	i "github.com/aakash-rajur/sqlxgen/internal/introspect"
	"github.com/aakash-rajur/sqlxgen/internal/introspect/prepare"
	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/aakash-rajur/sqlxgen/internal/utils/casing"
	"github.com/jmoiron/sqlx"
	"github.com/joomcode/errorx"
)

func (s source) IntrospectQueries(tx *sqlx.Tx) ([]i.Query, error) {
	validateQueryName, err := utils.CreateValidateEntityNames(s.args.QueryInclusions, s.args.QueryExclusions)

	if err != nil {
		return nil, err
	}

	qas := make([]QueryArgs, 0)

	qs := make([]i.Query, 0)

	for _, queryDir := range s.args.QueryDirs {
		qfds, err := s.fd.Find(queryDir, `[\w-]+.sql$`, false)

		if err != nil {
			return nil, errorx.InitializationFailed.Wrap(err, "failed to discover query files")
		}

		for _, qfd := range qfds {
			content, err := qfd.Load()

			if err != nil {
				errMsg := fmt.Sprintf("failed to load query file %s", qfd.GetFullPath())

				return nil, errorx.InitializationFailed.Wrap(err, errMsg)
			}

			fileName := qfd.GetFilename()

			if !validateQueryName(fileName) {
				continue
			}

			qa := QueryArgs{
				Query:    string(content),
				Filename: fileName,
				GenDir:   qfd.GetDir(),
			}

			qas = append(qas, qa)
		}

		for _, arg := range qas {
			q, err := introspectQuery(tx, arg)

			if err != nil {
				return nil, err
			}

			qs = append(qs, q)
		}
	}

	return qs, nil
}

func introspectQuery(tx *sqlx.Tx, args QueryArgs) (i.Query, error) {
	query := strings.TrimSpace(args.Query)

	introspectionQuery, err := generateIntrospectQuery(query)

	if err != nil {
		msg := msgWithFilename(args.Filename, "failed to generate introspection query")

		return i.Query{}, errorx.Decorate(err, msg)
	}

	jsonColumnMap, err := parseJsonColumns(query)

	if err != nil {
		msg := msgWithFilename(args.Filename, "failed to parse json columns")

		return i.Query{}, errorx.Decorate(err, msg)
	}

	nilParams, err := i.ParseParamsAsNil(query)

	if err != nil {
		msg := msgWithFilename(args.Filename, "failed to parse params as nil")

		return i.Query{}, errorx.Decorate(err, msg)
	}

	res, err := regexp.Compile(`--\n([\w\W\s\S]*?;)\n`)

	if err != nil {
		msg := msgWithFilename(args.Filename, "failed to compile query split regex")

		return i.Query{}, errorx.IllegalFormat.Wrap(err, msg)
	}

	statements := res.FindAllStringSubmatch(introspectionQuery, -1)

	if len(statements) != 3 {
		msg := msgWithFilename(args.Filename, "failed to generate introspection query")

		return i.Query{}, errorx.IllegalFormat.New(msg)
	}

	dropQ, createQ, selectQ := statements[0][1], statements[1][1], statements[2][1]

	_, err = tx.Exec(dropQ)

	if err != nil {
		msg := msgWithFilename(args.Filename, "failed to execute drop query")

		return i.Query{}, errorx.InternalError.Wrap(err, msg)
	}

	_, err = tx.NamedExec(createQ, nilParams)

	if err != nil {
		msg := msgWithFilename(args.Filename, "failed to execute create query")

		return i.Query{}, errorx.InternalError.Wrap(err, msg)
	}

	rows, err := tx.Queryx(selectQ)

	if err != nil {
		msg := msgWithFilename(args.Filename, "failed to execute introspection query")

		return i.Query{}, errorx.InternalError.Wrap(err, msg)
	}

	defer func(rows *sqlx.Rows) {
		err := rows.Close()

		if err != nil {
			println(err.Error())
		}
	}(rows)

	columns := make([]i.Column, 0)

	for rows.Next() {
		column := i.Column{}

		err := rows.StructScan(&column)

		if err != nil {
			return i.Query{}, errorx.InternalError.Wrap(err, "failed to scan introspection query result")
		}

		if jsonType, ok := jsonColumnMap[column.ColumnName]; ok {
			column.JsonType = getJsonType(jsonType)
		}

		columns = append(columns, column)
	}

	params, err := i.ParseParams(query)

	if err != nil {
		msg := msgWithFilename(args.Filename, "failed to parse query params")

		return i.Query{}, errorx.Decorate(err, msg)
	}

	fileName, _ := utils.SplitFilename(args.Filename)

	queryName, err := casing.PascalCase(fileName)

	if err != nil {
		msg := msgWithFilename(args.Filename, "failed to generate query name")

		return i.Query{}, errorx.Decorate(err, msg)
	}

	filename, err := casing.SnakeCase(args.Filename)

	if err != nil {
		msg := msgWithFilename(args.Filename, "failed to generate query name")

		return i.Query{}, errorx.Decorate(err, msg)
	}

	pascalName, err := casing.PascalCase(filename)

	if err != nil {
		return i.Query{}, errorx.InternalError.Wrap(err, "failed to generate query name")
	}

	camelName, err := casing.CamelCase(filename)

	if err != nil {
		return i.Query{}, errorx.InternalError.Wrap(err, "failed to generate query name")
	}

	q := i.Query{
		Filename:   args.Filename,
		QueryName:  queryName,
		PascalName: pascalName,
		CamelName:  camelName,
		Params:     params,
		Columns:    columns,
		SourceDir:  args.GenDir,
	}

	return q, nil
}

func generateIntrospectQuery(query string) (string, error) {
	preparedQuery, err := prepare.PrepareQuery(query)

	if err != nil {
		return "", errorx.InternalError.Wrap(err, "failed to prepare query")
	}

	tmpl, err := template.New("introspect").Parse(introspectQuerySql)

	if err != nil {
		return "", errorx.IllegalFormat.Wrap(err, "failed to parse introspect query template")
	}

	var introspectQueryBuffer bytes.Buffer

	err = tmpl.Execute(
		&introspectQueryBuffer,
		map[string]interface{}{
			"Query": preparedQuery,
		},
	)

	if err != nil {
		return "", errorx.IllegalFormat.Wrap(err, "failed to execute introspect query template")
	}

	introspectionQuery := introspectQueryBuffer.String()

	return introspectionQuery, nil
}

//go:embed query.sql
var introspectQuerySql string

type QueryArgs struct {
	Query    string
	Filename string
	GenDir   string
}

func msgWithFilename(filename string, msg string) string {
	return fmt.Sprintf("%s: %s", filename, msg)
}
