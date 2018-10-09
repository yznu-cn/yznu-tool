package controllers

import (
	"fmt"

	"github.com/yznu-cn/yznu-tool/yznu-go/auth"
	"github.com/yznu-cn/yznu-tool/yznu-go/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	fmt.Println("login--->")
	var userInfo models.UserInfo
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		handleErr(c, 400, err.Error())
		return
	}
	code := c.GetString("code")
	token, err := auth.Login(code, &userInfo)
	if err != nil {
		handleErr(c, 400, err.Error())
		return
	}
	handleOk(c, map[string]string{"token": token})
}
