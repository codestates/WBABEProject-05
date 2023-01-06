package page

type Sort struct {
	Name      string `json:"sort_name" binding:"eq=rating|eq=order_count|eq=base_time.updated_at" form:"sort_name,default=base_time.updated_at"`
	Direction int    `json:"direction" binding:"eq=-1|eq=1" form:"direction,default=-1"`
}
