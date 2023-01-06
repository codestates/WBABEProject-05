package receipt

import (
	"github.com/codestates/WBABEProject-05/common"
	"github.com/codestates/WBABEProject-05/common/convertor"
	"github.com/codestates/WBABEProject-05/model/common/query"
	"github.com/codestates/WBABEProject-05/model/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (r *receiptModel) InsertReceipt(receipt *entity.Receipt) (*entity.Receipt, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	if _, err := r.collection.InsertOne(ctx, receipt); err != nil {
		return nil, err
	}

	return receipt, nil
}

func (r *receiptModel) UpdateReceiptStatus(receipt *entity.Receipt) (int64, error) {
	opt := receipt.NewUpdateStatusBsonSetD()
	return r.updateStatusByBsonSetD(receipt, opt)
}

func (r *receiptModel) UpdateCancelReceipt(receipt *entity.Receipt) (int64, error) {
	opt := receipt.NewUpdateStatusCancelBsonSetD()
	return r.updateStatusByBsonSetD(receipt, opt)
}

func (r *receiptModel) SelectReceiptByID(receiptID string) (*entity.Receipt, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	ID, err := convertor.ConvertStringToOBJID(receiptID)
	if err != nil {
		return nil, err
	}

	var receipt *entity.Receipt
	filter := query.GetDefaultIDFilter(ID)
	if err := r.collection.FindOne(ctx, filter).Decode(&receipt); err != nil {
		return nil, err
	}
	return receipt, nil
}
func (r *receiptModel) SelectSortLimitedReceipt(ID, status, userRole string, pageQuery *query.PageQuery) ([]*entity.Receipt, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	objID, err := convertor.ConvertStringToOBJID(ID)
	if err != nil {
		return nil, err
	}

	filter, err := query.GetCheckedUserRoleStatusFilter(objID, status, userRole)
	if err != nil {
		return nil, err
	}

	opt := pageQuery.NewSortFindOptions()
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

func (r *receiptModel) SelectToDayTotalCount() (int64, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	filter, err := query.GetToDayGteFilter()
	if err != nil {
		return 0, err
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, err
}

func (r *receiptModel) SelectTotalCount(ID, status, userRole string) (int64, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return 0, err
	}

	filter, err := query.GetCheckedUserRoleStatusFilter(objID, status, userRole)
	if err != nil {
		return 0, err
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *receiptModel) updateStatusByBsonSetD(receipt *entity.Receipt, bsonSet bson.D) (int64, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	filter := query.GetDefaultIDFilter(receipt.ID)
	result, err := r.collection.UpdateOne(ctx, filter, bsonSet)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}
