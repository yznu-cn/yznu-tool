package controllers

import (
	"github.com/gin-gonic/gin"
)

func handleOk(c *gin.Context, data interface{}) {
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "success", "data": data})
}

func handleErr(c *gin.Context, code int, errMsg string) {
	c.JSON(200, map[string]interface{}{"code": code, "msg": errMsg})
}
