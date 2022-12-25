package user

import (
	error2 "github.com/codestates/WBABEProject-05/common/error"
	"github.com/codestates/WBABEProject-05/protocol"
	"github.com/codestates/WBABEProject-05/service/login"
	"github.com/gin-gonic/gin"
)

var instance *userControl

type userControl struct {
	userService login.UserServicer
}

func NewUserControl(svc login.UserServicer) *userControl {
	if instance != nil {
		return instance
	}
	instance = &userControl{
		userService: svc,
	}
	return instance
}

func (u *userControl) GetUser(c *gin.Context) {

}
func (u *userControl) PutUser(c *gin.Context) {

}
func (u *userControl) DeleteUser(c *gin.Context) {

}
func (u *userControl) PostUser(c *gin.Context) {
	reqU := &protocol.RequestPostUser{}
	err := c.ShouldBindJSON(reqU)
	if err != nil {
		protocol.Fail(error2.BadRequestError).Response(c)
		return
	}
	savedId, err := u.userService.RegisterUser(reqU)
	if err != nil {
		protocol.Fail(error2.NewError(err)).Response(c)
		return
	}
	protocol.SuccessData(gin.H{
		"user_id": savedId,
	}).Response(c)
}
