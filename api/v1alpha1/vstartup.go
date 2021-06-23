package v1alpha1

import (
	"context"
	"fmt"
	"os"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

var AdmissionPolicies []AdmissionPolicy

func Startup() {
	LoadPolicyFromCustomResources()
}

// LoadPolicyFromCustomResources loads policie from Custom resources that were created before the controller started.
func LoadPolicyFromCustomResources() {
	var apList AdmissionPolicyList

	getAdmissionPolicies(&apList)

	AdmissionPolicies = append(AdmissionPolicies, apList.Items...)

}

func getAdmissionPolicies(apList *AdmissionPolicyList) error {
	cl, err := client.New(config.GetConfigOrDie(), client.Options{})
	if err != nil {
		fmt.Println("failed to create client")
		os.Exit(1)
	}

	err = cl.List(context.Background(), apList)
	if err != nil {
		fmt.Printf("failed to list pods in namespace default: %v\n", err)
		return err
	}
	return nil
}
