package main

import "github.com/gin-gonic/gin"

/**
模板渲染  GO template 使用参数官网
*/
func main() {
	r := gin.Default()
	//加载模板
	r.LoadHTMLGlob("template/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "test.html", gin.H{
			"licslan": "LICSLAN",
		})
	})
	r.Run()

}
