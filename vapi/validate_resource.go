package vapi

import (
	"encoding/json"
	"fmt"

	"github.com/dop251/goja"
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

	// resourceType := strings.ToLower(aReviewRequest.Request.Kind.Kind)
	// fmt.Println("Resource type: ", resourceType)

	reviewObject, _ := json.Marshal(aReviewRequest.Request.Object)
	jsObject := string(reviewObject)
	// fmt.Println("Review Object is:", jsObject)

	for _, admissionPolicy := range AdmissionPolicies {
		for _, rule := range admissionPolicy.Spec.Rules {
			fmt.Println(rule.Rule)
			ValidateRule(rule.Rule, jsObject)

		}
	}

	err := json.Unmarshal(aJSON, &aResource)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func ValidateRule(rule string, jsObject string) {
	fmt.Println("Rule is: ", rule)
	fmt.Println("JSObject is: ", jsObject)
	vm := goja.New()
	_, err := vm.RunString(rule)
	if err != nil {
		panic(err)
	}

	validate, ok := goja.AssertFunction(vm.Get("validate"))
	if !ok {
		panic("Not a function")
	}

	res, err := validate(goja.Undefined(), vm.ToValue(jsObject))
	if err != nil {
		panic(err)
	}
	fmt.Println("Rule result is: ", res)
}
