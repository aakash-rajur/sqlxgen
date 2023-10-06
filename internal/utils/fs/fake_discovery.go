package fs

import (
	"fmt"
	"strings"

	"github.com/aakash-rajur/sqlxgen/internal/utils/array"
)

type FakeDiscover struct {
	Content  string `json:"content"`
	Dir      string `json:"dir"`
	Filename string `json:"filename"`
	FullPath string `json:"fullPath"`
}

func (fd FakeDiscover) GetDir() string {
	return fd.Dir
}

func (fd FakeDiscover) GetFilename() string {
	return fd.Filename
}

func (fd FakeDiscover) GetFullPath() string {
	return fd.FullPath
}

func (fd FakeDiscover) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Dir: %s", fd.Dir),
			fmt.Sprintf("Filename: %s", fd.Filename),
			fmt.Sprintf("FullPath: %s", fd.FullPath),
			fmt.Sprintf("Content: %s", fd.Content),
		},
		", ",
	)

	return fmt.Sprintf("FakeDiscover{%s}", content)
}

func (fd FakeDiscover) Load() ([]byte, error) {
	return []byte(fd.Content), nil
}

type fakeFileDiscovery []FakeDiscover

func (f fakeFileDiscovery) Find(_ string, _ string, _ bool) ([]Discovery, error) {
	ds := array.Map(
		f,
		func(fd FakeDiscover, _ int) Discovery {
			return fd
		},
	)

	return ds, nil
}

func NewFakeFileDiscovery(fds []FakeDiscover) FileDiscovery {
	return fakeFileDiscovery(fds)
}
