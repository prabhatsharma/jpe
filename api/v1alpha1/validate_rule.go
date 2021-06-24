package v1alpha1

import (
	"encoding/json"
	"fmt"

	"github.com/dop251/goja"
	v1 "k8s.io/api/admission/v1"
)

func ValidateRule(rule *Rule, aReviewRequest *v1.AdmissionReview) RuleResponse {
	var rr RuleResponse
	rr.Allowed = true
	rr.Status = "Failure"
	result := true

	reviewObject, _ := json.Marshal(aReviewRequest.Request.Object)
	jsObject := string(reviewObject)

	// fmt.Println("Rule is: ", rule)
	// fmt.Println("JSObject is: ", jsObject)
	vm := goja.New()
	_, err := vm.RunString(rule.Rule)
	if err != nil {
		panic(err)
	}

	validate, ok := goja.AssertFunction(vm.Get("validate"))
	if !ok {
		// panic("Not a function")

		rr.Message = rule.Name + " : Invalid Rule: Not a function. Vsalidation allowed. Should be in the form function validate(resource) { return true/false;}"
		return rr
	}

	res, err := validate(goja.Undefined(), vm.ToValue(jsObject))
	if err != nil {
		// panic(err)
		rr.Message = rule.Name + " : " + err.Error()
		return rr
	}

	fmt.Println("Rule result is: ", res)

	result = res.ToBoolean()

	if result {
		rr.Allowed = true
		rr.Status = "Success"
		rr.Message = rule.Name + " : Passed : " + rule.Message
	} else {
		rr.Allowed = false
		rr.Status = "Success"
		rr.Message = rule.Name + " : Failed : " + rule.Message
	}

	return rr
}
