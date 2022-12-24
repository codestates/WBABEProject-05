package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store struct {
	Id             primitive.ObjectID
	UserId         primitive.ObjectID
	Address        *Address
	StorePhone     string
	Menu           []*Menu
	RecommendMenus []*Menu
	BaseTime       *BaseTime
}
