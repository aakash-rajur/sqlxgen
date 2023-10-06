package casing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToTileCase(t *testing.T) {
	t.Parallel()

	type args struct {
		word string
	}

	testCases := []struct {
		name string
		args args
		want string
	}{
		{
			name: "snake case",
			args: args{
				word: "snake_case",
			},
			want: "Snake_case",
		},
		{
			name: "kebab case",
			args: args{
				word: "kebab-case",
			},
			want: "Kebab-case",
		},
		{
			name: "pascal case",
			args: args{
				word: "PascalCase",
			},
			want: "PascalCase",
		},
		{
			name: "camel case",
			args: args{
				word: "camelCase",
			},
			want: "CamelCase",
		},
		{
			name: "space case",
			args: args{
				word: "space case",
			},
			want: "Space case",
		},
		{
			name: "empty string",
			args: args{
				word: "",
			},
			want: "",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := toTitleCase(testCase.args.word)

			assert.Equal(t, testCase.want, got)
		})
	}
}
