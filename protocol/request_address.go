package protocol

type RequestAddress struct {
	Street  string `json:"street"`
	Detail  string `json:"detail"`
	ZipCode string `json:"zip_code"`
}
