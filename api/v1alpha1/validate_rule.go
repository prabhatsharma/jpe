package v1alpha1

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/dop251/goja"
	v1 "k8s.io/api/admission/v1"
)

func ValidateRule(rule *Rule, aReviewRequest *v1.AdmissionReview) RuleResponse {
	var rr RuleResponse
	rr.Allowed = true
	rr.Status = "Failure"
	result := true

	// Check if we are evaluating for the right resource kind. e.g. If rule is for pod and resource is not pod then skip.
	requestedResourceKind := strings.ToLower(aReviewRequest.Request.Kind.Kind)
	ruleResourceKind := strings.ToLower(rule.ResourceKind)

	if requestedResourceKind != ruleResourceKind {
		rr.Status = "Success"
		return rr
	}

	reviewObject, _ := json.Marshal(aReviewRequest.Request.Object)
	jsObject := string(reviewObject)

	vm := goja.New()
	_, err := vm.RunString(rule.Rule)
	if err != nil {
		panic(err)
	}

	validate, ok := goja.AssertFunction(vm.Get("validate"))
	if !ok {
		rr.Message = rule.Name + " : Invalid Rule: Not a function. Validation allowed. Should be in the form function validate(resource) { return true/false;}"
		return rr
	}

	// Interrupt javascript in 200 milliseconds. This will avoid any long running javascript loops improving security.
	time.AfterFunc(200*time.Millisecond, func() {
		vm.Interrupt("halt")
	})

	res, err := validate(goja.Undefined(), vm.ToValue(jsObject))
	if err != nil {
		// panic(err)
		rr.Message = rule.Name + " : " + err.Error()
		return rr
	}

	result = res.ToBoolean()

	if result {
		rr.Allowed = true
		rr.Status = "Success"
		rr.Message = rule.Name + " : Passed : " + rule.Message
	} else {
		rr.Allowed = false
		rr.Status = "Success"
		rr.Message = rule.Name + " : Failed : " + rule.Message

		logger.LogStuff("Rule result is: ", rr)
	}

	return rr
}
