package user

import (
	"github.com/codestates/WBABEProject-05/common"
	"github.com/codestates/WBABEProject-05/common/enum"
	"github.com/codestates/WBABEProject-05/model/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	if _, err := u.collection.InsertOne(ctx, user); err != nil {
		return enum.BlankSTR, err
	}

	return user.ID.Hex(), nil
}

func (u *userModel) UpdateUser(user *entity.User) (int64, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	filter := bson.M{"_id": user.ID}
	setField := user.NewBsonSetDForUpdateUser()
	updateRes, err := u.collection.UpdateOne(ctx, filter, setField)
	if err != nil {
		return 0, err
	}

	return updateRes.ModifiedCount, nil
}

func (u *userModel) UpdateUserPreOrder(user *entity.User) (int64, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	filter := bson.M{"_id": user.ID}
	setField := user.NewBsonSetDForUpdatePreOrder()
	updateRes, err := u.collection.UpdateOne(ctx, filter, setField)
	if err != nil {
		return 0, err
	}

	return updateRes.ModifiedCount, nil
}

func (u *userModel) SelectUser(id string) (*entity.User, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	obID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user *entity.User
	filter := bson.M{"_id": obID}
	if err := u.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userModel) DeleteUser(id string) (int64, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	obID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	filter := bson.M{"_id": obID}
	delRes, err := u.collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return delRes.DeletedCount, nil
}
