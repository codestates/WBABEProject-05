package entity

import "time"

type BaseTime struct {
	Created_at time.Time `bson:"created_at"`
	Updated_at time.Time `bson:"updated_at"`
}
