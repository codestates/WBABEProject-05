package service

import (
	"github.com/codestates/WBABEProject-05/logger"
	error2 "github.com/codestates/WBABEProject-05/protocol/error"
	"github.com/go-playground/validator"
	"strings"
)

var instance *validateService

type validateService struct {
	validate *validator.Validate
}

func NewService() *validateService {
	if instance != nil {
		return instance
	}
	instance = &validateService{
		validate: validator.New(),
	}
	return instance
}

func (svc *validateService) Struct(s interface{}) error {
	if err := svc.validate.Struct(s); err != nil {
		logger.AppLog.Error(err)
		return error2.BadRequestError.New()
	}
	return nil
}

func (svc *validateService) EmtyString(s string) error {
	STR := strings.Trim(s, " ")
	if STR == "" {
		return error2.BadRequestError.New()
	}
	return nil
}
