package entity

import (
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	Id      primitive.ObjectID
	StoreId primitive.ObjectID
	UserId  primitive.ObjectID
	Menu    dom.Menu
	// Rating 1~5점
	Rating  int
	Content string
	dom.BaseTime
}
