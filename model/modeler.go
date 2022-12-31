package model

var AppModel Modeler

type Modeler interface {
	Connect(uri string) error
	CreateIndexes(colName string, indexName ...string)
	CreateCompoundIndex(colName string, indexName ...string)
}

func SetModeler(modeler Modeler) {
	AppModel = modeler
}
