package utils

import (
	"strings"
)

func SplitFilename(filename string) (string, string) {
	splits := strings.Split(filename, ".")

	splitsLen := len(splits)

	if splitsLen == 0 {
		return "", ""
	}

	if len(splits) == 1 {
		return splits[0], ""
	}

	baseName := strings.Join(splits[:splitsLen-1], ".")

	return baseName, splits[splitsLen-1]
}
