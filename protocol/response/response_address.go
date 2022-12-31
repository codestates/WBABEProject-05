package response

type ResponseAddress struct {
	Street  string `json:"street" validate:"required"`
	Detail  string `json:"detail" validate:"required"`
	ZipCode string `json:"zip_code" validate:"required"`
}
