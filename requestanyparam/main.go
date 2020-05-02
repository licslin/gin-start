package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

/**
获取get请求参数
获取post请求参数
*/
func main() {
	r := gin.Default()


	//获取get请求参数  放开注释即可
	//http://localhost:8080/test?firstname=12
	/*r.GET("/test", func(c *gin.Context) {
		//get请求获取第一个参数
		firstName:=c.Query("firstname")
		//get请求获取第二个参数  没有默认值是licslan
		lastName:=c.DefaultQuery("lastname","licslan")
		age:=c.Query("age")
		c.String(http.StatusOK,"%s,%s,%s",firstName,lastName,age)
	})*/

	//获取body请求
	//curl -X POST 'http://127.0.0.1:8080/test' -d '{"name":"hwl","age":28}'
	r.POST("/test", func(c *gin.Context) {
		//获取body
		res,err:=ioutil.ReadAll(c.Request.Body)
		if err!=nil{
			c.String(http.StatusBadRequest,err.Error())
		}

		//回写回去
		c.Request.Body=ioutil.NopCloser(bytes.NewBuffer(res))

		//获取body中的值
		// curl -X POST 'http://127.0.0.1:8080/test' -d  'name=hwl&age=100'
		// curl -X POST 'http://127.0.0.1:8080/test' -d  'name=hwl'
		name:=c.PostForm("name")
		//如果没有的话
		age:=c.DefaultPostForm("age","28")
		c.String(http.StatusOK,"%s,%s,%s",name,age,string(res))
	})





	r.Run(":8080")
}
