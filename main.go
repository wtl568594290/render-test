package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	router.Run() // 默认监听 0.0.0.0:8080
}
