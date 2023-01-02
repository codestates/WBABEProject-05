package validator

import (
	error2 "github.com/codestates/WBABEProject-05/protocol/error"
	"strings"
)

func EmtyString(s string) error {
	STR := strings.Trim(s, " ")
	if STR == "" {
		return error2.BadRequestError.New()
	}
	return nil
}
