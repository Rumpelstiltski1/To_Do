package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLog(stage string) {
	var level slog.Level
	switch stage {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})

	Logger = slog.New(handler)
	slog.SetDefault(Logger)
}
