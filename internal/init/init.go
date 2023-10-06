package init

import (
	_ "embed"
	"os"
	"path"

	"github.com/joomcode/errorx"
)

func Init(workDir string) error {
	fullPath := path.Join(workDir, "sqlxgen.yml")

	err := os.WriteFile(fullPath, []byte(sqlxGenCfg), 0644)

	if err != nil {
		return errorx.Decorate(err, "failed to write config file")
	}

	return nil
}

//go:embed sqlxgen.example.yml
var sqlxGenCfg string
