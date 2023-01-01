package dom

type Address struct {
	Street  string `bson:"street,omitempty"`
	Detail  string `bson:"detail,omitempty"`
	ZipCode string `bson:"zip_code,omitempty"`
}
