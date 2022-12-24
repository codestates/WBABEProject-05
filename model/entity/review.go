package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	Id      primitive.ObjectID
	StoreId primitive.ObjectID
	UserId  primitive.ObjectID
	Menu    Menu
	// Rating 1~5점
	Rating  int
	Content string
	BaseTime
}
