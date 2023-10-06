package prepare

import (
	"regexp"
	"strings"

	"github.com/joomcode/errorx"
)

func handleNoComments(query string) (string, error) {
	commentRe, err := regexp.Compile(`-{2,}\s*([\w\W\s\S]*?)(\n|\z)`)

	if err != nil {
		return "", errorx.IllegalFormat.Wrap(err, "failed to compile comment regex")
	}

	prepared := strings.TrimSpace(commentRe.ReplaceAllString(query, "\n"))

	return prepared, nil
}
