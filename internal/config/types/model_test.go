package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModel_Merge(t *testing.T) {
	t.Parallel()

	type fields struct {
		Schemas []string
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
			name: "nil",
			fields: fields{
				Schemas: []string{"schema"},
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
			name: "only schemas",
			fields: fields{
				Schemas: []string{"schema"},
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
			name: "only include",
			fields: fields{
				Schemas: []string{"schema"},
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
			name: "only exclude",
			fields: fields{
				Schemas: []string{"schema"},
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
			name: "all",
			fields: fields{
				Schemas: []string{"schema"},
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
				Schemas: testCase.fields.Schemas,
				Include: testCase.fields.Include,
				Exclude: testCase.fields.Exclude,
			}

			assert.Equal(t, testCase.want, m.Merge(testCase.args.other))
		})
	}
}

func TestModel_String(t *testing.T) {
	t.Parallel()

	type fields struct {
		Schemas []string
		Include []string
		Exclude []string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "nil",
			fields: fields{
				Schemas: nil,
				Include: nil,
				Exclude: nil,
			},
			want: "Model{schemas: [], include: [], exclude: []}",
		},
		{
			name: "empty",
			fields: fields{
				Schemas: []string{},
				Include: []string{},
				Exclude: []string{},
			},
			want: "Model{schemas: [], include: [], exclude: []}",
		},
		{
			name: "non-empty",
			fields: fields{
				Schemas: []string{"schema"},
				Include: []string{"include"},
				Exclude: []string{"exclude"},
			},
			want: "Model{schemas: [schema], include: [include], exclude: [exclude]}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Model{
				Schemas: tt.fields.Schemas,
				Include: tt.fields.Include,
				Exclude: tt.fields.Exclude,
			}

			assert.Equal(t, tt.want, m.String())
		})
	}
}
