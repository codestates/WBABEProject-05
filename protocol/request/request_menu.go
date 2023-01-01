package request

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RequestMenu struct {
	StoreId     string `json:"store_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Possible    bool   `json:"possible" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Origin      string `json:"origin" validate:"required"`
	Description string `json:"description" validate:"required"`
	LimitCount  string `json:"limit_count"`
}

func (r *RequestMenu) NewMenu() (*entity.Menu, error) {
	sID, err := primitive.ObjectIDFromHex(r.StoreId)
	if err != nil {
		return nil, err
	}

	return &entity.Menu{
		ID:          primitive.NewObjectID(),
		StoreID:     sID,
		Name:        r.Name,
		LimitCount:  r.LimitCount,
		Possible:    r.Possible,
		Price:       r.Price,
		Origin:      r.Origin,
		Description: r.Description,
		BaseTime: &dom.BaseTime{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}, nil
}

func (r *RequestMenu) NewUpdateMenu(ID string) (*entity.Menu, error) {
	mID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	sID, err := primitive.ObjectIDFromHex(r.StoreId)
	if err != nil {
		return nil, err
	}
	return &entity.Menu{
		ID:          mID,
		StoreID:     sID,
		Name:        r.Name,
		LimitCount:  r.LimitCount,
		Possible:    r.Possible,
		Price:       r.Price,
		Origin:      r.Origin,
		Description: r.Description,
	}, nil
}
