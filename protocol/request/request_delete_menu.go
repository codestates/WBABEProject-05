package request

type RequestDeleteMenu struct {
	MenuID  string `json:"menu-id" form:"menu-id" binding:"required"`
	StoreID string `json:"store-id " form:"store-id" binding:"required"`
	UserID  string `json:"user-id" form:"user-id" binding:"required"`
}
