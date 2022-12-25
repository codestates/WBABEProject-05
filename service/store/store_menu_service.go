package store

import (
	"encoding/json"
	"github.com/codestates/WBABEProject-05/common/flag"
	"github.com/codestates/WBABEProject-05/config/db"
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/model/store"
	"github.com/codestates/WBABEProject-05/protocol"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"time"
)

type storeMenuService struct {
	storeModel store.StoreModeler
}

var instance *storeMenuService

func NewStoreMenuService(
	sd store.StoreModeler,
) *storeMenuService {
	if instance != nil {
		return instance
	}
	instance = &storeMenuService{
		storeModel: sd,
	}
	return instance
}

func (s *storeMenuService) RegisterMenu(menu *protocol.RequestPostMenu) (int, error) {
	sid, m, err := menu.ToStoreIdAndMenuNewId()
	if err != nil {
		return 0, err
	}
	savedCount, err := s.storeModel.InsertMenu(sid, m)
	if err != nil {
		return 0, err
	}
	return savedCount, nil
}
func (s *storeMenuService) DeleteMenuAndBackup(storeId, menuId string) (int, error) {
	sOId, err := primitive.ObjectIDFromHex(storeId)
	if err != nil {
		return 0, err
	}
	mOId, err := primitive.ObjectIDFromHex(menuId)
	if err != nil {
		return 0, err
	}

	store, err := s.storeModel.SelectMenuByIdAndDelete(sOId, mOId)
	if err != nil || store == nil {
		return 0, err
	}

	path := flag.Flags[flag.DatabaseFlag.Name]
	dbcfg := db.NewDbConfig(*path)
	err = db.WriteBackup(dbcfg.BackupPath, &store)
	if err != nil {
		// TODO Err
		zap.L().Error(err.Error())
		if m, err := json.Marshal(store); err == nil {
			zap.L().Error(string(m))
		}
		return 0, err
	}
	return 1, nil
}
func (s *storeMenuService) ModifyMenu(menuId string, menu *protocol.RequestPostMenu) (int, error) {
	sId, m, err := menu.ToStoreIdAndMenuMatchId(menuId)
	if err != nil {
		return 0, err
	}

	cnt, err := s.storeModel.UpdateMenu(sId, m)
	if err != nil {
		return 0, err
	}
	return cnt, nil
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

	//rlen := len(store.RecommendMenus)
	//var menus = make([]*entity.Menu, rlen)
	//if rlen > 0 {
	//	var pob []primitive.ObjectID
	//	for i, menuId := range store.RecommendMenus {
	//		obi, err := primitive.ObjectIDFromHex(menuId)
	//		if err != nil {
	//			return "", err
	//		}
	//		pob[i] = obi
	//	}
	//
	//	menus, err = s.storeModel.SelectMenusByIds(pob)
	//	if err != nil {
	//		return "", err
	//	}
	//}

	st := &entity.Store{
		Id:         primitive.NewObjectID(),
		UserId:     uid,
		Name:       store.Name,
		Address:    store.Address.ToAddress(),
		StorePhone: store.StorePhone,
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
