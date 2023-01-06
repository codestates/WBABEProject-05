package validator

import (
	"github.com/codestates/WBABEProject-05/common/enum"
	error2 "github.com/codestates/WBABEProject-05/common/error"
	"strings"
)

// CheckBlank string trim 한 값이 "" 인 경우 BadRequestError
func CheckBlank(STR string) error {
	s := strings.Trim(STR, " ")
	if s == enum.BlankSTR {
		return error2.BadRequestError
	}
	return nil
}
