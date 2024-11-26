package interfaces

import "user-service/pkg/model"

type UserRepository interface {
	Save(user model.User) (*model.User, error)
}
