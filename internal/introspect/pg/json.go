package pg

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
		aggFn, columnName := match[1], match[2]

		if len(match) == 4 && match[4] == "many" {
			matchesMap[columnName] = "array"

			continue
		}

		matchesMap[columnName] = aggFn
	}

	return matchesMap, nil
}

func getJsonType(jsonFn string) string {
	switch jsonFn {
	case "array_to_json", "json_build_array", "jsonb_build_array", "json_agg", "jsonb_agg":
		return "array"
	case "row_to_json", "json_build_object", "jsonb_build_object", "json_object", "jsonb_object", "jsonb_object_agg":
		return "object"
	}

	return "identity"
}
