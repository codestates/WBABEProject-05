package entity

import (
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store struct {
	Id             primitive.ObjectID `bson:"_id"`
	UserId         primitive.ObjectID `bson:"user_id"`
	Name           string             `bson:"name"`
	Address        *dom.Address       `bson:"address"`
	StorePhone     string             `bson:"store_phone"`
	Menu           []*dom.Menu        `bson:"menu,omitempty"`
	RecommendMenus []*dom.Menu        `bson:"recommend_menus,omitempty"`
	BaseTime       *dom.BaseTime      `bson:"base_time"`
}
