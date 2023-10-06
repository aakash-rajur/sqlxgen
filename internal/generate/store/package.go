package store

import (
	"bytes"
	_ "embed"
	"go/format"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/aakash-rajur/sqlxgen/internal/utils/casing"
	"github.com/aakash-rajur/sqlxgen/internal/utils/writer"
	"github.com/joomcode/errorx"
)

type Package struct {
	WriterCreator writer.Creator `json:"-"`
	PackageName   string         `json:"package_name"`
	PackageDir    string         `json:"package_dir"`
	GenDir        string         `json:"gen_dir"`
}

func (p Package) Generate() error {
	slog.Debug("generating store package")

	tmpl, err := template.New("store").Parse(storeTemplate)

	if err != nil {
		return errorx.IllegalFormat.Wrap(err, "unable to parse store template")
	}

	var storeFileBuffer bytes.Buffer

	err = tmpl.Execute(
		&storeFileBuffer,
		map[string]interface{}{
			"PackageName": p.PackageName,
		},
	)

	if err != nil {
		return errorx.InternalError.Wrap(err, "unable to execute store template")
	}

	formatted, err := format.Source(storeFileBuffer.Bytes())

	if err != nil {
		return errorx.InternalError.Wrap(err, "unable to format store template")
	}

	err = os.MkdirAll(p.GenDir, 0755)

	if err != nil {
		return errorx.InitializationFailed.Wrap(err, "unable to create directory for store")
	}

	genFileName := utils.FilenameWithGen("store.go")

	storeFilePath := path.Join(p.GenDir, genFileName)

	pen := p.WriterCreator(storeFilePath, string(formatted))

	err = pen.Write()

	if err != nil {
		return errorx.InternalError.Wrap(err, "unable to write store file")
	}

	slog.Debug("generated store package")

	return nil
}

func NewPackage(
	writerCreator writer.Creator,
	packageDir string,
	genDir string,
) (Package, error) {
	parentDir := filepath.Base(packageDir)

	packageName, err := casing.SnakeCase(parentDir)

	if err != nil {
		return Package{}, errorx.IllegalState.Wrap(err, "unable to generate package name")
	}

	p := Package{
		WriterCreator: writerCreator,
		PackageName:   packageName,
		PackageDir:    packageDir,
		GenDir:        genDir,
	}

	return p, nil
}

//go:embed store.go.tmpl
var storeTemplate string
