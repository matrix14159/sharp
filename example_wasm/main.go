package main

import (
	"log/slog"

	"sharp"
)

// run steps:
// 1. enter example_wasm directory, compile to wasm: GOOS=js GOARCH=wasm go build -o=./server/public/main.wasm
// 2. enter example_wasm/server, start http server: go run main.go
// 3. open chrome, navigate to http://localhost:12000/
// 4. check chrome console
func main() {
	opts := &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}
	handler := sharp.NewSimpleHandler(sharp.NewWasmWriter(), opts, "2006-01-02T15:04:05.000", false)

	slog.SetDefault(slog.New(handler))
	slog.Debug("debug")
	slog.Info("info")
	slog.Warn("warn")
	slog.Error("error")
}
