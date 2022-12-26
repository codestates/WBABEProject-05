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

// PostUser godoc
// @Summary call Post user, return saved id by json.
// @Description 회원가입을 할 수 있다.
// @name PostUser
// @Accept  json
// @Produce  json
// @Router /app/v1/users/join [post]
// @Param menu body protocol.RequestPostUser true "RequestPostUser JSON"
// @Success 200 {object} map[string]string
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
