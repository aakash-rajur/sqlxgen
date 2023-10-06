package utils

import (
	"cmp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValues(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		testName string
		input    map[string]interface{}
		want     []interface{}
		cmp      func(a, b any) int
	}{
		{
			testName: "empty map",
			input:    map[string]interface{}{},
			want:     []interface{}{},
			cmp:      nil,
		},
		{
			testName: "map with one string value",
			input:    map[string]interface{}{"key": "value"},
			want:     []interface{}{"value"},
			cmp: func(a, b any) int {
				return cmp.Compare(a.(string), b.(string))
			},
		},
		{
			testName: "map with multiple string values",
			input:    map[string]interface{}{"key1": "value1", "key2": "value2"},
			want:     []interface{}{"value1", "value2"},
			cmp: func(a, b any) int {
				return cmp.Compare(a.(string), b.(string))
			},
		},
		{
			testName: "map with one int value",
			input:    map[string]interface{}{"key": 1},
			want:     []interface{}{1},
			cmp: func(a, b any) int {
				return cmp.Compare(a.(int), b.(int))
			},
		},
		{
			testName: "map with multiple int values",
			input:    map[string]interface{}{"key1": 1, "key2": 2},
			want:     []interface{}{1, 2},
			cmp: func(a, b any) int {
				return cmp.Compare(a.(int), b.(int))
			},
		},
		{
			testName: "map with one float value",
			input:    map[string]interface{}{"key": 1.1},
			want:     []interface{}{1.1},
			cmp: func(a, b any) int {
				return cmp.Compare(a.(float64), b.(float64))
			},
		},
		{
			testName: "map with multiple float values",
			input:    map[string]interface{}{"key1": 1.1, "key2": 2.2},
			want:     []interface{}{1.1, 2.2},
			cmp: func(a, b any) int {
				return cmp.Compare(a.(float64), b.(float64))
			},
		},
		{
			testName: "map with one bool value",
			input:    map[string]interface{}{"key": true},
			want:     []interface{}{true},
			cmp:      nil,
		},
		{
			testName: "map with multiple bool values",
			input:    map[string]interface{}{"key1": true, "key2": false},
			want:     []interface{}{true, false},
			cmp: func(a, b any) int {
				ab := a.(bool)

				bb := b.(bool)

				ai, bi := 0, 0

				if !ab {
					ai = 1
				}

				if !bb {
					bi = 1
				}

				return ai - bi
			},
		},
		{
			testName: "map with one struct value",
			input:    map[string]interface{}{"key": struct{}{}},
			want:     []interface{}{struct{}{}},
			cmp:      nil,
		},
		{
			testName: "map with multiple struct values",
			input:    map[string]interface{}{"key1": struct{}{}, "key2": struct{}{}},
			want:     []interface{}{struct{}{}, struct{}{}},
			cmp:      nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			actual := Values(testCase.input)

			assert.ElementsMatch(t, testCase.want, actual, "want values do not match actual values")
		})
	}
}

func TestKeys(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		testName string
		input    map[string]interface{}
		want     []string
	}{
		{
			testName: "empty map",
			input:    map[string]interface{}{},
			want:     []string{},
		},
		{
			testName: "map with one key",
			input:    map[string]interface{}{"key": "value"},
			want:     []string{"key"},
		},
		{
			testName: "map with multiple keys",
			input:    map[string]interface{}{"key1": "value1", "key2": "value2"},
			want:     []string{"key1", "key2"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			actual := Keys(testCase.input)

			assert.ElementsMatch(t, testCase.want, actual, "want keys do not match actual keys")
		})
	}
}

func TestEntries(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		testName string
		input    map[string]interface{}
		want     []Entry[string, interface{}]
	}{
		{
			testName: "empty map",
			input:    map[string]interface{}{},
			want:     make([]Entry[string, interface{}], 0),
		},
		{
			testName: "map with one entry",
			input:    map[string]interface{}{"key": "value"},
			want:     []Entry[string, interface{}]{{"key", "value"}},
		},
		{
			testName: "map with multiple entries",
			input:    map[string]interface{}{"key1": "value1", "key2": "value2"},
			want:     []Entry[string, interface{}]{{"key1", "value1"}, {"key2", "value2"}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			actual := Entries(testCase.input)

			assert.ElementsMatch(t, testCase.want, actual, "want entries do not match actual entries")
		})
	}
}
