package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store struct {
	Id            primitive.ObjectID
	User          *User
	Address       *Address
	Menu          []*Menu
	RecommendMenu []*Menu
	BaseTime      *BaseTime
}
