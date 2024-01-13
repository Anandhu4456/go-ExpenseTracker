// go:build wireinject
// +build wireinject

package di

import (
	"go-Expense/pkg/api"
	"go-Expense/pkg/api/handlers"
	"go-Expense/pkg/config"
	"go-Expense/pkg/db"
	"go-Expense/pkg/repository"
	"go-Expense/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*api.Server, error) {
	wire.Build(
		db.ConnectToDB,
		repository.NewUserRepo,
		usecase.NewUserUsecase,
		handlers.NewUserHandler,

		api.NewServer,
	)
	return &api.Server{}, nil
}
