package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prabhatsharma/jpe/api/v1alpha1"
)

var (
// scheme = runtime.NewScheme()
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1alpha1.Startup()

	r.POST("/validate", v1alpha1.Validate)

	r.RunTLS(":8443", "./cert/server.crt", "./cert/server.key")
}
