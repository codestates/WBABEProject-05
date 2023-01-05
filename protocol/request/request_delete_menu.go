package request

type RequestDeleteMenu struct {
	MenuID  string `form:"menu-id" binding:"required"`
	StoreID string `form:"store-id" binding:"required"`
	UserID  string `form:"user-id" binding:"required"`
}
