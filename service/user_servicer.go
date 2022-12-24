package service

import "github.com/codestates/WBABEProject-05/protocol"

type UserServicer interface {
	RegisterUser(user *protocol.RequestPostUser) (string, error)
	ModifyUser()
	FindUser()
	DeleteUser()
}
