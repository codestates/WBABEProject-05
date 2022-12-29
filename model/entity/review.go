package entity

import (
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	Id      primitive.ObjectID `bson:"_id"`
	StoreId primitive.ObjectID `bson:"store_id"`
	UserId  primitive.ObjectID `bson:"user_id"`
	Menu    []*dom.Menu        `bson:"menu_ids"`
	// Rating 1~5점
	Rating   int          `bson:"rating"`
	Content  string       `bson:"content"`
	BaseTime dom.BaseTime `bson:"base_time"`
}
