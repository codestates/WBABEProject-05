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

const Store = "store"

func GetUserModel(col *mongo.Collection) *userModel {
	if usrMdInstance != nil {
		return usrMdInstance
	}
	//collection := mod.GetCollection(Store, "wbe")
	usrMdInstance = &userModel{
		collection: col,
	}
	return usrMdInstance
}

func (u *userModel) PostUser(user *entity.User) (string, error) {
	ctx, cancel := common.GetContext(common.ModelTimeOut)
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
