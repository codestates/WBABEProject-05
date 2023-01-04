package entity

import (
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name,omitempty"`
	NicName      string             `bson:"nic_name,omitempty"`
	Password     string             `bson:"password,omitempty"`
	PhoneNumber  string             `bson:"phone_number,omitempty"`
	Role         string             `bson:"role,omitempty"`
	PreOrderInfo *dom.PreOrderInfo  `bson:"pre_order_info"`
	BaseTime     *dom.BaseTime      `bson:"base_time"`
}

func (u *User) NewBsonSetDForUpdateUser() bson.D {
	return bson.D{
		{"$set",
			bson.D{
				{"name", u.Name},
				{"nic_name", u.NicName},
				{"phone_number", u.PhoneNumber},
				{"role", u.Role},
				{"base_time.updated_at", time.Now()},
			},
		},
	}
}

func (u *User) NewBsonSetDForUpdatePreOrder() bson.D {
	return bson.D{
		{"$set",
			bson.D{
				{"pre_order_info", u.PreOrderInfo},
				{"base_time.updated_at", time.Now()},
			},
		},
	}
}
