package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
静态文件夹绑定
*/
func main() {

	r:=gin.Default()
	//指向文件夹
	r.Static("/staticfile","./staticfile")
	//静态文件类型
	r.StaticFS("static",http.Dir("static"))
	//图标
	r.StaticFile("/favicon,ico","./favicon.ico")
	r.Run() //8080

	//cd ginwebserver/route
	// go build -o route && ./route   [in linux env]
	
}
