package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartServer01(port string) {
	router := gin.Default()
	router.GET("/someGet", getting)
	//router.POST("/somePost", posting)
	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)
	router.Run(port)
}
func getting(context *gin.Context) {
	context.String(http.StatusOK, "getting")
}
