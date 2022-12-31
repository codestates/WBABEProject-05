package error

type ApiError struct {
	Code int
	MSG  string
	Name string
}

func (e *ApiError) Error() string {
	return e.MSG
}

func NewError(e error) ApiError {
	return ApiError{
		Code: UnKnownCode,
		MSG:  e.Error(),
		Name: "error",
	}
}

func NewErrorAndCode(e string, c int) ApiError {
	return ApiError{
		Code: c,
		MSG:  e,
		Name: "error",
	}
}

func NewErrorCustom(e string, c int, name string) ApiError {
	return ApiError{
		Code: c,
		MSG:  e,
		Name: name,
	}
}

/* 401 ~ 499 error status */
var (
	UnauthorizedError   = NewErrorCustom("허가되지 않은 사용자 입니다.", UnauthorizedCode, Unauthorized)
	RestAccessFailError = NewErrorCustom("로그인이 필요합니다.", RestAccessFailCode, RestAccessFail)
	BadRequestError     = NewErrorCustom("부적절한 요청입니다.", BadRequestCode, BadRequest)
	DataNotFoundError   = NewErrorCustom("데이터를 찾을 수 없습니다.", DataNotFoundCode, DataNotFound)
	DuplicatedDataError = NewErrorCustom("이미 존재하는 데이터 입니다.", DuplicatedDataCode, DuplicatedData)
)

/* 501 ~ 599 서버 에러 */
var (
	NonInjectedError    = NewErrorCustom("의존성 주입이 이루어지지 않았습니다.", InternalServerErrCode, NonInjected)
	InternalServerError = NewErrorCustom("서버 로직을 수행하지 못했습니다.", SystemErrCode, InternalServerErr)
	FailRequestError    = NewErrorCustom("요청을 정상적으로 처리하지 못했습니다.", UnKnownCode, FailRequestErr)
	UnKnownError        = NewErrorCustom("알 수 없는 오류입니다.", NonInjectedCode, UnKnown)
)
