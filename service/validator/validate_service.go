package validator

import (
	error2 "github.com/codestates/WBABEProject-05/common/error"
	"strings"
)

func IsBlank(STR string) error {
	s := strings.Trim(STR, " ")
	if s == "" {
		return error2.BadRequestError.New()
	}
	return nil
}
