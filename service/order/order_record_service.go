package order

import (
	"errors"
	"fmt"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/menu"
	"github.com/codestates/WBABEProject-05/model/receipt"
	"github.com/codestates/WBABEProject-05/model/util"
	error2 "github.com/codestates/WBABEProject-05/protocol/error"
	"github.com/codestates/WBABEProject-05/protocol/page"
	"github.com/codestates/WBABEProject-05/protocol/request"
	"github.com/codestates/WBABEProject-05/protocol/response"
	"time"
)

type orderRecordService struct {
	receiptModel receipt.ReceiptModeler
	menuModel    menu.MenuModeler
}

var instance *orderRecordService

func NewOrderRecordService(rd receipt.ReceiptModeler, md menu.MenuModeler) *orderRecordService {
	if instance != nil {
		return instance
	}

	instance = &orderRecordService{receiptModel: rd, menuModel: md}
	return instance
}

func (o *orderRecordService) RegisterOrderRecord(order *request.RequestOrder) (string, error) {
	rct, err := order.ToNewReceipt()
	if err != nil {
		return "", err
	}

	menus, err := o.menuModel.SelectMenusByIDs(order.StoreId, order.Menus)
	if err != nil {
		return "", err
	}

	rct.Price = o.sumMenusPrice(menus)

	toDayCnt, err := o.receiptModel.SelectToDayTotalCount()
	if err != nil {
		return "", err
	}
	rct.Numbering = o.newNumbering(toDayCnt)

	insertedId, err := o.receiptModel.InsertReceipt(rct)
	if err != nil {
		return "", err
	}

	// OrderCount 의 증가는 비즈니스상 중요하지않아 고루틴 활용
	go func() {
		count, err := o.menuModel.UpdateMenusInCOrderCount(order.Menus)
		if err != nil || count == 0 {
			msg := fmt.Sprintf("does not update order count Menu IDs %v", order.Menus)
			logger.AppLog.Error(errors.New(msg))
		}
	}()

	return insertedId, nil
}

func (o *orderRecordService) ModifyOrderRecordFromCustomer(order *request.RequestPutCustomerOrder) (string, error) {
	foundOrder, err := o.receiptModel.SelectReceiptByID(order.ID)
	if err != nil {
		return "", err
	}
	if foundOrder.Status != entity.Waiting {
		return "", error2.AlreadyReceivedOrderError.New()
	}

	if util.ConvertOBJIDToString(foundOrder.CustomerID) != order.CustomerID {
		return "", error2.BadAccessOrderError.New()
	}

	if _, err := o.receiptModel.UpdateCancelReceipt(foundOrder); err != nil {
		return "", err
	}

	savedID, err := o.RegisterOrderRecord(order.ToRequestOrder())
	if err != nil {
		return "", err
	}

	return savedID, nil
}
func (o *orderRecordService) ModifyOrderRecordFromStore(order *request.RequestPutStoreOrder) (int, error) {
	foundOrder, err := o.receiptModel.SelectReceiptByID(order.ID)
	if err != nil {
		return 0, err
	}

	foundOrder.Status = order.Status

	updatedCnt, err := o.receiptModel.UpdateReceiptStatus(foundOrder)
	if err != nil {
		return 0, err
	}

	return int(updatedCnt), nil
}

func (o *orderRecordService) FindOrderRecordsSortedPage(ID, userRole string, pg *request.RequestPage) (*page.PageData[any], error) {
	skip := pg.CurrentPage * pg.ContentCount

	receipts, err := o.receiptModel.SelectSortLimitedReceipt(ID, userRole, pg.Sort, skip, pg.ContentCount)
	if err != nil {
		return nil, err
	}

	totalCount, err := o.receiptModel.SelectTotalCount(ID, userRole)
	if err != nil {
		return nil, err
	}

	pgInfo := pg.NewPageInfo(int(totalCount))

	return page.NewPageData(receipts, pgInfo), nil
}

func (o *orderRecordService) FindOrderRecord(orderID string) (*response.ResponseOrder, error) {
	foundReceipt, err := o.receiptModel.SelectReceiptByID(orderID)
	if err != nil {
		return nil, err
	}

	menuIDs := util.ConvertOBJIDsToStrings(foundReceipt.Menus)

	menus, err := o.menuModel.SelectMenusByIDs(foundReceipt.StoreID.Hex(), menuIDs)
	if err != nil {
		return nil, err
	}

	resOrder := response.FromReceiptAndMenus(foundReceipt, menus)

	return resOrder, nil
}

func (o *orderRecordService) FiendSelectedMenusTotalPrice(storeID string, menuIDs []string) (*response.ResponseCheckPrice, error) {
	menus, err := o.menuModel.SelectMenusByIDs(storeID, menuIDs)
	if err != nil {
		return nil, err
	}

	totalPrice := o.sumMenusPrice(menus)

	resCheckPrice := response.NewResponseCheckPrice(menus, totalPrice)

	return resCheckPrice, nil
}

func (o *orderRecordService) sumMenusPrice(menus []*entity.Menu) int {
	var totalPrice int
	for _, m := range menus {
		totalPrice += m.Price
	}
	return totalPrice
}

func (o *orderRecordService) newNumbering(toDayCnt int64) string {
	return fmt.Sprintf("%d-%d", time.Now().UnixNano(), toDayCnt)
}
