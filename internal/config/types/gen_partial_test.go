package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenPartial_Merge(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		g     *GenPartial
		other *GenPartial
		want  *GenPartial
	}{
		{
			name:  "nil",
			g:     nil,
			other: nil,
			want:  nil,
		},
		{
			name:  "nil g",
			g:     nil,
			other: &GenPartial{},
			want:  &GenPartial{},
		},
		{
			name:  "nil other",
			g:     &GenPartial{},
			other: nil,
			want:  &GenPartial{},
		},
		{
			name:  "empty",
			g:     &GenPartial{},
			other: &GenPartial{},
			want:  &GenPartial{},
		},
		{
			name:  "path",
			g:     &GenPartial{},
			other: &GenPartial{Path: "path"},
			want:  &GenPartial{Path: "path"},
		},
		{
			name:  "path overwrite",
			g:     &GenPartial{Path: "path"},
			other: &GenPartial{Path: "path2"},
			want:  &GenPartial{Path: "path2"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.g.Merge(testCase.other)

			assert.Equal(t, testCase.want, got)
		})
	}
}

func TestGenPartial_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		g    *GenPartial
		want string
	}{
		{
			name: "empty",
			g:    &GenPartial{},
			want: "GenPartial{path: }",
		},
		{
			name: "path",
			g:    &GenPartial{Path: "path"},
			want: "GenPartial{path: path}",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.g.String()

			assert.Equal(t, testCase.want, got)
		})
	}
}
