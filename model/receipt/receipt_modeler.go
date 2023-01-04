package receipt

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/protocol/page"
)

var ReceiptModel ReceiptModeler

type ReceiptModeler interface {
	InsertReceipt(receipt *entity.Receipt) (string, error)

	UpdateReceiptStatus(receipt *entity.Receipt) (int64, error)

	UpdateCancelReceipt(receipt *entity.Receipt) (int64, error)

	SelectReceiptByID(receiptID string) (*entity.Receipt, error)

	SelectToDayTotalCount() (int64, error)

	SelectSortLimitedReceipt(ID, status, userRole string, sort *page.Sort, skip, limit int) ([]*entity.Receipt, error)

	SelectTotalCount(ID, status, userRole string) (int64, error)
}
