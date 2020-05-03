package main

import "github.com/gin-gonic/gin"

/**
database/sql 是 Go 操作数据库的标准库之一，它提供了一系列接口方法，用于访问数据库（mysql，sqllite，oralce，postgresql），
它并不会提供数据库特有的方法，那些特有的方法交给数据库驱动去实现
而通常在工作中，我们更多的是用  https://github.com/jmoiron/sqlx 包来操作数据库，sqlx 是基于标准库 sql 的扩展，
并且我们可以通过 sqlx 操作各种类型的数据，如将查询的数据转为结构体等
github 地址：
https://github.com/go-sql-driver/mysql
https://github.com/jmoiron/sqlx
*/

func main() {
	//创建gin实例
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() //8080
}

//gin的请求路由
//1.多种请求类型
//2.绑定静态文件夹  可以作为静态资源服务器使用
//3.参数作为URL  RESTFUL
//4.泛绑定

//获取请求参数
//get请求参数获取
//post请求参数获取
//body值获取
//参数绑定结构体获取

//gin补充
//优雅关停
//模板渲染
//自动证书配置(证书过期自动续约功能)
