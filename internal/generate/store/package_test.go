package store

import (
	_ "embed"
	"path"
	"testing"

	"github.com/aakash-rajur/sqlxgen/internal/utils/writer"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/assert"
)

func TestPackage_Generate(t *testing.T) {
	tmpDir := t.TempDir()

	genDir := path.Join(tmpDir, "gen/tmdb_pg/store")

	mw := writer.NewMemoryWriters()

	storePackage, err := NewPackage(
		mw.Creator,
		"github.com/aakash-rajur/sqlxgen/gen/tmdb_pg/store",
		genDir,
	)

	if err != nil {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	err = storePackage.Generate()

	assert.Nil(t, err)

	assert.Equal(t, 1, len(mw.Writers))

	pen := mw.Writers[0]

	assert.Equal(t, path.Join(genDir, "store.gen.go"), pen.FullPath)

	cupaloy.SnapshotT(t, pen.Content)
}

func TestNewPackage(t *testing.T) {
	tmpDir := t.TempDir()

	genDir := path.Join(tmpDir, "gen/tmdb_pg/store")

	mw := writer.NewMemoryWriters()

	got, err := NewPackage(
		mw.Creator,
		"github.com/aakash-rajur/sqlxgen/gen/tmdb_pg/store",
		genDir,
	)

	assert.Nil(t, err)

	want := Package{
		WriterCreator: mw.Creator,
		PackageName:   "store",
		PackageDir:    "github.com/aakash-rajur/sqlxgen/gen/tmdb_pg/store",
		GenDir:        genDir,
	}

	assert.Equal(t, want.PackageName, got.PackageName)

	assert.Equal(t, want.PackageDir, got.PackageDir)

	assert.Equal(t, want.GenDir, got.GenDir)
}
