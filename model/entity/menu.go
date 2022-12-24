package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Menu struct {
	Id   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
	// OrderCount 총 주문수
	OrderCount int `bson:"order_count"`
	// LimitCount 한정수량 ex) "non" , "1", "10"
	LimitCount string `bson:"limit_count"`
	Possible   bool   `bson:"possible"`
	Price      int    `bson:"price"`
	// Origin 원산지
	Origin      string   `bson:"origin"`
	Description string   `bson:"description"`
	BaseTime    BaseTime `bson:"base_time"`
}
