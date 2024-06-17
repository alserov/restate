package middleware

import (
	"github.com/alserov/restate/gateway/internal/utils"
	"github.com/alserov/restate/gateway/internal/wrappers"
	"github.com/labstack/echo/v4"
)

func WithErrorHandler(fn echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := fn(c)
		lg := wrappers.ExtractLogger(c.Request().Context())

		if err != nil {
			msg, st := utils.FromError(lg, err)
			_ = c.JSON(st, map[string]string{
				"error": msg,
			})
		}

		return nil
	}
}
