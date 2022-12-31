package response

import (
	"github.com/codestates/WBABEProject-05/model/entity/dom"
	"time"
)

type ResponseBaseTime struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromBaseTime(t *dom.BaseTime) *ResponseBaseTime {
	return &ResponseBaseTime{
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
