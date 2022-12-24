package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Store struct {
	Id            primitive.ObjectID
	User          User
	Address       Address
	Menu          []Menu
	RecommendMenu []Menu
	BaseTime      BaseTime
}

type BaseTime struct {
	Created_at time.Time
	Updated_at time.Time
}
