package http_method

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// 定义一个组前缀
	// /v1/login 就会匹配到这个组
	v1 := router.Group("/api")
	{
		v1.GET("/register", register)
	}
	router.Run()
}
func register(ctx *gin.Context) {
	username := ctx.DefaultQuery("username", "default_name")
	password := ctx.Query("password")
	ctx.String(http.StatusOK, "your username is %s \npassword is %s", username, password) //http.StatusOK return code 200
}
