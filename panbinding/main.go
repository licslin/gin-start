package main

import "github.com/gin-gonic/gin"

/**
泛绑定  以user开头的请求都会进来
*/
func main() {
	r:=gin.Default()
	r.GET("/user/*action", func(c *gin.Context) {
		c.String(200,"hello world")
	})
	r.Run()
}
