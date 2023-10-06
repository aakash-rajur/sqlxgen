package array

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	t.Parallel()

	type args[T any, R any] struct {
		arr      []int
		callback func(v T, i int) R
	}

	type TestCase[T any, R any] struct {
		name string
		args args[T, R]
		want []R
	}

	t.Run("map int to string", func(t *testing.T) {
		testCases := []TestCase[int, string]{
			{
				name: "map int to string",
				args: args[int, string]{
					arr: []int{1, 2, 3, 4, 5, 6},
					callback: func(v int, i int) string {
						return fmt.Sprintf("%d", v)
					},
				},
				want: []string{"1", "2", "3", "4", "5", "6"},
			},
			{
				name: "map int to string",
				args: args[int, string]{
					arr: []int{10, 20, 30, 40, 50, 60},
					callback: func(v int, i int) string {
						return fmt.Sprintf("%d", v)
					},
				},
				want: []string{"10", "20", "30", "40", "50", "60"},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				got := Map(testCase.args.arr, testCase.args.callback)

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

	t.Run("map int to int", func(t *testing.T) {
		testCases := []TestCase[int, int]{
			{
				name: "map int to int",
				args: args[int, int]{
					arr: []int{1, 2, 3, 4, 5, 6},
					callback: func(v int, i int) int {
						return v * 2
					},
				},
				want: []int{2, 4, 6, 8, 10, 12},
			},
			{
				name: "map int to int",
				args: args[int, int]{
					arr: []int{10, 20, 30, 40, 50, 60},
					callback: func(v int, i int) int {
						return v * 2
					},
				},
				want: []int{20, 40, 60, 80, 100, 120},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				got := Map(testCase.args.arr, testCase.args.callback)

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
