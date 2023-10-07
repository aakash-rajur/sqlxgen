package mysql

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
	case "json_arrayagg", "array":
		return "array"
	case "json_objectagg", "json_object", "object":
		return "object"
	}

	return "identity"
}
