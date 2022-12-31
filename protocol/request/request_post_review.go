package request

import (
	"github.com/codestates/WBABEProject-05/common/util"
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RequestPostReview struct {
	StoreID string   `json:"store_id,omitempty"`
	UserID  string   `json:"user_id,omitempty"`
	Menu    []string `json:"menu_ids,omitempty"`
	Content string   `json:"content,omitempty"`
}

func (r *RequestPostReview) NewReview() (*entity.Review, error) {
	sID, err := primitive.ObjectIDFromHex(r.StoreID)
	if err != nil {
		return nil, err
	}

	uID, err := primitive.ObjectIDFromHex(r.UserID)
	if err != nil {
		return nil, err
	}

	mIDs, err := util.ConvertStringsToObjIDs(r.Menu)
	if err != nil {
		return nil, err
	}

	return &entity.Review{
		ID:      primitive.NewObjectID(),
		StoreID: sID,
		UserID:  uID,
		Menu:    mIDs,
		Content: r.Content,
		BaseTime: &dom.BaseTime{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}, nil
}
