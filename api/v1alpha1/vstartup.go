package v1alpha1

import (
	"context"
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/runtime"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

var (
	scheme1 = runtime.NewScheme()
)

var AdmissionPolicies []AdmissionPolicy

func Startup() {
	AddToScheme(scheme1)
	LoadPolicyFromCustomResources()
}

// LoadPolicyFromCustomResources loads policie from Custom resources that were created before the controller started.
func LoadPolicyFromCustomResources() {
	var apList AdmissionPolicyList

	getAdmissionPolicies(&apList)

	AdmissionPolicies = append(AdmissionPolicies, apList.Items...)

	logger.LogStuff("Existing admissionpolicies: ", AdmissionPolicies)

}

func getAdmissionPolicies(apList *AdmissionPolicyList) error {
	cl, err := client.New(config.GetConfigOrDie(), client.Options{Scheme: scheme1})
	if err != nil {
		fmt.Println("failed to create client")
		os.Exit(1)
	}

	err = cl.List(context.Background(), apList)

	if err != nil {
		logger.LogStuff("failed to list admissionpolicies in namespace default:", err)
		return err
	}

	return nil
}
