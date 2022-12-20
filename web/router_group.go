package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartServer04() {
	router := gin.Default()

	// 定义一个组前缀
	// /v1/login 就会匹配到这个组
	v1 := router.Group("/v1")
	{
		v1.GET("/home", homepage)
		//v1.POST("/login", loginEndpoint)
		//v1.POST("/submit", submitEndpoint)
		//v1.POST("/read", readEndpoint)
	}
	router.Run()
}
func homepage(ctx *gin.Context) {
	ctx.String(http.StatusOK, "home")
}
