package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/api2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "api2-response",
		})
	})
	r.Run(":2002")
}
