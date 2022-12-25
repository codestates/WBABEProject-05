package receipt

import (
	"github.com/codestates/WBABEProject-05/common"
	"github.com/codestates/WBABEProject-05/model/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

var instance *receiptModel

type receiptModel struct {
	collection *mongo.Collection
}

func GetReceiptModel(col *mongo.Collection) *receiptModel {
	if instance != nil {
		return instance
	}
	instance = &receiptModel{
		collection: col,
	}
	return instance
}

func (r *receiptModel) InsertReceipt(receipt *entity.Receipt) (string, error) {
	ctx, cancel := common.GetContext(common.ModelTimeOut)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, receipt)
	if err != nil {
		return "", err
	}
	return receipt.Id.Hex(), nil
}
func (r *receiptModel) UpdateReceipt() {

}
func (r *receiptModel) SelectReceipt() {

}
func (r *receiptModel) SelectReceipts() {

}
