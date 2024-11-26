//go:build wireinject
// +build wireinject

package di

import (
	http "user-service/pkg/api"
	handler "user-service/pkg/api/handler"

	db "user-service/pkg/db"
	repo "user-service/pkg/repository"

	"github.com/google/wire"
)

func Initialize() (*http.ServerHttp, error) {
	wire.Build(
		db.DBConnect,
		repo.NewUserRepo,
		handler.NewUserHandler,
		http.NewServerHttp,
	)
	return &http.ServerHttp{}, nil
}
