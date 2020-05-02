package main

import "github.com/gin-gonic/gin"

//gin 路由绑定
func main() {
	r:=gin.Default()
	//get request
	r.GET("/get", func(c *gin.Context) {
		c.String(200,"get")
	})
	//post request
	r.POST("/post", func(c *gin.Context) {
		c.String(200,"post")
	})
	//delete request
	r.Handle("DELETE","/delete", func(c *gin.Context) {
		c.String(200,"delete")
	})
	//随便什么方式请求
	r.Any("/any", func(c *gin.Context) {
		c.String(200,"any")
	})
	//开启服务器
	r.Run()
}
