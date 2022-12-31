package order

import (
	"github.com/codestates/WBABEProject-05/model/receipt"
	"github.com/codestates/WBABEProject-05/protocol/request"
)

type orderRecordService struct {
	receiptModel receipt.ReceiptModeler
}

var instance *orderRecordService

func NewOrderRecordService(rd receipt.ReceiptModeler) *orderRecordService {
	if instance != nil {
		return instance
	}

	instance = &orderRecordService{receiptModel: rd}
	return instance
}

func (o *orderRecordService) RegisterOrderRecord(order *request.RequestOrder) (string, error) {
	//rct, err := order.ToReceipt()
	//if err != nil {
	//	return "", err
	//}
	//
	//insertedId, err := o.receiptModel.InsertReceipt(rct)
	//if err != nil {
	//	return "", err
	//}
	//
	//return insertedId, nil
	return "", nil
}
func (o *orderRecordService) ModifyOrderRecord() {

}
func (o *orderRecordService) FindOrderRecordsSortedPage() {

}
func (o *orderRecordService) SelectReceipts() {

}
