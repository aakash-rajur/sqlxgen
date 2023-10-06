package fs

import (
	"fmt"
	"os"
	"strings"
)

type Discovery interface {
	GetDir() string

	GetFilename() string

	GetFullPath() string

	Load() ([]byte, error)
}

type discover struct {
	Dir      string `json:"dir"`
	Filename string `json:"filename"`
	FullPath string `json:"fullPath"`
}

func (d discover) GetDir() string {
	return d.Dir
}

func (d discover) GetFilename() string {
	return d.Filename
}

func (d discover) GetFullPath() string {
	return d.FullPath
}

func (d discover) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Dir: %s", d.Dir),
			fmt.Sprintf("Filename: %s", d.Filename),
			fmt.Sprintf("FullPath: %s", d.FullPath),
		},
		", ",
	)

	return fmt.Sprintf("Discovery{%s}", content)
}

func (d discover) Load() ([]byte, error) {
	return os.ReadFile(d.FullPath)
}

func NewDiscovery(
	dir string,
	filename string,
	fullPath string,
) Discovery {
	return discover{
		Dir:      dir,
		Filename: filename,
		FullPath: fullPath,
	}
}
