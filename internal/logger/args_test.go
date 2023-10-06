package logger

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgsLevel(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		args *Args
		want slog.Level
	}{
		{
			name: "default",
			args: DefaultLogArgs(),
			want: slog.LevelInfo,
		},
		{
			name: "trace",
			args: &Args{
				LogLevel: "trace",
			},
			want: slog.Level(-8),
		},
		{
			name: "verbose",
			args: &Args{
				LogLevel: "verbose",
			},
			want: slog.Level(-8),
		},
		{
			name: "debug",
			args: &Args{
				LogLevel: "debug",
			},
			want: slog.LevelDebug,
		},
		{
			name: "info",
			args: &Args{
				LogLevel: "info",
			},
			want: slog.LevelInfo,
		},
		{
			name: "warn",
			args: &Args{
				LogLevel: "warn",
			},
			want: slog.LevelWarn,
		},
		{
			name: "error",
			args: &Args{
				LogLevel: "error",
			},
			want: slog.LevelError,
		},
		{
			name: "unknown",
			args: &Args{
				LogLevel: "unknown",
			},
			want: slog.LevelInfo,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			level := testCase.args.Level()

			assert.Equal(t, testCase.want, level, "Args.Level()")
		})
	}
}

func TestArgsMerge(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		args  *Args
		other *Args
		want  *Args
	}{
		{
			name:  "nil",
			args:  nil,
			other: nil,
			want:  nil,
		},
		{
			name:  "nil args",
			args:  nil,
			other: &Args{},
			want:  &Args{},
		},
		{
			name:  "nil other",
			args:  &Args{},
			other: nil,
			want:  &Args{},
		},
		{
			name: "empty",
			args: &Args{},
			other: &Args{
				LogLevel: "info",
				Format:   "text",
				Writer:   nil,
			},
			want: &Args{
				LogLevel: "info",
				Format:   "text",
				Writer:   nil,
			},
		},
		{
			name: "merge",
			args: &Args{
				LogLevel: "info",
				Format:   "text",
				Writer:   nil,
			},
			other: &Args{
				LogLevel: "debug",
				Format:   "json",
				Writer:   nil,
			},
			want: &Args{
				LogLevel: "debug",
				Format:   "json",
				Writer:   nil,
			},
		},
		{
			name: "merge with nil",
			args: &Args{
				LogLevel: "info",
				Format:   "text",
				Writer:   nil,
			},
			other: &Args{
				LogLevel: "debug",
				Format:   "json",
				Writer:   os.Stdout,
			},
			want: &Args{
				LogLevel: "debug",
				Format:   "json",
				Writer:   os.Stdout,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			args := testCase.args.Merge(testCase.other)

			assert.Equal(t, testCase.want, args, "Args.Merge()")
		})
	}
}

func TestDefaultLogArgs(t *testing.T) {
	t.Parallel()

	args := DefaultLogArgs()

	assert.Equal(t, "info", args.LogLevel, "DefaultLogArgs().LogLevel")
	assert.Equal(t, "text", args.Format, "DefaultLogArgs().Format")
	assert.Equal(t, os.Stdout, args.Writer, "DefaultLogArgs().Writer")
}
