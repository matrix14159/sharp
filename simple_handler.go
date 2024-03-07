package sharp

import (
	"github.com/matrix14159/tint"
	"io"
	"log/slog"
)

func NewSimpleHandler(w io.Writer, opts *slog.HandlerOptions, timeLayout string, color bool) slog.Handler {
	if opts == nil {
		opts = &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo}
	}
	tmpOpts := &tint.Options{
		AddSource:   opts.AddSource,
		Level:       opts.Level,
		ReplaceAttr: opts.ReplaceAttr,
		TimeFormat:  timeLayout,
		NoColor:     !color,
	}
	return tint.NewHandler(w, tmpOpts)
}
