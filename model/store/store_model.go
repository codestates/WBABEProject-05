package store

import (
	"github.com/codestates/WBABEProject-05/common"
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/util"
	"github.com/codestates/WBABEProject-05/protocol/page"
	"go.mongodb.org/mongo-driver/bson"
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

	ID, err := util.ConvertStringToObjID(storeId)
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

	if _, err := s.collection.InsertOne(ctx, store); err != nil {
		return "", err
	}

	return store.ID.Hex(), nil
}

func (s *storeModel) UpdateStore(store *entity.Store) (int64, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	filter := bson.M{"_id": store.ID}
	opt := store.NewUpdateStoreBsonSetD()
	updateResult, err := s.collection.UpdateOne(ctx, filter, opt)
	if err != nil {
		return 0, err
	}

	return updateResult.ModifiedCount, nil
}

func (s *storeModel) SelectSortLimitedStore(sort *page.Sort, skip int, limit int) ([]*entity.Store, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	filter := bson.M{}
	opt := util.NewSortFindOptions(sort, skip, limit)
	receiptCursor, err := s.collection.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}

	var stores []*entity.Store
	if err = receiptCursor.All(ctx, &stores); err != nil {
		return nil, err
	}

	return stores, nil
}

func (s *storeModel) SelectTotalCount() (int64, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	count, err := s.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return count, nil
}
