package sharp

import (
	"context"
	"log/slog"
)

// NewSharpHandler wrap multi-handler as one handler
func NewSharpHandler(handler ...slog.Handler) slog.Handler {
	return &sharpHandler{handlers: handler}
}

type sharpHandler struct {
	handlers []slog.Handler
}

func (p *sharpHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, one := range p.handlers {
		if !one.Enabled(ctx, level) {
			return false
		}
	}
	return true
}

func (p *sharpHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, one := range p.handlers {
		if err := one.Handle(ctx, r); err != nil {
			return err
		}
	}
	return nil
}

func (p *sharpHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	for i, one := range p.handlers {
		p.handlers[i] = one.WithAttrs(attrs)
	}
	return p
}

func (p *sharpHandler) WithGroup(name string) slog.Handler {
	for i, one := range p.handlers {
		p.handlers[i] = one.WithGroup(name)
	}
	return p
}
