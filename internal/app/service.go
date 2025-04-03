package app

import (
	"github.com/rudolfoborges/asapi/internal/service"
	"go.uber.org/fx"
)

var serviceModule = fx.Module(
	"service",

	fx.Provide(service.NewAccountService),
)
