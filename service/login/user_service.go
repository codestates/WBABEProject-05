package login

import (
	"github.com/codestates/WBABEProject-05/model/user"
	"github.com/codestates/WBABEProject-05/protocol/request"
	"github.com/codestates/WBABEProject-05/protocol/response"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type userService struct {
	userModel user.UserModeler
}

var instance *userService

func NewUserService(modeler user.UserModeler) *userService {
	if instance != nil {
		return instance
	}
	instance = &userService{
		userModel: modeler,
	}
	return instance
}

func (u *userService) RegisterUser(usr *request.RequestUser) (string, error) {
	postUser := usr.NewPostUser()
	postUser.Password = u.hashPassword(usr.Password)
	savedId, err := u.userModel.PostUser(postUser)
	if err != nil {
		return "", err
	}
	return savedId, err
}
func (u *userService) ModifyUser(ID string, usr *request.RequestPutUser) (int, error) {
	updateUser, err := usr.NewUpdatePutUser(ID)
	if err != nil {
		return 0, err
	}

	updateCount, err := u.userModel.PutUser(updateUser)
	if err != nil {
		return 0, err
	}
	return updateCount, nil
}
func (u *userService) FindUser(id string) (*response.ResponseUser, error) {
	findUser, err := u.userModel.SelectUser(id)
	if err != nil {
		return nil, err
	}
	resUser := response.NewResponseUserFromUser(findUser)
	return resUser, nil
}
func (u *userService) DeleteUser(id string) (int, error) {
	deleteUser, err := u.userModel.DeleteUser(id)
	if err != nil {
		return 0, err
	}
	return deleteUser, nil
}

func (u *userService) hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func (u *userService) verifyPassword(userPassword string, providedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	if err != nil {
		check = false
		return check, err
	}
	return check, nil
}
