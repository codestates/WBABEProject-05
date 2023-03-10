package request

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RequestUser struct {
	Name        string `json:"name" binding:"required,min=2,max=15"`
	NicName     string `json:"nic_name" binding:"required,min=2,max=15"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Role        string `json:"role"  binding:"required,eq=store|eq=customer"`
}

func (r *RequestUser) ToPostUser() *entity.User {
	return &entity.User{
		ID:          primitive.NewObjectID(),
		Name:        r.Name,
		NicName:     r.NicName,
		Password:    r.Password,
		PhoneNumber: r.PhoneNumber,
		Role:        r.Role,
		BaseTime: &dom.BaseTime{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

func (r *RequestUser) ToPutUser(ID string) (*entity.User, error) {
	OBJID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:          OBJID,
		Name:        r.Name,
		NicName:     r.NicName,
		Password:    r.Password,
		PhoneNumber: r.PhoneNumber,
		Role:        r.Role,
	}, nil
}
