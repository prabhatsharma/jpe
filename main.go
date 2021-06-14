package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prabhatsharma/jpe/vapi"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/validate", vapi.Validate)

	r.RunTLS(":8443", "./cert/server.crt", "./cert/server.key")
}
