package utils

import (
	"os"
	"path"

	"github.com/joomcode/errorx"
	"golang.org/x/mod/modfile"
)

func GetProjectPackageName(dir string) (string, error) {
	filepath := path.Join(dir, "go.mod")

	file, err := os.ReadFile(filepath)

	if err != nil {
		return "", errorx.InitializationFailed.Wrap(err, "unable to read go.mod file")
	}

	mf, err := modfile.Parse(filepath, file, nil)

	if err != nil {
		return "", errorx.InitializationFailed.Wrap(err, "unable to parse go.mod file")
	}

	return mf.Module.Mod.Path, nil
}
