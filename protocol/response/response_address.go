package response

import "github.com/codestates/WBABEProject-05/model/entity/dom"

type ResponseAddress struct {
	Street  string `json:"street" validate:"required"`
	Detail  string `json:"detail" validate:"required"`
	ZipCode string `json:"zip_code" validate:"required"`
}

func FromAddr(addr *dom.Address) *ResponseAddress {
	return &ResponseAddress{
		Street:  addr.Street,
		Detail:  addr.Detail,
		ZipCode: addr.ZipCode,
	}
}
