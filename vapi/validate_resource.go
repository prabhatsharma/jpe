package vapi

import (
	"encoding/json"
	"fmt"
	"strings"

	v1 "k8s.io/api/admission/v1"
)

func ValidateResource(aReviewRequest *v1.AdmissionReview) {
	var aResource map[string]interface{}

	// e, err1 := json.Marshal(aReviewRequest)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }
	// fmt.Println("Validating resource:", string(e))

	aJSON, _ := aReviewRequest.Request.Object.MarshalJSON()

	resourceType := strings.ToLower(aReviewRequest.Request.Kind.Kind)
	fmt.Println("Resource type: ", resourceType)

	for _, admissionPolicy := range AdmissionPolicies {
		for _, rule := range admissionPolicy.Spec.Rules {
			fmt.Println(rule.Rule)
		}
	}

	err := json.Unmarshal(aJSON, &aResource)

	if err != nil {
		fmt.Println(err.Error())
	}
}
