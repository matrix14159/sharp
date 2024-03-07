package main

import (
	"log/slog"
	"os"
	"sharp"
)

func main() {
	opts := &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}
	handler1 := slog.NewTextHandler(os.Stdout, opts)
	handler2 := slog.NewJSONHandler(os.Stdout, opts)

	handler := sharp.NewSharpHandler(handler1, handler2)

	slog.SetDefault(slog.New(handler))
	slog.Debug("debug")
	slog.Info("hello2")
	slog.Warn("warn")
	slog.Error("world")
}
