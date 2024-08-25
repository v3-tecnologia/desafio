package infra

import (
	"log/slog"
	"os"
)

type Logger = *slog.Logger

func NewLogger() Logger {

	logger := slog.NewJSONHandler(os.Stdout, nil)

	return slog.New(logger)
}
