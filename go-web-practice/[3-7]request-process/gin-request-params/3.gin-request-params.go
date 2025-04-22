package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200,
			gin.H{
				"message": "pong",
			})
	})

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "who are you ?")
	})

	// 解析路径参数
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello, %s", name)
	})

	// 解析query参数 http://localhost:8081/user?name=huaqiao&role=student
	router.GET("/user", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "student")
		c.String(http.StatusOK, "%s is a %s", name, role)
	})

	// 解析Post参数，curl http://localhost:8081/form  -X POST -d 'username=geektutu&password=1234'
	router.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// Query和Post混合参数 curl "http://localhost:8081/posts?id=9876&page=7"  -X POST -d 'username=geektutu&password=1234'
	router.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")

		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})

	router.Run(":8081")
}
