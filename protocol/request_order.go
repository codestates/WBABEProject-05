package protocol

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RequestOrder struct {
	StoreId    string   `json:"store_id"`
	CustomerId string   `json:"customer_id"`
	Price      int      `json:"price"`
	Menu       []string `json:"menu_ids"`
	// Status 접수중/접수취소/추가접수/접수중/조리중/배달중/배달완료
	Status       string          `json:"status"`
	CustomerAddr *RequestAddress `json:"ordered_addr"`
}

func (r *RequestOrder) ToReceipt() (*entity.Receipt, error) {
	sid, err := primitive.ObjectIDFromHex(r.StoreId)
	if err != nil {
		return nil, err
	}
	cid, err := primitive.ObjectIDFromHex(r.CustomerId)
	if err != nil {
		return nil, err
	}
	mids := make([]primitive.ObjectID, len(r.Menu))
	for _, menuId := range r.Menu {
		mid, err := primitive.ObjectIDFromHex(menuId)
		if err != nil {
			return nil, err
		}
		mids = append(mids, mid)
	}

	// TODO 생각해보니 총금액을 BE에서 해야하지 않나?! 이 API 도 고려해야될듯
	rc := &entity.Receipt{
		Id:           primitive.NewObjectID(),
		StoreId:      sid,
		CustomerId:   cid,
		Menu:         mids,
		Price:        r.Price,
		Status:       r.Status,
		CustomerAddr: r.CustomerAddr.ToAddress(),
		BaseTime: &entity.BaseTime{
			Created_at: time.Now(),
			Updated_at: time.Now(),
		},
	}
	return rc, nil
}
