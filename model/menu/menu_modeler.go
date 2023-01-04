package menu

import (
	"github.com/codestates/WBABEProject-05/model/entity"
	"github.com/codestates/WBABEProject-05/protocol/page"
)

var MenuModel MenuModeler

type MenuModeler interface {
	InsertMenu(menu *entity.Menu) (string, error)

	UpdateMenu(menu *entity.Menu) (int64, error)

	UpdateMenuRating(menu *entity.Menu) (int64, error)

	SelectSortLimitedMenus(storeID string, sort *page.Sort, skip, limit int) ([]*entity.Menu, error)

	SelectSortLimitedMenusByName(name string, sort *page.Sort, skip, limit int) ([]*entity.Menu, error)

	SelectTotalCount(storeID string) (int64, error)

	SelectTotalCountByName(name string) (int64, error)

	SelectMenusByIDs(storeID string, menuIDs []string) ([]*entity.Menu, error)

	SelectMenuByID(menuID string) (*entity.Menu, error)

	SelectMenuByStoreIDAndName(storeID, name string) (*entity.Menu, error)

	SelectMenuByIDsAndDelete(menuID string) (*entity.Menu, error)

	UpdateMenusInCOrderCount(menus []string) (int64, error)
}
