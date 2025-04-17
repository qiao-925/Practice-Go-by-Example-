package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthRequired 身份验证中间件
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		// 在这里验证 token 的有效性
		if !isValidToken(token) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Next()
	}
}

func isValidToken(token string) bool {
	// 在这里实现 token 验证逻辑
	// 例如，查询数据库、调用外部 API 等
	return token == "valid_token"
}

/**
中间件应用范围
- 全局
- 按组
- 针对单个请求（bingo）
*/

func main() {
	r := gin.Default()
	r.GET("/protected", AuthRequired(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "访问受保护的资源"})
	})
	r.Run()
}
