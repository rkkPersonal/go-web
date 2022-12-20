package valid

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `form:"name" json:"name" binding:"required"`
	Pwd  int    `form:"pwd" json:"pwd" binding:"required"`
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api")
	{
		v1.POST("/registerJSON", registerJSON)
		v1.POST("/registerFORM", registerFORM)
	}
	router.Run(":8080")
}
func registerJSON(ctx *gin.Context) {
	var user_json User
	if err := ctx.ShouldBind(&user_json); err == nil {
		if user_json.Name == "zzh" && user_json.Pwd == 123 {
			ctx.JSON(http.StatusOK, gin.H{
				"status":   "register successfully",
				"username": user_json.Name,
				"password": user_json.Pwd,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	}
}
func registerFORM(ctx *gin.Context) {
	var user_form User
	if err := ctx.ShouldBind(&user_form); err == nil {
		if user_form.Name == "zzh" && user_form.Pwd == 123 {
			ctx.JSON(http.StatusOK, gin.H{
				"status":   "register successfully",
				"username": user_form.Name,
				"password": user_form.Pwd,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	}
}
