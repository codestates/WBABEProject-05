package service

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/store"
	"github.com/codestates/WBABEProject-05/protocol"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var stmsvc *storeMenuService

type storeMenuService struct {
	storeModel store.StoreModeler
}

func GetStoreMenuService(
	sd store.StoreModeler,
) *storeMenuService {
	if stmsvc != nil {
		return stmsvc
	}
	stmsvc = &storeMenuService{
		storeModel: sd,
	}
	return stmsvc
}

func (s *storeMenuService) RegisterMenu(menu *protocol.RequestPostMenu) (int, error) {
	id, err := primitive.ObjectIDFromHex(menu.StoreId)
	if err != nil {
		return 0, err
	}
	str := &entity.Store{
		Id: id,
		Menu: []*entity.Menu{
			{
				Name:        menu.Name,
				Price:       menu.Price,
				Origin:      menu.Origin,
				Possible:    menu.Possible,
				LimitCount:  menu.LimitCount,
				Description: menu.Description,
			},
		},
	}
	savedCount, aErr := s.storeModel.InsertMenu(str)
	if aErr != nil {
		return 0, err
	}
	return savedCount, nil
}
func (s *storeMenuService) DeleteMenuAndBackup() {

}
func (s *storeMenuService) ModifyMenu() {

}
func (s *storeMenuService) ModifyStoreAndRecommendMenus() {

}
func (s *storeMenuService) FindRecommendMenusSortedTimeDesc() {

}
func (s *storeMenuService) FindMenusSortedPage() {

}

func (s *storeMenuService) RegisterStore(store *protocol.RequestPostStore) (string, error) {
	uid, err := primitive.ObjectIDFromHex(store.UserId)
	if err != nil {
		return "", err
	}

	rlen := len(store.RecommendMenus)
	var menus = make([]*entity.Menu, rlen)
	if rlen > 0 {
		var pob []primitive.ObjectID
		for i, menuId := range store.RecommendMenus {
			obi, err := primitive.ObjectIDFromHex(menuId)
			if err != nil {
				return "", err
			}
			pob[i] = obi
		}

		menus, err = s.storeModel.SelectMenusByIds(pob)
		if err != nil {
			return "", err
		}
	}

	addr := &entity.Address{
		Street:  store.Address.Street,
		Detail:  store.Address.Detail,
		ZipCode: store.Address.ZipCode,
	}

	st := &entity.Store{
		Id:             primitive.NewObjectID(),
		UserId:         uid,
		Address:        addr,
		RecommendMenus: menus,
		StorePhone:     store.StorePhone,
		BaseTime: &entity.BaseTime{
			Created_at: time.Now(),
			Updated_at: time.Now(),
		},
	}

	savedId, err := s.storeModel.InsertStore(st)
	if err != nil {
		return "", err
	}
	return savedId, nil
}
