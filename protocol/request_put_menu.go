package protocol

type RequestPutMenu struct {
	menuId RequestMenuId   `json:"menu_id"`
	menu   RequestPostMenu `json:"menu"`
}
