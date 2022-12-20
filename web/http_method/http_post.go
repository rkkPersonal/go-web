package http_method

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api")
	{
		v1.POST("/register", register03)
	}
	router.Run(":8080")
}
func register03(ctx *gin.Context) {
	var user User
	err := ctx.ShouldBind(&user)
	if err == nil {
		log.Println(user.Name)
		log.Println(user.Pwd)
		ctx.JSON(http.StatusOK, gin.H{
			"username": user.Name,
			"password": user.Pwd,
		})
	}
}
