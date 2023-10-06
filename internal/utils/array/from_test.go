package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrom(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		t.Parallel()

		type args struct {
			length int
			fn     fromFn[int]
		}

		testCases := []struct {
			name string
			args args
			want []int
		}{
			{
				name: "empty array",
				args: args{
					length: 0,
					fn:     func(offset int) int { return offset },
				},
				want: []int{},
			},
			{
				name: "array of length 1",
				args: args{
					length: 1,
					fn:     func(offset int) int { return offset },
				},
				want: []int{0},
			},
			{
				name: "array of length 10",
				args: args{
					length: 10,
					fn:     func(offset int) int { return offset },
				},
				want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				got := From(testCase.args.length, testCase.args.fn)

				assert.Equal(t, testCase.want, got)
			})
		}
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()

		type args struct {
			length int
			fn     fromFn[string]
		}

		testCases := []struct {
			name string
			args args
			want []string
		}{
			{
				name: "empty array",
				args: args{
					length: 0,
					fn:     func(offset int) string { return "a" },
				},
				want: []string{},
			},
			{
				name: "array of length 1",
				args: args{
					length: 1,
					fn:     func(offset int) string { return "a" },
				},
				want: []string{"a"},
			},
			{
				name: "array of length 10",
				args: args{
					length: 10,
					fn:     func(offset int) string { return "a" },
				},
				want: []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a"},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				got := From(testCase.args.length, testCase.args.fn)

				assert.Equal(t, testCase.want, got)
			})
		}
	})

	t.Run("bool", func(t *testing.T) {
		t.Parallel()

		type args struct {
			length int
			fn     fromFn[bool]
		}

		testCases := []struct {
			name string
			args args
			want []bool
		}{
			{
				name: "empty array",
				args: args{
					length: 0,
					fn:     func(offset int) bool { return true },
				},
				want: []bool{},
			},
			{
				name: "array of length 1",
				args: args{
					length: 1,
					fn:     func(offset int) bool { return true },
				},
				want: []bool{true},
			},
			{
				name: "array of length 10",
				args: args{
					length: 10,
					fn:     func(offset int) bool { return true },
				},
				want: []bool{true, true, true, true, true, true, true, true, true, true},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				got := From(testCase.args.length, testCase.args.fn)

				assert.Equal(t, testCase.want, got)
			})
		}
	})

	t.Run("float64", func(t *testing.T) {
		t.Parallel()

		type args struct {
			length int
			fn     fromFn[float64]
		}

		testCases := []struct {
			name string
			args args
			want []float64
		}{
			{
				name: "empty array",
				args: args{
					length: 0,
					fn:     func(offset int) float64 { return float64(offset) },
				},
				want: []float64{},
			},
			{
				name: "array of length 1",
				args: args{
					length: 1,
					fn:     func(offset int) float64 { return float64(offset) },
				},
				want: []float64{0},
			},
			{
				name: "array of length 10",
				args: args{
					length: 10,
					fn:     func(offset int) float64 { return float64(offset) },
				},
				want: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				got := From(testCase.args.length, testCase.args.fn)

				assert.Equal(t, testCase.want, got)
			})
		}
	})
}
