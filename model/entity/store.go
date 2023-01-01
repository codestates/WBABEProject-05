package entity

import (
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Store struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty"`
	UserID         primitive.ObjectID   `bson:"user_id,omitempty"`
	Name           string               `bson:"name,omitempty"`
	Address        *dom.Address         `bson:"address"`
	StorePhone     string               `bson:"store_phone,omitempty"`
	RecommendMenus []primitive.ObjectID `bson:"recommend_menus,omitempty"`
	BaseTime       *dom.BaseTime        `bson:"base_time"`
}

func (s *Store) NewUpdateStoreBsonSetD() bson.D {
	return bson.D{
		{"$set",
			bson.D{
				{"name", s.Name},
				{"address", &dom.Address{
					Street:  s.Address.Street,
					Detail:  s.Address.Detail,
					ZipCode: s.Address.ZipCode,
				}},
				{"store_phone", s.StorePhone},
				{"recommend_menus", s.RecommendMenus},
				{"base_time.updated_at", time.Now()},
			},
		},
	}
}
