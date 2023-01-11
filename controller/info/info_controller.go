package info

import "github.com/gin-gonic/gin"

var InfoControl InfoController

type InfoController interface {
	GetInformation(c *gin.Context)
}
