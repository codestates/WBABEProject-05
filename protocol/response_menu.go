package protocol

import "github.com/codestates/WBABEProject-05/model/entity"

type ResponseMenu struct {
	Id          string `json:"menu_id"`
	Name        string `json:"name"`
	LimitCount  string `json:"limit_count"`
	Possible    bool   `json:"possible"`
	Price       int    `json:"price"`
	Origin      string `json:"origin"`
	Description string `json:"description"`
}

func NewResponseMenuFromMenu(menu entity.Menu) *ResponseMenu {
	return &ResponseMenu{
		Id:          menu.Id.Hex(),
		Name:        menu.Name,
		LimitCount:  menu.LimitCount,
		Possible:    menu.Possible,
		Price:       menu.Price,
		Origin:      menu.Origin,
		Description: menu.Description,
	}
}
