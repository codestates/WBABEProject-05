package model

import (
	"github.com/codestates/WBABEProject-05/common/enum"
	"github.com/codestates/WBABEProject-05/model/menu"
	"github.com/codestates/WBABEProject-05/model/receipt"
	"github.com/codestates/WBABEProject-05/model/review"
	"github.com/codestates/WBABEProject-05/model/store"
	"github.com/codestates/WBABEProject-05/model/user"
	"go.mongodb.org/mongo-driver/mongo"
)

// LoadModelImpls generate singleton modelers

func InjectModelsMongoDependency(m map[string]*mongo.Collection) {
	receipt.ReceiptModel = receipt.NewReceiptModel(m[enum.ReceiptCollectionName])
	review.ReviewModel = review.NewReviewModel(m[enum.ReviewCollectionName])
	store.StoreModel = store.NewStoreModel(m[enum.StoreCollectionName])
	user.UserModel = user.NewUserModel(m[enum.UserCollectionName])
	menu.MenuModel = menu.NewMenuModel(m[enum.MenuCollectionName])
}

func CreateIndexesInModels() {
	// User Collection
	AppModel.CreateIndexes(enum.UserCollectionName, true, "nic_name", "phone_number")
	AppModel.CreateCompoundIndex(enum.UserCollectionName, true, "nic_name", "phone_number")

	// Store Collection
	AppModel.CreateIndexes(enum.StoreCollectionName, true, "store_phone")
	AppModel.CreateCompoundIndex(enum.StoreCollectionName, false, "store_id", "user_id")

	// MenuCollection
	AppModel.CreateIndexes(enum.MenuCollectionName, false, "store_id", "name")
	AppModel.CreateCompoundIndex(enum.MenuCollectionName, true, "store_id", "name")

	// Receipt Collection
	AppModel.CreateIndexes(enum.ReceiptCollectionName, false, "user_id")

	// Review Collection
	AppModel.CreateIndexes(enum.ReviewCollectionName, false, "store_id", "user_id")
}
