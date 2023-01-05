package entity

import (
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	StoreID    primitive.ObjectID `bson:"store_id,omitempty"`
	CustomerID primitive.ObjectID `bson:"customer_id,omitempty"`
	MenuID     primitive.ObjectID `bson:"menu_id,omitempty"`
	OrderID    primitive.ObjectID `bson:"order_id,omitempty"`
	Content    string             `bson:"content,omitempty"`
	Rating     int                `bson:"rating,omitempty"`
	BaseTime   *dom.BaseTime      `bson:"base_time"`
}
