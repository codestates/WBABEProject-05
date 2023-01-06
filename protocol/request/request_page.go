package request

import (
	"github.com/codestates/WBABEProject-05/protocol/page"
)

type RequestPage struct {
	CurrentPage  int        `json:"current-page" form:"current-page,default=0"`
	ContentCount int        `json:"content-count" form:"content-count,default=5"`
	Sort         *page.Sort `json:"sorting"`
}

func (r *RequestPage) ToPageInfo(totalCount int) *page.PageInfo {
	var totalPages int
	totalPages = totalCount / r.ContentCount
	b := (totalCount % r.ContentCount) != 0
	if b {
		totalPages++
	}

	return &page.PageInfo{
		CurrentPage:   r.CurrentPage,
		ContentCount:  r.ContentCount,
		TotalPages:    totalPages,
		TotalContents: totalCount,
		Sorting:       r.Sort,
	}
}
