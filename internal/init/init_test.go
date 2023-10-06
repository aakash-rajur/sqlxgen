package init

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	tmpDir := t.TempDir()

	err := Init(tmpDir)

	assert.Nil(t, err)

	fullPath := path.Join(tmpDir, "sqlxgen.yml")

	got, err := os.ReadFile(fullPath)

	assert.Nil(t, err)

	assert.Equal(t, sqlxGenCfg, string(got))
}
