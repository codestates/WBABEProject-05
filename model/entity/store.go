package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store struct {
	Id             primitive.ObjectID `bson:"_id"`
	UserId         primitive.ObjectID `bson:"user_id"`
	Address        *Address           `bson:"address"`
	StorePhone     string             `bson:"store_phone"`
	Menu           []*Menu            `bson:"menu"`
	RecommendMenus []*Menu            `bson:"recommend_menus"`
	BaseTime       *BaseTime          `bson:"base_time"`
}
