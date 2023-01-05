package order

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/menu"
	"github.com/codestates/WBABEProject-05/model/receipt"
	"github.com/codestates/WBABEProject-05/model/user"
)

type orderRecordService struct {
	receiptModel receipt.ReceiptModeler
	menuModel    menu.MenuModeler
	userModel    user.UserModeler
}

var instance *orderRecordService

func NewOrderRecordService(rd receipt.ReceiptModeler, md menu.MenuModeler, ud user.UserModeler) *orderRecordService {
	if instance != nil {
		return instance
	}

	instance = &orderRecordService{receiptModel: rd, menuModel: md, userModel: ud}
	return instance
}

// command & query 둘다 사용
func (o *orderRecordService) sumMenusPrice(menus []*entity.Menu) int {
	var totalPrice int
	for _, m := range menus {
		totalPrice += m.Price
	}
	return totalPrice
}
