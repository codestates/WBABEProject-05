package store

import (
	"github.com/codestates/WBABEProject-05/common/enum"
	utilErr "github.com/codestates/WBABEProject-05/common/error"
	"github.com/codestates/WBABEProject-05/common/validator"
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/codestates/WBABEProject-05/protocol"
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
// @Tags 가게
// @Summary call Post store, return posted id by json.
// @Description 가게정보를 등록 할 수 있다.
// @name PostStore
// @Accept  json
// @Produce  json
// @Router /app/v1/stores/store [post]
// @Param RequestPostStore body request.RequestPostStore true "RequestPostStore JSON"
// @Success 201 {object} protocol.ApiResponse[any]
func (s *storeControl) PostStore(c *gin.Context) {
	reqS := &request.RequestPostStore{}
	if err := c.ShouldBindJSON(reqS); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	savedID, err := s.storeMenuService.RegisterStore(reqS)
	if err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessCodeAndData(
		http.StatusCreated,
		gin.H{"saved_id": savedID},
	).Response(c)
}

// PutStore godoc
// @Tags 가게
// @Summary call Put store, return modify count by json.
// @Description 가게를 수정할 수 있다.
// @name PutStore
// @Accept  json
// @Produce  json
// @Router /app/v1/stores/store [put]
// @Param store-id query string true "store-id"
// @Param RequestPutStore body request.RequestPutStore true "RequestPutStore JSON"
// @Success 200 {object} protocol.ApiResponse[any]
func (s *storeControl) PutStore(c *gin.Context) {
	var store *request.RequestPutStore
	if err := c.ShouldBind(&store); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	storeID := c.Query("store-id")
	if err := validator.CheckBlank(storeID); err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}

	cnt, err := s.storeMenuService.ModifyStore(storeID, store)
	if err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessData(gin.H{
		"updated_count": cnt,
	}).Response(c)
}

// PostMenu godoc
// @Tags 가게
// @Summary call Post menu, return saved id by json.
// @Description 메뉴를 등록할 수 있다.
// @name PostMenu
// @Accept  json
// @Produce  json
// @Router /app/v1/stores/store/menus/menu [post]
// @Param RequestMenu body request.RequestMenu true "RequestMenu JSON"
// @Success 201 {object} protocol.ApiResponse[any]
func (s *storeControl) PostMenu(c *gin.Context) {
	reqM := &request.RequestMenu{}
	if err := c.ShouldBindJSON(reqM); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	savedID, err := s.storeMenuService.RegisterMenu(reqM)
	if err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessCodeAndData(
		http.StatusCreated,
		gin.H{"saved_id": savedID},
	).Response(c)
}

// PutMenu godoc
// @Tags 가게
// @Summary call Put menu, return updated count by json.
// @Description 메뉴를 수정할 수 있다.
// @name PutMenu
// @Accept  json
// @Produce  json
// @Router /app/v1/stores/store/menus/menu [put]
// @Param menu-id query string true "menu-id"
// @Param RequestMenu body request.RequestMenu true "RequestMenu JSON"
// @Success 200 {object} protocol.ApiResponse[any]
func (s *storeControl) PutMenu(c *gin.Context) {
	reqM := &request.RequestMenu{}
	if err := c.ShouldBindJSON(reqM); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	menuID := c.Query("menu-id")
	if err := validator.CheckBlank(menuID); err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}

	cnt, err := s.storeMenuService.ModifyMenu(menuID, reqM)
	if err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}

	protocol.SuccessData(gin.H{
		"updated_count": cnt,
	}).Response(c)
}

// DeleteMenu godoc
// @Tags 가게
// @Summary call Delete menu, return deleted count by json.
// @Description 메뉴를 삭제할 수 있다.
// @name DeleteMenu
// @Accept  json
// @Produce  json
// @Router /app/v1/stores/store/menus/menu [delete]
// @Param RequestDeleteMenu query request.RequestDeleteMenu true "RequestDeleteMenu"
// @Success 200 {object} protocol.ApiResponse[any]
func (s *storeControl) DeleteMenu(c *gin.Context) {
	menu := &request.RequestDeleteMenu{}
	if err := c.ShouldBindQuery(menu); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	count, err := s.storeMenuService.DeleteMenuAndBackup(menu)
	if err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessData(gin.H{
		"deleted_count": count,
	}).Response(c)
}

