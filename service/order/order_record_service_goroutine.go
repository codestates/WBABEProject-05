package order

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/protocol/request"
	"sync"
)

func (o *orderRecordService) setTotalPriceAndNumbering(rct *entity.Receipt, order *request.RequestOrder) error {
	var wg sync.WaitGroup
	wg.Add(2)

	// setTotalPrice
	findMenusErrCH := make(chan error, 1)
	go func() {
		defer wg.Done()
		menus, err := o.menuModel.SelectMenusByIDs(order.StoreId, order.Menus)
		if err != nil {
			findMenusErrCH <- err
			return
		}
		rct.Price = o.sumMenusPrice(menus)
		findMenusErrCH <- nil
	}()

	// setNumbering
	findTotalCountErrCH := make(chan error, 1)
	go func() {
		defer wg.Done()
		toDayCnt, err := o.receiptModel.SelectToDayTotalCount()
		if err != nil {
			findTotalCountErrCH <- err
			return
		}
		rct.Numbering = o.newNumbering(toDayCnt)
		findTotalCountErrCH <- nil
	}()

	wg.Wait()
	if err := <-findMenusErrCH; err != nil {
		return err
	}
	if err := <-findTotalCountErrCH; err != nil {
		return err
	}
	return nil
}
