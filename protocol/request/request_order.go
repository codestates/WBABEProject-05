package request

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"github.com/codestates/WBABEProject-05/model/enum"
	"github.com/codestates/WBABEProject-05/model/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RequestOrder struct {
	StoreId      string          `json:"store_id" binding:"required"`
	CustomerId   string          `json:"customer_id" binding:"required"`
	Menus        []string        `json:"menu_ids" binding:"required"`
	CustomerAddr *RequestAddress `json:"ordered_addr" binding:"required"`
}

func (r *RequestOrder) ToNewReceipt() (*entity.Receipt, error) {
	sid, err := primitive.ObjectIDFromHex(r.StoreId)
	if err != nil {
		return nil, err
	}
	cid, err := primitive.ObjectIDFromHex(r.CustomerId)
	if err != nil {
		return nil, err
	}

	OBJMIDs, err := util.ConvertStringsToOBJIDs(r.Menus)
	if err != nil {
		return nil, err
	}

	rc := &entity.Receipt{
		ID:           primitive.NewObjectID(),
		StoreID:      sid,
		CustomerID:   cid,
		Menus:        OBJMIDs,
		Status:       enum.Waiting,
		CustomerAddr: r.CustomerAddr.ToAddress(),
		BaseTime: &dom.BaseTime{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	return rc, nil
}
