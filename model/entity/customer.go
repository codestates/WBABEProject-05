package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	Id          primitive.ObjectID
	User        User
	BaseAddress Address
	AddressList [5]Address
}
