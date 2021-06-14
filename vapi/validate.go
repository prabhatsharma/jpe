package vapi

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Validate(c *gin.Context) {
	vm := goja.New()
	v, err := vm.RunString("2 + 2")
	if err != nil {
		panic(err)
	}
	if num := v.Export().(int64); num != 4 {
		panic(num)
	} else {
		fmt.Println(v.Export().(int64))
	}

	var aReviewRequest v1.AdmissionReview
	c.BindJSON(&aReviewRequest)
	if strings.ToLower(aReviewRequest.Request.Kind.Kind) == "admissionpolicy" {
		LoadPolicyFromAdmissionReview(&aReviewRequest)
	} else {
		ValidateResource(&aReviewRequest)
	}

	// var aReviewResponse v1.AdmissionResponse

	status := &metav1.Status{
		Status:  "accepted",
		Reason:  "Will accept anything",
		Message: "Enjoy your life",
	}

	aReviewResponse := v1.AdmissionResponse{
		Allowed: true,
		Result:  status,
	}

	c.JSON(200, aReviewResponse)
}

func ValidateResource(aReview *v1.AdmissionReview) {

	var aResource map[string]interface{}

	aJSON, _ := aReview.Request.Object.MarshalJSON()

	err := json.Unmarshal(aJSON, &aResource)

	if err != nil {
		fmt.Println(err.Error())
	}
}
