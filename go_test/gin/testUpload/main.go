package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/testUpload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		c.SaveUploadedFile(file, "./"+file.Filename)

		c.JSON(200, gin.H{
			"msg": file,
		})
		// c.File("./" + file.Filename)
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
