package utils

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0)

	for _, value := range m {
		values = append(values, value)
	}

	return values
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0)

	for key := range m {
		keys = append(keys, key)
	}

	return keys
}

func Entries[K comparable, V any](m map[K]V) []Entry[K, V] {
	entries := make([]Entry[K, V], 0)

	for key, value := range m {
		entry := Entry[K, V]{key, value}

		entries = append(entries, entry)
	}

	return entries
}

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}
