package fs

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscover_GetDir(t *testing.T) {
	t.Parallel()

	type fields struct {
		Dir      string
		Filename string
		FullPath string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "valid 1",
			fields: fields{
				Dir:      "dir",
				Filename: "example.txt",
				FullPath: "fullPath",
			},
			want: "dir",
		},
		{
			name: "valid 2",
			fields: fields{
				Dir:      "/tmp/dir",
				Filename: "hello.txt",
				FullPath: "fullPath",
			},
			want: "/tmp/dir",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := discover{
				Dir:      tt.fields.Dir,
				Filename: tt.fields.Filename,
				FullPath: tt.fields.FullPath,
			}

			got := d.GetDir()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDiscover_GetFilename(t *testing.T) {
	t.Parallel()

	type fields struct {
		Dir      string
		Filename string
		FullPath string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "valid 1",
			fields: fields{
				Dir:      "dir",
				Filename: "example.txt",
				FullPath: "fullPath",
			},
			want: "example.txt",
		},
		{
			name: "valid 2",
			fields: fields{
				Dir:      "/tmp/dir",
				Filename: "hello.txt",
				FullPath: "fullPath",
			},
			want: "hello.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := discover{
				Dir:      tt.fields.Dir,
				Filename: tt.fields.Filename,
				FullPath: tt.fields.FullPath,
			}

			got := d.GetFilename()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDiscover_GetFullPath(t *testing.T) {
	t.Parallel()

	type fields struct {
		Dir      string
		Filename string
		FullPath string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "valid 1",
			fields: fields{
				Dir:      "/dir",
				Filename: "example.txt",
				FullPath: "/dir/example.txt",
			},
			want: "/dir/example.txt",
		},
		{
			name: "valid 2",
			fields: fields{
				Dir:      "/tmp/dir",
				Filename: "hello.txt",
				FullPath: "/tmp/dir/hello.txt",
			},
			want: "/tmp/dir/hello.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := discover{
				Dir:      tt.fields.Dir,
				Filename: tt.fields.Filename,
				FullPath: tt.fields.FullPath,
			}

			got := d.GetFullPath()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDiscover_String(t *testing.T) {
	t.Parallel()

	type fields struct {
		Dir      string
		Filename string
		FullPath string
	}

	testCases := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "valid file discovery",
			fields: fields{
				Dir:      "dir",
				Filename: "filename",
				FullPath: "dir/filename",
			},
			want: "Discovery{Dir: dir, Filename: filename, FullPath: dir/filename}",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			fd := discover{
				Dir:      testCase.fields.Dir,
				Filename: testCase.fields.Filename,
				FullPath: testCase.fields.FullPath,
			}

			got := fd.String()

			assert.Equal(t, testCase.want, got, "want %s but got %s", testCase.want, got)
		})
	}
}

func TestDiscover_Load(t *testing.T) {
	tmpDir := t.TempDir()

	err := os.WriteFile(path.Join(tmpDir, "filename"), []byte("content"), 0644)

	if err != nil {
		t.Fatalf("unable to write file: %s", err.Error())
	}

	fd := discover{
		Dir:      tmpDir,
		Filename: "filename",
		FullPath: path.Join(tmpDir, "filename"),
	}

	got, err := fd.Load()

	assert.Nil(t, err, "want no error but got %s", err)

	assert.Equal(t, []byte("content"), got, "want %s but got %s", "content", got)
}
