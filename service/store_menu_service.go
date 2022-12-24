package service

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/store"
	"github.com/codestates/WBABEProject-05/protocol"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var stmsvc *storeMenuService

type storeMenuService struct {
	storeModel store.StoreModeler
}

func GetStoreMenuService(modeler store.StoreModeler) *storeMenuService {
	if stmsvc != nil {
		return stmsvc
	}
	stmsvc = &storeMenuService{
		storeModel: modeler,
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
