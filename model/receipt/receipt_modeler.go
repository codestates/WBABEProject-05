package receipt

import (
	"github.com/codestates/WBABEProject-05/model/common/query"
	"github.com/codestates/WBABEProject-05/model/entity"
)

var ReceiptModel ReceiptModeler

type ReceiptModeler interface {
	InsertReceipt(receipt *entity.Receipt) (*entity.Receipt, error)

	UpdateReceiptStatus(receipt *entity.Receipt) (int64, error)

	UpdateCancelReceipt(receipt *entity.Receipt) (int64, error)

	SelectReceiptByID(receiptID string) (*entity.Receipt, error)

	SelectToDayTotalCount() (int64, error)

	SelectSortLimitedReceipt(ID, status, userRole string, pageQuery *query.PageQuery) ([]*entity.Receipt, error)

	SelectTotalCount(ID, status, userRole string) (int64, error)
}
