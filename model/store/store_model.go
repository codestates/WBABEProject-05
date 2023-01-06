package store

import (
	"github.com/codestates/WBABEProject-05/common"
	"github.com/codestates/WBABEProject-05/common/convertor"
	"github.com/codestates/WBABEProject-05/common/enum"
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/query"
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

func (s *storeModel) SelectStoreByID(storeID string) (*entity.Store, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	ID, err := convertor.ConvertStringToOBJID(storeID)
	if err != nil {
		return nil, err
	}

	var store *entity.Store
	filter := query.GetDefaultIDFilter(ID)
	if err := s.collection.FindOne(ctx, filter).Decode(&store); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *storeModel) SelectStoreByIDAndUserID(storeID, userID string) (*entity.Store, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	sID, err := convertor.ConvertStringToOBJID(storeID)
	if err != nil {
		return nil, err
	}

	uID, err := convertor.ConvertStringToOBJID(userID)
	if err != nil {
		return nil, err
	}

	var store *entity.Store
	filter := bson.D{{"_id", sID}, {"user_id", uID}}
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

	filter := query.GetDefaultIDFilter(store.ID)
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

func (s *storeModel) UpdatePullRecommendMenu(storeID, menuID string) (int64, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	sID, err := convertor.ConvertStringToOBJID(storeID)
	if err != nil {
		return 0, err
	}

	mID, err := convertor.ConvertStringToOBJID(menuID)
	if err != nil {
		return 0, err
	}

	filter := query.GetDefaultIDFilter(sID)
	opt := bson.M{"$pull": bson.M{"recommend_menus": mID}}
	updateOne, err := s.collection.UpdateOne(ctx, filter, opt)
	if err != nil {
		return 0, err
	}
	return updateOne.ModifiedCount, nil
}
