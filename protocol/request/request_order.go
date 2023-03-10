package request

import (
	"github.com/codestates/WBABEProject-05/common/convertor"
	"github.com/codestates/WBABEProject-05/common/enum"
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RequestOrder struct {
	StoreID      string          `json:"store_id" binding:"required"`
	CustomerID   string          `json:"customer_id" binding:"required"`
	Menus        []string        `json:"menu_ids" binding:"required"`
	CustomerAddr *RequestAddress `json:"ordered_addr" binding:"required"`
	PhoneNumber  string          `json:"phone_number" binding:"required"`
}

func (r *RequestOrder) ToPostReceipt() (*entity.Receipt, error) {
	sid, err := primitive.ObjectIDFromHex(r.StoreID)
	if err != nil {
		return nil, err
	}
	cid, err := primitive.ObjectIDFromHex(r.CustomerID)
	if err != nil {
		return nil, err
	}

	OBJMIDs, err := convertor.ConvertStringsToOBJIDs(r.Menus)
	if err != nil {
		return nil, err
	}

	return &entity.Receipt{
		ID:            primitive.NewObjectID(),
		StoreID:       sid,
		CustomerID:    cid,
		MenuIDs:       OBJMIDs,
		Status:        enum.Waiting,
		CustomerAddr:  r.CustomerAddr.ToAddress(),
		CustomerPhone: r.PhoneNumber,
		BaseTime: &dom.BaseTime{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}, nil
}

func (r *RequestOrder) ToUserPreOrderInfo() (*entity.User, error) {
	cid, err := primitive.ObjectIDFromHex(r.CustomerID)
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID: cid,
		PreOrderInfo: &dom.PreOrderInfo{
			Address:     r.CustomerAddr.ToAddress(),
			PhoneNumber: r.PhoneNumber,
		},
	}, nil
}
