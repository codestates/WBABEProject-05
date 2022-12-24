package model

import (
	"github.com/codestates/WBABEProject-05/model/receipt"
	"github.com/codestates/WBABEProject-05/model/review"
	"github.com/codestates/WBABEProject-05/model/store"
	"github.com/codestates/WBABEProject-05/model/user"
)

type Modeler interface {
	StoreModel() store.StoreModeler
	ReviewModel() review.ReviewModeler
	ReceiptModel() receipt.ReceiptModeler
	UserModel() user.UserModeler
}
