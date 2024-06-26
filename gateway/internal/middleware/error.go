package middleware

import (
	"github.com/alserov/restate/gateway/internal/middleware/wrappers"
	"github.com/alserov/restate/gateway/internal/utils"
	"github.com/labstack/echo/v4"
)

func WithErrorHandler(fn echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := fn(c)

		if err != nil {
			msg, st := utils.FromError(wrappers.ExtractLogger(wrappers.Ctx(c)), err)
			_ = c.JSON(st, map[string]string{
				"error": msg,
			})
		}

		return nil
	}
}
