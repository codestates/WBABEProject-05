package protocol

import "github.com/codestates/WBABEProject-05/model/entity"

type RequestAddress struct {
	Street  string `json:"street"`
	Detail  string `json:"detail"`
	ZipCode string `json:"zip_code"`
}

func (r *RequestAddress) ToAddress() *entity.Address {
	addr := &entity.Address{
		Street:  r.Street,
		Detail:  r.Detail,
		ZipCode: r.ZipCode,
	}
	return addr
}
