package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForEach(t *testing.T) {
	t.Parallel()

	type args[T any] struct {
		arr      []T
		callback func(v T, i int)
	}

	type TestCase[T any] struct {
		name string
		args args[T]
	}

	t.Run("for each int", func(t *testing.T) {
		got := make([]int, 0)

		ForEach(
			[]int{1, 2, 3, 4, 5, 6},
			func(v *int, i int) {
				got = append(got, 2*(*v))
			},
		)

		want := []int{2, 4, 6, 8, 10, 12}

		assert.Equal(t, want, got)
	})

	t.Run("for each string", func(t *testing.T) {
		got := make([]string, 0)

		ForEach(
			[]string{"a", "b", "c", "d", "e", "f"},
			func(v *string, i int) {
				got = append(got, *v+(*v))
			},
		)

		want := []string{"aa", "bb", "cc", "dd", "ee", "ff"}

		assert.Equal(t, want, got)
	})
}
