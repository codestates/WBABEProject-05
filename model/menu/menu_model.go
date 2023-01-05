package menu

import "go.mongodb.org/mongo-driver/mongo"

var instance *menuModel

type menuModel struct {
	collection *mongo.Collection
}

func NewMenuModel(col *mongo.Collection) *menuModel {
	if instance != nil {
		return instance
	}
	instance = &menuModel{
		collection: col,
	}
	return instance
}
