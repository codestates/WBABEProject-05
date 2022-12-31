package login

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/protocol/request"
)

var UserService UserServicer

type UserServicer interface {
	RegisterUser(user *request.RequestUser) (string, error)
	ModifyUser(id string, usr *request.RequestUser) (int, error)
	FindUser(id string) (*entity.User, error)
	DeleteUser(id string) (int, error)
}
