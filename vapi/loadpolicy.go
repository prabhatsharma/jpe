package vapi

import (
	"encoding/json"
	"fmt"

	v1 "k8s.io/api/admission/v1"
)

// LoadPolicyFromCustomResources loads policie from Custom resources that were created before the controller started.
func LoadPolicyFromCustomResources() {

}

// LoadPolicyFromAdmissionReview loads policy immediately into memory when a new CR is created
func LoadPolicyFromAdmissionReview(areview *v1.AdmissionReview) {
	var aPolicy AdmissionPolicy

	aJSON, _ := areview.Request.Object.MarshalJSON()

	err := json.Unmarshal(aJSON, &aPolicy)

	if err != nil {
		fmt.Println(err.Error())
	}

	reviewObject, _ := json.Marshal(areview)
	fmt.Println("Policy Review Object is:", string(reviewObject))

	ap, _ := json.Marshal(aPolicy)
	fmt.Println("Policy is: ", string(ap))

	AdmissionPolicies = append(AdmissionPolicies, aPolicy)

}
