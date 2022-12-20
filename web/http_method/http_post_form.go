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
		v1.POST("/register", register01)
	}
	router.Run(":8080")
}
func register01(ctx *gin.Context) {
	log.Println("someone register")
	username := ctx.DefaultPostForm("username", "default_name")
	password := ctx.PostForm("password")
	ctx.JSON(http.StatusOK, gin.H{
		"status":   "posted",
		"username": username,
		"password": password,
	})

}
