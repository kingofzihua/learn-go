package main

import "github.com/gin-gonic/gin"

// LoggerMiddleware --- 定义一些中间件 ---
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// log something
		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// auth check
		c.Next()
	}
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// cors headers
		c.Next()
	}
}
