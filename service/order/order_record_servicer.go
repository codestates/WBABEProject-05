package order

import (
	"github.com/codestates/WBABEProject-05/protocol/page"
	"github.com/codestates/WBABEProject-05/protocol/request"
	"github.com/codestates/WBABEProject-05/protocol/response"
)

var OrderRecordService OrderRecordServicer

type OrderRecordServicer interface {
	RegisterOrderRecord(order *request.RequestOrder) (string, error)

	ModifyOrderRecordFromCustomer(order *request.RequestPutCustomerOrder) (string, error)

	ModifyOrderRecordFromStore(order *request.RequestPutStoreOrder) (int, error)

	FindOrderRecordsSortedPage(ID, userRole string, page *request.RequestPage) (*page.PageData[any], error)

	FindOrderRecord(orderID string) (*response.ResponseOrder, error)

	FiendSelectedMenusTotalPrice(storeID string, menuIDs []string) (*response.ResponseCheckPrice, error)
}
