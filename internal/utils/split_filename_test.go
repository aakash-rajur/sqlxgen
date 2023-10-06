package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitFilename(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		testName string
		filename string
		base     string
		ext      string
	}{
		{
			testName: "empty string",
			filename: "",
			base:     "",
			ext:      "",
		},
		{
			testName: "no extension",
			filename: "foo",
			base:     "foo",
			ext:      "",
		},
		{
			testName: "one dot with no extension",
			filename: "foo.",
			base:     "foo",
			ext:      "",
		},
		{
			testName: "one dot with extension",
			filename: "foo.bar",
			base:     "foo",
			ext:      "bar",
		},
		{
			testName: "two dots with extension",
			filename: "foo.bar.baz",
			base:     "foo.bar",
			ext:      "baz",
		},
		{
			testName: "leading dot",
			filename: ".foo",
			base:     "",
			ext:      "foo",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			base, ext := SplitFilename(tc.filename)

			assert.Equal(t, tc.base, base)

			assert.Equal(t, tc.ext, ext)
		})
	}
}
