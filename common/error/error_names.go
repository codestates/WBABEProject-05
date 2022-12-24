package error

/* 4001 ~ 4999 error status */
const (
	Unauthorized   = "인증 오류"
	RestAccessFail = "로그인 필요"
	BadRequest     = "부적절한 요청 오류"
	DataNotFound   = "존재하지 않는 데이터"
	InternalLogin  = "서버 로직 오류"
	System         = "시스템 오류"
	UnKnown        = "알 수 없는 오류"

	NonInjected = "의존성 주입 오류"
)
