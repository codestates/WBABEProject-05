package protocol

type RequestPutMenu struct {
	menuId string          `json:"menu_id"`
	menu   RequestPostMenu `json:"menu"`
}
