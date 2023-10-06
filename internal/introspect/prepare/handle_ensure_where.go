package prepare

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/joomcode/errorx"
)

func handleEnsureWhereClause(query string) (string, error) {
	withRe, err := regexp.Compile(`[)\s]+where[(\s]+`)

	if err != nil {
		return "", errorx.IllegalFormat.Wrap(err, "failed to compile with regex")
	}

	hasWhere := withRe.MatchString(query)

	if hasWhere {
		return query, nil
	}

	re, err := regexp.Compile(`(\s+(returning|group by|order by|limit|offset)\s+|;\z)`)

	if err != nil {
		return "", errorx.IllegalFormat.Wrap(err, "failed to compile with regex")
	}

	firstMatch := re.FindString(query)

	if firstMatch == "" {
		return query, nil
	}

	withPartial := fmt.Sprintf("\nwhere false\n%s", firstMatch)

	prepared := strings.TrimSpace(strings.Replace(query, firstMatch, withPartial, 1))

	return prepared, nil
}
