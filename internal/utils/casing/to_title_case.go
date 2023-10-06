package casing

import (
	"strings"
)

func toTitleCase(word string) string {
	if word == "" {
		return ""
	}

	return strings.ToUpper(word[:1]) + word[1:]
}
