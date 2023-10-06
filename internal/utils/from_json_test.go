package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromJson(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		jsons []string
		want  []map[string]interface{}
		err   error
	}{
		{
			name:  "empty jsons",
			jsons: []string{},
			want:  []map[string]interface{}{},
			err:   nil,
		},
		{
			name:  "invalid jsons",
			jsons: []string{"{", "}", "foo"},
			want:  nil,
			err:   errors.New("unexpected end of JSON input"),
		},
		{
			name:  "valid jsons",
			jsons: []string{`{"foo": "bar"}`, `{"bar": "baz"}`},
			want: []map[string]interface{}{
				{"foo": "bar"},
				{"bar": "baz"},
			},
			err: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := FromJson[map[string]interface{}](testCase.jsons)

			if testCase.err != nil {
				assert.EqualError(t, err, testCase.err.Error())
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, testCase.want, got)
		})
	}
}
