package controllers

import (
	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	handleOk(c, "Hello World")
}
