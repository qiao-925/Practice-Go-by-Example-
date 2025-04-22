package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func main() {
	router := gin.Default()
	/**

	请求体上下文一次性or not
	curl --location 'http://localhost:8080/ShouldBind' \
	--header 'Content-Type: application/json' \
	--data '{
	    "foo":"111"
	}'
	*/
	router.POST("/ShouldBind", func(c *gin.Context) {
		objA := formA{}
		objB := formB{}
		// c.ShouldBind 使用了 c.Request.Body，不可重用。
		if errA := c.ShouldBind(&objA); errA == nil {
			c.String(http.StatusOK, `the body should be formA`)
			// 因为现在 c.Request.Body 是 EOF，所以这里会报错。
		} else if errB := c.ShouldBind(&objB); errB == nil {
			c.String(http.StatusOK, `the body should be formB`)
		} else {
			c.String(http.StatusOK, `the body should be formA or formB`)
		}
	})

	/**
	curl --location 'http://localhost:8080/ShouldBindWith' \
	--header 'Content-Type: application/json' \
	--data '{
	    "foo":"111"
	}'
	*/
	router.POST("/ShouldBindWith", func(c *gin.Context) {
		objA := formA{}
		objB := formB{}
		// 读取 c.Request.Body 并将结果存入上下文。
		if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
			c.String(http.StatusOK, `the body should be formA`)
			// 这时, 复用存储在上下文中的 body。
		}

		if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
			c.String(http.StatusOK, `the body should be formB JSON`)
			// 可以接受其他格式
		}

		if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
			c.String(http.StatusOK, `the body should be formB XML`)
		} else {
			c.String(http.StatusOK, `the body should be formA or formB`)
		}
	})

	router.Run()
}
