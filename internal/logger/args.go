package logger

import (
	"io"
	"log/slog"
	"os"
)

type Args struct {
	LogLevel string `json:"level" yaml:"level"`
	Format   string `json:"format" yaml:"format"`
	Writer   io.Writer
}

func (l *Args) Level() slog.Level {
	switch l.LogLevel {
	case "trace", "verbose":
		return slog.Level(-8)
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	}
	return slog.LevelInfo
}

func (l *Args) Merge(other *Args) *Args {
	if l == nil {
		return other
	}

	if other == nil {
		return l
	}

	if other.LogLevel != "" {
		l.LogLevel = other.LogLevel
	}

	if other.Format != "" {
		l.Format = other.Format
	}

	if other.Writer != nil {
		l.Writer = other.Writer
	}

	return l
}

func DefaultLogArgs() *Args {
	return &Args{
		LogLevel: "info",
		Format:   "text",
		Writer:   os.Stdout,
	}
}
