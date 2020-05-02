package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

/**
获取get请求参数
获取post请求参数

公用结构体返回的回调函数
*/
func main() {
	r := gin.Default()
	//route 返回相同的回调方法 test
	r.GET("/test", test)
	r.POST("/test",test)


	//curl -X GET 'http://127.0.0.1:8080/test?name=hwl'
	//curl -X GET 'http://127.0.0.1:8080/test?name=licslan&address=wuhan'
	//curl -X GET 'http://127.0.0.1:8080/test?name=licslan&address=wuhan&birthday=2000-11-18 00:00:00'
	//person binding error:parsing time "2000-11-18" as "2006-01-02T15:04:05Z07:00": cannot parse "" as "T"

	//curl -X POST 'http://127.0.0.1:8080/test' -d 'name=licslan&address=wuhan&birthday=2000-01-12'


	//curl -H "Content-Type:application/json" -X POST "http://127.0.0.1:8080/test" -d '{"name":"LICSLAN","address":"wuhan","birthday":"1991-05-16"}'
	r.Run(":8080")
}

type Person struct {
	//由参数转换为结构体
	Name string  `form:"name"`
	Address string `form:"address"`
	Brithday time.Time `form:"birthday" time_format:"2006-01-02"`
	// 2006-01-02  在go语言里面是个特殊的时间  不要随便改 ^_^
}

func test(c *gin.Context)  {
	var person Person
	//根据请求的content-type来做不同binding操作
	if err:=c.ShouldBind(&person);err==nil{
		c.String(200,"%v",person)
	}else {
		c.String(200,"person binding error:%v",err)
	}
}

//请求参数验证  判断请求参数是否合法  结构体验证  自定义验证  升级验证  支持多语言错误信息
