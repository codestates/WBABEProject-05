package request

type RequestCheckPrice struct {
	StoreID string   `json:"store-id" form:"store-id" binding:"required"`
	Menus   []string `json:"menu-ids" form:"menu-ids" binding:"required"`
}
