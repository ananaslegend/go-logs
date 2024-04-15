package logs

import (
	"context"
	"log/slog"
)

func WithLogMetric(ctx context.Context, metricName string, msg any) context.Context {
	if c, ok := ctx.Value(key{}).(logCtx); ok {
		c[metricName] = msg

		return ctx
	}

	return context.WithValue(ctx, key{}, logCtx{metricName: msg})
}

type ContextMiddleware struct {
	next slog.Handler
}

func NewContextMiddleware(next slog.Handler) *ContextMiddleware {
	return &ContextMiddleware{next: next}
}

func (h *ContextMiddleware) Handle(ctx context.Context, rec slog.Record) error {
	if c, ok := ctx.Value(key{}).(logCtx); ok {
		for metric, msg := range c {
			rec.Add(metric, msg)
		}
	}

	return h.next.Handle(ctx, rec)
}

func (h *ContextMiddleware) Enabled(ctx context.Context, rec slog.Level) bool {
	return h.next.Enabled(ctx, rec)
}

func (h *ContextMiddleware) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ContextMiddleware{next: h.next.WithAttrs(attrs)}
}

func (h *ContextMiddleware) WithGroup(name string) slog.Handler {
	return &ContextMiddleware{next: h.next.WithGroup(name)}
}

type logCtx map[string]any
type key struct{}
