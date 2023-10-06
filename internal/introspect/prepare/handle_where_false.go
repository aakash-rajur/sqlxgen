package prepare

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/joomcode/errorx"
)

func handleWhereFalseClause(query string) (string, error) {
	whereContexts, err := splitWhereContexts(query)

	if err != nil {
		return "", errorx.InternalError.Wrap(err, "failed to find where contexts")
	}

	whereRe, err := whereFalseRegexPattern()

	if err != nil {
		return "", errorx.IllegalFormat.Wrap(err, "failed to compile where regex")
	}

	cursor := query

	for _, whereContext := range whereContexts {
		preparedPartial := whereRe.ReplaceAllString(whereContext, "$1 false and (\n$2\n)\n$3")

		cursor = strings.ReplaceAll(cursor, whereContext, preparedPartial)
	}

	return cursor, nil
}

func whereFalseRegexPattern() (*regexp.Regexp, error) {
	followedByTokens := strings.Join(
		[]string{
			`returning`,
			`group by`,
			`order by`,
			`limit`,
			`offset`,
			`;\s*\z`,
			`\z`,
		},
		"|",
	)

	followedByTokensPartial := fmt.Sprintf(`(%s)`, followedByTokens)

	pattern := strings.Join(
		[]string{
			`(where)`,              // where anchor
			`\s*([\w\W\s\S]*?)\s*`, // all where conditions
			followedByTokensPartial,
		},
		"",
	)

	return regexp.Compile(pattern)
}
