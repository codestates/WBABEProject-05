package request

type RequestPutCustomerOrder struct {
	ID           string          `json:"order_id" binding:"required"`
	StoreID      string          `json:"store_id" binding:"required"`
	CustomerID   string          `json:"customer_id" binding:"required"`
	MenuIDs      []string        `json:"menu_ids" binding:"required"`
	CustomerAddr *RequestAddress `json:"ordered_addr" binding:"required"`
	PhoneNumber  string          `json:"phone_number" binding:"required"`
}

func (r *RequestPutCustomerOrder) ToPutRequestOrder() *RequestOrder {
	return &RequestOrder{
		StoreID:      r.StoreID,
		CustomerID:   r.CustomerID,
		Menus:        r.MenuIDs,
		CustomerAddr: r.CustomerAddr,
		PhoneNumber:  r.PhoneNumber,
	}
}
