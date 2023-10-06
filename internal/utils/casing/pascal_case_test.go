package casing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPascalCase(t *testing.T) {
	t.Parallel()

	type args struct {
		identifier string
	}

	testCases := []struct {
		name  string
		args  args
		want  string
		error error
	}{
		{
			name: "snake case",
			args: args{
				identifier: "snake_case",
			},
			want:  "SnakeCase",
			error: nil,
		},
		{
			name: "kebab case",
			args: args{
				identifier: "kebab-case",
			},
			want:  "KebabCase",
			error: nil,
		},
		{
			name: "pascal case",
			args: args{
				identifier: "PascalCase",
			},
			want:  "PascalCase",
			error: nil,
		},
		{
			name: "camel case",
			args: args{
				identifier: "camelCase",
			},
			want:  "CamelCase",
			error: nil,
		},
		{
			name: "space case",
			args: args{
				identifier: "space case",
			},
			want:  "SpaceCase",
			error: nil,
		},
		{
			name: "empty string",
			args: args{
				identifier: "",
			},
			want:  "",
			error: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := PascalCase(testCase.args.identifier)

			if testCase.error != nil {
				errMsgLeft := testCase.error.Error()

				errMsgRight := err.Error()

				assert.Contains(t, errMsgRight, errMsgLeft)
			} else {
				assert.Nil(t, err)

				assert.Equal(t, testCase.want, got)
			}
		})
	}
}
