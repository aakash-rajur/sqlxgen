package config

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/aakash-rajur/sqlxgen/internal/config/types"
	"github.com/aakash-rajur/sqlxgen/internal/generate"
	mysqlgen "github.com/aakash-rajur/sqlxgen/internal/generate/mysql"
	pggen "github.com/aakash-rajur/sqlxgen/internal/generate/pg"
	gentypes "github.com/aakash-rajur/sqlxgen/internal/generate/types"
	i "github.com/aakash-rajur/sqlxgen/internal/introspect"
	mysqlintrospect "github.com/aakash-rajur/sqlxgen/internal/introspect/mysql"
	pgintrospect "github.com/aakash-rajur/sqlxgen/internal/introspect/pg"
	"github.com/aakash-rajur/sqlxgen/internal/utils/array"
	"github.com/aakash-rajur/sqlxgen/internal/utils/fs"
	"github.com/aakash-rajur/sqlxgen/internal/utils/writer"
	"github.com/jmoiron/sqlx"
	"github.com/joomcode/errorx"
)

type Config struct {
	Writer   writer.Writer
	Name     *string         `json:"name" yaml:"name"`
	Engine   *string         `json:"engine" yaml:"engine"`
	Database *types.Database `json:"database" yaml:"database"`
	Source   *types.Source   `json:"source" yaml:"source"`
	Gen      *types.Gen      `json:"gen" yaml:"gen"`
}

func (c *Config) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Name: %v", *c.Name),
			fmt.Sprintf("Engine: %v", *c.Engine),
			fmt.Sprintf("Database: %v", c.Database),
			fmt.Sprintf("Source: %v", c.Source),
			fmt.Sprintf("Gen: %v", c.Gen),
		},
		", ",
	)

	return fmt.Sprintf("Config{%s}", content)
}

func (c *Config) Generate(
	connect Connect,
	fd fs.FileDiscovery,
	writerCreator writer.Creator,
	workDir string,
) error {
	engine := *c.Engine

	slog.Debug("connecting to database")

	dbUrl, err := c.Database.GetUrl(engine)

	if err != nil {
		return errorx.InitializationFailed.Wrap(err, "failed to get database url")
	}

	db, err := connect(engine, dbUrl)

	if err != nil {
		msg := fmt.Sprintf("failed to connect to database for %s", *c.Name)

		return errorx.Decorate(err, msg)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()

		if err != nil {
			panic(err)
		}

		slog.Debug("closed database connection")
	}(db)

	slog.Debug("connected to database")

	slog.Debug("starting transaction")

	tx, err := db.Beginx()

	if err != nil {
		panic(err)
	}

	defer func(tx *sqlx.Tx) {
		err := tx.Rollback()

		if err != nil {
			panic(err)
		}

		slog.Debug("rolled back transaction")
	}(tx)

	if engine == "postgres" {
		return c.generatePg(fd, writerCreator, tx, workDir)
	}

	if engine == "mysql" {
		return c.generateMysql(fd, writerCreator, tx, workDir)
	}

	return errorx.IllegalArgument.New("unsupported engine %s", engine)
}

func (c *Config) generatePg(
	fd fs.FileDiscovery,
	writerCreator writer.Creator,
	tx *sqlx.Tx,
	workDir string,
) error {
	pgi := pgintrospect.NewIntrospect(
		fd,
		pgintrospect.IntrospectArgs{
			Schemas:         c.Source.Models.Schemas,
			TableInclusions: c.Source.Models.Include,
			TableExclusions: c.Source.Models.Exclude,
			QueryDirs:       c.Source.Queries.Paths,
			QueryInclusions: c.Source.Queries.Include,
			QueryExclusions: c.Source.Queries.Exclude,
		},
	)

	pt := pggen.NewTranslate()

	return c.generate(writerCreator, pgi, pt, tx, workDir)
}

func (c *Config) generateMysql(
	fd fs.FileDiscovery,
	writerCreator writer.Creator,
	tx *sqlx.Tx,
	workDir string,
) error {
	mysqli := mysqlintrospect.NewIntrospect(
		fd,
		mysqlintrospect.IntrospectArgs{
			Schemas:         c.Source.Models.Schemas,
			TableInclusions: c.Source.Models.Include,
			TableExclusions: c.Source.Models.Exclude,
			QueryDirs:       c.Source.Queries.Paths,
			QueryInclusions: c.Source.Queries.Include,
			QueryExclusions: c.Source.Queries.Exclude,
		},
	)

	mt := mysqlgen.NewTranslate()

	return c.generate(writerCreator, mysqli, mt, tx, workDir)
}

func (c *Config) generate(
	writerCreator writer.Creator,
	introspect i.Introspect,
	translate gentypes.Translate,
	tx *sqlx.Tx,
	workDir string,
) error {
	slog.Debug("introspecting database")

	tables, err := introspect.IntrospectSchema(tx)

	if err != nil {
		return err
	}

	tableNames := array.Map(
		tables,
		func(table i.Table, _ int) string {
			return table.TableName
		},
	)

	slog.Info("found tables", "count", len(tables), "tables", tableNames)

	slog.Debug("introspecting queries")

	queries, err := introspect.IntrospectQueries(tx)

	if err != nil {
		return err
	}

	queryNames := array.Map(
		queries,
		func(query i.Query, _ int) string {
			return query.Filename
		},
	)

	slog.Info("found queries", "count", len(queries), "queries", queryNames)

	gen := generate.Generate{
		WriterCreator:   writerCreator,
		ProjectDir:      workDir,
		StorePackageDir: c.Gen.Store.Path,
		ModelPackageDir: c.Gen.Model.Path,
		Tables:          tables,
		Queries:         queries,
		Translate:       translate,
	}

	err = gen.Generate()

	if err != nil {
		return err
	}

	return nil
}

func (c *Config) Merge(other *Config) *Config {
	if other == nil {
		return c
	}

	if c == nil {
		return other
	}

	c.Writer = other.Writer

	if other.Name != nil {
		c.Name = other.Name
	}

	if other.Engine != nil {
		c.Engine = other.Engine
	}

	c.Database = c.Database.Merge(other.Database)
	c.Source = c.Source.Merge(other.Source)
	c.Gen = c.Gen.Merge(other.Gen)

	return c
}

type Connect func(engine string, connectionUrl string) (*sqlx.DB, error)
