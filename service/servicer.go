package service

type Servicer interface {
	OrderReceiptServicer() (OrderReceiptServicer, error)
	MenuReviewServicer() (MenuReviewServicer, error)
	StoreMenuServicer() (StoreMenuServicer, error)
}
