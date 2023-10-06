package array

func ForEach[T any](
	slice []T,
	fn ForEachFn[T],
) {
	for index, item := range slice {
		fn(&item, index)
	}
}

type ForEachFn[T any] func(each *T, index int)
