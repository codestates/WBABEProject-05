package service

var stmsvc *storeMenuService

type storeMenuService struct {
}

func GetStoreMenuService() *storeMenuService {
	if stmsvc != nil {
		return stmsvc
	}
	stmsvc = &storeMenuService{}
	return stmsvc
}

func (s *storeMenuService) RegisterMenu() {

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
