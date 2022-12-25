package store

import (
	"github.com/codestates/WBABEProject-05/common"
	"github.com/codestates/WBABEProject-05/model/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO 싱글톤을 메서드보다 변수로 갖다쓰는게 낫지않나? init 과 함께 사용하면 될것같은데 나중에 해보자
var instance *storeModel

type storeModel struct {
	collection *mongo.Collection
}

const Store = "store"

func GetStoreModel(col *mongo.Collection) *storeModel {
	if instance != nil {
		return instance
	}
	//collection := mod.GetCollection(Store, "wbe")
	instance = &storeModel{
		collection: col,
	}
	return instance
}

func (s *storeModel) InsertMenu(storeId primitive.ObjectID, menu *entity.Menu) (int, error) {
	ctx, cancel := common.GetContext(common.ModelTimeOut)
	defer cancel()

	filter := bson.D{{"_id", storeId}}
	update := bson.D{{"$push", bson.M{"menu": menu}}}
	result, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	return int(result.ModifiedCount), nil
}
func (s *storeModel) DeleteMenu() {

}
func (s *storeModel) UpdateMenu(storeId primitive.ObjectID, menu *entity.Menu) (int, error) {
	ctx, cancel := common.GetContext(common.ModelTimeOut)
	defer cancel()

	filter := bson.M{"_id": storeId, "menu": bson.M{"$elemMatch": bson.M{"_id": menu.Id}}}
	opt := bson.M{"$set": bson.M{"menu.$": menu}}
	result, err := s.collection.UpdateOne(ctx, filter, opt)
	if err != nil {
		return 0, err
	}

	return int(result.ModifiedCount), nil
}
func (s *storeModel) SelectMenus() {

}
func (s *storeModel) SelectMenu() {

}

// TODO 테스트필요
func (s *storeModel) SelectMenusByIds(storeId primitive.ObjectID, menuIds []primitive.ObjectID) ([]*entity.Menu, error) {
	ctx, cancel := common.GetContext(common.ModelTimeOut)
	defer cancel()

	filter := bson.M{"_id": storeId, "menu": bson.M{"$elemMatch": bson.M{"_id": menuIds}}}
	opt := bson.M{"menu": true}
	menuCursor, err := s.collection.Find(ctx, filter, options.Find().SetProjection(opt).SetLimit(5))
	if err != nil {
		return []*entity.Menu{}, err
	}

	var menus []*entity.Menu
	if err = menuCursor.All(ctx, &menus); err != nil {
		return []*entity.Menu{}, err
	}
	return menus, nil
}

func (s *storeModel) InsertStore(store *entity.Store) (string, error) {
	ctx, cancel := common.GetContext(common.ModelTimeOut)
	defer cancel()

	_, err := s.collection.InsertOne(ctx, store)
	if err != nil {
		return "", err
	}
	return store.Id.Hex(), nil
}

func (s *storeModel) SelectMenuByIdAndDelete(storeId, menuId primitive.ObjectID) (*entity.Store, error) {
	ctx, cancel := common.GetContext(common.ModelTimeOut)
	defer cancel()

	var store *entity.Store
	filter := bson.M{"_id": storeId, "menu": bson.M{"$elemMatch": bson.M{"_id": menuId}}}
	opt := bson.M{"$pop": bson.M{"menu": -1}}
	err := s.collection.FindOneAndUpdate(ctx, filter, opt).Decode(&store)
	if err != nil {
		return nil, err
	}
	return store, nil
}
