package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilenameWithGen(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		filename string
		want     string
	}{
		{
			name:     "empty filename",
			filename: "",
			want:     ".gen.",
		},
		{
			name:     "no extension",
			filename: "filename",
			want:     "filename.gen.",
		},
		{
			name:     "with extension",
			filename: "filename.txt",
			want:     "filename.gen.txt",
		},
		{
			name:     "with multiple extensions",
			filename: "filename.txt.gz",
			want:     "filename.txt.gen.gz",
		},
		{
			name:     "with multiple dots",
			filename: "filename.with.dots.txt",
			want:     "filename.with.dots.gen.txt",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := FilenameWithGen(testCase.filename)

			assert.Equal(t, testCase.want, got, "want filename with gen does not match actual filename with gen")
		})
	}
}
