package request

import "github.com/codestates/WBABEProject-05/model/entity/dom"

type RequestAddress struct {
	Street  string `json:"street" validate:"required"`
	Detail  string `json:"detail" validate:"required"`
	ZipCode string `json:"zip_code" validate:"required"`
}

func (r *RequestAddress) ToAddress() *dom.Address {
	addr := &dom.Address{
		Street:  r.Street,
		Detail:  r.Detail,
		ZipCode: r.ZipCode,
	}
	return addr
}

func FromAddress(addr *dom.Address) *RequestAddress {
	return &RequestAddress{
		Street:  addr.Street,
		Detail:  addr.Detail,
		ZipCode: addr.ZipCode,
	}
}
