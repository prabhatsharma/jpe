package v1alpha1

import (
	"encoding/json"
	"fmt"

	v1 "k8s.io/api/admission/v1"
)

// UpdatePolicy loads policy immediately into memory when a new CR is created
func UpdatePolicy(aReview *v1.AdmissionReview) RuleResponse {
	var rr RuleResponse
	var aPolicy AdmissionPolicy

	aJSON, _ := aReview.Request.Object.MarshalJSON()

	err := json.Unmarshal(aJSON, &aPolicy)

	if err != nil {
		fmt.Println(err.Error())
	}

	for index, policy := range AdmissionPolicies {
		if aPolicy.Name == policy.Name {
			AdmissionPolicies[index] = aPolicy // Update the policy
			continue
		}
	}

	rr.Allowed = true
	rr.Message = "Policy Updated"
	rr.Status = "Success"

	PrintPolicies()

	return rr

}
