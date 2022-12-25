package user

import "github.com/codestates/WBABEProject-05/model/entity"

var UserModel UserModeler

type UserModeler interface {
	PostUser(user *entity.User) (string, error)
	PutUser()
	SelectUser()
	DeleteUser()
}
