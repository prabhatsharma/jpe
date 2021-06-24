package v1alpha1

import (
	v1 "k8s.io/api/admission/v1"
)

func ValidateResource(aReviewRequest *v1.AdmissionReview) RuleResponse {
	var rr RuleResponse
	rr.Allowed = true // By default we pass everything

	for _, admissionPolicy := range AdmissionPolicies {
		for _, rule := range admissionPolicy.Spec.Rules {
			// fmt.Println(rule.Rule)
			rr = ValidateRule(&rule, aReviewRequest)
			if !rr.Allowed {
				rr.Message = "Policy/Rule: " + admissionPolicy.Name + "/" + rr.Message
				logger.LogStuff("Deny rule is", rr)
				return rr // return immediately if any rule fails
			}
		}
	}

	rr.Message = "All rules in all policies passed. "
	return rr // succeeded
}
