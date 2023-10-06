package queries

import (
	"encoding/json"
	"log/slog"

	"github.com/aakash-rajur/sqlxgen/internal/generate/types"
	"github.com/aakash-rajur/sqlxgen/internal/introspect"
	"github.com/aakash-rajur/sqlxgen/internal/utils/writer"
)

type Package struct {
	QueryTemplate    string
	StorePackageDir  string
	StorePackageName string
	Queries          []queryModel
}

func (p *Package) Generate() error {
	for _, query := range p.Queries {
		err := query.generate(p.QueryTemplate)

		if err != nil {
			queryJson, err2 := json.Marshal(query)

			if err2 != nil {
				return err
			}

			slog.Debug("unable to generate", "query", queryJson)

			return err
		}
	}

	return nil
}

func NewPackage(
	writerCreator writer.Creator,
	translate types.Translate,
	projectDir string,
	storePackageDir string,
	storePackageName string,
	queries []introspect.Query,
) (Package, error) {
	queryModels := make([]queryModel, len(queries))

	for i, query := range queries {
		qm, err := newQueryModel(
			writerCreator,
			translate,
			projectDir,
			storePackageDir,
			storePackageName,
			query,
		)

		if err != nil {
			return Package{}, err
		}

		queryModels[i] = qm
	}

	p := Package{
		QueryTemplate:    translate.QueryTemplate(),
		StorePackageDir:  storePackageDir,
		StorePackageName: storePackageName,
		Queries:          queryModels,
	}

	return p, nil
}
