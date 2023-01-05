package request

type RequestPutCustomerOrder struct {
	ID           string          `json:"order_id" binding:"required"`
	StoreID      string          `json:"store_id" binding:"required"`
	CustomerID   string          `json:"customer_id" binding:"required"`
	MenuIDs      []string        `json:"menu_ids" binding:"required"`
	CustomerAddr *RequestAddress `json:"ordered_addr" binding:"required"`
}

func (r *RequestPutCustomerOrder) ToPutRequestOrder() *RequestOrder {
	return &RequestOrder{
		StoreId:      r.StoreID,
		CustomerId:   r.CustomerID,
		Menus:        r.MenuIDs,
		CustomerAddr: r.CustomerAddr,
	}
}
