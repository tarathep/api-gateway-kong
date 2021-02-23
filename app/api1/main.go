package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/api1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "api1-response",
		})
	})
	r.Run(":2001")
}
