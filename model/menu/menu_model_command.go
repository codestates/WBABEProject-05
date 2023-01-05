package menu

import (
	"github.com/codestates/WBABEProject-05/common"
	"github.com/codestates/WBABEProject-05/common/enum"
	mongo2 "github.com/codestates/WBABEProject-05/model/common"
	"github.com/codestates/WBABEProject-05/model/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *menuModel) InsertMenu(menu *entity.Menu) (string, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	if _, err := m.collection.InsertOne(ctx, menu); err != nil {
		return enum.BlankSTR, err
	}

	return menu.ID.Hex(), nil
}

func (m *menuModel) UpdateMenu(menu *entity.Menu) (int64, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	filter := bson.M{"_id": menu.ID}
	opt := menu.NewUpdateMenuBsonSetD()
	result, err := m.collection.UpdateOne(ctx, filter, opt)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (m *menuModel) UpdateMenuRating(menu *entity.Menu) (int64, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	filter := bson.M{"_id": menu.ID}
	opt := menu.NewUpdateMenuBsonSetDAboutReview()
	result, err := m.collection.UpdateOne(ctx, filter, opt)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (m *menuModel) UpdateMenusInCOrderCount(menus []string) (int64, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	IDs, err := mongo2.ConvertStringsToOBJIDs(menus)
	if err != nil {
		return 0, err
	}

	filter := bson.M{"_id": bson.M{"$in": IDs}}
	opt := bson.M{"$inc": bson.M{"order_count": 1}}
	result, err := m.collection.UpdateMany(ctx, filter, opt)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}
