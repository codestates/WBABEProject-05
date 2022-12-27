package store

import (
	utilErr "github.com/codestates/WBABEProject-05/common/error"
	"github.com/codestates/WBABEProject-05/protocol"
	"github.com/codestates/WBABEProject-05/service/store"
	"github.com/gin-gonic/gin"
)

var instance *storeControl

type storeControl struct {
	storeMenuService store.StoreMenuServicer
}

func NewStoreControl(svc store.StoreMenuServicer) *storeControl {
	if instance != nil {
		return instance
	}
	instance = &storeControl{
		storeMenuService: svc,
	}
	return instance
}

// PostMenu godoc
// @Summary call Post menu in store, return saved id by json.
// @Description 메뉴를 등록할 수 있다.
// @name PostMenu
// @Accept  json
// @Produce  json
// @Router /app/v1/stores/menu [post]
// @Param menu body protocol.RequestPostMenu true "RequestPostMenu JSON"
// @Success 200 {object} map[string]string
func (s *storeControl) PostMenu(c *gin.Context) {
	reqM := &protocol.RequestPostMenu{}
	err := c.ShouldBindJSON(reqM)
	if err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	modiCount, err := s.storeMenuService.RegisterMenu(reqM)
	if err != nil {
		protocol.Fail(utilErr.NewError(err)).Response(c)
		return
	}
	protocol.SuccessData(gin.H{
		"posted_count": modiCount,
	}).Response(c)
}

// DeleteMenu godoc
// @Summary call Delete menu in store, return deleted count by json.
// @Description 메뉴를 삭제할 수 있다.
// @name DeleteMenu
// @Accept  json
// @Produce  json
// @Router /app/v1/stores/menu [delete]
// @Param store-id query string true "store-id"
// @Param menu-id query string true "menu-id"
// @Success 200 {object} map[string]int
func (s *storeControl) DeleteMenu(c *gin.Context) {
	storeId := c.Query("store-id")
	menuId := c.Query("menu-id")
	if menuId == "" || storeId == "" {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}
	count, err := s.storeMenuService.DeleteMenuAndBackup(storeId, menuId)
	if err != nil {
		protocol.Fail(utilErr.NewError(err)).Response(c)
		return
	}
	protocol.SuccessData(gin.H{
		"deleted_count": count,
	}).Response(c)
}

func (s *storeControl) PutSoreAndRecommendMenu(c *gin.Context) {
	s.storeMenuService.ModifyStoreAndRecommendMenus()
}

// PutMenu godoc
// @Summary call Put menu, return updated count by json.
// @Description 메뉴를 수정할 수 있다.
// @name PutMenu
// @Accept  json
// @Produce  json
// @Router /app/v1/stores/menu [put]
// @Param menu-id query string true "menu-id"
// @Param menu body protocol.RequestPostMenu true "RequestPostMenu JSON"
// @Success 200 {object} map[string]string
func (s *storeControl) PutMenu(c *gin.Context) {
	reqM := &protocol.RequestPostMenu{}
	mid := c.Query("menu-id")
	err := c.ShouldBindJSON(reqM)
	if err != nil || mid == "" {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}
	cnt, err := s.storeMenuService.ModifyMenu(mid, reqM)
	if err != nil {
		protocol.Fail(utilErr.NewError(err)).Response(c)
		return
	}

	protocol.SuccessData(gin.H{
		"updated_count": cnt,
	}).Response(c)
}

func (s *storeControl) GetRecommendMenusSortedTimeDesc(c *gin.Context) {
	s.storeMenuService.FindRecommendMenusSortedTimeDesc()
}

func (s *storeControl) GetMenuSortedPages(c *gin.Context) {
	s.storeMenuService.FindMenusSortedPage()
}

// PostStore godoc
// @Summary call Post store, return posted id by json.
// @Description 가게정보를 등록 할 수 있다.
// @name PostStore
// @Accept  json
// @Produce  json
// @Router /app/v1/stores [post]
// @Param store body protocol.RequestPostStore true "RequestPostStore JSON"
// @Success 200 {object} map[string]string
func (s *storeControl) PostStore(c *gin.Context) {
	reqS := &protocol.RequestPostStore{}
	err := c.ShouldBindJSON(reqS)
	if err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}
	savedId, err := s.storeMenuService.RegisterStore(reqS)
	if err != nil {
		protocol.Fail(utilErr.NewError(err)).Response(c)
		return
	}
	protocol.SuccessData(gin.H{
		"saved_id": savedId,
	}).Response(c)
}
