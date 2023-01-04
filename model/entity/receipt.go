package entity

import (
	"github.com/codestates/WBABEProject-05/common/enum"
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Receipt struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty"`
	StoreID       primitive.ObjectID   `bson:"store_id,omitempty"`
	CustomerID    primitive.ObjectID   `bson:"customer_id,omitempty"`
	MenuIDs       []primitive.ObjectID `bson:"menu,omitempty"`
	Price         int                  `bson:"price,omitempty"`
	Status        string               `bson:"status,omitempty"`
	CustomerAddr  *dom.Address         `bson:"ordered_addr"`
	CustomerPhone string               `bson:"customer_phone,omitempty"`
	Numbering     string               `bson:"numbering,omitempty"`
	BaseTime      *dom.BaseTime        `bson:"base_time"`
}

func (s *Receipt) NewUpdateStatusBsonSetD() bson.D {
	return bson.D{
		{"$set",
			bson.D{
				{"status", s.Status},
				{"base_time.updated_at", time.Now()},
			},
		},
	}
}

func (s *Receipt) NewUpdateStatusCancelBsonSetD() bson.D {
	return bson.D{
		{"$set",
			bson.D{
				{"Status", enum.Cancel},
				{"base_time.updated_at", time.Now()},
			},
		},
	}
}
