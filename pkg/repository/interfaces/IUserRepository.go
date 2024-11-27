package interfaces

import "github.com/yashdiq/nitip_user-service/pkg/model"

type UserRepository interface {
	Save(user model.User) (*model.User, error)
}
