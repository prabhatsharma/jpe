package vapi

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var AdmissionPolicies []AdmissionPolicy

func Startup() {
	LoadPolicyFromCustomResources()

}

// LoadPolicyFromCustomResources loads policie from Custom resources that were created before the controller started.
func LoadPolicyFromCustomResources() {
	KUBERNETES_SERVICE_HOST := os.Getenv("KUBERNETES_SERVICE_HOST")

	fmt.Println(KUBERNETES_SERVICE_HOST)

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	

}
