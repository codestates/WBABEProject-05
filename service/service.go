package service

import (
	error2 "github.com/codestates/WBABEProject-05/common/error"
)

var instance *service

type service struct {
	orderReceiptService OrderReceiptServicer
	menuReviewService   MenuReviewServicer
	storeMenuService    StoreMenuServicer
}

func GetService(
	oSvc OrderReceiptServicer,
	rSvc MenuReviewServicer,
	sSvc StoreMenuServicer,
) *service {
	if instance != nil {
		return instance
	}
	instance := &service{
		orderReceiptService: oSvc,
		menuReviewService:   rSvc,
		storeMenuService:    sSvc,
	}
	return instance
}

func (s *service) OrderReceiptServicer() (OrderReceiptServicer, error) {
	if s.orderReceiptService != nil {
		return s.orderReceiptService, nil
	}
	// TODO logger
	return nil, error2.NonInjectedError.Err
}

func (s *service) MenuReviewServicer() (MenuReviewServicer, error) {
	if s.menuReviewService != nil {
		return s.menuReviewService, nil
	}
	// TODO logger
	return nil, error2.NonInjectedError.Err
}

func (s *service) StoreMenuServicer() (StoreMenuServicer, error) {
	if s.storeMenuService != nil {
		return s.storeMenuService, nil
	}
	// TODO logger
	return nil, error2.NonInjectedError.Err
}
