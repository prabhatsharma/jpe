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
	// scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
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

	policies, _ := json.Marshal(AdmissionPolicies)

	fmt.Println("Existing admissionpolicies: ", string(policies))

	var ap AdmissionPolicy

	getAdmissionPolicy(&ap)

	policy, _ := json.Marshal(ap)

	fmt.Println("\nSimplePolicy: ", string(policy))

}

func getAdmissionPolicies(apList *AdmissionPolicyList) error {
	cl, err := client.New(config.GetConfigOrDie(), client.Options{})
	if err != nil {
		fmt.Println("failed to create client")
		os.Exit(1)
	}

	// var api *AdmisssionPolicy

	err = cl.List(context.Background(), apList)

	if err != nil {
		fmt.Printf("\nfailed to list admissionpolicies in namespace default: %v\n", err)
		return err
	}

	return nil
}

func getAdmissionPolicy(ap *AdmissionPolicy) error {
	scheme := runtime.NewScheme()
	cl, err := client.New(config.GetConfigOrDie(), client.Options{Scheme: scheme})
	if err != nil {
		fmt.Println("failed to create client")
		os.Exit(1)
	}

	// var api *AdmisssionPolicy

	okey := client.ObjectKey{
		Name: "simplepolicy",
	}

	err = cl.Get(context.Background(), okey, ap)

	if err != nil {
		fmt.Printf("\nfailed to get admissionpolicy : %v\n", err)
		return err
	}

	return nil
}
