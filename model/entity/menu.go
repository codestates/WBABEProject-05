package entity

import (
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Menu struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	StoreID          primitive.ObjectID `bson:"store_id,omitempty"`
	Name             string             `bson:"name,omitempty"`
	LimitCount       string             `bson:"limit_count,omitempty"`
	Possible         bool               `bson:"possible,omitempty"`
	Price            int                `bson:"price,omitempty"`
	Origin           string             `bson:"origin,omitempty"`
	Description      string             `bson:"description,omitempty"`
	Rating           float64            `bson:"rating,omitempty"`
	OrderCount       int                `bson:"order_count,omitempty"`
	ReviewCount      int                `bson:"review_count,omitempty"`
	TotalReviewScore int                `bson:"total_review_score,omitempty"`
	BaseTime         *dom.BaseTime      `bson:"base_time"`
}

func (m *Menu) NewUpdateMenuBsonSetD() bson.D {
	return bson.D{
		{"$set",
			bson.D{
				{"name", m.Name},
				{"limit_count", m.LimitCount},
				{"possible", m.Possible},
				{"price", m.Price},
				{"origin", m.Origin},
				{"description", m.Description},
				{"base_time.updated_at", time.Now()},
			},
		},
	}
}

func (m *Menu) NewUpdateMenuBsonSetDAboutReview() bson.D {
	return bson.D{
		{"$set",
			bson.D{
				{"rating", m.Rating},
				{"review_count", m.ReviewCount},
				{"total_review_score", m.TotalReviewScore},
				{"base_time.updated_at", time.Now()},
			},
		},
	}
}
