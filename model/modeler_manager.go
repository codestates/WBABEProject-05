package model

import (
	"github.com/codestates/WBABEProject-05/model/menu"
	"github.com/codestates/WBABEProject-05/model/receipt"
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
	//TODO ReviewModel

	store.StoreModel = store.NewStoreModel(m[StoreCollectionName])
	user.UserModel = user.NewUserModel(m[UserCollectionName])

	menu.MenuModel = menu.NewMenuModel(m[MenuCollectionName])
}

func CreateIndexesInModels() {
	AppModel.CreateIndexes(UserCollectionName, "nic_name", "phone_number")
	AppModel.CreateCompoundIndex(UserCollectionName, "nic_name", "phone_number")

	AppModel.CreateIndexes(StoreCollectionName, "store_phone")

	AppModel.CreateCompoundIndex(MenuCollectionName, "store_id", "name")
}
