package config

import (
	"log/slog"
	"maps"
	"os"
	"path"
	"strings"

	"github.com/joho/godotenv"
	"github.com/joomcode/errorx"
)

func loadAndExpand(workDir string, sqlxGenAltPath string) (string, error) {
	env := loadEnvFromFile(workDir)

	cfgPath := getSqlxGenPath(workDir, env, sqlxGenAltPath)

	content, err := os.ReadFile(cfgPath)

	if err != nil {
		return "", errorx.Decorate(err, "failed to read config file")
	}

	expanded := os.Expand(
		string(content),
		func(key string) string {
			value, ok := env[key]

			if !ok {
				slog.Warn("environment variable not found", "key", key)

				return ""
			}

			return value
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

func loadEnvFromFile(workDir string) map[string]string {
	env := loadEnvFromOS()

	envFile := path.Join(workDir, ".env")

	_, err := os.Stat(envFile)

	if err != nil {
		if !os.IsNotExist(err) {
			slog.Warn("local environment file not found")
		}

		return env
	}

	dotEnv, err := godotenv.Read(envFile)

	if err != nil {
		slog.Warn("failed to read environment file")

		return env
	}

	if dotEnv == nil {
		return env
	}

	maps.Copy(env, dotEnv)

	return env
}

func loadEnvFromOS() map[string]string {
	env := make(map[string]string)

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)

		env[pair[0]] = pair[1]
	}

	return env
}
