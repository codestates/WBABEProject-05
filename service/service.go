package service

import (
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/go-playground/validator"
)

var validate = validator.New()

// service validate 등등
type service struct {
}

func (svc *service) ValidateStruct(s interface{}) error {
	if err := validate.Struct(s); err != nil {
		logger.AppLog.Error(err)
		return err
	}
	return nil
}
