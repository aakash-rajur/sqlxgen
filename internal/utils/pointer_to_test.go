package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointerTo(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		testName string
		value    interface{}
	}{
		{
			testName: "nil",
			value:    nil,
		},
		{
			testName: "int",
			value:    42,
		},
		{
			testName: "string",
			value:    "foo",
		},
		{
			testName: "struct",
			value:    struct{}{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ptr := PointerTo(tc.value)

			assert.Equal(t, tc.value, *ptr)
		})
	}
}
