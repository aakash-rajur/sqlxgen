package pg

import (
	"regexp"
	"strings"
)

func parseJsonColumns(query string) (map[string]string, error) {
	jsonAnnotations, err := parseJsonAnnotations(query)

	if err != nil {
		return nil, err
	}

	re, err := regexp.Compile(`([a-zA-Z0-9_]+_b?json|jsonb?_[_a-zA-Z0-9]+)\s*\([\w\W\s\S]*?\)\s*as\s+"?([a-zA-Z0-9_]+)"?,?`)

	if err != nil {
		return nil, err
	}

	matches := re.FindAllStringSubmatch(strings.ToLower(query), -1)

	matchesMap := make(map[string]string)

	for _, match := range matches {
		aggFn, columnName := match[1], match[2]

		annotation, ok := jsonAnnotations[columnName]

		if ok {
			matchesMap[columnName] = annotation

			continue
		}

		matchesMap[columnName] = aggFn
	}

	return matchesMap, nil
}

func parseJsonAnnotations(query string) (map[string]string, error) {
	re, err := regexp.Compile(`-{2,}\s*column:\s*([a-zA-Z_]+)\s+json_type:\s*(array|object)`)

	if err != nil {
		return nil, err
	}

	matches := re.FindAllStringSubmatch(strings.ToLower(query), -1)

	matchesMap := make(map[string]string)

	for _, match := range matches {
		columnName, jsonType := match[1], match[2]

		matchesMap[columnName] = jsonType
	}

	return matchesMap, nil
}

func getJsonType(jsonFn string) string {
	switch jsonFn {
	case "array_to_json", "json_build_array", "jsonb_build_array", "json_agg", "jsonb_agg", "array":
		return "array"
	case "row_to_json", "json_build_object", "jsonb_build_object", "json_object", "jsonb_object", "jsonb_object_agg", "object":
		return "object"
	}

	return "identity"
}
