package entity

import (
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	NicName     string             `bson:"nic_name"`
	Password    string             `bson:"password"`
	PhoneNumber string             `bson:"phone_number"`
	Role        string             `bson:"role"`
	BaseTime    *dom.BaseTime      `bson:"base_time"`
}
