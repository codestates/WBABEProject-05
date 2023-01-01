package receipt

import (
	"github.com/codestates/WBABEProject-05/common"
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/protocol/page"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var instance ReceiptModeler

type receiptModel struct {
	collection *mongo.Collection
}

func NewReceiptModel(col *mongo.Collection) ReceiptModeler {
	if instance != nil {
		return instance
	}
	instance = &receiptModel{
		collection: col,
	}
	return instance
}

func (r *receiptModel) InsertReceipt(receipt *entity.Receipt) (string, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	if _, err := r.collection.InsertOne(ctx, receipt); err != nil {
		return "", err
	}

	return receipt.ID.Hex(), nil
}
func (r *receiptModel) UpdateReceiptStatus(receipt *entity.Receipt) (int, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	filter := bson.M{"_id": receipt.ID}
	opt := receipt.NewUpdateStatusOrderBsonSetD()
	result, err := r.collection.UpdateOne(ctx, filter, opt)
	if err != nil {
		return 0, err
	}
	return int(result.ModifiedCount), nil
}

func (r *receiptModel) UpdateCancelReceipt(receipt *entity.Receipt) (int, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	filter := bson.M{"_id": receipt.ID}
	opt := receipt.NewUpdateOrderCancelBsonSetD()
	result, err := r.collection.UpdateOne(ctx, filter, opt)
	if err != nil {
		return 0, err
	}
	return int(result.ModifiedCount), nil
}

func (r *receiptModel) SelectReceiptByID(receiptID string) (*entity.Receipt, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(receiptID)
	if err != nil {
		return nil, err
	}

	var receipt *entity.Receipt
	filter := bson.M{"_id": id}
	if err := r.collection.FindOne(ctx, filter).Decode(&receipt); err != nil {
		return nil, err
	}
	return receipt, nil
}
func (r *receiptModel) SelectSortLimitedReceipt(ID, userRole string, sort *page.Sort, skip, limit int) ([]*entity.Receipt, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{}
	switch userRole {
	case entity.CustomerRole:
		filter = bson.M{"customer_id": objID}
	case entity.StoreRole:
		filter = bson.M{"store_id": objID}
	}
	opt := options.Find().SetSort(bson.M{sort.Name: sort.Direction}).SetSkip(int64(skip)).SetLimit(int64(limit))
	receiptCursor, err := r.collection.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}

	var receipts []*entity.Receipt
	if err = receiptCursor.All(ctx, &receipts); err != nil {
		return nil, err
	}

	return receipts, nil
}

func (r *receiptModel) SelectToDayTotalCount() (int, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	KST, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		return 0, err
	}
	now := time.Now()
	startTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, KST).UTC()
	filter := bson.M{"base_time.created_at": bson.M{"$gte": startTime}}
	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return int(count), err
}

func (r *receiptModel) SelectTotalCount(ID, userRole string) (int, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return 0, err
	}

	filter := bson.M{}
	switch userRole {
	case entity.CustomerRole:
		filter = bson.M{"customer_id": objID}
	case entity.StoreRole:
		filter = bson.M{"store_id": objID}
	}
	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}
