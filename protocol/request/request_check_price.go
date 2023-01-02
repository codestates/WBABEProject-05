package request

type RequestCheckPrice struct {
	StoreID string   `json:"store_id" form:"store-id" binding:"required"`
	Menus   []string `json:"menu_ids" form:"menu-ids" binding:"required"`
}
