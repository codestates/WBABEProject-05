package util

import (
	"github.com/codestates/WBABEProject-05/model/enum"
	"github.com/codestates/WBABEProject-05/protocol/page"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewFilterCheckedUserRole(OBJID primitive.ObjectID, status, userRole string) (bson.D, error) {
	var filter []bson.E
	switch userRole {
	case enum.CustomerRole:
		filter = append(filter, bson.E{"customer_id", OBJID})
		//return bson.M{"customer_id": OBJID}, nil
	case enum.StoreRole:
		filter = append(filter, bson.E{"store_id", OBJID})
		//return bson.M{"store_id": OBJID}, nil
	}
	if status != "" {
		filter = append(filter, bson.E{"status", status})
	}
	return filter, nil
}

func NewToDayGteFilter() (bson.M, error) {
	KST, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		return nil, err
	}
	now := time.Now()
	startTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, KST).UTC()
	return bson.M{"base_time.created_at": bson.M{"$gte": startTime}}, nil
}

func NewSortFindOptions(sort *page.Sort, skip int, limit int) *options.FindOptions {
	opt := options.Find().SetSort(bson.M{sort.Name: sort.Direction}).SetSkip(int64(skip)).SetLimit(int64(limit))
	return opt
}
