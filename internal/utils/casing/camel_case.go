package casing

import (
	"regexp"
	"strings"

	"github.com/aakash-rajur/sqlxgen/internal/utils/array"
	"github.com/joomcode/errorx"
)

func CamelCase(identifier string) (string, error) {
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

	pascal := array.Reduce(
		words,
		func(accumulator string, word string, index int) string {
			if index == 0 {
				return accumulator + strings.ToLower(word)
			}

			return accumulator + toTitleCase(word)
		},
		"",
	)

	return pascal, nil
}