// GetMenuSortedPagesByStoreID godoc
// @Tags 가게
// @Summary call Get sorted menu page, return sorted menu pages data by json.
// @Description 특정 가게 메뉴 리스트를 보여준다. 정렬 가능 - name: rating, order_count, base_time.updated_at | direction: 1, -1
// @name GetMenuSortedPagesByStoreID
// @Accept  json
// @Produce  json
// @Router /app/v1/stores/store/menus [get]
// @Param store-id query string true "store-id"
// @Param RequestPage query request.RequestPage true "RequestPage"
// @Param Sort query page.Sort true "Sort"
// @Success 200 {object} protocol.ApiResponse[any]
func (s *storeControl) GetMenuSortedPagesByStoreID(c *gin.Context) {
	page := &request.RequestPage{}
	if err := c.ShouldBindQuery(page); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	storeID := c.Query("store-id")
	if err := validator.CheckBlank(storeID); err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}

	menus, err := s.storeMenuService.FindMenusSortedPage(storeID, page)
	if err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessData(menus).Response(c)
}

// GetMenuSortedPagesByName godoc
// @Tags 가게
// @Summary call Get store menus page, return sorted menus page data by json.
// @Description 메뉴 이름으로 검색해 특정 가게 메뉴 리스트를 보여준다. 정렬 가능 - name: rating, order_count, base_time.updated_at | direction: 1, -1
// @name GetMenuSortedPagesByName
// @Accept  json
// @Produce  json
// @Router /app/v1/stores/store/menus/menu [get]
// @Param name query string true "name"
// @Param RequestPage query request.RequestPage true "RequestPage"
// @Param Sort query page.Sort true "Sort"
// @Success 200 {object} protocol.ApiResponse[any]
func (s *storeControl) GetMenuSortedPagesByName(c *gin.Context) {
	page := &request.RequestPage{}
	if err := c.ShouldBindQuery(page); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	name := c.Query("name")
	if err := validator.CheckBlank(name); err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}

	menus, err := s.storeMenuService.FindMenusSortedPageByName(name, page)
	if err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessData(menus).Response(c)
}

// GetRecommendMenus godoc
// @Tags 가게
// @Summary call Get store and recommend menus, return store and recommend menus data by json.
// @Description 특정 가게의 추천 메뉴 상세 정보 보여준다.
// @name GetRecommendMenus
// @Accept  json
// @Produce  json
// @Router /app/v1/stores/store/recommends [get]
// @Param store-id query string true "store-id"
// @Success 200 {object} protocol.ApiResponse[any]
func (s *storeControl) GetRecommendMenus(c *gin.Context) {
	storeID := c.Query("store-id")
	if err := validator.CheckBlank(storeID); err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}

	recommendMenus, err := s.storeMenuService.FindRecommendMenus(storeID)
	if err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}

	protocol.SuccessData(recommendMenus).Response(c)
}

// GetStoresSortedPage godoc
// @Tags 가게
// @Summary call Get store pages, return store pages data by json.
// @Description 가게들 정보를 보여준다. 정렬 가능 - name: base_time.updated_at | direction: 1, -1
// @name GetStoresSortedPage
// @Accept json
// @Produce json
// @Router /app/v1/stores [get]
// @Param RequestPage query request.RequestPage true "RequestPage"
// @Param Sort query page.Sort true "Sort"
// @Success 200 {object} protocol.ApiResponse[any]
func (s *storeControl) GetStoresSortedPage(c *gin.Context) {
	page := &request.RequestPage{}
	if err := c.ShouldBindQuery(page); err != nil {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	if page.Sort.Name != enum.SortBaseTimeUpdateAt {
		protocol.Fail(utilErr.BadRequestError).Response(c)
		return
	}

	stores, err := s.storeMenuService.FindStoresSortedPage(page)
	if err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessData(stores).Response(c)
}

// GetStore godoc
// @Tags 가게
// @Summary call Get store, return store by json.
// @Description 특정 가게의 정보를 보여준다.
// @name GetStore
// @Accept  json
// @Produce  json
// @Router /app/v1/stores/store [get]
// @Param store-id query string true "store_id"
// @Success 200 {object} protocol.ApiResponse[any]
func (s *storeControl) GetStore(c *gin.Context) {
	storeID := c.Query("store-id")
	if err := validator.CheckBlank(storeID); err != nil {
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}

	foundStore, err := s.storeMenuService.FindStore(storeID)
	if err != nil {
		logger.AppLog.Info(err)
		protocol.Fail(utilErr.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessData(foundStore).Response(c)
}
