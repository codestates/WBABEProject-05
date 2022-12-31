package entity

import (
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	Waiting       = "주문대기"
	Cancel        = "주문취소"
	AddOrder      = "추가주문"
	OrderReceived = "주문접수완료"
	Cooking       = "조리중"
	Delivering    = "배달중"
	Completion    = "배달완료"
)

// Status 접수중/접수취소/추가접수/접수중/조리중/배달중/배달완료
type Receipt struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty"`
	StoreID    primitive.ObjectID   `bson:"store_id,omitempty"`
	CustomerID primitive.ObjectID   `bson:"customer_id,omitempty"`
	Menus      []primitive.ObjectID `bson:"menu,omitempty"`
	Price      int                  `bson:"price,omitempty"`
	// Status 접수중/접수취소/추가접수/접수중/조리중/배달중/배달완료
	Status       string        `bson:"status,omitempty"`
	CustomerAddr *dom.Address  `bson:"ordered_addr"`
	Numbering    string        `bson:"numbering"`
	BaseTime     *dom.BaseTime `bson:"base_time"`
}

func (s *Receipt) NewUpdateStatusOrderBsonSetD() bson.D {
	return bson.D{
		{"$set",
			bson.D{
				{"status", s.Status},
				{"base_time.updated_at", time.Now()},
			},
		},
	}
}

func (s *Receipt) NewUpdateOrderCancelBsonSetD() bson.D {
	return bson.D{
		{"$set",
			bson.D{
				{"Status", Cancel},
				{"base_time.updated_at", time.Now()},
			},
		},
	}
}
