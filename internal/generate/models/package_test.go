package models

import (
	"path"
	"testing"

	"github.com/aakash-rajur/sqlxgen/internal/generate/types"
	"github.com/aakash-rajur/sqlxgen/internal/introspect"
	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/aakash-rajur/sqlxgen/internal/utils/writer"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/assert"
)

func TestNewPackage(t *testing.T) {
	tables, err := utils.FromJson[introspect.Table](
		[]string{actorTableJson, movieTableJson},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	ft := types.NewFakeTranslate(
		`package {{ .PackageName }}
var content = `+"`"+`{
  "packageName": {{ .PackageName | ToJson }}, 
  "imports": {{ .Imports | ToJson }},
	"model": {{ .Model | ToJson }},
  "insertFields": {{ .InsertFields | ToJson }}
	"updateFields": {{ .UpdateFields | ToJson }}
	"selectFields": {{ .SelectFields | ToJson }}
}`+"`",
		"",
	)

	got, err := NewPackage(
		nil,
		ft,
		"store",
		"store",
		"gen/models",
		"gen/models",
		tables,
	)

	assert.Nil(t, err)

	cupaloy.SnapshotT(t, got)
}

func TestPackage_Generate(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()

	mw := writer.NewMemoryWriters()

	tables, err := utils.FromJson[introspect.Table](
		[]string{actorTableJson, movieTableJson},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	ft := types.NewFakeTranslate(
		`package {{ .PackageName }}
var content = `+"`"+`{
  "packageName": {{ .PackageName | ToJson }}, 
  "imports": {{ .Imports | ToJson }},
	"model": {{ .Model | ToJson }},
  "insertFields": {{ .InsertFields | ToJson }}
	"updateFields": {{ .UpdateFields | ToJson }}
	"selectFields": {{ .SelectFields | ToJson }}
}`+"`",
		"",
	)

	pkg, err := NewPackage(
		mw.Creator,
		ft,
		"gen/store",
		"gen/store",
		"gen/models",
		tmpDir,
		tables,
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = pkg.Generate()

	assert.Nil(t, err)

	paths := []string{
		path.Join(tmpDir, "actor.gen.go"),
		path.Join(tmpDir, "movie.gen.go"),
	}

	for i, p := range paths {
		testName, _ := utils.SplitFilename(path.Base(p))

		t.Run(testName, func(t *testing.T) {
			got := mw.Writers[i]

			assert.Equal(t, p, got.FullPath)

			cupaloy.SnapshotT(t, got.Content)
		})
	}
}
