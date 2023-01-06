package order

import (
	"fmt"
	"github.com/codestates/WBABEProject-05/common/convertor"
	"github.com/codestates/WBABEProject-05/common/enum"
	error2 "github.com/codestates/WBABEProject-05/common/error"
	"github.com/codestates/WBABEProject-05/common/util"
	"github.com/codestates/WBABEProject-05/common/validator"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/protocol/request"
	"github.com/codestates/WBABEProject-05/protocol/response"
	"sync"
	"time"
)

func (o *orderRecordService) RegisterOrderRecord(order *request.RequestOrder) (*response.ResponsePostOrder, error) {
	if err := validator.CheckExistsMenus(order.StoreID, order.Menus); err != nil {
		return nil, err
	}

	rct, err := order.ToPostReceipt()
	if err != nil {
		return nil, err
	}

	// 이전 주문 정보 저장도 비즈니스상 중요하지 않아보여 따로 컨트롤하지 않는 고루틴 처리
	go o.updateUserPreOrderInfo(order)

	if err := o.setTotalPriceAndNumbering(rct, order); err != nil {
		return nil, err
	}

	savedOrder, err := o.receiptModel.InsertReceipt(rct)
	if err != nil {
		return nil, err
	}

	// OrderCount 의 증가는 비즈니스상 중요하지않아보여 따로 컨틀롤하지 않는 고루틴 활용
	go o.updateOrderCount(order)

	return response.FromReceipt(savedOrder), nil
}

func (o *orderRecordService) ModifyOrderRecordFromCustomer(order *request.RequestPutCustomerOrder) (*response.ResponsePostOrder, error) {
	if err := validator.CheckRoleIsCustomer(order.CustomerID); err != nil {
		return nil, err
	}

	foundOrder, err := o.receiptModel.SelectReceiptByID(order.ID)
	if err != nil {
		return nil, error2.DoesNotExistsOrderErr
	}

	if convertor.ConvertOBJIDToString(foundOrder.CustomerID) != order.CustomerID {
		return nil, error2.BadAccessOrderError
	}

	if err := o.checkOrderStatus(order, foundOrder); err != nil {
		return nil, err
	}

	if _, err := o.receiptModel.UpdateCancelReceipt(foundOrder); err != nil {
		return nil, err
	}

	resPostOrder, err := o.RegisterOrderRecord(order.ToPutRequestOrder())
	if err != nil {
		return nil, error2.DoesNotReOrderError
	}

	return resPostOrder, nil
}

func (o *orderRecordService) ModifyOrderRecordFromStore(order *request.RequestPutStoreOrder) (int, error) {
	if err := validator.CheckStoreUser(order.StoreID, order.UserID); err != nil {
		return 0, err
	}

	foundOrder, err := o.receiptModel.SelectReceiptByID(order.ID)
	if err != nil {
		return 0, err
	}

	if convertor.ConvertOBJIDToString(foundOrder.StoreID) != order.StoreID {
		return 0, error2.BadAccessOrderError
	}

	if foundOrder.Status == enum.Cancel || foundOrder.Status == enum.Completion {
		return 0, error2.DoesNotModifyOrderError
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
	menus, err := o.menuModel.SelectMenusByIDs(order.StoreID, order.Menus)
	if err != nil {
		menusErr <- err
		return
	}
	rct.Price = o.sumMenusPrice(menus)
	menusErr <- nil
}

// checkOrderStatus 메뉴 추가의 경우 배달중은 실패 , 메뉴 변경의 경우 조리중,배달중인경우 실패 -> 즉, 기본으로 배달중은 실패, 추가적으로 완료도 실패시키자
func (o *orderRecordService) checkOrderStatus(order *request.RequestPutCustomerOrder, foundOrder *entity.Receipt) error {
	reqMIDs := util.ConvertSliceToExistMap(order.MenuIDs)
	isChange := o.isChangeOrderMenus(foundOrder, reqMIDs)
	switch {
	case isChange && foundOrder.Status == enum.Cooking:
		// 변경이면서 이미 조리중
		return error2.DoesNotModifyOrderError
	case !isChange && len(foundOrder.MenuIDs) == len(order.MenuIDs):
		// 주문 메뉴가 현재 똑같음
		return error2.DuplicatedDataError
	case foundOrder.Status == enum.Completion || foundOrder.Status == enum.Delivering || foundOrder.Status == enum.Cancel:
		// 이미 완료거나 배달중 , 취소주문
		return error2.DoesNotModifyOrderError
	}
	return nil
}

func (o *orderRecordService) isChangeOrderMenus(foundOrder *entity.Receipt, reqMIDs map[string]int) bool {
	for _, ID := range foundOrder.MenuIDs {
		// exist 가 false 는 -> 변경으로 볼 수 있다.
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
