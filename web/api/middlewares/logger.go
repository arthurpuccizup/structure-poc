package middlewares

import (
	"context"
	"github.com/labstack/echo/v4"
	"poc/internal/logging"
	"strconv"
	"time"
)

func ContextLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		logger, err := logging.NewLogger()
		if err != nil {
			return err
		}
		defer logger.Sync()

		sugar := logger.Sugar().With("request-id", echoCtx.Response().Header().Get("x-request-id"))

		ctx := context.WithValue(echoCtx.Request().Context(), logging.LoggerFlag, sugar)
		echoCtx.SetRequest(echoCtx.Request().Clone(ctx))

		return next(echoCtx)
	}
}

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		start := time.Now()
		err := next(echoCtx)

		ctx := echoCtx.Request().Context()

		req := echoCtx.Request()
		resp := echoCtx.Response()

		if err != nil {
			logging.LogErrorFromCtx(ctx, err)
			echoCtx.Error(err)
		}

		if logger, ok := logging.LoggerFromContext(ctx); ok {
			logger.Infow("finished request",
				"path", req.RequestURI,
				"method", req.Method,
				"status", strconv.Itoa(resp.Status),
				"time", time.Since(start).String())
		}

		return nil
	}
}
