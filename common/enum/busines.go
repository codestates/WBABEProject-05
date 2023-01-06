package enum

const (
	CustomerRole = "customer"
	StoreRole    = "store"

	Waiting       = "주문대기"
	Cancel        = "주문취소"
	OrderReceived = "주문접수완료"
	Cooking       = "조리중"
	Delivering    = "배달중"
	Completion    = "배달완료"

	SortBaseTimeUpdateAt = "base_time.updated_at"
	SortRating           = "rating"
	SortOrderCount       = "order_count"
	SortASC              = 1
	SortDESC             = -1

	BlankSTR = ""
)

var (
	OrderStatusMap = map[string]string{
		Waiting:       "주문대기",
		Cancel:        "주문취소",
		OrderReceived: "주문접수완료",
		Cooking:       "조리중",
		Delivering:    "배달중",
		Completion:    "배달완료",
	}
)
