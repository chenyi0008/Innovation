package service

import (
	"Innovation/model"
	"Innovation/utils"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Username string
	Password string
}

func LoginService(c *gin.Context) {
	var request loginRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(500, model.NewResponse(0, err.Error(), nil))
	}

	if user, flag := model.UserLogin(request.Username, request.Password); flag {
		token := utils.GetToken(user.Account, user.ID)
		ctx := context.Background()

		if err := utils.GetRedisHelper().Set(ctx, token, utils.Marshal(user), time.Hour*24).Err(); err != nil {
			c.JSON(500, model.NewResponse(0, err.Error(), nil))
			return
		}
		c.JSON(200, model.NewResponse(1, "登录成功", "Bearer "+token))
		return
	} else {
		c.JSON(200, model.NewResponse(0, "账号或密码输入有误", nil))
		return
	}

}

type registerRequest struct {
	Username string
	Password string
}

func RegisterService(c *gin.Context) {
	var request registerRequest
	c.BindJSON(&request)

	user := &model.User{
		Account:  request.Username,
		Password: request.Password,
	}

	if len(user.Account) == 0 || len(user.Password) == 0 {
		c.JSON(200, model.NewResponse(0, "账号或密码不能为空", nil))
		return
	}

	if model.UserIfExist(request.Username) {
		c.JSON(200, model.NewResponse(0, "此用户名已被注册", nil))
		return
	}

	model.UserCreate(user)
	c.JSON(200, model.NewResponse(1, "注册成功", nil))
}
