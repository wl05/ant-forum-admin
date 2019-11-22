package user

import (
	. "ant-forum/handler/v1"
	"ant-forum/model"
	"ant-forum/pkg/auth"
	"ant-forum/pkg/errno"
	"ant-forum/pkg/token"
	"ant-forum/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// @Summary Login generates the authentication token
// @Description 登录
// @Tags login
// @Accept  json
// @Produce  json
// @Param user body user.LoginRequest true "login"
// @Success 200 {object} user.LoginResponse "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ"}}"
// @Router /v1/login [post]
func Login(c *gin.Context) {
	log.Info("User login function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	// Binding the data with the user struct.
	var u LoginRequest
	var user *model.UserModel
	if err := c.Bind(&u); err != nil {
		fmt.Println("err", err)
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// Get the user information by the login username.
	d, err := user.GetUserByName(u.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	// Compare the login password with the user password.
	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// Sign the json web token.
	t, err := token.Sign(c, token.Context{ID: d.Id, Username: d.Username}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, nil, model.Token{Token: t})
}
