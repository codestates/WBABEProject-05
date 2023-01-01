package model

var AppModel Modeler

type Modeler interface {
	Connect(uri string) error
	CreateIndexes(colName string, unique bool, indexName ...string)
	CreateCompoundIndex(colName string, unique bool, indexName ...string)
}

func SetModeler(modeler Modeler) {
	AppModel = modeler
}
