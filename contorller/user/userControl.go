package user

import (
	error2 "github.com/codestates/WBABEProject-05/common/error"
	"github.com/codestates/WBABEProject-05/protocol"
	"github.com/codestates/WBABEProject-05/protocol/request"
	"github.com/codestates/WBABEProject-05/service/login"
	"github.com/codestates/WBABEProject-05/service/validator"
	"github.com/gin-gonic/gin"
	"net/http"
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

// GetUser godoc
// @Tags 사용자정보
// @Summary call Get user, return user by json.
// @Description 사용자 정보를 보여준다.
// @name GetUser
// @Accept  json
// @Produce  json
// @Router /app/v1/users/user [get]
// @Param user-id query string true "user-id"
// @Success 200 {object} protocol.ApiResponse[any]
func (u *userControl) GetUser(c *gin.Context) {
	userID := c.Query("user-id")
	if err := validator.CheckBlank(userID); err != nil {
		protocol.Fail(error2.NewAppError(err)).Response(c)
		return
	}

	user, err := u.userService.FindUser(userID)
	if err != nil {
		protocol.Fail(error2.DataNotFoundError).Response(c)
		return
	}

	protocol.SuccessData(user).Response(c)
}

// PutUser godoc
// @Tags 사용자정보
// @Summary call Put user, return updated count by json.
// @Description 사용자 정보를 수정 할 수 있다.
// @name PutUser
// @Accept  json
// @Produce  json
// @Router /app/v1/users/user [put]
// @Param user-id query string true "user-id"
// @Param user body request.RequestUser true "RequestUser JSON"
// @Success 200 {object} protocol.ApiResponse[any]
func (u *userControl) PutUser(c *gin.Context) {
	reqU := &request.RequestPutUser{}
	if err := c.ShouldBindJSON(reqU); err != nil {
		protocol.Fail(error2.BadRequestError).Response(c)
		return
	}

	userID := c.Query("user-id")
	if err := validator.CheckBlank(userID); err != nil {
		protocol.Fail(error2.NewAppError(err)).Response(c)
		return
	}

	cnt, err := u.userService.ModifyUser(userID, reqU)
	if err != nil {
		protocol.Fail(error2.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessData(gin.H{
		"updated_count": cnt,
	}).Response(c)
}

// DeleteUser godoc
// @Tags 사용자정보
// @Summary call Delete user, return delete count by json.
// @Description 사용자 정보를 삭제 할 수 있다.
// @name DeleteUser
// @Accept  json
// @Produce  json
// @Router /app/v1/users/user [delete]
// @Param user-id query string true "user-id"
// @Success 200 {object} protocol.ApiResponse[any]
func (u *userControl) DeleteUser(c *gin.Context) {
	userID := c.Query("user-id")
	if err := validator.CheckBlank(userID); err != nil {
		protocol.Fail(error2.NewAppError(err)).Response(c)
		return
	}

	cnt, err := u.userService.DeleteUser(userID)
	if err != nil {
		protocol.Fail(error2.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessData(gin.H{
		"deleted_count": cnt,
	}).Response(c)
}

// PostUser godoc
// @Tags 사용자정보
// @Summary call Post user, return saved id by json.
// @Description 회원가입을 할 수 있다.
// @name PostUser
// @Accept  json
// @Produce  json
// @Router /app/v1/users/user [post]
// @Param menu body request.RequestUser true "RequestUser JSON"
// @Success 201 {object} protocol.ApiResponse[any]
func (u *userControl) PostUser(c *gin.Context) {
	reqU := &request.RequestUser{}
	if err := c.ShouldBindJSON(reqU); err != nil {
		protocol.Fail(error2.BadRequestError).Response(c)
		return
	}

	savedID, err := u.userService.RegisterUser(reqU)
	if err != nil {
		protocol.Fail(error2.NewAppError(err)).Response(c)
		return
	}

	protocol.SuccessCodeAndData(http.StatusCreated, gin.H{"saved_id": savedID}).Response(c)
}
