package echomiddleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type calcMiddleware struct {
	logger *logrus.Logger
}

func New(logger *logrus.Logger) *calcMiddleware {
	return &calcMiddleware{
		logger: logger,
	}
}

func (c calcMiddleware) WithLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		timeStart := time.Now()
		next(ctx)
		code := ctx.Response().Status
		if code >= 400 && code <= 599 {
			c.logger.Errorf("%d %s %s %s", code, ctx.Request().Method, ctx.Request().URL, time.Since(timeStart))
		} else {
			c.logger.Infof("%d %s %s %s", code, ctx.Request().Method, ctx.Request().URL, time.Since(timeStart))
		}

		return nil
	}
}

func (c calcMiddleware) WithAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		uID := ctx.Request().Header.Get("Authorization")

		if uID == "" {
			ctx.Response().WriteHeader(echo.ErrUnauthorized.Code)
			return echo.ErrUnauthorized
		}

		return next(ctx)
	}
}

func (c calcMiddleware) WithRecover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				c.logger.Errorf("Recovered: %v", r)
				ctx.Response().WriteHeader(echo.ErrInternalServerError.Code)
			}
		}()

		return next(ctx)
	}
}
