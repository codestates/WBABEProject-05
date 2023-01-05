package order

import (
	"fmt"
	"github.com/codestates/WBABEProject-05/common/enum"
	error2 "github.com/codestates/WBABEProject-05/common/error"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/codestates/WBABEProject-05/model/common"
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/protocol/request"
	util2 "github.com/codestates/WBABEProject-05/service/common"
	"sync"
	"time"
)

func (o *orderRecordService) RegisterOrderRecord(order *request.RequestOrder) (string, error) {
	rct, err := order.ToNewReceipt()
	if err != nil {
		return enum.BlankSTR, err
	}

	// 이전 주문 정보 저장도 비즈니스상 중요하지 않아보여 따로 컨트롤하지 않는 고루틴 처리
	go o.updateUserPreOrderInfo(order)

	if err := o.setTotalPriceAndNumbering(rct, order); err != nil {
		return enum.BlankSTR, err
	}

	savedNumbering, err := o.receiptModel.InsertReceipt(rct)
	if err != nil {
		return enum.BlankSTR, err
	}

	// OrderCount 의 증가는 비즈니스상 중요하지않아보여 따로 컨틀롤하지 않는 고루틴 활용
	go o.updateOrderCount(order)

	return savedNumbering, nil
}

func (o *orderRecordService) ModifyOrderRecordFromCustomer(order *request.RequestPutCustomerOrder) (string, error) {
	foundOrder, err := o.receiptModel.SelectReceiptByID(order.ID)
	if err != nil {
		return enum.BlankSTR, err
	}

	// 메뉴 추가의 경우 배달중은 실패 , 메뉴 변경의 경우 조리중,배달중인경우 실패 -> 즉, 기본으로 배달중은 실패, 추가적으로 완료도 실패시키자
	if err := o.checkOrderStatus(order, foundOrder); err != nil {
		return enum.BlankSTR, err
	}

	if common.ConvertOBJIDToString(foundOrder.CustomerID) != order.CustomerID {
		return enum.BlankSTR, error2.BadAccessOrderError.New()
	}

	if _, err := o.receiptModel.UpdateCancelReceipt(foundOrder); err != nil {
		return enum.BlankSTR, err
	}

	savedID, err := o.RegisterOrderRecord(order.ToRequestOrder())
	if err != nil {
		return enum.BlankSTR, err
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

func (o *orderRecordService) setTotalPriceAndNumbering(rct *entity.Receipt, order *request.RequestOrder) error {
	var wg sync.WaitGroup
	wg.Add(2)

	// setTotalPrice
	findMenusErrCH := make(chan error, 1)
	go o.setTotalPrice(rct, order, findMenusErrCH, &wg)

	// setNumbering
	findTotalCountErrCH := make(chan error, 1)
	go o.setNumbering(&wg, findTotalCountErrCH, rct)

	wg.Wait()

	if err := <-findMenusErrCH; err != nil {
		return err
	}
	if err := <-findTotalCountErrCH; err != nil {
		return err
	}
	return nil
}

func (o *orderRecordService) setNumbering(wg *sync.WaitGroup, countErr chan error, rct *entity.Receipt) {
	defer wg.Done()
	toDayCnt, err := o.receiptModel.SelectToDayTotalCount()
	if err != nil {
		countErr <- err
		return
	}
	rct.Numbering = o.newNumbering(toDayCnt)
	countErr <- nil
}

func (o *orderRecordService) setTotalPrice(rct *entity.Receipt, order *request.RequestOrder, menusErr chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	menus, err := o.menuModel.SelectMenusByIDs(order.StoreId, order.Menus)
	if err != nil {
		menusErr <- err
		return
	}
	rct.Price = o.sumMenusPrice(menus)
	menusErr <- nil
}

func (o *orderRecordService) checkOrderStatus(order *request.RequestPutCustomerOrder, foundOrder *entity.Receipt) error {
	reqMIDs := util2.ConvertSliceToExistMap(order.MenuIDs)
	isChange := o.isChangeOrderMenus(foundOrder, reqMIDs)
	switch {
	case isChange:
		if foundOrder.Status == enum.Cooking {
			return error2.DoseNotModifyOrderError.New()
		}
	case foundOrder.Status == enum.Completion || foundOrder.Status == enum.Delivering:
		return error2.DoseNotModifyOrderError.New()
	}
	return nil
}

func (o *orderRecordService) isChangeOrderMenus(foundOrder *entity.Receipt, reqMIDs map[string]int) bool {
	for _, ID := range foundOrder.MenuIDs {
		// false -> 변경으로 볼 수 있다.
		if _, exist := reqMIDs[ID.Hex()]; !exist {
			return true
		}
	}
	return false
}

func (o *orderRecordService) updateUserPreOrderInfo(order *request.RequestOrder) {
	preOrderInfo, err := order.ToUserPreOrderInfo()
	if err != nil {
		logger.AppLog.Error(err)
	}
	_, err = o.userModel.UpdateUserPreOrder(preOrderInfo)
	if err != nil {
		logger.AppLog.Error(err)
	}
}

func (o *orderRecordService) updateOrderCount(order *request.RequestOrder) {
	count, err := o.menuModel.UpdateMenusInCOrderCount(order.Menus)
	if err != nil || count == 0 {
		MSG := fmt.Sprintf("does not update order count Menu IDs %v", order.Menus)
		logger.AppLog.Error(MSG)
	}
}

func (o *orderRecordService) newNumbering(toDayCnt int64) string {
	return fmt.Sprintf("%d-%d", time.Now().UnixNano(), toDayCnt)
}
