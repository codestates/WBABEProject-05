package review

import (
	"github.com/codestates/WBABEProject-05/common"
	"github.com/codestates/WBABEProject-05/model/entity"
	"go.mongodb.org/mongo-driver/mongo"
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
