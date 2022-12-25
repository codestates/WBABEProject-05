package login

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/user"
	"github.com/codestates/WBABEProject-05/protocol"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
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

func (u *userService) RegisterUser(usr *protocol.RequestPostUser) (string, error) {
	hashPassword := HashPassword(usr.Password)
	userEntity := &entity.User{
		Id:          primitive.NewObjectID(),
		Name:        usr.Name,
		NicName:     usr.NicName,
		Password:    hashPassword,
		PhoneNumber: usr.PhoneNumber,
		Role:        usr.Role,
		BaseTime: &entity.BaseTime{
			Created_at: time.Now(),
			Updated_at: time.Now(),
		},
	}
	savedId, err := u.userModel.PostUser(userEntity)
	if err != nil {
		return "", err
	}
	return savedId, err
}
func (u *userService) ModifyUser() {

}
func (u *userService) FindUser() {

}
func (u *userService) DeleteUser() {

}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	if err != nil {
		check = false
		return check, err
	}
	return check, nil
}
