package menu

import (
	"context"
	"github.com/codestates/WBABEProject-05/common"
	"github.com/codestates/WBABEProject-05/common/convertor"
	"github.com/codestates/WBABEProject-05/model/common/query"
	"github.com/codestates/WBABEProject-05/model/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *menuModel) SelectSortLimitedMenus(storeID string, pageQuery *query.PageQuery) ([]*entity.Menu, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	ID, err := convertor.ConvertStringToOBJID(storeID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"store_id": ID}
	opt := pageQuery.NewSortFindOptions()
	menus, err := m.findMenus(ctx, filter, opt)
	if err != nil {
		return nil, err
	}

	return menus, nil
}

func (m *menuModel) SelectSortLimitedMenusByName(name string, pageQuery *query.PageQuery) ([]*entity.Menu, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	filter := bson.M{"name": bson.M{"$regex": name}}
	opt := pageQuery.NewSortFindOptions()
	menus, err := m.findMenus(ctx, filter, opt)
	if err != nil {
		return nil, err
	}

	return menus, nil

}

func (m *menuModel) SelectTotalCount(storeID string) (int64, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	ID, err := convertor.ConvertStringToOBJID(storeID)
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

	sID, err := convertor.ConvertStringToOBJID(storeID)
	if err != nil {
		return nil, err
	}

	inID, err := convertor.ConvertStringsToOBJIDs(menuIDs)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"store_id": sID, "_id": bson.M{"$in": inID}}
	opt := options.Find().SetSort(
		bson.M{
			"base_time.updated_at": -1,
		})

	menus, err := m.findMenus(ctx, filter, opt)
	if err != nil {
		return nil, err
	}

	return menus, nil
}

func (m *menuModel) SelectMenuByID(menuID string) (*entity.Menu, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	mID, err := convertor.ConvertStringToOBJID(menuID)
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

	sID, err := convertor.ConvertStringToOBJID(storeID)
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

	mID, err := convertor.ConvertStringToOBJID(menuID)
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

func (m *menuModel) findMenus(ctx context.Context, filter bson.M, opt *options.FindOptions) ([]*entity.Menu, error) {
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
