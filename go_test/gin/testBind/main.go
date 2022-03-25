package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type PostParams struct {
	Name string `json:"name"`
	Age  int    `json:"age" binding:"required,mustBig"`
	Sex  bool   `json:"sex"`
}

func mustBig(fl validator.FieldLevel) bool {
	if fl.Field().Interface().(int) <= 18 {
		return false
	}
	return true
}

func main() {
	r := gin.Default()
	r.POST("/testBind", func(c *gin.Context) {
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			v.RegisterValidation("mustBig", mustBig)
		}

		var p PostParams

		err := c.ShouldBindJSON(&p)
		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "错误了",
				"data": gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "成功了",
				"data": p,
			})
		}
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
