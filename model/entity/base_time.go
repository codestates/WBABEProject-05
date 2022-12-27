package entity

import "time"

type BaseTime struct {
	//GO 명명규칙에 따라 Camel case로 변경하는 것이 좋을 것 같습니다.
	Created_at time.Time `bson:"created_at"`
	Updated_at time.Time `bson:"updated_at"`
}
