package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 单个文件
	r.POST("/upload1", func(c *gin.Context) {
		file, _  := c.FormFile("")
		c.SaveUploadFile(file, dst)
		c.String(http.StatusOK, "%s upload!", file.Filename)
	})

	// 多个文件
	r.POST("/upload2", func(c *gin.Context) {
		// MultiPart Form
		form, _ := c.MultiPartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			// c.SaveUploadFile(file, dst)
		}
		c.String(http.StatusOK, "%d is uploaded!", len(files))
	})

	return
}