package model

import (
	"github.com/codestates/WBABEProject-05/model/receipt"
	"github.com/codestates/WBABEProject-05/model/review"
	"github.com/codestates/WBABEProject-05/model/store"
)

type Modeler interface {
	GetStoreModel() store.StoreModeler
	GetReviewModel() review.ReviewModeler
	GetReceiptModel() receipt.ReceiptModeler
}
