package middlewares

import (
	"context"
	"github.com/labstack/echo/v4"
	"poc/internal/observ"
)

func ContextLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger, err := observ.NewLogger()
		if err != nil {
			return err
		}
		defer logger.Sync()

		sugar := logger.Sugar().With("request-id", c.Response().Header().Get("x-request-id"))
		ctx := context.WithValue(c.Request().Context(), observ.LoggerFlag, sugar)
		c.SetRequest(c.Request().Clone(ctx))

		return next(c)
	}
}
