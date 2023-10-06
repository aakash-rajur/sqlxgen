package prepare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleEnsureWhereClause(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input string
		want  string
		err   error
	}{
		{
			name:  "no where clause 1",
			input: `select * from users;`,
			want: `select * from users
where false
;`,
			err: nil,
		},
		{
			name: "no where clause 2",
			input: `select * from users
limit :limit -- :limit type: int
offset :offset; -- :offset type: int`,
			want: `select * from users
where false

limit :limit -- :limit type: int
offset :offset; -- :offset type: int`,
			err: nil,
		},
		{
			name: "with where clause",
			input: `select * from users where true
and id = :id -- :id type: int
limit :limit -- :limit type: int
offset :offset; -- :offset type: int`,
			want: `select * from users where true
and id = :id -- :id type: int
limit :limit -- :limit type: int
offset :offset; -- :offset type: int`,
			err: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := handleEnsureWhereClause(testCase.input)

			assert.Nil(t, err)

			assert.Equal(t, testCase.want, got)
		})
	}
}
