package contorller

import (
	"github.com/codestates/WBABEProject-05/logger"
	"github.com/codestates/WBABEProject-05/protocol"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinShouldBindJson[T any](c *gin.Context, t *T) *T {
	err := c.ShouldBindJSON(t)
	if err != nil {
		logger.AppLog.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return nil
	} else {
		return t
	}
}

func GinResponseToJson[T any](c *gin.Context, t *protocol.ApiResponse[T]) {
	c.Header("Content-Type", "application/json")
	c.JSON(t.Code, t)
}
