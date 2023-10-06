package prepare

import (
	"regexp"

	"github.com/joomcode/errorx"
)

func handleSequence(query string) (string, error) {
	re, err := regexp.Compile(`(nextval)(\s*\()`)

	if err != nil {
		return "", errorx.IllegalFormat.Wrap(err, "failed to compile sequence regex")
	}

	prepared := re.ReplaceAllString(query, "currval$2")

	return prepared, nil
}
