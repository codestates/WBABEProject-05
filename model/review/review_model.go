package review

import (
	"context"
	"github.com/codestates/WBABEProject-05/common"
	"github.com/codestates/WBABEProject-05/common/convertor"
	"github.com/codestates/WBABEProject-05/common/enum"
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/query"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var instance *reviewModel

type reviewModel struct {
	collection *mongo.Collection
}

func NewReviewModel(col *mongo.Collection) *reviewModel {
	if instance != nil {
		return instance
	}
	instance = &reviewModel{
		collection: col,
	}
	return instance
}

func (r *reviewModel) InsertReview(review *entity.Review) (string, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	if _, err := r.collection.InsertOne(ctx, review); err != nil {
		return enum.BlankSTR, err
	}
	return review.ID.Hex(), nil
}

func (r *reviewModel) SelectSortLimitedReviewsByMenuID(menuID string, pageQuery *query.PageQuery) ([]*entity.Review, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	ID, err := convertor.ConvertStringToOBJID(menuID)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"menu_id", ID}}
	opt := pageQuery.NewSortFindOptions()
	reviews, err := r.findReviews(ctx, filter, opt)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *reviewModel) SelectSortLimitedReviewsByUserID(ID, userRole string, pageQuery *query.PageQuery) ([]*entity.Review, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	OBJID, err := convertor.ConvertStringToOBJID(ID)
	if err != nil {
		return nil, err
	}

	filter, err := query.GetCheckedUserRoleStatusFilter(OBJID, enum.BlankSTR, userRole)
	if err != nil {
		return nil, err
	}

	opt := pageQuery.NewSortFindOptions()
	reviews, err := r.findReviews(ctx, filter, opt)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *reviewModel) SelectTotalCountByMenuID(menuID string) (int64, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	ID, err := convertor.ConvertStringToOBJID(menuID)
	if err != nil {
		return 0, err
	}

	count, err := r.collection.CountDocuments(ctx, bson.M{"menu_id": ID})
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *reviewModel) SelectTotalCountByUserID(ID, userRole string) (int64, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	OBJID, err := convertor.ConvertStringToOBJID(ID)
	if err != nil {
		return 0, err
	}

	filter, err := query.GetCheckedUserRoleStatusFilter(OBJID, enum.BlankSTR, userRole)
	if err != nil {
		return 0, err
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *reviewModel) findReviews(ctx context.Context, filter bson.D, opt *options.FindOptions) ([]*entity.Review, error) {
	reviewCursor, err := r.collection.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}

	var reviews []*entity.Review
	if err = reviewCursor.All(ctx, &reviews); err != nil {
		return nil, err
	}
	return reviews, nil
}
