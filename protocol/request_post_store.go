package protocol

type RequestPostStore struct {
	UserId         string          `json:"user_id"`
	Address        *RequestAddress `json:"address"`
	RecommendMenus []string        `json:"recommend_menus"`
	StorePhone     string          `json:"store_phone"`
}
