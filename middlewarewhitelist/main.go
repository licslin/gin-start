package main

import "github.com/gin-gonic/gin"

/**
gin中间件使用  自定义中间件  白名单
*/

//自定义中间件方法
func IPAuthMiddleWare() gin.HandlerFunc {

	return func(context *gin.Context) {
		ipList := []string{
			"127.0.0.1",
		}
		flag := false
		clientIp := context.ClientIP()
		for _, host := range ipList {
			if clientIp == host {
				flag = true
				break
			}
		}
		if !flag {
			context.String(401, "%s,not in ipList", clientIp)
			context.Abort()
		}
	}
}
func main() {
	r := gin.Default()
	//自定义  r.Use(IPAuthMiddleWare())
	r.Use(IPAuthMiddleWare())
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	r.Run()
}
