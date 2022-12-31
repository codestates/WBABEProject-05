package request

import (
	"github.com/codestates/WBABEProject-05/common/util"
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RequestOrder struct {
	StoreId      string          `json:"store_id"`
	CustomerId   string          `json:"customer_id"`
	Menus        []string        `json:"menu_ids"`
	CustomerAddr *RequestAddress `json:"ordered_addr"`
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

	objMIDs, err := util.ConvertStringsToObjIDs(r.Menus)
	if err != nil {
		return nil, err
	}

	rc := &entity.Receipt{
		ID:           primitive.NewObjectID(),
		StoreID:      sid,
		CustomerID:   cid,
		Menus:        objMIDs,
		Status:       entity.Waiting,
		CustomerAddr: r.CustomerAddr.ToAddress(),
		BaseTime: &dom.BaseTime{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	return rc, nil
}
