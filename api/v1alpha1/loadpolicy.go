package v1alpha1

import (
	"encoding/json"
	"fmt"

	"github.com/dop251/goja"
	v1 "k8s.io/api/admission/v1"
)

// AddToExistingPolicies loads policy immediately into memory when a new CR is created
func AddToExistingPolicies(aReview *v1.AdmissionReview) RuleResponse {
	var rr RuleResponse
	var aPolicy AdmissionPolicy

	aJSON, _ := aReview.Request.Object.MarshalJSON()

	err := json.Unmarshal(aJSON, &aPolicy)

	if err != nil {
		fmt.Println(err.Error())
	}

	for _, rule := range aPolicy.Spec.Rules {
		vm := goja.New()
		_, err = vm.RunString(rule.Rule)
		if err != nil {
			rr.Allowed = false
			rr.Message = "Policy did not pass checks. Check for errors in policy"
			rr.Status = "Failure"

			return rr
		}
	}

	reviewObject, _ := json.Marshal(aReview)
	fmt.Println("Policy Review Object is:", string(reviewObject))

	ap, _ := json.Marshal(aPolicy)
	fmt.Println("Policy is: ", string(ap))

	AdmissionPolicies = append(AdmissionPolicies, aPolicy)

	rr.Allowed = true
	rr.Message = "Policy loaded"
	rr.Status = "Success"

	return rr

}
