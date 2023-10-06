package prepare

import (
	"strings"
)

func PrepareQuery(query string) (string, error) {
	handler := []func(string) (string, error){
		handleNoComments,
		handleWhereFalseClause,
		handleEnsureWhereClause,
		handleSequence,
	}

	cursor := strings.TrimSpace(query)

	for _, handler := range handler {
		prepared, err := handler(cursor)

		if err != nil {
			return "", err
		}

		cursor = prepared
	}

	return cursor, nil
}
