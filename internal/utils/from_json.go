package utils

import (
	"encoding/json"
)

func FromJson[T any](jsons []string) ([]T, error) {
	instances := make([]T, len(jsons))

	for i, source := range jsons {
		var instance T

		err := json.Unmarshal([]byte(source), &instance)

		if err != nil {
			return nil, err
		}

		instances[i] = instance
	}

	return instances, nil
}
