package web

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func StartServer06() {
	router := gin.Default()
	// 定义一个组前缀
	// /v1/home 就会匹配到这个组
	v1 := router.Group("/v1", middleware)
	{
		v1.GET("/home", homepage01)
	}
	router.Run()
}
func homepage01(ctx *gin.Context) {
	ctx.String(http.StatusOK, "home")
	log.Println("exec home")
}
func middleware(c *gin.Context) {
	log.Println("arrive at middleware")
	// 执行该中间件之前，先跳到流程的下一个方法
	c.Next()
	// 流程中的其他逻辑已经执行完了
	log.Println("exec middleware")
	//你可以写一些逻辑代码
}
