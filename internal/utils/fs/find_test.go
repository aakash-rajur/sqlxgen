package fs

import (
	"os"
	"path"
	"testing"

	"github.com/joomcode/errorx"
	"github.com/stretchr/testify/assert"
)

func TestFileDiscoverFiles(t *testing.T) {
	t.Parallel()

	cwd, err := os.Getwd()

	if err != nil {
		t.Fatalf("unable to get current working directory: %s", err.Error())
	}

	type args struct {
		dir     string
		pattern string
		shallow bool
	}

	testCases := []struct {
		name  string
		args  args
		want  []discover
		error error
	}{
		{
			name: "valid directory shallow",
			args: args{
				dir:     "..",
				pattern: `[\w-]+.go$`,
				shallow: true,
			},
			want: []discover{
				{
					Dir:      "..",
					Filename: "filename_with_gen.go",
					FullPath: "../filename_with_gen.go",
				},
				{
					Dir:      "..",
					Filename: "filename_with_gen_test.go",
					FullPath: "../filename_with_gen_test.go",
				},
			},
		},
		{
			name: "valid directory not shallow",
			args: args{
				dir:     "..",
				pattern: `[\w-]+.go$`,
				shallow: false,
			},
			want: []discover{
				{
					Dir:      "../fs",
					Filename: "find.go",
					FullPath: path.Join(cwd, "../fs/find.go"),
				},
				{
					Dir:      "../fs",
					Filename: "find_test.go",
					FullPath: path.Join(cwd, "../fs/find_test.go"),
				},
				{
					Dir:      "../array",
					Filename: "filter.go",
					FullPath: path.Join(cwd, "../array/filter.go"),
				},
				{
					Dir:      "../array",
					Filename: "map.go",
					FullPath: path.Join(cwd, "../array/map.go"),
				},
				{
					Dir:      "../array",
					Filename: "reduce.go",
					FullPath: path.Join(cwd, "../array/reduce.go"),
				},
			},
		},
		{
			name: "invalid directory",
			args: args{
				dir:     "./invalid",
				pattern: `[\w-]+.go$`,
				shallow: true,
			},
			want:  nil,
			error: errorx.IllegalState.New("unable to read directory, cause: open ./invalid: no such file or directory"),
		},
	}

	fd := NewFileDiscovery()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := fd.Find(testCase.args.dir, testCase.args.pattern, testCase.args.shallow)

			if testCase.error != nil {
				if err == nil {
					t.Fatal("want error but got nil")
				}

				errMsgLeft := testCase.error.Error()

				errMsgRight := err.Error()

				assert.Contains(t, errMsgRight, errMsgLeft, "want error %s but got %s", errMsgLeft, errMsgRight)
			} else {
				assert.Nil(t, err, "want no error but got %s", err)

				for _, want := range testCase.want {
					assert.Contains(t, actual, want, "couldn't find %v", want)
				}
			}
		})
	}
}
