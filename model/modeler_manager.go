package model

import (
	"github.com/codestates/WBABEProject-05/model/menu"
	"github.com/codestates/WBABEProject-05/model/receipt"
	"github.com/codestates/WBABEProject-05/model/review"
	"github.com/codestates/WBABEProject-05/model/store"
	"github.com/codestates/WBABEProject-05/model/user"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ReceiptCollectionName = "receipt"
	ReviewCollectionName  = "review"
	StoreCollectionName   = "store"
	UserCollectionName    = "user"
	MenuCollectionName    = "menu"
)

// LoadModelImpls generate singleton modelers

func InjectModelsMongoDependency(m map[string]*mongo.Collection) {
	receipt.ReceiptModel = receipt.NewReceiptModel(m[ReceiptCollectionName])
	review.ReviewModel = review.NewReviewModel(m[ReviewCollectionName])
	store.StoreModel = store.NewStoreModel(m[StoreCollectionName])
	user.UserModel = user.NewUserModel(m[UserCollectionName])
	menu.MenuModel = menu.NewMenuModel(m[MenuCollectionName])
}

func CreateIndexesInModels() {
	// User Collection
	AppModel.CreateIndexes(UserCollectionName, true, "nic_name", "phone_number")
	AppModel.CreateCompoundIndex(UserCollectionName, true, "nic_name", "phone_number")

	// Store Collection
	AppModel.CreateIndexes(StoreCollectionName, true, "store_phone")

	// MenuCollection
	AppModel.CreateIndexes(MenuCollectionName, false, "store_id", "name")
	AppModel.CreateCompoundIndex(MenuCollectionName, true, "store_id", "name")

	// Receipt Collection
	AppModel.CreateIndexes(ReceiptCollectionName, false, "user_id")

	// Review Collection
	AppModel.CreateIndexes(ReviewCollectionName, false, "store_id", "user_id")
}
