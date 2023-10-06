package array

func From[T any](
	length int,
	fn fromFn[T],
) []T {
	result := make([]T, length)

	for i := 0; i < length; i += 1 {
		result[i] = fn(i)
	}

	return result
}

type fromFn[T any] func(offset int) T
