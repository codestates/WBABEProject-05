package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

// Receipt 거래내역
type Receipt struct {
	Id         primitive.ObjectID
	StoreId    primitive.ObjectID
	CustomerId primitive.ObjectID
	Price      int
	Menu       []Menu
	// Status 접수중/접수취소/추가접수/접수중/조리중/배달중/배달완료
	Status   string
	BaseTime BaseTime
}
