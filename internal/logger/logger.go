package logger

import (
	"io"
	"log/slog"
	"os"
)

func NewLogger(args *Args) *slog.Logger {
	handler := newHandler(args)

	return slog.New(handler)
}

func newHandler(args *Args) slog.Handler {
	options := slog.HandlerOptions{Level: args}

	var writer io.Writer = os.Stdout
	
	if args.Writer != nil {
		writer = args.Writer
	}

	if args.Format == "json" {
		return slog.NewJSONHandler(writer, &options)
	}

	return slog.NewTextHandler(writer, &options)
}
