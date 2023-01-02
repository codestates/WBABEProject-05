package request

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RequestPostReview struct {
	StoreID    string `json:"store_id" validate:"required"`
	CustomerID string `json:"customer_id" validate:"required"`
	MenuID     string `json:"menu_id" validate:"required"`
	Content    string `json:"content" validate:"required, min=5, max=100"`
	Rating     int    `json:"rating" validate:"required, min=1, max=5"`
}

func (r *RequestPostReview) NewReview() (*entity.Review, error) {
	sID, err := primitive.ObjectIDFromHex(r.StoreID)
	if err != nil {
		return nil, err
	}

	cID, err := primitive.ObjectIDFromHex(r.CustomerID)
	if err != nil {
		return nil, err
	}

	mID, err := primitive.ObjectIDFromHex((r.MenuID))

	if err != nil {
		return nil, err
	}

	return &entity.Review{
		ID:         primitive.NewObjectID(),
		StoreID:    sID,
		CustomerID: cID,
		MenuID:     mID,
		Content:    r.Content,
		Rating:     r.Rating,
		BaseTime: &dom.BaseTime{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}, nil
}
