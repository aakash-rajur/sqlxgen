package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGen_Merge(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		g     *Gen
		other *Gen
		want  *Gen
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
			other: &Gen{},
			want:  &Gen{},
		},
		{
			name:  "nil other",
			g:     &Gen{},
			other: nil,
			want:  &Gen{},
		},
		{
			name:  "empty",
			g:     &Gen{},
			other: &Gen{},
			want:  &Gen{},
		},
		{
			name:  "store",
			g:     &Gen{},
			other: &Gen{Store: &GenPartial{Path: "path"}},
			want:  &Gen{Store: &GenPartial{Path: "path"}},
		},
		{
			name:  "store overwrite",
			g:     &Gen{Store: &GenPartial{Path: "path"}},
			other: &Gen{Store: &GenPartial{Path: "path2"}},
			want:  &Gen{Store: &GenPartial{Path: "path2"}},
		},
		{
			name:  "model",
			g:     &Gen{},
			other: &Gen{Model: &GenPartial{Path: "path"}},
			want:  &Gen{Model: &GenPartial{Path: "path"}},
		},
		{
			name:  "model overwrite",
			g:     &Gen{Model: &GenPartial{Path: "path"}},
			other: &Gen{Model: &GenPartial{Path: "path2"}},
			want:  &Gen{Model: &GenPartial{Path: "path2"}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.g.Merge(testCase.other)

			assert.Equal(t, testCase.want, got)
		})
	}
}

func TestGen_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		g    *Gen
		want string
	}{
		{
			name: "nil",
			g:    nil,
			want: "Gen{nil}",
		},
		{
			name: "empty",
			g:    &Gen{},
			want: "Gen{store: GenPartial{nil}, model: GenPartial{nil}}",
		},
		{
			name: "store",
			g:    &Gen{Store: &GenPartial{Path: "path"}},
			want: `Gen{store: GenPartial{path: path}, model: GenPartial{nil}}`,
		},
		{
			name: "model",
			g:    &Gen{Model: &GenPartial{Path: "path"}},
			want: `Gen{store: GenPartial{nil}, model: GenPartial{path: path}}`,
		},
		{
			name: "store and model",
			g:    &Gen{Store: &GenPartial{Path: "path"}, Model: &GenPartial{Path: "path"}},
			want: `Gen{store: GenPartial{path: path}, model: GenPartial{path: path}}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.g.String()

			assert.Equal(t, testCase.want, got)
		})
	}
}
