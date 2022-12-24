package protocol

type RequestMenuId struct {
	menuId string `form:"menu_id" validate:"required"`
}
