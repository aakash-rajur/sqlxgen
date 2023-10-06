package array

import (
	"strconv"
	"testing"
)

func TestReduce(t *testing.T) {
	t.Parallel()

	t.Run("reduce sum", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}

		reduced := Reduce(
			slice,
			func(reduced int, each int, index int) int {
				return reduced + each
			},
			0,
		)

		if reduced != 15 {
			t.Errorf("want %d, got %d", 15, reduced)
		}
	})

	t.Run("reduce product", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}

		reduced := Reduce(
			slice,
			func(reduced int, each int, index int) int {
				return reduced * each
			},
			1,
		)

		if reduced != 120 {
			t.Errorf("want %d, got %d", 120, reduced)
		}
	})

	t.Run("reduce string merge", func(t *testing.T) {
		slice := []string{"a", "b", "c", "d", "e"}

		reduced := Reduce(
			slice,
			func(reduced string, each string, index int) string {
				return reduced + each
			},
			"",
		)

		if reduced != "abcde" {
			t.Errorf("want %s, got %s", "abcde", reduced)
		}
	})

	t.Run("reduce string merge with index", func(t *testing.T) {
		slice := []string{"a", "b", "c", "d", "e"}

		reduced := Reduce(
			slice,
			func(reduced string, each string, index int) string {
				return reduced + each + strconv.Itoa(index)
			},
			"",
		)

		if reduced != "a0b1c2d3e4" {
			t.Errorf("want %s, got %s", "a0b1c2d3e4", reduced)
		}
	})
}
