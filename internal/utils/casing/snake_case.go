package casing

import (
	"regexp"
	"strings"

	"github.com/joomcode/errorx"
)

func SnakeCase(identifier string) (string, error) {
	pascalRe, err := regexp.Compile(`([A-Z])([a-z])`)

	if err != nil {
		return "", errorx.IllegalFormat.Wrap(err, "failed to compile regex")
	}

	unPascal := strings.TrimSpace(pascalRe.ReplaceAllString(identifier, " ${1}${2}"))

	wordRe, err := regexp.Compile(`[-_\s]`)

	if err != nil {
		return "", errorx.IllegalFormat.Wrap(err, "failed to compile regex")
	}

	words := wordRe.Split(unPascal, -1)

	snake := strings.ToLower(strings.Join(words, "_"))

	return snake, nil
}
