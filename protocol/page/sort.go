package page

type Sort struct {
	Name      string `json:"sort_name" validate:"eq=recommend_count|eq=rating|eq=re_order_count" form:"sort-name,default=base_time.updated_at"`
	Direction int    `json:"direction" form:"direction,default=-1"`
}
