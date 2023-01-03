package login

import (
	"github.com/codestates/WBABEProject-05/protocol/request"
	"github.com/codestates/WBABEProject-05/protocol/response"
)

var UserService UserServicer

type UserServicer interface {
	RegisterUser(user *request.RequestUser) (string, error)

	ModifyUser(ID string, user *request.RequestPutUser) (int, error)

	FindUser(ID string) (*response.ResponseUser, error)

	DeleteUser(ID string) (int, error)
}
