package request

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RequestMenu struct {
	StoreId        string  `json:"store_id" validate:"required"`
	Name           string  `json:"name" validate:"required"`
	LimitCount     string  `bson:"limit_count,omitempty"`
	Possible       bool    `bson:"possible,omitempty"`
	Price          int     `bson:"price,omitempty"`
	Origin         string  `bson:"origin,omitempty"`
	Description    string  `bson:"description,omitempty"`
	RecommendCount int     `bson:"recommend_count,omitempty"`
	Rating         float64 `bson:"rating,omitempty"`
	ReOrderCount   int     `bson:"re_order_count,omitempty"`
}

func (r *RequestMenu) NewMenu() (*entity.Menu, error) {
	sID, err := primitive.ObjectIDFromHex(r.StoreId)
	if err != nil {
		return nil, err
	}

	return &entity.Menu{
		ID:             primitive.NewObjectID(),
		StoreID:        sID,
		Name:           r.Name,
		LimitCount:     r.LimitCount,
		Possible:       r.Possible,
		Price:          r.Price,
		Origin:         r.Origin,
		Description:    r.Description,
		RecommendCount: r.RecommendCount,
		Rating:         r.Rating,
		ReOrderCount:   r.ReOrderCount,
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
		ID:             mID,
		StoreID:        sID,
		Name:           r.Name,
		LimitCount:     r.LimitCount,
		Possible:       r.Possible,
		Price:          r.Price,
		Origin:         r.Origin,
		Description:    r.Description,
		RecommendCount: r.RecommendCount,
		Rating:         r.Rating,
		ReOrderCount:   r.ReOrderCount,
	}, nil
}
