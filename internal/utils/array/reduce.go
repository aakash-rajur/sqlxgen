package array

func Reduce[T any, R any](
	slice []T,
	reducer reduceFn[T, R],
	initial R,
) R {
	reduced := initial

	for index, item := range slice {
		reduced = reducer(reduced, item, index)
	}

	return reduced
}

type reduceFn[T any, R any] func(reduced R, each T, index int) R
