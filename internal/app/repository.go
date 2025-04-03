package app

import (
	"github.com/rudolfoborges/asapi/internal/repository"
	"go.uber.org/fx"
)

var repositoryModule = fx.Module(
	"repository",

	fx.Provide(repository.NewAccountRepositoryPostgres),
)
