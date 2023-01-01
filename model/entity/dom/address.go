package dom

type Address struct {
	Street  string `bson:"street"`
	Detail  string `bson:"detail"`
	ZipCode string `bson:"zip_code"`
}
