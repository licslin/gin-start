package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

/**
gin中间件使用
*/
func main() {

	//gin.Default()
	//底层使用了2个中间件
	//engine := New()
	//engine.Use(Logger(), Recovery())

	//将日志输入到指定的文件里面
	f, _ := os.Create("licslan.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	r := gin.New()
	//基于控制台输出日志  打印请求路径和状态码等信息
	//r.Use(gin.Logger(),gin.Recovery())
	r.Use(gin.Logger())
	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name", "licslan")
		//panic("testing......")
		c.String(200, "%s", name)
	})
	r.Run()

}
