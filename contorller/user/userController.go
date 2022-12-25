package user

import "github.com/gin-gonic/gin"

var UserControl UserController

type UserController interface {
	GetUser(c *gin.Context)
	PutUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	PostUser(c *gin.Context)
}
