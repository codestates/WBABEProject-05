package response

import (
	"github.com/codestates/WBABEProject-05/model/entity"
)

type ResponseOrder struct {
	ID           string            `json:"_id,omitempty"`
	StoreID      string            `json:"store_id,omitempty"`
	CustomerID   string            `json:"customer_id,omitempty"`
	Menu         []*ResponseMenu   `json:"menu,omitempty"`
	Price        int               `json:"price,omitempty"`
	Status       string            `json:"status,omitempty"`
	CustomerAddr *ResponseAddress  `json:"ordered_addr"`
	Numbering    string            `json:"numbering"`
	BaseTime     *ResponseBaseTime `json:"base_time"`
}

func FromReceiptAndMenus(receipt *entity.Receipt, menus []*entity.Menu) *ResponseOrder {
	return &ResponseOrder{
		ID:           receipt.ID.Hex(),
		StoreID:      receipt.StoreID.Hex(),
		CustomerID:   receipt.CustomerID.Hex(),
		Menu:         FromMenus(menus),
		Price:        receipt.Price,
		Status:       receipt.Status,
		CustomerAddr: FromAddr(receipt.CustomerAddr),
		Numbering:    receipt.Numbering,
		BaseTime:     FromBaseTime(receipt.BaseTime),
	}
}
