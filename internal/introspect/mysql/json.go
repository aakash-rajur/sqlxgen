package mysql

import (
	"regexp"
	"strings"
)

func parseJsonColumns(query string) (map[string]string, error) {
	re, err := regexp.Compile(`([a-zA-Z0-9_]+_b?json|jsonb?_[_a-zA-Z0-9]+)\s*\([\w\W\s\S]*?\)\s*as\s+"?([a-zA-Z0-9_]+)"?,?(\s*--\s*:(many|one))?`)

	if err != nil {
		return nil, err
	}

	matches := re.FindAllStringSubmatch(strings.ToLower(query), -1)

	matchesMap := make(map[string]string)

	for _, match := range matches {
		if len(match) == 4 && match[4] == "many" {
			matchesMap[match[2]] = "array"

			continue
		}

		matchesMap[match[2]] = match[1]
	}

	return matchesMap, nil
}

func getJsonType(jsonFn string) string {
	switch jsonFn {
	case "json_arrayagg":
		return "array"
	case "json_objectagg", "json_object":
		return "object"
	}

	return "identity"
}
