package request

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RequestPostReview struct {
	StoreID    string `json:"store_id" binding:"required"`
	CustomerID string `json:"customer_id" binding:"required"`
	MenuID     string `json:"menu_id" binding:"required"`
	OrderID    string `json:"order_id" binding:"required"`
	Content    string `json:"content" binding:"required,min=5,max=100"`
	Rating     int    `json:"rating" binding:"required,min=1,max=5"`
}

func (r *RequestPostReview) ToPostReview() (*entity.Review, error) {
	sID, err := primitive.ObjectIDFromHex(r.StoreID)
	if err != nil {
		return nil, err
	}

	cID, err := primitive.ObjectIDFromHex(r.CustomerID)
	if err != nil {
		return nil, err
	}

	mID, err := primitive.ObjectIDFromHex(r.MenuID)
	if err != nil {
		return nil, err
	}

	oID, err := primitive.ObjectIDFromHex(r.OrderID)
	if err != nil {
		return nil, err
	}

	return &entity.Review{
		ID:         primitive.NewObjectID(),
		StoreID:    sID,
		CustomerID: cID,
		MenuID:     mID,
		OrderID:    oID,
		Content:    r.Content,
		Rating:     r.Rating,
		BaseTime: &dom.BaseTime{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}, nil
}
