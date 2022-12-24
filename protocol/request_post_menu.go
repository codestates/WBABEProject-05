package protocol

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RequestPostMenu struct {
	StoreId     string `json:"store_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Origin      string `json:"origin" validate:"required"`
	Possible    bool   `json:"possible" validate:"required"`
	LimitCount  string `json:"limit_count,omitempty"`
	Description string `json:"description,omitempty"`
}

func (r *RequestPostMenu) ToStoreIdAndMenu() (primitive.ObjectID, *entity.Menu, error) {
	id, err := primitive.ObjectIDFromHex(r.StoreId)
	if err != nil {
		return primitive.ObjectID{}, nil, err
	}
	menu := &entity.Menu{
		Id:          primitive.NewObjectID(),
		Name:        r.Name,
		Price:       r.Price,
		Origin:      r.Origin,
		Possible:    r.Possible,
		LimitCount:  r.LimitCount,
		Description: r.Description,
		BaseTime: entity.BaseTime{
			Created_at: time.Now(),
			Updated_at: time.Now(),
		},
	}
	return id, menu, nil
}
