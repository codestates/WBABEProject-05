package user

import (
	"github.com/codestates/WBABEProject-05/model/entity"
)

var UserModel UserModeler

type UserModeler interface {
	PostUser(user *entity.User) (string, error)
	PutUser(user *entity.User) (int64, error)
	SelectUser(id string) (*entity.User, error)
	DeleteUser(id string) (int64, error)
}
