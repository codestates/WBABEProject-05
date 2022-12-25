package login

import "github.com/codestates/WBABEProject-05/protocol"

var UserService UserServicer

type UserServicer interface {
	RegisterUser(user *protocol.RequestPostUser) (string, error)
	ModifyUser()
	FindUser()
	DeleteUser()
}
