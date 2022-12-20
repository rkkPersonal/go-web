package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartServer03(port string) {
	router := gin.Default()

	// 注册一个动态路由
	// 可以匹配 /user/joy
	// 不能匹配 /user 和 /user/
	router.GET("/user/:name", func(c *gin.Context) {
		// 使用 c.Param(key) 获取 url 参数
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 注册一个高级的动态路由
	// 该路由会匹配 /user/john/ 和 /user/john/send
	// 如果没有任何路由匹配到 /user/john, 那么他就会重定向到 /user/john/，从而被该方法匹配到
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run(port)
}
