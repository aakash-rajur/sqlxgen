package introspect

import (
	"cmp"
	"regexp"
	"slices"

	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/joomcode/errorx"
)

func ParseParams(query string) ([]Column, error) {
	re1, err := regexp.Compile(`\s*:(\w+)\s+type:\s+(\w+),?\s*`)

	if err != nil {
		return nil, errorx.IllegalFormat.Wrap(err, "failed to compile params type regex")
	}

	typeMap := make(map[string]string)

	matches := re1.FindAllStringSubmatch(query, -1)

	for _, match := range matches {
		typeMap[match[1]] = match[2]
	}

	uniqueParams := make(map[string]Column)

	re2, err := regexp.Compile(`[^:][:@](\w+)`)

	if err != nil {
		return nil, errorx.IllegalFormat.Wrap(err, "failed to compile params regex")
	}

	matches = re2.FindAllStringSubmatch(query, -1)

	for _, match := range matches {
		columnName := match[1]

		if _, ok := uniqueParams[columnName]; ok {
			continue
		}

		column := Column{
			ColumnName:        columnName,
			Type:              "any",
			TypeId:            "0",
			Nullable:          true,
			PkName:            "",
			PkOrdinalPosition: 0,
			JsonType:          "identity",
		}

		if columnType, ok := typeMap[columnName]; ok {
			column.Type = columnType
		}

		uniqueParams[columnName] = column
	}

	params := utils.Values(uniqueParams)

	slices.SortStableFunc(
		params,
		func(i, j Column) int { return cmp.Compare(i.ColumnName, j.ColumnName) },
	)

	return params, nil
}

func ParseParamsAsNil(query string) (map[string]any, error) {
	re, err := regexp.Compile(`[^:][:@](\w+)`)

	if err != nil {
		return nil, errorx.IllegalFormat.Wrap(err, "failed to compile params regex")
	}

	matches := re.FindAllStringSubmatch(query, -1)

	params := make(map[string]any)

	for _, match := range matches {
		params[match[1]] = nil
	}

	return params, nil
}
