//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/yashdiq/nitip_user-service/pkg/api"
	handler "github.com/yashdiq/nitip_user-service/pkg/api/handler"

	db "github.com/yashdiq/nitip_user-service/pkg/db"
	repo "github.com/yashdiq/nitip_user-service/pkg/repository"

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
export GOPATH='/Users/iqra/go'