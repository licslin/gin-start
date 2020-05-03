package main

import "github.com/gin-gonic/gin"

/**
验证请求参数  结构体验证
*/
func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, "%v", err)
			return
		}
		c.String(200, "%v", person)
	})
	r.Run()

}

//gt  validator 验证  go
type Person struct {
	Age     int    `form:"age" binding:"required,gt=10"`
	Name    string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
}

//curl -X GET "http://127.0.0.1:8080/test?age=1&name=licslan&address=wuhan"
//Key: 'Person.Age' Error:Field validation for 'Age' failed on the 'gt' tag{1 licslan wuhan}
//curl -X GET "http://127.0.0.1:8080/test?age=11&name=licslan&address=wuhan"
//{11 licslan wuhan}
