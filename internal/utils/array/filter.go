package array

func Filter[T any](slice []T, filter filterFn[T]) []T {
	filtered := make([]T, 0)

	for index, item := range slice {
		if !filter(item, index) {
			continue
		}

		filtered = append(filtered, item)
	}

	return filtered
}

type filterFn[T any] func(each T, index int) bool
