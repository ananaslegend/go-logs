package logs

import (
	"context"
	"errors"
	"log/slog"
)

func WrapError(ctx context.Context, err error) error {
	c := logCtx{}
	if x, ok := ctx.Value(key{}).(logCtx); ok {
		c = x
	}

	return &errorWithLogCtx{
		next: err,
		ctx:  c,
	}
}

func ErrorCtx(ctx context.Context, err error) context.Context {
	var e *errorWithLogCtx
	if errors.As(err, &e) {
		return context.WithValue(ctx, key{}, e.ctx)
	}

	return ctx
}

func ErrorMsg(err error) slog.Attr {
	return slog.String("error", err.Error())

}

type errorWithLogCtx struct {
	next error
	ctx  logCtx
}

func (e *errorWithLogCtx) Error() string {
	return e.next.Error()
}
