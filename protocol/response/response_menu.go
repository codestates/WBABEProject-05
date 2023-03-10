package response

import "github.com/codestates/WBABEProject-05/model/entity"

type ResponseMenu struct {
	ID          string  `json:"menu_id"`
	Name        string  `json:"name"`
	LimitCount  string  `json:"limit_count"`
	Possible    bool    `json:"possible"`
	Price       int     `json:"price"`
	Origin      string  `json:"origin"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	OrderCount  int     `json:"order_count"`
}

func FromMenus(menus []*entity.Menu) []*ResponseMenu {
	var res []*ResponseMenu
	for _, menu := range menus {
		res = append(res, FromMenu(menu))
	}
	return res
}

func FromMenu(menu *entity.Menu) *ResponseMenu {
	return &ResponseMenu{
		ID:          menu.ID.Hex(),
		Name:        menu.Name,
		LimitCount:  menu.LimitCount,
		Possible:    menu.Possible,
		Price:       menu.Price,
		Origin:      menu.Origin,
		Description: menu.Description,
		Rating:      menu.Rating,
		OrderCount:  menu.OrderCount,
	}
}
