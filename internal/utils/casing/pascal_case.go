package casing

import (
	"regexp"

	"github.com/aakash-rajur/sqlxgen/internal/utils/array"
	"github.com/joomcode/errorx"
)

func PascalCase(identifier string) (string, error) {
	if identifier == "" {
		return "", nil
	}

	camelRe, err := regexp.Compile(`([a-z])([A-Z])`)

	if err != nil {
		return "", errorx.IllegalFormat.Wrap(err, "failed to compile regex")
	}

	unCamel := camelRe.ReplaceAllString(identifier, "${1} ${2}")

	wordRe, err := regexp.Compile(`[-_\s]`)

	if err != nil {
		return "", errorx.IllegalFormat.Wrap(err, "failed to compile regex")
	}

	words := wordRe.Split(unCamel, -1)

	pascal := array.Reduce(
		words,
		func(accumulator string, word string, _ int) string {
			return accumulator + toTitleCase(word)
		},
		"",
	)

	return pascal, nil
}
