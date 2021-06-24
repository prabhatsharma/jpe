package v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

var (
	scheme1  = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
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

	policies, _ := json.Marshal(AdmissionPolicies)

	fmt.Println("Existing admissionpolicies: ", string(policies))

}

func getAdmissionPolicies(apList *AdmissionPolicyList) error {
	cl, err := client.New(config.GetConfigOrDie(), client.Options{Scheme: scheme1})
	if err != nil {
		fmt.Println("failed to create client")
		os.Exit(1)
	}

	err = cl.List(context.Background(), apList)

	if err != nil {
		fmt.Printf("\nfailed to list admissionpolicies in namespace default: %v\n", err)
		return err
	}

	return nil
}
