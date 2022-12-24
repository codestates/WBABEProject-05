package protocol

type RequestPostStore struct {
	UserId     string          `json:"user_id"`
	Name       string          `json:"name"`
	Address    *RequestAddress `json:"address"`
	StorePhone string          `json:"store_phone"`
}
