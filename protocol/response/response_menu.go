package response

import "github.com/codestates/WBABEProject-05/model/entity"

type ResponseMenu struct {
	Name        string  `json:"name" validate:"required"`
	LimitCount  string  `bson:"limit_count,omitempty"`
	Possible    bool    `bson:"possible,omitempty"`
	Price       int     `bson:"price,omitempty"`
	Origin      string  `bson:"origin,omitempty"`
	Description string  `bson:"description,omitempty"`
	Rating      float64 `bson:"rating,omitempty"`
	OrderCount  int     `bson:"order_count,omitempty"`
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
