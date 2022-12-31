package store

import (
	"github.com/codestates/WBABEProject-05/common"
	"github.com/codestates/WBABEProject-05/model/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (s *storeModel) SelectStoreByID(storeId string) (*entity.Store, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	ID, err := primitive.ObjectIDFromHex(storeId)
	if err != nil {
		return nil, err
	}

	var store *entity.Store
	filter := bson.M{"_id": ID}
	if err := s.collection.FindOne(ctx, filter).Decode(&store); err != nil {
		return nil, err
	}
	return store, nil
}

func (s *storeModel) SelectStoreByPhone(storePhone string) (*entity.Store, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	var store *entity.Store
	filter := bson.D{{"store_phone", storePhone}}
	if err := s.collection.FindOne(ctx, filter).Decode(&store); err != nil {
		return nil, err
	}
	return store, nil
}

func (s *storeModel) InsertStore(store *entity.Store) (string, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	_, err := s.collection.InsertOne(ctx, store)
	if err != nil {
		return "", err
	}
	return store.ID.Hex(), nil
}

func (s *storeModel) UpdateStore(store *entity.Store) (int, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	filter := bson.M{"_id": store.ID}
	opt := store.NewUpdateStoreBsonSetD()
	updateResult, err := s.collection.UpdateOne(ctx, filter, opt)
	if err != nil {
		return 0, err
	}

	return int(updateResult.ModifiedCount), nil
}
