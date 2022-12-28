package entity

import (
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Receipt 거래내역
type Receipt struct {
	Id         primitive.ObjectID   `bson:"id"`
	StoreId    primitive.ObjectID   `bson:"store_id"`
	CustomerId primitive.ObjectID   `bson:"customer_id"`
	Menu       []primitive.ObjectID `bson:"menu"`
	Price      int                  `bson:"price"`
	// Status 접수중/접수취소/추가접수/접수중/조리중/배달중/배달완료
	Status       string        `bson:"status"`
	CustomerAddr *dom.Address  `bson:"ordered_addr"`
	BaseTime     *dom.BaseTime `bson:"base_time"`
}
