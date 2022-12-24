package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Menu struct {
	Id      primitive.ObjectID
	StoreId primitive.ObjectID
	Name    string
	// OrderCount 총 주문수
	OrderCount int
	// LimitCount 한정수량 ex) "non" , "1", "10"
	LimitCount string
	Price      int
	// Origin 원산지
	Origin   string
	BaseTime BaseTime
}
