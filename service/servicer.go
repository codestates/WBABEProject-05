package service

var Service Servicer

type Servicer interface {
	ValidateStruct(s interface{}) error
}
