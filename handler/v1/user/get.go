package user

import (
	. "ant-forum/handler/v1"
	"ant-forum/model"
	"ant-forum/pkg/errno"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get an user by the user identifier
// @Description Get an user by username
// @Tags user
// @Accept  json
// @Produce  json
// @Param username path string true "Username"
// @Success 200 {object} model.UserModel "{"code":0,"message":"OK","data":{"username":"kong","password":"$2a$10$E0kwtmtLZbwW/bDQ8qI8e.eHPqhQOW9tvjwpyo/p05f/f4Qvr3OmS"}}"
// @Router /v1/user/{username} [get]
func GetUserByName(c *gin.Context) {
	username := c.Param("username")
	// Get the user by the `username` from the database.
	user, err := model.GetUserByName(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}

// @Summary Get an user by the user id
// @Description Get an user by id
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200 {object} model.UserInfo "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /v1/user/{id} [get]
func GetUserById(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	// Get the user by the `id` from the database.
	fmt.Println("userId",userId)
	user, err := model.GetUserById(uint64(userId))
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	SendResponse(c, nil, user)
}
