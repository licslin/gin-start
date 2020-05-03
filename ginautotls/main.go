package main

import "github.com/gin-gonic/gin"
import "github.com/gin-gonic/autotls"

/**
自动证书配置
*/
func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello licslan")
	})

	//自动下载证书
	autotls.Run(r, "www.xxx.com")
	//r.Run()

}
