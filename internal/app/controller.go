package app

import (
	"github.com/rudolfoborges/asapi/internal/controller"
	"go.uber.org/fx"
)

var controllerModule = fx.Module(
	"controller",

	fx.Provide(controller.NewAccountController),
)
