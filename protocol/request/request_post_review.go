package request

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RequestPostReview struct {
	StoreID    string `json:"store_id,omitempty"`
	CustomerId string `json:"customer_id,omitempty"`
	Menu       string `json:"menu_id,omitempty"`
	Content    string `json:"content,omitempty"`
	Rating     int    `json:"rating" validate:"required, min=0, max=5"`
}

func (r *RequestPostReview) NewReview() (*entity.Review, error) {
	sID, err := primitive.ObjectIDFromHex(r.StoreID)
	if err != nil {
		return nil, err
	}

	cID, err := primitive.ObjectIDFromHex(r.CustomerId)
	if err != nil {
		return nil, err
	}

	mID, err := primitive.ObjectIDFromHex((r.Menu))

	if err != nil {
		return nil, err
	}

	return &entity.Review{
		ID:         primitive.NewObjectID(),
		StoreID:    sID,
		CustomerID: cID,
		Menu:       mID,
		Content:    r.Content,
		Rating:     r.Rating,
		BaseTime: &dom.BaseTime{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}, nil
}
