package middlewares

import (
	"context"
	"github.com/labstack/echo/v4"
)

func Validator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "api-request-context", "")
		c.SetRequest(c.Request().Clone(ctx))
		return next(c)
	}
}
