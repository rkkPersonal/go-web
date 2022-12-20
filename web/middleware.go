package web

import (
	"github.com/gin-gonic/gin"
	"log"
	//"net/http"
)

func StartServer05() {
	router := gin.Default()

	// 注册一个路由，使用了 middleware1，middleware2 两个中间件
	router.GET("/someGet", middleware1, middleware2, handler)

	// 默认绑定 :8080
	router.Run()
}
func middleware1(c *gin.Context) {
	log.Println("exec middleware1")

	//你可以写一些逻辑代码

	// 执行该中间件之后的逻辑
	c.Next()
}
func middleware2(c *gin.Context) {
	log.Println("arrive at middleware2")
	// 执行该中间件之前，先跳到流程的下一个方法
	c.Next()
	// 流程中的其他逻辑已经执行完了
	log.Println("exec middleware2")

	//你可以写一些逻辑代码
}
func handler(c *gin.Context) {
	log.Println("exec handler")
}
