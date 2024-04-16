# go-logs 
Utils for better logging using default slog.Logger and context.Context.

```bash
go get -u github.com/ananaslegend/go-logs
```

Example of using the logs package:
```go
func main() {
    handler := slog.Handler(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true}))
    handler = logs.NewContextMiddleware(handler)
    logger := slog.New(handler)

    ctx, _ := context.WithCancel(context.Background())

    ctx = logs.WithMetric(ctx, "metric", "value")

    err := func() error {
        ctx = logs.WithMetric(ctx, "error metric", "error value")
        err := fmt.Errorf("database error")
        return logs.WrapError(ctx, err) // storing error ctx inside of error
    }()

    logger.ErrorContext(
        logs.ErrorCtx(ctx, err), // trying to extract error ctx from error
        "error proceeding", 
        logs.ErrorMsg(err), 
    )   
}
```