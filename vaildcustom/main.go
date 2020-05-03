package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"time"
)

/**
自定义验证请求参数  have some problems with validator/v10 TODO
多语言支持
*/
func main() {
	r := gin.Default()

	//FIXME
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	r.GET("/bookable", func(c *gin.Context) {
		var book Booking
		if err := c.ShouldBind(&book); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "ok!", "booking": book})

	})
	r.Run()

}

//gt  validator 验证  go  请查看官方文档
type Booking struct {
	CheckIn time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	//CheckIn time.Time `form:"check_in" binding:"required" time_format:"2006-01-02"`
	//CheckOut time.Time `form:"check_out" binding:"required,gitfield=CheckIn" time_format:"2006-01-02"`
	//自定义实现验证器
	CheckOut time.Time `form:"check_out" binding:"required,gitfield=CheckIn" time_format:"2006-01-02"`
}

func bookableDate(
	v *validator.Validate,
	topStruct reflect.Value,
	currentStrctOrField reflect.Value,
	field reflect.Value,
	fieldType reflect.Type,
	fieldKind reflect.Kind,
	param string) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if date.Unix() > today.Unix() {
			return false
		}
	}
	return true

}
