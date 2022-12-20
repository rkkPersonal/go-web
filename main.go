package main

import (
	_ "go-web/init" // import 下划线（如：import hello/imp）的作用：当导入一个包时，该包下的文件里所有init()函数都会被执行，然而，有些时候我们并不需要把整个包都导入进来，仅仅是是希望它执行init()函数而已。这个时候就可以使用 import 引用该包。即使用【import _ 包路径】只是引用该包，仅仅是为了调用init()函数，所以无法通过包名来调用包中的其他函数。 示例：
	_ "go-web/web"
)

/**
https://github.com/zhongzhh8/Notes/blob/master/Gin%20-%20%E9%AB%98%E6%80%A7%E8%83%BD%20Golang%20Web%20%E6%A1%86%E6%9E%B6%E7%9A%84%E4%BB%8B%E7%BB%8D%E5%92%8C%E4%BD%BF%E7%94%A8.md
*/
func main() {
	//web.StartServer03(":8088")
}
