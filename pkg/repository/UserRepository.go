package repository

import (
	"user-service/pkg/model"
	"user-service/pkg/repository/interfaces"

	"github.com/gocql/gocql"
)

type userRepo struct {
	session *gocql.Session
}

func NewUserRepo(session *gocql.Session) interfaces.UserRepository {
	return &userRepo{
		session: session,
	}
}

func (u *userRepo) Save(user model.User) (*model.User, error) {
	var query string = "INSERT INTO up_users(id, username, email, phone) VALUES(?,?,?,?)"

	if err := u.session.Query(query, user.Id, user.Username, user.Email, user.Phone).Exec(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) GetById(id string) (*model.User, error) {

	var user model.User

	var query string = `SELECT id,title,content FROM up_users where id=?`

	if err := u.session.Query(query, id).Scan(&user.Id, &user.Username, &user.Email); err != nil {

		if err == gocql.ErrNotFound {
			return nil, err
		}

		return nil, err
	}

	return &user, nil
}