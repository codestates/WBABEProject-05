package user

import (
	"fmt"
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
	fmt.Println(reqU)
	err := c.ShouldBindJSON(reqU)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reqU)
	savedId, err := u.userService.RegisterUser(reqU)
	if err != nil {
		// TODO ERR
		fmt.Println(err)
		return
	}
	protocol.SuccessData(gin.H{
		"user_id": savedId,
	}).Response(c)
}
