package info

import "github.com/gin-gonic/gin"

type InfoController interface {
	GetInformation(g *gin.Context)
}
