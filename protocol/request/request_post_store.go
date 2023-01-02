package request

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RequestPostStore struct {
	UserID     string          `json:"user_id" binding:"required"`
	Name       string          `json:"name" binding:"required,min=2,max=15"`
	Address    *RequestAddress `json:"address" binding:"required"`
	StorePhone string          `json:"store_phone" binding:"required"`
}

func (r *RequestPostStore) NewPostStore() (*entity.Store, error) {
	userID, err := primitive.ObjectIDFromHex(r.UserID)
	if err != nil {
		return nil, err
	}
	return &entity.Store{
		ID:         primitive.NewObjectID(),
		UserID:     userID,
		Name:       r.Name,
		Address:    r.Address.ToAddress(),
		StorePhone: r.StorePhone,
		BaseTime: &dom.BaseTime{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}, err
}
