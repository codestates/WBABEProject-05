package protocol

import (
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RequestPostMenu struct {
	StoreId     string `json:"store_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Origin      string `json:"origin" validate:"required"`
	Possible    bool   `json:"possible" validate:"required"`
	LimitCount  string `json:"limit_count,omitempty"`
	Description string `json:"description,omitempty"`
}

// TODO 생각해보니 수정은 기존에것을 가지고와서 해야할듯!!!!! 업데이트 시간만 바꿔줘야하기도하고 , 필드들만 따로 명시해 수정할거 아니면 통으로 수정되기에!!!!!!
func (r *RequestPostMenu) ToStoreIdAndMenuNewId() (primitive.ObjectID, *dom.Menu, error) {
	id, err := primitive.ObjectIDFromHex(r.StoreId)
	if err != nil {
		return primitive.ObjectID{}, nil, err
	}
	menu := &dom.Menu{
		Id:          primitive.NewObjectID(),
		Name:        r.Name,
		Price:       r.Price,
		Origin:      r.Origin,
		Possible:    r.Possible,
		LimitCount:  r.LimitCount,
		Description: r.Description,
		BaseTime: &dom.BaseTime{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	return id, menu, nil
}

func (r *RequestPostMenu) ToStoreIdAndMenuMatchId(menuId string) (primitive.ObjectID, *dom.Menu, error) {
	sid, err := primitive.ObjectIDFromHex(r.StoreId)
	if err != nil {
		return primitive.ObjectID{}, nil, err
	}
	mid, err := primitive.ObjectIDFromHex(menuId)
	if err != nil {
		return primitive.ObjectID{}, nil, err
	}

	menu := &dom.Menu{
		Id:          mid,
		Name:        r.Name,
		Price:       r.Price,
		Origin:      r.Origin,
		Possible:    r.Possible,
		LimitCount:  r.LimitCount,
		Description: r.Description,
		BaseTime: &dom.BaseTime{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	return sid, menu, nil
}
