package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 省略的代码 ...
func Handler(c *gin.Context) {
	// 使用 String 方法即可
	c.String(200, "Success")
}

func HandlerJson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success", "format": "H_JSON"})
}

func WriteHtml() {
	engine := gin.Default()
	engine.GET("/html-test", startPage)
	// 注册一个目录，gin 会把该目录当成一个静态的资源目录
	// 该目录下的资源看可以按照路径访问
	// 如 static 目录下有图片路径 index/logo.png , 你可以通过 GET /static/index/logo.png 访问到
	engine.Static("/static", "./static")
	// 注册一个路径，gin 加载模板的时候会从该目录查找
	// 参数是一个匹配字符，如 views/*/* 的意思是 模板目录有两层
	// gin 在启动时会自动把该目录的文件编译一次缓存，不用担心效率问题
	engine.LoadHTMLGlob("views/*/*")

	engine.Run(":8080")
}

func startPage(c *gin.Context) {
	// 输出 html
	// 三个参数，200 是http状态码
	// "shop/search" 要加载的模板路径，对应目录路径 views/shop/search.html
	// gin.H{"keywords":"macbook pro"} 是模板变量
	c.HTML(200, "shop/search", gin.H{"keywords": "macbook pro"})
}
