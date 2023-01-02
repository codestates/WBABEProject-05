package request

type RequestCheckPrice struct {
	StoreID string   `json:"store_id" form:"store-id" validate:"required"`
	Menus   []string `json:"menu_ids" form:"menu-ids" validate:"required"`
}
