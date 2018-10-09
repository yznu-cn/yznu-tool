package router

import (
	"github.com/yznu-cn/yznu-tool/yznu-go/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/hello", controllers.HelloWorld)
		apiv1.POST("/login", controllers.Login)
	}
	return r
}
