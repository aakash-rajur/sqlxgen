package config

import (
	"log/slog"
	"os"
	"path"

	"github.com/joho/godotenv"
	"github.com/joomcode/errorx"
)

func loadAndExpand(workDir string, sqlxGenAltPath string) (string, error) {
	envFile := path.Join(workDir, ".env")

	_, err := os.Stat(envFile)

	env := make(map[string]string)

	if err == nil {
		env, err = godotenv.Read(envFile)
	} else if !os.IsNotExist(err) {
		slog.Warn("failed to read environment variables")
	}

	cfgPath := getSqlxGenPath(workDir, env, sqlxGenAltPath)

	content, err := os.ReadFile(cfgPath)

	if err != nil {
		return "", errorx.Decorate(err, "failed to read config file")
	}

	expanded := os.Expand(
		string(content),
		func(key string) string {
			return env[key]
		},
	)

	return expanded, nil
}

func getSqlxGenPath(
	workDir string,
	env map[string]string,
	sqlxGenAltPath string,
) string {
	if sqlxGenAltPath != "" {
		if path.IsAbs(sqlxGenAltPath) {
			return sqlxGenAltPath
		}

		return path.Join(workDir, sqlxGenAltPath)
	}

	cfgEnvPath, ok := env["SQLXGEN_CONFIG_PATH"]

	if !ok || cfgEnvPath == "" {
		return path.Join(workDir, "sqlxgen.yml")
	}

	if path.IsAbs(cfgEnvPath) {
		return cfgEnvPath
	}

	return path.Join(workDir, cfgEnvPath)
}
