package fs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFakeDiscover_GetDir(t *testing.T) {
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

func TestFakeDiscover_GetFilename(t *testing.T) {
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

func TestFakeDiscover_GetFullPath(t *testing.T) {
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

func TestFakeDiscover_String(t *testing.T) {
	t.Parallel()

	type fields struct {
		Dir      string
		Filename string
		FullPath string
		Content  string
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
				Content:  "content",
			},
			want: "FakeDiscover{Dir: dir, Filename: filename, FullPath: dir/filename, Content: content}",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			fd := FakeDiscover{
				Dir:      testCase.fields.Dir,
				Filename: testCase.fields.Filename,
				FullPath: testCase.fields.FullPath,
				Content:  testCase.fields.Content,
			}

			got := fd.String()

			assert.Equal(t, testCase.want, got, "want %s but got %s", testCase.want, got)
		})
	}
}

func TestFakeDiscover_Load(t *testing.T) {
	t.Parallel()

	type fields struct {
		Dir      string
		Filename string
		FullPath string
		Content  string
	}

	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "valid 1",
			fields: fields{
				Dir:      "dir",
				Filename: "example.txt",
				FullPath: "fullPath",
				Content:  "content",
			},
			want:    []byte("content"),
			wantErr: false,
		},
		{
			name: "valid 2",
			fields: fields{
				Dir:      "/tmp/dir",
				Filename: "hello.txt",
				FullPath: "fullPath",
				Content:  "content",
			},
			want:    []byte("content"),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fd := FakeDiscover{
				Dir:      tt.fields.Dir,
				Filename: tt.fields.Filename,
				FullPath: tt.fields.FullPath,
				Content:  tt.fields.Content,
			}

			got, err := fd.Load()

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
