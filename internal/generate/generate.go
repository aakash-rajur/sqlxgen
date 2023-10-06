package generate

import (
	"log/slog"
	"path"

	"github.com/aakash-rajur/sqlxgen/internal/generate/models"
	"github.com/aakash-rajur/sqlxgen/internal/generate/queries"
	"github.com/aakash-rajur/sqlxgen/internal/generate/store"
	"github.com/aakash-rajur/sqlxgen/internal/generate/types"
	"github.com/aakash-rajur/sqlxgen/internal/introspect"
	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/aakash-rajur/sqlxgen/internal/utils/writer"
	"github.com/joomcode/errorx"
)

type Generate struct {
	WriterCreator   writer.Creator
	ProjectDir      string
	StorePackageDir string
	ModelPackageDir string
	Tables          []introspect.Table
	Queries         []introspect.Query
	Translate       types.Translate
}

func (gen Generate) Generate() error {
	projectPackageName, err := utils.GetProjectPackageName(gen.ProjectDir)

	if err != nil {
		return errorx.InitializationFailed.Wrap(err, "unable to get project package name")
	}

	storePackage, err := gen.generateStorePackage(projectPackageName)

	if err != nil {
		return errorx.InternalError.Wrap(err, "unable to generate store package")
	}

	_, err = gen.generateModelPackage(storePackage, projectPackageName)

	if err != nil {
		return errorx.InternalError.Wrap(err, "unable to generate models package")
	}

	_, err = gen.generateQueryPackage(storePackage)

	if err != nil {
		return errorx.InternalError.Wrap(err, "unable to generate queries package")
	}

	return nil
}

func (gen Generate) generateStorePackage(projectPackageName string) (store.Package, error) {
	storePackageDir := path.Join(projectPackageName, gen.StorePackageDir)

	storeGenDir := path.Join(gen.ProjectDir, gen.StorePackageDir)

	storePackage, err := store.NewPackage(gen.WriterCreator, storePackageDir, storeGenDir)

	if err != nil {
		return store.Package{}, errorx.InitializationFailed.Wrap(err, "unable to initialize store package")
	}

	err = storePackage.Generate()

	if err != nil {
		return store.Package{}, err
	}

	return storePackage, nil
}

func (gen Generate) generateModelPackage(
	storePackage store.Package,
	projectPackageName string,
) (models.Package, error) {
	slog.Debug("generating models package")

	modelPackageDir := path.Join(projectPackageName, gen.ModelPackageDir)

	modelGenDir := path.Join(gen.ProjectDir, gen.ModelPackageDir)

	modelPackage, err := models.NewPackage(
		gen.WriterCreator,
		gen.Translate,
		storePackage.PackageDir,
		storePackage.PackageName,
		modelPackageDir,
		modelGenDir,
		gen.Tables,
	)

	if err != nil {
		return models.Package{}, errorx.InitializationFailed.Wrap(err, "unable to initialize models package")
	}

	err = modelPackage.Generate()

	if err != nil {
		return models.Package{}, err
	}

	slog.Debug("generated models package")

	return modelPackage, nil
}

func (gen Generate) generateQueryPackage(storePackage store.Package) (queries.Package, error) {
	slog.Debug("generating queries package")

	queryPackage, err := queries.NewPackage(
		gen.WriterCreator,
		gen.Translate,
		gen.ProjectDir,
		storePackage.PackageDir,
		storePackage.PackageName,
		gen.Queries,
	)

	if err != nil {
		return queries.Package{}, errorx.InitializationFailed.Wrap(err, "unable to initialize queries package")
	}

	err = queryPackage.Generate()

	if err != nil {
		return queries.Package{}, err
	}

	slog.Debug("generated queries package")

	return queryPackage, nil
}
