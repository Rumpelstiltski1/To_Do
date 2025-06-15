package logger

import (
	"io"
	"log"
	"log/slog"
	"os"
)

var Logger *slog.Logger

var file *os.File

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

	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Не удалось открыть файл для логирования", err)
	}
	MultiWriter := io.MultiWriter(file, os.Stdout)

	handler := slog.NewJSONHandler(MultiWriter, &slog.HandlerOptions{Level: level})
	Logger = slog.New(handler)
	slog.SetDefault(Logger)
}

func CloseFile() {
	if file != nil {
		file.Close()
	}
}
