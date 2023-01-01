package request

type RequestCheckPrice struct {
	StoreID string   `json:"store_id" form:"store-id"`
	Menus   []string `json:"menus" form:"menu-ids"`
}
