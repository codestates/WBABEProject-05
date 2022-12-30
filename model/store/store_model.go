package store

import (
	"fmt"
	"github.com/codestates/WBABEProject-05/common"
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/entity/dom"
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

func NewStoreModel(col *mongo.Collection) *storeModel {
	if instance != nil {
		return instance
	}
	instance = &storeModel{
		collection: col,
	}
	return instance
}

func (s *storeModel) InsertMenu(storeId primitive.ObjectID, menu *dom.Menu) (int, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
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
func (s *storeModel) UpdateMenu(storeId primitive.ObjectID, menu *dom.Menu) (int, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	filter := bson.M{"_id": storeId, "menu": bson.M{"$elemMatch": bson.M{"_id": menu.Id}}}
	opt := bson.M{"$set": bson.M{"menu.$": menu}}
	result, err := s.collection.UpdateOne(ctx, filter, opt)
	if err != nil {
		return 0, err
	}

	return int(result.ModifiedCount), nil
}
func (s *storeModel) SelectStore(storeId primitive.ObjectID) (*entity.Store, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	var store *entity.Store
	filter := bson.D{{"_id", storeId}}
	if err := s.collection.FindOne(ctx, filter).Decode(&store); err != nil {
		return nil, err
	}
	return store, nil
}

// TODO 테스트필요
func (s *storeModel) SelectMenusByIds(storeId primitive.ObjectID, menuIds []primitive.ObjectID) ([]*dom.Menu, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	filter := bson.M{"_id": storeId, "menu": bson.M{"$elemMatch": bson.M{"_id": menuIds}}}
	opt := bson.M{"menu": true}
	menuCursor, err := s.collection.Find(ctx, filter, options.Find().SetProjection(opt).SetLimit(5))
	if err != nil {
		return nil, err
	}

	var menus []*dom.Menu
	if err = menuCursor.All(ctx, &menus); err != nil {
		return nil, err
	}
	return menus, nil
}

func (s *storeModel) SelectMenusSortBy(storeId primitive.ObjectID, sort string) ([]*dom.Menu, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	//var menus *[]*dom.Menu
	filter := bson.D{{"_id", storeId}}
	prj := bson.M{"menu": true}

	//opt := options.Find().SetProjection(prj)
	//menuCursor, err := s.collection.Find(ctx, filter, opt)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var store entity.Store
	//if err = menuCursor.All(ctx, &store); err != nil {
	//	return nil, err
	//}
	//fmt.Println(store)
	//for i, menu := range store.Menu {
	//	fmt.Println(i, menu)
	//}

	var store entity.Store

	opt := options.FindOne().SetProjection(prj).SetSort(bson.M{"": ""}) //객체내에 배열sort 어렵네...
	err := s.collection.FindOne(ctx, filter, opt).Decode(&store)
	if err != nil {
		return nil, err
	}
	fmt.Println(store)
	for i, menu := range store.Menu {
		fmt.Println("------")
		fmt.Println(i, menu)
	}
	return nil, nil
}

func (s *storeModel) InsertStore(store *entity.Store) (string, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	_, err := s.collection.InsertOne(ctx, store)
	if err != nil {
		return "", err
	}
	return store.Id.Hex(), nil
}

func (s *storeModel) SelectMenuByIdAndDelete(storeId, menuId primitive.ObjectID) (*entity.Store, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	var store *entity.Store
	filter := bson.M{"_id": storeId, "menu": bson.M{"$elemMatch": bson.M{"_id": menuId}}}
	opt := bson.M{"$pop": bson.M{"menu": -1}}
	if err := s.collection.FindOneAndUpdate(ctx, filter, opt).Decode(&store); err != nil {
		return nil, err
	}
	return store, nil
}
