package main

import "github.com/gin-gonic/gin"

/**
url 作为参数使用
*/
func main() {

	r:=gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) {
		 c.JSON(200,gin.H{
		 	"name":c.Param("name"),
		 	"id":c.Param("id"),
		 })
	})
	r.Run()

	//restful style
	//http://localhost:8080/licslan/100
	//{"id":"100","name":"licslan"}
	
}
