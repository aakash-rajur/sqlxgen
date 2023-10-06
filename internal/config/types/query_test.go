package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_Merge(t *testing.T) {
	t.Parallel()

	type fields struct {
		Paths   []string
		Include []string
		Exclude []string
	}

	type args struct {
		other *Model
	}

	testCases := []struct {
		name   string
		fields fields
		args   args
		want   *Model
	}{
		{
			name: "other nil",
			fields: fields{
				Paths:   []string{"schema"},
				Include: []string{"include"},
				Exclude: []string{"exclude"},
			},
			args: args{
				other: nil,
			},
			want: &Model{
				Schemas: []string{"schema"},
				Include: []string{"include"},
				Exclude: []string{"exclude"},
			},
		},
		{
			name: "other paths",
			fields: fields{
				Paths:   []string{"schema"},
				Include: []string{"include"},
				Exclude: []string{"exclude"},
			},
			args: args{
				other: &Model{
					Schemas: []string{"other"},
				},
			},
			want: &Model{
				Schemas: []string{"other"},
				Include: []string{"include"},
				Exclude: []string{"exclude"},
			},
		},
		{
			name: "other include",
			fields: fields{
				Paths:   []string{"schema"},
				Include: []string{"include"},
				Exclude: []string{"exclude"},
			},
			args: args{
				other: &Model{
					Include: []string{"other"},
				},
			},
			want: &Model{
				Schemas: []string{"schema"},
				Include: []string{"other"},
				Exclude: []string{"exclude"},
			},
		},
		{
			name: "other exclude",
			fields: fields{
				Paths:   []string{"schema"},
				Include: []string{"include"},
				Exclude: []string{"exclude"},
			},
			args: args{
				other: &Model{
					Exclude: []string{"other"},
				},
			},
			want: &Model{
				Schemas: []string{"schema"},
				Include: []string{"include"},
				Exclude: []string{"other"},
			},
		},
		{
			name: "other paths, include, exclude",
			fields: fields{
				Paths:   []string{"schema"},
				Include: []string{"include"},
				Exclude: []string{"exclude"},
			},
			args: args{
				other: &Model{
					Schemas: []string{"other"},
					Include: []string{"other"},
					Exclude: []string{"other"},
				},
			},
			want: &Model{
				Schemas: []string{"other"},
				Include: []string{"other"},
				Exclude: []string{"other"},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			m := &Model{
				Schemas: testCase.fields.Paths,
				Include: testCase.fields.Include,
				Exclude: testCase.fields.Exclude,
			}

			assert.Equal(t, testCase.want, m.Merge(testCase.args.other))
		})
	}
}

func TestQuery_String(t *testing.T) {
	t.Parallel()

	type fields struct {
		Paths   []string
		Include []string
		Exclude []string
	}

	testCases := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "nil",
			fields: fields{
				Paths:   nil,
				Include: nil,
				Exclude: nil,
			},
			want: "Query{paths: [], include: [], exclude: []}",
		},
		{
			name: "empty",
			fields: fields{
				Paths:   []string{},
				Include: []string{},
				Exclude: []string{},
			},
			want: "Query{paths: [], include: [], exclude: []}",
		},
		{
			name: "all",
			fields: fields{
				Paths:   []string{"schema"},
				Include: []string{"include"},
				Exclude: []string{"exclude"},
			},
			want: "Query{paths: [schema], include: [include], exclude: [exclude]}",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			q := &Query{
				Paths:   testCase.fields.Paths,
				Include: testCase.fields.Include,
				Exclude: testCase.fields.Exclude,
			}

			assert.Equal(t, testCase.want, q.String())
		})
	}
}
