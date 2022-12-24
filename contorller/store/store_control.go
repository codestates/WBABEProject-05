package store

import (
	utilErr "github.com/codestates/WBABEProject-05/common/error"
	"github.com/codestates/WBABEProject-05/protocol"
	"github.com/codestates/WBABEProject-05/service"
	"github.com/gin-gonic/gin"
)

var instance *storeControl

type storeControl struct {
	storeMenuService service.StoreMenuServicer
}

func GetStoreControl(svc service.StoreMenuServicer) *storeControl {
	if instance != nil {
		return instance
	}
	instance = &storeControl{
		storeMenuService: svc,
	}
	return instance
}

func (s *storeControl) PostMenu(c *gin.Context) {
	reqM := &protocol.RequestPostMenu{}
	err := c.ShouldBindJSON(reqM)
	if err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	//store, terr := reqM.ToStore()
	//if terr != nil {
	//	protocol.Fail(*terr).Response(c)
	//	return
	//}
	modiCount, err := s.storeMenuService.RegisterMenu(reqM)
	if err != nil {
		//TODO ERR
		return
	}
	protocol.SuccessData(gin.H{
		"posted_count": modiCount,
	}).Response(c)
}

func (s *storeControl) DeleteMenu(c *gin.Context) {
	storeId := c.Query("store-id")
	menuId := c.Query("menu-id")
	if menuId == "" || storeId == "" {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}
	count, err := s.storeMenuService.DeleteMenuAndBackup(storeId, menuId)
	if err != nil {
		//TODO ERR
		return
	}
	protocol.SuccessData(gin.H{
		"deleted_count": count,
	}).Response(c)
}

func (s *storeControl) PutSoreAndRecommendMenu(c *gin.Context) {
	s.storeMenuService.ModifyStoreAndRecommendMenus()
}

func (s *storeControl) PutMenu(c *gin.Context) {
	s.storeMenuService.ModifyMenu()
}

func (s *storeControl) GetRecommendMenusSortedTimeDesc(c *gin.Context) {
	s.storeMenuService.FindRecommendMenusSortedTimeDesc()
}

func (s *storeControl) GetMenuSortedPages(c *gin.Context) {
	s.storeMenuService.FindMenusSortedPage()
}

func (s *storeControl) PostStore(c *gin.Context) {
	reqS := &protocol.RequestPostStore{}
	err := c.ShouldBindJSON(reqS)
	if err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}
	savedId, err := s.storeMenuService.RegisterStore(reqS)
	if err != nil {
		//TODO ERR
		return
	}
	protocol.SuccessData(gin.H{
		"saved_id": savedId,
	}).Response(c)
}
