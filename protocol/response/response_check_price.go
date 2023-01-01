package response

import "github.com/codestates/WBABEProject-05/model/entity"

type ResponseCheckPrice struct {
	Menus      []*ResponseMenu `json:"menus"`
	TotalPrice int             `json:"total_price"`
}

func NewResponseCheckPrice(menus []*entity.Menu, price int) *ResponseCheckPrice {
	return &ResponseCheckPrice{
		Menus:      FromMenus(menus),
		TotalPrice: price,
	}
}
