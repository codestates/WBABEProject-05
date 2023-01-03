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
