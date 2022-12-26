package info

import (
	"github.com/codestates/WBABEProject-05/common/flag"
	info3 "github.com/codestates/WBABEProject-05/config/info"
	"github.com/codestates/WBABEProject-05/protocol"
	"github.com/gin-gonic/gin"
)

var instance *infoControl

type infoControl struct {
}

func NewInfoControl() *infoControl {
	if instance != nil {
		return instance
	}
	instance = &infoControl{}
	return instance
}

// GetInformation godoc
// @Summary call App Information, return Info by json.
// @Description App 에 대해 간략적인 정보를(소개) 제공해 준다.
// @name GetInformation
// @Accept  json
// @Produce  json
// @Router /home/info [get]
// @Success 200 {object} Info
func (h *infoControl) GetInformation(c *gin.Context) {
	path := flag.Flags[flag.InformationFlag.Name]
	info := info3.NewInfo(*path)

	protocol.SuccessData(info).Response(c)
}
