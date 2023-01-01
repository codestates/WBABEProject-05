package page

type PageData[T any] struct {
	Data     T
	PageInfo *PageInfo
}

type PageInfo struct {
	CurrentPage   int   `json:"current_page"`
	ContentCount  int   `json:"page_count"`
	TotalPages    int   `json:"total_pages"`
	TotalContents int   `json:"total_contents"`
	Sorting       *Sort `json:"sorting"`
}

func NewPageData[T any](d T, p *PageInfo) *PageData[any] {
	return &PageData[any]{
		Data:     d,
		PageInfo: p,
	}
}
