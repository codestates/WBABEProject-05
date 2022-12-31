package error

import "errors"

type ApiError struct {
	Code int
	MSG  string
	Name string
}

func (e *ApiError) Error() string {
	return e.MSG
}

func (e *ApiError) New() error {
	return errors.New(e.MSG)
}

func NewApiError(e error) ApiError {
	return NewApiErrorAndCode(UnKnownCode, e.Error())
}

func NewApiErrorAndCode(c int, e string) ApiError {
	return NewApiErrorCustom(c, e, "error")
}

func NewApiErrorCustom(c int, e string, name string) ApiError {
	return ApiError{
		Code: c,
		MSG:  e,
		Name: name,
	}
}

/* 401 ~ 499 error status */
var (
	UnauthorizedError         = NewApiErrorCustom(UnauthorizedCode, "허가되지 않은 사용자 입니다.", Unauthorized)
	RestAccessFailError       = NewApiErrorCustom(RestAccessFailCode, "로그인이 필요합니다.", RestAccessFail)
	BadRequestError           = NewApiErrorCustom(BadRequestCode, "부적절한 요청입니다.", BadRequest)
	DataNotFoundError         = NewApiErrorCustom(DataNotFoundCode, "데이터를 찾을 수 없습니다.", DataNotFound)
	DuplicatedDataError       = NewApiErrorCustom(DuplicatedDataCode, "이미 존재하는 데이터 입니다.", DuplicatedData)
	AlreadyReceivedOrderError = NewApiErrorCustom(DoesNotModifyOrderCode, "주문대기 상태에만 수정 가능 합니다.", DoesNotModifyOrder)
	BadAccessOrderError       = NewApiErrorCustom(DoesNotModifyOrderCode, "자신의 주문만 수정 가능 합니다.", DoesNotModifyOrder)
)

/* 501 ~ 599 서버 에러 */
var (
	NonInjectedError    = NewApiErrorCustom(InternalServerErrCode, "의존성 주입이 이루어지지 않았습니다.", NonInjected)
	InternalServerError = NewApiErrorCustom(SystemErrCode, "서버 로직을 수행하지 못했습니다.", InternalServerErr)
	FailRequestError    = NewApiErrorCustom(UnKnownCode, "요청을 정상적으로 처리하지 못했습니다.", FailRequestErr)
	UnKnownError        = NewApiErrorCustom(NonInjectedCode, "알 수 없는 오류입니다.", UnKnown)
)
