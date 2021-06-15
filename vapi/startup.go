package vapi

import (
	"fmt"
	"os"
)

var AdmissionPolicies []AdmissionPolicy

func Startup() {
	LoadPolicyFromCustomResources()

}

// LoadPolicyFromCustomResources loads policie from Custom resources that were created before the controller started.
func LoadPolicyFromCustomResources() {
	KUBERNETES_SERVICE_HOST := os.Getenv("KUBERNETES_SERVICE_HOST")

	fmt.Println(KUBERNETES_SERVICE_HOST)

}
