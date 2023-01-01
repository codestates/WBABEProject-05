package request

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RequestPutStore struct {
	UserID         string          `json:"user_id" validate:"required"`
	Name           string          `json:"name" validate:"required"`
	Address        *RequestAddress `json:"address" validate:"required"`
	StorePhone     string          `json:"store_phone" validate:"required"`
	RecommendMenus []string        `json:"recommend_menus"`
}

func (r *RequestPutStore) NewPutStore(storeID string) (*entity.Store, error) {
	sOBJID, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, err
	}

	uObjID, err := primitive.ObjectIDFromHex(r.UserID)
	if err != nil {
		return nil, err
	}

	rsMIDS, err := util.ConvertStringsToOBJIDs(r.RecommendMenus)
	if err != nil {
		return nil, err
	}

	return &entity.Store{
		ID:             sOBJID,
		UserID:         uObjID,
		Name:           r.Name,
		Address:        r.Address.ToAddress(),
		StorePhone:     r.StorePhone,
		RecommendMenus: rsMIDS,
	}, nil
}
