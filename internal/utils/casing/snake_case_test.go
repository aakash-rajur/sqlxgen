package casing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnakeCase(t *testing.T) {
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
			want:  "snake_case",
			error: nil,
		},
		{
			name: "kebab case",
			args: args{
				identifier: "kebab-case",
			},
			want:  "kebab_case",
			error: nil,
		},
		{
			name: "pascal case",
			args: args{
				identifier: "PascalCase",
			},
			want:  "pascal_case",
			error: nil,
		},
		{
			name: "camel case",
			args: args{
				identifier: "camelCase",
			},
			want:  "camel_case",
			error: nil,
		},
		{
			name: "space case",
			args: args{
				identifier: "space case",
			},
			want:  "space_case",
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
			got, err := SnakeCase(testCase.args.identifier)

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
