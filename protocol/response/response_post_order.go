package response

import "github.com/codestates/WBABEProject-05/model/entity"

type ResponsePostOrder struct {
	ID        string `json:"order_id"`
	Numbering string `json:"order_numbering"`
}

func FromReceipt(order *entity.Receipt) *ResponsePostOrder {
	return &ResponsePostOrder{
		ID:        order.ID.Hex(),
		Numbering: order.Numbering,
	}
}
