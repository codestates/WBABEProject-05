package store

import (
	"github.com/codestates/WBABEProject-05/common"
	"github.com/codestates/WBABEProject-05/model/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (s *storeModel) InsertMenu(store *entity.Store) (int, error) {
	ctx, cancel := common.GetContext(common.ModelTimeOut)
	defer cancel()
	filter := bson.D{{"id", store.Id}}
	update := bson.D{{"$set", bson.D{{"menu", store.Menu}}}}
	result, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	return int(result.ModifiedCount), nil
}
func (s *storeModel) DeleteMenu() {

}
func (s *storeModel) UpdateMenu() {

}
func (s *storeModel) SelectMenus() {

}
func (s *storeModel) SelectMenu() {

}
