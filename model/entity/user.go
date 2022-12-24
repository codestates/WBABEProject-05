package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID
	Name        string
	NicName     string
	Password    string
	PhoneNumber string
	Role        string
	BaseTime    *BaseTime
}
