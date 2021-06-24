package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prabhatsharma/jpe/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"

	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
)

var (
	scheme = runtime.NewScheme()
	// setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

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
