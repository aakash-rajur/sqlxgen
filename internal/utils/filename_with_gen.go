package utils

import (
	"strings"
)

func FilenameWithGen(filename string) string {
	name, ext := SplitFilename(filename)

	return strings.Join([]string{name, "gen", ext}, ".")
}
