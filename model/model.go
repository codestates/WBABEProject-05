package model

import (
	"github.com/Hooneats/go-gin-pr4/util"
	"github.com/codestates/WBABEProject-05/common"
	"github.com/codestates/WBABEProject-05/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoModel Modeler
var MongoCollection map[string]*mongo.Collection

// model Collection 은 Store , Customer , Receipt , Review
type model struct {
	client     *mongo.Client
	collection map[string]*mongo.Collection
}

var mongoM *model

func LoadMongoModel(uri string) error {
	m := newModel()
	err := m.Connect(uri)
	m.checkClient()
	if err != nil {
		logger.AppLog.Error(err)
		return err
	} else {
		m.exposeModel()
	}

	return nil
}

func (m *model) Connect(uri string) error {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	opt := options.Client().SetMaxPoolSize(100).SetTimeout(common.DatabaseClientTimeOut)
	client, err := mongo.Connect(ctx, opt.ApplyURI(uri))
	if err != nil {
		logger.AppLog.Error(err)
		return err
	} else {
		mongoM.client = client
	}
	return nil
}

func (m *model) CreateIndex(colName string, indexName ...string) {
	ctx, cancel := util.GetContext(util.ModelTimeOut)
	defer cancel()

	var indexModels []mongo.IndexModel
	for _, name := range indexName {
		idxModel := mongo.IndexModel{
			Keys: bson.M{name: 1}, Options: options.Index().SetUnique(true),
		}
		indexModels = append(indexModels, idxModel)
	}
	_, err := MongoCollection[colName].Indexes().CreateMany(ctx, indexModels)
	if err != nil {
		logger.AppLog.Error(err)
		return
	}
}

// LoadMongoCollections 이미 (들어간)로드된 collection 은 넣지 않음
func LoadMongoCollections(colNames []string, dbName string) {
	for _, name := range colNames {
		PutCollection(name, dbName)
	}
}

// PutCollection 이미 (들어간)로드된 collection 은 넣지 않음
func PutCollection(collection, dbName string) {
	if MongoCollection == nil {
		MongoCollection = make(map[string]*mongo.Collection)
	}

	if _, exists := MongoCollection[collection]; exists {
		return
	}

	db := mongoM.client.Database(dbName)
	col := db.Collection(collection)
	MongoCollection[collection] = col
}

func (m *model) exposeModel() {
	MongoModel = mongoM
}

func newModel() *model {
	mongoM = &model{}
	return mongoM
}

func (m *model) checkClient() error {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	err := m.client.Ping(ctx, nil)
	if err != nil {
		logger.AppLog.Error(err)
		return err
	}
	return nil
}
