package user

import (
	. "ant-forum/handler/v1"
	"ant-forum/pkg/token"
	"ant-forum/model"
	"ant-forum/pkg/errno"
	"ant-forum/util"

	"fmt"
	"github.com/lexkong/log/lager"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// @Summary Get an user by the user id
// @Description Get an user by id
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200 {object} model.UserInfo "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /v1/user/{id} [get]
func GetUserById(c *gin.Context) {
	log.Info("User get function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	userId, _ := strconv.Atoi(c.Param("id"))
	// Get the user by the `id` from the database.
	fmt.Println("userId", userId)
	user, err := model.GetUserById(uint64(userId))
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	resUsr := &model.UserInfo{Id: user.Id, Username: user.Username, Avatar: user.Avatar}
	SendResponse(c, nil, resUsr)
}

// @Summary Get an user
// @Description Get an user
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} model.UserInfo "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /v1/auth/info [get]
func GetUserInfo(c *gin.Context) {
	res, _ := token.ParseRequest(c)
	user, err := model.GetUserById(uint64(res.ID))
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	resUsr := &model.UserInfo{Id: user.Id, Username: user.Username, Avatar: user.Avatar}
	SendResponse(c, nil, resUsr)
}
