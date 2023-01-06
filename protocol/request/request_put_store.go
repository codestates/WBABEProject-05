package request

import (
	"github.com/codestates/WBABEProject-05/common/convertor"
	"github.com/codestates/WBABEProject-05/model/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RequestPutStore struct {
	UserID         string          `json:"user_id" binding:"required"`
	Name           string          `json:"name" binding:"required,min=2,max=15"`
	Address        *RequestAddress `json:"address" binding:"required"`
	StorePhone     string          `json:"store_phone" binding:"required"`
	RecommendMenus []string        `json:"recommend_menus,omitempty"`
}

func (r *RequestPutStore) ToPutStore(storeID string) (*entity.Store, error) {
	sOBJID, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, err
	}

	uObjID, err := primitive.ObjectIDFromHex(r.UserID)
	if err != nil {
		return nil, err
	}

	rsMIDS, err := convertor.ConvertStringsToOBJIDs(r.RecommendMenus)
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
