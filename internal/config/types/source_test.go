package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSource_Merge(t *testing.T) {
	type fields struct {
		Models  *Model
		Queries *Query
	}

	type args struct {
		other *Source
	}

	testCases := []struct {
		name   string
		fields fields
		args   args
		want   *Source
	}{
		{
			name: "source nil",
			fields: fields{
				Models:  nil,
				Queries: nil,
			},
			args: args{
				other: nil,
			},
			want: &Source{
				Models:  nil,
				Queries: nil,
			},
		},
		{
			name: "source only models",
			fields: fields{
				Models: &Model{
					Schemas: []string{"schema"},
					Include: []string{"include"},
					Exclude: []string{"exclude"},
				},
				Queries: nil,
			},
			args: args{
				other: &Source{
					Models:  nil,
					Queries: nil,
				},
			},
			want: &Source{
				Models: &Model{
					Schemas: []string{"schema"},
					Include: []string{"include"},
					Exclude: []string{"exclude"},
				},
				Queries: nil,
			},
		},
		{
			name: "source only queries",
			fields: fields{
				Models: nil,
				Queries: &Query{
					Paths:   []string{"queries"},
					Include: []string{"include"},
					Exclude: []string{"exclude"},
				},
			},
			args: args{
				other: &Source{
					Models:  nil,
					Queries: nil,
				},
			},
			want: &Source{
				Models: nil,
				Queries: &Query{
					Paths:   []string{"queries"},
					Include: []string{"include"},
					Exclude: []string{"exclude"},
				},
			},
		},
		{
			name: "source models and queries",
			fields: fields{
				Models: &Model{
					Schemas: []string{"schema"},
					Include: []string{"include"},
					Exclude: []string{"exclude"},
				},
				Queries: &Query{
					Paths:   []string{"queries"},
					Include: []string{"include"},
					Exclude: []string{"exclude"},
				},
			},
			args: args{
				other: &Source{
					Models:  nil,
					Queries: nil,
				},
			},
			want: &Source{
				Models: &Model{
					Schemas: []string{"schema"},
					Include: []string{"include"},
					Exclude: []string{"exclude"},
				},
				Queries: &Query{
					Paths:   []string{"queries"},
					Include: []string{"include"},
					Exclude: []string{"exclude"},
				},
			},
		},
		{
			name: "other only models",
			fields: fields{
				Models: &Model{
					Schemas: []string{"schema"},
					Include: []string{"include"},
					Exclude: []string{"exclude"},
				},
				Queries: &Query{
					Paths:   []string{"queries"},
					Include: []string{"include"},
					Exclude: []string{"exclude"},
				},
			},
			args: args{
				other: &Source{
					Models: &Model{
						Schemas: []string{"other"},
						Include: []string{"other"},
						Exclude: []string{"other"},
					},
					Queries: nil,
				},
			},
			want: &Source{
				Models: &Model{
					Schemas: []string{"other"},
					Include: []string{"other"},
					Exclude: []string{"other"},
				},
				Queries: &Query{
					Paths:   []string{"queries"},
					Include: []string{"include"},
					Exclude: []string{"exclude"},
				},
			},
		},
		{
			name: "other only queries",
			fields: fields{
				Models: &Model{
					Schemas: []string{"schema"},
					Include: []string{"include"},
					Exclude: []string{"exclude"},
				},
				Queries: &Query{
					Paths:   []string{"queries"},
					Include: []string{"include"},
					Exclude: []string{"exclude"},
				},
			},
			args: args{
				other: &Source{
					Models: nil,
					Queries: &Query{
						Paths:   []string{"other"},
						Include: []string{"other"},
						Exclude: []string{"other"},
					},
				},
			},
			want: &Source{
				Models: &Model{
					Schemas: []string{"schema"},
					Include: []string{"include"},
					Exclude: []string{"exclude"},
				},
				Queries: &Query{
					Paths:   []string{"other"},
					Include: []string{"other"},
					Exclude: []string{"other"},
				},
			},
		},
		{
			name: "other models and queries",
			fields: fields{
				Models: &Model{
					Schemas: []string{"schema"},
					Include: []string{"include"},
					Exclude: []string{"exclude"},
				},
			},
			args: args{
				other: &Source{
					Models: &Model{
						Schemas: []string{"other"},
						Include: []string{"other"},
						Exclude: []string{"other"},
					},
					Queries: &Query{
						Paths:   []string{"other"},
						Include: []string{"other"},
						Exclude: []string{"other"},
					},
				},
			},
			want: &Source{
				Models: &Model{
					Schemas: []string{"other"},
					Include: []string{"other"},
					Exclude: []string{"other"},
				},
				Queries: &Query{
					Paths:   []string{"other"},
					Include: []string{"other"},
					Exclude: []string{"other"},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := &Source{
				Models:  testCase.fields.Models,
				Queries: testCase.fields.Queries,
			}

			got := s.Merge(testCase.args.other)

			assert.Equal(t, testCase.want, got)
		})
	}
}

func TestSource_String(t *testing.T) {
	type fields struct {
		Models  *Model
		Queries *Query
	}

	testCases := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "source nil",
			fields: fields{
				Models:  nil,
				Queries: nil,
			},
			want: "Source{models: Model{nil}, queries: Query{nil}}",
		},
		{
			name: "source only models",
			fields: fields{
				Models: &Model{
					Schemas: []string{"schema"},
					Include: []string{"include"},
					Exclude: []string{"exclude"},
				},
				Queries: nil,
			},
			want: "Source{models: Model{schemas: [schema], include: [include], exclude: [exclude]}, queries: Query{nil}}",
		},
		{
			name: "source only queries",
			fields: fields{
				Models: nil,
				Queries: &Query{
					Paths:   []string{"queries"},
					Include: []string{"include"},
					Exclude: []string{"exclude"},
				},
			},
			want: "Source{models: Model{nil}, queries: Query{paths: [queries], include: [include], exclude: [exclude]}}",
		},
		{
			name: "source models and queries",
			fields: fields{
				Models: &Model{
					Schemas: []string{"schema"},
					Include: []string{"include"},
					Exclude: []string{"exclude"},
				},
				Queries: &Query{
					Paths:   []string{"queries"},
					Include: []string{"include"},
					Exclude: []string{"exclude"},
				},
			},
			want: "Source{models: Model{schemas: [schema], include: [include], exclude: [exclude]}, queries: Query{paths: [queries], include: [include], exclude: [exclude]}}",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := &Source{
				Models:  testCase.fields.Models,
				Queries: testCase.fields.Queries,
			}

			got := s.String()

			assert.Equal(t, testCase.want, got)
		})
	}
}
