package http_method

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	Name string `form:"name"` //必须有`form:"name"`才能绑定
	Pwd  int    `form:"pwd"`
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api")
	{
		v1.GET("/register", register02)
	}
	router.Run(":8080")
}
func register02(ctx *gin.Context) {
	log.Println("someone register")
	var user User
	err := ctx.ShouldBindQuery(&user)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"username": user.Name,
			"password": user.Pwd,
		})
	}
}
