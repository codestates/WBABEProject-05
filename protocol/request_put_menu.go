package protocol

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RequestPutMenu struct {
	menuId string           `json:"menu_id" validate:"required"`
	menu   *RequestPostMenu `json:"menu" validate:"required"`
}

func (r *RequestPutMenu) ToStoreIdAndMenu() ([12]byte, *entity.Menu, error) {
	storeId, menuWithNewId, err := r.menu.ToStoreIdAndMenuNewId()
	if err != nil {
		return [12]byte{}, nil, err
	}

	mId, err := primitive.ObjectIDFromHex(r.menuId)
	if err != nil {
		return [12]byte{}, nil, err
	}

	menuWithNewId.Id = mId
	//mid, err := primitive.ObjectIDFromHex(r.menuId)
	//if err != nil {
	//	return primitive.ObjectID{}, nil, err
	//}
	//sid, err := primitive.ObjectIDFromHex(r.StoreId)
	//if err != nil {
	//	return primitive.ObjectID{}, nil, err
	//}
	//
	//menu := &entity.Menu{
	//	Id:          mid,
	//	Name:        r.Name,
	//	Price:       r.Price,
	//	Origin:      r.Origin,
	//	Possible:    r.Possible,
	//	LimitCount:  r.LimitCount,
	//	Description: r.Description,
	//	BaseTime: entity.BaseTime{
	//		Created_at: time.Now(),
	//		Updated_at: time.Now(),
	//	},
	//}

	//return sid, menu, nil
	return storeId, menuWithNewId, nil
}
