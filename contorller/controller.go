package contorller

import (
	"github.com/codestates/WBABEProject-05/contorller/info"
)

type Controller interface {
	GetInfoControl() info.InfoController
}
