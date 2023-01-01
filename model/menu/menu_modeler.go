package menu

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/protocol/page"
)

var MenuModel MenuModeler

// TODO 디테일한 점이 많아 서비스 만들어야할듯하다
type MenuModeler interface {
	InsertMenu(menu *entity.Menu) (string, error)

	UpdateMenu(menu *entity.Menu) (int64, error)

	UpdateAboutRating(menu *entity.Menu) (int64, error)

	SelectSortLimitedMenus(storeID string, sort *page.Sort, skip, limit int) ([]*entity.Menu, error)

	SelectTotalCount(storeID string) (int64, error)

	SelectMenusByIDs(storeID string, menuIDs []string) ([]*entity.Menu, error)

	SelectMenuByID(menuID string) (*entity.Menu, error)

	SelectMenuByIdsAndDelete(menuId string) (*entity.Menu, error)

	UpdateMenusInCOrderCount(menus []string) (int64, error)
}
