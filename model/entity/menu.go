package entity

import (
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Menu struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	StoreID        primitive.ObjectID `bson:"store_id"`
	Name           string             `bson:"name,omitempty"`
	LimitCount     string             `bson:"limit_count,omitempty"`
	Possible       bool               `bson:"possible,omitempty"`
	Price          int                `bson:"price,omitempty"`
	Origin         string             `bson:"origin,omitempty"`
	Description    string             `bson:"description,omitempty"`
	RecommendCount int                `bson:"recommend_count,omitempty"`
	Rating         float64            `bson:"rating,omitempty"`
	ReOrderCount   int                `bson:"re_order_count,omitempty"`
	BaseTime       *dom.BaseTime      `bson:"base_time"`
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
				{"recommend_count", m.RecommendCount},
				{"rating", m.Rating},
				{"re_order_count", m.ReOrderCount},
				{"base_time.updated_at", time.Now()},
			},
		},
	}
}
