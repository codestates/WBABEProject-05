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

func (u *userService) RegisterUser(user *request.RequestUser) (string, error) {
	postUser := user.NewPostUser()
	postUser.Password = u.hashPassword(user.Password)
	savedID, err := u.userModel.PostUser(postUser)
	if err != nil {
		return "", err
	}
	return savedID, err
}

func (u *userService) ModifyUser(ID string, usr *request.RequestPutUser) (int, error) {
	updateUser, err := usr.NewUpdatePutUser(ID)
	if err != nil {
		return 0, err
	}

	updateCount, err := u.userModel.UpdateUser(updateUser)
	if err != nil {
		return 0, err
	}
	return int(updateCount), nil
}

func (u *userService) FindUser(ID string) (*response.ResponseUser, error) {
	foundUser, err := u.userModel.SelectUser(ID)
	if err != nil {
		return nil, err
	}
	resUser := response.NewResponseUserFromUser(foundUser)
	return resUser, nil
}

func (u *userService) DeleteUser(ID string) (int, error) {
	deletedCount, err := u.userModel.DeleteUser(ID)
	if err != nil {
		return 0, err
	}
	return int(deletedCount), nil
}

func (u *userService) hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func (u *userService) verifyPassword(userPassword string, providedPassword string) (bool, error) {
	var check bool
	if err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword)); err != nil {
		return check, err
	}

	check = true
	return check, nil
}
