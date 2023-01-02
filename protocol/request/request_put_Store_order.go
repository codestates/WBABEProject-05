package request

type RequestPutStoreOrder struct {
	ID     string `json:"order_id" binding:"required"`
	Status string `json:"status" binding:"required,eq=주문대기|eq=주문취소|eq=주문접수완료|eq=조리중|eq=배달중|eq=배달완료"`
}
