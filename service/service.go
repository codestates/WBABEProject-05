package service

import (
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/go-playground/validator"
)

var instance *service

// service validate 등등
type service struct {
	validate *validator.Validate
}

func NewService() *service {
	if instance != nil {
		return instance
	}
	instance = &service{
		validate: validator.New(),
	}
	return instance
}

func (svc *service) ValidateStruct(s interface{}) error {
	if err := svc.validate.Struct(s); err != nil {
		logger.AppLog.Error(err)
		return err
	}
	return nil
}
