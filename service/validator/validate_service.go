package validator

import (
	"github.com/codestates/WBABEProject-05/common/enum"
	error2 "github.com/codestates/WBABEProject-05/common/error"
	"strings"
)

func IsBlank(STR string) error {
	s := strings.Trim(STR, " ")
	if s == enum.BlankSTR {
		return error2.BadRequestError.New()
	}
	return nil
}
