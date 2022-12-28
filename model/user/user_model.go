package user

import (
	"github.com/codestates/WBABEProject-05/common"
	"github.com/codestates/WBABEProject-05/model/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

var usrMdInstance *userModel

type userModel struct {
	collection *mongo.Collection
}

func NewUserModel(col *mongo.Collection) *userModel {
	if usrMdInstance != nil {
		return usrMdInstance
	}
	usrMdInstance = &userModel{
		collection: col,
	}
	return usrMdInstance
}

func (u *userModel) PostUser(user *entity.User) (string, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()
	_, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}
	return user.Id.Hex(), nil
}
func (u *userModel) PutUser() {

}
func (u *userModel) SelectUser() {

}
func (u *userModel) DeleteUser() {

}
