package request

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RequestUser struct {
	Name        string `json:"name" validate:"required"`
	NicName     string `json:"nic_name" validate:"required"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Role        string `json:"role"  validate:"required, eq=store|eq=user"`
}

func (r *RequestUser) NewPostUser() *entity.User {
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

func (r *RequestUser) NewUpdateUser(ID string) (*entity.User, error) {
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
