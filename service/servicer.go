package service

var Validator ValidateServicer

type ValidateServicer interface {
	Struct(s interface{}) error
	EmtyString(s string) error
}
