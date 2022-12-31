package menu

import (
	"github.com/codestates/WBABEProject-05/common"
	"github.com/codestates/WBABEProject-05/common/util"
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/protocol/page"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func (m *menuModel) DeleteMenu() {

}
func (m *menuModel) UpdateMenu(menu *entity.Menu) (int, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	filter := bson.M{"_id": menu.ID}
	opt := menu.NewUpdateMenuBsonSetD()
	result, err := m.collection.UpdateOne(ctx, filter, opt)
	if err != nil {
		return 0, err
	}

	return int(result.ModifiedCount), nil
}

func (m *menuModel) SelectSortLimitedMenus(storeID string, sort *page.Sort, skip, limit int) ([]*entity.Menu, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	ID, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"store_id": ID}
	opt := options.Find().SetSort(bson.M{sort.Name: sort.Direction}).SetSkip(int64(skip)).SetLimit(int64(limit))
	//opt := options.Find().SetSort(bson.M{sort.Name: sort.Direction})
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

func (m *menuModel) SelectTotalCount() (int, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	count, err := m.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (m *menuModel) SelectMenusByIds(storeID string, menuIDs []string) ([]*entity.Menu, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	sID, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, err
	}

	inID, err := util.ConvertStringsToObjIDs(menuIDs)
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

func (m *menuModel) SelectMenuByIdsAndDelete(menuId string) (*entity.Menu, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	mID, err := primitive.ObjectIDFromHex(menuId)
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
