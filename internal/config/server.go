package config

import (
	"context"
	"os"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func ServerProvider(lc fx.Lifecycle) *echo.Echo {
	e := echo.New()

	lc.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			go func() {
				if err := e.Start(os.Getenv("PORT")); err != nil {
					e.Logger.Fatal("shutting down the server")
				}
			}()

			return nil
		},
		OnStop: func(c context.Context) error {
			return e.Shutdown(c)
		},
	})

	return e
}
