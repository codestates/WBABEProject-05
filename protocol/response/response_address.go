package response

import "github.com/codestates/WBABEProject-05/model/entity/dom"

type ResponseAddress struct {
	Street  string `json:"street"`
	Detail  string `json:"detail"`
	ZipCode string `json:"zip_code"`
}

func FromAddr(addr *dom.Address) *ResponseAddress {
	return &ResponseAddress{
		Street:  addr.Street,
		Detail:  addr.Detail,
		ZipCode: addr.ZipCode,
	}
}
