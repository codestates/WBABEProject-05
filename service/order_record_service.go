package service

import (
	"github.com/codestates/WBABEProject-05/model/receipt"
	"github.com/codestates/WBABEProject-05/protocol"
)

var orsvc *orderReceiptService

type orderReceiptService struct {
	receiptModel receipt.ReceiptModeler
}

func GetOrderService(
	rd receipt.ReceiptModeler,
) *orderReceiptService {
	if orsvc != nil {
		return orsvc
	}
	orsvc = &orderReceiptService{
		receiptModel: rd,
	}
	return orsvc
}

func (o *orderReceiptService) RegisterOrderRecord(order *protocol.RequestOrder) (string, error) {
	receipt, err := order.ToReceipt()
	if err != nil {
		return "", err
	}

	insertedId, err := o.receiptModel.InsertReceipt(receipt)
	if err != nil {
		return "", err
	}
	return insertedId, nil
}
func (o *orderReceiptService) ModifyOrderRecord() {

}
func (o *orderReceiptService) FindOrderRecordsSortedPage() {

}
func (o *orderReceiptService) SelectReceipts() {

}
