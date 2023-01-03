package user

import "github.com/gin-gonic/gin"

var UserControl UserController

type UserController interface {
	GetUser(c *gin.Context)

	// PutUser 초기버전: 비밀번호 업데이트는 기능 x
	PutUser(c *gin.Context)

	DeleteUser(c *gin.Context)

	PostUser(c *gin.Context)
}
