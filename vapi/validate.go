package vapi

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/admission/v1"
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

	var aReview v1.AdmissionReview
	c.BindJSON(&aReview)
	if strings.ToLower(aReview.Request.Kind.Kind) == "admissionpolicy" {
		LoadPolicyFromAdmissionReview(&aReview)
	} else {
		ValidateResource(&aReview)
	}

	c.JSON(200, aReview)
}

func ValidateResource(aReview *v1.AdmissionReview) {

	var aResource map[string]interface{}

	aJSON, _ := aReview.Request.Object.MarshalJSON()

	err := json.Unmarshal(aJSON, &aResource)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(aResource)

}
