package user

import "github.com/codestates/WBABEProject-05/model/entity"

type UserModeler interface {
	PostUser(user *entity.User) (string, error)
	PutUser()
	SelectUser()
	DeleteUser()
}
