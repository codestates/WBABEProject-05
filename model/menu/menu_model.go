package menu

import (
	"github.com/codestates/WBABEProject-05/common"
	"github.com/codestates/WBABEProject-05/model/entity"
	mongo2 "github.com/codestates/WBABEProject-05/model/util"
	"github.com/codestates/WBABEProject-05/protocol/page"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO 싱글톤을 메서드보다 변수로 갖다쓰는게 낫지않나? init 과 함께 사용하면 될것같은데 나중에 해보자
var instance *menuModel

type menuModel struct {
	collection *mongo.Collection
}

func NewMenuModel(col *mongo.Collection) *menuModel {
	if instance != nil {
		return instance
	}
	instance = &menuModel{
		collection: col,
	}
	return instance
}

func (m *menuModel) InsertMenu(menu *entity.Menu) (string, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	if _, err := m.collection.InsertOne(ctx, menu); err != nil {
		return "", err
	}

	return menu.ID.Hex(), nil
}

func (m *menuModel) UpdateMenu(menu *entity.Menu) (int64, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	filter := bson.M{"_id": menu.ID}
	opt := menu.NewUpdateMenuBsonSetDWithPost()
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

func (m *menuModel) SelectSortLimitedMenus(storeID string, sort *page.Sort, skip, limit int) ([]*entity.Menu, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	ID, err := mongo2.ConvertStringToOBJID(storeID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"store_id": ID}
	opt := mongo2.NewSortFindOptions(sort, skip, limit)
	menusCursor, err := m.collection.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}

	var menus []*entity.Menu
	if err = menusCursor.All(ctx, &menus); err != nil {
		return nil, err
	}

	return menus, nil
}

func (m *menuModel) SelectSortLimitedMenusByName(name string, sort *page.Sort, skip, limit int) ([]*entity.Menu, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	filter := bson.M{"name": bson.M{"$regex": name}}
	opt := mongo2.NewSortFindOptions(sort, skip, limit)
	menusCursor, err := m.collection.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}

	var menus []*entity.Menu
	if err = menusCursor.All(ctx, &menus); err != nil {
		return nil, err
	}

	return menus, nil

}

func (m *menuModel) SelectTotalCount(storeID string) (int64, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	ID, err := mongo2.ConvertStringToOBJID(storeID)
	if err != nil {
		return 0, err
	}

	filter := bson.M{"store_id": ID}
	count, err := m.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (m *menuModel) SelectTotalCountByName(name string) (int64, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	filter := bson.M{"name": bson.M{"$regex": name}}
	count, err := m.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (m *menuModel) SelectMenusByIDs(storeID string, menuIDs []string) ([]*entity.Menu, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	sID, err := mongo2.ConvertStringToOBJID(storeID)
	if err != nil {
		return nil, err
	}

	inID, err := mongo2.ConvertStringsToOBJIDs(menuIDs)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"store_id": sID, "_id": bson.M{"$in": inID}}
	opt := options.Find().SetSort(
		bson.M{
			"base_time.updated_at": -1,
		})
	menuCursor, err := m.collection.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}

	var menus []*entity.Menu
	if err = menuCursor.All(ctx, &menus); err != nil {
		return nil, err
	}

	return menus, nil
}

func (m *menuModel) SelectMenuByID(menuID string) (*entity.Menu, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	mID, err := mongo2.ConvertStringToOBJID(menuID)
	if err != nil {
		return nil, err
	}

	var menu *entity.Menu
	filter := bson.M{"_id": mID}
	if err := m.collection.FindOne(ctx, filter).Decode(&menu); err != nil {
		return nil, err
	}

	return menu, nil
}

func (m *menuModel) SelectMenuByStoreIDAndName(storeID, name string) (*entity.Menu, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	sID, err := mongo2.ConvertStringToOBJID(storeID)
	if err != nil {
		return nil, err
	}

	var menu *entity.Menu
	filter := bson.D{{"store_id", sID}, {"name", name}}
	if err := m.collection.FindOne(ctx, filter).Decode(&menu); err != nil {
		return nil, err
	}

	return menu, nil
}

func (m *menuModel) SelectMenuByIDsAndDelete(menuID string) (*entity.Menu, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	mID, err := mongo2.ConvertStringToOBJID(menuID)
	if err != nil {
		return nil, err
	}

	var menu *entity.Menu
	filter := bson.M{"_id": mID}
	if err := m.collection.FindOneAndDelete(ctx, filter).Decode(&menu); err != nil {
		return nil, err
	}
	return menu, nil
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
