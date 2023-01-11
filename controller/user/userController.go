package user

import "github.com/gin-gonic/gin"

var UserControl UserController

// UserController 초기버전: 회원에 대한 요구사항은 없기에 기본적인 CRUD 만 구현
type UserController interface {
	GetUser(c *gin.Context)

	// PutUser 초기버전: 비밀번호 업데이트는 기능 x
	PutUser(c *gin.Context)

	DeleteUser(c *gin.Context)

	PostUser(c *gin.Context)
}
