package config

import (
	"github.com/labstack/echo/v4"
	"github.com/rudolfoborges/asapi/internal/controller"
	"go.uber.org/fx"
)

type ControllerIn struct {
	fx.In
	AccountController *controller.AccountController
}

func Routes(e *echo.Echo, in ControllerIn) *echo.Echo {
	v1 := e.Group("/v1")
	v1.POST("/accounts", in.AccountController.CreateAccount)

	return e
}
