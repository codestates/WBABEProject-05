package entity

import (
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Status 접수중/접수취소/추가접수/접수중/조리중/배달중/배달완료
type Receipt struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty"`
	StoreID    primitive.ObjectID   `bson:"store_id,omitempty"`
	CustomerID primitive.ObjectID   `bson:"customer_id,omitempty"`
	Menu       []primitive.ObjectID `bson:"menu,omitempty"`
	Price      int                  `bson:"price,omitempty"`
	// Status 접수중/접수취소/추가접수/접수중/조리중/배달중/배달완료
	Status       string        `bson:"status,omitempty"`
	CustomerAddr *dom.Address  `bson:"ordered_addr"`
	BaseTime     *dom.BaseTime `bson:"base_time"`
}
