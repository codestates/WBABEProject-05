package request

type RequestDeleteMenu struct {
	MenuID  string `json:"menu_id" form:"menu-id" binding:"required"`
	StoreID string `json:"store_id "form:"store-id" binding:"required"`
	UserID  string `json:"user_id" form:"user-id" binding:"required"`
}
