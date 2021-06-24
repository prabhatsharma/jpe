package v1alpha1

import (
	"encoding/json"
	"fmt"

	v1 "k8s.io/api/admission/v1"
)

// DeletePolicy loads policy immediately into memory when a new CR is created
func DeletePolicy(aReview *v1.AdmissionReview) RuleResponse {
	var rr RuleResponse
	var aPolicy AdmissionPolicy

	aJSON, _ := aReview.Request.Object.MarshalJSON()

	err := json.Unmarshal(aJSON, &aPolicy)

	if err != nil {
		fmt.Println(err.Error())
	}

	var policyIndex int

	for index, policy := range AdmissionPolicies {
		if aPolicy.Name == policy.Name {
			policyIndex = index
			continue
		}
	}

	// Remove the policy from the central repo (memory)
	AdmissionPolicies = append(AdmissionPolicies[:policyIndex], AdmissionPolicies[policyIndex+1:]...)

	rr.Allowed = true
	rr.Message = "Policy Deleted"
	rr.Status = "Success"

	PrintPolicies()

	return rr

}
