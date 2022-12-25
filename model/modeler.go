package model

var AppModel Modeler

type Modeler interface {
	Connect(uri string) error
	CreateIndex(colName string, indexName ...string)
}

func SetModeler(modeler Modeler) {
	AppModel = modeler
}
