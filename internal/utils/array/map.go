package array

func Map[T any, R any](slice []T, mapper mapperFn[T, R]) []R {
	mapped := make([]R, len(slice))

	for index, item := range slice {
		mapped[index] = mapper(item, index)
	}

	return mapped
}

type mapperFn[T any, R any] func(each T, index int) R
