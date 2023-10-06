package utils

import (
	"regexp"
	"testing"

	"github.com/aakash-rajur/sqlxgen/internal/utils/array"
	"github.com/joomcode/errorx"
	"github.com/stretchr/testify/assert"
)

func TestCompileRegexes(t *testing.T) {
	t.Parallel()

	type args struct {
		patterns []string
	}

	testCases := []struct {
		name string
		args args
		err  error
		want []*regexp.Regexp
	}{
		{
			name: "empty",
			args: args{
				patterns: []string{},
			},
			err:  nil,
			want: []*regexp.Regexp{},
		},
		{
			name: "invalid",
			args: args{
				patterns: []string{
					"[",
				},
			},
			err:  errorx.IllegalFormat.New("failed to compile regex '['"),
			want: nil,
		},
		{
			name: "valid 1",
			args: args{
				patterns: []string{
					"^[a-z]+$",
				},
			},
			err: nil,
			want: []*regexp.Regexp{
				regexp.MustCompile("^[a-z]+$"),
			},
		},
		{
			name: "valid 2",
			args: args{
				patterns: []string{
					"^.+$",
					"^public.migrations$",
					"^public.entity_status_logs_vof_v2_backup$",
					"^list-project-2.sql$",
				},
			},
			err: nil,
			want: []*regexp.Regexp{
				regexp.MustCompile("^.+$"),
				regexp.MustCompile("^public.migrations$"),
				regexp.MustCompile("^public.entity_status_logs_vof_v2_backup$"),
				regexp.MustCompile("^list-project-2.sql$"),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			regexes, err := compileRegexes(testCase.args.patterns)

			if testCase.err != nil {
				errMsgLeft := testCase.err.Error()

				errMsgRight := err.Error()

				assert.Containsf(t, errMsgRight, errMsgLeft, "want error %s but got %s", errMsgLeft, errMsgRight)
			} else {
				assert.Nil(t, err)

				assert.Equal(t, testCase.want, regexes)
			}
		})
	}
}

func TestCreateValidateEntityNames(t *testing.T) {
	t.Parallel()

	tableNames := []string{
		"actors",
		"companies",
		"crew",
		"hyper_parameter",
		"movies",
		"movies_actors",
		"movies_companies",
		"movies_countries",
		"movies_crew",
		"movies_genres",
		"movies_languages",
		"t_movies",
		"t_movies_credits",
	}

	queryNames := []string{
		"list-hyper-parameters.sql",
		"list-movies.sql",
		"get-hyper-parameter.sql",
		"get-movie.sql",
		"get-company.sql",
		"get-actor.sql",
		"get-crew.sql",
		"example1.sql",
	}

	type args struct {
		inclusions []string
		exclusions []string
	}

	testCases := []struct {
		name        string
		args        args
		err         error
		entityNames []string
		want        []bool
	}{
		{
			name: "all table names",
			args: args{
				inclusions: []string{},
				exclusions: []string{},
			},
			err:         nil,
			entityNames: tableNames,
			want:        array.From(len(tableNames), func(_ int) bool { return true }),
		},
		{
			name: "all query names",
			args: args{
				inclusions: []string{},
				exclusions: []string{},
			},
			err:         nil,
			entityNames: queryNames,
			want:        array.From(len(queryNames), func(_ int) bool { return true }),
		},
		{
			name: "exclude some table names",
			args: args{
				inclusions: []string{},
				exclusions: []string{"^t_movies$", "^t_movies_credits$"},
			},
			err:         nil,
			entityNames: tableNames,
			want:        array.From(len(tableNames), func(offset int) bool { return offset < 11 }),
		},
		{
			name: "exclude some query names",
			args: args{
				inclusions: []string{},
				exclusions: []string{"^example1.sql$"},
			},
			err:         nil,
			entityNames: queryNames,
			want:        array.From(len(queryNames), func(offset int) bool { return offset < 7 }),
		},
		{
			name: "include some table names",
			args: args{
				inclusions: []string{"^movies_.*$"},
				exclusions: []string{},
			},
			err:         nil,
			entityNames: tableNames,
			want:        array.From(len(tableNames), func(offset int) bool { return offset > 4 && offset < 11 }),
		},
		{
			name: "include some query names",
			args: args{
				inclusions: []string{"get-.*$"},
				exclusions: []string{},
			},
			err:         nil,
			entityNames: queryNames,
			want:        array.From(len(queryNames), func(offset int) bool { return offset > 1 && offset < 7 }),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			validateEntityName, err := CreateValidateEntityNames(testCase.args.inclusions, testCase.args.exclusions)

			if testCase.err != nil {
				errMsgLeft := testCase.err.Error()

				errMsgRight := err.Error()

				assert.Containsf(t, errMsgRight, errMsgLeft, "want error %s but got %s", errMsgLeft, errMsgRight)
			} else {
				assert.Nil(t, err)

				for i, entityName := range testCase.entityNames {
					got := validateEntityName(entityName)

					assert.Equal(t, testCase.want[i], got, "want %v but got %v for name", testCase.want[i], got, entityName)
				}
			}
		})
	}
}
