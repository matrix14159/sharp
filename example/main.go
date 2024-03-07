package main

import (
	"io"
	"log/slog"
	"os"
	"sharp"
)

func main() {
	sharpHandlerExample()
	simpleHandlerExample()
	fileExample()
	zincsearchExample()
}

func sharpHandlerExample() {
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

func fileExample() {
	opts := &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}

	w := sharp.NewFileRotateWriter("./var/log/app.log", 50, 10, 30)
	ws := io.MultiWriter(os.Stdout, w)
	handler := sharp.NewSimpleHandler(ws, opts, "2006-01-02T15:04:05.000", true)

	slog.SetDefault(slog.New(handler))
	slog.Debug("debug")
	slog.Info("info")
	slog.Warn("warn")
	slog.Error("error")
}

func zincsearchExample() {
	opts := &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}

	w := sharp.NewZincsearchWriter("http://localhost:4080", "app.log", "admin", "admin", true)
	ws := io.MultiWriter(os.Stdout, w)
	handler := slog.NewJSONHandler(ws, opts)

	slog.SetDefault(slog.New(handler))
	slog.Debug("debug")
	slog.Info("info")
	slog.Warn("warn")
	slog.Error("error")
	//time.Sleep(1 * time.Second)
}
