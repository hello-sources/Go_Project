package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Who are you?")
	})

	// /:param表示参数
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "Hello %s", name)
	})

	// 获取Query参数
	r.GET("/tests", func(c *gin.Context) {
		name := c.Query("name")
		roles := c.DefaultQuery("role", "NBA球员")
		c.String(http.StatusOK, "%s is a %s", name, roles)
	})

	// POST
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "蔡徐坤是球星")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// Query和POST混合参数
	r.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		username := c.PostForm("username")
		password := c.DefaultPostForm("username", "000000")

		c.JSON(http.StatusOK, gin.H{
			"id": id,
			"page": page,
			"username": username,
			"password": password,
		})
	})

	// Map字典参数
	r.POST("/map", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostForm("names")

		c.JSON(http.StatusOK, gin.H{
			"ids": ids,
			"names": names,
		})
	})

	// 重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	r.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})

	// 路由分组
	defaultHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}
	v1 := r.Group("/v1")
	{
		v1.GET("/posts", defaultHandler)
		v1.GET("/series", defaultHandler)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/posts", defaultHandler)
		v2.GET("/series", defaultHandler)
	}

	r.Run(":9999")
}
