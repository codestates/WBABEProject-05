package request

type RequestPutCustomerOrder struct {
	ID           string          `json:"order_id,omitempty"`
	StoreID      string          `json:"store_id,omitempty"`
	CustomerID   string          `json:"customer_id,omitempty"`
	Menus        []string        `json:"menu_ids,omitempty"`
	CustomerAddr *RequestAddress `json:"ordered_addr"`
}

func (r *RequestPutCustomerOrder) ToRequestOrder() *RequestOrder {
	return &RequestOrder{
		StoreId:      r.StoreID,
		CustomerId:   r.CustomerID,
		Menus:        r.Menus,
		CustomerAddr: r.CustomerAddr,
	}
}