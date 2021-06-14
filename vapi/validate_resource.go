package vapi

import (
	"encoding/json"
	"fmt"

	v1 "k8s.io/api/admission/v1"
)

func ValidateResource(aReviewRequest *v1.AdmissionReview) {
	var aResource map[string]interface{}

	e, err1 := json.Marshal(aReviewRequest)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println("Validating resource:", string(e))

	aJSON, _ := aReviewRequest.Request.Object.MarshalJSON()

	err := json.Unmarshal(aJSON, &aResource)

	if err != nil {
		fmt.Println(err.Error())
	}
}
