package error

import "errors"

type Error struct {
	Code int
	Err  error
	Name string
}

func NewError(e error) Error {
	return Error{
		Code: BadRequestCode,
		Err:  e,
		Name: "error",
	}
}

func NewErrorAndCode(e error, c int) Error {
	return Error{
		Code: c,
		Err:  e,
		Name: "error",
	}
}

func NewErrorCustom(e error, c int, name string) Error {
	return Error{
		Code: c,
		Err:  e,
		Name: name,
	}
}

/* 401 ~ 499 error status */
var (
	UnauthorizedError   = NewErrorCustom(errors.New("허가되지 않은 사용자 입니다."), UnauthorizedCode, Unauthorized)
	RestAccessFailError = NewErrorCustom(errors.New("로그인이 필요합니다."), RestAccessFailCode, RestAccessFail)
	BadRequestError     = NewErrorCustom(errors.New("부적절한 요청입니다."), BadRequestCode, BadRequest)
	DataNotFoundError   = NewErrorCustom(errors.New("데이터를 찾을 수 없습니다."), DataNotFoundCode, DataNotFound)
)

/* 501 ~ 599 서버 에러 */
var (
	NonInjectedError    = NewErrorCustom(errors.New("의존성 주입이 이루어지지 않았습니다."), InternalServerErrCode, NonInjected)
	InternalServerError = NewErrorCustom(errors.New("서버 로직을 수행하지 못했습니다."), SystemErrCode, InternalServerErr)
	SystemError         = NewErrorCustom(errors.New("요청을 정상적으로 처리하지 못했습니다."), UnKnownCode, SystemErr)
	UnKnownError        = NewErrorCustom(errors.New("알 수 없는 오류입니다."), NonInjectedCode, UnKnown)
)
