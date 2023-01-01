package request

type RequestPutCustomerOrder struct {
	ID           string          `json:"order_id" validate:"required"`
	StoreID      string          `json:"store_id" validate:"required"`
	CustomerID   string          `json:"customer_id" validate:"required"`
	MenuIDs      []string        `json:"menu_ids" validate:"required"`
	CustomerAddr *RequestAddress `json:"ordered_addr" validate:"required"`
}

func (r *RequestPutCustomerOrder) ToRequestOrder() *RequestOrder {
	return &RequestOrder{
		StoreId:      r.StoreID,
		CustomerId:   r.CustomerID,
		Menus:        r.MenuIDs,
		CustomerAddr: r.CustomerAddr,
	}
}
