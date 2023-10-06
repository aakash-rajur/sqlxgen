package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Parallel()

	type args[T comparable] struct {
		arr      []T
		callback func(v T, i int) bool
	}

	type TestCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}

	t.Run("filter int", func(t *testing.T) {
		t.Parallel()

		testCases := []TestCase[int]{
			{
				name: "filter even numbers",
				args: args[int]{
					arr: []int{1, 2, 3, 4, 5, 6},
					callback: func(v int, i int) bool {
						return v%2 == 0
					},
				},
				want: []int{2, 4, 6},
			},
			{
				name: "filter odd numbers",
				args: args[int]{
					arr: []int{1, 2, 3, 4, 5, 6},
					callback: func(v int, i int) bool {
						return v%2 == 1
					},
				},
				want: []int{1, 3, 5},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				got := Filter(testCase.args.arr, testCase.args.callback)

				assert.Equal(t, testCase.want, got)
			})
		}
	})

	t.Run("filter string", func(t *testing.T) {
		t.Parallel()

		testCases := []TestCase[string]{
			{
				name: "filter even characters",
				args: args[string]{
					arr: []string{"a", "b", "c", "d", "e", "f"},
					callback: func(v string, i int) bool {
						return i%2 == 0
					},
				},
				want: []string{"a", "c", "e"},
			},
			{
				name: "filter odd characters",
				args: args[string]{
					arr: []string{"a", "b", "c", "d", "e", "f"},
					callback: func(v string, i int) bool {
						return i%2 == 1
					},
				},
				want: []string{"b", "d", "f"},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				got := Filter(testCase.args.arr, testCase.args.callback)

				if len(got) != len(testCase.want) {
					t.Errorf("want %v, got %v", testCase.want, got)
				}

				for i := 0; i < len(got); i++ {
					if got[i] != testCase.want[i] {
						t.Errorf("want %v, got %v", testCase.want, got)
					}
				}
			})
		}
	})
}
