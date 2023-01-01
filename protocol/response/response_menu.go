package response

import "github.com/codestates/WBABEProject-05/model/entity"

type ResponseMenu struct {
	ID          string  `json:"menu_id"`
	Name        string  `json:"name"`
	LimitCount  string  `json:"limit_count,omitempty"`
	Possible    bool    `json:"possible,omitempty"`
	Price       int     `json:"price,omitempty"`
	Origin      string  `json:"origin,omitempty"`
	Description string  `json:"description,omitempty"`
	Rating      float64 `json:"rating,omitempty"`
	OrderCount  int     `json:"order_count,omitempty"`
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
