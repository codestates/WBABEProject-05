package model

import (
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
)

// LoadModelImpls generate singleton modelers

func InjectModelsMongoDependency(m map[string]*mongo.Collection) {
	receipt.ReceiptModel = receipt.NewReceiptModel(m[ReceiptCollectionName])
	//TODO ReviewModel

	store.StoreModel = store.NewStoreModel(m[StoreCollectionName])
	user.UserModel = user.NewUserModel(m[UserCollectionName])
}

func CreateIndexesInModels() {
	AppModel.CreateIndex(UserCollectionName, "nic_name", "phone_number")
}
