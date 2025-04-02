package app

import (
	"github.com/labstack/echo/v4"
	"github.com/rudolfoborges/asapi/internal/config"
	"go.uber.org/fx"
)

func Run() {
	fx.New(
		fx.Provide(config.ServerProvider),
		fx.Provide(config.DatabaseProvider),

		fx.Invoke(func(e *echo.Echo) {}),
	).Run()
}
