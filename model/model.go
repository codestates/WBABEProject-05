package model

import (
	"github.com/Hooneats/go-gin-pr4/util"
	"github.com/codestates/WBABEProject-05/common/flag"
	"github.com/codestates/WBABEProject-05/config/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var instance *Model

// model Collection 은 Store , Customer , Receipt , Review
type Model struct {
	client     *mongo.Client
	collection map[string]*mongo.Collection
}

func LoadModel() (*Model, error) {
	dbOpt := GetDbConfig()
	md, err := GetModel(dbOpt)
	if err != nil {
		return nil, err
	}
	return md, nil
}

func GetModel(dbOpt *db.Database) (*Model, error) {
	if instance != nil {
		return instance, nil
	}
	ctx, cancel := util.GetContext(util.ModelTimeOut)
	defer cancel()

	m := &Model{}
	opt := options.Client().SetMaxPoolSize(100).SetTimeout(util.DatabaseTimeOut)

	client, err := mongo.Connect(ctx, opt.ApplyURI(dbOpt.MongoUri))
	if err != nil {
		return nil, err
	}
	m.client = client

	err = m.client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	instance = m
	return instance, nil
}

func GetDbConfig() *db.Database {
	path := flag.Flags[flag.DatabaseFlag.Name]
	dbOpt := db.NewDbConfig(*path)
	return dbOpt
}

func (m *Model) GetCollection(collection string, dbName string) *mongo.Collection {
	if m.collection == nil {
		m.collection = make(map[string]*mongo.Collection)
	}
	if col, exists := m.collection[collection]; exists {
		return col
	}

	db := m.client.Database(dbName)
	col := db.Collection(collection)
	m.collection[collection] = col
	return col
}

func (m *Model) CreateIndex(colName, dbName string, indexName ...string) {
	ctx, cancel := util.GetContext(util.ModelTimeOut)
	defer cancel()

	col := m.GetCollection(colName, dbName)
	var indexModels []mongo.IndexModel
	for _, name := range indexName {
		idxModel := mongo.IndexModel{
			Keys: bson.M{name: 1}, Options: options.Index().SetUnique(true),
		}
		indexModels = append(indexModels, idxModel)
	}
	_, err := col.Indexes().CreateMany(ctx, indexModels)
	if err != nil {
		log.Println(err)
		return
	}
}
