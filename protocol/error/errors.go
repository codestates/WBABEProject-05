package error

import "errors"

type Error struct {
	Code int
	Err  error
	Name string
}

func NewError(e error, c int) Error {
	return Error{
		Code: c,
		Err:  e,
		Name: "error",
	}
}

func NewErrorAndName(e error, c int, name string) Error {
	return Error{
		Code: c,
		Err:  e,
		Name: name,
	}
}

/* 4001 ~ 4999 error status */
var (
	UnauthorizedError   = NewErrorAndName(errors.New("허가되지 않은 사용자 입니다."), 4001, Unauthorized)
	RestAccessFailError = NewErrorAndName(errors.New("로그인이 필요합니다."), 4002, RestAccessFail)
	BadRequestError     = NewErrorAndName(errors.New("부적절한 요청입니다."), 4003, BadRequest)
	DataNotFoundError   = NewErrorAndName(errors.New("데이터를 찾을 수 없습니다."), 4004, DataNotFound)
)

/* 9000 ~ 9999 서버 에러 */
var (
	InternalLoginError = NewErrorAndName(errors.New("서버 로직을 수행하지 못했습니다."), 9997, InternalLogin)
	SystemError        = NewErrorAndName(errors.New("요청을 정상적으로 처리하지 못했습니다."), 9998, System)
	UnKnownError       = NewErrorAndName(errors.New("알 수 없는 오류입니다."), 9999, UnKnown)
)
