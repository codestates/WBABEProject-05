package request

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"github.com/codestates/WBABEProject-05/model/util"
	"github.com/codestates/WBABEProject-05/model/util/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RequestOrder struct {
	StoreId      string          `json:"store_id" validate:"required"`
	CustomerId   string          `json:"customer_id" validate:"required"`
	Menus        []string        `json:"menu_ids" validate:"required"`
	CustomerAddr *RequestAddress `json:"ordered_addr" validate:"required"`
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
