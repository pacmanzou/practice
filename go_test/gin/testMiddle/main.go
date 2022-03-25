package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func middle() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法前")
		c.Next()
		fmt.Println("我在方法后")
	}
}

func main() {
	r := gin.Default()
	v1router := r.Group("v1").Use(middle())
	v1router.GET("test", func(c *gin.Context) {
		fmt.Println("我在分组方法内部")
		c.JSON(200, gin.H{
			"success": true,
		})
	})
	r.Run(":8080")
}
