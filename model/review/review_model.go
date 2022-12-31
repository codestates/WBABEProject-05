package review

import (
	"github.com/codestates/WBABEProject-05/common"
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/protocol/page"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *reviewModel) SelectReviews() {

}
func (r *reviewModel) InsertReview(review *entity.Review) (string, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, review)
	if err != nil {
		return "", err
	}
	return review.ID.Hex(), nil
}

func (r *reviewModel) SelectSortLimitedReviewsByMenuID(menuID string, sort *page.Sort, skip int, limit int) ([]*entity.Review, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	ID, err := primitive.ObjectIDFromHex(menuID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"menu_id": ID}
	opt := options.Find().SetSort(bson.M{sort.Name: sort.Direction}).SetSkip(int64(skip)).SetLimit(int64(limit))
	//opt := options.Find().SetSort(bson.M{sort.Name: sort.Direction})
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
func (r *reviewModel) SelectSortLimitedReviewsByUserID(userID string, sort *page.Sort, skip int, limit int) ([]*entity.Review, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	ID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"user_id": ID}
	opt := options.Find().SetSort(bson.M{sort.Name: sort.Direction}).SetSkip(int64(skip)).SetLimit(int64(limit))
	//opt := options.Find().SetSort(bson.M{sort.Name: sort.Direction})
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
func (r *reviewModel) SelectTotalCountByMenuID(menuID string) (int, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	ID, err := primitive.ObjectIDFromHex(menuID)
	if err != nil {
		return 0, err
	}

	count, err := r.collection.CountDocuments(ctx, bson.M{"menu_id": ID})
	if err != nil {
		return 0, err
	}

	return int(count), nil
}
func (r *reviewModel) SelectTotalCountByUserID(userID string) (int, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	ID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return 0, err
	}

	count, err := r.collection.CountDocuments(ctx, bson.M{"user_id": ID})
	if err != nil {
		return 0, err
	}

	return int(count), nil
}
