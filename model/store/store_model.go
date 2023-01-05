package store

import (
	"github.com/codestates/WBABEProject-05/common"
	"github.com/codestates/WBABEProject-05/common/enum"
	common2 "github.com/codestates/WBABEProject-05/model/common"
	"github.com/codestates/WBABEProject-05/model/common/query"
	"github.com/codestates/WBABEProject-05/model/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

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

	ID, err := common2.ConvertStringToOBJID(storeId)
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
		return enum.BlankSTR, err
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

func (s *storeModel) SelectSortLimitedStore(pageQuery *query.PageQuery) ([]*entity.Store, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	filter := bson.M{}
	opt := pageQuery.NewSortFindOptions()
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
