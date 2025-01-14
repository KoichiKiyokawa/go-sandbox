package logger

import (
	"bytes"
	"io"
	"log/slog"
	"os"

	"github.com/danielgtaylor/huma/v2"
)

//nolint:gochecknoglobals
var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func Info(msg string) {
	logger.Info(msg)
}

func Error(msg string) {
	logger.Error(msg)
}

// https://github.com/danielgtaylor/huma/issues/470#issuecomment-2224007448
type humaContext huma.Context

type statusAwareContext struct {
	humaContext
	requestBody []byte
	bodyWriter  *bytes.Buffer
}

func (c *statusAwareContext) ResponseBody() string {
	if c.bodyWriter == nil {
		return ""
	}

	return c.bodyWriter.String()
}

func (c *statusAwareContext) BodyReader() io.Reader {
	c.requestBody, _ = io.ReadAll(c.humaContext.BodyReader())

	return io.NopCloser(bytes.NewBuffer(c.requestBody))
}

func (c *statusAwareContext) BodyWriter() io.Writer {
	if c.bodyWriter == nil {
		c.bodyWriter = bytes.NewBuffer(nil)
	}

	return io.MultiWriter(c.humaContext.BodyWriter(), c.bodyWriter)
}

func LogMiddleware(ctx huma.Context, next func(huma.Context)) {
	sctx := &statusAwareContext{humaContext: ctx, requestBody: nil, bodyWriter: nil}
	next(sctx)

	logger.Info(
		"Request and Response Log",
		"request", map[string]any{
			"method": ctx.Method(),
			"url":    ctx.URL(),
			"body":   string(sctx.requestBody),
		},
		"response", map[string]any{
			"status": ctx.Status(),
			"body":   sctx.ResponseBody(),
		},
	)
}
