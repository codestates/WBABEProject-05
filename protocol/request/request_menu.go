package request

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RequestMenu struct {
	UserID      string `json:"user_id" binding:"required"`
	StoreID     string `json:"store_id" binding:"required"`
	Name        string `json:"name" binding:"required,min=2,max=15"`
	Possible    bool   `json:"possible" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Origin      string `json:"origin" binding:"required"`
	Description string `json:"description" binding:"required,min=1,max=50"`
	LimitCount  string `json:"limit_count,omitempty"`
}

func (r *RequestMenu) ToPostMenu() (*entity.Menu, error) {
	sID, err := primitive.ObjectIDFromHex(r.StoreID)
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

func (r *RequestMenu) ToPutMenu(ID string) (*entity.Menu, error) {
	mID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	sID, err := primitive.ObjectIDFromHex(r.StoreID)
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
