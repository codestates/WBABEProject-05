package query

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PageQuery struct {
	Name      string
	Direction int
	Skip      int
	Limit     int
}

func NewPageQuery(name string, direction, skip, limit int) *PageQuery {
	return &PageQuery{
		Name:      name,
		Direction: direction,
		Skip:      skip,
		Limit:     limit,
	}
}

func (p *PageQuery) NewSortFindOptions() *options.FindOptions {
	opt := options.Find().SetSort(bson.M{p.Name: p.Direction}).SetSkip(int64(p.Skip)).SetLimit(int64(p.Limit))
	return opt
}
