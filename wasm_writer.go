//go:build js && wasm

package sharp

import (
	"io"

	"github.com/matrix14159/tint"
)

// NewWasmWriter work will SimpleHandler will write to browser console
func NewWasmWriter() io.Writer {
	return tint.NewWasmConsole()
}
