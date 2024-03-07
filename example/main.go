package main

import (
	"log/slog"
	"os"
	"sharp"
)

func main() {
	//sharpHandlerExample()
	simpleHandlerExample()
}

func sharpHandlerExample() {
	opts := &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}
	handler1 := slog.NewTextHandler(os.Stdout, opts)
	handler2 := slog.NewJSONHandler(os.Stdout, opts)
	handler3 := sharp.NewSimpleHandler(os.Stdout, opts, "2006-01-02T15:04:05.000", true)

	handler := sharp.NewSharpHandler(handler1, handler2, handler3)

	slog.SetDefault(slog.New(handler))
	slog.Debug("debug")
	slog.Info("info")
	slog.Warn("warn")
	slog.Error("error")
}

func simpleHandlerExample() {
	opts := &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}
	handler := sharp.NewSimpleHandler(os.Stdout, opts, "2006-01-02T15:04:05.000", true)

	slog.SetDefault(slog.New(handler))
	slog.Debug("debug")
	slog.Info("info")
	slog.Warn("warn")
	slog.Error("error")
}
