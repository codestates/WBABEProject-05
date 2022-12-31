package request

type RequestOrder struct {
	StoreId      string          `json:"store_id"`
	CustomerId   string          `json:"customer_id"`
	Price        int             `json:"price"`
	Menu         []string        `json:"menu_ids"`
	CustomerAddr *RequestAddress `json:"ordered_addr"`
}

//func (r *RequestOrder) ToReceipt() (*entity2.Receipt, error) {
//	sid, err := primitive.ObjectIDFromHex(r.StoreId)
//	if err != nil {
//		return nil, err
//	}
//	cid, err := primitive.ObjectIDFromHex(r.CustomerId)
//	if err != nil {
//		return nil, err
//	}
//	mids := make([]primitive.ObjectID, len(r.Menu))
//	for _, menuId := range r.Menu {
//		mid, err := primitive.ObjectIDFromHex(menuId)
//		if err != nil {
//			return nil, err
//		}
//		mids = append(mids, mid)
//	}
//
//	// TODO 생각해보니 총금액을 BE에서 해야하지 않나?! 이 API 도 고려해야될듯
//	rc := &entity2.Receipt{
//		Id:           primitive.NewObjectID(),
//		StoreId:      sid,
//		CustomerId:   cid,
//		Menu:         mids,
//		Price:        r.Price,
//		Status:       r.Status,
//		CustomerAddr: r.CustomerAddr.ToAddress(),
//		BaseTime: &dom.BaseTime{
//			CreatedAt: time.Now(),
//			UpdatedAt: time.Now(),
//		},
//	}
//	return rc, nil
//}
