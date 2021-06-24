package v1alpha1

import (
	v1 "k8s.io/api/admission/v1"
)

// DeletePolicy loads policy immediately into memory when a new CR is created
func DeletePolicy(aReview *v1.AdmissionReview) RuleResponse {
	logger.LogStuff("delete request is: ", aReview)

	// json2Logger.LogJSON("delete request is: ", aReview)
	var rr RuleResponse
	// var requestedDeletePolicy AdmissionPolicy

	// aJSON, _ := aReview.Request.Name

	// err := json.Unmarshal(aJSON, &requestedDeletePolicy)
	requestedDeletePolicy := aReview.Request.Name

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	for index, existingPolicy := range AdmissionPolicies {
		logger.LogStuff("requestedDeletePolicy and existingPolicy are: ", requestedDeletePolicy, existingPolicy)
		if requestedDeletePolicy == existingPolicy.Name {
			// Remove the policy from the central repo (memory)
			AdmissionPolicies = append(AdmissionPolicies[:index], AdmissionPolicies[index+1:]...)
			continue
		}
	}

	rr.Allowed = true
	rr.Message = "Policy Deleted"
	rr.Status = "Success"

	PrintPolicies()

	return rr

}
