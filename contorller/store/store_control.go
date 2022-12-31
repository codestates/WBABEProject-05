package store

import (
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/codestates/WBABEProject-05/protocol"
	utilErr "github.com/codestates/WBABEProject-05/protocol/error"
	"github.com/codestates/WBABEProject-05/protocol/request"
	"github.com/codestates/WBABEProject-05/service/store"
	"github.com/gin-gonic/gin"
	"net/http"
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

// PostStore godoc
// @Summary call Post store, return posted id by json.
// @Description 가게정보를 등록 할 수 있다.
// @name PostStore
// @Accept  json
// @Produce  json
// @Router /app/v1/stores [post]
// @Param store body protocol.RequestPostStore true "RequestPostStore JSON"
// @Success 201 {object} protocol.ApiResponse[any]
func (s *storeControl) PostStore(c *gin.Context) {
	reqS := &request.RequestPostStore{}
	if err := c.ShouldBindJSON(reqS); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	savedId, err := s.storeMenuService.RegisterStore(reqS)
	if err != nil {
		protocol.Fail(utilErr.NewApiError(err)).Response(c)
		return
	}
	protocol.SuccessCodeAndData(
		http.StatusCreated,
		gin.H{"saved_id": savedId},
	).Response(c)
}

func (s *storeControl) PutSore(c *gin.Context) {
	var store *request.RequestPutStore
	storeID := c.Query("store-id")
	if err := c.ShouldBind(&store); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}
	cnt, err := s.storeMenuService.ModifyStore(storeID, store)
	if err != nil {
		protocol.Fail(utilErr.NewApiError(err)).Response(c)
		return
	}
	protocol.SuccessData(gin.H{
		"updated_count": cnt,
	}).Response(c)
}

// PostMenu godoc
// @Summary call Post menu in store, return saved id by json.
// @Description 메뉴를 등록할 수 있다.
// @name PostMenu
// @Accept  json
// @Produce  json
// @Router /app/v1/stores/menu [post]
// @Param menu body request.RequestMenu true "RequestMenu JSON"
// @Success 200 {object} protocol.ApiResponse[any]
func (s *storeControl) PostMenu(c *gin.Context) {
	reqM := &request.RequestMenu{}
	if err := c.ShouldBindJSON(reqM); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	modiCount, err := s.storeMenuService.RegisterMenu(reqM)
	if err != nil {
		protocol.Fail(utilErr.NewApiError(err)).Response(c)
		return
	}
	protocol.SuccessCodeAndData(
		http.StatusCreated,
		gin.H{"posted_id": modiCount},
	).Response(c)
}

// PutMenu godoc
// @Summary call Put menu, return updated count by json.
// @Description 메뉴를 수정할 수 있다.
// @name PutMenu
// @Accept  json
// @Produce  json
// @Router /app/v1/stores/menu [put]
// @Param menu-id query string true "menu-id"
// @Param menu body request.RequestMenu true "RequestMenu JSON"
// @Success 200 {object} protocol.ApiResponse[any]
func (s *storeControl) PutMenu(c *gin.Context) {
	reqM := &request.RequestMenu{}
	mid := c.Query("menu-id")
	if err := c.ShouldBindJSON(reqM); err != nil || mid == "" {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}
	cnt, err := s.storeMenuService.ModifyMenu(mid, reqM)
	if err != nil {
		protocol.Fail(utilErr.NewApiError(err)).Response(c)
		return
	}

	protocol.SuccessData(gin.H{
		"updated_count": cnt,
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
// @Success 200 {object} protocol.ApiResponse[any]
func (s *storeControl) DeleteMenu(c *gin.Context) {
	menuId := c.Query("menu-id")
	if menuId == "" {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}
	count, err := s.storeMenuService.DeleteMenuAndBackup(menuId)
	if err != nil {
		protocol.Fail(utilErr.NewApiError(err)).Response(c)
		return
	}
	protocol.SuccessData(gin.H{
		"deleted_count": count,
	}).Response(c)
}

func (s *storeControl) GetMenuSortedPages(c *gin.Context) {
	page := &request.RequestPage{}
	srtID := c.Query("store-id")
	if err := c.ShouldBindQuery(page); err != nil || srtID == "" {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}
	menus, err := s.storeMenuService.FindMenusSortedPage(srtID, page)
	if err != nil {
		protocol.Fail(utilErr.NewApiError(err)).Response(c)
		return
	}
	protocol.SuccessData(menus).Response(c)
}

func (s *storeControl) GetRecommendMenus(c *gin.Context) {
	storeID := c.Query("store-id")
	resStore, err := s.storeMenuService.FindRecommendMenus(storeID)
	if err != nil {
		protocol.Fail(utilErr.NewApiError(err)).Response(c)
		return
	}

	protocol.SuccessData(resStore).Response(c)
}

func (s *storeControl) GetStoresSortedPage(c *gin.Context) {
	page := &request.RequestPage{}
	if err := c.ShouldBindQuery(page); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	stores, err := s.storeMenuService.FindStoresSortedPage(page)
	if err != nil {
		protocol.Fail(utilErr.NewApiError(err)).Response(c)
		return
	}
	protocol.SuccessData(stores).Response(c)
}

// GetStoreInSwagForTest godoc
// @Summary call Get store, return store by json.
// @Description 특정 store 의 모든 정보를 스웨거 테스트를 위해 보여준다.
// @name GetStoreInSwagForTest
// @Accept  json
// @Produce  json
// @Router /app/v1/stores/swag/store [get]
// @Param store-id query string true "store_id"
// @Success 200 {object} protocol.ApiResponse[entity.Store]
func (s *storeControl) GetStoreInSwagForTest(c *gin.Context) {
	strId := c.Query("store-id")
	str, err := s.storeMenuService.FindStore(strId)
	if err != nil {
		logger.AppLog.Info(err)
		protocol.Fail(utilErr.NewApiError(err)).Response(c)
		return
	}
	protocol.SuccessData(str).Response(c)
}
