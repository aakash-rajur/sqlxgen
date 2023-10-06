package config

import (
	"fmt"
	"log/slog"

	"github.com/aakash-rajur/sqlxgen/internal/logger"
	"github.com/aakash-rajur/sqlxgen/internal/utils/fs"
	"github.com/aakash-rajur/sqlxgen/internal/utils/writer"
	"github.com/joomcode/errorx"
	"gopkg.in/yaml.v3"
)

type SqlxGen struct {
	Connect       Connect          `json:"-" yaml:"-"`
	Fd            fs.FileDiscovery `json:"-" yaml:"-"`
	WriterCreator writer.Creator   `json:"-" yaml:"-"`
	Version       *string          `json:"version" yaml:"version"`
	ProjectDir    *string          `json:"projectDir" yaml:"projectDir"`
	LogArgs       *logger.Args     `json:"log" yaml:"log"`
	Configs       []Config         `json:"configs" yaml:"configs"`
}

func (gen *SqlxGen) InitLogger() {
	p := logger.NewLogger(gen.LogArgs)

	l := p.With("version", *gen.Version)

	slog.SetDefault(l)

	slog.Info("logger initialized")
}

func (gen *SqlxGen) Generate() error {
	version := *gen.Version

	if version != "1" {
		return errorx.IllegalArgument.New("unsupported sqlxgen version: %s", version)
	}

	slog.Debug("starting all config generations")

	slog.Info(fmt.Sprintf("have %d configs to generate", len(gen.Configs)))

	for _, cfg := range gen.Configs {
		slog.Info("generating", "name", *cfg.Name, "engine", *cfg.Engine)

		slog.Debug("config", "config", &cfg)

		err := cfg.Generate(gen.Connect, gen.Fd, gen.WriterCreator, *gen.ProjectDir)

		if err != nil {
			return err
		}

		slog.Info("generation ended", "name", *cfg.Name, "engine", *cfg.Engine)
	}

	slog.Debug("ended all config generations")

	return nil
}

func NewSqlxGen(args SqlxGenArgs) (*SqlxGen, error) {
	content, err := loadAndExpand(args.WorkingDir, args.SqlxAltPath)

	if err != nil {
		return nil, errorx.Decorate(err, "failed to load config")
	}

	sqlxGen := &SqlxGen{ProjectDir: &args.WorkingDir}

	err = yaml.Unmarshal([]byte(content), sqlxGen)

	if err != nil {
		return nil, errorx.Decorate(err, "failed to unmarshal config")
	}

	safeGenCfg, err := withDefaults(args, sqlxGen)

	if err != nil {
		return nil, err
	}

	return safeGenCfg, nil
}

func withDefaults(args SqlxGenArgs, sqlxGen *SqlxGen) (*SqlxGen, error) {
	logArgs := logger.DefaultLogArgs()

	safeGenCfg := &SqlxGen{
		Connect:       args.Connect,
		Fd:            args.Fd,
		WriterCreator: args.WriterCreator,
		Version:       sqlxGen.Version,
		ProjectDir:    sqlxGen.ProjectDir,
		LogArgs:       logArgs.Merge(sqlxGen.LogArgs),
		Configs:       make([]Config, 0),
	}

	for _, cfg := range sqlxGen.Configs {
		engine := *cfg.Engine

		if engine == "postgres" {
			pgConfig := defaultPgConfig()

			safeCfg := pgConfig.Merge(&cfg)

			safeGenCfg.Configs = append(safeGenCfg.Configs, *safeCfg)

			continue
		}

		if engine == "mysql" {
			mysqlConfig := defaultMysqlConfig()

			safeCfg := mysqlConfig.Merge(&cfg)

			safeGenCfg.Configs = append(safeGenCfg.Configs, *safeCfg)

			continue
		}

		return nil, errorx.IllegalArgument.New("unsupported engine: %s", engine)
	}

	return safeGenCfg, nil
}

type SqlxGenArgs struct {
	Connect       Connect
	Fd            fs.FileDiscovery
	WriterCreator writer.Creator
	WorkingDir    string
	SqlxAltPath   string
}
