package fs

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/joomcode/errorx"
)

type FileDiscovery interface {
	Find(dir string, pattern string, shallow bool) ([]Discovery, error)
}

type fileDiscovery struct{}

func (f fileDiscovery) Find(dir string, pattern string, shallow bool) ([]Discovery, error) {
	filepathRe, err := regexp.Compile(pattern)

	if err != nil {
		return nil, errorx.IllegalFormat.Wrap(err, "failed to compile regex")
	}

	found := make([]Discovery, 0)

	if shallow {
		files, err := os.ReadDir(dir)

		if err != nil {
			return nil, errorx.IllegalState.Wrap(err, "unable to read directory")
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			if !filepathRe.MatchString(file.Name()) {
				continue
			}

			fd := NewDiscovery(
				dir,
				file.Name(),
				filepath.Join(dir, file.Name()),
			)

			found = append(found, fd)
		}

		return found, nil
	}

	err = filepath.Walk(
		dir,
		func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return errorx.IllegalState.Wrap(err, "unable to walk path")
			}

			if info.IsDir() {
				return nil
			}

			if !filepathRe.MatchString(p) {
				return nil
			}

			fullPath, err := filepath.Abs(p)

			if err != nil {
				return nil
			}

			fd := NewDiscovery(
				filepath.Dir(p),
				filepath.Base(p),
				fullPath,
			)

			found = append(found, fd)

			return nil
		},
	)

	if err != nil {
		return nil, errorx.InternalError.Wrap(err, "unable to find files")
	}

	return found, nil
}

func NewFileDiscovery() FileDiscovery {
	return fileDiscovery{}
}
