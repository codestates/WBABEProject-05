package order

import (
	"github.com/codestates/WBABEProject-05/common/convertor"
	"github.com/codestates/WBABEProject-05/model/common/query"
	"github.com/codestates/WBABEProject-05/protocol/page"
	"github.com/codestates/WBABEProject-05/protocol/request"
	"github.com/codestates/WBABEProject-05/protocol/response"
	util2 "github.com/codestates/WBABEProject-05/service/common"
)

func (o *orderRecordService) FindOrderRecordsSortedPage(ID, status, userRole string, pg *request.RequestPage) (*page.PageData[any], error) {
	skip := util2.NewSkipNumber(pg.CurrentPage, pg.ContentCount)

	pageQuery := query.NewPageQuery(pg.Sort.Name, pg.Sort.Direction, skip, pg.ContentCount)

	receipts, err := o.receiptModel.SelectSortLimitedReceipt(ID, status, userRole, pageQuery)
	if err != nil {
		return nil, err
	}

	totalCount, err := o.receiptModel.SelectTotalCount(ID, status, userRole)
	if err != nil {
		return nil, err
	}

	pgInfo := pg.ToPageInfo(int(totalCount))

	return page.NewPageData(receipts, pgInfo), nil
}

func (o *orderRecordService) FindOrderRecord(orderID string) (*response.ResponseOrder, error) {
	foundReceipt, err := o.receiptModel.SelectReceiptByID(orderID)
	if err != nil {
		return nil, err
	}

	menuIDs := convertor.ConvertOBJIDsToStrings(foundReceipt.MenuIDs)

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
