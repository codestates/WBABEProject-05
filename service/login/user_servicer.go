package login

import (
	"github.com/codestates/WBABEProject-05/protocol/request"
	"github.com/codestates/WBABEProject-05/protocol/response"
)

var UserService UserServicer

type UserServicer interface {
	RegisterUser(user *request.RequestUser) (string, error)
	ModifyUser(id string, usr *request.RequestPutUser) (int, error)
	FindUser(id string) (*response.ResponseUser, error)
	DeleteUser(id string) (int, error)
}
